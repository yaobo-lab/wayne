package clear

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"wayne/clear/dal/query"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DataBase *gorm.DB

func TestMain(m *testing.M) {
	var err error
	DataBase, err = gorm.Open(mysql.Open("root:@Pwdabc7@tcp(10.0.1.10:3306)/wayne_db?charset=utf8&parseTime=True&loc=Local&allowNativePasswords=true"))
	if err != nil {
		panic(err)
	}
	query.SetDefault(DataBase)
	m.Run()
}

// 清理部署
func TestClearDeployment(t *testing.T) {

	rep := query.DeploymentTemplate

	list, err := rep.Where(rep.ID.Gt(0)).Find()

	if err != nil {
		panic(err)
	}

	for _, item := range list {
		item.Template = strings.Replace(item.Template, "extensions/v1beta1", "apps/v1", -1)
		_, err := rep.Where(rep.ID.Eq(item.ID)).Update(rep.Template, item.Template)
		if err != nil {
			fmt.Printf("ID:%d \n", item.ID)
			panic(err)
		}
	}

}

func TestDemoIngress(t *testing.T) {
	rep := query.IngressTemplate
	list, err := rep.Where(rep.ID.Gt(367)).Find()
	if err != nil {
		panic(err)
	}

	for _, item := range list {

		fmt.Println(item.Template)
		fmt.Printf("\n tpl end ========================= \n\n")

		oldModel := OldIngress{}
		err := json.Unmarshal([]byte(item.Template), &oldModel)
		if err != nil {
			panic(err)
		}

		bo, err := json.Marshal(oldModel)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(bo))
		fmt.Printf("\n old end ========================= \n\n")

		newModel := NewIngress{}
		err = json.Unmarshal([]byte(item.Template), &newModel)
		if err != nil {
			panic(err)
		}

		//更新接口地址
		newModel.APIVersion = "networking.k8s.io/v1"

		if newModel.Metadata.Annotations == nil {
			newModel.Metadata.Annotations = map[string]string{}
		}

		//修改 Annotations
		for key, anns := range newModel.Metadata.Annotations {
			switch key {
			case "configuration.konghq.com":
				delete(newModel.Metadata.Annotations, key)
				newModel.Metadata.Annotations["konghq.com/override"] = "config-script"
			case "plugins.konghq.com":
				delete(newModel.Metadata.Annotations, key)
				newModel.Metadata.Annotations["konghq.com/plugins"] = anns
			}
		}
		newModel.Metadata.Annotations["kubernetes.io/ingress.class"] = "kong"

		//修改转发
		for i, rule := range newModel.Spec.Rules {
			for k, _ := range rule.HTTP.Paths {
				newModel.Spec.Rules[i].HTTP.Paths[k].PathType = "Prefix"
				newModel.Spec.Rules[i].HTTP.Paths[k].Backend.Service.Name = oldModel.Spec.Rules[i].HTTP.Paths[k].Backend.ServiceName
				newModel.Spec.Rules[i].HTTP.Paths[k].Backend.Service.Port.Number = oldModel.Spec.Rules[i].HTTP.Paths[k].Backend.ServicePort
			}
		}

		b, err := json.Marshal(newModel)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
		fmt.Printf("\n new end ========================= \n\n")

	}
}

// 清理ingress
func TestClearIngress(t *testing.T) {
	rep := query.IngressTemplate
	list, err := rep.Where(rep.ID.Gt(0)).Find()
	if err != nil {
		panic(err)
	}

	for _, item := range list {

		oldModel := OldIngress{}
		err := json.Unmarshal([]byte(item.Template), &oldModel)
		if err != nil {
			panic(err)
		}

		newModel := NewIngress{}
		err = json.Unmarshal([]byte(item.Template), &newModel)
		if err != nil {
			panic(err)
		}

		//更新接口地址
		newModel.APIVersion = "networking.k8s.io/v1"
		if newModel.Metadata.Annotations == nil {
			newModel.Metadata.Annotations = map[string]string{}
		}
		//修改 Annotations
		for key, anns := range newModel.Metadata.Annotations {
			switch key {
			case "configuration.konghq.com":
				delete(newModel.Metadata.Annotations, key)
				newModel.Metadata.Annotations["konghq.com/override"] = "config-script"
			case "plugins.konghq.com":
				delete(newModel.Metadata.Annotations, key)
				newModel.Metadata.Annotations["konghq.com/plugins"] = anns
			}
		}
		newModel.Metadata.Annotations["kubernetes.io/ingress.class"] = "kong"

		for i, rule := range newModel.Spec.Rules {
			for k, _ := range rule.HTTP.Paths {
				newModel.Spec.Rules[i].HTTP.Paths[k].PathType = "Prefix"
				newModel.Spec.Rules[i].HTTP.Paths[k].Backend.Service.Name = oldModel.Spec.Rules[i].HTTP.Paths[k].Backend.ServiceName
				newModel.Spec.Rules[i].HTTP.Paths[k].Backend.Service.Port.Number = oldModel.Spec.Rules[i].HTTP.Paths[k].Backend.ServicePort
			}
		}

		b, err := json.Marshal(newModel)
		if err != nil {
			panic(err)
		}

		item.Template = string(b)

		_, err = rep.Where(rep.ID.Eq(item.ID)).Update(rep.Template, item.Template)
		if err != nil {
			fmt.Printf("ID:%d \n", item.ID)
			panic(err)
		}

	}
}

// 检查k8s 用了多少个域名
func TestIngressHost(t *testing.T) {
	rep := query.IngressTemplate
	list, err := rep.Where(rep.ID.Gt(0)).Find()
	if err != nil {
		panic(err)
	}

	hosts := map[string]string{}

	for _, item := range list {

		newModel := NewIngress{}
		err = json.Unmarshal([]byte(item.Template), &newModel)
		if err != nil {
			panic(err)
		}
		hosts[newModel.Spec.Rules[0].Host] = ""
	}

	fmt.Println(hosts)
}
