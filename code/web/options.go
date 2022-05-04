package web

import "github.com/frankywahl/ddns53"

// Option is an argument that can be passwed to the constructor
type Option func(s *Server) error

// WithLogger defines a logger for the server
func WithLogger(l ddns53.Logger) func(s *Server) error {
	return func(s *Server) error {
		s.logger = l
		return nil
	}
}

// WithAuth defines the basic auth required
func WithAuth(username, password string) Option {
	return func(s *Server) error {
		s.username = username
		s.password = password
		return nil
	}
}

// WithZoneID defines the hosted zone that we are allowed to affect
func WithZoneID(id string) Option {
	return func(s *Server) error {
		s.zoneID = id
		return nil
	}
}

// WithFQDN defines the Fully Qualified Domain name that can be changed
func WithFQDN(fqdn string) Option {
	return func(s *Server) error {
		s.fqdn = fqdn
		return nil
	}
}
