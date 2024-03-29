package store

import (
	"database/sql"
	"log"
	"politics/.gen/model"
	"politics/.gen/table"
	"politics/internal/types"

	. "github.com/go-jet/jet/v2/sqlite"
)

type Storage struct {
	db *sql.DB
}

type Store interface {
	SaveQuery(text string) error
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) GetQuery(id int32) (model.Query, error) {
	query := model.Query{}
	log.Printf("Getting query %v", id)
	selectStmt := table.Query.SELECT(table.Query.AllColumns).FROM(table.Query).WHERE(table.Query.ID.EQ(Int(int64(id))))
	err := selectStmt.Query(s.db, &query)
	if err != nil {
		log.Fatalf("Error getting query: %v", err)
		return model.Query{}, err
	}
	return query, nil
}

func (s *Storage) GetQueries() ([]model.Query, error) {
	var queries []model.Query
	selectStmt := table.Query.SELECT(table.Query.AllColumns).FROM(table.Query)
	err := selectStmt.Query(s.db, &queries)
	if err != nil {
		log.Fatalf("Error getting queries: %v", err)
		return nil, err
	}
	return queries, nil
}

func (s *Storage) SaveQuery(text string) (model.Query, error) {
	query := model.Query{
		Text: text,
	}

	insertStmt := table.Query.INSERT(table.Query.MutableColumns).MODELS([]model.Query{query}).RETURNING(table.Query.AllColumns)

	dest := model.Query{}
	err := insertStmt.Query(s.db, &dest)
	if err != nil {
		log.Fatalf("Error inserting query: %v", err)
		return model.Query{}, err
	}
	return dest, nil
}

func (s *Storage) GetParties() ([]model.Party, error) {
	var parties []model.Party
	selectStmt := table.Party.SELECT(table.Party.AllColumns)
	err := selectStmt.Query(s.db, &parties)
	if err != nil {
		log.Fatalf("Error getting parties: %v", err)
		return nil, err
	}
	return parties, nil
}

func (s *Storage) SaveStance(stance model.Stance) error {
	insertStmt := table.Stance.INSERT(table.Stance.MutableColumns).MODELS([]model.Stance{stance})
	_, err := insertStmt.Exec(s.db)
	if err != nil {
		log.Fatalf("Error saving stance: %v", err)
		return err
	}
	return nil
}

func (s *Storage) GetStances() ([]model.Stance, error) {
	var stances []model.Stance
	selectStmt := table.Stance.SELECT(table.Stance.AllColumns).FROM(table.Stance)
	err := selectStmt.Query(s.db, &stances)
	if err != nil {
		log.Fatalf("Error getting stances: %v", err)
		return nil, err
	}
	return stances, nil
}

func (s *Storage) GetStancesForQuery(queryID int32) ([]types.EnrichedStanced, error) {
	var stances []types.EnrichedStanced
	selectStmt := table.Stance.SELECT(
		table.Stance.AllColumns,
		table.Party.AllColumns,
	).FROM(
		table.Stance.
			INNER_JOIN(table.Party, table.Stance.PartyID.EQ(table.Party.ID)).
			INNER_JOIN(table.Query, table.Stance.QueryID.EQ(table.Query.ID)),
	).WHERE(table.Stance.QueryID.EQ(Int(int64(queryID))))

	err := selectStmt.Query(s.db, &stances)
	if err != nil {
		log.Fatalf("Error getting stances for query: %v", err)

	}
	return stances, nil
}
