package model

type Model struct {
	ID        uint32  `json:"id" gorm:"primary_key"`
	CreatedAt float64 `json:"created_at"`
}

type Blog struct {
	Model
	UserId    string `json:"user_id"`
	UserName  string `json:"user_name"`
	UserImage string `json:"user_image"`
	Name      string `json:"name"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
}
