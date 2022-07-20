package domain

import "time"

type User struct {
	ID          string    `json:"id" firestore:"id" gorm:"column:id; type:string; primaryKey;"`
	Identity    string    `json:"identity,omitempty" map:"identity" validate:"lte=30" firestore:"identity"`
	Email       string    `json:"email,omitempty" map:"email" validate:"lte=256" firestore:"email"`
	FullName    string    `json:"full_name,omitempty" map:"full_name" firestore:"full_name"`
	BirthPlace  string    `json:"birth_place,omitempty" map:"birth_place" firestore:"birth_place"`
	BirthDate   string    `json:"birth_date,omitempty" map:"birth_date" firestore:"birth_date"`
	FullAddress string    `json:"full_address,omitempty" map:"full_address" firestore:"full_address"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at" firestore:"created_at"`
}

type UserDetail struct {
	ID        string         `json:"id,omitempty" map:"id" firestore:"id"`
	Form      UserDetailForm `json:"form,omitempty" map:"form" firestore:"form"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at" firestore:"created_at"`
}

type UserDetailForm struct {
	Personal User `json:"personal,omitempty" map:"personal" validate:"required" firestore:"personal"`
}
