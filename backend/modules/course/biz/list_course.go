package coursebiz

import (
	"btcn03-api/common"
	coursemodel "btcn03-api/modules/course/model"
	"context"
)

type ListCourseStore interface {
	ListDataWithCondition(
		ctx context.Context,
		paging *common.Paging,
		moreKeys ...string,
	) ([]coursemodel.Course, error)
}

type listCourseBiz struct {
	store ListCourseStore
}

func NewListCourseBiz(store ListCourseStore) *listCourseBiz {
	return &listCourseBiz{store: store}
}

func (biz *listCourseBiz) ListCourse(ctx context.Context,
	paging *common.Paging,
) ([]coursemodel.Course, error) {
	result, err := biz.store.ListDataWithCondition(ctx, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(coursemodel.EntityName, err)
	}

	return result, nil
}
