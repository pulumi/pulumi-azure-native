


package v20201201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:synapse/v20201201:getWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupWorkspaceResult struct {
	AdlaResourceId                   string                                    `pulumi:"adlaResourceId"`
	ConnectivityEndpoints            map[string]string                         `pulumi:"connectivityEndpoints"`
	DefaultDataLakeStorage           *DataLakeStorageAccountDetailsResponse    `pulumi:"defaultDataLakeStorage"`
	Encryption                       *EncryptionDetailsResponse                `pulumi:"encryption"`
	ExtraProperties                  map[string]interface{}                    `pulumi:"extraProperties"`
	Id                               string                                    `pulumi:"id"`
	Identity                         *ManagedIdentityResponse                  `pulumi:"identity"`
	Location                         string                                    `pulumi:"location"`
	ManagedResourceGroupName         *string                                   `pulumi:"managedResourceGroupName"`
	ManagedVirtualNetwork            *string                                   `pulumi:"managedVirtualNetwork"`
	ManagedVirtualNetworkSettings    *ManagedVirtualNetworkSettingsResponse    `pulumi:"managedVirtualNetworkSettings"`
	Name                             string                                    `pulumi:"name"`
	PrivateEndpointConnections       []PrivateEndpointConnectionResponse       `pulumi:"privateEndpointConnections"`
	ProvisioningState                string                                    `pulumi:"provisioningState"`
	PurviewConfiguration             *PurviewConfigurationResponse             `pulumi:"purviewConfiguration"`
	SqlAdministratorLogin            *string                                   `pulumi:"sqlAdministratorLogin"`
	SqlAdministratorLoginPassword    *string                                   `pulumi:"sqlAdministratorLoginPassword"`
	Tags                             map[string]string                         `pulumi:"tags"`
	Type                             string                                    `pulumi:"type"`
	VirtualNetworkProfile            *VirtualNetworkProfileResponse            `pulumi:"virtualNetworkProfile"`
	WorkspaceRepositoryConfiguration *WorkspaceRepositoryConfigurationResponse `pulumi:"workspaceRepositoryConfiguration"`
	WorkspaceUID                     string                                    `pulumi:"workspaceUID"`
}

func LookupWorkspaceOutput(ctx *pulumi.Context, args LookupWorkspaceOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceResult, error) {
			args := v.(LookupWorkspaceArgs)
			r, err := LookupWorkspace(ctx, &args, opts...)
			var s LookupWorkspaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceResultOutput)
}

type LookupWorkspaceOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWorkspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceArgs)(nil)).Elem()
}


type LookupWorkspaceResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceResult)(nil)).Elem()
}

func (o LookupWorkspaceResultOutput) ToLookupWorkspaceResultOutput() LookupWorkspaceResultOutput {
	return o
}

func (o LookupWorkspaceResultOutput) ToLookupWorkspaceResultOutputWithContext(ctx context.Context) LookupWorkspaceResultOutput {
	return o
}

func (o LookupWorkspaceResultOutput) AdlaResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.AdlaResourceId }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) ConnectivityEndpoints() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) map[string]string { return v.ConnectivityEndpoints }).(pulumi.StringMapOutput)
}

func (o LookupWorkspaceResultOutput) DefaultDataLakeStorage() DataLakeStorageAccountDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *DataLakeStorageAccountDetailsResponse { return v.DefaultDataLakeStorage }).(DataLakeStorageAccountDetailsResponsePtrOutput)
}

func (o LookupWorkspaceResultOutput) Encryption() EncryptionDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *EncryptionDetailsResponse { return v.Encryption }).(EncryptionDetailsResponsePtrOutput)
}

func (o LookupWorkspaceResultOutput) ExtraProperties() pulumi.MapOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) map[string]interface{} { return v.ExtraProperties }).(pulumi.MapOutput)
}

func (o LookupWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Identity() ManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *ManagedIdentityResponse { return v.Identity }).(ManagedIdentityResponsePtrOutput)
}

func (o LookupWorkspaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) ManagedResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ManagedResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) ManagedVirtualNetwork() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ManagedVirtualNetwork }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) ManagedVirtualNetworkSettings() ManagedVirtualNetworkSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *ManagedVirtualNetworkSettingsResponse {
		return v.ManagedVirtualNetworkSettings
	}).(ManagedVirtualNetworkSettingsResponsePtrOutput)
}

func (o LookupWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupWorkspaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) PurviewConfiguration() PurviewConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *PurviewConfigurationResponse { return v.PurviewConfiguration }).(PurviewConfigurationResponsePtrOutput)
}

func (o LookupWorkspaceResultOutput) SqlAdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.SqlAdministratorLogin }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) SqlAdministratorLoginPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.SqlAdministratorLoginPassword }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) VirtualNetworkProfile() VirtualNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *VirtualNetworkProfileResponse { return v.VirtualNetworkProfile }).(VirtualNetworkProfileResponsePtrOutput)
}

func (o LookupWorkspaceResultOutput) WorkspaceRepositoryConfiguration() WorkspaceRepositoryConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *WorkspaceRepositoryConfigurationResponse {
		return v.WorkspaceRepositoryConfiguration
	}).(WorkspaceRepositoryConfigurationResponsePtrOutput)
}

func (o LookupWorkspaceResultOutput) WorkspaceUID() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.WorkspaceUID }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceResultOutput{})
}
