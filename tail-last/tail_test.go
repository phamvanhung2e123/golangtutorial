package main

import "testing"

func TestFileOffset(t *testing.T) {
	f := func(fileName string, expectedOffset int64, expectError bool) {
		offset, err := fileOffset(fileName, 10)

		if expectError && err == nil {
			t.Errorf("Expected error, got result number '%d'. Input file: %s", offset, fileName)
		} else if !expectError && err != nil {
			t.Errorf("Not expected error, got '%s'. Input file: %s", err, fileName)
		} else {
			if expectedOffset != offset {
				t.Errorf("Expected %d lines, got %d in file '%s'", offset, expectedOffset, fileName)
			}
		}
	}

	f("fixtures/file.txt", -20, false)
	f("fixtures/file1.txt", -60, false)
	f("fixtures/file2.txt",  -25, false)
	f("fixtures/some-unknown-file.txt", 0, true)
}