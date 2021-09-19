package main

// This is the television interface we want to use with both TV types
type television interface {
	volumeUp() int
	volumeDown() int
	channelUp() int
	channeDown() int
	turnOn() int
	turnOff() int
	goToChannel(ch int)
}
