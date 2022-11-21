package controller

import (
	"belajargorm/connection"
	"belajargorm/models"
	"encoding/json"
	"net/http"
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
		w.Write([]byte("Sukses Post Data"))
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

		err := DB.First(&models.Barang{}, "id=?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Delete(&models.Barang{}, DB.Where("id=?", id[2]))

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return

	}
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
		var datarequest models.Barang
		if err := decoder.Decode(&datarequest); err != nil {
			http.Error(w, "Error Decode JSON", 500)
			return
		}

		err := DB.First(&models.Barang{}, "id = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Model(&models.Barang{}).Where("id = ?", id[2]).Updates(&datarequest)

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return

	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

}

func Limit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Barang
		DB.Find(&data)

		u := r.URL.String()
		var id []string = strings.Split(u, "/")
		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		idd, error := strconv.Atoi(id[2]) //castinya
		if error != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		DB.Limit(idd).Find(&data)

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

func Detail(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		u := r.URL.String()
		var id []string = strings.Split(u, "/")
		if id[2] == "" {
			http.Error(w, "Id kosong", 500)
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
		w.Header().Set("Content-Type", "application/json")
		w.Write(datajson)
		w.WriteHeader(http.StatusOK)
		return
	}
}

func GetJual(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Jual
		//DB.Find(&data)
		DB.Model(&models.Jual{}).Preload("Barang").Find(&data)
		//DB.Model : wadah
		//preload : untuk memanggil struct barang yang ada di struct jual

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

func PostJual(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Jual
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON", 500)
			return
		}
		cek := DB.Create(&data).Error
		if cek != nil {
			http.Error(w, "Id barang tidak ada", 500)
		}

		w.Write([]byte("Sukses Post Data"))
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func GetKategori(w http.ResponseWriter, r *http.Request) {
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

func PostKategori(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Kategori
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON", 500)
			return
		}
		cek := DB.Create(&data).Error
		if cek != nil {
			http.Error(w, "Id barang tidak ada", 500)
		}

		w.Write([]byte("Sukses Post Data"))
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func GetJoin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		u := r.URL.String()
		var param []string = strings.Split(u, "/")
		if param[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		tamp, err := strconv.Atoi(param[2])
		if err != nil {
			http.Error(w, "Error convert", 500)
			return
		}

		type join struct {
			Id       int    `json:"id"`
			Nama_brg string `json:"nama_brg"`
			Harga    int    `json:"harga"`
			Id_ktg   int    `json:"id_ktg"`
			Nama_ktg string `json:"nama_ktg"`
			Id_jual  int    `json:"id_jual"`
		}

		var tampJoin []join
		DB.Model(&models.Kategori{}).Select("barangs.id, barangs.nama_brg, barangs.harga, kategoris.id_ktg, kategoris.nama_ktg, juals.id_jual").Joins("INNER JOIN barangs ON kategoris.id_ktg=barangs.kategori_id INNER JOIN juals ON juals.barang_id=barangs.id").Where(&models.Kategori{Id_ktg: tamp}).Scan(&tampJoin)
		datajson, err := json.Marshal(tampJoin)
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
