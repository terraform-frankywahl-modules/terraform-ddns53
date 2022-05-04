package web

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/frankywahl/ddns53"
)

// Server defines a web server
type Server struct {
	Handler  http.Handler
	logger   ddns53.Logger
	username string
	password string
	zoneID   string
	fqdn     string
}

// New creates and initializes a server
func New(options ...Option) (*Server, error) {
	srv := &Server{
		logger: ddns53.NoopLogger,
	}

	for _, opt := range options {
		if err := opt(srv); err != nil {
			return &Server{}, err
		}
	}
	srv.routes()

	return srv, nil
}

func (s *Server) routes() {
	router := http.NewServeMux()
	router.HandleFunc("/", s.rootHandler())
	router.HandleFunc("/nic/update", s.Auth(s.ddns53))
	s.Handler = router
}

func (s *Server) ddns53(w http.ResponseWriter, r *http.Request) {
	ip, ok := r.URL.Query()["myip"]
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if err := ddns53.UpdateRecord(r.Context(), s.zoneID, s.fqdn, ip[0]); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return

	}

	fmt.Fprintf(w, `{"status":"done"}`)

}

func parseBasicHeaderValue(s string) (username, password string, err error) {
	h := strings.SplitN(s, " ", 2)
	if len(h) != 2 {
		return "", "", errors.New("request does not have a valid authorization header")
	}
	if strings.ToLower(h[0]) != "basic" {
		return "", "", errors.New("request does not have a valid basic auth header")
	}
	b, err := base64.StdEncoding.DecodeString(h[1])
	if err != nil {
		return "", "", errors.New("could not base64 decode auth header")
	}
	v := string(b)
	vc := strings.SplitN(v, ":", 2)
	return vc[0], vc[1], nil
}

func (s *Server) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		u, p, err := parseBasicHeaderValue(authorization)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, err.Error())
			return
		}
		if s.username != u || s.password != p {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Authentication failed")
			return
		}

		next(w, r)
	}
}

func (s *Server) rootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		s.health()(w, r)
	}

}

func healthEndpoint(s *Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.health()
	}
}

// health returns the status of the server
func (s *Server) health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"status":"ok"}`)
	}
}
