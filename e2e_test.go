package mdprinter_test

import (
	"os"
	"os/exec"
	"testing"
)

const (
	binPath    = "./bin/mdprinter"
	inputPath  = "examples/example.md"
	outputPath = "test_output/example.pdf"
	dataPath   = "examples/exampleData.json"
)

func RunE2E(t *testing.T, args ...string) {
	// Run the program as a subprocess
	// with the input file and output file as arguments
	cmd := exec.Command(binPath, args...)
	err := cmd.Run()
	if err != nil {
		t.Errorf("Failed to run the program: %v", err)
		return
	}

	// Check if the output file exists
	_, err = os.Stat(outputPath)
	if err != nil {
		t.Errorf("Failed to create the output file: %v", err)
		return
	}

	// Clean up the output file
	err = os.Remove(outputPath)
	if err != nil {
		t.Errorf("Failed to clean up the output file: %v", err)
		return
	}
}

// This is an end-to-end test that will run the entire program.
func TestE2E_1(t *testing.T) {
	RunE2E(t, inputPath, "-o", outputPath)
}

func TestE2E_2(t *testing.T) {
	RunE2E(t, "-s=air", "-o="+outputPath, inputPath)
}

func TestE2E_3(t *testing.T) {
	RunE2E(t, "-s", "air", "-o", outputPath, inputPath)
}
