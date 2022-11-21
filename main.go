package main

import (
	"fmt"
	"net/http"
	"pt9/controller"
)

func main() {
	http.HandleFunc("/get", controller.Get)
	http.HandleFunc("/post", controller.Post)
	http.HandleFunc("/update/", controller.Update)
	http.HandleFunc("/delete/", controller.Delete)
	http.HandleFunc("/getlimit", controller.GetLimit)
	http.HandleFunc("/detail/", controller.Detail)

	http.HandleFunc("/getjual", controller.GetJual)
	http.HandleFunc("/postjual", controller.PostJual)
	http.HandleFunc("/getkategori", controller.GetKategori)
	http.HandleFunc("/postkategori", controller.PostKategori)
	http.HandleFunc("/tampil/", controller.Tampil)

	fmt.Println("Running Service")

	if err := http.ListenAndServe(":5001", nil); err != nil {
		fmt.Println("Error Starting Service")
	}
}
