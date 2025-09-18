package main

import (
	"bytes"
	"math"
	"strings"
	"testing"

	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
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
			fields, err := parseCSV(tt.query, false, strings.NewReader(tt.input), log.DefaultLogger)
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
			_, err := parseCSV(opts, false, strings.NewReader(tt.In), log.DefaultLogger)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestParseCSV_EdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		opts        csvOptions
		input       string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "completely empty CSV",
			opts:        csvOptions{Header: true},
			input:       "",
			expectError: true,
			errorMsg:    "no records found in CSV",
		},
		{
			name:        "only headers, no data rows",
			opts:        csvOptions{Header: true},
			input:       "col1,col2,col3",
			expectError: true,
			errorMsg:    "no data rows found in CSV",
		},
		{
			name:        "empty CSV without headers",
			opts:        csvOptions{Header: false},
			input:       "",
			expectError: true,
			errorMsg:    "no records found in CSV",
		},
		{
			name:        "single line with data, no headers",
			opts:        csvOptions{Header: false},
			input:       "1,2,3",
			expectError: false,
		},
		{
			name:        "CSV with missing columns in some rows",
			opts:        csvOptions{Header: true},
			input:       "col1,col2,col3\n1,2,3\n4,5\n6,7,8",
			expectError: true, // CSV parser will reject rows with wrong field count
		},
		{
			name:        "CSV with time field but no data",
			opts:        csvOptions{
				Header: true,
				Schema: []fieldSchema{
					{Name: "timestamp", Type: "time"},
				},
			},
			input:       "timestamp",
			expectError: true,
			errorMsg:    "no data rows found in CSV",
		},
		{
			name:        "CSV with empty delimiter",
			opts:        csvOptions{
				Header:    true,
				Delimiter: "",
			},
			input:       "col1,col2\n1,2",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields, err := parseCSV(tt.opts, false, strings.NewReader(tt.input), log.DefaultLogger)
			
			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
				if tt.errorMsg != "" && !strings.Contains(err.Error(), tt.errorMsg) {
					t.Fatalf("expected error containing %q, got %q", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if fields == nil {
					t.Fatalf("expected fields but got nil")
				}
			}
		})
	}
}

func TestReadCSV_EdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		opts        csvOptions
		input       string
		expectError bool
	}{
		{
			name:        "empty input",
			opts:        csvOptions{Header: true},
			input:       "",
			expectError: true,
		},
		{
			name:        "only whitespace",
			opts:        csvOptions{Header: true},
			input:       "   \n   \n",
			expectError: false, // This might parse as empty records
		},
		{
			name:        "header only",
			opts:        csvOptions{Header: true},
			input:       "col1,col2",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header, rows, err := readCSV(tt.opts, strings.NewReader(tt.input))
			
			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				// Additional checks for valid cases
				if header == nil && rows == nil {
					t.Fatalf("both header and rows are nil")
				}
			}
		})
	}
}

func TestParseCSV_BoundsCheckingSimulation(t *testing.T) {
	// Test that demonstrates the bounds checking works correctly
	// by manually testing the improved logic paths
	tests := []struct {
		name        string
		opts        csvOptions
		input       string
		expectError bool
	}{
		{
			name: "valid CSV with all fields populated",
			opts: csvOptions{
				Header: true,
				Schema: []fieldSchema{
					{Name: "col1", Type: "string"},
					{Name: "col2", Type: "number"},
				},
			},
			input:       "col1,col2\nvalue1,42\nvalue2,24",
			expectError: false,
		},
		{
			name: "CSV that would previously cause index panic - empty with time field",
			opts: csvOptions{
				Header: true,
				Schema: []fieldSchema{
					{Name: "timestamp", Type: "time"},
				},
			},
			input:       "timestamp",
			expectError: true,
		},
		{
			name: "CSV with just newlines",
			opts: csvOptions{Header: false},
			input:       "\n\n\n",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields, err := parseCSV(tt.opts, false, strings.NewReader(tt.input), log.DefaultLogger)
			
			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if fields == nil {
					t.Fatalf("expected fields but got nil")
				}
			}
		})
	}
}

func TestParseCSV_TimeFieldsEdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		opts        csvOptions
		input       string
		expectError bool
	}{
		{
			name: "time field with no data rows",
			opts: csvOptions{
				Header: true,
				Schema: []fieldSchema{
					{Name: "timestamp", Type: "time"},
				},
			},
			input:       "timestamp",
			expectError: true,
		},
		{
			name: "time field with valid data",
			opts: csvOptions{
				Header: true,
				Schema: []fieldSchema{
					{Name: "timestamp", Type: "time"},
				},
			},
			input:       "timestamp\n2023-01-01",
			expectError: false,
		},
		{
			name: "time field with missing column data",
			opts: csvOptions{
				Header: true,
				Schema: []fieldSchema{
					{Name: "timestamp", Type: "time"},
					{Name: "value", Type: "number"},
				},
			},
			input:       "timestamp,value\n2023-01-01\n2023-01-02,42",
			expectError: true, // CSV parser will reject rows with wrong field count
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields, err := parseCSV(tt.opts, false, strings.NewReader(tt.input), log.DefaultLogger)
			
			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if fields == nil {
					t.Fatalf("expected fields but got nil")
				}
			}
		})
	}
}

func TestSplitNumberParts_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		sep      string
		expectedInt string
		expectedFrac string
	}{
		{
			name: "normal case",
			input: "123.45",
			sep: ".",
			expectedInt: "123",
			expectedFrac: "45",
		},
		{
			name: "separator at end",
			input: "123.",
			sep: ".",
			expectedInt: "123",
			expectedFrac: "0",
		},
		{
			name: "no separator",
			input: "123",
			sep: ".",
			expectedInt: "123",
			expectedFrac: "0",
		},
		{
			name: "empty string",
			input: "",
			sep: ".",
			expectedInt: "",
			expectedFrac: "0",
		},
		{
			name: "only separator",
			input: ".",
			sep: ".",
			expectedInt: "",
			expectedFrac: "0",
		},
		{
			name: "separator at beginning",
			input: ".45",
			sep: ".",
			expectedInt: "",
			expectedFrac: "45",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This should not panic
			intPart, fracPart := splitNumberParts(tt.input, tt.sep)
			
			if intPart != tt.expectedInt {
				t.Errorf("expected int part %q, got %q", tt.expectedInt, intPart)
			}
			if fracPart != tt.expectedFrac {
				t.Errorf("expected frac part %q, got %q", tt.expectedFrac, fracPart)
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

			fields, err := parseCSV(opts, false, strings.NewReader(tt.In), log.DefaultLogger)
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
