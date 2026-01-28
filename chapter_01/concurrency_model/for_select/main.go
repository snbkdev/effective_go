package main

func SelectExample(dataCh chan Data, stopCh chan struct{}) {
	defer close(dataCh)

	for{
		select{
		case data:= <- dataCh:
			processData(data)
		case <- stopCh:
			return
		}
	}
}

func processData(data Data) {}

type Data struct {}