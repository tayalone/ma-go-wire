package domain

// Domain of Post
type Domain struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	UserID int    `gorm:"not null" json:"userId"`
	Title  string `gorm:"not null" json:"title"`
	Body   string `gorm:"not null" json:"body"`
}
