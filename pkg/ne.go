package jq

import "log"

type Ne struct {
	Eq
}

func (n *Ne) Eval() bool {
	log.Println("doing not equal")
	return n.Value != n.Other
}
