package day05

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReaderProcess(t *testing.T) {
	reader := NewReader("test_input.txt")
	got := reader.Process()
	want := 35

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Reader.Process() mismatch (-want +got):\n%s", diff)
	}
}
