package main

import (
	"fmt"
	"testing"

	"github.com/bsromr/modern-webapp-w-devops/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := Routes(&app)
	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing; test passed
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, type is %T", v))
	}
}
