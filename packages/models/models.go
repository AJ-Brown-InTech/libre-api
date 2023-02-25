package models

import (
	"time"

)

type Account struct{
	UserName 	string `json:"username" db:username`
	UserId 		int `json:"user_id" db:user_id`
	Uuid 		string `json:"uuid" db:"uuid"`
	Dob			string `json:"dob" db:"dob"`
	FirstName 	string `json:"first_name,omitempty" db:first_name`
	LastName 	string `json:"last_name,omitempty" db:last_name`
	Password 	string `json:"password" db:"password"`
	Email 		string `json:"email" db:"email"`
	Session 	time.Time `json:"session,omitempty" db:"session"`
	Rating 		int `json:"rating" db:"rating"`
	Post 		[]int `json:"post,omitempty" db:"post"`
	Active 		bool	`json:"active" db:"active"`
	Bio 		string `json:"bio,omitempty" db:"bio"`
	Link 		string `json:"link,omitempty" db:"bio"`
	Verfied 	bool `json:"verified" db:"verified"`
	CreadtedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at" db:"updated_at"`
}

type Followers struct{
	UserId 		int `json:"user_id" db:user_id`
	FollowerId 	int `json:"follower_id" db:"follower_id"`
}