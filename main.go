package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/caarlos0/env"
	"github.com/hashicorp/cdktf-provider-ionoscloud-go/ionoscloud/v2"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

var cfg *config

type config struct {
	// AccessKey is the access key for the S3 backend.
	AccessKey string `env:"ACCESS_KEY" envDefault:""`

	// SecretKey is the secret key for the S3 backend.
	SecretKey string `env:"SECRET_KEY" envDefault:""`

	// BucketName is the bucket name for the S3 backend.
	BucketName string `env:"BUCKET_NAME" envDefault:""`

	// Region is the region for the S3 backend.
	Region string `env:"REGION" envDefault:"de"`

	// Endpoint is the endpoint for the S3 backend.
	Endpoint string `env:"ENDPOINT" envDefault:"https://S3-eu-central-1.ionoscloud.com"`

	// Key is the key for the S3 backend.
	Key string `env:"KEY" envDefault:"dev/terraform.tfstate"`
}

func init() {
	cfg = &config{}
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
}

func K8sExampleStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	ionoscloud.NewIonoscloudProvider(
		stack,
		jsii.String("ionoscloud"),
		&ionoscloud.IonoscloudProviderConfig{},
	)

	datacenter := ionoscloud.NewDatacenter(
		stack,
		jsii.String("datacenter"),
		&ionoscloud.DatacenterConfig{
			Name:     jsii.String("demo"),
			Location: jsii.String("de/fra"),
		},
	)

	k8sCluster := ionoscloud.NewK8SCluster(
		stack,
		jsii.String("k8sClusterExample"),
		&ionoscloud.K8SClusterConfig{
			Name:       jsii.String("k8sClusterExample"),
			K8SVersion: jsii.String("1.23.9"),
		},
	)

	ionoscloud.NewK8SNodePool(
		stack,
		jsii.String("k8sNodePoolExample"),
		&ionoscloud.K8SNodePoolConfig{
			DatacenterId: datacenter.Id(),
			K8SClusterId: k8sCluster.Id(),
			Name:         jsii.String("k8sNodePoolExample"),
			K8SVersion:   k8sCluster.K8SVersion(),
			MaintenanceWindow: &ionoscloud.K8SNodePoolMaintenanceWindow{
				DayOfTheWeek: jsii.String("Monday"),
				Time:         jsii.String("09:00:00Z"),
			},
			AutoScaling: &ionoscloud.K8SNodePoolAutoScaling{
				MinNodeCount: jsii.Number(1),
				MaxNodeCount: jsii.Number(1),
			},
			CpuFamily:        datacenter.CpuArchitecture().Get(jsii.Number(0)).CpuFamily(),
			AvailabilityZone: jsii.String("AUTO"),
			StorageType:      jsii.String("SSD"),
			NodeCount:        jsii.Number(1),
			CoresCount:       jsii.Number(1),
			RamSize:          jsii.Number(2048),
			StorageSize:      jsii.Number(40),
		},
	)

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	stack := K8sExampleStack(app, "example")

	cdktf.NewS3Backend(stack, &cdktf.S3BackendProps{
		Region:    jsii.String(cfg.Region),
		Key:       jsii.String(cfg.Key),
		Bucket:    jsii.String(cfg.BucketName),
		AccessKey: jsii.String(cfg.AccessKey),
		SecretKey: jsii.String(cfg.SecretKey),
		Endpoint:  jsii.String(cfg.Endpoint),

		// Some extra settings to make the remote state work.
		SkipMetadataApiCheck:      jsii.Bool(true),
		SkipRegionValidation:      jsii.Bool(true),
		SkipCredentialsValidation: jsii.Bool(true),
		ForcePathStyle:            jsii.Bool(true),
	})

	app.Synth()
}
