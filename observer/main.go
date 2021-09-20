package main

func main() {
	listener1 := DataListener{
		Name: "Listener1",
	}

	listener2 := DataListener{
		Name: "Listener2",
	}

	subj := &DataSubject{}
	subj.registerObserver(listener1)
	subj.registerObserver(listener2)

	subj.changeItem("Changed Item")
}
