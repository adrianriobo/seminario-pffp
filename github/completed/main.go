package main

import (
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
	defer logger.Sync()
	sugar.Debug("PFFP!")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("PFFP!"))
}

func main() {
	http.HandleFunc("/seminario", helloHandler)

	port := ":8080"
	log.Printf("Servidor escuchando en http://localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error iniciando el servidor: %v", err)
	}
}
