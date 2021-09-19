package main

func main() {
	tv1 := &SammysangTV{
		currentChan:   13,
		currentVolume: 35,
		tvOn:          true,
	}
	tv2 := &SohneeTV{
		vol:     20,
		channel: 9,
		isOn:    true,
	}

	tv2.turnOn()
	tv2.volumeUp()
	tv2.volumeDown()
	tv2.turnOff()

	ssAdapt := &sammysangAdapter{
		sstv: tv1,
	}

	ssAdapt.turnOn()
	ssAdapt.volumeUp()
	ssAdapt.volumeDown()
	ssAdapt.turnOff()
}
