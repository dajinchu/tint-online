package handler

import (
  "fmt"
	"net/http"
	"io/ioutil"
	builder "github.com/dajinchu/tint/builder/yaml"
	"github.com/cjcodell1/tint/machine"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
    panic(err)
	}
	fmt.Printf("%s", body)
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
	var m machine.Machine
	var machineType = "one-way-tm"
	m, err = builder.Build(string(body), machineType)
	fmt.Fprintf(w, "%d", test(m, "0 1 0"))
}

type Result int

const (
	Accept Result = 1
  Reject = 0
	Error = -1
)

func test(m machine.Machine, input string) Result {
	var err error
	conf := m.Start(input)
	for {
		// print verbosely
		// if verboseFlag {
		// 	fmt.Println(conf.Print())
		// }

		// check if accept or reject and break
		if m.IsAccept(conf) {
			return 1
		} else if m.IsReject(conf) {
			return 0
		}

		// step
		conf, err = m.Step(conf)
		if err != nil {
			return -1
		}
	}
}