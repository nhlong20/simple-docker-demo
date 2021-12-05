package coursebiz

import (
	"btcn03-api/common"
	coursemodel "btcn03-api/modules/course/model"
	"context"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *coursemodel.Course) error
}

type createCourseBiz struct {
	store CreateRestaurantStore
}

func NewCreateCourseBiz(store CreateRestaurantStore) *createCourseBiz {
	return &createCourseBiz{store: store}
}

func (biz *createCourseBiz) CreateNewCourse(ctx context.Context, data *coursemodel.Course) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(coursemodel.EntityName, err)
	}

	return nil
}
