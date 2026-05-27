package ingress

import (
	"encoding/json"

	"wayne/internal/api/base"
	"wayne/internal/model"
)

type IngressController struct {
	base.APIController
}

func (c *IngressController) Prepare() {

	c.APIController.Prepare()

	perAction := ""
	_, method := c.GetControllerAndAction()
	switch method {
	case "Get", "List":
		perAction = model.PermissionRead
	case "Create":
		perAction = model.PermissionCreate
	case "Update":
		perAction = model.PermissionUpdate
	case "Delete":
		perAction = model.PermissionDelete
	}
	if perAction != "" {
		c.CheckPermission(model.PermissionTypeIngress, perAction)
	}
}

// swagger:route GET /api/v1/apps/{appid}/ingresses/names ingress reqGetNamesIngress
// get all id and names
// responses:
//
//	200: respSuccessDescription
func (c *IngressController) GetNames() {
	filters := make(map[string]interface{})
	deleted := c.GetDeleteFromQuery()

	filters["Deleted"] = deleted
	if c.AppId != 0 {
		filters["App__Id"] = c.AppId
	}

	ingresses, err := model.IngressModel.GetNames(filters)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(ingresses)
}

// swagger:route GET /api/v1/apps/{appid}/ingresses ingress reqListIngress
// get all Ingress
// responses:
//
//	200: respSuccessDescription
func (c *IngressController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}

	ingrs := []model.Ingress{}
	if c.AppId != 0 {
		param.Query["App__Id"] = c.AppId
	} else if !c.User.Admin {
		param.Query["App__AppUsers__User__Id__exact"] = c.User.Id
		perName := model.PermissionModel.MergeName(model.PermissionTypeIngress, model.PermissionRead)
		param.Query["App__AppUsers__Group__Permissions__Permission__Name__contains"] = perName
		param.Groupby = []string{"Id"}
	}

	total, err := model.GetTotal(new(model.Ingress), param)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.GetAll(new(model.Ingress), &ingrs, param)
	if err != nil {
		c.HandleError(err)
		return
	}
	for key, one := range ingrs {
		ingrs[key].AppId = one.App.Id
	}

	c.Success(param.NewPage(total, ingrs))
}

// swagger:route POST /api/v1/apps/{appid}/ingresses ingress reqCreateIngress
// create Ingress
// responses:
//
//	200: respSuccessDescription
func (c *IngressController) Create() {
	var ingr model.Ingress
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ingr)
	if err != nil {
		c.HandleError(err)
		return
	}

	ingr.User = c.User.Name
	_, err = model.IngressModel.Add(&ingr)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(ingr)
}

// swagger:route GET /api/v1/apps/{appid}/ingresses/{id} ingress reqGetIngress
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *IngressController) Get() {
	id := c.GetIDFromURL()

	ingr, err := model.IngressModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(ingr)
	return
}

// swagger:route PUT /api/v1/apps/{appid}/ingresses/{id} ingress reqUpdateIngress
// update the Ingress
// responses:
//
//	200: respSuccessDescription
func (c *IngressController) Update() {
	id := c.GetIDFromURL()
	var ingr model.Ingress
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ingr)
	if err != nil {
		c.HandleError(err)
		return
	}

	ingr.Id = int64(id)
	err = model.IngressModel.UpdateById(&ingr)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(ingr)
}

// swagger:route PUT /api/v1/apps/{appid}/ingresses/updateorders ingress reqUpdateOrdersIngress
// batch update the Orders
// responses:
//
//	200: respSuccessDescription
func (c *IngressController) UpdateOrders() {
	var ingr []*model.Ingress
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ingr)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.IngressModel.UpdateOrders(ingr)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success("ok!")
}

// swagger:route DELETE /api/v1/apps/{appid}/ingresses/{id} ingress reqDeleteIngress
// delete the Ingress
// responses:
//
//	200: respSuccessDescription
func (c *IngressController) Delete() {
	id := c.GetIDFromURL()

	logical := c.GetLogicalFromQuery()

	err := model.IngressModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
