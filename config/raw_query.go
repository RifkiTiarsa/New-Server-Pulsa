package config

const (
	SelectMemberById = `SELECT id, member_id, name, phone_number, address, balance, pin, createdAt, updatedAt WHERE id = ?`
	SelectMemberList = `SELECT id, member_id, name, phone_number, address, balance, pin, createdAt, updatedAt`
	InsertMember     = `INSERT INTO member (member_id, name, phone_number, address, balance, pin, updatedAt) VALUES (?,?,?,?,?,?) RETURNING id, balance, pin, createdAt`
	UpdateMember     = `UPDATE member SET member_id=?, name=?, phone_number=?, address=?, updatedAt=? WHERE id=?`
	DeleteMember     = `DELETE FROM member WHERE id=?`
)
