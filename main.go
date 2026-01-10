package main

func main() {
	setupDatabase()
	defer db.Close()
	handleFunc()
}
