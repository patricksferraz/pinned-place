package db

import (
	"fmt"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/gorm"
)

type MigrateOrm struct {
	Db *gorm.DB
	m  *gormigrate.Gormigrate
}

func NewMigrateOrm(db *gorm.DB) *MigrateOrm {
	m := MigrateOrm{
		Db: db,
	}
	m.load()
	return &m
}

func (m *MigrateOrm) load() {
	m.m = gormigrate.New(m.Db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202203301940",
			Migrate: func(db *gorm.DB) error {
				type Base struct {
					ID        string    `gorm:"type:uuid;primaryKey"`
					CreatedAt time.Time `gorm:"column:created_at;autoUpdateTime"`
					UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime"`
				}
				type Place struct {
					Base
					Name  *string `gorm:"column:name;not null"`
					Token *string `gorm:"column:token;type:varchar(25);not null"`
				}

				return db.AutoMigrate(&Place{})
			},
			Rollback: func(db *gorm.DB) error {
				return db.Migrator().DropTable("places")
			},
		},
	})
}

func (m *MigrateOrm) Migrate() error {
	if err := m.m.Migrate(); err != nil {
		return fmt.Errorf("could not migrate: %v", err)
	}
	return nil
}
