package models

import "time"

type (
	Users struct {
		ID          int `json:"id" gorm:"primaryKey;autoIncrement"`
		Username    string `json:"username" gorm:"column:username;type:varchar(20);not null;unique"`
		Email       string `json:"email" gorm:"column:email;type:varchar(100);not null;unique"`
		PhoneNumber string `json:"phone_number" gorm:"column:phone_number;type:varchar(15);not null;unique"`
		FullName    string `json:"full_name" gorm:"column:full_name;type:varchar(255);not null"`
		Address     string `json:"address" gorm:"column:address;type:text;not null"`
		Dob         string `json:"dob" gorm:"column:dob;type:date;not null"`
		Password    string `json:"password" gorm:"column:password;type:varchar(255);not null"`
		CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;autoCreateTime"`
		UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;autoUpdateTime"`
	}

	UserSessions struct {
		ID          int `json:"id" gorm:"primaryKey;autoIncrement"`
		UserID      int `json:"user_id" gorm:"column:user_id;type:int;not null"`
		Token       string `json:"token" gorm:"column:token;type:varchar(255);not null"`
		RefreshToken string `json:"refresh_token" gorm:"column:refresh_token;type:varchar(255);not null"`
		TokenExpired time.Time `json:"token_expired" gorm:"column:token_expired;type:datetime;not null"`
		RefreshTokenExpired time.Time `json:"refresh_token_expired" gorm:"column:refresh_token_expired;type:datetime;not null"`
		CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;autoCreateTime"`
		UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;autoUpdateTime"`
	}
)