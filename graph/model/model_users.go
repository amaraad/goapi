package model

type User struct {
	ID        int64  `json:"id" gorm:"primary_key;auto_increment;not_null"`
	FirstName string `json:"firstname" gorm:type:varchar(120);not null`
	LastName  string `json:"lastname" gorm:type:varchar(120);not null`
	Email     string `json:"email" gorm:type:varchar(120);not null`
	Password  string `json:"password" gorm:type:varchar(120);not null`
}
