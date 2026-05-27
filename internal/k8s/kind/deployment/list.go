package deployment

import (
	"wayne/internal/k8s/client"
	"wayne/internal/k8s/kind/dataselector"
	"wayne/pkg/dto"
)

func GetDeploymentPage(indexer *client.CacheFactory, namespace string, q *dto.QueryParam) (*dto.Page, error) {
	kubeDeployments, err := GetDeploymentList(indexer, namespace)
	if err != nil {
		return nil, err
	}

	deployments := make([]Deployment, 0)

	for i := 0; i < len(kubeDeployments); i++ {
		deploy, err := toDeployment(kubeDeployments[i], indexer)
		if err != nil {
			return nil, err
		}
		deployments = append(deployments, *deploy)
	}

	return dataselector.DataSelectPage(toCells(deployments), q), nil
}

func toCells(deploy []Deployment) []dataselector.DataCell {
	cells := make([]dataselector.DataCell, len(deploy))
	for i := range deploy {
		cells[i] = DeploymentCell(deploy[i])
	}
	return cells
}
