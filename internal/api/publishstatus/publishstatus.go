package publishstatus

import (
	"wayne/internal/api/base"
	"wayne/internal/model"
)

type PublishStatusController struct {
	base.APIController
}

func (c *PublishStatusController) Prepare() {

	c.APIController.Prepare()
}

// swagger:route GET /api/v1/publishstatus publishstatus reqListPublishStatus
// get all PublishStatus
// responses:
//
//	200: respSuccessDescription
func (c *PublishStatusController) List() {
	resourceType := c.GetIntParamFromQuery("type")
	resourceId := c.GetIntParamFromQuery("resourceId")

	status, err := model.PublishStatusModel.GetAll(model.PublishType(resourceType), int64(resourceId))
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(status)
}

// swagger:route DELETE /api/v1/publishstatus publishstatus/{id} reqDeletePublishStatus
// delete the publishstatus
// responses:
//
//	200: respSuccessDescription
func (c *PublishStatusController) Delete() {
	id := c.GetIDFromURL()

	err := model.PublishStatusModel.DeleteById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
