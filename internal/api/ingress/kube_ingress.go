package ingress

import (
	"encoding/json"
	"fmt"

	"wayne/internal/api/base"
	"wayne/internal/k8s/kind/ingress"
	"wayne/internal/model"

	networkingv1 "k8s.io/api/networking/v1"
)

type KubeIngressController struct {
	base.APIController
}

func (c *KubeIngressController) Prepare() {
	c.APIController.Prepare()

	methodActionMap := map[string]string{
		"Create": model.PermissionCreate,
	}
	_, method := c.GetControllerAndAction()
	c.PreparePermission(methodActionMap, method, model.PermissionTypeKubeIngress)
}

// swagger:route POST /api/v1/kubernetes/apps/{appid}/ingresses/{ingressId}/tpls/{tplId}/clusters/{cluster} ingress reqCreateKubeIngress
// deploy tpl
// responses:
//
//	200: respSuccessDescription
func (c *KubeIngressController) Create() {
	ingressId := c.GetIntParamFromURL(":ingressId")
	tplId := c.GetIntParamFromURL(":tplId")

	appId := c.GetIntParamFromURL(":appid")
	app, err := model.AppModel.GetById(appId)
	if err != nil {
		c.HandleError(err)
		return
	}

	if app == nil || app.Id == 0 {
		c.HandleError(fmt.Errorf("appId is empty:%d", appId))
		return
	}

	var kubeIngress networkingv1.Ingress
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &kubeIngress)
	if err != nil {
		c.HandleError(err)
		return
	}

	clusterName := c.Ctx.Input.Param(":cluster")
	k8sClient := c.Client(clusterName)

	namespaceModel, err := model.NamespaceModel.GetNamespaceByAppId(c.AppId)
	if err != nil {
		c.HandleError(err)
		return
	}

	clusterModel, err := model.ClusterModel.GetParsedMetaDataByName(clusterName)
	if err != nil {
		c.HandleError(err)
		return
	}

	// add ingress predeploy
	ingressPreDeploy(&kubeIngress, clusterModel, namespaceModel)

	// ingressDetail include endpoints
	_, err = ingress.CreateOrUpdateIngress(k8sClient, &kubeIngress)
	if err != nil {
		c.HandleError(err)
		return
	}

	publishStatus := model.PublishStatus{
		ResourceId:  int64(ingressId),
		TemplateId:  int64(tplId),
		Type:        model.PublishTypeIngress,
		Cluster:     clusterName,
		AppId:       app.Id,
		NamespaceId: app.Namespace.Id,
	}

	err = model.PublishStatusModel.Publish(&publishStatus)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success("ok")
}

func ingressPreDeploy(kubeIngress *networkingv1.Ingress, cluster *model.Cluster, namespace *model.Namespace) {
	preDefinedAnnotationMap := make(map[string]string)

	annotationResult := make(map[string]string, 0)

	for k, v := range kubeIngress.Annotations {
		preDefinedAnnotationMap[k] = v
	}

	for k, v := range cluster.MetaDataObj.IngressAnnotations {
		preDefinedAnnotationMap[k] = v
	}

	for k, v := range namespace.MetaDataObj.IngressAnnotations {
		preDefinedAnnotationMap[k] = v
	}
	for k, v := range preDefinedAnnotationMap {
		annotationResult[k] = v
	}

	kubeIngress.Annotations = annotationResult
}
