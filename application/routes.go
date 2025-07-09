package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/themanojk/reflekt/handler"
)

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(JSONMiddleware)

	router.Route("/devices", a.loadOrderRoutes)
	a.router = router
}

func (a *App) loadOrderRoutes(router chi.Router) {
	deviceHandler := handler.NewDevice(a.store)

	router.Post("/", deviceHandler.Create)
	router.Get("/{macAddress}", deviceHandler.GetByMacAddress)
}
