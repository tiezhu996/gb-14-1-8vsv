package dto

import "github.com/blueship581/codelearn/internal/model"

// TrialRunRequest 试运行请求（与提交评测请求同构）。
type TrialRunRequest struct {
	Language string `json:"language" binding:"required,oneof=python javascript java"`
	Code     string `json:"code" binding:"required,min=1,max=20000"`
}

// TrialCaseResultResponse 试运行单用例结果：无期望输出、无对错判定。
type TrialCaseResultResponse struct {
	TestCaseIndex int    `json:"test_case_index"`
	Input         string `json:"input"`
	Actual        string `json:"actual"`
	RuntimeMs     int64  `json:"runtime_ms"`
	ErrorMessage  string `json:"error_message"`
}

// TrialRunResponse 试运行响应：每条示例的实际输出与耗时，不产生提交记录。
type TrialRunResponse struct {
	ProblemID    string                    `json:"problem_id"`
	ProblemTitle string                    `json:"problem_title"`
	Language     string                    `json:"language"`
	Results      []TrialCaseResultResponse `json:"results"`
	TotalRuntime int64                     `json:"total_runtime_ms"`
}

// ToTrialRunResponse 将试运行结果转换为响应。
func ToTrialRunResponse(problemID, problemTitle, language string, results []model.TrialCaseResult) TrialRunResponse {
	out := make([]TrialCaseResultResponse, 0, len(results))
	var total int64
	for _, r := range results {
		out = append(out, TrialCaseResultResponse{
			TestCaseIndex: r.TestCaseIndex,
			Input:         r.Input,
			Actual:        r.Actual,
			RuntimeMs:     r.RuntimeMs,
			ErrorMessage:  r.ErrorMessage,
		})
		total += r.RuntimeMs
	}
	return TrialRunResponse{
		ProblemID:    problemID,
		ProblemTitle: problemTitle,
		Language:     language,
		Results:      out,
		TotalRuntime: total,
	}
}
