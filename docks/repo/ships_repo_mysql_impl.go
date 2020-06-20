package repo

import (
	"database/sql"
	_ "encoding/json"
	"errors"
	"fmt"
	"godoksip/model"
	"godoksip/ships"
	_ "net/http"
)
//var _ ships.ShipsRepo = &ShipsRepoMysqlImpl{}
type ShipsRepoMysqlImpl struct {
	db *sql.DB
}
func CreateShipsRepoMysqlImpl(db *sql.DB) ships.ShipsRepo {
	return &ShipsRepoMysqlImpl{db}
}
func (s *ShipsRepoMysqlImpl) GetByIdShips(id int) (*model.Ships, error) {
	qry := "SELECT id, ships_type, ships_weight FROM ships where id = ?"
	ships := model.Ships{}
	err := s.db.QueryRow(qry, id).Scan(&ships.ID, &ships.ShipsType, &ships.ShipsWeight)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[ShipsRepoMysqlImpl.GetByIdShips]: %w", err)
	}
	return &ships, nil
}
func (s *ShipsRepoMysqlImpl) InsertShips(ship *model.Ships) error {
	qry := "INSERT INTO ships(id, ships_type, ships_weight) VALUES (?, ?, ?)"
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("[ShipsRepoMysqlImpl.InsertShips]: %w", err)
	}
	_, err = tx.Exec(qry, ship.ID, ship.ShipsType, ship.ShipsWeight)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[ShipsRepoMysqlImpl.InsertShips]: '"+qry+"' : %w", err)
	}
	tx.Commit()
	return nil
}
func (s *ShipsRepoMysqlImpl) GetAllShips() (*[]model.Ships, error) {
	qry := "SELECT id, ships_type, ships_weight FROM ships"
	ships := model.Ships{}
	err := s.db.QueryRow(qry).Scan(&ships.ID, &ships.ShipsType, &ships.ShipsWeight)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[ShipsRepoMysqlImpl.GetByIdShips]: %w", err)
	}
	return &[]model.Ships{}, nil
}
//func (s *ShipsRepoMysqlImpl) UpdateShips(id int) (*model.Ships, error) {
//	qry := "UPDATE ships SET ships_type = ?, ships_weight = ? WHERE (id = ?)"
//	tx, err := s.db.Begin()
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