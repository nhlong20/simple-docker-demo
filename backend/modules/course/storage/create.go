package coursestorage

import (
	"btcn03-api/common"
	coursemodel "btcn03-api/modules/course/model"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *coursemodel.Course) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
