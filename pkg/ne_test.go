package jq

import (
	"encoding/json"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNeTrue(t *testing.T) {
	js := `
	{
		"ne": {
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
	assert.True(t, q.Eval(), "Not equal should be true")
}

func TestNeFalse(t *testing.T) {
	js := `
	{
		"ne": {
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
	assert.False(t, q.Eval(), "Not equal should be false")
}
