package main

import (
	"fmt"
	"net/http"
	"relation/controller"
)

func main() {
	http.HandleFunc("/get", controller.Get)
	http.HandleFunc("/post", controller.Post)
	http.HandleFunc("/getjual", controller.GetJual)
	http.HandleFunc("/postjual", controller.PostJual)
	http.HandleFunc("/getbarang/", controller.GetBarang)
	http.HandleFunc("/postkat", controller.PostKategori)

	fmt.Println("Start Runnning")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Println("Can't Start Server")
		return
	}

}
