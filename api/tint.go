package main

import (
  "fmt"
	"net/http"
	"gopkg.in/yaml.v2"
	 "github.com/dajinchu/tint/builder/yaml"
	 "github.com/dajinchu/tint/machine"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
    panic(err)
	}

	fmt.Printf("%s", b)
  fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
	// m, err = yaml.Build(config, machineFlag)
}