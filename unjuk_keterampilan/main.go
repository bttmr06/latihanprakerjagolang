package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shopspring/decimal"
)

var db *gorm.DB
var err error

// pegawai is a representation of a pegawai
type Pegawai struct {
	ID   int             `form:"id" json:"id"`
	NIP  string          `form:"nip" json:"nip"`
	Nama string          `form:"nama" json:"nama"`
	Gaji decimal.Decimal `form:"gaji" json:"gaji" sql:"type:decimal(16,2);"`
}

// Result is an array of pegawai
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Main
func main() {
	db, err = gorm.Open("mysql", "root:12345678@/golang_restapi_prakerja?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection failed", err)
	} else {
		log.Println("Connection established")
	}

	db.AutoMigrate(&Pegawai{})
	handleRequests()
}

func handleRequests() {
	log.Println("Start the development server at http://127.0.0.1:8888")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		res := Result{Code: 404, Message: "Method not found"}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	myRouter.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		res := Result{Code: 403, Message: "Method not allowed"}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/pegawai", createPegawai).Methods("POST")
	myRouter.HandleFunc("/api/pegawai", getPegawai).Methods("GET")
	myRouter.HandleFunc("/api/pegawai/{id}", getPegawaiSingle).Methods("GET")
	myRouter.HandleFunc("/api/pegawai/{id}", updatePegawai).Methods("PUT")
	myRouter.HandleFunc("/api/pegawai/{id}", deletePegawai).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8888", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func createPegawai(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var pegawai Pegawai
	json.Unmarshal(payloads, &pegawai)

	db.Create(&pegawai)

	res := Result{Code: 200, Data: pegawai, Message: "Success create pegawai"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getPegawai(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: get pegawai")

	pegawai := []Pegawai{}
	db.Find(&pegawai)

	res := Result{Code: 200, Data: pegawai, Message: "Success get pegawai"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func getPegawaiSingle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pegawaiID := vars["id"]

	var pegawai Pegawai

	db.First(&pegawai, pegawaiID)

	res := Result{Code: 200, Data: pegawai, Message: "Success get pegawai"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func updatePegawai(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pegawaiID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var pegawaiUpdates Pegawai
	json.Unmarshal(payloads, &pegawaiUpdates)

	var pegawai Pegawai
	db.First(&pegawai, pegawaiID)
	db.Model(&pegawai).Updates(pegawaiUpdates)

	res := Result{Code: 200, Data: pegawai, Message: "Success update Pegawai"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func deletePegawai(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pegawaiID := vars["id"]

	var pegawai Pegawai

	db.First(&pegawai, pegawaiID)
	db.Delete(&pegawai)

	res := Result{Code: 200, Message: "Success delete pegawai"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
