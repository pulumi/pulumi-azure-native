


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20220601preview:getWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWorkspaceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupWorkspaceResult struct {
	AllowPublicAccessWhenBehindVnet *bool                                    `pulumi:"allowPublicAccessWhenBehindVnet"`
	ApplicationInsights             *string                                  `pulumi:"applicationInsights"`
	ContainerRegistry               *string                                  `pulumi:"containerRegistry"`
	Description                     *string                                  `pulumi:"description"`
	DiscoveryUrl                    *string                                  `pulumi:"discoveryUrl"`
	Encryption                      *EncryptionPropertyResponse              `pulumi:"encryption"`
	FriendlyName                    *string                                  `pulumi:"friendlyName"`
	HbiWorkspace                    *bool                                    `pulumi:"hbiWorkspace"`
	Id                              string                                   `pulumi:"id"`
	Identity                        *ManagedServiceIdentityResponse          `pulumi:"identity"`
	ImageBuildCompute               *string                                  `pulumi:"imageBuildCompute"`
	KeyVault                        *string                                  `pulumi:"keyVault"`
	Location                        *string                                  `pulumi:"location"`
	MlFlowTrackingUri               string                                   `pulumi:"mlFlowTrackingUri"`
	Name                            string                                   `pulumi:"name"`
	NotebookInfo                    NotebookResourceInfoResponse             `pulumi:"notebookInfo"`
	PrimaryUserAssignedIdentity     *string                                  `pulumi:"primaryUserAssignedIdentity"`
	PrivateEndpointConnections      []PrivateEndpointConnectionResponse      `pulumi:"privateEndpointConnections"`
	PrivateLinkCount                int                                      `pulumi:"privateLinkCount"`
	ProvisioningState               string                                   `pulumi:"provisioningState"`
	PublicNetworkAccess             *string                                  `pulumi:"publicNetworkAccess"`
	ScheduledPurgeDate              string                                   `pulumi:"scheduledPurgeDate"`
	ServiceManagedResourcesSettings *ServiceManagedResourcesSettingsResponse `pulumi:"serviceManagedResourcesSettings"`
	ServiceProvisionedResourceGroup string                                   `pulumi:"serviceProvisionedResourceGroup"`
	SharedPrivateLinkResources      []SharedPrivateLinkResourceResponse      `pulumi:"sharedPrivateLinkResources"`
	Sku                             *SkuResponse                             `pulumi:"sku"`
	SoftDeletedAt                   string                                   `pulumi:"softDeletedAt"`
	StorageAccount                  *string                                  `pulumi:"storageAccount"`
	StorageHnsEnabled               bool                                     `pulumi:"storageHnsEnabled"`
	SystemData                      SystemDataResponse                       `pulumi:"systemData"`
	Tags                            map[string]string                        `pulumi:"tags"`
	TenantId                        string                                   `pulumi:"tenantId"`
	Type                            string                                   `pulumi:"type"`
	V1LegacyMode                    *bool                                    `pulumi:"v1LegacyMode"`
	WorkspaceId                     string                                   `pulumi:"workspaceId"`
}


func (val *LookupWorkspaceResult) Defaults() *LookupWorkspaceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllowPublicAccessWhenBehindVnet) {
		allowPublicAccessWhenBehindVnet_ := false
		tmp.AllowPublicAccessWhenBehindVnet = &allowPublicAccessWhenBehindVnet_
	}
	if isZero(tmp.HbiWorkspace) {
		hbiWorkspace_ := false
		tmp.HbiWorkspace = &hbiWorkspace_
	}
	if isZero(tmp.V1LegacyMode) {
		v1LegacyMode_ := false
		tmp.V1LegacyMode = &v1LegacyMode_
	}
	return &tmp
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

func (o LookupWorkspaceResultOutput) AllowPublicAccessWhenBehindVnet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *bool { return v.AllowPublicAccessWhenBehindVnet }).(pulumi.BoolPtrOutput)
}

func (o LookupWorkspaceResultOutput) ApplicationInsights() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ApplicationInsights }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) ContainerRegistry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ContainerRegistry }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) DiscoveryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.DiscoveryUrl }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Encryption() EncryptionPropertyResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *EncryptionPropertyResponse { return v.Encryption }).(EncryptionPropertyResponsePtrOutput)
}

func (o LookupWorkspaceResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) HbiWorkspace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *bool { return v.HbiWorkspace }).(pulumi.BoolPtrOutput)
}

func (o LookupWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupWorkspaceResultOutput) ImageBuildCompute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ImageBuildCompute }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) KeyVault() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.KeyVault }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) MlFlowTrackingUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.MlFlowTrackingUri }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) NotebookInfo() NotebookResourceInfoResponseOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) NotebookResourceInfoResponse { return v.NotebookInfo }).(NotebookResourceInfoResponseOutput)
}

func (o LookupWorkspaceResultOutput) PrimaryUserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.PrimaryUserAssignedIdentity }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupWorkspaceResultOutput) PrivateLinkCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) int { return v.PrivateLinkCount }).(pulumi.IntOutput)
}

func (o LookupWorkspaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) ScheduledPurgeDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.ScheduledPurgeDate }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) ServiceManagedResourcesSettings() ServiceManagedResourcesSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *ServiceManagedResourcesSettingsResponse {
		return v.ServiceManagedResourcesSettings
	}).(ServiceManagedResourcesSettingsResponsePtrOutput)
}

func (o LookupWorkspaceResultOutput) ServiceProvisionedResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.ServiceProvisionedResourceGroup }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) SharedPrivateLinkResources() SharedPrivateLinkResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []SharedPrivateLinkResourceResponse { return v.SharedPrivateLinkResources }).(SharedPrivateLinkResourceResponseArrayOutput)
}

func (o LookupWorkspaceResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupWorkspaceResultOutput) SoftDeletedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.SoftDeletedAt }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) StorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.StorageAccount }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) StorageHnsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) bool { return v.StorageHnsEnabled }).(pulumi.BoolOutput)
}

func (o LookupWorkspaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupWorkspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWorkspaceResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) V1LegacyMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *bool { return v.V1LegacyMode }).(pulumi.BoolPtrOutput)
}

func (o LookupWorkspaceResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceResultOutput{})
}
