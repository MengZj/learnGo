package main

import (
	"log"
	"os"
	_ "rssMatch/matcher"
	"rssMatch/searcher"
)
func init() {
	log.SetOutput(os.Stdout)
}

func main(){
	searcher.Run("Python")
}
