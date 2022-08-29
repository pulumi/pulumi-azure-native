


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:databricks/v20210401preview:getWorkspace", args, &rv, opts...)
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
	Authorizations             []WorkspaceProviderAuthorizationResponse `pulumi:"authorizations"`
	CreatedBy                  *CreatedByResponse                       `pulumi:"createdBy"`
	CreatedDateTime            string                                   `pulumi:"createdDateTime"`
	Encryption                 *WorkspacePropertiesResponseEncryption   `pulumi:"encryption"`
	Id                         string                                   `pulumi:"id"`
	Location                   string                                   `pulumi:"location"`
	ManagedResourceGroupId     string                                   `pulumi:"managedResourceGroupId"`
	Name                       string                                   `pulumi:"name"`
	Parameters                 *WorkspaceCustomParametersResponse       `pulumi:"parameters"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse      `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                                   `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                                  `pulumi:"publicNetworkAccess"`
	RequiredNsgRules           *string                                  `pulumi:"requiredNsgRules"`
	Sku                        *SkuResponse                             `pulumi:"sku"`
	StorageAccountIdentity     *ManagedIdentityConfigurationResponse    `pulumi:"storageAccountIdentity"`
	SystemData                 SystemDataResponse                       `pulumi:"systemData"`
	Tags                       map[string]string                        `pulumi:"tags"`
	Type                       string                                   `pulumi:"type"`
	UiDefinitionUri            *string                                  `pulumi:"uiDefinitionUri"`
	UpdatedBy                  *CreatedByResponse                       `pulumi:"updatedBy"`
	WorkspaceId                string                                   `pulumi:"workspaceId"`
	WorkspaceUrl               string                                   `pulumi:"workspaceUrl"`
}


func (val *LookupWorkspaceResult) Defaults() *LookupWorkspaceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Parameters = tmp.Parameters.Defaults()

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

func (o LookupWorkspaceResultOutput) Authorizations() WorkspaceProviderAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []WorkspaceProviderAuthorizationResponse { return v.Authorizations }).(WorkspaceProviderAuthorizationResponseArrayOutput)
}

func (o LookupWorkspaceResultOutput) CreatedBy() CreatedByResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *CreatedByResponse { return v.CreatedBy }).(CreatedByResponsePtrOutput)
}

func (o LookupWorkspaceResultOutput) CreatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.CreatedDateTime }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Encryption() WorkspacePropertiesResponseEncryptionPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *WorkspacePropertiesResponseEncryption { return v.Encryption }).(WorkspacePropertiesResponseEncryptionPtrOutput)
}

func (o LookupWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) ManagedResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.ManagedResourceGroupId }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Parameters() WorkspaceCustomParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *WorkspaceCustomParametersResponse { return v.Parameters }).(WorkspaceCustomParametersResponsePtrOutput)
}

func (o LookupWorkspaceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupWorkspaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) RequiredNsgRules() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.RequiredNsgRules }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupWorkspaceResultOutput) StorageAccountIdentity() ManagedIdentityConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *ManagedIdentityConfigurationResponse { return v.StorageAccountIdentity }).(ManagedIdentityConfigurationResponsePtrOutput)
}

func (o LookupWorkspaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupWorkspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) UiDefinitionUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.UiDefinitionUri }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) UpdatedBy() CreatedByResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *CreatedByResponse { return v.UpdatedBy }).(CreatedByResponsePtrOutput)
}

func (o LookupWorkspaceResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) WorkspaceUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.WorkspaceUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceResultOutput{})
}
