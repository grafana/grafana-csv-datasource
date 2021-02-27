package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"sync"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type fieldSchema struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type csvOptions struct {
	Delimiter     string        `json:"delimiter"`
	Header        bool          `json:"header"`
	IgnoreUnknown bool          `json:"ignoreUnknown"`
	Schema        []fieldSchema `json:"schema"`
	SkipRows      int           `json:"skipRows"`
}

func parseCSV(opts csvOptions, r io.Reader) ([]*data.Field, error) {
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

	if len(opts.Delimiter) == 1 {
		rd.Comma = rune(opts.Delimiter[0])
	}

	records, err := rd.ReadAll()
	if err != nil {
		return nil, err
	}

	var rows [][]string
	var header []string

	if opts.Header {
		header = records[0]
		rows = records[1:]
	} else {
		if len(records) > 0 {
			for i := 0; i < len(records[0]); i++ {
				header = append(header, fmt.Sprintf("Field %d", i+1))
			}
		}
		rows = records
	}

	fields := make([]*data.Field, 0)

	// Create fields from schema.
	for _, name := range header {
		sch, ok := schemaContains(opts.Schema, name)
		if !ok {
			if opts.IgnoreUnknown {
				// Add a null field to maintain index.
				fields = append(fields, nil)
				continue
			} else {
				sch = fieldSchema{Name: name, Type: "string"}
			}
		}
		f := data.NewFieldFromFieldType(fieldType(sch.Type), len(rows))
		f.Name = name
		fields = append(fields, f)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(fields))

	for fieldIdx := range fields {
		go func(fieldIdx int) {
			defer wg.Done()

			f := fields[fieldIdx]
			if f == nil {
				return
			}

			var timeLayout string
			if f.Type() == data.FieldTypeNullableTime {
				layout, err := detectTimeLayoutNaive(rows[0][fieldIdx])
				if err == nil {
					timeLayout = layout
				}
			}

			for rowIdx := 0; rowIdx < f.Len(); rowIdx++ {
				value := rows[rowIdx][fieldIdx]

				switch f.Type() {
				case data.FieldTypeNullableFloat64:
					n, err := strconv.ParseFloat(value, 10)
					if err != nil {
						f.Set(rowIdx, nil)
						continue
					}
					f.Set(rowIdx, &n)
				case data.FieldTypeNullableBool:
					n, err := strconv.ParseBool(value)
					if err != nil {
						f.Set(rowIdx, nil)
						continue
					}
					f.Set(rowIdx, &n)
				case data.FieldTypeNullableTime:
					n, err := strconv.ParseInt(value, 10, 64)
					if err == nil {
						t := time.Unix(n, 0)
						f.Set(rowIdx, &t)
						continue
					}

					if timeLayout != "" {
						t, err := time.Parse(timeLayout, value)
						if err == nil {
							f.Set(rowIdx, &t)
							continue
						}
					}

					f.Set(rowIdx, nil)
				default:
					f.Set(rowIdx, &value)
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

func schemaContains(fields []fieldSchema, name string) (fieldSchema, bool) {
	for _, sch := range fields {
		if sch.Name == name {
			return sch, true
		}
	}
	return fieldSchema{}, false
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
		"2006-01-02 15:04:05.999999Z",
		"2006-01-02T15:04",
		"2006-01-02T15:04:05 MST",
		"2006-01-02T15:04:05.999999",
		"2006-01-02T15:04:05.999999 -07:00",
		"2006-01-02T15:04:05.999999Z",
		"2006/01/02",
		"2006/1/2",
		"2006/01/02 15:04",
		"2006/01/02 15:04:05 MST",
		"2006/01/02 15:04:05.999999",
		"2006/01/02 15:04:05.999999 -07:00",
		"2006/01/02 15:04:05.999999Z",
		"2006/01/02T15:04",
		"2006/01/02T15:04:05 MST",
		"2006/01/02T15:04:05.999999",
		"2006/01/02T15:04:05.999999 -07:00",
		"2006/01/02T15:04:05.999999Z",
		"01/02/2006",
		"01/02/2006 15:04",
		"01/02/2006 15:04:05 MST",
		"01/02/2006 15:04:05.999999",
		"01/02/2006 15:04:05.999999 -07:00",
		"01/02/2006 15:04:05.999999Z",
		"01/02/2006T15:04",
		"01/02/2006T15:04:05 MST",
		"01/02/2006T15:04:05.999999",
		"01/02/2006T15:04:05.999999 -07:00",
		"01/02/2006T15:04:05.999999Z",
	}

	for _, layout := range layouts {
		if _, err := time.Parse(layout, str); err == nil {
			return layout, nil
		}
	}

	return "", errors.New("unsupported time format")
}
