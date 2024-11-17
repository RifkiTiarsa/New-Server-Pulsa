package usecase

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"server-pulsa/entity"
	"server-pulsa/repository"
	"strconv"
	"time"
)

type memberUsecase struct {
	repo repository.MemberRepository
}

func (m *memberUsecase) generateMemberId() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(99999))
	return fmt.Sprintf("SP%05d", n.Int64())
}

func (m *memberUsecase) generateRandomPin() (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(9000))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()) + 1000, nil
}

// Create implements MemberUsecase.
func (m *memberUsecase) Create(payload entity.Member) (entity.Member, error) {

	payload.MemberID = m.generateMemberId()

	payload.Balance = 0

	pinInt, _ := m.generateRandomPin()
	payload.Pin = strconv.Itoa(pinInt)

	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()

	fmt.Println(payload)

	return m.repo.Create(payload)
}

// Delete implements MemberUsecase.
func (m *memberUsecase) Delete(id int) error {
	member, err := m.repo.FindByID(id)
	if err != nil {
		return err
	}

	if err = m.repo.Delete(member.ID); err != nil {
		return err
	}

	return nil
}

// FindAll implements MemberUsecase.
func (m *memberUsecase) FindAll() ([]entity.Member, error) {
	return m.repo.FindAll()
}

// FindByID implements MemberUsecase.
func (m *memberUsecase) FindByID(id int) (entity.Member, error) {
	return m.repo.FindByID(id)
}

// Update implements MemberUsecase.
func (m *memberUsecase) Update(payload entity.Member) (entity.Member, error) {
	payload.UpdatedAt = time.Now()

	_, err := m.repo.FindByID(payload.ID)
	if err != nil {
		return entity.Member{}, err
	}

	return m.repo.Update(payload)

}

type MemberUsecase interface {
	Create(payload entity.Member) (entity.Member, error)
	Delete(id int) error
	FindAll() ([]entity.Member, error)
	FindByID(id int) (entity.Member, error)
	Update(payload entity.Member) (entity.Member, error)
}

func NewMemberUsecase(repo repository.MemberRepository) MemberUsecase {
	return &memberUsecase{repo: repo}
}
