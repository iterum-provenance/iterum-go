package manager

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/prometheus/common/log"

	"github.com/iterum-provenance/iterum-go/util"
)

type upstreamState struct {
	Finished bool `json:"finished"`
}

// UpstreamChecker is the structure responsible for checking the
// state of upstream pipeline components by asking the manager
// It notifies all its consumber which can be registered
type UpstreamChecker struct {
	// Register is a channel via which you can register additional consumers
	Register  chan chan bool
	state     chan upstreamState
	consumers []chan bool
	retries   int

	ManagerURL   string
	PipelineHash string
	ProcessName  string
}

// NewUpstreamChecker creates a new instance of UpstreamChecker and intializes its necessary values
func NewUpstreamChecker(managerURL, pipelineHash, processName string, retries int) UpstreamChecker {
	return UpstreamChecker{
		Register:     make(chan chan bool, 10),
		state:        make(chan upstreamState, 10),
		retries:      retries,
		ManagerURL:   managerURL, // see env.go
		PipelineHash: pipelineHash,
		ProcessName:  processName,
	}
}

// register adds a channel to be notified when messages arrive
func (checker *UpstreamChecker) register(target chan bool) {
	checker.consumers = append(checker.consumers, target)
}

// notify sends a message to all of checker's consumers
func (checker UpstreamChecker) notify() {
	for _, consumer := range checker.consumers {
		consumer <- true
	}
}

// _getData takes a url and a pointer to an interface. It fires a GET request
// and tries to json.Unmarshal the response into the target. If and error occurs
// it is returned
func _getData(url string, target interface{}) (err error) {
	defer util.ReturnErrOnPanic(&err)
	var client http.Client

	resp, err := client.Get(url)
	util.PanicIfErr(err, fmt.Sprintf("Get request to manager failed due to '%v'", err))
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Getting upstream state from manager did not return 'StatusOK'")
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	util.PanicIfErr(err, fmt.Sprintf("Reading response body failed due to: '%v'", err))

	err = json.Unmarshal(bodyBytes, target)
	util.PanicIfErr(err, fmt.Sprintf("Parsing response body failed due to: '%v'", err))

	return
}

// Code to have a mockup check
var dummyCounter = 5

func dummyCheck(url string, state *upstreamState) error {
	log.Warnf("Running dummy upstream state checker")
	if dummyCounter > 0 {
		dummyCounter--
		state.Finished = false
		return nil
	}
	state.Finished = true
	return nil
}

// periodicCheck polls the manager every 'interval' amount of seconds
func (checker UpstreamChecker) periodicCheck(wg *sync.WaitGroup, interval int) {
	defer wg.Done()
	retries := checker.retries
	handleErr := func(err error) {
		log.Warnf("Getting upstream state from manager failed due to '%v'\nretrying...\n", err)
		retries--
		if retries < 0 {
			log.Fatalf("Cannot get state from manager due to '%v'\n", err)
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}

	endpoint := checker.ManagerURL + "/pipeline/" + checker.PipelineHash + "/" + checker.ProcessName + "/upstream_finished"
	for {
		newState := upstreamState{}
		var err error
		if strings.Contains(checker.ManagerURL, "dummy") {
			err = dummyCheck(endpoint, &newState)
		} else {
			err = _getData(endpoint, &newState)
		}
		if err != nil {
			handleErr(err)
			continue
		}
		checker.state <- newState
		if newState.Finished {
			break
		}
		retries = checker.retries
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

// StartBlocking starts the main loop of the UpstreamChecker
func (checker *UpstreamChecker) StartBlocking() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go checker.periodicCheck(&wg, 5)

outer:
	for {
		select {
		case consumer := <-checker.Register:
			log.Infof("Registered new consumer\n")
			checker.register(consumer)
		case state := <-checker.state:
			log.Debugf("Newly arrived state: %v\n", state)
			if state.Finished {
				checker.notify()
				break outer
			}
		}
	}
	wg.Wait()
}

// Start asychronously calls StartBlocking via Gorouting
func (checker *UpstreamChecker) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		checker.StartBlocking()
	}()
}
