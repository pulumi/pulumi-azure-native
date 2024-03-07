package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicefabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := servicefabric.NewManagedClusterApplication(ctx, "managedClusterApplication", &servicefabric.ManagedClusterApplicationArgs{
ApplicationName: pulumi.String("myApp"),
ClusterName: pulumi.String("myCluster"),
Location: pulumi.String("eastus"),
Parameters: pulumi.StringMap{
"param1": pulumi.String("value1"),
},
ResourceGroupName: pulumi.String("resRg"),
Tags: pulumi.StringMap{
"a": pulumi.String("b"),
},
UpgradePolicy: servicefabric.ApplicationUpgradePolicyResponse{
ApplicationHealthPolicy: interface{}{
ConsiderWarningAsError: pulumi.Bool(true),
DefaultServiceTypeHealthPolicy: &servicefabric.ServiceTypeHealthPolicyArgs{
MaxPercentUnhealthyPartitionsPerService: pulumi.Int(0),
MaxPercentUnhealthyReplicasPerPartition: pulumi.Int(0),
MaxPercentUnhealthyServices: pulumi.Int(0),
},
MaxPercentUnhealthyDeployedApplications: pulumi.Int(0),
ServiceTypeHealthPolicyMap: servicefabric.ServiceTypeHealthPolicyMap{
"myService": &servicefabric.ServiceTypeHealthPolicyArgs{
MaxPercentUnhealthyPartitionsPerService: pulumi.Int(30),
MaxPercentUnhealthyReplicasPerPartition: pulumi.Int(30),
MaxPercentUnhealthyServices: pulumi.Int(30),
},
},
},
ForceRestart: pulumi.Bool(false),
InstanceCloseDelayDuration: pulumi.Float64(600),
RecreateApplication: pulumi.Bool(false),
RollingUpgradeMonitoringPolicy: &servicefabric.RollingUpgradeMonitoringPolicyArgs{
FailureAction: pulumi.String("Rollback"),
HealthCheckRetryTimeout: pulumi.String("00:10:00"),
HealthCheckStableDuration: pulumi.String("00:05:00"),
HealthCheckWaitDuration: pulumi.String("00:02:00"),
UpgradeDomainTimeout: pulumi.String("00:15:00"),
UpgradeTimeout: pulumi.String("01:00:00"),
},
UpgradeMode: pulumi.String("UnmonitoredAuto"),
UpgradeReplicaSetCheckTimeout: pulumi.Float64(3600),
},
Version: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedclusters/myCluster/applicationTypes/myAppType/versions/1.0"),
})
if err != nil {
return err
}
return nil
})
}
