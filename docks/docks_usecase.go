package docks

import "godoksip/model"

type DocksUsecase interface {
	InsertDocks(docks *model.Docks) error
	GetByIdDocks(id int) (*model.Docks, error)
	//UpdateShips()
	//DeleteShips()
	//GetAllShips() (*[]model.Ships, error)
}
