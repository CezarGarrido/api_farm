package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/CezarGarrido/FarmVue/ApiFarm/entities"
)

// LogRepo explain...
type AreaRepo interface {
	/*Fetch(ctx context.Context, num int64) ([]*entities.Area, error)*/
	//GetByID(ctx context.Context, id int64) (*entities.Log, error)*/
	Create(ctx context.Context, area *entities.Area) (int64, error)
	GetAllByProprietarioID(ctx context.Context, id int64) ([]*entities.Area, error)
	Delete(ctx context.Context, id int64) (bool, error)
	/*Update(ctx context.Context, p *entities.Log) (*entities.Log, error)
	Delete(ctx context.Context, id int64) (bool, error)*/
}

// NewSQLLogRepo retunrs implement of post repository interface
func NewSQLAreaRepo(Conn *sql.DB) AreaRepo {
	return &postgresAreaRepo{Conn: Conn}
}

type postgresAreaRepo struct {
	Conn *sql.DB
}

func (m *postgresAreaRepo) GetAllByProprietarioID(ctx context.Context, id int64) ([]*entities.Area, error) {
	query := "SELECT id,proprietario_id,descricao,area_total,geo_json,created_at,updated_at FROM cadastros.areas WHERE proprietario_id=$1"
	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}
	if len(rows) > 0 {
		return rows, nil
	} else {
		return nil, err
	}
}
func (m *postgresAreaRepo) Create(ctx context.Context, area *entities.Area) (int64, error) {
	query := "INSERT INTO cadastros.areas (proprietario_id,descricao,area_total,geo_json,created_at) VALUES ($1,$2,$3,$4,$5) RETURNING id"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	var id int64
	err = stmt.QueryRowContext(
		ctx,
		area.ProprietarioID,
		area.Descricao,
		area.AreaTotal,
		area.GeoJSON,
		time.Now(),
	).Scan(&id)
	defer stmt.Close()
	if err != nil {
		return -1, err
	}
	return id, nil
}
func (m *postgresAreaRepo) Fetch(ctx context.Context, num int64) ([]*entities.Area, error) {
	query := "SELECT id,proprietario_id,descricao,area_total,geo_json,created_at,updated_at FROM cadastros.areas LIMIT ?"
	return m.fetch(ctx, query, num)
}
func (m *postgresAreaRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*entities.Area, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payload := make([]*entities.Area, 0)
	for rows.Next() {
		data := new(entities.Area)
		err := rows.Scan(
			&data.ID,
			&data.ProprietarioID,
			&data.Descricao,
			&data.AreaTotal,
			&data.GeoJSON,
			&data.CreatedAt,
			&data.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *postgresAreaRepo) Delete(ctx context.Context, id int64) (bool, error) {
	query := "DELETE FROM cadastros.areas WHERE id=$1"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}