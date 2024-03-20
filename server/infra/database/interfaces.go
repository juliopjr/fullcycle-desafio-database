package database

import (
	"github.com/juliopjr/fullcycle-desafio-database/server/entity"
)

type CotationInterface interface {
	Create(*entity.Cotation) error
}
