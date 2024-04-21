package storage

import "github.com/evsedov/GoCalculator/orchestrator/entities"

func (s *storage) CreateNew(expression *entities.Expression) (err error) {
	if err = s.DB.Create(&expression).Error; err != nil {
		return err
	}

	return nil
}
