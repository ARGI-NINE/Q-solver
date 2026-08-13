package solution

import "testing"

func TestExtractLeetCodeSolution(t *testing.T) {
	input := "说明\n```cpp\nclass Solution {\npublic:\n    int answer() { return 42; }\n};\n```\n更多说明"
	want := "class Solution {\npublic:\n    int answer() { return 42; }\n};"
	got, err := extractLeetCodeSolution(input)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestExtractLeetCodeSolutionRejectsACM(t *testing.T) {
	_, err := extractLeetCodeSolution("int main() { return 0; }")
	if err == nil {
		t.Fatal("expected non-LeetCode output to be rejected")
	}
}
