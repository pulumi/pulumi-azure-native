


package v20150801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListConnectionKeys(ctx *pulumi.Context, args *ListConnectionKeysArgs, opts ...pulumi.InvokeOption) (*ListConnectionKeysResult, error) {
	var rv ListConnectionKeysResult
	err := ctx.Invoke("azure-native:web/v20150801preview:listConnectionKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConnectionKeysArgs struct {
	ConnectionName    string            `pulumi:"connectionName"`
	Id                *string           `pulumi:"id"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	Type              *string           `pulumi:"type"`
	ValidityTimeSpan  *string           `pulumi:"validityTimeSpan"`
}

type ListConnectionKeysResult struct {
	ConnectionKey   *string                `pulumi:"connectionKey"`
	ParameterValues map[string]interface{} `pulumi:"parameterValues"`
}

func ListConnectionKeysOutput(ctx *pulumi.Context, args ListConnectionKeysOutputArgs, opts ...pulumi.InvokeOption) ListConnectionKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListConnectionKeysResult, error) {
			args := v.(ListConnectionKeysArgs)
			r, err := ListConnectionKeys(ctx, &args, opts...)
			var s ListConnectionKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListConnectionKeysResultOutput)
}

type ListConnectionKeysOutputArgs struct {
	ConnectionName    pulumi.StringInput    `pulumi:"connectionName"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Kind              pulumi.StringPtrInput `pulumi:"kind"`
	Location          pulumi.StringPtrInput `pulumi:"location"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	Tags              pulumi.StringMapInput `pulumi:"tags"`
	Type              pulumi.StringPtrInput `pulumi:"type"`
	ValidityTimeSpan  pulumi.StringPtrInput `pulumi:"validityTimeSpan"`
}

func (ListConnectionKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectionKeysArgs)(nil)).Elem()
}

type ListConnectionKeysResultOutput struct{ *pulumi.OutputState }

func (ListConnectionKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectionKeysResult)(nil)).Elem()
}

func (o ListConnectionKeysResultOutput) ToListConnectionKeysResultOutput() ListConnectionKeysResultOutput {
	return o
}

func (o ListConnectionKeysResultOutput) ToListConnectionKeysResultOutputWithContext(ctx context.Context) ListConnectionKeysResultOutput {
	return o
}

func (o ListConnectionKeysResultOutput) ConnectionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListConnectionKeysResult) *string { return v.ConnectionKey }).(pulumi.StringPtrOutput)
}

func (o ListConnectionKeysResultOutput) ParameterValues() pulumi.MapOutput {
	return o.ApplyT(func(v ListConnectionKeysResult) map[string]interface{} { return v.ParameterValues }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterOutputType(ListConnectionKeysResultOutput{})
}
