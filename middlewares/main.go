package middlewares

import (
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/weisurya/ready-to-serve-go-service/setting"
)

type Handler struct {
	Logger  *log.Logger
	DB      *sqlx.DB
	Context Context
}

type Context struct {
	Req *http.Request
	Res http.ResponseWriter
}

func CreateHandler() *Handler {
	return &Handler{
		Logger: setting.GetLog(),
		DB:     setting.GetDB(),
	}
}

func (h *Handler) Log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.Logger.Printf("Resp. Time: %s\n", time.Now().Sub(startTime))

		next(w, r)
	}
}
