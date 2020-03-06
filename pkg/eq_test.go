package jq

import (
	"encoding/json"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqTrue(t *testing.T) {
	js := `
	{
		"eq": {
			"value": "one",
			"other": "one"
		}
	}
	`

	q := &Expr{}

	err := json.NewDecoder(strings.NewReader(js)).Decode(q)
	if err != nil {
		log.Println("error", err)
	}
	assert.True(t, q.Eval(), "Equal should be true")
}

func TestEqFalse(t *testing.T) {
	js := `
	{
		"eq": {
			"value": "one",
			"other": "two"
		}
	}
	`

	q := &Expr{}

	err := json.NewDecoder(strings.NewReader(js)).Decode(q)
	if err != nil {
		log.Println("error", err)
	}
	assert.False(t, q.Eval(), "Equal should be false")
}
