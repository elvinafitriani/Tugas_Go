package main

import (
	"db-warung/controller"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/get", controller.Get)
	http.HandleFunc("/post", controller.Post)
	http.HandleFunc("/delete/", controller.Delete)
	http.HandleFunc("/update/", controller.Update)
	http.HandleFunc("/getlimit", controller.GetLimit)
	http.HandleFunc("/getdetail/", controller.Detail)

	http.HandleFunc("/postJual", controller.PostJual)
	http.HandleFunc("/getJual", controller.GetJual)
	http.HandleFunc("/postKategori", controller.PostKategori)
	http.HandleFunc("/getKategori", controller.GetKategori)
	http.HandleFunc("/getview/", controller.GetRelasi)

	fmt.Println("running service")

	if err := http.ListenAndServe(":5003", nil); err != nil {
		fmt.Println("Error Starting Service")
	}

}

// crud
//relasi
//database gorm
//api->fe
