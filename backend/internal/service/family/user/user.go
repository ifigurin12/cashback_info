package user

import (
	familyuserrepository "cashback_info/internal/repository/family/user"
	familyuser "cashback_info/internal/repository/model/family/user"

	"github.com/google/uuid"
)

type FamilyUserService interface {
	Create(familyID, userID uuid.UUID) error
	Delete(familyID, userID uuid.UUID) error
}

type familyUserService struct {
	familyUserRepo familyuserrepository.FamilyUserRepository
}

func NewFamilyUserService(repo familyuserrepository.FamilyUserRepository) FamilyUserService {
	return &familyUserService{familyUserRepo: repo}
}

func (s *familyUserService) Create(familyID, userID uuid.UUID) error {
	return s.familyUserRepo.Create(familyuser.FamilyUserDB{FamilyID: familyID, UserID: userID})
}

func (s *familyUserService) Delete(familyID, userID uuid.UUID) error {
	return s.familyUserRepo.Delete(familyID, userID)
}
