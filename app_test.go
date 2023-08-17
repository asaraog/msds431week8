package main

import (
	"testing"
)

func TestProcessRaw(t *testing.T) {

	processedout, _ := ProcessRaw("\n")
	if len(processedout) == 1 {
		t.Errorf("Error in data splitting. ")
	}

}
