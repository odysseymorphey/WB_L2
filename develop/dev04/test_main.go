package main

import (
	"strings"
	"testing"
)

func TestAnagram(t *testing.T) {
	got := SearchAnagrams([]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стиль"})
	want := make(map[string][]string)
	want["акптя"] = []string{"пятак", "пятка", "тяпка"}
	want["иклост"] = []string{"листок", "слиток", "столик"}

	if len(want) != len(got) {
		t.Errorf("invalid size")
	}

	for key, valueInGot := range want {
		valueInWant, exist := got[key]
		if !exist {
			t.Errorf("the key %s is missing in got", key)
		}
		if len(valueInGot) != len(valueInWant) {
			t.Errorf("the size of slice in got is %d, and in want is %d", len(valueInGot), len(valueInWant))
		}
		for i := 0; i < len(valueInWant); i++ {
			if strings.Compare(valueInGot[i], valueInWant[i]) != 0 {
				t.Errorf("the strings are different")
			}
		}
	}
}