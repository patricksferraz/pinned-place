package service

import (
	"context"

	"github.com/c-4u/pinned-place/domain/entity"
	"github.com/c-4u/pinned-place/domain/repo"
	"github.com/c-4u/pinned-place/infra/client/kafka/topic"
	"github.com/c-4u/pinned-place/utils"
)

type Service struct {
	Repo repo.RepoInterface
}

func NewService(repo repo.RepoInterface) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) CreatePlace(ctx context.Context, name *string) (*string, error) {
	place, err := entity.NewPlace(name)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreatePlace(ctx, place); err != nil {
		return nil, err
	}

	// TODO: adds retry
	event, err := entity.NewEvent(place)
	if err != nil {
		return nil, err
	}

	eMsg, err := event.ToJson()
	if err != nil {
		return nil, err
	}

	err = s.Repo.PublishEvent(ctx, utils.PString(topic.NEW_PLACE), utils.PString(string(eMsg)), place.ID)
	if err != nil {
		return nil, err
	}

	return place.ID, nil
}

func (s *Service) FindPlace(ctx context.Context, placeID *string) (*entity.Place, error) {
	place, err := s.Repo.FindPlace(ctx, placeID)
	if err != nil {
		return nil, err
	}

	return place, nil
}

func (s *Service) SearchPlaces(ctx context.Context, pageToken *string, pageSize *int) ([]*entity.Place, *string, error) {
	pagination, err := entity.NewPagination(pageToken, pageSize)
	if err != nil {
		return nil, nil, err
	}

	searchPlaces, err := entity.NewSearchPlaces(pagination)
	if err != nil {
		return nil, nil, err
	}

	places, nextPageToken, err := s.Repo.SearchPlaces(ctx, searchPlaces)
	if err != nil {
		return nil, nil, err
	}

	return places, nextPageToken, nil
}
