package main

// BasicControl ...
type BasicControl struct {
	forwardBackward string
	leftRight       string
}

// Message ...
type Message struct {
	ButtonName string `json:"ButtonName"`
}
