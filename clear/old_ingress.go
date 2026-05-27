package clear

/*
{
  "apiVersion": "extensions/v1beta1",
  "kind": "Ingress",
  "metadata": {
    "name": "productmanagement-api-ingress",
    "labels": {
      "wayne-app": "productmanagement",
      "wayne-ns": "oa-service",
      "app": "productmanagement-api-ingress"
    },
    "annotations": {
      "configuration.konghq.com": "strip-config",
      "plugins.konghq.com": "oidc-token-productmanagement-api"
    }
  },
  "spec": {
    "tls": [

    ],
    "rules": [
      {
        "host": "gw.meda.test",
        "http": {
          "paths": [
            {
              "backend": {
                "serviceName": "productmanagement-api-service",
                "servicePort": 80
              },
              "path": "/productmanagement"
            }
          ]
        }
      }
    ]
  }
}
*/
type OldIngress struct {
	APIVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec"`
}

type Metadata struct {
	Name        string      `json:"name"`
	Labels      Labels      `json:"labels"`
	Annotations Annotations `json:"annotations"`
}

type Annotations map[string]string
type Labels map[string]string

type Backend struct {
	ServiceName string `json:"serviceName"`
	ServicePort int    `json:"servicePort"`
}
type Paths struct {
	Backend Backend `json:"backend"`
	Path    string  `json:"path"`
}
type HTTP struct {
	Paths []Paths `json:"paths"`
}
type Rules struct {
	Host string `json:"host"`
	HTTP HTTP   `json:"http"`
}
type Spec struct {
	TLS   []interface{} `json:"tls"`
	Rules []Rules       `json:"rules"`
}
