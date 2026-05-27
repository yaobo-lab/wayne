package appstarred

import (
	"encoding/json"

	"wayne/internal/api/base"
	"wayne/internal/model"
)

type AppStarredController struct {
	base.APIController
}

func (c *AppStarredController) Prepare() {
	c.APIController.Prepare()
}

// swagger:route POST /api/v1/apps/stars app reqCreateAppStarred
// create AppStarred
// responses:
//
//	200: respSuccessDescription
//	403: respFailureDescription
func (c *AppStarredController) Create() {
	var appStarred model.AppStarred
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &appStarred)
	if err != nil {
		c.HandleError(err)
		return
	}
	appStarred.User = c.User

	objectid, err := model.AppStarredModel.Add(&appStarred)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(objectid)
}

// swagger:route DELETE /api/v1/apps/stars/{appId} app reqDeleteAppStarred
// delete the AppStarred
// responses:
//
//	200: respSuccessDescription
//	403: respFailureDescription
func (c *AppStarredController) Delete() {
	appId := c.GetIntParamFromURL(":appId")

	err := model.AppStarredModel.DeleteByAppId(c.User.Id, int64(appId))
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
