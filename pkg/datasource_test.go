package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

func TestToColumns(t *testing.T) {
	input := [][]string{
		[]string{"foo", "bar", "baz"},
		[]string{"foo", "bar", "baz"},
		[]string{"foo", "bar", "baz"},
	}

	want := [][]string{
		[]string{"foo", "foo", "foo"},
		[]string{"bar", "bar", "bar"},
		[]string{"baz", "baz", "baz"},
	}

	got := toColumns(input)

	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestParseCSV(t *testing.T) {
	for _, tt := range []struct {
		query  queryModel
		input  string
		output []*data.Field
	}{
		{query: queryModel{}, input: "foo,bar,baz\n1,2,3", output: []*data.Field{
			data.NewField("foo", nil, []string{"1"}),
		}},
	} {
		t.Run("", func(t *testing.T) {
			fields, err := parseCSV(tt.query, strings.NewReader(tt.input))
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
