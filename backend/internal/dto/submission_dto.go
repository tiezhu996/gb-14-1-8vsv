package dto

import (
	"github.com/blueship581/codelearn/internal/model"
)

// SubmitRequest 提交代码请求。
type SubmitRequest struct {
	Language string `json:"language" binding:"required,oneof=python javascript java"`
	Code     string `json:"code" binding:"required,min=1,max=20000"`
}

// TestRunRequest 试运行请求（与提交同构，但不落提交记录）。
type TestRunRequest struct {
	Language string `json:"language" binding:"required,oneof=python javascript java"`
	Code     string `json:"code" binding:"required,min=1,max=20000"`
}

// TestRunResultResponse 单条示例试运行结果（不含期望输出、不判对错）。
type TestRunResultResponse struct {
	Index        int    `json:"index"`
	Input        string `json:"input"`
	Actual       string `json:"actual"`
	RuntimeMs    int64  `json:"runtime_ms"`
	ErrorMessage string `json:"error_message"`
}

// TestRunResponse 试运行响应：每条示例的实际输出与耗时。
type TestRunResponse struct {
	ProblemID      string                  `json:"problem_id"`
	Language       string                  `json:"language"`
	Results        []TestRunResultResponse `json:"results"`
	TotalRuntimeMs int64                   `json:"total_runtime_ms"`
	TodayTestRuns  int64                   `json:"today_test_runs"`
}

// ToTestRunResponse 将试运行结果转换为响应。
func ToTestRunResponse(problemID, language string, results []model.TestRunResult, totalRuntimeMs, todayTestRuns int64) TestRunResponse {
	out := make([]TestRunResultResponse, 0, len(results))
	for _, r := range results {
		out = append(out, TestRunResultResponse{
			Index:        r.Index,
			Input:        r.Input,
			Actual:       r.Actual,
			RuntimeMs:    r.RuntimeMs,
			ErrorMessage: r.ErrorMessage,
		})
	}
	return TestRunResponse{
		ProblemID:      problemID,
		Language:       language,
		Results:        out,
		TotalRuntimeMs: totalRuntimeMs,
		TodayTestRuns:  todayTestRuns,
	}
}

// JudgeResultResponse 单个测试用例评测结果。
type JudgeResultResponse struct {
	TestCaseIndex int    `json:"test_case_index"`
	Input         string `json:"input"`
	Expected      string `json:"expected"`
	Actual        string `json:"actual"`
	Passed        bool   `json:"passed"`
	ErrorMessage  string `json:"error_message"`
}

// SubmissionResponse 提交记录响应。
type SubmissionResponse struct {
	ID            string                `json:"id"`
	UserID        string                `json:"user_id"`
	Username      string                `json:"username"`
	ProblemID     string                `json:"problem_id"`
	ProblemTitle  string                `json:"problem_title"`
	Language      string                `json:"language"`
	Code          string                `json:"code"`
	Status        string                `json:"status"`
	Score         int                   `json:"score"`
	PointsAwarded int64                 `json:"points_awarded"`
	RuntimeMs     int64                 `json:"runtime_ms"`
	Results       []JudgeResultResponse `json:"results"`
	ErrorMessage  string                `json:"error_message"`
	CreatedAt     string                `json:"created_at"`
}

// ToSubmissionResponse 将模型转换为响应。
func ToSubmissionResponse(s *model.Submission) SubmissionResponse {
	results := make([]JudgeResultResponse, 0, len(s.Results))
	for _, r := range s.Results {
		results = append(results, JudgeResultResponse{
			TestCaseIndex: r.TestCaseIndex,
			Input:         r.Input,
			Expected:      r.Expected,
			Actual:        r.Actual,
			Passed:        r.Passed,
			ErrorMessage:  r.ErrorMessage,
		})
	}
	return SubmissionResponse{
		ID:            s.ID.Hex(),
		UserID:        s.UserID.Hex(),
		Username:      s.Username,
		ProblemID:     s.ProblemID.Hex(),
		ProblemTitle:  s.ProblemTitle,
		Language:      s.Language,
		Code:          s.Code,
		Status:        s.Status,
		Score:         s.Score,
		PointsAwarded: s.PointsAwarded,
		RuntimeMs:     s.RuntimeMs,
		Results:       results,
		ErrorMessage:  s.ErrorMessage,
		CreatedAt:     s.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
