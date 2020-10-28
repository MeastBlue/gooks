package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func creatBookTable(db *sqlx.DB) {
	request := fmt.Sprintf(`create table if not exists book(
        id serial primary key,
        title varchar(255) not null,
        author varchar(255) not null
        )`)

	db.MustExec(request)
}
