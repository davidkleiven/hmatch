package pkg_test

import (
	"testing"

	"github.com/davidkleiven/hmatch/internal/pkg_test_utils"
	"github.com/davidkleiven/hmatch/pkg"
)

func TestRead(t *testing.T) {
	riffDataReader := pkg_test_utils.NewMockRiffReader()

	result, err := pkg.Read(riffDataReader)

	if err != nil {
		t.Errorf("Received error %v", err)
	}

	if len(result) != 1 {
		t.Errorf("Expected 1 chunk result. Got %d", len(result))
	}

	if len(result[0].Left) != 128 {
		t.Errorf("Expected length to be 256 got %d", len(result[0].Left))
	}
}
