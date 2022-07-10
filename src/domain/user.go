package domain

import "time"

type User struct {
	ID          int       `gorm:"column:id; type:int; primaryKey;" json:"id"`
	Address     string    `gorm:"column:address" json:"address"`
	BirthDate   string    `gorm:"column:birth_date" json:"birth_date"`
	BirthPlace  string    `gorm:"column:birth_place" json:"birth_place"`
	Education   int       `gorm:"column:education" json:"education"`
	Email       string    `gorm:"column:email" json:"email"`
	Firstname   string    `gorm:"column:firstname" json:"firstname"`
	Identity    string    `gorm:"column:identity" json:"identity"`
	Lastname    string    `gorm:"column:lastname" json:"lastname"`
	PhoneCode   string    `gorm:"column:phone_code" json:"phone_code"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phone_number"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
}
