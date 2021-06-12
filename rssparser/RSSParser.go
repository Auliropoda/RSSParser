package rssparser

import (
	"context"
	"fmt"
	"lab8/rssparser/parser"
	"lab8/rssparser/repository"
	"log"

	"github.com/pkg/errors"
)

type RSSParser struct {
	repo *repository.PostgreRSSRepository
}

func NewRSSParser() (*RSSParser, error) {
	repo, err := repository.NewPostgreRSSRepository()
	if err != nil {
		return nil, errors.Wrap(err, "Unable to create new RSS Parser")
	}
	return &RSSParser{repo}, nil
}

func (p *RSSParser) SaveData() error {
	err := p.repo.CreateTable()
	if err != nil {
		return errors.Wrap(err, "Unable to create table")
	}

	array, err := parser.Parse()
	if err != nil {
		return errors.Wrap(err, "Unable to Parse")
	}

	for _, element := range array {
		err = p.repo.AddOneElementToTable(context.Background(), element)
		if err != nil {
			return errors.Wrap(err, "Unable to save data from RSS parser")
		}
	}
	return nil
}

func (p *RSSParser) ShowData() {
	result, err := p.repo.ReadFromTable(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, res := range result {
		fmt.Println(res)
	}
}

func (p *RSSParser) CLoseConnection() {
	p.repo.ClosePool()
}

func (p *RSSParser) DropTable() error {
	return p.repo.DropTable()
}
