package pkg

import (
	"errors"
	"slices"
	"testing"
)

func TestCombine(t *testing.T) {
	for _, test := range []struct {
		positions []DrawbarSetting
		want      DrawbarSetting
		desc      string
	}{
		{
			positions: []DrawbarSetting{{1, 2, 3, 4, 5, 6, 7, 8, 8}},
			want:      DrawbarSetting{1, 2, 3, 4, 5, 6, 7, 8, 8},
			desc:      "Combining one should return the same result",
		},
		{
			positions: []DrawbarSetting{{0, 0, 2, 0, 0, 0, 0, 0, 0}, {0, 0, 2, 0, 0, 0, 0, 0, 0}},
			want:      DrawbarSetting{0, 0, 3, 0, 0, 0, 0, 0, 0},
			desc:      "Combining two should produce twice the power (10*log(2) = 3 dB higher)",
		},
		{
			positions: []DrawbarSetting{{0, 0, 1, 0, 0, 0, 0, 0, 0}, {0, 0, 2, 0, 0, 0, 0, 0, 0}},
			want:      DrawbarSetting{0, 0, 3, 0, 0, 0, 0, 0, 0},
			desc:      "Combining two of the same with uneven strength should also be 3 10*log(1+10^(3/10)) = 4.8 dB higher than the first",
		},
		{
			positions: []DrawbarSetting{{1, 2, 3, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 2, 5, 3}},
			want:      DrawbarSetting{1, 2, 3, 0, 0, 0, 2, 5, 3},
			desc:      "Combining two non overlapping combinations should be equal to summing the drawbar positions",
		},
		{
			positions: []DrawbarSetting{{8, 8, 8, 0, 0, 0, 0, 0, 0}, {0, 0, 8, 0, 0, 0, 0, 0, 0}},
			want:      DrawbarSetting{7, 7, 8, 0, 0, 0, 0, 0, 0},
			desc:      "One is overflowing, the two first should be normalized",
		},
	} {
		result := Combine(test.positions)

		if slices.Compare(result[:], test.want[:]) != 0 {
			t.Errorf("Test %s:\nwanted\n%v\ngot\n%v\n", test.desc, test.want, result)
		}
	}
}

func TestParseDrawbarPositionInvalidDigits(t *testing.T) {
	for _, test := range []struct {
		input string
		desc  string
		err   error
	}{
		{
			input: "00 0000 000",
			desc:  "Contains spaces",
			err:   ErrInvalidCharacters,
		},
		{
			input: "0000000",
			desc:  "Contains only eight digits",
			err:   ErrInvalidLengthOfSubtoken,
		},
		{
			input: "123456789,00000012",
			desc:  "Second item contains just one item",
			err:   ErrInvalidLengthOfSubtoken,
		},
		{
			input: "123456789,000000123",
			desc:  "Should be fine",
			err:   nil,
		},
	} {
		_, err := ParseDrawbarPositions(test.input)

		if !errors.Is(err, test.err) {
			t.Errorf("Test %s shuld have error %v\ngot\n%v", test.desc, test.err, err)
		}
	}
}

func TestParseDrawbarPositions(t *testing.T) {
	input := "352000000,003421000"
	result, err := ParseDrawbarPositions(input)

	if err != nil {
		t.Errorf("Unexpected error %s\n", err)
	}

	want := []DrawbarSetting{
		{3, 5, 2, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 4, 2, 1, 0, 0, 0},
	}

	if len(result) != len(want) {
		t.Errorf("Wanted %d got %d\n", len(want), len(result))
	}

	for i := range want {
		if slices.Compare(result[i][:], want[i][:]) != 0 {
			t.Errorf("Wanted\n%v\ngot\n%v\n", want, result)
			break
		}
	}

}
