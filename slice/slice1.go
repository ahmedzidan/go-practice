package main

func main() {
	slice := []string{"2", "2", "2"}

	for i, v := range slice {
		slice[1] = "1"
		if i == 1 {
			println("Value is", v)
		}
	}
}
