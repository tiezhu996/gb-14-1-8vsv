package service

import (
	"context"
	"os/exec"
	"testing"

	"github.com/blueship581/codelearn/internal/constants"
	"github.com/blueship581/codelearn/internal/model"
	"github.com/blueship581/codelearn/internal/util"
)

func newTestJudge(t *testing.T) *JudgeService {
	t.Helper()
	return NewJudgeService(util.NewLogger("error"), constants.DefaultJudgeTimeout)
}

func hasPython() bool {
	_, err := exec.LookPath("python3")
	return err == nil
}

func TestJudgePythonAccepted(t *testing.T) {
	if !hasPython() {
		t.Skip("python3 not available")
	}
	j := newTestJudge(t)
	code := "a, b = map(int, input().split())\nprint(a + b)"
	tcs := []model.TestCase{
		{Input: "1 2", Output: "3"},
		{Input: "10 20", Output: "30"},
	}
	results, status, score, _, _ := j.Judge(context.Background(), constants.LanguagePython, code, tcs, 10)
	if status != constants.SubmissionAccepted {
		t.Errorf("status = %q, want accepted; results=%+v", status, results)
	}
	if score != 100 {
		t.Errorf("score = %d, want 100", score)
	}
	if len(results) != 2 || !results[0].Passed || !results[1].Passed {
		t.Errorf("results not all passed: %+v", results)
	}
}

func TestJudgePythonPartial(t *testing.T) {
	if !hasPython() {
		t.Skip("python3 not available")
	}
	j := newTestJudge(t)
	code := "a, b = map(int, input().split())\nprint(a + b)"
	tcs := []model.TestCase{
		{Input: "1 2", Output: "3"},
		{Input: "10 20", Output: "999"},
	}
	_, status, score, _, _ := j.Judge(context.Background(), constants.LanguagePython, code, tcs, 10)
	if status != constants.SubmissionPartial {
		t.Errorf("status = %q, want partial", status)
	}
	if score != 50 {
		t.Errorf("score = %d, want 50", score)
	}
}

func TestJudgePythonRuntimeError(t *testing.T) {
	if !hasPython() {
		t.Skip("python3 not available")
	}
	j := newTestJudge(t)
	code := "raise ValueError('boom')"
	tcs := []model.TestCase{{Input: "", Output: "3"}}
	_, status, _, _, errMsg := j.Judge(context.Background(), constants.LanguagePython, code, tcs, 10)
	if status != constants.SubmissionRuntimeError {
		t.Errorf("status = %q, want runtime_error", status)
	}
	if errMsg == "" {
		t.Error("expected error message")
	}
}

func TestJudgePythonTimeout(t *testing.T) {
	if !hasPython() {
		t.Skip("python3 not available")
	}
	j := newTestJudge(t)
	code := "while True:\n    pass"
	tcs := []model.TestCase{{Input: "", Output: "3"}}
	_, status, _, _, _ := j.Judge(context.Background(), constants.LanguagePython, code, tcs, 1)
	if status != constants.SubmissionTimeout {
		t.Errorf("status = %q, want timeout", status)
	}
}

func TestRunSamplesPython(t *testing.T) {
	if !hasPython() {
		t.Skip("python3 not available")
	}
	j := newTestJudge(t)
	code := "a, b = map(int, input().split())\nprint(a + b)"
	results, total := j.RunSamples(context.Background(), constants.LanguagePython, code, []string{"1 2", "10 20"}, 10)
	if len(results) != 2 {
		t.Fatalf("results len = %d, want 2", len(results))
	}
	if got := util.NormalizeOutput(results[0].Actual); got != "3" {
		t.Errorf("results[0].Actual = %q, want %q", got, "3")
	}
	if got := util.NormalizeOutput(results[1].Actual); got != "30" {
		t.Errorf("results[1].Actual = %q, want %q", got, "30")
	}
	for i, r := range results {
		if r.ErrorMessage != "" {
			t.Errorf("results[%d].ErrorMessage = %q, want empty", i, r.ErrorMessage)
		}
	}
	if total <= 0 {
		t.Errorf("total runtime = %d, want > 0", total)
	}
}

func TestRunSamplesContinuesAfterError(t *testing.T) {
	if !hasPython() {
		t.Skip("python3 not available")
	}
	j := newTestJudge(t)
	// 第一条输入非法触发运行错误，第二条仍应继续运行。
	code := "n = int(input())\nprint(n * 2)"
	results, _ := j.RunSamples(context.Background(), constants.LanguagePython, code, []string{"abc", "21"}, 10)
	if len(results) != 2 {
		t.Fatalf("results len = %d, want 2", len(results))
	}
	if results[0].ErrorMessage == "" {
		t.Error("results[0].ErrorMessage empty, want runtime error")
	}
	if got := util.NormalizeOutput(results[1].Actual); got != "42" {
		t.Errorf("results[1].Actual = %q, want %q", got, "42")
	}
}

func TestCalcScore(t *testing.T) {
	tests := []struct {
		name   string
		passed int
		total  int
		want   int
	}{
		{name: "all pass", passed: 3, total: 3, want: 100},
		{name: "half", passed: 1, total: 2, want: 50},
		{name: "zero total", passed: 0, total: 0, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcScore(tt.passed, tt.total); got != tt.want {
				t.Errorf("calcScore(%d,%d) = %d, want %d", tt.passed, tt.total, got, tt.want)
			}
		})
	}
}
