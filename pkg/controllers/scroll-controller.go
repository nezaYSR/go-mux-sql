package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nezaYSR/go-mux-sql/pkg/models"
	"github.com/nezaYSR/go-mux-sql/pkg/utils"
)

var NewScroll models.Scroll

func GetScroll(w http.ResponseWriter, r *http.Request) {
	newScroll := models.GetAllScrolls()
	res, _ := json.Marshal(newScroll)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetScrollById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	scrollId := vars["scrollId"]
	id, err := strconv.ParseInt(scrollId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	scrollDetails, _ := models.GetScrollById(id)
	res, _ := json.Marshal(scrollDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateScroll(w http.ResponseWriter, r *http.Request) {
	prepareTo := &models.Scroll{}
	utils.ParseBody(r, prepareTo)
	s, err := prepareTo.CreateScroll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(s)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteScroll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	scrollId := vars["scrollId"]
	id, err := strconv.ParseInt(scrollId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// deleteScroll := models.DeleteScroll(id)
	// res, _ := json.Marshal(deleteScroll)
	models.DeleteScroll(id)

	response := map[string]bool{"success": true}
	res, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateScroll(w http.ResponseWriter, r *http.Request) {
	var updateScroll = &models.Scroll{}
	utils.ParseBody(r, updateScroll)
	vars := mux.Vars(r)
	scrollId := vars["scrollId"]
	id, err := strconv.ParseInt(scrollId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	scrollDetails, db := models.GetScrollById(id)
	if updateScroll.Title != "" {
		scrollDetails.Title = updateScroll.Title
	}
	if updateScroll.MagicianName != "" {
		scrollDetails.MagicianName = updateScroll.MagicianName
	}
	if updateScroll.Element != "" {
		scrollDetails.Element = updateScroll.Element
	}
	if updateScroll.Rarity != "" {
		scrollDetails.Rarity = updateScroll.Rarity
	}
	if updateScroll.Price >= 0 {
		scrollDetails.Price = updateScroll.Price
	}

	db.Save(&scrollDetails)
	res, _ := json.Marshal(scrollDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
