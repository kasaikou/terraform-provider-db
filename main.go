package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/kasaikou/terraform-provider-db/provider"
)

var version = ""

func main() {
	var debugFlag bool
	flag.BoolVar(&debugFlag, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	err := providerserver.Serve(context.Background(), provider.New(version), providerserver.ServeOpts{
		Address: "hashicorp.com/kasaikou/db",
		Debug:   debugFlag,
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
