package main

import (
	"encoding/csv"
	"io"
)

// Record Color Entry
type Recorder struct {
	CsvWriter *csv.Writer
}

// Create New Recorder
func NewRecorder(writer *io.Writer) *Recorder {
	var s Recorder
	s.CsvWriter = csv.NewWriter(*writer)
	return &s
}
