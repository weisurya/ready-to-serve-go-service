package config

import (
	"crypto/tls"
	"net/http"
	"time"
)

type ServerSetting struct {
	Port         string
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

func (serverSetting ServerSetting) InitiateHTTPSServer(mux *http.ServeMux) *http.Server {

	// https://blog.cloudflare.com/exposing-go-on-the-internet/
	tlsConfig := &tls.Config{

		// Causes servers to use Go's default ciphersuite preferences,
		// which are tuned to avoid attacks. Does nothing on clients.
		PreferServerCipherSuites: true,

		// Only use curves which have assembly implementations
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519, // Go 1.8 only
		},

		MinVersion: tls.VersionTLS12,

		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,

			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,

			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only

			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305, // Go 1.8 only

			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,

			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,

			// Best disabled, as they don't provide Forward Secrecy,
			// but might be necessary for some clients
			// tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			// tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	return &http.Server{
		Addr: ":" + serverSetting.Port,

		ReadTimeout: time.Duration(int64(time.Second) * int64(serverSetting.ReadTimeout)),

		WriteTimeout: time.Duration(int64(time.Second) * int64(serverSetting.WriteTimeout)),

		IdleTimeout: time.Duration(int64(time.Second) * int64(serverSetting.IdleTimeout)),

		TLSConfig: tlsConfig,

		Handler: mux,
	}
}
