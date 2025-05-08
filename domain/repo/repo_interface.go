package repo

import (
	"context"

	"github.com/patricksferraz/pinned-place/domain/entity"
)

type RepoInterface interface {
	CreatePlace(ctx context.Context, place *entity.Place) error
	FindPlace(ctx context.Context, placeID *string) (*entity.Place, error)
	SavePlace(ctx context.Context, place *entity.Place) error
	SearchPlaces(ctx context.Context, searchPlace *entity.SearchPlaces) ([]*entity.Place, *string, error)

	PublishEvent(ctx context.Context, topic, msg, key *string) error
}
