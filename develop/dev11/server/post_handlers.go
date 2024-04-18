package server

import (
	"encoding/json"
	"fmt"
	"github.com/odysseymorphey/WB_L2/develop/dev11/event"
	"github.com/odysseymorphey/WB_L2/develop/dev11/tools"
	"io"
	"log"
	"net/http"
)

const (
	permissionError     int = 2
	valid               int = 1
	invalidData         int = 0
	internalServerError int = -1
)

func getDataFromRequest(r *http.Request) (event.Event, error) {
	evt := event.Event{}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return evt, err
	}
	err = json.Unmarshal(data, &evt)
	if err != nil {
		return evt, err
	}
	return evt, nil
}

func (s *Server) validatePost(w http.ResponseWriter, event *event.Event, actionType string) int {
	s.Mu.RLock()
	data, ok := s.Cache[event.EventName]
	s.Mu.RUnlock()
	result := invalidData

	date := tools.IsValidDate(event)
	switch actionType {
	case "create":
		if !ok && date && event.EventName != "" {
			result = valid
		}
	case "update":
		if ok && date {
			result = valid
		}
	case "delete":
		if ok {
			if data.UserId == event.UserId {
				result = valid
			} else {
				result = permissionError
			}
		}
	default:
		result = internalServerError
	}
	return result
}

func (s *Server) validateAndRespond(w http.ResponseWriter, code int) bool {
	if code == valid {
		return true
	}
	switch code {
	case internalServerError:
		tools.MakeJsonRespond(w, 503, tools.JsonError("internal server error"))
	case invalidData:
		tools.MakeJsonRespond(w, 400, tools.JsonError("invalid data"))
	case permissionError:
		tools.MakeJsonRespond(w, 500, tools.JsonError("permisson error"))
	}
	return false
}

func (s *Server) postRequestCheck(w http.ResponseWriter, r *http.Request, request string) (event.Event, error) {
	evt := event.Event{}
	if r.Method != http.MethodPost {
		errorString := "method not allowed"
		tools.MakeJsonRespond(w, 500, tools.JsonError(errorString))
		return evt, fmt.Errorf(errorString)
	}
	evt, err := getDataFromRequest(r)
	if err != nil {
		log.Println(err)
		tools.MakeJsonRespond(w, 503, tools.JsonError("internal server error"))
		return evt, err
	}
	validate := s.validatePost(w, &evt, request)
	if !s.validateAndRespond(w, validate) {
		return evt, fmt.Errorf("something being wrong")
	}
	return evt, nil
}

func (s *Server) createAndUpdate(w http.ResponseWriter, r *http.Request, request string) {
	evt, err := s.postRequestCheck(w, r, request)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.setEvent(evt)
	tools.MakeJsonRespond(w, 200, tools.JsonResult("ok"))
}

func (s *Server) CreateEvent(w http.ResponseWriter, r *http.Request) {
	s.createAndUpdate(w, r, "create")
}

func (s *Server) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	s.createAndUpdate(w, r, "update")
}

func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	evt, err := s.postRequestCheck(w, r, "delete")
	if err != nil {
		return
	}
	s.deleteEvent(evt.EventName)
	tools.MakeJsonRespond(w, 200, tools.JsonResult("ok"))
}
