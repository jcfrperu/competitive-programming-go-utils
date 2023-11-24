package cp

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
