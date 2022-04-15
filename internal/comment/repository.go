package comment

import (
	"context"
	"github.com/arivictor/go-api/internal/database"
	uuid "github.com/satori/go.uuid"
)

type Repository struct {
	database *database.Database
}

func NewRepository(database *database.Database) Repository {
	return Repository{database: database}
}

type Record struct {
	ID     string
	Slug   string `json:"slug"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

type Records []Record

func (r Repository) Get(ctx context.Context, id string) (Record, error) {
	var record Record

	var query = `SELECT id, slug, body, author FROM comments WHERE id = $1`
	row := r.database.Client.QueryRowContext(ctx, query, id)

	err := row.Scan(&record.ID, &record.Slug, &record.Body, &record.Author)
	if err != nil {
		return record, err
	}
	return record, nil
}

func (r Repository) Delete(ctx context.Context, id string) (Record, error) {
	var record Record

	query := `DELETE FROM comments where id = $1`
	rows := r.database.Client.QueryRowContext(ctx, query, id)

	err := rows.Scan(&record.ID, &record.Slug, &record.Body, &record.Author)
	if err != nil {
		return record, err
	}
	return record, nil
}

func (r Repository) Update(ctx context.Context, id string, comment Comment) (Record, error) {
	record := Record{
		ID:     id,
		Slug:   comment.Slug,
		Body:   comment.Body,
		Author: comment.Author,
	}
	query := `UPDATE comments SET slug = :slug, author = :author, body = :body WHERE id = :id`
	rows, err := r.database.Client.NamedQueryContext(ctx, query, record)
	defer rows.Close()

	if err != nil {
		return Record{}, err
	}
	return record, nil
}

func (r Repository) Create(ctx context.Context, comment Comment) (Record, error) {
	record := Record{
		ID:     uuid.NewV4().String(),
		Slug:   comment.Slug,
		Body:   comment.Body,
		Author: comment.Author,
	}
	query := `INSERT INTO comments (id, slug, author, body) VALUES (:id, :slug, :author, :body)`
	rows, err := r.database.Client.NamedQueryContext(ctx, query, record)
	defer rows.Close()

	if err != nil {
		return Record{}, err
	}
	return record, nil
}

func (r Repository) List(ctx context.Context) (Records, error) {
	var records Records
	rows, _ := r.database.Client.Query(`SELECT id, slug, body, author FROM comments`)
	defer rows.Close()

	for rows.Next() {
		var record Record
		err := rows.Scan(&record.ID, &record.Slug, &record.Body, &record.Author)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}
