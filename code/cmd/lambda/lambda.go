package main

import (
	"flag"
	"log"
	"os"

	"github.com/akrylysov/algnhsa"
	"github.com/frankywahl/ddns53/web"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

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
		log.Fatal(err)
	}

	algnhsa.ListenAndServe(srv.Handler, nil)
}
