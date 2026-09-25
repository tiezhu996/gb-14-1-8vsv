package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// JudgeResult 单个测试用例的评测结果。
type JudgeResult struct {
	TestCaseIndex int    `bson:"test_case_index" json:"test_case_index"`
	Input         string `bson:"input" json:"input"`
	Expected      string `bson:"expected" json:"expected"`
	Actual        string `bson:"actual" json:"actual"`
	Passed        bool   `bson:"passed" json:"passed"`
	ErrorMessage  string `bson:"error_message" json:"error_message"`
}

// TestRunResult 试运行单条示例结果：只记录实际输出与耗时，不比对期望输出、不判对错。
// 与 JudgeResult 不同：该结构仅用于试运行接口响应，不落库。
type TestRunResult struct {
	Index        int    `json:"index"`
	Input        string `json:"input"`
	Actual       string `json:"actual"`
	RuntimeMs    int64  `json:"runtime_ms"`
	ErrorMessage string `json:"error_message"`
}

// Submission 提交评测实体：状态机 pending -> judging -> accepted/partial/runtime_error/timeout。
type Submission struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id"`
	Username     string             `bson:"username" json:"username"`
	ProblemID    primitive.ObjectID `bson:"problem_id" json:"problem_id"`
	ProblemTitle string             `bson:"problem_title" json:"problem_title"`
	Language     string             `bson:"language" json:"language"`
	Code         string             `bson:"code" json:"code"`
	Status       string             `bson:"status" json:"status"`
	Score        int                `bson:"score" json:"score"`
	PointsAwarded int64             `bson:"points_awarded" json:"points_awarded"`
	RuntimeMs    int64              `bson:"runtime_ms" json:"runtime_ms"`
	Results      []JudgeResult      `bson:"results" json:"results"`
	ErrorMessage string             `bson:"error_message" json:"error_message"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
}
