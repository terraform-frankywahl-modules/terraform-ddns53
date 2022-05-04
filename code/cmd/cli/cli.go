package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/frankywahl/ddns53/web"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	port := 8080

	logger := logrus.New()
	// logger.SetFormatter(&logrus.JSONFormatter{})
	username := flag.String("username", os.Getenv("USERNAME"), "Username for the DynDNS service")
	password := flag.String("password", os.Getenv("PASSWORD"), "Password for the DynDNS service")
	zoneId := flag.String("zone-id", os.Getenv("ZONE_ID"), "Hosted Zone ID of the record to update")
	fqdn := flag.String("fqdn", os.Getenv("FQDN"), "Fully Qualified domain name to update")
	flag.Parse()

	srv, err := web.New(
		web.WithLogger(logger),
		web.WithAuth(*username, *password),
		web.WithFQDN(*fqdn),
		web.WithZoneID(*zoneId),
	)
	if err != nil {
		return err
	}

	s := &http.Server{
		Handler: srv.Handler,
		Addr:    fmt.Sprintf(":%d", port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	logger.Infof("Starting server on port %s", s.Addr)
	return s.ListenAndServe()
}
