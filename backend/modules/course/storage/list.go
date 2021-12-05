package coursestorage

import (
	"btcn03-api/common"
	coursemodel "btcn03-api/modules/course/model"
	"context"
)

func (s *sqlStore) ListDataWithCondition(
	ctx context.Context,
	paging *common.Paging,
	moreKeys ...string,
) ([]coursemodel.Course, error) {
	db := s.db
	var result []coursemodel.Course

	if err := db.Table(coursemodel.Course{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	offset := (paging.Page - 1) * paging.Limit

	if err := db.
		Limit(paging.Limit).
		Offset(offset).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
