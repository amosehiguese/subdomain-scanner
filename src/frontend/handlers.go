package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var (
	templates = template.Must(template.New("").
		ParseGlob("templates/*.html"))
)

func (fe *frontendServer) homeHandler(w http.ResponseWriter, r *http.Request) {
	zapLog := r.Context().Value(ctxKeyLog{}).(*zap.Logger)
	zapLog.Sugar().Info("home")
	if err := templates.ExecuteTemplate(w, "home", nil); err != nil {
		zapLog.With(
			zap.Error(err),
		).Error("internal server error")
		return
	}
}

func (fe *frontendServer) scanHandler(w http.ResponseWriter, r *http.Request) {
	zapLog := r.Context().Value(ctxKeyLog{}).(*zap.Logger)
	domain := mux.Vars(r)["domain"]

	zapLog.With(
		zap.String("domain", domain),
	).Debug("processing request")

	result, err := fe.scan(r, domain)
	if err != nil {
		zapLog.With(
			zap.Error(err),
		).Error("internal server error")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"status": "failed",
			"msg":    "internal server error",
			"code":   http.StatusInternalServerError,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":     "success",
		"subdomains": result,
	})
}
