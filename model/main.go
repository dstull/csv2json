package model

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"io"
	"os"
)

type Encoder struct {
	src  io.ReadCloser
	dest io.WriteCloser
}

func FromArgs(args []string) (*Encoder, error) {
	fl := flag.NewFlagSet("csv2json", flag.ExitOnError)

	src := fl.String("i", "", "Input file (default: stdin)")
	dest := fl.String("o", "", "Output file (default: stdout)")
	_ = fl.Parse(args)

	encoder := Encoder{
		src:  os.Stdin,
		dest: os.Stdout,
	}
	if *src != "" && *src != "-" {
		readHandle, err := os.Open(*src)
		if err != nil {
			return nil, err
		}
		encoder.src = readHandle
	}
	if *dest != "" && *dest != "-" {
		writeHandle, err := os.Create(*dest)
		if err != nil {
			return nil, err
		}
		encoder.dest = writeHandle
	}
	return &encoder, nil
}

func (encoder *Encoder) Encode() error {
	return encoder.asColumnAndLines()
}

func (encoder *Encoder) asColumnAndLines() (err error) {
	defer deferClose(&err, encoder.dest.Close)

	data, err := makeAsColumneAndLines(encoder.src)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(encoder.dest)
	return enc.Encode(&data)
}

type JSONStructure struct {
	Columns []string   `json:"columns"`
	Lines   [][]string `json:"lines"`
}

func makeAsColumneAndLines(src io.ReadCloser) (jsonStr *JSONStructure, err error) {
	defer deferClose(&err, src.Close)

	cr := csv.NewReader(src)
	cr.Comment = '#'
	cr.FieldsPerRecord = -1

	columns, err := cr.Read()

	lines := [][]string{}

	for {
		fields, err := cr.Read()
		if err == io.EOF {
			jsonStr := JSONStructure{
				Columns: columns,
				Lines:   lines,
			}
			return &jsonStr, nil
		}

		if err != nil {
			return nil, err
		}

		line := make([]string, len(fields))

		for i, field := range fields {
			line[i] = field
		}

		lines = append(lines, line)
	}
}

func deferClose(err *error, f func() error) {
	newErr := f()
	if *err == nil {
		*err = newErr
	}
}
