package repository

import (
	"database/sql"
	"server-pulsa/config"
	"server-pulsa/entity"
)

type memberRepository struct {
	db *sql.DB
}

// Create implements MemberRepository.
func (m *memberRepository) Create(payload *entity.Member) (entity.Member, error) {
	if err := m.db.QueryRow(config.InsertMember, payload.MemberID, payload.Name, payload.Phone, payload.Address, payload.Balance, payload.Pin, payload.UpdatedAt).Scan(&payload.ID, &payload.CreatedAt); err != nil {
		return entity.Member{}, err
	}

	return *payload, nil
}

// Delete implements MemberRepository.
func (m *memberRepository) Delete(id int) error {
	row, err := m.db.Exec(config.DeleteMember, id)
	if err != nil {
		return err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// FindAll implements MemberRepository.
func (m *memberRepository) FindAll() ([]entity.Member, error) {
	var members []entity.Member

	rows, err := m.db.Query(config.SelectMemberList)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var payload entity.Member

		if err := rows.Scan(&payload.ID, &payload.MemberID, &payload.Name, &payload.Phone, &payload.Address, &payload.Balance, &payload.Pin, &payload.CreatedAt, &payload.UpdatedAt); err != nil {
			return nil, err
		}

		members = append(members, payload)
	}

	return members, nil
}

// FindByID implements MemberRepository.
func (m *memberRepository) FindByID(id int) (*entity.Member, error) {
	var member entity.Member

	if err := m.db.QueryRow(config.SelectMemberById, id).Scan(&member.ID, &member.MemberID, &member.Name, &member.Phone, &member.Address, &member.Balance, &member.Pin, &member.CreatedAt, &member.UpdatedAt); err != nil {
		return &entity.Member{}, err
	}

	return &member, nil
}

// Update implements MemberRepository.
func (m *memberRepository) Update(payload *entity.Member) (entity.Member, error) {
	row, err := m.db.Exec(config.UpdateMember, payload.MemberID, payload.Name, payload.Phone, payload.Address, payload.UpdatedAt, payload.ID)
	if err != nil {
		return entity.Member{}, err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return entity.Member{}, err
	}

	if rowsAffected == 0 {
		return entity.Member{}, sql.ErrNoRows
	}

	return *payload, nil
}

type MemberRepository interface {
	FindByID(id int) (*entity.Member, error)
	FindAll() ([]entity.Member, error)
	Create(payload *entity.Member) (entity.Member, error)
	Update(payload *entity.Member) (entity.Member, error)
	Delete(id int) error
}

func NewMemberRepository(db *sql.DB) MemberRepository {
	return &memberRepository{db: db}
}
