package database

import (
	"context"
	"errors"
	"time"

	"github.com/juliopjr/fullcycle-desafio-database/server/entity"
	"gorm.io/gorm"
)

func NewQuotationDB(db *gorm.DB) *quotation {
	return &quotation{db}
}

type quotation struct {
	db *gorm.DB
}

func (e *quotation) Create(cotatiton *entity.Quotation) error {
	ctxDB, cancelCtxDB := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelCtxDB()
	chDB := make(chan error)

	go func() {
		chDB <- e.db.WithContext(ctxDB).Create(cotatiton).Error
	}()
	select {
	case <-ctxDB.Done():
		return errors.New("DB timeout")
	case err := <-chDB:
		if err != nil {
			return errors.New("DB internal error")
		}
	}
	return nil
}
