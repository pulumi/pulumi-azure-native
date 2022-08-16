


package v20180202

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	var rv LookupGroupResult
	err := ctx.Invoke("azure-native:migrate/v20180202:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGroupArgs struct {
	GroupName         string `pulumi:"groupName"`
	ProjectName       string `pulumi:"projectName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupGroupResult struct {
	Assessments      []string `pulumi:"assessments"`
	CreatedTimestamp string   `pulumi:"createdTimestamp"`
	ETag             *string  `pulumi:"eTag"`
	Id               string   `pulumi:"id"`
	Machines         []string `pulumi:"machines"`
	Name             string   `pulumi:"name"`
	Type             string   `pulumi:"type"`
	UpdatedTimestamp string   `pulumi:"updatedTimestamp"`
}

func LookupGroupOutput(ctx *pulumi.Context, args LookupGroupOutputArgs, opts ...pulumi.InvokeOption) LookupGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGroupResult, error) {
			args := v.(LookupGroupArgs)
			r, err := LookupGroup(ctx, &args, opts...)
			var s LookupGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGroupResultOutput)
}

type LookupGroupOutputArgs struct {
	GroupName         pulumi.StringInput `pulumi:"groupName"`
	ProjectName       pulumi.StringInput `pulumi:"projectName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupArgs)(nil)).Elem()
}


type LookupGroupResultOutput struct{ *pulumi.OutputState }

func (LookupGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupResult)(nil)).Elem()
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutput() LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutputWithContext(ctx context.Context) LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) Assessments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.Assessments }).(pulumi.StringArrayOutput)
}

func (o LookupGroupResultOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) Machines() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.Machines }).(pulumi.StringArrayOutput)
}

func (o LookupGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupResultOutput{})
}
