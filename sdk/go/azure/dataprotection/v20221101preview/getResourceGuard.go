


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupResourceGuard(ctx *pulumi.Context, args *LookupResourceGuardArgs, opts ...pulumi.InvokeOption) (*LookupResourceGuardResult, error) {
	var rv LookupResourceGuardResult
	err := ctx.Invoke("azure-native:dataprotection/v20221101preview:getResourceGuard", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceGuardArgs struct {
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	ResourceGuardsName string `pulumi:"resourceGuardsName"`
}

type LookupResourceGuardResult struct {
	ETag       *string                     `pulumi:"eTag"`
	Id         string                      `pulumi:"id"`
	Identity   *DppIdentityDetailsResponse `pulumi:"identity"`
	Location   *string                     `pulumi:"location"`
	Name       string                      `pulumi:"name"`
	Properties ResourceGuardResponse       `pulumi:"properties"`
	SystemData SystemDataResponse          `pulumi:"systemData"`
	Tags       map[string]string           `pulumi:"tags"`
	Type       string                      `pulumi:"type"`
}

func LookupResourceGuardOutput(ctx *pulumi.Context, args LookupResourceGuardOutputArgs, opts ...pulumi.InvokeOption) LookupResourceGuardResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceGuardResult, error) {
			args := v.(LookupResourceGuardArgs)
			r, err := LookupResourceGuard(ctx, &args, opts...)
			var s LookupResourceGuardResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceGuardResultOutput)
}

type LookupResourceGuardOutputArgs struct {
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceGuardsName pulumi.StringInput `pulumi:"resourceGuardsName"`
}

func (LookupResourceGuardOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceGuardArgs)(nil)).Elem()
}

type LookupResourceGuardResultOutput struct{ *pulumi.OutputState }

func (LookupResourceGuardResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceGuardResult)(nil)).Elem()
}

func (o LookupResourceGuardResultOutput) ToLookupResourceGuardResultOutput() LookupResourceGuardResultOutput {
	return o
}

func (o LookupResourceGuardResultOutput) ToLookupResourceGuardResultOutputWithContext(ctx context.Context) LookupResourceGuardResultOutput {
	return o
}

func (o LookupResourceGuardResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceGuardResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupResourceGuardResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGuardResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupResourceGuardResultOutput) Identity() DppIdentityDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupResourceGuardResult) *DppIdentityDetailsResponse { return v.Identity }).(DppIdentityDetailsResponsePtrOutput)
}

func (o LookupResourceGuardResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceGuardResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupResourceGuardResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGuardResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupResourceGuardResultOutput) Properties() ResourceGuardResponseOutput {
	return o.ApplyT(func(v LookupResourceGuardResult) ResourceGuardResponse { return v.Properties }).(ResourceGuardResponseOutput)
}

func (o LookupResourceGuardResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupResourceGuardResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupResourceGuardResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupResourceGuardResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupResourceGuardResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGuardResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceGuardResultOutput{})
}
