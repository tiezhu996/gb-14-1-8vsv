package model

// TrialCaseResult 试运行单用例结果：只包含实际输出与耗时，
// 不带期望输出、不判定对错（与正式评测的 JudgeResult 区分）。
type TrialCaseResult struct {
	TestCaseIndex int    `json:"test_case_index"`
	Input         string `json:"input"`
	Actual        string `json:"actual"`
	RuntimeMs     int64  `json:"runtime_ms"`
	ErrorMessage  string `json:"error_message"`
}
