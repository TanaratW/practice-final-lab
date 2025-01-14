package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model

	StudentID	string	`valid:"required~StudentID is required, matches(^[BMD]\\d{7}$)"`
	FirstName	string	`valid:"required~FirstName is required"`
	LastName	string	`valid:"required"`
	Email			string	`valid:"required~Email is required, email"`
	Phone			string	`valid:"required~PhoneNumber is required, matches(^[0]\\d{9}$)"`
	GenderID	uint		`valid:"required~Gender is required"`
	Gender		Gender	`gorm:"foreignKey:GenderID"`

}