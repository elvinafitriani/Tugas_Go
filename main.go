package main

import (
	"fmt"
	"net/http"
	"relasi/controller"
)

func main() {
	http.HandleFunc("/get", controller.Get)
	http.HandleFunc("/post", controller.Post)
	http.HandleFunc("/delete/", controller.Delete)
	http.HandleFunc("/detail/", controller.Detail)

	http.HandleFunc("/getjual", controller.GetJual)
	http.HandleFunc("/postjual", controller.PostJual)

	http.HandleFunc("/postkategori", controller.PostKtg)
	http.HandleFunc("/getkategori", controller.GetKtg)

	http.HandleFunc("/getjoin/", controller.Join)

	fmt.Println("Running Service")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Printf("Error Starting Service")
	}
}
