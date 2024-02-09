package state

import (
	"encoding/json"
	"log"
	"os"
)

const filePath = "~/.emuk.json"

func Save() {
	fp, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer fp.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(fp).Encode(struct {
		a int
	}{
		a: 24,
	})
	if err != nil {
		log.Fatal(err)
	}
}
