package server

import (
	"fmt"
	"github.com/odysseymorphey/WB_L2/develop/dev11/event"
	mw "github.com/odysseymorphey/WB_L2/develop/dev11/middleware"
	"log"
	"net/http"
	"sync"
)

type Server struct {
	Mu    sync.RWMutex
	Cache map[string]event.Event
}

func NewServer() *Server {
	srv := &Server{
		Cache: make(map[string]event.Event),
	}

	srv.SetupHandlers()

	return srv
}

func (s *Server) setEvent(event event.Event) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Cache[event.EventName] = event
}

func (s *Server) deleteEvent(eventName string) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	delete(s.Cache, eventName)
}
func (s *Server) SetupHandlers() {
	s.setupGetHandlers()
	s.setupPostHandlers()
}

func (s *Server) setupGetHandlers() {
	http.HandleFunc("/event_by_name", mw.Logger(s.EventByName))
	http.HandleFunc("/events_for_day", mw.Logger(s.EventsForDay))
	http.HandleFunc("/events_for_week", mw.Logger(s.EventsForWeek))
	http.HandleFunc("/events_for_month", mw.Logger(s.EventsForMonth))
}

func (s *Server) setupPostHandlers() {
	http.HandleFunc("/create_event", mw.Logger(s.CreateEvent))
	http.HandleFunc("/update_event", mw.Logger(s.UpdateEvent))
	http.HandleFunc("/delete_event", mw.Logger(s.DeleteEvent))
}
func (s *Server) Run() {
	fmt.Println("Server listen on: 0.0.0.0:8080")
	log.Println(http.ListenAndServe("0.0.0.0:8080", nil))
}
