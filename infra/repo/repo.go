package repo

import (
	"context"
	"fmt"

	"github.com/c-4u/pinned-place/domain/entity"
	"github.com/c-4u/pinned-place/infra/client/kafka"
	"github.com/c-4u/pinned-place/infra/db"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Repository struct {
	Orm *db.DbOrm
	Kp  *kafka.KafkaProducer
}

func NewRepository(orm *db.DbOrm, kp *kafka.KafkaProducer) *Repository {
	return &Repository{
		Orm: orm,
		Kp:  kp,
	}
}

func (r *Repository) CreatePlace(ctx context.Context, place *entity.Place) error {
	err := r.Orm.Db.Create(place).Error
	return err
}

func (r *Repository) FindPlace(ctx context.Context, placeID *string) (*entity.Place, error) {
	var e entity.Place

	r.Orm.Db.First(&e, "id = ?", *placeID)

	if e.ID == nil {
		return nil, fmt.Errorf("no place found")
	}

	return &e, nil
}

func (r *Repository) SavePlace(ctx context.Context, place *entity.Place) error {
	err := r.Orm.Db.Save(place).Error
	return err
}

func (r *Repository) SearchPlaces(ctx context.Context, searchPlace *entity.SearchPlaces) ([]*entity.Place, *string, error) {
	var e []*entity.Place

	q := r.Orm.Db
	if *searchPlace.PageToken != "" {
		q = q.Where("token < ?", *searchPlace.PageToken)
	}
	err := q.Order("token DESC").
		Limit(*searchPlace.PageSize).
		Find(&e).Error
	if err != nil {
		return nil, nil, err
	}

	if len(e) < *searchPlace.PageSize {
		return e, nil, nil
	}

	return e, e[len(e)-1].Token, nil
}

func (r *Repository) PublishEvent(ctx context.Context, topic, msg, key *string) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: topic, Partition: ckafka.PartitionAny},
		Value:          []byte(*msg),
		Key:            []byte(*key),
	}
	err := r.Kp.Producer.Produce(message, r.Kp.DeliveryChan)
	if err != nil {
		return err
	}
	return nil
}
