package storage

import (
	"encoding/json"

	"github.com/evsedov/GoCalculator/orchestrator/entities"
)

func (s *storage) CreateNew(expression *entities.Expression) (err error) {
	if err = s.DB.Create(&expression).Error; err != nil {
		return err
	}

	return nil
}

func (s *storage) GetExpressionsByEmail(email string) (expressions []byte, err error) {
	var dbExpressions []entities.Expression
	err = s.DB.Table("expression").Where("email = ?", email).Find(&dbExpressions).Error
	if err != nil {
		return nil, err
	}

	JSONExpressions, err := json.Marshal(dbExpressions)
	if err != nil {
		return nil, err
	}

	return JSONExpressions, nil
}
