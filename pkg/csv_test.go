package main

import (
	"bytes"
	"math"
	"strings"
	"testing"

	"github.com/grafana/grafana-plugin-sdk-go/data"
)

func TestParseCSV(t *testing.T) {
	for _, tt := range []struct {
		query  csvOptions
		input  string
		output []*data.Field
	}{
		{query: csvOptions{}, input: "foo,bar,baz\n1,2,3", output: []*data.Field{
			data.NewField("foo", nil, []string{"1"}),
		}},
	} {
		t.Run("", func(t *testing.T) {
			fields, err := parseCSV(tt.query, false, strings.NewReader(tt.input))
			if err != nil {
				t.Fatal(err)
			}

			f1 := data.Frame{Fields: fields}
			f2 := data.Frame{Fields: tt.output}

			b1, _ := f1.MarshalArrow()
			b2, _ := f2.MarshalArrow()

			if bytes.Equal(b1, b2) {
				t.Fatal("unexpected output")
			}
		})
	}
}

func TestParseTimeNaive(t *testing.T) {
	for _, tt := range []struct {
		input  string
		output string
	}{
		{input: "2018-08-19", output: "2006-01-02"},
		{input: "2018-08-19 12:11", output: "2006-01-02 15:04"},
		{input: "2018-08-19 12:11:35", output: "2006-01-02 15:04:05.999999"},
		{input: "2018-08-19 12:11:35.22", output: "2006-01-02 15:04:05.999999"},
		{input: "2018-08-19 12:11:35Z", output: "2006-01-02 15:04:05.999999Z07:00"},
		{input: "2018-08-19 12:11:35+01:00", output: "2006-01-02 15:04:05.999999Z07:00"},
		{input: "2018-08-19 12:11:35.220Z", output: "2006-01-02 15:04:05.999999Z07:00"},
		{input: "2018-08-19 12:11:35.220+01:00", output: "2006-01-02 15:04:05.999999Z07:00"},
		{input: "2018-08-19 07:11:35.220 -05:00", output: "2006-01-02 15:04:05.999999 -07:00"},
		{input: "2018-07-05 12:54:00 UTC", output: "2006-01-02 15:04:05 MST"},
		{input: "2018-08-19T12:11:35", output: "2006-01-02T15:04:05.999999"},
		{input: "2018-08-19T12:11:35.22", output: "2006-01-02T15:04:05.999999"},
		{input: "2018-08-19T12:11:35Z", output: "2006-01-02T15:04:05.999999Z07:00"},
		{input: "2018-08-19T12:11:35+01:00", output: "2006-01-02T15:04:05.999999Z07:00"},
		{input: "2018-08-19T12:11:35.220Z", output: "2006-01-02T15:04:05.999999Z07:00"},
		{input: "2018-08-19T12:11:35.220+01:00", output: "2006-01-02T15:04:05.999999Z07:00"},
		{input: "2018/08/19", output: "2006/1/2"},
		{input: "2018/08/19 12:11", output: "2006/01/02 15:04"},
		{input: "2018/08/19 12:11:35", output: "2006/01/02 15:04:05.999999"},
		{input: "2018/08/19 12:11:35Z", output: "2006/01/02 15:04:05.999999Z07:00"},
		{input: "2018/08/19 12:11:35-05:00", output: "2006/01/02 15:04:05.999999Z07:00"},
		{input: "2018/08/19 12:11:35.22", output: "2006/01/02 15:04:05.999999"},
		{input: "2018/08/19T12:11:35Z", output: "2006/01/02T15:04:05.999999Z07:00"},
		{input: "2018/08/19T12:11:35-05:00", output: "2006/01/02T15:04:05.999999Z07:00"},
		{input: "2018/9/8", output: "2006/1/2"},
		{input: "08/19/2018", output: "01/02/2006"},
		{input: "08/19/2018 12:11:35Z", output: "01/02/2006 15:04:05.999999Z07:00"},
		{input: "08/19/2018 12:11:35-05:00", output: "01/02/2006 15:04:05.999999Z07:00"},
		{input: "08/19/2018T12:11:35Z", output: "01/02/2006T15:04:05.999999Z07:00"},
		{input: "08/19/2018T12:11:35-05:00", output: "01/02/2006T15:04:05.999999Z07:00"},
	} {
		t.Run(tt.input, func(t *testing.T) {
			got, err := detectTimeLayoutNaive(tt.input)
			if err != nil {
				t.Fatal(err)
			}

			if got != tt.output {
				t.Fatalf("want = %q; got = %q", tt.output, got)
			}
		})
	}
}

func TestParseLazyQuotes(t *testing.T) {
	opts := csvOptions{
		Delimiter: ",",
	}

	for _, tt := range []struct {
		In string
	}{
		{In: `"I","can't","even"`},
		{In: `'I','can"t','even'`},
		{In: `I,can't,even`},
		{In: `I,can"t,even`},
	} {
		t.Run("", func(t *testing.T) {
			_, err := parseCSV(opts, false, strings.NewReader(tt.In))
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestDecimalSeparator(t *testing.T) {
	for _, tt := range []struct {
		In  string
		Out float64
		Sep string
	}{
		{In: `10`, Out: 10, Sep: ","},
		{In: `10.4`, Out: 10.4, Sep: "."},
		{In: `10,4`, Out: 104, Sep: "."},
		{In: `10,4`, Out: 10.4, Sep: ","},
		{In: `10.4`, Out: 104, Sep: ","},
		{In: `10 000,12`, Out: 10000.12, Sep: ","},
		{In: `10.000,12`, Out: 10000.12, Sep: ","},
		{In: `10.000.000,12`, Out: 10000000.12, Sep: ","},
		{In: `10,000.12`, Out: 10000.12, Sep: ","},
		{In: `10,000,000.12`, Out: 10000000.12, Sep: ","},
		{In: `10,000,12`, Out: math.NaN(), Sep: ","},
		{In: `10,000,12`, Out: 1000012, Sep: "."},
		{In: `10.000.12`, Out: math.NaN(), Sep: "."},
		{In: `10.000.12`, Out: 1000012, Sep: ","},
	} {
		t.Run(tt.In+"_"+tt.Sep, func(t *testing.T) {
			opts := csvOptions{
				Delimiter:        ";",
				DecimalSeparator: tt.Sep,
				Schema: []fieldSchema{
					{Name: "Field 1", Type: "number"},
				},
			}

			fields, err := parseCSV(opts, false, strings.NewReader(tt.In))
			if err != nil {
				t.Fatal(err)
			}

			if len(fields) != 1 {
				t.Fatalf("unexpected number of fields: %v", len(fields))
			}

			numberField := fields[0]

			if numberField.Len() != 1 {
				t.Fatalf("unexpected field size: %v", numberField.Len())
			}

			got, err := numberField.FloatAt(0)
			if err != nil {
				t.Fatal(err)
			}

			if !math.IsNaN(got) && math.IsNaN(tt.Out) {
				t.Fatalf("want = %v; got = %v", tt.Out, got)
			}

			if !math.IsNaN(got) && tt.Out != got {
				t.Errorf("want = %v; got = %v", tt.Out, got)
			}
		})
	}
}
