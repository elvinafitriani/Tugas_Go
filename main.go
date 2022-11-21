package main

import (
	"crud/controller"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/get", controller.Get)
	http.HandleFunc("/post", controller.Post)
	http.HandleFunc("/getlimit", controller.Getlimits)
	http.HandleFunc("/put/", controller.Update)

	http.HandleFunc("/postJual", controller.PostJual)
	http.HandleFunc("/getJual", controller.GetJual)
	http.HandleFunc("/view/", controller.View)

	fmt.Println("Running Service")

	if err := http.ListenAndServe(":5000", nil); err != nil { //return listenandserve bertipe error
		fmt.Println("Error Starting Service")
	}

}
