


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:operationalinsights/v20211201preview:getWorkspace", args, &rv, opts...)
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
	CreatedDate                         string                              `pulumi:"createdDate"`
	CustomerId                          string                              `pulumi:"customerId"`
	DefaultDataCollectionRuleResourceId *string                             `pulumi:"defaultDataCollectionRuleResourceId"`
	ETag                                *string                             `pulumi:"eTag"`
	Features                            *WorkspaceFeaturesResponse          `pulumi:"features"`
	ForceCmkForQuery                    *bool                               `pulumi:"forceCmkForQuery"`
	Id                                  string                              `pulumi:"id"`
	Location                            string                              `pulumi:"location"`
	ModifiedDate                        string                              `pulumi:"modifiedDate"`
	Name                                string                              `pulumi:"name"`
	PrivateLinkScopedResources          []PrivateLinkScopedResourceResponse `pulumi:"privateLinkScopedResources"`
	ProvisioningState                   string                              `pulumi:"provisioningState"`
	PublicNetworkAccessForIngestion     *string                             `pulumi:"publicNetworkAccessForIngestion"`
	PublicNetworkAccessForQuery         *string                             `pulumi:"publicNetworkAccessForQuery"`
	RetentionInDays                     *int                                `pulumi:"retentionInDays"`
	Sku                                 *WorkspaceSkuResponse               `pulumi:"sku"`
	SystemData                          SystemDataResponse                  `pulumi:"systemData"`
	Tags                                map[string]string                   `pulumi:"tags"`
	Type                                string                              `pulumi:"type"`
	WorkspaceCapping                    *WorkspaceCappingResponse           `pulumi:"workspaceCapping"`
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

func (o LookupWorkspaceResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) CustomerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.CustomerId }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) DefaultDataCollectionRuleResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.DefaultDataCollectionRuleResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Features() WorkspaceFeaturesResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *WorkspaceFeaturesResponse { return v.Features }).(WorkspaceFeaturesResponsePtrOutput)
}

func (o LookupWorkspaceResultOutput) ForceCmkForQuery() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *bool { return v.ForceCmkForQuery }).(pulumi.BoolPtrOutput)
}

func (o LookupWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) ModifiedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.ModifiedDate }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) PrivateLinkScopedResources() PrivateLinkScopedResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []PrivateLinkScopedResourceResponse { return v.PrivateLinkScopedResources }).(PrivateLinkScopedResourceResponseArrayOutput)
}

func (o LookupWorkspaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) PublicNetworkAccessForIngestion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.PublicNetworkAccessForIngestion }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) PublicNetworkAccessForQuery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.PublicNetworkAccessForQuery }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o LookupWorkspaceResultOutput) Sku() WorkspaceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *WorkspaceSkuResponse { return v.Sku }).(WorkspaceSkuResponsePtrOutput)
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

func (o LookupWorkspaceResultOutput) WorkspaceCapping() WorkspaceCappingResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *WorkspaceCappingResponse { return v.WorkspaceCapping }).(WorkspaceCappingResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceResultOutput{})
}
