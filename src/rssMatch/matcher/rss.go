package matcher

import (
	"encoding/xml"
	"errors"
	"log"
	"net/http"
	"regexp"
	"rssMatch/searcher"
	"rssMatch/types"
)

type rssMatcher struct {}

func init() {
	var matcher rssMatcher
	searcher.Register("rss",&matcher)
}

func (rss *rssMatcher) Search(item string,feed *searcher.Feed) ([]*types.Result,error){
	var results []*types.Result

	log.Printf("Search Feed Type[%s] Site[%s] For URI[%s]\n", feed.Type, feed.Name, feed.Url)

	rssdoc,err := download(feed)
	if err != nil {
		return nil,err
	}
	for _,channelitem := range rssdoc.Channel.Item {
		match,err := regexp.MatchString(item,channelitem.Title)
		if err != nil {
			return nil,err
		}
		if match {
			results = append(results,&types.Result{
				Field:"Title",
				Content:channelitem.Title,
			})
		}
		match,err = regexp.MatchString(item,channelitem.Description)
		if err != nil {
			return nil,err
		}
		if match {
			results = append(results,&types.Result{
				Field:"Description",
				Content:channelitem.Description,
			})
		}

	}
	return results,nil


}

func download(feed *searcher.Feed) (*types.RssDocument,error) {
	if feed.Url == "" {
		return nil,nil
	}

	resp, err := http.Get(feed.Url)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil,errors.New("bad statusCode")
	}
	var rssdoc types.RssDocument
	err = xml.NewDecoder(resp.Body).Decode(&rssdoc)
	return &rssdoc,err

}
