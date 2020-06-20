package ships

import "godoksip/model"

type ShipsRepo interface {
	InsertShips(ships *model.Ships) error
	GetByIdShips(id int) (*model.Ships, error)
	GetAllShips() (*[]model.Ships, error)
	//UpdateShips()
	//DeleteShips()
}
