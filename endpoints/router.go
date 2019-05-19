package endpoints

import (
	"net/http"

	root "github.com/weisurya/ready-to-serve-go-service/endpoints/root"
	m "github.com/weisurya/ready-to-serve-go-service/middlewares"
)

func InitiateRouter() *http.ServeMux {
	mux := http.NewServeMux()

	handler := m.CreateHandler()

	mux.HandleFunc("/", handler.Log(root.Get))

	return mux
}
