package comparisonStrategySolution

import "testing"

func TestDefaultReader_Get_KeyNotMatchBecauseOfCase_ResultShouldBeEmpty(t *testing.T) {
	r := NewReader()

	result := r.Get("heLLo")

	if result != "" {
		t.Errorf("result: %s", result)
	}
}

func TestDefaultReader_Get_KeyMatch_ResultShouldBeValue(t *testing.T) {
	r := NewReader()
	result := r.Get("hello")

	if result != "world" {
		t.Errorf("result: %s", result)
	}
}

func TestDefaultReader_Get_KeyNotMatchBecauseOfCase_WithComparator_ResultShouldBeValue(t *testing.T) {
	insensitiveComparator := NewCaseInsensitiveComparator()
	r := NewReader().WithComparator(insensitiveComparator)

	result := r.Get("heLLo")

	if result != "world" {
		t.Errorf("result: %s", result)
	}
}

func TestDefaultReader_Get_KeyMatch_WithComparator_ResultShouldBeValue(t *testing.T) {
	insensitiveComparator := NewCaseInsensitiveComparator()
	r := NewReader().WithComparator(insensitiveComparator)

	result := r.Get("hello")

	if result != "world" {
		t.Errorf("result: %s", result)
	}
}
