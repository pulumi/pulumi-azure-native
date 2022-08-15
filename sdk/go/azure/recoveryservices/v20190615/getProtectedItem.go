


package v20190615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProtectedItem(ctx *pulumi.Context, args *LookupProtectedItemArgs, opts ...pulumi.InvokeOption) (*LookupProtectedItemResult, error) {
	var rv LookupProtectedItemResult
	err := ctx.Invoke("azure-native:recoveryservices/v20190615:getProtectedItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProtectedItemArgs struct {
	ContainerName     string  `pulumi:"containerName"`
	FabricName        string  `pulumi:"fabricName"`
	Filter            *string `pulumi:"filter"`
	ProtectedItemName string  `pulumi:"protectedItemName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	VaultName         string  `pulumi:"vaultName"`
}


type LookupProtectedItemResult struct {
	ETag       *string           `pulumi:"eTag"`
	Id         string            `pulumi:"id"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}

func LookupProtectedItemOutput(ctx *pulumi.Context, args LookupProtectedItemOutputArgs, opts ...pulumi.InvokeOption) LookupProtectedItemResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProtectedItemResult, error) {
			args := v.(LookupProtectedItemArgs)
			r, err := LookupProtectedItem(ctx, &args, opts...)
			var s LookupProtectedItemResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProtectedItemResultOutput)
}

type LookupProtectedItemOutputArgs struct {
	ContainerName     pulumi.StringInput    `pulumi:"containerName"`
	FabricName        pulumi.StringInput    `pulumi:"fabricName"`
	Filter            pulumi.StringPtrInput `pulumi:"filter"`
	ProtectedItemName pulumi.StringInput    `pulumi:"protectedItemName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	VaultName         pulumi.StringInput    `pulumi:"vaultName"`
}

func (LookupProtectedItemOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProtectedItemArgs)(nil)).Elem()
}


type LookupProtectedItemResultOutput struct{ *pulumi.OutputState }

func (LookupProtectedItemResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProtectedItemResult)(nil)).Elem()
}

func (o LookupProtectedItemResultOutput) ToLookupProtectedItemResultOutput() LookupProtectedItemResultOutput {
	return o
}

func (o LookupProtectedItemResultOutput) ToLookupProtectedItemResultOutputWithContext(ctx context.Context) LookupProtectedItemResultOutput {
	return o
}

func (o LookupProtectedItemResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProtectedItemResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupProtectedItemResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectedItemResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProtectedItemResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProtectedItemResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupProtectedItemResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectedItemResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProtectedItemResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupProtectedItemResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupProtectedItemResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProtectedItemResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupProtectedItemResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectedItemResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProtectedItemResultOutput{})
}
