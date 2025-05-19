package family

import (
	entityfamily "cashback_info/internal/model/family"
	familyentity "cashback_info/internal/model/family"
	userentity "cashback_info/internal/model/user"
	familyrepository "cashback_info/internal/repository/family"
	userrepository "cashback_info/internal/repository/family/user"
	repofamily "cashback_info/internal/repository/model/family"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FamilyService interface {
	CreateFamily(title string, userID uuid.UUID) error
	GetFamilyByID(id uuid.UUID) (*familyentity.Family, error)
	DeleteFamily(ID uuid.UUID) error
	GetFamilyByUserID(userID uuid.UUID) (*entityfamily.Family, error)
}

type familyService struct {
	familyRepo           familyrepository.FamilyRepository
	familyUserRepository userrepository.FamilyUserRepository
}

func NewFamilyService(familyRepo familyrepository.FamilyRepository, familyUserRepository userrepository.FamilyUserRepository) FamilyService {
	return &familyService{familyRepo: familyRepo, familyUserRepository: familyUserRepository}
}

func (f *familyService) CreateFamily(title string, userID uuid.UUID) error {
	_, err := f.familyRepo.Create(repofamily.FamilyDB{
		Title:    title,
		LeaderID: userID,
	})

	return err
}

func (f *familyService) GetFamilyByID(id uuid.UUID) (*familyentity.Family, error) {
	family, err := f.familyRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	leaderRoleType := userentity.GenerateRoleTypeFromString(string(family.Leader.RoleType))
	if leaderRoleType == nil {
		return nil, fmt.Errorf("invalid role type: %s", family.Leader.RoleType)
	}

	result := familyentity.Family{
		ID:    family.ID,
		Title: family.Title,
		Leader: userentity.User{
			ID:       family.Leader.ID,
			Login:    family.Leader.Login,
			Email:    family.Leader.Email,
			Phone:    family.Leader.Phone,
			RoleType: *leaderRoleType,
		},
	}

	for _, member := range family.Members {
		memberRoleType := userentity.GenerateRoleTypeFromString(string(member.RoleType))
		if memberRoleType == nil {
			return nil, fmt.Errorf("invalid role type: %s", member.RoleType)
		}

		result.Members = append(result.Members, userentity.User{
			ID:       member.ID,
			Login:    member.Login,
			Email:    member.Email,
			Phone:    member.Phone,
			RoleType: *memberRoleType,
		})
	}

	return &result, nil
}

func (f *familyService) DeleteFamily(ID uuid.UUID) error {
	return f.familyRepo.Delete(ID)
}

func (f *familyService) GetFamilyByUserID(userID uuid.UUID) (*entityfamily.Family, error) {
	family, err := f.familyRepo.GetByLeaderID(userID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	familyUser, err := f.familyUserRepository.GetByUserID(userID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	if family != nil && family.Leader.ID == userID {
		members := make([]userentity.User, len(family.Members))
		for i, member := range family.Members {
			memberRole := userentity.GenerateRoleTypeFromString(string(member.RoleType))
			if memberRole == nil {
				return nil, fmt.Errorf("invalid role type: %s", member.RoleType)
			}

			members[i] = userentity.User{
				ID:       member.ID,
				Login:    member.Login,
				Email:    member.Email,
				Phone:    member.Phone,
				RoleType: *memberRole,
			}
		}

		leaderRoleType := userentity.GenerateRoleTypeFromString(string(family.Leader.RoleType))
		if leaderRoleType == nil {
			return nil, fmt.Errorf("invalid role type: %s", family.Leader.RoleType)
		}

		return &entityfamily.Family{
			ID:    family.ID,
			Title: family.Title,
			Leader: userentity.User{
				ID:       family.Leader.ID,
				Login:    family.Leader.Login,
				Email:    family.Leader.Email,
				Phone:    family.Leader.Phone,
				RoleType: *userentity.GenerateRoleTypeFromString(string(family.Leader.RoleType)),
			},
			Members: members,
		}, nil
	} else if familyUser != nil {
		members := make([]userentity.User, len(familyUser.Family.Members))

		for i, member := range familyUser.Family.Members {
			memberRole := userentity.GenerateRoleTypeFromString(string(member.RoleType))
			if memberRole == nil {
				return nil, fmt.Errorf("invalid role type: %s", member.RoleType)
			}

			members[i] = userentity.User{
				ID:       member.ID,
				Login:    member.Login,
				Email:    member.Email,
				Phone:    member.Phone,
				RoleType: *memberRole,
			}
		}

		leaderRoleType := userentity.GenerateRoleTypeFromString(string(familyUser.Family.Leader.RoleType))
		if leaderRoleType == nil {
			return nil, fmt.Errorf("invalid role type: %s", familyUser.Family.Leader.RoleType)
		}

		return &entityfamily.Family{
			ID:    family.ID,
			Title: family.Title,
			Leader: userentity.User{
				ID:       familyUser.Family.Leader.ID,
				Login:    familyUser.Family.Leader.Login,
				Email:    familyUser.Family.Leader.Email,
				Phone:    familyUser.Family.Leader.Phone,
				RoleType: *leaderRoleType,
			},
			Members: members,
		}, nil
	} else {
		return nil, nil
	}
}
