package main

import (
	"log"

	"github.com/giantswarm/helm-values-gen/cmd"
)

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	cmd.Execute()
}
