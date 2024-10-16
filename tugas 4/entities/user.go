package entities

type LoginRequest struct {
    Username string `json:"username" form:"username" binding:"required"`
    Password string `json:"password" form:"password" binding:"required"`
}

type User struct {
	ID        uint `gorm:"primarykey" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
func (User) TableName() string {
	return "users"
}
