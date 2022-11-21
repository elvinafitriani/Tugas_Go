package controller

import (
	"encoding/json"
	"net/http"
	"pt9/connection"
	"pt9/models"
	"strconv"
	"strings"

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

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "error encode ke json ", 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "error not found ", 404)
}

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Barang
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}
		DB.Create(&data)
		w.Write([]byte("Sukses post data"))
		w.WriteHeader(200)
		return

	}
	http.Error(w, "Error Not Found ", 404)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		u := r.URL.String()
		var id []string = strings.Split(u, "/")
		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := DB.First(&models.Barang{}, "id = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Delete(&models.Barang{}, DB.Where("id = ?", id[2]))

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		u := r.URL.String()
		var id []string = strings.Split(u, "/")
		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var data models.Barang
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}

		err := DB.First(&models.Barang{}, "id = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		DB.Model(&models.Barang{}).Where("id = ? ", id[2]).Updates(&data)

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return

	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

}

func GetLimit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var nama []models.Barang
		DB.Limit(5).Find(&nama)
		datajson, err := json.Marshal(nama)
		if err != nil {
			http.Error(w, "error encode ke json ", 500)
			return
		}
		w.WriteHeader(200)
		w.Write(datajson)
		return

	}
}

func Detail(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		u := r.URL.String()
		var id []string = strings.Split(u, "/")
		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		var nama []models.Barang
		DB.Find(&nama, "id=?", id[2])
		datajson, err := json.Marshal(nama)
		if err != nil {
			http.Error(w, "error encode to json", 500)
			return
		}
		w.WriteHeader(200)
		w.Write(datajson)
		return
	}
}

func GetJual(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Jual
		DB.Model(&models.Jual{}).Preload("Barang").Find(&data)
		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "error encode ke json ", 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "error not found ", 404)
}

func PostJual(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Jual
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}
		cek := DB.Create(&data).Error
		if cek != nil {
			http.Error(w, "Id barang tidak ada", 500)
		}

		w.Write([]byte("Sukses post data"))
		w.WriteHeader(http.StatusCreated)
		return

	}
	http.Error(w, "Error Not Found ", 404)
}

func GetKategori(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Kategori

		DB.Model(&models.Kategori{}).Preload("Barang").Find(&data)
		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "error encode ke json ", 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "error not found ", 404)
}

func PostKategori(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Kategori
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}
		cek := DB.Create(&data).Error
		if cek != nil {
			http.Error(w, "Id barang tidak ada", 500)
		}

		w.Write([]byte("Sukses post data"))
		w.WriteHeader(http.StatusCreated)
		return

	}
	http.Error(w, "Error Not Found ", 404)
}

func Tampil(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		url := r.URL.String()
		var inp []string = strings.Split(url, "/")
		put, err := strconv.Atoi(inp[2])

		type All struct {
			Id_barang     int    `json:"id_barang"`
			Nama_barang   string `json:"nama_barang"`
			Harga_barang  int    `json:"harga_barang"`
			Id_kategori   int    `json:"id_kategori"`
			Nama_kategori string `json:"nama_kategori"`
			Id_Jual       int    `json:"id_jual"`
		}

		if err != nil {
			http.Error(w, "Error decode JSON", http.StatusInternalServerError)
		}

		var data []All
		DB.Model(&models.Kategori{}).Select("barangs.id_barang, barangs.nama_barang, barangs.harga_barang, kategoris.id_kategori, kategoris.nama_kategori, juals.id_jual").
			Joins("inner join barangs on kategoris.id_kategori=barangs.kategori_id inner join juals on juals.barang_id = barangs.id_barang").
			Where(&models.Kategori{Id_kategori: put}).Scan(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "error encode ke json ", 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "error not found ", 404)
}
