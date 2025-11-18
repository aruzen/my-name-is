package repository

import (
	"backend/internal/domain"
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HueRepository struct {
	db *pgxpool.Pool
}

func NewHueRepository(db *pgxpool.Pool) *HueRepository {
	return &HueRepository{db: db}
}

// Save は hue_records テーブルへ新しいレコードを保存する。
func (r *HueRepository) Save(ctx context.Context, record domain.HueRecord) error {
	const query = `
		INSERT INTO hue_records (id, user_name, choices)
		VALUES ($1, $2, $3)
	`

	choiceJSON, err := json.Marshal(record.ChoiceMap())
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, record.ID(), record.Name().String(), choiceJSON)
	return err
}

// FindRange は作成順で並んだレコードの指定範囲を返す。
func (r *HueRepository) FindRange(ctx context.Context, recordRange domain.RecordRange) ([]domain.HueRecord, error) {
	const query = `
		SELECT id, user_name, choices
		FROM hue_records
		ORDER BY created_at, id
		OFFSET $1
		LIMIT $2
	`

	rows, err := r.db.Query(ctx, query, recordRange.Begin(), recordRange.Count())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []domain.HueRecord
	for rows.Next() {
		record, err := scanHueRecord(rows)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func scanHueRecord(row rowScanner) (domain.HueRecord, error) {
	var (
		id         uuid.UUID
		userName   string
		choiceJSON []byte
	)

	if err := row.Scan(&id, &userName, &choiceJSON); err != nil {
		return domain.HueRecord{}, err
	}

	name, err := domain.NewName(userName)
	if err != nil {
		return domain.HueRecord{}, err
	}

	var raw map[string]string
	if err := json.Unmarshal(choiceJSON, &raw); err != nil {
		return domain.HueRecord{}, err
	}

	choices, err := domain.NewHueChoices(raw)
	if err != nil {
		return domain.HueRecord{}, err
	}

	return domain.NewHueRecordFromPersistence(id, name, choices)
}
