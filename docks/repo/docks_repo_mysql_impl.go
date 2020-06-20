package repo

import (
	"database/sql"
	_ "encoding/json"
	"errors"
	"fmt"
	"godoksip/docks"
	"godoksip/model"
	_ "net/http"
)
//var _ ships.ShipsRepo = &ShipsRepoMysqlImpl{}
type DocksRepoMysqlImpl struct {
	db *sql.DB
}
func CreateDocksRepoMysqlImpl(db *sql.DB) docks.DocksRepo {
	return &DocksRepoMysqlImpl{db}
}
func (s *DocksRepoMysqlImpl) GetByIdDocks(id int) (*model.Docks, error) {
	qry := "SELECT id, docks_number, docks_status FROM docks where id = ?"
	docks := model.Docks{}
	err := s.db.QueryRow(qry, id).Scan(&docks.ID, &docks.DocksNumber, &docks.DocksStatus)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[DocksRepoMysqlImpl.GetByIdShips]: %w", err)
	}
	return &docks, nil
}
func (s *DocksRepoMysqlImpl) InsertDocks(docks *model.Docks) error {
	qry := "INSERT INTO docks(id, docks_number, docks_status) VALUES (?, ?, ?)"
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("[DocksRepoMysqlImpl.InsertDocks]: %w", err)
	}
	_, err = tx.Exec(qry, docks.ID, docks.DocksNumber, docks.DocksStatus)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[DocksRepoMysqlImpl.InsertDocks]: '"+qry+"' : %w", err)
	}
	tx.Commit()
	return nil
}
//func (s *ShipsRepoMysqlImpl) UpdateShips(id int) (*model.Ships, error) {
//	//qry := "UPDATE ships SET ships_type = ?, ships_weight = ? WHERE id = ?"
//	//tx, err := s.db.Begin()
//	//if err != nil {
//	//	return fmt.Errorf("[ShipsRepoMysqlImpl.InsertShips]: %w", err)
//	//}
//	//_, err = db.Exec(qry, ship.ShipsType, ship.ShipsWeight, val["id"])
//	//if err != nil {
//	//	tx.Rollback()
//	//	return fmt.Errorf("[ShipsRepoMysqlImpl.InsertShips]: '"+qry+"' : %w", err)
//	//}
//	//tx.Commit()
//	//return nil
//
//}
//func (s *ShipsRepoMysqlImpl) GetAllShips() (*[]model.Ships, error) {
//	qry := "SELECT id, ships_type, ships_weight FROM ships"
//	rows, err := s.db.Query(qry)
//	var result []model.Ships
//	if err != nil {
//		return nil, fmt.Errorf("[ShipsRepoMysqlImpl.GetAllShips]: %w", err)
//	}
//	defer rows.Close()
//	ships := model.Ships{}
//	for rows.Next() {
//		rows.Scan(&ships.ID, &ships.ShipsType, &ships.ShipsWeight)
//		result = append(result, ships)
//	}
//}
//func (s *ShipsRepoMysqlImpl) DeleteShips()