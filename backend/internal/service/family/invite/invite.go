package invite

import (
	entityfamily "cashback_info/internal/model/family"
	entityinvite "cashback_info/internal/model/family/invite"
	"cashback_info/internal/model/user"
	inviterepository "cashback_info/internal/repository/family/invite"
	repoinvite "cashback_info/internal/repository/model/family/invite"
	"fmt"

	"github.com/google/uuid"
)

type FamilyInviteService interface {
	Create(familyID, userID uuid.UUID) error
	ListByFamilyID(familyID uuid.UUID) ([]entityinvite.FamilyInvite, error)
	ListByUserID(userID uuid.UUID) ([]entityinvite.FamilyInvite, error)
	DeleteByID(ID uuid.UUID) error
	DeleteByUserID(userID uuid.UUID) error
}

type familyInviteService struct {
	familyInviteRepo inviterepository.FamilyInviteRepository
}

func NewFamilyInviteService(repo inviterepository.FamilyInviteRepository) FamilyInviteService {
	return &familyInviteService{familyInviteRepo: repo}
}

func (s *familyInviteService) Create(familyID, userID uuid.UUID) error {
	return s.familyInviteRepo.Create(repoinvite.FamilyInviteDB{FamilyID: familyID, UserID: userID})
}

func (s *familyInviteService) ListByFamilyID(
	familyID uuid.UUID,
) ([]entityinvite.FamilyInvite, error) {
	familyInvites, err := s.familyInviteRepo.ListByFamilyID(familyID)

	result := make([]entityinvite.FamilyInvite, len(familyInvites))
	for i, familyInvite := range familyInvites {
		leaderRoleType := user.GenerateRoleTypeFromString(string(familyInvite.Family.Leader.RoleType))
		if leaderRoleType == nil {
			return nil, fmt.Errorf("invalid role type: %s", familyInvite.Family.Leader.RoleType)
		}

		result[i] = entityinvite.FamilyInvite{
			ID: familyInvite.ID,
			Family: entityfamily.Family{
				ID:    familyInvite.Family.ID,
				Title: familyInvite.Family.Title,
				Leader: user.User{
					ID:       familyInvite.Family.Leader.ID,
					Login:    familyInvite.Family.Leader.Login,
					Email:    familyInvite.Family.Leader.Email,
					Phone:    familyInvite.Family.Leader.Phone,
					RoleType: *leaderRoleType,
				},
			},
			User: user.User{
				ID:       familyInvite.User.ID,
				Login:    familyInvite.User.Login,
				Email:    familyInvite.User.Email,
				Phone:    familyInvite.User.Phone,
				RoleType: familyInvite.User.RoleType,
			},
		}
	}
	return result, err
}

func (s *familyInviteService) ListByUserID(userID uuid.UUID) ([]entityinvite.FamilyInvite, error) {
	familyInvites, err := s.familyInviteRepo.ListByUserID(userID)

	result := make([]entityinvite.FamilyInvite, len(familyInvites))
	for i, familyInvite := range familyInvites {
		leaderRoleType := user.GenerateRoleTypeFromString(string(familyInvite.Family.Leader.RoleType))
		if leaderRoleType == nil {
			return nil, fmt.Errorf("invalid role type: %s", familyInvite.Family.Leader.RoleType)
		}

		result[i] = entityinvite.FamilyInvite{
			ID: familyInvite.ID,
			Family: entityfamily.Family{
				ID:    familyInvite.Family.ID,
				Title: familyInvite.Family.Title,
				Leader: user.User{
					ID:       familyInvite.Family.Leader.ID,
					Login:    familyInvite.Family.Leader.Login,
					Email:    familyInvite.Family.Leader.Email,
					Phone:    familyInvite.Family.Leader.Phone,
					RoleType: *leaderRoleType,
				},
			},
			User: user.User{
				ID:       familyInvite.User.ID,
				Login:    familyInvite.User.Login,
				Email:    familyInvite.User.Email,
				Phone:    familyInvite.User.Phone,
				RoleType: familyInvite.User.RoleType,
			},
		}
	}
	return result, err
}

func (s *familyInviteService) DeleteByID(ID uuid.UUID) error {
	return s.familyInviteRepo.DeleteByID(ID)
}

func (s *familyInviteService) DeleteByUserID(userID uuid.UUID) error {
	return s.familyInviteRepo.DeleteByUserID(userID)
}
