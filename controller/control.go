package controller

import (
	"encoding/json"
	"net/http"
	"relasi/connection"
	"relasi/models"
	"strings"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = connection.ConnectToDB()
}

func Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Barang

		DB.Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)

		var data []models.Barang

		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON", 500)
			return
		}
		DB.Create(&data)
		w.Write([]byte("Suscces Post Data"))
		w.WriteHeader(200)
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

		err := DB.First(&models.Barang{}, "id = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		var results models.Barang

		DB.Where("id = ?", id[2]).Find(&models.Barang{}).Scan(&results)
		datajson, err := json.Marshal(results)

		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		w.Write(datajson)
		w.WriteHeader(http.StatusOK)
		return
	}
}

func GetJual(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Jual

		DB.Model(&models.Jual{}).Preload("Barang").Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func GetKtg(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Kategori

		DB.Model(&models.Kategori{}).Preload("Barang").Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func Join(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		u := r.URL.String()
		var id []string = strings.Split(u, "/")

		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		var result []models.Join

		//DB.Raw("SELECT barangs.id_barang,barangs.nama_barang,barangs.harga,kategoris.id_ktg,kategoris.nama_ktg, juals.id_jual FROM kategoris RIGHT JOIN barangs ON barangs.kategori_id = kategoris.id_ktg LEFT JOIN juals ON barangs.id_barang = juals.barang_id").Where("barangs.id_barang = ?", id[2]).Scan(&result)
		DB.Table("Kategoris").Select("barangs.id_barang,barangs.nama_barang,barangs.harga,kategoris.id_ktg,kategoris.nama_ktg, juals.id_jual ").Joins("RIGHT JOIN barangs ON barangs.kategori_id = kategoris.id_ktg").Joins("LEFT JOIN juals ON barangs.id_barang = juals.barang_id").Where("barangs.id_barang = ?", id[2]).Scan(&result)

		datajson, err := json.Marshal(result)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
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

		DB.Delete(&models.Barang{}, "id = ?", id[2])

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return

	}
}

func PostJual(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)

		var data []models.Jual

		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON", 500)
			return
		}

		DB.Create(&data)
		w.Write([]byte("Suscces Post Data"))
		w.WriteHeader(200)
		return
	}
}

func PostKtg(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)

		var data []models.Kategori

		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON", 500)
			return
		}
		DB.Create(&data)
		w.Write([]byte("Suscces Post Data"))
		w.WriteHeader(200)
		return
	}
}
