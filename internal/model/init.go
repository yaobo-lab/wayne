package model

import (
	"sync"

	"github.com/beego/beego/v2/adapter/orm"
)

var (
	globalOrm orm.Ormer
	once      sync.Once

	UserModel                     *userModel
	AppModel                      *appModel
	AppUserModel                  *appUserModel
	AppStarredModel               *appStarredModel
	NamespaceUserModel            *namespaceUserModel
	ClusterModel                  *clusterModel
	DeploymentModel               *deploymentModel
	DeploymentTplModel            *deploymentTplModel
	PermissionModel               *permissionModel
	GroupModel                    *groupModel
	NamespaceModel                *namespaceModel
	ConfigMapModel                *configMapModel
	ConfigMapTplModel             *configMapTplModel
	ServiceModel                  *serviceModel
	ServiceTplModel               *serviceTplModel
	CronjobModel                  *cronjobModel
	CronjobTplModel               *cronjobTplModel
	SecretModel                   *secretModel
	SecretTplModel                *secretTplModel
	PublishStatusModel            *publishStatusModel
	PersistentVolumeClaimModel    *persistentVolumeClaimModel
	PersistentVolumeClaimTplModel *persistentVolumeClaimTplModel

	ApiKeyModel         *apiKeyModel
	StatefulsetModel    *statefulsetModel
	StatefulsetTplModel *statefulsetTplModel
	ConfigModel         *configModel
	DaemonSetModel      *daemonSetModel
	DaemonSetTplModel   *daemonSetTplModel

	IngressModel         *ingressModel
	IngressTemplateModel *ingressTemplateModel
	HPAModel             *hpaModel
	HPATemplateModel     *hpaTemplateModel
)

func init() {
	// init orm tables
	orm.RegisterModel(
		new(User),
		new(App),
		new(AppStarred),
		new(AppUser),
		new(NamespaceUser),
		new(Cluster),
		new(Namespace),
		new(Deployment),
		new(DeploymentTemplate),
		new(Service),
		new(ServiceTemplate),
		new(Group),
		new(Permission),
		new(Secret),
		new(SecretTemplate),
		new(ConfigMap),
		new(ConfigMapTemplate),
		new(Cronjob),
		new(CronjobTemplate),
		new(PublishStatus),
		new(PersistentVolumeClaim),
		new(PersistentVolumeClaimTemplate),
		new(APIKey),
		new(Statefulset),
		new(StatefulsetTemplate),
		new(Config),
		new(DaemonSet),
		new(DaemonSetTemplate),

		new(Ingress),
		new(IngressTemplate),
		new(HPA),
		new(HPATemplate))

	// init models
	UserModel = &userModel{}
	AppModel = &appModel{}
	AppUserModel = &appUserModel{}
	AppStarredModel = &appStarredModel{}
	NamespaceUserModel = &namespaceUserModel{}
	ClusterModel = &clusterModel{}
	NamespaceModel = &namespaceModel{}
	DeploymentModel = &deploymentModel{}
	DeploymentTplModel = &deploymentTplModel{}
	GroupModel = &groupModel{}
	SecretModel = &secretModel{}
	SecretTplModel = &secretTplModel{}
	ConfigMapModel = &configMapModel{}
	ConfigMapTplModel = &configMapTplModel{}
	CronjobModel = &cronjobModel{}
	CronjobTplModel = &cronjobTplModel{}
	PublishStatusModel = &publishStatusModel{}
	PersistentVolumeClaimModel = &persistentVolumeClaimModel{}
	PersistentVolumeClaimTplModel = &persistentVolumeClaimTplModel{}

	ApiKeyModel = &apiKeyModel{}
	StatefulsetModel = &statefulsetModel{}
	StatefulsetTplModel = &statefulsetTplModel{}
	ConfigModel = &configModel{}
	DaemonSetModel = &daemonSetModel{}
	DaemonSetTplModel = &daemonSetTplModel{}

	IngressModel = &ingressModel{}
	IngressTemplateModel = &ingressTemplateModel{}
	HPAModel = &hpaModel{}
	HPATemplateModel = &hpaTemplateModel{}
}

// singleton init ormer ,only use for normal db operation
// if you begin transaction，please use orm.NewOrm()
func Ormer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}
