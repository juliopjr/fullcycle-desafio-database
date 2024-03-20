package database

import (
	"github.com/juliopjr/fullcycle-desafio-database/server/entity"
)

type QuotationInterface interface {
	Create(*entity.Quotation) error
}
