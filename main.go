package main

import (
	"fmt"
	"net/http"
	"pelatihan/controller"
)

func main() {

	http.HandleFunc("/get", controller.Get)
	http.HandleFunc("/post", controller.Post)
	http.HandleFunc("/delete/", controller.Delete)
	http.HandleFunc("/update/", controller.Update)
	http.HandleFunc("/detail/", controller.Detail)
	http.HandleFunc("/getlimit/", controller.GetLimit)

	http.HandleFunc("/postJual", controller.PostJual)
	http.HandleFunc("/getJual", controller.GetJual)

	http.HandleFunc("/postKategori", controller.PostKategori)
	http.HandleFunc("/getKategori", controller.GetKategori)

	http.HandleFunc("/getJoin/", controller.GetJoin)

	fmt.Println("Starting Sevice")

	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Println("Error starting service")
	}

}
