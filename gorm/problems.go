package gorm

// Problem : User generated problem
type Problem struct {
	Title       string
	Description string
	SubProblems []Problem
	Comments    []Comment
}
