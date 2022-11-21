package controller

import (
	"encoding/json"
	"net/http"
	"pelatihan/connection"
	"pelatihan/models"
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
		var data models.Barang
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

		//delete
		DB.Model(&models.Barang{}).Where("id = ?", id[2]).Delete(&data)

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
		w.Write(datajson)
		w.WriteHeader(http.StatusOK)
		return
	}
}

func GetLimit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var l []models.Barang
		DB.Limit(5).Find(&l)

		datajson, err := json.Marshal(l)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(datajson)
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

		cek := DB.Create(&data)
		if cek.Error != nil {
			http.Error(w, "Id Barang Tidak Ada", 500)
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
		//DB.Find(&data)
		DB.Model(&models.Jual{}).Preload("Barang").Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(datajson)
		w.WriteHeader(http.StatusAccepted)
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

		DB.Create(&data)
		// if cek.Error != nil {
		// 	http.Error(w, "Id Barang Tidak Ada", 500)
		// 	return
		// }
		w.Write([]byte("Sukses Post Data"))
		w.WriteHeader(http.StatusCreated)
		return

	}
	http.Error(w, "Error Not Found", 404)
}

func GetKategori(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Kategori
		//DB.Find(&data)
		DB.Model(&models.Kategori{}).Preload("Barang").Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(datajson)
		w.WriteHeader(http.StatusAccepted)
		return
	}

	http.Error(w, "Error Not Found", 404)
}

func GetJoin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		u := r.URL.String()
		var id []string = strings.Split(u, "/")
		if id[2] == "" {
			http.Error(w, "Id kosong", 500)
			return
		}
		var data []models.Join
		// DB.Raw("SELECT barangs.*,kategoris.nama_ktg,juals.id_jual FROM kategoris RIGHT JOIN barangs ON kategoris.id_ktg=barangs.kategori_id LEFT JOIN juals ON barangs.id_brg=juals.barang_id;").Scan(&data)
		DB.Table("Kategoris").Select("barangs.id_brg,barangs.nama_brg,barangs.harga,kategoris.id_ktg,kategoris.nama_ktg,juals.id_jual").Joins("RIGHT JOIN barangs ON barangs.kategori_id = kategoris.id_ktg").Joins("LEFT JOIN juals ON barangs.id_brg = juals.barang_id").Where("barangs.id_brg = ?", id[2]).Scan(&data)
		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(datajson)
		w.WriteHeader(http.StatusAccepted)
		return
	}

	http.Error(w, "Error Not Found", 404)
}
