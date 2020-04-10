package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model
type Memo struct {
	ID   uint
	Text string
}

func post(w http.ResponseWriter, r *http.Request) {
	// Read request body
	m := map[string]string{"text": ""}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&m); err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// Create record
	memo := Memo{Text: m["text"]}
	if err := db.Create(&memo).Error; err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// Encode json
	resp, err := json.Marshal(&memo)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// Write response
	w.Write(resp)
}

func get(w http.ResponseWriter, r *http.Request) {
	// Get all records
	memos := []Memo{}
	if err := db.Find(&memos).Error; err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// Encode json
	resp, err := json.Marshal(&memos)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// Write response
	w.Write(resp)
}

func put(w http.ResponseWriter, r *http.Request) {
	// Convert id from string to int
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// Read request body
	m := map[string]string{"text": ""}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&m); err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// Get recorde by id
	memo := Memo{}
	if err := db.First(&memo, id).Error; err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// Update memo item
	memo.Text = m["text"]
	if err := db.Save(&memo).Error; err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// Encode json
	resp, err := json.Marshal(&memo)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// Write response
	w.Write(resp)
}

func delete(w http.ResponseWriter, r *http.Request) {
	// Convert id from string to int
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// Delete record by id
	if err := db.Where("id = ?", id).Delete(Memo{}).Error; err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
