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

	if len(result.Left) != 128 {
		t.Errorf("Expected length to be 256 got %d", len(result.Left))
	}

	if result.SampleRate != 44100.0 {
		t.Errorf("Expected sample rate 44100 Hz got %f", result.SampleRate)
	}
}

func TestReadWrongFormat(t *testing.T) {
	riffDataReader := pkg_test_utils.NewFailingMockRiffReader()
	result, err := pkg.Read(riffDataReader)

	if err == nil {
		t.Errorf("Should have failed")
	}

	if len(result.Left) != 0 || len(result.Right) != 0 {
		t.Errorf("Should have no data")
	}
}
