package domain

// Domain of Comment
type Domain struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	PostID uint   `gorm:"not null;foreignKey:PostID" json:"postId"`
	Name   string `gorm:"not null" json:"name"`
	Email  string `gorm:"not null" json:"email"`
	Body   string `gorm:"not null" json:"body"`
}
