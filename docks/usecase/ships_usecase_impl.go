package usecase

import (
	"errors"
	"fmt"
	"godoksip/model"
	"godoksip/ships"
)

type ShipsUsecaseImpl struct {
	shipsRepo ships.ShipsRepo
}

func CreateShipsUsecase(shipsRepo ships.ShipsRepo) ships.ShipsUsecase {
	return &ShipsUsecaseImpl{shipsRepo}
}

func (s *ShipsUsecaseImpl) GetByIdShips(id int) (*model.Ships, error) {
	return s.shipsRepo.GetByIdShips(id)
}

func (s *ShipsUsecaseImpl) GetAllShips() (*[]model.Ships, error) {
	return s.shipsRepo.GetAllShips()
}

func (s *ShipsUsecaseImpl) InsertShips(ships *model.Ships) error {
	shipsVal, err := s.shipsRepo.GetByIdShips(ships.ID)
	if err != nil {
		return fmt.Errorf("[CreateShipsUsecase.InsertShips]: %w", err)
	}
	fmt.Println(shipsVal)

	if shipsVal != nil {
		return errors.New("Ship id already exists, please enter another id")
	}
	return s.shipsRepo.InsertShips(ships)
}