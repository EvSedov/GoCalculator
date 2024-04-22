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
	err = s.DB.Table("expressions").Where("email = ?", email).Find(&dbExpressions).Error
	if err != nil {
		return nil, err
	}

	expressions, err = json.Marshal(dbExpressions)
	if err != nil {
		return nil, err
	}

	return expressions, nil
}

func (s *storage) GetToCalculate() (expression []byte, err error) {
	var dbExpression entities.Expression
	err = s.DB.Table("expressions").Where("state = ?", "valid").First(&dbExpression).Error
	if err != nil {
		return nil, err
	}

	err = s.DB.Table("expressions").Where("expression_id = ?", dbExpression.ExpressionId).Update("state", "in_progress").Error
	if err != nil {
		return nil, err
	}

	expression, err = json.Marshal(dbExpression)
	if err != nil {
		return nil, err
	}

	return expression, nil
}

func (s *storage) UpdateCalculated(expression *entities.Expression) (err error) {

	if err = s.DB.Where("expression_id = ?", expression.ExpressionId).Updates(&expression).Error; err != nil {
		return err
	}

	return nil
}
