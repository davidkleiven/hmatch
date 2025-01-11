package standard

import (
	"slices"
	"testing"
)

func TestFindRegistration(t *testing.T) {
	result := FindRegistration("tib")
	wantNames := []string{"Tibia 16'", "Tibia 8'", "Tibia 2'"}
	got := make([]string, len(result))
	for i, v := range result {
		got[i] = v.Name
	}
	if slices.Compare(wantNames, got) != 0 {
		t.Errorf("Wanted %v\ngot%v\n", wantNames, got)
	}
}
