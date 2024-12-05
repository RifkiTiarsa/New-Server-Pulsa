package config

const (
	SelectMemberById   = `SELECT id, member_id, name, phone_number, address, balance, pin, createdAt, updated_at FROM member WHERE id = ?`
	SelectMemberList   = `SELECT id, member_id, name, phone_number, address, balance, pin, createdAt, updated_at FROM member`
	InsertMember       = `INSERT INTO member (member_id, name, phone_number, address, balance, pin, createdAt, updated_at) VALUES (?,?,?,?,?,?,?,?)`
	UpdateMember       = `UPDATE member SET name = ?, phone_number = ?, address = ?, updated_at = ? WHERE id = ?`
	DeleteMember       = `DELETE FROM member WHERE id = ?`
	CheckBalanceMember = `SELECT balance FROM member WHERE id = ?`
)
