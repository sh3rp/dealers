package main

func main() {
	city := City{}
	city.Populate(10, 10)
	city.PopulateJunkies()

	city.UpdateJunkies()
}
