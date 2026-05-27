package pvc

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/adapter/httplib"

	"wayne/internal/api/base"
	"wayne/internal/k8s/kind/pvc"
	"wayne/internal/model"
	"wayne/pkg/des"
	"wayne/pkg/hack"
)

type RobinPersistentVolumeClaimController struct {
	base.APIController
}

func (c *RobinPersistentVolumeClaimController) Prepare() {

	c.APIController.Prepare()

	methodActionMap := map[string]string{
		"GetPvcStatus":      model.PermissionRead,
		"ActiveImage":       model.PermissionRead,
		"InActiveImage":     model.PermissionRead,
		"OfflineImageUser":  model.PermissionRead,
		"LoginInfo":         model.PermissionRead,
		"Verify":            model.PermissionRead,
		"ListSnapshot":      model.PermissionRead,
		"CreateSnapshot":    model.PermissionRead,
		"DeleteAllSnapshot": model.PermissionRead,
		"DeleteSnapshot":    model.PermissionRead,
		"RollbackSnapshot":  model.PermissionRead,
	}
	_, method := c.GetControllerAndAction()
	c.PreparePermission(methodActionMap, method, model.PermissionTypeKubePersistentVolumeClaim)
}

// swagger:route GET /api/v1/kubernetes/apps/{appid}/persistentvolumeclaims/robin/{pvc}/status/namespaces/{namespace}/clusters/{cluster} pvc reqGetPvcStatusRobinPersistentVolumeClaim
// find PersistentVolumeClaim by cluster
// responses:
//
//	200: respSuccessDescription
func (c *RobinPersistentVolumeClaimController) GetPvcStatus() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":pvc")
	cli := c.Client(cluster)
	image, imageType, err := pvc.GetImageNameAndTypeByPvc(cli, name, namespace)
	if err != nil {
		c.HandleError(err)
		return
	}
	status := struct {
		Status    []string `json:"status"`
		RbdImage  string   `json:"rbdImage"`
		ImageType string   `json:"imageType"`
	}{RbdImage: image, ImageType: imageType}

	robinMetaData, err := clusterRobinMetaData(cluster)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = doRobinRequestDeserialization(httplib.Get(fmt.Sprintf("%s/v1/device/%s/status?type=%s",
		robinMetaData.Url,
		image,
		imageType)), &status, robinMetaData.Token)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(status)
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

// swagger:route POST /api/v1/kubernetes/apps/{appid}/persistentvolumeclaims/robin/{pvc}/rbd/namespaces/{namespace}/clusters/{cluster} pvc reqActiveImageRobinPersistentVolumeClaim
// active rbd images
// responses:
//
//	200: respSuccessDescription
func (c *RobinPersistentVolumeClaimController) ActiveImage() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":pvc")

	cli := c.Client(cluster)
	image, imageType, err := pvc.GetImageNameAndTypeByPvc(cli, name, namespace)
	if err != nil {
		c.HandleError(err)
		return
	}
	robinMetaData, err := clusterRobinMetaData(cluster)
	if err != nil {
		c.HandleError(err)
		return
	}

	password, err := des.DesEncrypt(hack.Slice(name), hack.Slice(robinMetaData.PasswordDesKey))
	if err != nil {
		c.HandleError(err)
		return
	}

	activeUser := struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}{
		User:     name,
		Password: base64.StdEncoding.EncodeToString(password),
	}
	body, _ := json.Marshal(&activeUser)

	result, err := doRobinRequest(
		httplib.Put(
			fmt.Sprintf(
				"%s/v1/device/%s?type=%s",
				robinMetaData.Url,
				image,
				imageType)).Body(body), robinMetaData.Token)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(result)
}

// swagger:route DELETE /api/v1/kubernetes/apps/{appid}/persistentvolumeclaims/robin/{pvc}/rbd/namespaces/{namespace}/clusters/{cluster} pvc reqInActiveImageRobinPersistentVolumeClaim
// inActive rbd images
// responses:
//
//	200: respSuccessDescription
func (c *RobinPersistentVolumeClaimController) InActiveImage() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":pvc")

	cli := c.Client(cluster)
	image, imageType, err := pvc.GetImageNameAndTypeByPvc(cli, name, namespace)
	if err != nil {
		c.HandleError(err)
		return
	}

	robinMetaData, err := clusterRobinMetaData(cluster)
	if err != nil {
		c.HandleError(err)
		return
	}

	result, err := doRobinRequest(httplib.Delete(
		fmt.Sprintf(
			"%s/v1/device/%s?type=%s",
			robinMetaData.Url,
			image,
			imageType)), robinMetaData.Token)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(result)

}

// swagger:route DELETE /api/v1/kubernetes/apps/{appid}/persistentvolumeclaims/robin/{pvc}/user/namespaces/{namespace}/clusters/{cluster} pvc reqOfflineImageUserRobinPersistentVolumeClaim
// offline image user
// responses:
//
//	200: respSuccessDescription
func (c *RobinPersistentVolumeClaimController) OfflineImageUser() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":pvc")

	cli := c.Client(cluster)
	image, imageType, err := pvc.GetImageNameAndTypeByPvc(cli, name, namespace)
	if err != nil {
		c.HandleError(err)
		return
	}
	robinMetaData, err := clusterRobinMetaData(cluster)
	if err != nil {
		c.HandleError(err)
		return
	}

	result, err := doRobinRequest(
		httplib.Delete(fmt.Sprintf(
			"%s/v1/device/%s/user?type=%s",
			robinMetaData.Url,
			image,
			imageType)), robinMetaData.Token)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(result)

}

// swagger:route GET /api/v1/kubernetes/apps/{appid}/persistentvolumeclaims/robin/{pvc}/user/namespaces/{namespace}/clusters/{cluster} pvc reqLoginInfoRobinPersistentVolumeClaim
// get user info
// responses:
//
//	200: respSuccessDescription
func (c *RobinPersistentVolumeClaimController) LoginInfo() {
	name := c.Ctx.Input.Param(":pvc")

	cluster := c.Ctx.Input.Param(":cluster")

	robinMetaData, err := clusterRobinMetaData(cluster)
	if err != nil {
		c.HandleError(err)
		return
	}

	password, err := des.DesEncrypt(hack.Slice(name), hack.Slice(robinMetaData.PasswordDesKey))
	if err != nil {
		c.HandleError(err)
		return
	}

	robinServer, err := url.Parse(robinMetaData.Url)
	if err != nil {
		c.HandleError(err)
		return
	}

	activeUser := struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Server   string `json:"server,omitempty"`
		Port     int    `json:"port,omitempty"`
	}{
		User:     name,
		Password: base64.StdEncoding.EncodeToString(password),
		Server:   robinServer.Hostname(),
		Port:     robinMetaData.SftpPort,
	}

	c.Success(activeUser)
}

// swagger:route GET /api/v1/kubernetes/apps/{appid}/persistentvolumeclaims/robin/{pvc}/verify/namespaces/{namespace}/clusters/{cluster} pvc reqVerifyRobinPersistentVolumeClaim
// verify file
// responses:
//
//	200: respSuccessDescription
func (c *RobinPersistentVolumeClaimController) Verify() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":pvc")

	cli := c.Client(cluster)
	image, imageType, err := pvc.GetImageNameAndTypeByPvc(cli, name, namespace)
	if err != nil {
		c.HandleError(err)
		return
	}
	robinMetaData, err := clusterRobinMetaData(cluster)
	if err != nil {
		c.HandleError(err)
		return
	}

	result, err := doRobinRequest(
		httplib.Get(fmt.Sprintf(
			"%s/v1/device/%s/verify?type=%s",
			robinMetaData.Url,
			image,
			imageType)), robinMetaData.Token)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(result)
}

func doRobinRequestDeserialization(request *httplib.BeegoHTTPRequest, obj interface{}, token string) error {
	resp, err := request.Header("token", token).
		DoRequest()
	if err != nil {
		return err
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(hack.String(result))
	}

	err = json.Unmarshal(result, &obj)
	if err != nil {
		return err
	}

	return nil
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

// swagger:route GET /api/v1/kubernetes/apps/{appid}/persistentvolumeclaims/robin/{pvc}/snapshot/namespaces/{namespace}/clusters/{cluster} pvc reqListSnapshotRobinPersistentVolumeClaim
// list snapshot
// responses:
//
//	200: respSuccessDescription
func (c *RobinPersistentVolumeClaimController) ListSnapshot() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":pvc")

	cli := c.Client(cluster)
	image, err := pvc.GetRbdImageByPvc(cli, name, namespace)
	if err != nil {
		c.HandleError(err)
		return
	}
	robinMetaData, err := clusterRobinMetaData(cluster)
	if err != nil {
		c.HandleError(err)
		return
	}

	result, err := doRobinRequest(
		httplib.Get(fmt.Sprintf(
			"%s/v1/snaps/%s",
			robinMetaData.Url,
			image)), robinMetaData.Token)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(result)
}

// swagger:route POST /api/v1/kubernetes/apps/{appid}/persistentvolumeclaims/robin/{pvc}/snapshot/{version}/namespaces/{namespace}/clusters/{cluster} pvc reqCreateSnapshotRobinPersistentVolumeClaim
// create snapshot
// responses:
//
//	200: respSuccessDescription
func (c *RobinPersistentVolumeClaimController) CreateSnapshot() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":pvc")
	version := c.Ctx.Input.Param(":version")

	cli := c.Client(cluster)
	image, err := pvc.GetRbdImageByPvc(cli, name, namespace)
	if err != nil {
		c.HandleError(err)
		return
	}
	robinMetaData, err := clusterRobinMetaData(cluster)
	if err != nil {
		c.HandleError(err)
		return
	}

	result, err := doRobinRequest(
		httplib.Put(fmt.Sprintf(
			"%s/v1/snap/%s/%s",
			robinMetaData.Url,
			image,
			version)), robinMetaData.Token)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(result)

}

// swagger:route DELETE /api/v1/kubernetes/apps/{appid}/persistentvolumeclaims/robin/{pvc}/snapshot/namespaces/{namespace}/clusters/{cluster} pvc reqDeleteAllSnapshotRobinPersistentVolumeClaim
// delete all snapshot
// responses:
//
//	200: respSuccessDescription
func (c *RobinPersistentVolumeClaimController) DeleteAllSnapshot() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":pvc")

	cli := c.Client(cluster)
	image, err := pvc.GetRbdImageByPvc(cli, name, namespace)
	if err != nil {
		c.HandleError(err)
		return
	}
	robinMetaData, err := clusterRobinMetaData(cluster)
	if err != nil {
		c.HandleError(err)
		return
	}

	result, err := doRobinRequest(
		httplib.Delete(fmt.Sprintf(
			"%s/v1/snaps/%s",
			robinMetaData.Url,
			image)), robinMetaData.Token)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(result)
}

// swagger:route DELETE /api/v1/kubernetes/apps/{appid}/persistentvolumeclaims/robin/{pvc}/snapshot/{version}/namespaces/{namespace}/clusters/{cluster} pvc reqDeleteSnapshotRobinPersistentVolumeClaim
// delete snapshot
// responses:
//
//	200: respSuccessDescription
func (c *RobinPersistentVolumeClaimController) DeleteSnapshot() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":pvc")
	version := c.Ctx.Input.Param(":version")

	cli := c.Client(cluster)
	image, err := pvc.GetRbdImageByPvc(cli, name, namespace)
	if err != nil {
		c.HandleError(err)
		return
	}
	robinMetaData, err := clusterRobinMetaData(cluster)
	if err != nil {
		c.HandleError(err)
		return
	}

	result, err := doRobinRequest(
		httplib.Delete(fmt.Sprintf(
			"%s/v1/snap/%s/%s",
			robinMetaData.Url,
			image,
			version)), robinMetaData.Token)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(result)
}

// swagger:route PUT /api/v1/kubernetes/apps/{appid}/persistentvolumeclaims/robin/{pvc}/snapshot/{version}/namespaces/{namespace}/clusters/{cluster} pvc reqRollbackSnapshotRobinPersistentVolumeClaim
// rollback to snapshot version
// responses:
//
//	200: respSuccessDescription
func (c *RobinPersistentVolumeClaimController) RollbackSnapshot() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":pvc")
	version := c.Ctx.Input.Param(":version")

	cli := c.Client(cluster)
	image, err := pvc.GetRbdImageByPvc(cli, name, namespace)
	if err != nil {
		c.HandleError(err)
		return
	}
	robinMetaData, err := clusterRobinMetaData(cluster)
	if err != nil {
		c.HandleError(err)
		return
	}

	result, err := doRobinRequest(
		httplib.Post(fmt.Sprintf(
			"%s/v1/snap/%s/%s",
			robinMetaData.Url,
			image,
			version)), robinMetaData.Token)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(result)
}
