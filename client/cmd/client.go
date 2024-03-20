package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/juliopjr/fullcycle-desafio-database/client/service"
	"github.com/juliopjr/fullcycle-desafio-database/common/dto"
	commonService "github.com/juliopjr/fullcycle-desafio-database/common/service"
)

func main() {
	const url, timeout = "http://localhost:8080/cotacao", 300 * time.Millisecond
	requester := commonService.NewRequester(url, timeout)
	data, err := requester.Execute()

	if err != nil {
		log.Println(err.Error())
		return
	}

	var quotation dto.Quotation
	err = json.Unmarshal(*data, &quotation)
	if err != nil {
		panic(err)
	}

	text := fmt.Sprintf("Dólar: " + quotation.Bid + "\n")

	const fileName = "cotacao.txt"
	register := service.NewRegister(fileName)
	register.OnFile(&text)
}
