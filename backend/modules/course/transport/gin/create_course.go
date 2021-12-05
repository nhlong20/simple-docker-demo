package coursegin

import (
	"btcn03-api/common"
	"btcn03-api/component/appctx"
	coursebiz "btcn03-api/modules/course/biz"
	coursemodel "btcn03-api/modules/course/model"
	coursestorage "btcn03-api/modules/course/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCourse(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var newData coursemodel.Course

		if err := c.ShouldBind(&newData); err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}
		store := coursestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := coursebiz.NewCreateCourseBiz(store)

		if err := biz.CreateNewCourse(c.Request.Context(), &newData); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(newData.Id))
	}
}
