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
	Accept []string      `json:"accept"`
	Reject []string      `json:"reject"`
	Machine string       `json:"machine"`
	MachineType string   `json:"machineType"`
}

type Output struct {
	AcceptResults []Result `json:"acceptResults"`
	RejectResults []Result `json:"rejectResults"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
    panic(err)
	}

	var i Input
	err = json.Unmarshal(body, &i)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
	}

	var m machine.Machine
	var machineType = "one-way-tm"
	m, err = builder.Build(i.Machine, machineType)

	// Run tests
	var out Output
	for _, s := range i.Accept {
		out.AcceptResults = append(out.AcceptResults, test(m, s))
	}

	var wBody []byte
	wBody, err = json.Marshal(out)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Could not marshal output", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(wBody))
}

type Result struct {
	Status ResultStatus `json:"status"`
	Verbose string      `json:"verbose"`
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
			break
		}

		// step
		conf, err = m.Step(conf)
		if err != nil {
			status = Error
			break
		}
		verbose.WriteString("\n")
	}
	return Result {
		Status: status,
		Verbose: verbose.String(),
	}
}