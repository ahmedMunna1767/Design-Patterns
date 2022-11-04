package main

import "fmt"

type Client struct {
}

// 3rd party client who can only consume lightning port from computers
// when only mac has lightning port but windows doesn't have
func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
