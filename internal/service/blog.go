package service

type BlogRequest struct {
	Name    string `form:"name" binding:"max=100"`
	Summary string `form:"summary" binding:"max=10000"`
	Content string `form:"content"`
}
