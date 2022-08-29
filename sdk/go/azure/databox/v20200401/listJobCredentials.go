


package v20200401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListJobCredentials(ctx *pulumi.Context, args *ListJobCredentialsArgs, opts ...pulumi.InvokeOption) (*ListJobCredentialsResult, error) {
	var rv ListJobCredentialsResult
	err := ctx.Invoke("azure-native:databox/v20200401:listJobCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListJobCredentialsArgs struct {
	JobName           string `pulumi:"jobName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListJobCredentialsResult struct {
	NextLink *string                          `pulumi:"nextLink"`
	Value    []UnencryptedCredentialsResponse `pulumi:"value"`
}

func ListJobCredentialsOutput(ctx *pulumi.Context, args ListJobCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListJobCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListJobCredentialsResult, error) {
			args := v.(ListJobCredentialsArgs)
			r, err := ListJobCredentials(ctx, &args, opts...)
			var s ListJobCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListJobCredentialsResultOutput)
}

type ListJobCredentialsOutputArgs struct {
	JobName           pulumi.StringInput `pulumi:"jobName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListJobCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListJobCredentialsArgs)(nil)).Elem()
}


type ListJobCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListJobCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListJobCredentialsResult)(nil)).Elem()
}

func (o ListJobCredentialsResultOutput) ToListJobCredentialsResultOutput() ListJobCredentialsResultOutput {
	return o
}

func (o ListJobCredentialsResultOutput) ToListJobCredentialsResultOutputWithContext(ctx context.Context) ListJobCredentialsResultOutput {
	return o
}

func (o ListJobCredentialsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListJobCredentialsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListJobCredentialsResultOutput) Value() UnencryptedCredentialsResponseArrayOutput {
	return o.ApplyT(func(v ListJobCredentialsResult) []UnencryptedCredentialsResponse { return v.Value }).(UnencryptedCredentialsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListJobCredentialsResultOutput{})
}
