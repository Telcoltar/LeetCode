package Solution1768

import (
	"testing"
)

func TestSolution(t *testing.T) {
	result := solution("abc", "def")
	expected := "adbecf"
	if result != expected {
		t.Errorf("solution(\"abc\", \"def\") failed, got %s, expected %s", result, expected)
	}
}
