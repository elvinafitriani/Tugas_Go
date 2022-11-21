package main

import (
	"belajargorm/controller"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/get", controller.Get)
	http.HandleFunc("/post", controller.Post)
	http.HandleFunc("/delete/", controller.Delete)
	http.HandleFunc("/update/", controller.Update)
	http.HandleFunc("/getLimit/", controller.Limit)
	http.HandleFunc("/getDetail/", controller.Detail)

	http.HandleFunc("/postJual", controller.PostJual)
	http.HandleFunc("/getJual", controller.GetJual)
	http.HandleFunc("/postKategori", controller.PostKategori)
	http.HandleFunc("/getKategori", controller.GetKategori)

	http.HandleFunc("/getJoin/",controller.GetJoin)

	fmt.Println("Succes Starting Sevice")

	if err := http.ListenAndServe(":5002", nil); err != nil {
		fmt.Println("Error starting service")
	}
}
