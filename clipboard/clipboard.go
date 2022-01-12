package clipboard

func Read() ([]byte, error) {
	return read()
}

func Write(buff []byte) error {
	return write(buff)
}
