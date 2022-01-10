package utils

func Ok(err error) {
	if err != nil {
		panic(err)
	}
}
