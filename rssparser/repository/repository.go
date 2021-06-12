package repository

import (
	"context"
	"io/ioutil"
	"lab8/rssparser/models"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

const userName string = "auli"
const password string = "bq47w8mz"
const dbName string = "test"
const host string = "127.0.0.1"
const port string = "5432"

const pathToCreateTable string = "rssparser/repository/infoRSS.sql"
const pathToDropTable string = "rssparser/repository/dropRSS.sql"

type PostgreRSSRepository struct {
	pool *pgxpool.Pool
}

func NewPostgreRSSRepository() (*PostgreRSSRepository, error) {
	DBUri := "postgresql://" + userName + ":" + password + "@" + host + ":" + port + "/" + dbName
	config, err := pgxpool.ParseConfig(DBUri)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to create config")
	}

	ctx := context.Background()
	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to create pool")
	}
	return &PostgreRSSRepository{pool: pool}, nil
}

func (repo *PostgreRSSRepository) AddOneElementToTable(ctx context.Context, element models.FeedElement) error {
	rows, err := repo.pool.Query(ctx, "insert into infoRSS (title, description, link, published) values($1::text, $2::text, $3::text, $4::text)",
		element.Title,
		element.Description,
		element.Link,
		element.Published)
	if err != nil {
		return errors.Wrap(err, "Unable to addOneElemetToTable")
	}
	rows.Close()

	return nil
}

func (repo *PostgreRSSRepository) ClosePool() {
	repo.pool.Close()
}

func (repo *PostgreRSSRepository) ReadFromTable(ctx context.Context) ([]models.FeedElement, error) {
	rows, err := repo.pool.Query(ctx, "SELECT title, description, link, published from infoRSS")
	if err != nil {
		return []models.FeedElement{}, errors.Wrap(err, "Unable to ReadFromTable")
	}

	defer rows.Close()

	var array []models.FeedElement
	for rows.Next() {
		var el models.FeedElement
		err = rows.Scan(&el.Title, &el.Description, &el.Link, &el.Published)
		if err != nil {
			return []models.FeedElement{}, errors.Wrap(err, "Unable to read from rows")
		}

		array = append(array, el)
	}
	return array, nil
}

func (p *PostgreRSSRepository) CreateTable() error {
	file, err := ioutil.ReadFile(pathToCreateTable)
	if err != nil {
		return errors.Wrap(err, "Unable to read File")
	}

	_, err = p.pool.Exec(context.Background(), string(file))
	if err != nil {
		return errors.Wrap(err, "Unable to Exec")
	}
	return nil
}

func (p *PostgreRSSRepository) DropTable() error {
	file, err := ioutil.ReadFile(pathToDropTable)
	if err != nil {
		return errors.Wrap(err, "Unable to read File")
	}

	_, err = p.pool.Exec(context.Background(), string(file))
	if err != nil {
		return errors.Wrap(err, "Unable to Exec")
	}
	return nil
}
