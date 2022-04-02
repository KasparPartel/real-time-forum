package helper

// CheckError checks for error and panics
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
