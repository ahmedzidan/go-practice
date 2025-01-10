package main

type Reader interface {
	read(b []byte) (n int, err error)
}
type File struct {
	name string
}

func (File) read(b []byte) (n int, err error) {
	return 0, nil
}

func retrive(r Reader) {
	r.read([]byte{})
}

func main() {

	f := File{"data.txt"}
	retrive(f)
}
