package panicHandler

import (
	"fmt"
	"log"
)

func PanicIfErr(funcName string, err error) {
	if err != nil {
		PanicNicely(funcName, err)
	}
}

func ErrorNicely(funcName string, err error) string {
	return fmt.Sprintf("> %s threw '%s'", funcName, err.Error())
}

func PanicNicely(funcName string, err error) {
	log.Panic(ErrorNicely(funcName, err))
}
