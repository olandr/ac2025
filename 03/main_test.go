package main

import (
	"os"
	"testing"
)

func TestProcessWithSamples(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedCount int64
	}{
		{
			name:          "sample1 - 987654321111111",
			input:         "987654321111111\n",
			expectedCount: 987654321111,
		},
		{
			name:          "sample2 - 811111111111119",
			input:         "811111111111119\n",
			expectedCount: 811111111119,
		},
		{
			name:          "sample2 - 234234234234278",
			input:         "234234234234278\n",
			expectedCount: 434234234278,
		},
		{
			name:          "sample2 - 818181911112111",
			input:         "818181911112111\n",
			expectedCount: 888911112111,
		},
		{
			name: "sample - multiple lines",
			input: `987654321111111
811111111111119
234234234234278
818181911112111
`,
			expectedCount: 3121910778619, // This is the sum checked in the main code
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary input and output
			inFile, err := os.CreateTemp("", "test-input-*.txt")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(inFile.Name())
			defer inFile.Close()

			outFile, err := os.CreateTemp("", "test-output-*.txt")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(outFile.Name())
			defer outFile.Close()

			// Write test input
			if _, err := inFile.WriteString(tt.input); err != nil {
				t.Fatal(err)
			}
			if _, err := inFile.Seek(0, 0); err != nil {
				t.Fatal(err)
			}

			// Run process with battery toggle length of 12
			actualCount := process(inFile, outFile, 12)

			// Assert the count matches expected
			if actualCount != tt.expectedCount {
				t.Errorf("process() returned count = %d, expected %d", actualCount, tt.expectedCount)
			}
		})
	}
}
