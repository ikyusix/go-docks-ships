package usecase

import (
	"errors"
	"fmt"
	"godoksip/docks"
	"godoksip/model"
)

type DocksUsecaseImpl struct {
	docksRepo docks.DocksRepo
}

func CreateDocksUsecase(docksRepo docks.DocksRepo) docks.DocksUsecase {
	return &DocksUsecaseImpl{docksRepo}
}
func (d *DocksUsecaseImpl) GetByIdDocks(id int) (*model.Docks, error) {
	return d.docksRepo.GetByIdDocks(id)
}
func (d *DocksUsecaseImpl) InsertDocks(docks *model.Docks) error {
	docksVal, err := d.docksRepo.GetByIdDocks(docks.ID)
	if err != nil {
		return fmt.Errorf("[CreateDocksUsecase.InsertDocks]: %w", err)
	}
	fmt.Println(docksVal)

	if docksVal != nil {
		return errors.New("[CreateDocksUsecase.InsertDocks]")
	}
	return d.docksRepo.InsertDocks(docks)
}
//func (s *ShipsUsecaseImpl) GetAllShips() (*[]model.Ships, error) {
//	return s.shipsRepo.GetAllShips()
//}