package tsError

import (
	"testing"
)

func TestGetError(t *testing.T) {
	errResponse, err := GetError("ERROR-01", "got error message")
	if err != nil {
		t.Errorf("not able to marshal the error message")
	}
	res := string(errResponse)
	if res != "\n\bERROR-01\x12\x11got error message" {
		t.Errorf("%q", errResponse)
	}
}
