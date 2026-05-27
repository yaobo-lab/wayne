package ingress

import (
	"encoding/json"
	"fmt"

	"wayne/internal/api/base"
	"wayne/internal/model"
	"wayne/pkg/hack"

	networkingv1 "k8s.io/api/networking/v1"
)

type IngressTplController struct {
	base.APIController
}

func (c *IngressTplController) Prepare() {

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
		c.CheckPermission(model.PermissionTypeStatefulset, perAction)
	}
}

// swagger:route GET /api/v1/apps/{appid}/ingresses/tpls ingress reqListIngressTpl
// get all ingressTpl
// responses:
//
//	200: respSuccessDescription
func (c *IngressTplController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}

	isOnline := c.GetIsOnlineFromQuery()

	ingressId := c.Input().Get("ingressId")
	if ingressId != "" {
		param.Query["ingress_id"] = ingressId
	}

	var ingrsTpls []model.IngressTemplate
	total, err := model.ListTemplate(&ingrsTpls, param, model.TableNameIngressTemplate, model.PublishTypeIngress, isOnline)
	if err != nil {
		c.HandleError(err)
		return
	}
	for index, tpl := range ingrsTpls {
		ingrsTpls[index].IngressId = tpl.Ingress.Id
	}

	c.Success(param.NewPage(total, ingrsTpls))
}

// swagger:route POST /api/v1/apps/{appid}/ingresses/tpls ingress reqCreateIngressTpl
// create ingressTpl
// responses:
//
//	200: respSuccessDescription
func (c *IngressTplController) Create() {
	var ingrTpl model.IngressTemplate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ingrTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	err = validIngressTemplate(ingrTpl.Template)
	if err != nil {
		c.HandleError(err)
		return
	}

	ingrTpl.User = c.User.Name

	_, err = model.IngressTemplateModel.Add(&ingrTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(ingrTpl)
}

func validIngressTemplate(ingrTplStr string) error {
	ingr := networkingv1.Ingress{}
	err := json.Unmarshal(hack.Slice(ingrTplStr), &ingr)
	if err != nil {
		return fmt.Errorf("ingress template format error.%v", err.Error())
	}
	return nil
}

// swagger:route GET /api/v1/apps/{appid}/ingresses/tpls/{id} ingress reqGetIngressTpl
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *IngressTplController) Get() {
	id := c.GetIDFromURL()

	ingrTpl, err := model.IngressTemplateModel.GetById(id)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(ingrTpl)
}

// swagger:route PUT /api/v1/apps/{appid}/ingresses/tpls/{id} ingress reqUpdateIngressTpl
// update the IngressTpl
// responses:
//
//	200: respSuccessDescription
func (c *IngressTplController) Update() {
	id := c.GetIDFromURL()
	var ingrTpl model.IngressTemplate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ingrTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	if err = validIngressTemplate(ingrTpl.Template); err != nil {
		c.HandleError(err)
		return
	}

	ingrTpl.Id = int64(id)
	err = model.IngressTemplateModel.UpdateById(&ingrTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(ingrTpl)
}

// swagger:route DELETE /api/v1/apps/{appid}/ingresses/tpls/{id} ingress reqDeleteIngressTpl
// delete the ingressTpl
// responses:
//
//	200: respSuccessDescription
func (c *IngressTplController) Delete() {
	id := c.GetIDFromURL()
	logical := c.GetLogicalFromQuery()

	err := model.IngressTemplateModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
