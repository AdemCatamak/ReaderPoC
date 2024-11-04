package inheritanceSolution

import "testing"

func TestCaseInsensitiveReader_Get_KeyNotMatchBecauseOfCase_ResultShouldBeValue(t *testing.T) {
	r := NewCaseInsensitiveReader()

	result := r.Get("heLLo")

	if result != "world" {
		t.Errorf("result: %s", result)
	}
}

func TestCaseInsensitiveReader_Get_KeyMatch_ResultShouldBeValue(t *testing.T) {
	r := NewCaseInsensitiveReader()

	result := r.Get("hello")

	if result != "world" {
		t.Errorf("result: %s", result)
	}
}
