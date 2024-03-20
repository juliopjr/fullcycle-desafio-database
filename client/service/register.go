package service

import (
	"log"
	"os"
)

func NewRegister(fileName string) *register {
	return &register{fileName}
}

type register struct {
	fileName string
}

func (e *register) OnFile(text *string) {
	file, err := os.OpenFile(e.fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(*text)
	if err != nil {
		log.Println("Não foi possível realizar o registro:", err.Error())
	}
}
