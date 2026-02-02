package main

func CloseTraditional() []error {
	errCh1 := Close1()
	errCh2 := Close2()
	errCh3 := Close3()

	var errors []error

	err := <- errCh1
	if err != nil {
		errors = append(errors, err)
	}

	err = <- errCh2
	if err != nil {
		errors = append(errors, err)
	}

	err = <- errCh3
	if err != nil {
		errors = append(errors, err)
	}

	return errors
}

func doClose() error {
	return nil
}

func Close1() chan error {
	errorCh := make(chan error, 1)

	go func(){
		err := doClose()

		errorCh <- err
	}()

	return errorCh
}

func Close2() chan error {
	errorCh := make(chan error, 1)

	go func() {
		var err error

		errorCh <- err
	}()

	return errorCh
}

func Close3() chan error {
	errorCh := make(chan error, 1)

	go func() {
		var err error

		errorCh <- err
	}()

	return errorCh
}