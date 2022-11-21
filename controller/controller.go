package controller

import (
	"encoding/json"
	"net/http"
	"relation/connection"
	"relation/model"
	"strings"

	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	Db = connection.Connect()
}

func Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []model.Barang
		Db.Find(&data)

		dataJson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Can't Encoding Data", 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Method Not Found", 404)
	return
}

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data []model.Barang
		decode := json.NewDecoder(r.Body).Decode(&data)

		if decode != nil {
			http.Error(w, "Can't Decode Data", 500)
			return
		}

		Db.Create(&data)
		w.Write([]byte("Post Data Success"))
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Methhod Not Found", 404)
	return
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		url := r.URL.String()
		var data model.Barang
		var id []string = strings.Split(url, "/")

		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := Db.First(&model.Barang{}).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		Db.Where("id_bar = ?", id[2]).Updates(&data)
		w.Write([]byte("Updates Data Success"))
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		url := r.URL.String()
		var data []model.Barang
		var id []string = strings.Split(url, "/")

		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusFound)
			return
		}
		err := Db.First("id_bar = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		Db.Where("id_bar = ?", id[2]).Delete(&data)
		w.Write([]byte("Delete Success"))
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}

func GetJual(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []model.Jual
		Db.Model(&model.Jual{}).Preload("Barang").Find(&data)

		dataJson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Can't Encoding Data", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Method Not Found", 404)
	return
}

func PostJual(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data []model.Jual
		decode := json.NewDecoder(r.Body).Decode(&data)
		if decode != nil {
			http.Error(w, "Can't Decode Data", 500)
			return
		}
		Db.Create(&data)
		w.Write([]byte("Success Post Data"))
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Method Not Found", 404)
	return
}

func GetBarang(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// var data []model.Barang
		url := r.URL.String()
		var id []string = strings.Split(url, "/")
		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		var DataScan []model.Select
		Db.Table("barangs").Select("barangs.id_bar, barangs.nama_barang, barangs.harga, barangs.kategori_id, kategoris.nama_kat, juals.id_jual").Joins("inner join juals on juals.id_barang = barangs.id_bar").Joins("inner join kategoris on kategoris.id_kategori = barangs.kategori_id").Where("id_bar = ?", id[2]).Scan(&DataScan)
		dataJson, err := json.Marshal(DataScan)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}

func PostKategori(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data model.Kategori
		decode := json.NewDecoder(r.Body).Decode(&data)
		if decode != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if err := Db.Create(&data).Error; err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Post Data Success"))
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}
