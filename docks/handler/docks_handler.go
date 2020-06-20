package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"godoksip/docks"
	"godoksip/model"
	_ "godoksip/ships"
	"io/ioutil"
	"net/http"
	"strconv"
)

type DocksHandler struct {
	docksUsecase docks.DocksUsecase
}

func CreateDocksHandler(r *mux.Router, docksUsecase docks.DocksUsecase) {
	docksHandler := DocksHandler{docksUsecase}
	r.HandleFunc("/docks", docksHandler.insertDocks).Methods(http.MethodPost)
	r.HandleFunc("/docks/{id}", docksHandler.getByIdDocks).Methods(http.MethodGet)
	//r.HandleFunc("/ships", shipsHandler.viewShips).Methods(http.MethodGet)
	//r.HandleFunc("/ships/{id}", shipsHandler.updateShips).Methods(http.MethodPut)
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

func (d *DocksHandler) insertDocks(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		handleError(resp, "Ooops, Something went wrong")
		fmt.Println("[DocksHandler.insertDocks]: " + err.Error())
		return
	}

	var docks = model.Docks{}
	err = json.Unmarshal(body, &docks)
	if err != nil {
		handleError(resp, "Ooops, Something went wrong")
		fmt.Println("[DocksHandler.insertDocks]: " + err.Error())
		return
	}

	err = d.docksUsecase.InsertDocks(&docks)
	if err != nil {
		handleError(resp, err.Error())
		fmt.Println("[DocksHandler.insertDocks]: " + err.Error())
		return
	}
	handleSuccess(resp, nil)
}
func (d *DocksHandler) getByIdDocks(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)
	strId := muxVar["id"]

	id, err := strconv.Atoi(strId)
	if err != nil {
		handleError(resp, "Id must number")
		return
	}

	docks, err := d.docksUsecase.GetByIdDocks(id)
	if err != nil {
		handleError(resp, err.Error())
		return
	}
	handleSuccess(resp, docks)
}
//func (s *ShipsHandler) viewShips(resp http.ResponseWriter, req *http.Request){
//	var result []TableStruct.Tables
//	rows, err := db.Query("Select * from table_res")
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//
//	for rows.Next() {
//		var each = TableStruct.Tables{}
//		var err = rows.Scan(&each.ID, &each.Status)
//		if err != nil {
//			fmt.Println(err.Error())
//			return
//		}
//		result = append(result, each)
//	}
//	defer rows.Close() // must close also for the rows
//
//	json, err := json.Marshal(result)
//	if err != nil {
//		resp.WriteHeader(http.StatusInternalServerError)
//		resp.Write([]byte("Ooops, something wrong"))
//		fmt.Println(err.Error())
//		return
//	}
//	resp.Header().Set("Content-Type", "application/json")
//	resp.Write(json)
//}
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