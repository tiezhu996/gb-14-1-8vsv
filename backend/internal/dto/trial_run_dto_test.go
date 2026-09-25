package dto

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/blueship581/codelearn/internal/model"
)

func TestTrialRunResponseNoExpectedOrPassed(t *testing.T) {
	results := []model.TrialCaseResult{
		{TestCaseIndex: 0, Input: "1 2", Actual: "3\n", RuntimeMs: 12},
		{TestCaseIndex: 1, Input: "10 20", Actual: "30\n", RuntimeMs: 9},
	}
	resp := ToTrialRunResponse("pid", "A+B", "python", results)
	raw, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	s := string(raw)
	if strings.Contains(s, "expected") {
		t.Errorf("trial response must not expose expected output: %s", s)
	}
	if strings.Contains(s, "passed") {
		t.Errorf("trial response must not judge pass/fail: %s", s)
	}
	if !strings.Contains(s, `"actual":"3\n"`) || !strings.Contains(s, `"runtime_ms":12`) {
		t.Errorf("response missing actual/runtime: %s", s)
	}
	if resp.TotalRuntime != 21 {
		t.Errorf("total runtime = %d, want 21", resp.TotalRuntime)
	}
}
