package entities


type Todo struct {
	ID      uint `gorm:"primarykey" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}
func (Todo) TableName() string {
	return "todos"
}