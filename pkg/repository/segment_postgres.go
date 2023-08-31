package repository

import (
	"avito-app"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/jmoiron/sqlx"
)

type SegmentPostgres struct {
	db *sqlx.DB
}

func NewSegmentPostgres(db *sqlx.DB) *SegmentPostgres {
	return &SegmentPostgres{db: db}
}

func (r *SegmentPostgres) CreateSegment(seg avito.Segment) (int, error) {
	var id, misc int

	// check if segment exists
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE seg_name = ($1) LIMIT 1", segmentsTable)
	row := r.db.QueryRow(query, seg.Seg_name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	if int(id) != 0 {
		return 0, fmt.Errorf("Segment '%s' already exists", seg.Seg_name)
	}

	query = fmt.Sprintf("INSERT INTO %s (seg_name) values ($1) RETURNING SEG_ID", segmentsTable)
	row = r.db.QueryRow(query, seg.Seg_name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	if seg.Chance != 0 {
		// get the total number of users into the 'id' variable
		slice := []string{}
		query = fmt.Sprintf("SELECT DISTINCT user_id FROM %s", userSegTable)
		if err := r.db.Select(&slice, query); err != nil {
			return 0, err
		}

		// get the amount of users to add to the new segment
		var res = math.Ceil(float64(len(slice)) * float64(seg.Chance) / 100)
		rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })

		for index := 0; index < int(res); index++ {
			// add random users to the new segment
			query = fmt.Sprintf("INSERT INTO %s (user_id, seg_id) values ($1,$2) RETURNING USER_ID", userSegTable)
			row = r.db.QueryRow(query, slice[index], id)
			if err := row.Scan(&misc); err != nil {
				return 0, err
			}

			// put records into "operations" table
			query = fmt.Sprintf("INSERT INTO %s (user_id, seg_id, operation, at_timestamp, TTL) values ($1, $2, $3, $4, $5) RETURNING USER_ID", operationsTable)
			row = r.db.QueryRow(query, slice[index], id, "ADD", time.Now(), 0)
			if err := row.Scan(&misc); err != nil {
				return 0, err
			}
		}
	}

	return id, nil
}

func (r *SegmentPostgres) DeleteSegment(seg avito.Segment) (int, error) {
	var id int

	// check if segment exists
	query := fmt.Sprintf("SELECT seg_id FROM %s WHERE seg_name = ($1) LIMIT 1", segmentsTable)
	row := r.db.QueryRow(query, seg.Seg_name)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("Segment '%s' does not exist", seg.Seg_name)
	}

	// find users in that segment
	usrSlice := []int{}
	query = fmt.Sprintf("SELECT DISTINCT user_id FROM %s WHERE seg_id = ($1)", userSegTable)
	err := r.db.Select(&usrSlice, query, id)
	if err != nil {
		return 0, err
	}

	// put records into "operations" table
	for _, usrid := range usrSlice {
		query = fmt.Sprintf("INSERT INTO %s (user_id, seg_id, operation, at_timestamp, TTL) values ($1, $2, $3, $4, $5) RETURNING USER_ID", operationsTable)
		row = r.db.QueryRow(query, usrid, id, "DEL", time.Now(), 0)
		if err := row.Scan(&id); err != nil {
			return 0, err
		}
	}

	// delete the segment
	query = fmt.Sprintf("DELETE FROM %s WHERE seg_name = ($1)", segmentsTable)
	row = r.db.QueryRow(query, seg.Seg_name)

	if err := row.Scan(&id); id == 0 && err != nil {
		return 0, err
	}

	return 200, nil
}
