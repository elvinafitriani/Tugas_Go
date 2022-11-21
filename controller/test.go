package controller

import (
	"crud/connection"
	"crud/models"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	// "github.com/jinzhu/gorm"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = connection.ConnectToDb()
}

func Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Barang
		DB.Find(&data)

		datajson, err := json.Marshal(data) //proses encode ke json
		if err != nil {
			http.Error(w, "Error Enconde to JSON", 500)
			return //fungsinya menghentikan fungsi itu
		}

		w.Header().Set("Content-Type", "application/json") //supaya output berupa data json
		w.Write(datajson)
		w.WriteHeader(200)
		return //fungsinya untuk menghentikan fungsi itu
	}
	http.Error(w, "Error Not Found", 404)
}

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Barang
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode Json", 500)
			return
		}

		DB.Create(&data)
		w.Write([]byte("Sukses Post Data"))
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func Getlimits(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Barang
		DB.Limit(5).Find(&data)

		dataJson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Conten-Type", "application/json")
		w.Write(dataJson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		var data models.Barang
		url := r.URL.String()
		var id []string = strings.Split(url, "/")
		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		decode := json.NewDecoder(r.Body)
		errors := decode.Decode(&data)

		if errors != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		err := DB.First(&models.Barang{}, "id = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		DB.Find(&models.Barang{}, "id=?", id[2]).Updates(&data)
		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	return
}

func PostJual(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Jual
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode Json", 500)
			return
		}

		cek := DB.Create(&data).Error
		if cek != nil {
			http.Error(w, "Tidak Ada Id Barang", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Sukses Post Data"))
		w.WriteHeader(http.StatusCreated)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func GetJual(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Jual
		// DB.Find(&data)
		DB.Model(&models.Jual{}).Preload("Barang").Find(&data)

		datajson, err := json.Marshal(data) //proses encode ke json
		if err != nil {
			http.Error(w, "Error Enconde to JSON", 500)
			return //fungsinya menghentikan fungsi itu
		}

		w.Header().Set("Content-Type", "application/json") //supaya output berupa data json
		w.Write(datajson)
		w.WriteHeader(200)
		return //fungsinya untuk menghentikan fungsi itu
	}
	http.Error(w, "Error Not Found", 404)
}

func View(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		url := r.URL.String()
		var inp []string = strings.Split(url, "/")
		put, err := strconv.Atoi(inp[2])
		type tamp struct {
			Id      int    `json:"id"`
			NamaBrg string `json:"nama"`
			Harga   int    `json:"harga"`
			IdKat   int    `json:"kategori_id"`
			Nama    string `json:"nama_kategori"`
			IdJual  int    `json:"id_jual"`
		}
		var buff []tamp
		DB.Model(&models.Barang{}).Select("barangs.id, barangs.nama_brg, barangs.harga, kategoris.id_kat, kategoris.nama, juals.id_jual").Joins("inner join kategoris on kategoris.id_kat = barangs.kategori_id inner join juals on juals.barang_id = barangs.id").Where(&models.Barang{Id: put}).Scan(&buff)
		datajson, err := json.Marshal(buff)
		if err != nil {
			http.Error(w, "Error Decode Json", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json") //supaya output berupa data json
		w.Write(datajson)
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, "Error Not Found", http.StatusNotFound)
}
