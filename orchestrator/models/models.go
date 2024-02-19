package models

import (
	"gorm.io/gorm"
)

type (
	Expression struct {
		gorm.Model
		RequestID    string `json:"request_id" gorm:"text;not null;default:null"`
		ExpressionId string `json:"expression_id" gorm:"text;not null;default:null"`
		Expression   string `json:"expression" gorm:"text;not null;default:null"`
		State        string `json:"state" gorm:"text;not null;default:null"`
		Result       string `json:"result" gorm:"text;not null;default:null"`
		Message      string `json:"message" gorm:"text;not null;default:null"`
	}
)
