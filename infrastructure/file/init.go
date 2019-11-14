package file

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
)

func Load(path string) []int {
	//filename is the path to the json config file
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Panic(err)
		}
	}()

	var init []int
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&init)
	if err != nil {
		log.Panic(err)
	}

	return init
}
