package model

type ErrorCode int

const (
	Success       ErrorCode = 0
	UserExists    ErrorCode = 301
	UserNotExists ErrorCode = 302
)
