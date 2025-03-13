package main

import (
	"fmt"
	"log"
	"net/http"

	"go.uber.org/zap"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	sugar := logger.Sugar()
	defer func() {
		if err := logger.Sugar(); err != nil {
			fmt.Println("Error when closing:", err)
		}
	}()

	sugar.Debug("PFFP!")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("PFFP!")); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/seminario", helloHandler)

	port := ":8080"
	log.Printf("Servidor escuchando en http://localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error iniciando el servidor: %v", err)
	}
}
