package record

func expectNoError(err error) {
	if err != nil {
		panic(err.Error())
	}
}