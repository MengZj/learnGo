package searcher

import (
	"encoding/json"
	"os"
)

//Feed keep decode result of json
type Feed struct {
	Name 	string `json:"site"`
	Url 	string `json:"link"`
	Type 	string `json:"type"`
}

func Retrieve(file string) ([]Feed,error) {
	f, err := os.Open(file)
	if err != nil {
		return nil,err
	}
	defer f.Close()
	var feeds []Feed
	err = json.NewDecoder(f).Decode(&feeds)
	return feeds,nil
}