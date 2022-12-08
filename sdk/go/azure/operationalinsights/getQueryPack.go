


package operationalinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupQueryPack(ctx *pulumi.Context, args *LookupQueryPackArgs, opts ...pulumi.InvokeOption) (*LookupQueryPackResult, error) {
	var rv LookupQueryPackResult
	err := ctx.Invoke("azure-native:operationalinsights:getQueryPack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueryPackArgs struct {
	QueryPackName     string `pulumi:"queryPackName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupQueryPackResult struct {
	Id                string            `pulumi:"id"`
	Location          string            `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	QueryPackId       string            `pulumi:"queryPackId"`
	Tags              map[string]string `pulumi:"tags"`
	TimeCreated       string            `pulumi:"timeCreated"`
	TimeModified      string            `pulumi:"timeModified"`
	Type              string            `pulumi:"type"`
}

func LookupQueryPackOutput(ctx *pulumi.Context, args LookupQueryPackOutputArgs, opts ...pulumi.InvokeOption) LookupQueryPackResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQueryPackResult, error) {
			args := v.(LookupQueryPackArgs)
			r, err := LookupQueryPack(ctx, &args, opts...)
			var s LookupQueryPackResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupQueryPackResultOutput)
}

type LookupQueryPackOutputArgs struct {
	QueryPackName     pulumi.StringInput `pulumi:"queryPackName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupQueryPackOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueryPackArgs)(nil)).Elem()
}


type LookupQueryPackResultOutput struct{ *pulumi.OutputState }

func (LookupQueryPackResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueryPackResult)(nil)).Elem()
}

func (o LookupQueryPackResultOutput) ToLookupQueryPackResultOutput() LookupQueryPackResultOutput {
	return o
}

func (o LookupQueryPackResultOutput) ToLookupQueryPackResultOutputWithContext(ctx context.Context) LookupQueryPackResultOutput {
	return o
}

func (o LookupQueryPackResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryPackResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupQueryPackResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryPackResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupQueryPackResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryPackResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupQueryPackResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryPackResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupQueryPackResultOutput) QueryPackId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryPackResult) string { return v.QueryPackId }).(pulumi.StringOutput)
}

func (o LookupQueryPackResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupQueryPackResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupQueryPackResultOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryPackResult) string { return v.TimeCreated }).(pulumi.StringOutput)
}

func (o LookupQueryPackResultOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryPackResult) string { return v.TimeModified }).(pulumi.StringOutput)
}

func (o LookupQueryPackResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryPackResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQueryPackResultOutput{})
}
