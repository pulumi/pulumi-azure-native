


package v20151101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:operationalinsights/v20151101preview:getWorkspace", args, &rv, opts...)
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
	CustomerId        string            `pulumi:"customerId"`
	ETag              *string           `pulumi:"eTag"`
	Id                string            `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	PortalUrl         string            `pulumi:"portalUrl"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	RetentionInDays   *int              `pulumi:"retentionInDays"`
	Sku               *SkuResponse      `pulumi:"sku"`
	Source            string            `pulumi:"source"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
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

func (o LookupWorkspaceResultOutput) CustomerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.CustomerId }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) PortalUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.PortalUrl }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o LookupWorkspaceResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupWorkspaceResultOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Source }).(pulumi.StringOutput)
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
