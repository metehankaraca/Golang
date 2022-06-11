package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := "8000"
	klasor := ""

	fmt.Printf("Olusturulacak port numarasini giriniz (Default Port No:8000): ")
	fmt.Scanf("%s", &port)

	if port == "" {
		port = "8000"
	} else if len(port) < 4 {
		port = "8000"
		fmt.Println("Port degerini 4 karakterden az girdiginiz icin default atama yapilmistir.")
	} else if len(port) > 4 {
		port = "8000"
		fmt.Println("Port degerini 4 karakterden cok girdiginiz icin default atama yapilmistir.")
	}

	fmt.Printf("Klasor ismi giriniz: ")
	fmt.Scanf("%s", klasor)
	http.Handle("/", http.FileServer(http.Dir(klasor)))
	fmt.Printf("http:localhost:%s olusturuldu.\n", port)
	http.ListenAndServe(":"+port, nil)
}
