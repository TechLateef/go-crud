package entities

//Book struct represent books table in database
type ToDoModels struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"userid"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	UserID      uint64 `gorm:"not null" json:"usersid"`
	Users       Users  `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
