


package v20200701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetBastionShareableLink(ctx *pulumi.Context, args *GetBastionShareableLinkArgs, opts ...pulumi.InvokeOption) (*GetBastionShareableLinkResult, error) {
	var rv GetBastionShareableLinkResult
	err := ctx.Invoke("azure-native:network/v20200701:getBastionShareableLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetBastionShareableLinkArgs struct {
	BastionHostName   string                 `pulumi:"bastionHostName"`
	ResourceGroupName string                 `pulumi:"resourceGroupName"`
	Vms               []BastionShareableLink `pulumi:"vms"`
}


type GetBastionShareableLinkResult struct {
	NextLink *string                        `pulumi:"nextLink"`
	Value    []BastionShareableLinkResponse `pulumi:"value"`
}

func GetBastionShareableLinkOutput(ctx *pulumi.Context, args GetBastionShareableLinkOutputArgs, opts ...pulumi.InvokeOption) GetBastionShareableLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBastionShareableLinkResult, error) {
			args := v.(GetBastionShareableLinkArgs)
			r, err := GetBastionShareableLink(ctx, &args, opts...)
			var s GetBastionShareableLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBastionShareableLinkResultOutput)
}

type GetBastionShareableLinkOutputArgs struct {
	BastionHostName   pulumi.StringInput             `pulumi:"bastionHostName"`
	ResourceGroupName pulumi.StringInput             `pulumi:"resourceGroupName"`
	Vms               BastionShareableLinkArrayInput `pulumi:"vms"`
}

func (GetBastionShareableLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBastionShareableLinkArgs)(nil)).Elem()
}


type GetBastionShareableLinkResultOutput struct{ *pulumi.OutputState }

func (GetBastionShareableLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBastionShareableLinkResult)(nil)).Elem()
}

func (o GetBastionShareableLinkResultOutput) ToGetBastionShareableLinkResultOutput() GetBastionShareableLinkResultOutput {
	return o
}

func (o GetBastionShareableLinkResultOutput) ToGetBastionShareableLinkResultOutputWithContext(ctx context.Context) GetBastionShareableLinkResultOutput {
	return o
}

func (o GetBastionShareableLinkResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBastionShareableLinkResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o GetBastionShareableLinkResultOutput) Value() BastionShareableLinkResponseArrayOutput {
	return o.ApplyT(func(v GetBastionShareableLinkResult) []BastionShareableLinkResponse { return v.Value }).(BastionShareableLinkResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBastionShareableLinkResultOutput{})
}
