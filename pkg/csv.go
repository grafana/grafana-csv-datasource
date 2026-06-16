package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type fieldSchema struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type csvOptions struct {
	Delimiter        string        `json:"delimiter"`
	Header           bool          `json:"header"`
	IgnoreUnknown    bool          `json:"ignoreUnknown"`
	Schema           []fieldSchema `json:"schema"`
	SkipRows         int           `json:"skipRows"`
	DecimalSeparator string        `json:"decimalSeparator"`
	Timezone         string        `json:"timezone"`
}

func parseCSV(opts csvOptions, regex bool, r io.Reader, logger log.Logger) ([]*data.Field, error) {
	header, rows, err := readCSV(opts, r)
	if err != nil {
		return nil, backend.DownstreamError(err)
	}

	// Check if we have any data to process
	if len(rows) == 0 {
		return nil, backend.DownstreamErrorf("no data rows found in CSV")
	}

	location, err := time.LoadLocation(opts.Timezone)
	if err != nil {
		return nil, backend.DownstreamError(err)
	}

	fields := makeFieldsFromSchema(header, opts.Schema, len(rows), opts.IgnoreUnknown, regex)

	wg := sync.WaitGroup{}
	wg.Add(len(fields))

	for fieldIdx := range fields {
		go func(fieldIdx int) {
			defer wg.Done()

			f := fields[fieldIdx]
			if f == nil {
				// Ignoring unknown field.
				return
			}

			var timeLayout string
			if f.Type() == data.FieldTypeNullableTime {
				// Ensure we have data and the field index is valid
				if len(rows) > 0 && len(rows[0]) > fieldIdx {
					layout, err := detectTimeLayoutNaive(rows[0][fieldIdx])
					if err != nil {
						logger.Warn(fmt.Sprintf("Parse csv error: %s", err.Error()), "timeField", rows[0][fieldIdx])
						return 
					}
					timeLayout = layout
				} else {
					logger.Warn("No data available for time layout detection")
					return
				}
			}

			for rowIdx := 0; rowIdx < f.Len(); rowIdx++ {
				// Ensure the row and field indices are valid
				if rowIdx < len(rows) && fieldIdx < len(rows[rowIdx]) {
					if err := parseCell(rows[rowIdx][fieldIdx], opts, timeLayout, rowIdx, f, location); err != nil {
						// Ignore any cells that couldn't be parsed.
						f.Set(rowIdx, nil)
					}
				} else {
					// Set nil for missing data
					f.Set(rowIdx, nil)
				}
			}
		}(fieldIdx)
	}

	wg.Wait()

	// Remove ignored fields from result.
	var res []*data.Field
	for _, f := range fields {
		if f != nil {
			res = append(res, f)
		}
	}

	return res, nil
}

func readCSV(opts csvOptions, r io.Reader) ([]string, [][]string, error) {
	// Read one byte at a time until we've counted newlines equal to the number
	// of skipped rows.
	for i := 0; i < opts.SkipRows; i++ {
		buf := make([]byte, 1)
		for {
			_, err := r.Read(buf)
			if err != nil || buf[0] == '\n' {
				break
			}
		}
	}

	rd := csv.NewReader(r)
	rd.LazyQuotes = true

	if len(opts.Delimiter) >= 1 {
		rd.Comma = rune(opts.Delimiter[0])
	}

	records, err := rd.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	var rows [][]string
	var header []string

	if len(records) == 0 {
		return nil, nil, errors.New("no records found in CSV")
	}

	if opts.Header {
		header = records[0]
		rows = records[1:]
	} else {
		if len(records) > 0 && len(records[0]) > 0 {
			for i := 0; i < len(records[0]); i++ {
				header = append(header, fmt.Sprintf("Field %d", i+1))
			}
		}
		rows = records
	}

	return header, rows, nil
}

func parseCell(value string, opts csvOptions, timeLayout string, rowIdx int, f *data.Field, tz *time.Location) error {
	switch f.Type() {
	case data.FieldTypeNullableFloat64:
		intPart, fracPart := splitNumberParts(value, opts.DecimalSeparator)

		// Only one separator is allowed.
		if strings.Contains(intPart, opts.DecimalSeparator) {
			return errors.New("multiple decimal separators")
		}

		converted := strings.Map(func(r rune) rune {
			switch r {
			case ',', '.', ' ':
				if string(r) == opts.DecimalSeparator {
					return r
				}
				// Returning a negative value removes the character.
				return -1
			default:
				return r
			}
		}, intPart) + "." + fracPart

		n, err := strconv.ParseFloat(converted, 64)
		if err != nil {
			return err
		}

		f.Set(rowIdx, &n)
	case data.FieldTypeNullableBool:
		n, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		f.Set(rowIdx, &n)
	case data.FieldTypeNullableTime:
		n, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			t := time.Unix(n, 0)
			f.Set(rowIdx, &t)
			return nil
		}

		if timeLayout != "" {
			t, err := time.ParseInLocation(timeLayout, value, tz)
			if err == nil {
				f.Set(rowIdx, &t)
				return nil
			}
		}

		return errors.New("unsupported time format")
	default:
		f.Set(rowIdx, &value)
	}

	return nil
}

// makeFieldsFromSchema returns a field for every column defined in the header.
func makeFieldsFromSchema(header []string, schema []fieldSchema, size int, ignoreUnknown bool, regex bool) []*data.Field {
	fields := make([]*data.Field, len(header))

	for i, name := range header {
		t, ok := typeFromName(name, schema, regex)
		if !ok {
			if ignoreUnknown {
				continue
			}
			t = data.FieldTypeNullableString
		}

		f := data.NewFieldFromFieldType(t, size)
		f.Name = name
		fields[i] = f
	}

	return fields
}

func typeFromName(name string, schemas []fieldSchema, regex bool) (data.FieldType, bool) {
	sch, ok := findSchema(schemas, name, regex)
	if !ok {
		return data.FieldTypeUnknown, false
	}
	return fieldType(sch.Type), true
}

func fieldType(str string) data.FieldType {
	switch str {
	case "number":
		return data.FieldTypeNullableFloat64
	case "boolean":
		return data.FieldTypeNullableBool
	case "time":
		return data.FieldTypeNullableTime
	default:
		return data.FieldTypeNullableString
	}
}

func findSchema(fields []fieldSchema, name string, regex bool) (res fieldSchema, ok bool) {
	for _, sch := range fields {
		if sch.Name == "" {
			continue
		}

		if regex {
			re, err := regexp.Compile(sch.Name)
			if err != nil {
				return
			}
			if re.MatchString(name) {
				res = sch
				ok = true
			}
		} else {
			if sch.Name == name {
				res = sch
				ok = true
			}
		}
	}
	return
}

func splitNumberParts(str string, sep string) (string, string) {
	idx := strings.LastIndex(str, sep)

	if idx < 0 {
		return str, "0"
	}

	// Check bounds to prevent slice out of range panic
	if idx+1 >= len(str) {
		return str[:idx], "0"
	}

	return str[:idx], str[idx+1:]
}

// detectTimeLayoutNaive attempts to parse the string from a set of layouts, and
// returns the first one that matched.
func detectTimeLayoutNaive(str string) (string, error) {
	layouts := []string{
		"2006-01-02",
		"2006-01-02 15:04",
		"2006-01-02 15:04:05 MST",
		"2006-01-02 15:04:05.999999",
		"2006-01-02 15:04:05.999999 -07:00",
		"2006-01-02 15:04:05.999999Z07:00",
		"2006-01-02T15:04",
		"2006-01-02T15:04:05 MST",
		"2006-01-02T15:04:05.999999",
		"2006-01-02T15:04:05.999999 -07:00",
		"2006-01-02T15:04:05.999999Z07:00",
		"2006/1/2",
		"2006/01/02",
		"2006/01/02 15:04",
		"2006/01/02 15:04:05 MST",
		"2006/01/02 15:04:05.999999",
		"2006/01/02 15:04:05.999999 -07:00",
		"2006/01/02 15:04:05.999999Z07:00",
		"2006/01/02T15:04",
		"2006/01/02T15:04:05 MST",
		"2006/01/02T15:04:05.999999",
		"2006/01/02T15:04:05.999999 -07:00",
		"2006/01/02T15:04:05.999999Z07:00",
		"01/02/2006",
		"01/02/2006 15:04",
		"01/02/2006 15:04:05 MST",
		"01/02/2006 15:04:05.999999",
		"01/02/2006 15:04:05.999999 -07:00",
		"01/02/2006 15:04:05.999999Z07:00",
		"01/02/2006T15:04",
		"01/02/2006T15:04:05 MST",
		"01/02/2006T15:04:05.999999",
		"01/02/2006T15:04:05.999999 -07:00",
		"01/02/2006T15:04:05.999999Z07:00",
	}

	for _, layout := range layouts {
		if _, err := time.Parse(layout, str); err == nil {
			return layout, nil
		}
	}

	return "", backend.DownstreamError(errors.New("unsupported time format"))
}
