package main

func main() {
	err := startPluginEngine()
	if err != nil {
		panic(err) // Panic to terminate the code execution
	}
}
