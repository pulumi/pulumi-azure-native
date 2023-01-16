


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	var rv LookupServerResult
	err := ctx.Invoke("azure-native:sql/v20220501preview:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerArgs struct {
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServerName        string  `pulumi:"serverName"`
}


type LookupServerResult struct {
	AdministratorLogin            *string                                   `pulumi:"administratorLogin"`
	Administrators                *ServerExternalAdministratorResponse      `pulumi:"administrators"`
	FederatedClientId             *string                                   `pulumi:"federatedClientId"`
	FullyQualifiedDomainName      string                                    `pulumi:"fullyQualifiedDomainName"`
	Id                            string                                    `pulumi:"id"`
	Identity                      *ResourceIdentityResponse                 `pulumi:"identity"`
	KeyId                         *string                                   `pulumi:"keyId"`
	Kind                          string                                    `pulumi:"kind"`
	Location                      string                                    `pulumi:"location"`
	MinimalTlsVersion             *string                                   `pulumi:"minimalTlsVersion"`
	Name                          string                                    `pulumi:"name"`
	PrimaryUserAssignedIdentityId *string                                   `pulumi:"primaryUserAssignedIdentityId"`
	PrivateEndpointConnections    []ServerPrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess           *string                                   `pulumi:"publicNetworkAccess"`
	RestrictOutboundNetworkAccess *string                                   `pulumi:"restrictOutboundNetworkAccess"`
	State                         string                                    `pulumi:"state"`
	Tags                          map[string]string                         `pulumi:"tags"`
	Type                          string                                    `pulumi:"type"`
	Version                       *string                                   `pulumi:"version"`
	WorkspaceFeature              string                                    `pulumi:"workspaceFeature"`
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
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput    `pulumi:"serverName"`
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

func (o LookupServerResultOutput) Administrators() ServerExternalAdministratorResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *ServerExternalAdministratorResponse { return v.Administrators }).(ServerExternalAdministratorResponsePtrOutput)
}

func (o LookupServerResultOutput) FederatedClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.FederatedClientId }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) FullyQualifiedDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.FullyQualifiedDomainName }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o LookupServerResultOutput) KeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.KeyId }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) MinimalTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.MinimalTlsVersion }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) PrimaryUserAssignedIdentityId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.PrimaryUserAssignedIdentityId }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) PrivateEndpointConnections() ServerPrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupServerResult) []ServerPrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(ServerPrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupServerResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) RestrictOutboundNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.RestrictOutboundNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) WorkspaceFeature() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.WorkspaceFeature }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerResultOutput{})
}
