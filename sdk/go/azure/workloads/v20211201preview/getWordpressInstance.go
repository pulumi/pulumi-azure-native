


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWordpressInstance(ctx *pulumi.Context, args *LookupWordpressInstanceArgs, opts ...pulumi.InvokeOption) (*LookupWordpressInstanceResult, error) {
	var rv LookupWordpressInstanceResult
	err := ctx.Invoke("azure-native:workloads/v20211201preview:getWordpressInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWordpressInstanceArgs struct {
	PhpWorkloadName   string `pulumi:"phpWorkloadName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWordpressInstanceResult struct {
	DatabaseName      *string            `pulumi:"databaseName"`
	DatabaseUser      *string            `pulumi:"databaseUser"`
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SiteUrl           string             `pulumi:"siteUrl"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
	Version           string             `pulumi:"version"`
}

func LookupWordpressInstanceOutput(ctx *pulumi.Context, args LookupWordpressInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupWordpressInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWordpressInstanceResult, error) {
			args := v.(LookupWordpressInstanceArgs)
			r, err := LookupWordpressInstance(ctx, &args, opts...)
			var s LookupWordpressInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWordpressInstanceResultOutput)
}

type LookupWordpressInstanceOutputArgs struct {
	PhpWorkloadName   pulumi.StringInput `pulumi:"phpWorkloadName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWordpressInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWordpressInstanceArgs)(nil)).Elem()
}


type LookupWordpressInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupWordpressInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWordpressInstanceResult)(nil)).Elem()
}

func (o LookupWordpressInstanceResultOutput) ToLookupWordpressInstanceResultOutput() LookupWordpressInstanceResultOutput {
	return o
}

func (o LookupWordpressInstanceResultOutput) ToLookupWordpressInstanceResultOutputWithContext(ctx context.Context) LookupWordpressInstanceResultOutput {
	return o
}

func (o LookupWordpressInstanceResultOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWordpressInstanceResult) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

func (o LookupWordpressInstanceResultOutput) DatabaseUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWordpressInstanceResult) *string { return v.DatabaseUser }).(pulumi.StringPtrOutput)
}

func (o LookupWordpressInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWordpressInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWordpressInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWordpressInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWordpressInstanceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWordpressInstanceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWordpressInstanceResultOutput) SiteUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWordpressInstanceResult) string { return v.SiteUrl }).(pulumi.StringOutput)
}

func (o LookupWordpressInstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWordpressInstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupWordpressInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWordpressInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWordpressInstanceResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWordpressInstanceResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWordpressInstanceResultOutput{})
}
