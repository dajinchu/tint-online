package handler

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	builder "github.com/dajinchu/tint/builder/yaml"
	"github.com/cjcodell1/tint/machine"
)

type Input struct {
	accept []string
	reject []string
	machine string
	machineType string
}

type Output struct {
	acceptResults []Result
	rejectResults []Result
}

func Handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
    panic(err)
	}

	var i Input
	err = json.Unmarshal(body, &i)
	if err != nil {
		// respond with 400
	}

	var m machine.Machine
	var machineType = "one-way-tm"
	m, err = builder.Build(i.machine, machineType)

	var out Output
	for _, s := range i.accept {
		out.acceptResults = append(out.acceptResults, test(m, s))
	}

	var wBody []byte
	wBody, err = json.Marshal(out)
	if err != nil {
		// respond with 500
	}
	fmt.Fprintf(w, string(wBody))
}

type Result struct {
	status ResultStatus
	verbose string
}
type ResultStatus int

const (
	Accept ResultStatus = 1
  Reject = 0
	Error = -1
)

func test(m machine.Machine, input string) Result {
	var err error
	var status ResultStatus
	var verbose strings.Builder
	conf := m.Start(input)
	for {
		// print verbosely
		verbose.WriteString(conf.Print())

		// check if accept or reject and break
		if m.IsAccept(conf) {
			status = Accept
			break
		} else if m.IsReject(conf) {
			status = Reject
		}

		// step
		conf, err = m.Step(conf)
		if err != nil {
			status = Error
		}
		verbose.WriteString("\n")
	}
	return Result {
		status: status,
		verbose: verbose.String(),
	}
}