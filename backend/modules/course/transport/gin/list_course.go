package coursegin

import (
	"btcn03-api/common"
	"btcn03-api/component/appctx"
	coursebiz "btcn03-api/modules/course/biz"
	coursestorage "btcn03-api/modules/course/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCourse(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := paging.Process(); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := coursestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := coursebiz.NewListCourseBiz(store)

		result, err := biz.ListCourse(c.Request.Context(), &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
