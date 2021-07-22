// model.go

package main

import (
	"context"
	"database/sql"
	//"fmt"
	"strings"
)

type contact struct {
	Id         int      `json:"id,omitempty"`
	First_name string   `json:"first_name"`
	Last_name  string   `json:"last_name"`
	Email      string   `json:"email"`
	Phones     []string `json:"phone_numbers"`
}

func (c *contact) createContact(db *sql.DB) error {
	//TODO
	//Not tested well, Idea here was to save the data in two tables, contact and phones, and
	//have these insert queries in single transaction scope, so that if someting breaks, it still remains clean.

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}
	err1 := tx.QueryRow("with ins as ( insert into contacts ( first_name, last_name,email) values ($1,S2,$3) returning id) select id from ins;", c.First_name, c.Last_name, c.Email).Scan(&c.Id)

	if err1 != nil {
		tx.Rollback()
		return err1
	}
	if c.Id > 0 && len(c.Phones) > 0 {
		err2 := tx.QueryRow("insert into contact_phones(contact_id,phone_number) values select $1, regexp_split_to_table($2, E',') as c ;")
		if err2 != nil {
			tx.Rollback()
			return nil
		}
	}
	err = tx.Commit()

	return nil
}

func getContacts(db *sql.DB, start, count int) ([]contact, error) {

	rows, err := db.Query(
		"SELECT id, first_name,last_name, email, (SELECT array_to_string( array( SELECT concat('\"',phone_number,'\"') FROM contact_phones p Where p.contact_id = c.id ), ',' )) as phones from contacts c  LIMIT $1 OFFSET $2", count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	contacts := []contact{}

	var phStr string

	for rows.Next() {
		var c contact
		c.Phones = make([]string, 0)

		if err := rows.Scan(&c.Id, &c.First_name, &c.Last_name, &c.Email, &phStr); err != nil {
			return nil, err
		}
		if len(phStr) > 0 {
			c.Phones = strings.Split(phStr, ",")
		}
		contacts = append(contacts, c)
	}
	return contacts, nil
}
