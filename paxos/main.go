package main

import (
	"fmt"
	"sync"
)

type Proposal struct {
	ID    int
	Value string
}

type Acceptor struct {
	mu            sync.Mutex
	promiseID     int
	promiseValue  string
	acceptedID    int
	acceptedValue string
}

func (a *Acceptor) Prepare(n int) (int, string) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if n > a.promiseID {
		a.promiseID = n
		return a.acceptedID, a.acceptedValue
	}

	return -1, ""
}

func (a *Acceptor) Accept(n int, v string) bool {
	a.mu.Lock()
	defer a.mu.Unlock()

	if n >= a.promiseID {
		a.promiseID = n
		a.acceptedID = n
		a.acceptedValue = v
		return true
	}

	return false
}

func main() {
	acceptor := Acceptor{}
	proposals := []Proposal{
		{ID: 1, Value: "value1"},
		{ID: 2, Value: "value2"},
		{ID: 3, Value: "value3"},
	}

	for _, prop := range proposals {
		// Prepare Phase
		acceptedID, acceptedValue := acceptor.Prepare(prop.ID)
		fmt.Printf("Prepare Phase: Proposal ID=%d, Accepted ID=%d, Accepted Value=%s\n", prop.ID, acceptedID, acceptedValue)

		// Accept Phase
		accepted := acceptor.Accept(prop.ID, prop.Value)
		fmt.Printf("Accept Phase: Proposal ID=%d, Accepted=%t\n", prop.ID, accepted)
	}
}
