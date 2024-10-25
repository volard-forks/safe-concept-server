package lib

func PanicIfError(e error) {
	if e != nil {
		panic(e)
	}
}
