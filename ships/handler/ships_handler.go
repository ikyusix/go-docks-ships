package handler

import (
	_ "database/sql"
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"godoksip/model"
	"godoksip/ships"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ShipsHandler struct {
	shipsUsecase ships.ShipsUsecase
}

func CreateShipsHandler(r *mux.Router, shipsUsecase ships.ShipsUsecase) {
	shipsHandler := ShipsHandler{shipsUsecase}
	r.HandleFunc("/ships", shipsHandler.insertShips).Methods(http.MethodPost)
	r.HandleFunc("/ships", shipsHandler.getAllShips).Methods(http.MethodGet)
	//r.HandleFunc("/ships/{id}", shipsHandler.updateShips).Methods(http.MethodPut)
	r.HandleFunc("/ships/{id}", shipsHandler.getByIdShips).Methods(http.MethodGet)
	//r.HandleFunc("/ships/{id}", shipsHandler.deleteShips).Methods(http.MethodDelete)
}

func handleSuccess(resp http.ResponseWriter, data interface{}) {
	returnData := model.ResponseWrapper{
		Success:true,
		Message: "SUCCESS",
		Data: data,
	}
	jsonData, err := json.Marshal(returnData)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong"))
		fmt.Printf("[ShipsHandler.handleSuccess]: %v \n", err)
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
}
func handleError(resp http.ResponseWriter, message string) {
	data := model.ResponseWrapper{
		Success: false,
		Message: message,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong"))
		fmt.Printf("[MejaHandler.handleError]: %v \n", err)
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
}

func (s *ShipsHandler) insertShips(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		handleError(resp, "Ooops, Something went wrong")
		fmt.Println("[ShipsHandler.insertShips]: " + err.Error())
		return
	}

	var ships = model.Ships{}
	err = json.Unmarshal(body, &ships)
	if err != nil {
		handleError(resp, "Ooops, Something went wrong")
		fmt.Println("[ShipsHandler.insertShips]: " + err.Error())
		return
	}

	err = s.shipsUsecase.InsertShips(&ships)
	if err != nil {
		handleError(resp, err.Error())
		fmt.Println("[ShipsHandler.insertShips]: " + err.Error())
		return
	}
	handleSuccess(resp, nil)
}
func (s *ShipsHandler) getByIdShips(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)
	strId := muxVar["id"]

	id, err := strconv.Atoi(strId)
	if err != nil {
		handleError(resp, "Id must number")
		return
	}

	ships, err := s.shipsUsecase.GetByIdShips(id)
	if err != nil {
		handleError(resp, err.Error())
		return
	}
	handleSuccess(resp, ships)
}
func (s *ShipsHandler) getAllShips(resp http.ResponseWriter, req *http.Request){
	var result []model.Ships
	if req.Method == "GET" {
		for rows.Next() {
			var each = model.Ships{}
			var err = rows.Scan(&each.ID, &each.ShipsType, &each.ShipsWeight)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			result = append(result, each)
		}

		json, err := json.Marshal(result)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte("Ooops, something wrong"))
			fmt.Println(err.Error())
			return
		}
		resp.Header().Set("Content-Type", "application/json")
		resp.Write(json)
	} else {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		resp.Write([]byte("Unsupported method"))
		return
	}

}
//func (s *ShipsHandler) updateShips(resp http.ResponseWriter, req *http.Request){
//	muxVar := mux.Vars(req)
//	strId := muxVar["id"]
//
//	id, err := strconv.Atoi(strId)
//	if err != nil {
//		handleError(resp, "Id must number")
//		return
//	}
//
//	ships, err := s.shipsUsecase.GetByIdShips(id)
//	if err != nil {
//		handleError(resp, err.Error())
//		return
//	}
//	handleSuccess(resp, ships)
//}
//func (s *ShipsHandler) deleteShips(resp http.ResponseWriter, res *http.Request){}