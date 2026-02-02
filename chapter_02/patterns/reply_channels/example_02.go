package main

func CloseReplyChannel() []error {
	errorCh := make(chan error, 3)

	go CloseA(errorCh)
	go CloseB(errorCh)
	go CloseC(errorCh)

	var errors []error
	for x := 0; x < 3; x++ {
		err := <- errorCh
		errors = append(errors, err)
	}

	return errors
}

func CloseA(errorCh chan error) {
	err := doClose()

	errorCh <- err
}

func CloseB(errorCh chan error) {
	var err error
	
	errorCh <- err
}

func CloseC(errorCh chan error) {
	var err error

	errorCh <- err
}