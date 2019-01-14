package persistence

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jojoarianto/go-ddd-api/domain"
	"github.com/jojoarianto/go-ddd-api/domain/repository"
)

// NewsRepositoryImpl Implements repository.NewsRepository
type NewsRepositoryImpl struct {
	Conn *sql.DB
}

// NewNewsRepositoryWithRDB returns initialized NewsRepositoryImpl
func NewNewsRepositoryWithRDB(conn *sql.DB) repository.NewsRepository {
	return &NewsRepositoryImpl{Conn: conn}
}

// Get returns domain.User
func (r *NewsRepositoryImpl) Get(id int) (*domain.News, error) {
	row, err := r.queryRow("select id, title from news where id=?", id)
	if err != nil {
		return nil, err
	}
	u := &domain.News{}
	err = row.Scan(&u.ID, &u.Title)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *NewsRepositoryImpl) query(q string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := r.Conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Query(args...)
}

func (r *NewsRepositoryImpl) queryRow(q string, args ...interface{}) (*sql.Row, error) {
	stmt, err := r.Conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryRow(args...), nil
}
