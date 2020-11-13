package handler

import (
  "fmt"
	"net/http"
	"io/ioutil"
	// "gopkg.in/yaml.v2"
	//  "github.com/dajinchu/tint/builder/yaml"
	//  "github.com/dajinchu/tint/machine"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
    panic(err)
	}

	fmt.Printf("%s", body)
  fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
	// m, err = yaml.Build(config, machineFlag)
}