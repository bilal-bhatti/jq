package jq

import "log"

type Eq struct {
	Value string `json:"value,omitempty"`
	Other string `json:"other,omitempty"`
}

func (e *Eq) Eval() bool {
	log.Println("doing equal")
	return e.Value == e.Other
}
