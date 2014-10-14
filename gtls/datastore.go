package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type LinkStore struct {
	db *sql.DB
}

func NewLinkStore(host string, port uint16, user, password string) *LinkStore {
	ls := new(LinkStore)

	config := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, "postgres")
	db, err := sql.Open("postgres", config)
	if err != nil {
		log.Fatal(90839873, err)
	}

	ls.db = db

	err = db.Ping()
	if err != nil {
		log.Fatal(3080120555, err)
	}
	return ls
}

func (ls LinkStore) AddShortlink(code, url string) error {
	err := ls.db.QueryRow(`INSERT INTO links(code, url) VALUES($1, $2)`, code, url).Scan()
	return err
}

func (ls LinkStore) GetShortlink(code string) (string, error) {
	var url string
	err := ls.db.QueryRow("SELECT url FROM links WHERE code = $1", code).Scan(&url)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil

		} else {
			return "", err
		}
	}
	return url, nil
}

func (ls LinkStore) GetAllShortlinks() ([]ShortlinkTemplateData, error) {
	links := []ShortlinkTemplateData{}

	rows, err := ls.db.Query("SELECT code, url from links")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// scan all rows
	for rows.Next() {
		var code string
		var url string
		err := rows.Scan(&code, &url)
		if err != nil {
			return nil, err
		}
		links = append(links, ShortlinkTemplateData{code, url})
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return links, nil
}
