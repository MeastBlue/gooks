package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func creatUserTable(db *sqlx.DB) {
	request := fmt.Sprintf(`create table if not exists users(
        id serial primary key,
        lastname varchar(255) not null,
        firstname varchar(255) not null,
        is_admin boolean not null
        )`)

	db.MustExec(request)
}
