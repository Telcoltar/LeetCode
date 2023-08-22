package Solution1071

import (
	"testing"
)

func TestSolution1(t *testing.T) {
	gcd := solution("ABCABC", "ABC")
	expected := "ABC"
	if gcd != expected {
		t.Errorf("solution(\"ABCABC\", \"ABC\") failed, got %s, expected %s", gcd, expected)
	}
}
