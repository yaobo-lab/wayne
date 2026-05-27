export const defaultIngress = `{
  "apiVersion": "networking.k8s.io/v1",
  "kind": "Ingress",
  "metadata": {
    "name": ""
  },
  "spec": {
    "tls": [

    ],
    "rules": [
      {
        "host": "",
        "http": {
          "paths": [
            {
              "pathType": "Prefix",
              "path": "/",
              "backend": {
                "service": {
                  "name": "",
                  "port": {
                    "number": 80
                  }
                }
              }
            }
          ]
        }
      }
    ]
  }
}`;
