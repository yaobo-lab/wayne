package clear

/*
{
  "apiVersion": "networking.k8s.io/v1",
  "kind": "Ingress",
  "metadata": {
    "name": "devops-platform-api",
    "labels": {
      "wayne-app": "devops-platform",
      "wayne-ns": "public-service",
      "app": "devops-platform-api"
    },
    "annotations": {
      "kubernetes.io/ingress.class": "kong",
      "konghq.com/override": "config-script",
      "konghq.com/plugins": "oidc-token-svc-devops-api"
    }
  },
  "spec": {
    "tls": [

    ],
    "rules": [
      {
        "host": "k8s-dev.yourdomain.com",
        "http": {
          "paths": [
            {
              "path": "/devopsapi",
              "backend": {
                "service": {
                  "name": "devops-platform-api",
                  "port": {
                    "number": 80
                  }
                }
              },
              "pathType": "Prefix"
            }
          ]
        }
      }
    ]
  }
}
*/
type NewIngress struct {
	APIVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"metadata"`
	Spec       NewSpec  `json:"spec"`
}

type NewSpec struct {
	TLS   []interface{} `json:"tls"`
	Rules []NewRules    `json:"rules"`
}
type NewPort struct {
	Number int `json:"number"`
}
type NewService struct {
	Name string  `json:"name"`
	Port NewPort `json:"port"`
}
type NewBackend struct {
	Service NewService `json:"service"`
}
type NewPaths struct {
	Path     string     `json:"path"`
	Backend  NewBackend `json:"backend"`
	PathType string     `json:"pathType"`
}
type NewHTTP struct {
	Paths []NewPaths `json:"paths"`
}
type NewRules struct {
	Host string  `json:"host"`
	HTTP NewHTTP `json:"http"`
}
