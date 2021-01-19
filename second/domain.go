package main

import "time"

type Metrics struct {
	Id    string    `json:"id,omitempty"`
	Value string    `json:"value"`
	Type  string    `json:"type"`
	Time  time.Time `json:"time,omitempty"`
}
