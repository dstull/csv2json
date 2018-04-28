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

	src := fl.String("src", "", "Source file (default: stdin)")
	dest := fl.String("dest", "", "Destination file (default: stdout)")
	_ = fl.Parse(args)

	e := Encoder{
		src:  os.Stdin,
		dest: os.Stdout,
	}
	if *src != "" && *src != "-" {
		f, err := os.Open(*src)
		if err != nil {
			return nil, err
		}
		e.src = f
	}
	if *dest != "" && *dest != "-" {
		f, err := os.Create(*dest)
		if err != nil {
			return nil, err
		}
		e.dest = f
	}
	return &e, nil
}

func (e *Encoder) Encode() error {
	return e.asColumnAndLines()
}

func (e *Encoder) asColumnAndLines() (err error) {
	defer deferClose(&err, e.dest.Close)

	data, err := makeAsColumneAndLines(e.src)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(e.dest)
	return enc.Encode(&data)
}

func makeAsColumneAndLines(src io.ReadCloser) (data [][]string, err error) {
	defer deferClose(&err, src.Close)

	cr := csv.NewReader(src)
	cr.Comment = '#'
	cr.FieldsPerRecord = -1

	return cr.ReadAll()
	// defer deferClose(&err, src.Close)
	//
	// cr := csv.NewReader(src)
	// cr.Comment = '#'
	// cr.FieldsPerRecord = -1
	// cr.ReuseRecord = true
	//
	// fields, err := cr.Read()
	//
	// // Save headers for each row of dict
	// dataHeader := make(map[int]string, len(fields))
	// for i, field := range fields {
	// 	dataHeader[i] = field
	// }
	//
	// for {
	// 	fields, err = cr.Read()
	// 	if err == io.EOF {
	// 		return data, nil
	// 	}
	//
	// 	if err != nil {
	// 		return nil, err
	// 	}
	//
	// 	datum := make(map[string]string, len(fields))
	// 	for i, val := range fields {
	// 		datum[dataHeader[i]] = val
	// 	}
	// 	data = append(data, datum)
	// }
}

func deferClose(err *error, f func() error) {
	newErr := f()
	if *err == nil {
		*err = newErr
	}
}
