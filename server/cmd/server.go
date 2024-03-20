package main

import (
	"log"
	"net/http"

	"github.com/juliopjr/fullcycle-desafio-database/server/infra/database"
	"github.com/juliopjr/fullcycle-desafio-database/server/infra/webserver"
)

func main() {
	log.Println("-> Programa iniciado.")
	db := database.Initialize()
	quotationDB := database.NewQuotationDB(db)
	cotacaoHandler := webserver.NewQuotationHandler(quotationDB)

	http.Handle("/cotacao", cotacaoHandler)
	log.Println("-> Rotas preparadas.")

	log.Println("-> Servidor operante, aguardando requisições.")
	http.ListenAndServe(":8080", nil)
}
