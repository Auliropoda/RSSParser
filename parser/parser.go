package parser

import (
	"io/ioutil"
	"lab8/models"
	"net/http"

	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
)

const nameAdress string = "https://www.forbes.com/investing/feed2/"

func getData() (string, error) {
	resp, err := http.Get(nameAdress)
	if err != nil {
		return "", errors.Wrap(err, "Unable to get data")
	}

	defer resp.Body.Close() // вызывается при выходе из функции

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "Unable to read data")
	}

	return string(data), nil
}

func Parse() ([]models.FeedElement, error) {
	data, err := getData()
	if err != nil {
		return []models.FeedElement{}, errors.Wrap(err, "No Data")
	}
	fp := gofeed.NewParser()
	feed, err := fp.ParseString(data)
	if err != nil {
		return []models.FeedElement{}, errors.Wrap(err, "Unable to Parse data")
	}

	var feedElements []models.FeedElement
	for _, value := range feed.Items {
		feedElements = append(feedElements, models.FeedElement{
			Title:       value.Title,
			Description: value.Description,
			Link:        value.Link,
			Published:   value.Published,
		})
		//fmt.Println(value.Title, value.Description, value.Link, value.Published)
	}
	return feedElements, nil
}
