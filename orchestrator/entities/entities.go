package entities

import (
	"gorm.io/gorm"
)

type (
	Expression struct {
		gorm.Model
		Email        string `json:"email" gorm:"text"`
		RequestID    string `json:"request_id" gorm:"text;not null;default:null"`
		ExpressionId string `json:"expression_id" gorm:"text;not null;default:null"`
		Expression   string `json:"expression" gorm:"text"`
		State        string `json:"state" gorm:"text"`
		Result       string `json:"result" gorm:"text"`
		Message      string `json:"message" gorm:"text"`
	}

	User struct {
		Id       uint   `json:"-" gorm:"primaryKey"`
		Name     string `json:"name"`
		Email    string `json:"email" gorm:"unique"`
		Password []byte
	}
)
