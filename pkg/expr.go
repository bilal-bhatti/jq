package jq

import "log"

type Condition interface {
	Eval() bool
}

type Expr struct {
	Eq  *Eq     `json:"eq,omitempty"`
	Ne  *Ne     `json:"ne,omitempty"`
	Not *Expr   `json:"not,omitempty"`
	And []*Expr `json:"and,omitempty"`
	Or  []*Expr `json:"or,omitempty"`
}

func (e *Expr) Eval() bool {
	if e.Eq != nil {
		return e.Eq.Eval()
	}
	if e.Ne != nil {
		return e.Ne.Eval()
	}
	if e.Not != nil {
		return e.not()
	}
	if e.And != nil {
		return e.and()
	}
	if e.Or != nil {
		return e.or()
	}

	panic("something bad")
}

func (e *Expr) not() bool {
	log.Println("doing not")
	return !e.Not.Eval()
}

func (e *Expr) and() bool {
	log.Println("doing and")
	for _, e := range e.And {
		if e.Eval() == false {
			return false
		}
	}

	return true
}

func (e *Expr) or() bool {
	log.Println("doing or")
	for _, e := range e.Or {
		if e.Eval() == true {
			return true
		}
	}

	return false
}
