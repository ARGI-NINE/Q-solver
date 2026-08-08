package domain

import (
	"strings"
	"testing"
)

func TestACMCppModeIsAvailable(t *testing.T) {
	found := false
	for _, category := range GetCategories() {
		for _, item := range category.Items {
			if item.ID == "dev-acm-cpp" {
				found = true
				if item.Label != "ACM C++ 输出" {
					t.Fatalf("unexpected ACM mode label: %q", item.Label)
				}
			}
		}
	}
	if !found {
		t.Fatal("ACM C++ mode is missing from domain categories")
	}
}

func TestACMCppPromptRequiresCompleteProgramWithoutSpeedRunRules(t *testing.T) {
	prompt := GetPrompt("dev-acm-cpp")
	for _, required := range []string{"GNU C++17", "int main()", "标准输入输出", "Markdown cpp 代码块"} {
		if !strings.Contains(prompt, required) {
			t.Errorf("ACM prompt is missing %q", required)
		}
	}
	for _, forbidden := range []string{"秒杀", "只输出 class Solution"} {
		if strings.Contains(prompt, forbidden) {
			t.Errorf("ACM prompt unexpectedly contains speed-run restriction %q", forbidden)
		}
	}
}
