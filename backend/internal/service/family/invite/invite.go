package invite

import (
	"cashback_info/internal/model/family"
	"cashback_info/internal/model/family/invite"
	"cashback_info/internal/model/user"
	repository "cashback_info/internal/repository/family/invite"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type FamilyInviteService interface {
	ListByFamilyID(
		ctx context.Context,
		familyID uuid.UUID,
	) ([]invite.FamilyInvite, error)
}

type familyInviteService struct {
	repo repository.FamilyInviteRepository
}

func NewFamilyInvitesService(repo repository.FamilyInviteRepository) *familyInviteService {
	return &familyInviteService{repo: repo}
}

func (s *familyInviteService) ListByFamilyID(
	ctx context.Context,
	familyID uuid.UUID,
) ([]invite.FamilyInvite, error) {
	familyInvites, err := s.repo.ListByFamilyID(familyID)

	result := make([]invite.FamilyInvite, len(familyInvites))
	for i, familyInvite := range familyInvites {
		leaderRoleType := user.GenerateRoleTypeFromString(string(familyInvite.Family.Leader.RoleType))
		if leaderRoleType == nil {
			return nil, fmt.Errorf("invalid role type: %s", familyInvite.Family.Leader.RoleType)
		}

		result[i] = invite.FamilyInvite{
			ID: familyInvite.ID,
			Family: family.Family{
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

func (s *familyInviteService) ListByUserID(
	ctx context.Context,
	familyID uuid.UUID,
) ([]invite.FamilyInvite, error) {
	familyInvites, err := s.repo.ListByUserID(familyID)

	result := make([]invite.FamilyInvite, len(familyInvites))
	for i, familyInvite := range familyInvites {
		leaderRoleType := user.GenerateRoleTypeFromString(string(familyInvite.Family.Leader.RoleType))
		if leaderRoleType == nil {
			return nil, fmt.Errorf("invalid role type: %s", familyInvite.Family.Leader.RoleType)
		}

		result[i] = invite.FamilyInvite{
			ID: familyInvite.ID,
			Family: family.Family{
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
