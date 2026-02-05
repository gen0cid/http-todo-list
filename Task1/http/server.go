package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	HttpHandlers *HTTPHandlers
}

func NewHTTPServer(httpHandlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		HttpHandlers: httpHandlers,
	}
}
func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/tasks").Methods(http.MethodPost).HandlerFunc(s.HttpHandlers.HandleCreateTask)
	router.Path("/tasks/{title}").Methods(http.MethodGet).HandlerFunc(s.HttpHandlers.HandleGetTask)
	router.Path("/tasks").Methods(http.MethodGet).Queries("completed", "false").HandlerFunc(s.HttpHandlers.HandleGetAllUncompletedTasks)

	router.Path("/tasks").Methods(http.MethodGet).HandlerFunc(s.HttpHandlers.HandleGetAllTasks)

	router.Path("/tasks/{title}").Methods(http.MethodPatch).HandlerFunc(s.HttpHandlers.HandleCompleteTask)
	router.Path("/tasks/{title}").Methods(http.MethodDelete).HandlerFunc(s.HttpHandlers.HandleDeleteTask)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
