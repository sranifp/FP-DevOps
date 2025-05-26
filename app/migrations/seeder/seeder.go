package seeder

import (
	"FP-DevOps/migrations/seeder/seeders"

	"gorm.io/gorm"
)

func RunSeeders(db *gorm.DB) error {
	if err := seeders.UserSeeder(db); err != nil {
		return err
	}
	return nil
}
