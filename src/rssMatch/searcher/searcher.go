package searcher

import (
	"log"
	"rssMatch/types"
	"sync"
)


type Match interface {
	Search(searchTerm string,fe *Feed) ([]*types.Result,error)
}
var matchs = make(map[string]Match)
var resultChan chan *types.Result

const filename  = "data.json"

func Run(searchTerm string) {
	var waitGroup sync.WaitGroup
	feeds,err := Retrieve(filename)
	if err != nil {
		log.Fatalf("retrieve file faild %v\n",err)
	}
	waitGroup.Add(len(feeds))
	for _, feed := range feeds {
		go func() {
			match := matchs[feed.Type]
			results,err :=  match.Search(searchTerm,&feed)
			if err != nil {
				log.Println(err)
			}
			for  _, result := range results {
				resultChan <- result
			}
			waitGroup.Done()
		}()
	}

	go func() {
		waitGroup.Wait()
		close(resultChan)
	}()
	Display(resultChan)

}

func Register(name string, match Match) {
	if _,ok := matchs[name];ok {
		log.Printf("Matcher %s already registered.\n",name)
		return
	}
	matchs[name] = match
	log.Printf("Matcher %s registered success.\n",name)
}


func Display(result <- chan *types.Result) {
	for re := range result {
		log.Printf("%s:%s\n",re.Field,re.Content)
	}
}