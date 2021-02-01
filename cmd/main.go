package main

import "emerson-argueta/m/v2/shared/infrastructure/http"

func main() {
	// gtin12 := "850008366079"
	httpServer := http.NewServer()

	if err := httpServer.Open(); err != nil {
		panic(err)
	}

}
