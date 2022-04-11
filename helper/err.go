package helper

/**
 * todo: for panic
 */
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
