package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/caarlos0/env"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/kubernetescluster"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

var cfg *config

type config struct {
	// ResourceGroupName is the resource group name for the backend.
	ResourceGroupName string `env:"RESOURCE_GROUP_NAME" envDefault:""`

	// StorageAccountName is the storage account name for the backend.
	StorageAccountName string `env:"STORAGE_ACCOUNT_NAME" envDefault:""`

	// ContainerName is the container name for the backend.
	ContainerName string `env:"CONTAINER_NAME" envDefault:""`

	// Key is the key for the backend.
	Key string `env:"KEY" envDefault:"dev/terraform.tfstate"`
}

func init() {
	cfg = &config{}
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
}

func K8sStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	nodePool := kubernetescluster.NewKubernetesDefault(stack, jsii.String("node-pool"), &kubernetescluster.KubernetesClusterNodePoolConfig{})

	cluster := kubernetescluster.NewKubernetesCluster(scope, jsii.String("demo"), &kubernetescluster.KubernetesClusterConfig{
		Name: jsii.String("demo"),
	})

	cdktf.NewTerraformOutput(stack, jsii.String("kubeconfig"), &cdktf.TerraformOutputConfig{
		Value: cluster.KubeConfigRaw(),
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	stack := K8sStack(app, "example")

	cdktf.NewAzurermBackend(stack, &cdktf.AzurermBackendConfig{
		ResourceGroupName:  jsii.String(cfg.ResourceGroupName),
		StorageAccountName: jsii.String(cfg.StorageAccountName),
		ContainerName:      jsii.String(cfg.ContainerName),
		Key:                jsii.String(cfg.Key),
	})

	app.Synth()
}
