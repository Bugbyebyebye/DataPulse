package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"store-service/dao"
	"strconv"
)

func (*StoreHandle) DeleteWarehouseTable(ctx *gin.Context) {
	idStr := ctx.Request.Header.Get("id")
	id, _ := strconv.Atoi(idStr)
	tableIdStr := ctx.PostForm("table_id")
	tableId, _ := strconv.Atoi(tableIdStr)

	err := dao.DeleteWarehouseTable(id, tableId)
	if err != nil {
		ctx.JSON(http.StatusOK, res.Fail(4001, "删除失败"))
		return
	}

	ctx.JSON(http.StatusOK, res.Success("删除成功"))
}
