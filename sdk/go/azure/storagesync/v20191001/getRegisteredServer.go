


package v20191001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupRegisteredServer(ctx *pulumi.Context, args *LookupRegisteredServerArgs, opts ...pulumi.InvokeOption) (*LookupRegisteredServerResult, error) {
	var rv LookupRegisteredServerResult
	err := ctx.Invoke("azure-native:storagesync/v20191001:getRegisteredServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegisteredServerArgs struct {
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	ServerId               string `pulumi:"serverId"`
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
}


type LookupRegisteredServerResult struct {
	AgentVersion              *string `pulumi:"agentVersion"`
	ClusterId                 *string `pulumi:"clusterId"`
	ClusterName               *string `pulumi:"clusterName"`
	DiscoveryEndpointUri      *string `pulumi:"discoveryEndpointUri"`
	FriendlyName              *string `pulumi:"friendlyName"`
	Id                        string  `pulumi:"id"`
	LastHeartBeat             *string `pulumi:"lastHeartBeat"`
	LastOperationName         *string `pulumi:"lastOperationName"`
	LastWorkflowId            *string `pulumi:"lastWorkflowId"`
	ManagementEndpointUri     *string `pulumi:"managementEndpointUri"`
	MonitoringConfiguration   *string `pulumi:"monitoringConfiguration"`
	Name                      string  `pulumi:"name"`
	ProvisioningState         *string `pulumi:"provisioningState"`
	ResourceLocation          *string `pulumi:"resourceLocation"`
	ServerCertificate         *string `pulumi:"serverCertificate"`
	ServerId                  *string `pulumi:"serverId"`
	ServerManagementErrorCode *int    `pulumi:"serverManagementErrorCode"`
	ServerOSVersion           *string `pulumi:"serverOSVersion"`
	ServerRole                *string `pulumi:"serverRole"`
	ServiceLocation           *string `pulumi:"serviceLocation"`
	StorageSyncServiceUid     *string `pulumi:"storageSyncServiceUid"`
	Type                      string  `pulumi:"type"`
}

func LookupRegisteredServerOutput(ctx *pulumi.Context, args LookupRegisteredServerOutputArgs, opts ...pulumi.InvokeOption) LookupRegisteredServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegisteredServerResult, error) {
			args := v.(LookupRegisteredServerArgs)
			r, err := LookupRegisteredServer(ctx, &args, opts...)
			var s LookupRegisteredServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegisteredServerResultOutput)
}

type LookupRegisteredServerOutputArgs struct {
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerId               pulumi.StringInput `pulumi:"serverId"`
	StorageSyncServiceName pulumi.StringInput `pulumi:"storageSyncServiceName"`
}

func (LookupRegisteredServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegisteredServerArgs)(nil)).Elem()
}


type LookupRegisteredServerResultOutput struct{ *pulumi.OutputState }

func (LookupRegisteredServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegisteredServerResult)(nil)).Elem()
}

func (o LookupRegisteredServerResultOutput) ToLookupRegisteredServerResultOutput() LookupRegisteredServerResultOutput {
	return o
}

func (o LookupRegisteredServerResultOutput) ToLookupRegisteredServerResultOutputWithContext(ctx context.Context) LookupRegisteredServerResultOutput {
	return o
}

func (o LookupRegisteredServerResultOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.AgentVersion }).(pulumi.StringPtrOutput)
}

func (o LookupRegisteredServerResultOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

func (o LookupRegisteredServerResultOutput) ClusterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.ClusterName }).(pulumi.StringPtrOutput)
}

func (o LookupRegisteredServerResultOutput) DiscoveryEndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.DiscoveryEndpointUri }).(pulumi.StringPtrOutput)
}

func (o LookupRegisteredServerResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o LookupRegisteredServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRegisteredServerResultOutput) LastHeartBeat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.LastHeartBeat }).(pulumi.StringPtrOutput)
}

func (o LookupRegisteredServerResultOutput) LastOperationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.LastOperationName }).(pulumi.StringPtrOutput)
}

func (o LookupRegisteredServerResultOutput) LastWorkflowId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.LastWorkflowId }).(pulumi.StringPtrOutput)
}

func (o LookupRegisteredServerResultOutput) ManagementEndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.ManagementEndpointUri }).(pulumi.StringPtrOutput)
}

func (o LookupRegisteredServerResultOutput) MonitoringConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.MonitoringConfiguration }).(pulumi.StringPtrOutput)
}

func (o LookupRegisteredServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRegisteredServerResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupRegisteredServerResultOutput) ResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.ResourceLocation }).(pulumi.StringPtrOutput)
}

func (o LookupRegisteredServerResultOutput) ServerCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.ServerCertificate }).(pulumi.StringPtrOutput)
}

func (o LookupRegisteredServerResultOutput) ServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.ServerId }).(pulumi.StringPtrOutput)
}

func (o LookupRegisteredServerResultOutput) ServerManagementErrorCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *int { return v.ServerManagementErrorCode }).(pulumi.IntPtrOutput)
}

func (o LookupRegisteredServerResultOutput) ServerOSVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.ServerOSVersion }).(pulumi.StringPtrOutput)
}

func (o LookupRegisteredServerResultOutput) ServerRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.ServerRole }).(pulumi.StringPtrOutput)
}

func (o LookupRegisteredServerResultOutput) ServiceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.ServiceLocation }).(pulumi.StringPtrOutput)
}

func (o LookupRegisteredServerResultOutput) StorageSyncServiceUid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.StorageSyncServiceUid }).(pulumi.StringPtrOutput)
}

func (o LookupRegisteredServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegisteredServerResultOutput{})
}
