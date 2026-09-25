// 提交评测 API
import request from '../utils/request'
import type { PageData, Submission, TestRunResponse } from '../types'

export function submitCode(problemId: string, payload: { language: string; code: string }) {
  return request.post<unknown, Submission>(`/problems/${problemId}/submit`, payload)
}

// 试运行：只回显实际输出与耗时，不判对错、不留提交记录
export function testRunCode(problemId: string, payload: { language: string; code: string }) {
  return request.post<unknown, TestRunResponse>(`/problems/${problemId}/test-run`, payload)
}

export function getSubmission(id: string) {
  return request.get<unknown, Submission>(`/submissions/${id}`)
}

export function listSubmissions(params: { page?: number; page_size?: number; problem_id?: string; status?: string }) {
  return request.get<unknown, PageData<Submission>>('/submissions', { params })
}
