


package dbformysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	var rv LookupServerResult
	err := ctx.Invoke("azure-native:dbformysql:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupServerResult struct {
	AdministratorLogin         *string                                   `pulumi:"administratorLogin"`
	ByokEnforcement            string                                    `pulumi:"byokEnforcement"`
	EarliestRestoreDate        *string                                   `pulumi:"earliestRestoreDate"`
	FullyQualifiedDomainName   *string                                   `pulumi:"fullyQualifiedDomainName"`
	Id                         string                                    `pulumi:"id"`
	Identity                   *ResourceIdentityResponse                 `pulumi:"identity"`
	InfrastructureEncryption   *string                                   `pulumi:"infrastructureEncryption"`
	Location                   string                                    `pulumi:"location"`
	MasterServerId             *string                                   `pulumi:"masterServerId"`
	MinimalTlsVersion          *string                                   `pulumi:"minimalTlsVersion"`
	Name                       string                                    `pulumi:"name"`
	PrivateEndpointConnections []ServerPrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess        *string                                   `pulumi:"publicNetworkAccess"`
	ReplicaCapacity            *int                                      `pulumi:"replicaCapacity"`
	ReplicationRole            *string                                   `pulumi:"replicationRole"`
	Sku                        *SkuResponse                              `pulumi:"sku"`
	SslEnforcement             *string                                   `pulumi:"sslEnforcement"`
	StorageProfile             *StorageProfileResponse                   `pulumi:"storageProfile"`
	Tags                       map[string]string                         `pulumi:"tags"`
	Type                       string                                    `pulumi:"type"`
	UserVisibleState           *string                                   `pulumi:"userVisibleState"`
	Version                    *string                                   `pulumi:"version"`
}

func LookupServerOutput(ctx *pulumi.Context, args LookupServerOutputArgs, opts ...pulumi.InvokeOption) LookupServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerResult, error) {
			args := v.(LookupServerArgs)
			r, err := LookupServer(ctx, &args, opts...)
			var s LookupServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerResultOutput)
}

type LookupServerOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerArgs)(nil)).Elem()
}


type LookupServerResultOutput struct{ *pulumi.OutputState }

func (LookupServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerResult)(nil)).Elem()
}

func (o LookupServerResultOutput) ToLookupServerResultOutput() LookupServerResultOutput {
	return o
}

func (o LookupServerResultOutput) ToLookupServerResultOutputWithContext(ctx context.Context) LookupServerResultOutput {
	return o
}

func (o LookupServerResultOutput) AdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.AdministratorLogin }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) ByokEnforcement() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.ByokEnforcement }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) EarliestRestoreDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.EarliestRestoreDate }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) FullyQualifiedDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.FullyQualifiedDomainName }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o LookupServerResultOutput) InfrastructureEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.InfrastructureEncryption }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) MasterServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.MasterServerId }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) MinimalTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.MinimalTlsVersion }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) PrivateEndpointConnections() ServerPrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupServerResult) []ServerPrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(ServerPrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupServerResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) ReplicaCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *int { return v.ReplicaCapacity }).(pulumi.IntPtrOutput)
}

func (o LookupServerResultOutput) ReplicationRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.ReplicationRole }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupServerResultOutput) SslEnforcement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.SslEnforcement }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *StorageProfileResponse { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

func (o LookupServerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) UserVisibleState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.UserVisibleState }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerResultOutput{})
}
