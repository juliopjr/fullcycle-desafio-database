package webserver

import (
	"encoding/json"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/juliopjr/fullcycle-desafio-database/common/dto"
	"github.com/juliopjr/fullcycle-desafio-database/common/service"
	"github.com/juliopjr/fullcycle-desafio-database/server/entity"
	"github.com/juliopjr/fullcycle-desafio-database/server/infra/database"
)

var requestCounter int64

func ReportRequests() {
	for {
		time.Sleep(10 * time.Second)
		log.Println("? -> TOTAL REQUESTS:", requestCounter)
	}
}

func IncrementCounter() {
	atomic.AddInt64(&requestCounter, 1)
}

func NewQuotationHandler(db database.QuotationInterface) *quotatiionHandler {
	buffer := 25
	chProcessSucess := make(chan dto.Quotation, buffer)
	chProcessFails := make(chan error, buffer)
	chToDB := make(chan dto.Quotation, buffer)
	chSignalToExecute := make(chan bool, buffer)

	e := &quotatiionHandler{db, chSignalToExecute, chToDB, chProcessSucess, chProcessFails}

	go e.getData()
	go e.dbListener()

	go ReportRequests()

	return e
}

type quotatiionHandler struct {
	db database.QuotationInterface

	chSignalToExecute chan bool
	chToDB            chan dto.Quotation
	chProcessSucess   chan dto.Quotation
	chProcessFails    chan error
}

func (e *quotatiionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	go IncrementCounter()
	// go e.getData()
	// go e.dbListener()
	e.chSignalToExecute <- true

	ctxReq := r.Context()

	w.Header().Set("Content-Type", "application/json")

	for {
		select {
		case <-ctxReq.Done():
			log.Println("!-> Client cancelou")
			return

		case err := <-e.chProcessFails:
			log.Println("!-> Falha no processo:", err)
			// continuar loop até timeout

		case quotation := <-e.chProcessSucess:
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(quotation)
			return

		case <-time.After(1 * time.Second):
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(struct{ ErrorMessage string }{"timeout"})
			return
		}
	}
}

func (e *quotatiionHandler) getData() {
	for {
		<-e.chSignalToExecute
		apiData := entity.NewApiDolarQuotation()
		// 200ms didnt work, 300ms sometimes works
		requester := service.NewRequester(apiData.GetUrl(), 200*time.Millisecond)
		data, err := requester.Execute()
		if err != nil {
			e.chProcessFails <- err
			continue
		}
		if err := json.Unmarshal(*data, &apiData); err != nil {
			e.chProcessFails <- err
			continue
		}

		quotation := dto.Quotation{Bid: apiData.GetBid()}
		e.chProcessSucess <- quotation
		e.chToDB <- quotation
	}
}

func (e *quotatiionHandler) dbListener() {
	quotation := <-e.chToDB
	err := e.db.Create(&entity.Quotation{Bid: quotation.Bid})
	if err != nil {
		e.chProcessFails <- err
	}
}
