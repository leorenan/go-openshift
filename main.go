package main

import (
	"fmt"
	"net/http"
)

func olaHandler(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")
	if nome == "" {
		nome = "Red Hat"
	}

	fmt.Fprintln(w, "<h1>Olá a PoC Red Hat OpenShift! </h1>")
	fmt.Fprintln(w, "<h1>Bem vindo", nome, "!</h1>")
	fmt.Fprintln(w, "<h1>Obrigado! </h1>")
}

func listenServe(port string) {
	fmt.Printf("Listen na porta %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe ERRO: " + err.Error())
	}
}

func main() {
	http.HandleFunc("/", olaHandler)
	go listenServe("8080")
	select {}
}
