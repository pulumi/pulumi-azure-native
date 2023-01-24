


package v20210114preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20210114preview:getWorkspace", args, &rv, opts...)
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
	ApplicationGroupReferences []string                                             `pulumi:"applicationGroupReferences"`
	CloudPcResource            bool                                                 `pulumi:"cloudPcResource"`
	Description                *string                                              `pulumi:"description"`
	Etag                       string                                               `pulumi:"etag"`
	FriendlyName               *string                                              `pulumi:"friendlyName"`
	Id                         string                                               `pulumi:"id"`
	Identity                   *ResourceModelWithAllowedPropertySetResponseIdentity `pulumi:"identity"`
	Kind                       *string                                              `pulumi:"kind"`
	Location                   *string                                              `pulumi:"location"`
	ManagedBy                  *string                                              `pulumi:"managedBy"`
	Name                       string                                               `pulumi:"name"`
	ObjectId                   string                                               `pulumi:"objectId"`
	Plan                       *ResourceModelWithAllowedPropertySetResponsePlan     `pulumi:"plan"`
	Sku                        *ResourceModelWithAllowedPropertySetResponseSku      `pulumi:"sku"`
	Tags                       map[string]string                                    `pulumi:"tags"`
	Type                       string                                               `pulumi:"type"`
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

func (o LookupWorkspaceResultOutput) ApplicationGroupReferences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []string { return v.ApplicationGroupReferences }).(pulumi.StringArrayOutput)
}

func (o LookupWorkspaceResultOutput) CloudPcResource() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) bool { return v.CloudPcResource }).(pulumi.BoolOutput)
}

func (o LookupWorkspaceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Identity() ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *ResourceModelWithAllowedPropertySetResponseIdentity { return v.Identity }).(ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput)
}

func (o LookupWorkspaceResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Plan() ResourceModelWithAllowedPropertySetResponsePlanPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *ResourceModelWithAllowedPropertySetResponsePlan { return v.Plan }).(ResourceModelWithAllowedPropertySetResponsePlanPtrOutput)
}

func (o LookupWorkspaceResultOutput) Sku() ResourceModelWithAllowedPropertySetResponseSkuPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *ResourceModelWithAllowedPropertySetResponseSku { return v.Sku }).(ResourceModelWithAllowedPropertySetResponseSkuPtrOutput)
}

func (o LookupWorkspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceResultOutput{})
}
