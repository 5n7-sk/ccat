package main

// Option represents application options
type Option struct {
	Number bool `short:"n" long:"number" description:"Show contents with line numbers"`
}

// Config represents the settings for this application
type Config struct {
	Theme string `json:"theme"`
}
