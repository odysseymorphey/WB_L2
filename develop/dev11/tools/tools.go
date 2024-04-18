package tools

import (
	"fmt"
	"github.com/odysseymorphey/WB_L2/develop/dev11/event"
	"log"
	"net/http"
	"time"
)

func IsValidDate(event *event.Event) bool {
	value, err := time.Parse("2006-01-02", string(event.Date))
	if err != nil {
		log.Println(err)
		return false
	}
	event.Time = value
	return true
}

func JsonResult(respondText string) []byte {
	return []byte(fmt.Sprintf(`{"result": %s}`, respondText))
}

func JsonError(respondText string) []byte {
	return []byte(fmt.Sprintf(`{"error": "%s"}`, respondText))
}

func MakeJsonRespond(w http.ResponseWriter, code int, data []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(data)
	if err != nil {
		log.Println(err)
	}
}
