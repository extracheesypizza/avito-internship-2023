package repository

import (
	"avito-app"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetUserSegments(usr_id int) ([]string, error) {
	slice := []string{}
	var id int

	// update user segments
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND seg_id in (SELECT DISTINCT seg_id FROM %s WHERE (operation = 'DEL' AND user_id = $1 AND (current_timestamp > at_timestamp))) ", userSegTable, operationsTable)
	row := r.db.QueryRow(query, usr_id)
	if err := row.Scan(&id); id != 0 && err != nil {
		fmt.Println(id)
		return []string{}, err
	}

	// get user segments
	query = fmt.Sprintf("SELECT DISTINCT seg_name FROM %s WHERE seg_id in (SELECT DISTINCT seg_id FROM %s WHERE (operation = 'ADD' AND user_id = $1 AND ((current_timestamp BETWEEN at_timestamp AND (at_timestamp + TTL * interval '1 second') OR TTL = 0))))", segmentsTable, operationsTable)
	err := r.db.Select(&slice, query, usr_id)
	return slice, err
}

func (r *UserPostgres) AddUserToSegment(usr avito.User) (int, error) {
	var s_id, id int

	for _, x := range usr.Seg_names {
		// check if segment exists
		query := fmt.Sprintf("SELECT seg_id FROM %s WHERE seg_name = ($1) LIMIT 1", segmentsTable)
		row := r.db.QueryRow(query, x)
		if err := row.Scan(&s_id); err != nil {
			return 0, fmt.Errorf("Segment '%s' does not exist", x)
		}

		// check if user was in a group
		query = fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE user_id = ($1) AND seg_id = ($2)", userSegTable)
		row = r.db.QueryRow(query, usr.Id, s_id)

		if err := row.Scan(&id); err != nil {
			return 0, err
		} else if int(id) > 0 {
			return 0, fmt.Errorf("User '%d' is already in segment '%s'", usr.Id, x)
		}

		// create a new row
		query = fmt.Sprintf("INSERT INTO %s (user_id, seg_id) values ($1, $2) RETURNING USER_ID", userSegTable)
		row = r.db.QueryRow(query, usr.Id, s_id)

		if err := row.Scan(&id); err != nil {
			return 0, err
		}

		// put a record into "operations" table
		query = fmt.Sprintf("INSERT INTO %s (user_id, seg_id, operation, at_timestamp, TTL) values ($1, $2, $3, current_timestamp, $4) RETURNING USER_ID", operationsTable)
		row = r.db.QueryRow(query, usr.Id, s_id, "ADD", usr.TTL)
		if err := row.Scan(&id); err != nil {
			return 0, err
		}

		if usr.TTL > 0 {
			query = fmt.Sprintf("INSERT INTO %s (user_id, seg_id, operation, at_timestamp, TTL) values ($1, $2, $3, current_timestamp + $4 * interval '1 second', 0) RETURNING USER_ID", operationsTable)
			row = r.db.QueryRow(query, usr.Id, s_id, "DEL", usr.TTL)
			if err := row.Scan(&id); err != nil {
				return 0, err
			}
		}
	}

	return 200, nil
}

func (r *UserPostgres) DeleteUserFromSegment(usr avito.User) (int, error) {
	var s_id, id int

	for _, x := range usr.Seg_names {
		// check if segment exists
		query := fmt.Sprintf("SELECT seg_id FROM %s WHERE seg_name = ($1)", segmentsTable)
		row := r.db.QueryRow(query, x)

		if err := row.Scan(&s_id); err != nil {
			return 0, fmt.Errorf("Segment '%s' does not exist", x)
		}

		// check if user was in a group
		query = fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE user_id = ($1) AND seg_id = ($2)", userSegTable)
		row = r.db.QueryRow(query, usr.Id, s_id)

		if err := row.Scan(&id); err != nil {
			return 0, err
		} else if int(id) == 0 {
			return 0, fmt.Errorf("User '%d' is not in segment '%s'", usr.Id, x)
		}

		// delete all corresponding rows
		query = fmt.Sprintf("DELETE FROM %s WHERE (user_id = ($1) AND seg_id = ($2))", userSegTable)
		row = r.db.QueryRow(query, usr.Id, s_id)

		if err := row.Scan(&id); id == 0 && err != nil {
			return 0, err
		}

		// put a record into "operations" table
		query = fmt.Sprintf("INSERT INTO %s (user_id, seg_id, operation, at_timestamp, TTL) values ($1, $2, $3, current_timestamp, 0) RETURNING USER_ID", operationsTable)
		row = r.db.QueryRow(query, usr.Id, s_id, "DEL")
		if err := row.Scan(&id); err != nil {
			return 0, err
		}
	}

	return 200, nil
}

func (r *UserPostgres) GetUserActions(usr avito.User) ([]string, error) {

	type Action struct {
		User_id      int    `db:"user_id"`
		Seg_id       int    `db:"seg_id"`
		Operation    string `db:"operation"`
		At_timestamp string `db:"at_timestamp"`
	}

	slice := []string{}
	temp := []Action{}

	// get user segments
	query := fmt.Sprintf("SELECT \"user_id\", \"seg_id\", \"operation\", \"at_timestamp\" FROM %s WHERE (user_id = $1 AND EXTRACT(month FROM at_timestamp ) >= EXTRACT(month FROM $2 * interval '1 month') AND EXTRACT(year FROM at_timestamp) >= EXTRACT(year FROM $3 * interval '1 year'))", operationsTable)
	err := r.db.Select(&temp, query, usr.Id, usr.Month, usr.Year)

	for _, x := range temp {
		s := strconv.Itoa(int(x.User_id)) + ";" + strconv.Itoa(int(x.Seg_id)) + ";" + x.Operation + ";" + x.At_timestamp + ";"
		slice = append(slice, s)
	}

	return slice, err
}
