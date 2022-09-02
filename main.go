package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/cdktf-provider-ionoscloud-go/ionoscloud/v2"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

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
			CpuFamily:        jsii.String("INTEL_XEON"),
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

	K8sExampleStack(app, "example")

	app.Synth()
}
