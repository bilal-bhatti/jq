package jq

import (
	"encoding/json"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotTrue(t *testing.T) {
	js := `
	{
		"not": {
				"eq": {
					"value": "one",
					"other": "two"
				}
			}
		}
	}
	`

	q := &Expr{}

	err := json.NewDecoder(strings.NewReader(js)).Decode(q)
	if err != nil {
		log.Println("error", err)
	}
	assert.True(t, q.Eval(), "Not should be true")
}

func TestAndTrue(t *testing.T) {
	js := `
	{
		"and": [
			{
				"eq": {
					"value": "one",
					"other": "one"
				}
			},
			{
				"eq": {
					"value": "two",
					"other": "two"
				}
			}
		]
	}
	`

	q := &Expr{}

	err := json.NewDecoder(strings.NewReader(js)).Decode(q)
	if err != nil {
		log.Println("error", err)
	}
	assert.True(t, q.Eval(), "And should be true")
}

func TestAndFalse(t *testing.T) {
	js := `
	{
		"and": [
			{
				"eq": {
					"value": "one",
					"other": "one"
				}
			},
			{
				"eq": {
					"value": "two",
					"other": "tow"
				}
			}
		]
	}
	`

	q := &Expr{}

	err := json.NewDecoder(strings.NewReader(js)).Decode(q)
	if err != nil {
		log.Println("error", err)
	}
	assert.False(t, q.Eval(), "And should be true")
}

func TestOrTrue(t *testing.T) {
	js := `
	{
		"or": [
			{
				"eq": {
					"value": "one",
					"other": "one"
				}
			},
			{
				"eq": {
					"value": "two",
					"other": "tow"
				}
			}
		]
	}
	`

	q := &Expr{}

	err := json.NewDecoder(strings.NewReader(js)).Decode(q)
	if err != nil {
		log.Println("error", err)
	}
	assert.True(t, q.Eval(), "Or expression should be true")
}

func TestOrFalse(t *testing.T) {
	js := `
	{
		"or": [
			{
				"eq": {
					"value": "one",
					"other": "oen"
				}
			},
			{
				"eq": {
					"value": "two",
					"other": "tow"
				}
			}
		]
	}
	`

	q := &Expr{}

	err := json.NewDecoder(strings.NewReader(js)).Decode(q)
	if err != nil {
		log.Println("error", err)
	}
	assert.False(t, q.Eval(), "Or expression should be false")
}

func TestComplexTrue(t *testing.T) {
	js := `
	{
		"or": [
			{
				"eq": {
					"value": "one",
					"other": "two"
				}
			},
			{
				"and": [
					{
						"ne": {
							"value": "one",
							"other": "bar"
						}
					},
					{
						"eq": {
							"value": "two",
							"other": "two"
						}
					},
					{
						"or": [
							{
								"eq": {
									"value": "one",
									"other": "two"
								}
							},
							{
								"eq": {
									"value": "one",
									"other": "one"
								}
							}
						]
					}
				]
			}
		]
	}
	`

	q := &Expr{}

	err := json.NewDecoder(strings.NewReader(js)).Decode(q)
	if err != nil {
		log.Println("error", err)
	}
	assert.True(t, q.Eval(), "Or expression should be true")
}
