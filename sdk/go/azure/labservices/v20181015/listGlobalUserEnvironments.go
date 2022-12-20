


package v20181015

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListGlobalUserEnvironments(ctx *pulumi.Context, args *ListGlobalUserEnvironmentsArgs, opts ...pulumi.InvokeOption) (*ListGlobalUserEnvironmentsResult, error) {
	var rv ListGlobalUserEnvironmentsResult
	err := ctx.Invoke("azure-native:labservices/v20181015:listGlobalUserEnvironments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListGlobalUserEnvironmentsArgs struct {
	LabId    *string `pulumi:"labId"`
	UserName string  `pulumi:"userName"`
}


type ListGlobalUserEnvironmentsResult struct {
	Environments []EnvironmentDetailsResponse `pulumi:"environments"`
}

func ListGlobalUserEnvironmentsOutput(ctx *pulumi.Context, args ListGlobalUserEnvironmentsOutputArgs, opts ...pulumi.InvokeOption) ListGlobalUserEnvironmentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListGlobalUserEnvironmentsResult, error) {
			args := v.(ListGlobalUserEnvironmentsArgs)
			r, err := ListGlobalUserEnvironments(ctx, &args, opts...)
			var s ListGlobalUserEnvironmentsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListGlobalUserEnvironmentsResultOutput)
}

type ListGlobalUserEnvironmentsOutputArgs struct {
	LabId    pulumi.StringPtrInput `pulumi:"labId"`
	UserName pulumi.StringInput    `pulumi:"userName"`
}

func (ListGlobalUserEnvironmentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListGlobalUserEnvironmentsArgs)(nil)).Elem()
}


type ListGlobalUserEnvironmentsResultOutput struct{ *pulumi.OutputState }

func (ListGlobalUserEnvironmentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListGlobalUserEnvironmentsResult)(nil)).Elem()
}

func (o ListGlobalUserEnvironmentsResultOutput) ToListGlobalUserEnvironmentsResultOutput() ListGlobalUserEnvironmentsResultOutput {
	return o
}

func (o ListGlobalUserEnvironmentsResultOutput) ToListGlobalUserEnvironmentsResultOutputWithContext(ctx context.Context) ListGlobalUserEnvironmentsResultOutput {
	return o
}

func (o ListGlobalUserEnvironmentsResultOutput) Environments() EnvironmentDetailsResponseArrayOutput {
	return o.ApplyT(func(v ListGlobalUserEnvironmentsResult) []EnvironmentDetailsResponse { return v.Environments }).(EnvironmentDetailsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListGlobalUserEnvironmentsResultOutput{})
}
