package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Error struct {
	Message string `json:"Error"`
}

type Item struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Amount int     `json:"amount"`
	Price  float64 `json:"price"`
}

var Items []Item

func main() {
	fmt.Println("API was started")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/items", GetAllItems).Methods("GET")
	myRouter.HandleFunc("/item/{id}", GetItemById).Methods("GET")

	myRouter.HandleFunc("/item", AddNewItem).Methods("POST")

	myRouter.HandleFunc("/item/{id}", UpdateItemById).Methods("PUT")

	myRouter.HandleFunc("/item/{id}", DeleteItemById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", myRouter))

}

func GetAllItems(writer http.ResponseWriter, request *http.Request) {
	if len(Items) < 1 {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(Error{Message: "No one items in stock!"})
	} else {
		json.NewEncoder(writer).Encode(Items)
	}
}

func GetItemById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])

	item := getItem(id)
	if item != nil {
		json.NewEncoder(writer).Encode(item)
	} else {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(Error{Message: "Item with that id not found!"})
	}
}

func getItem(id int) *Item {
	for _, item := range Items {
		if item.Id == id {
			return &item
		}
	}
	return nil
}

func AddNewItem(writer http.ResponseWriter, request *http.Request) {
	var item Item
	reqBody, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(reqBody, &item)

	Items = append(Items, item)
	writer.WriteHeader(http.StatusCreated)
}

func UpdateItemById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	reqBody, _ := ioutil.ReadAll(request.Body)
	id, _ := strconv.Atoi(vars["id"])

	if updateItemById(id, reqBody) {
		writer.WriteHeader(http.StatusAccepted)
	} else {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(Error{Message: "Item with that id not found!"})
	}
}

func updateItemById(id int, reqBody []byte) bool {
	for index, item := range Items {
		if item.Id == id {
			json.Unmarshal(reqBody, &Items[index])
			return true
		}
	}
	return false
}

func DeleteItemById(writer http.ResponseWriter, request *http.Request) {
	idStr := mux.Vars(request)["id"]
	id, _ := strconv.Atoi(idStr)
	if deleteItem(id) {
		writer.WriteHeader(http.StatusAccepted)
	} else {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(Error{Message: "Item with that id not found!"})
	}
}

func deleteItem(id int) bool {
	for index, item := range Items {
		if item.Id == id {
			Items = append(Items[:index], Items[index+1:]...)
			return true
		}
	}
	return false
}
