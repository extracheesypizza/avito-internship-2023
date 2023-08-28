package repository

import (
	"avito-app"
	"fmt"
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
	var id int

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
