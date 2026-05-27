package pv

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/beego/beego/v2/adapter/httplib"
	v1 "k8s.io/api/core/v1"

	"wayne/internal/api/base"
	"wayne/internal/model"
	"wayne/pkg/hack"
)

type RobinPersistentVolumeController struct {
	base.APIController
}

func (c *RobinPersistentVolumeController) Prepare() {

	c.APIController.Prepare()

	methodActionMap := map[string]string{
		"ListRbdImages":  model.PermissionRead,
		"CreateRbdImage": model.PermissionCreate,
	}
	_, method := c.GetControllerAndAction()
	c.PreparePermission(methodActionMap, method, model.PermissionTypeKubePersistentVolume)

}

// swagger:route GET /api/v1/kubernetes/persistentvolumes/robin/rbd.images/clusters/{cluster} pv reqListRbdImagesRobinPersistentVolume
// find rbd images by cluster
// responses:
//
//	200: respSuccessDescription
func (c *RobinPersistentVolumeController) ListRbdImages() {
	cluster := c.Ctx.Input.Param(":cluster")
	robinMetaData, err := clusterRobinMetaData(cluster)
	if err != nil {
		c.HandleError(err)
		return
	}

	result, err := doRobinRequest(
		httplib.Get(fmt.Sprintf(
			"%s/v1/images",
			robinMetaData.Url)).Param("type", "all"), robinMetaData.Token)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)

}

func clusterRobinMetaData(cluster string) (*model.ClusterRobinMetaData, error) {
	metaData, err := model.ClusterModel.ClusterMetaData(cluster)
	if err != nil {
		return nil, err
	}
	if metaData.Robin == nil {
		return nil, errors.New("No Robin metaData configured! ")
	}

	if metaData.Robin.Url == "" {
		return nil, errors.New("Robin url is null. ")
	}

	if metaData.Robin.Token == "" {
		return nil, errors.New("Robin token is null. ")
	}

	return metaData.Robin, nil
}

// swagger:route POST /api/v1/kubernetes/persistentvolumes/robin/rbd.images/clusters/{cluster} pv reqCreateRbdImageRobinPersistentVolume
// create rbd images
// responses:
//
//	200: respSuccessDescription
func (c *RobinPersistentVolumeController) CreateRbdImage() {
	var pvTpl v1.PersistentVolume
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &pvTpl)
	if err != nil {
		c.HandleError(err)
		return
	}

	storage := pvTpl.Spec.Capacity[v1.ResourceStorage]

	rbdImage := struct {
		Name string `json:"name"`
		Size int64  `json:"size"` // 单位，MB
	}{

		Size: storage.Value() / (1024 * 1024),
	}

	imageType := ""
	if pvTpl.Spec.RBD != nil {
		imageType = "rbd"
		rbdImage.Name = pvTpl.Spec.RBD.RBDImage
	} else if pvTpl.Spec.CephFS != nil {
		imageType = "cephfs"
		paths := strings.Split(pvTpl.Spec.CephFS.Path, "/")
		rbdImage.Name = paths[len(paths)-1]
	}

	if imageType == "" {
		c.HandleError(fmt.Errorf("imageType is empty"))
		return
	}

	cluster := c.Ctx.Input.Param(":cluster")

	robinMetaData, err := clusterRobinMetaData(cluster)
	if err != nil {
		c.HandleError(err)
		return
	}

	body, _ := json.Marshal(&rbdImage)

	result, err := doRobinRequest(
		httplib.Put(fmt.Sprintf(
			"%s/v1/image?type=%s",
			robinMetaData.Url,
			imageType)).Body(body), robinMetaData.Token)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(result)

}

func doRobinRequest(request *httplib.BeegoHTTPRequest, token string) (interface{}, error) {
	resp, err := request.Header("token", token).
		DoRequest()
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(hack.String(result))
	}

	var obj interface{}

	err = json.Unmarshal(result, &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
