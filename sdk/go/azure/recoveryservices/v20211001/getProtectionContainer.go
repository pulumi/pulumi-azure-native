


package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProtectionContainer(ctx *pulumi.Context, args *LookupProtectionContainerArgs, opts ...pulumi.InvokeOption) (*LookupProtectionContainerResult, error) {
	var rv LookupProtectionContainerResult
	err := ctx.Invoke("azure-native:recoveryservices/v20211001:getProtectionContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProtectionContainerArgs struct {
	ContainerName     string `pulumi:"containerName"`
	FabricName        string `pulumi:"fabricName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VaultName         string `pulumi:"vaultName"`
}


type LookupProtectionContainerResult struct {
	ETag       *string           `pulumi:"eTag"`
	Id         string            `pulumi:"id"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}

func LookupProtectionContainerOutput(ctx *pulumi.Context, args LookupProtectionContainerOutputArgs, opts ...pulumi.InvokeOption) LookupProtectionContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProtectionContainerResult, error) {
			args := v.(LookupProtectionContainerArgs)
			r, err := LookupProtectionContainer(ctx, &args, opts...)
			var s LookupProtectionContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProtectionContainerResultOutput)
}

type LookupProtectionContainerOutputArgs struct {
	ContainerName     pulumi.StringInput `pulumi:"containerName"`
	FabricName        pulumi.StringInput `pulumi:"fabricName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VaultName         pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupProtectionContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProtectionContainerArgs)(nil)).Elem()
}


type LookupProtectionContainerResultOutput struct{ *pulumi.OutputState }

func (LookupProtectionContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProtectionContainerResult)(nil)).Elem()
}

func (o LookupProtectionContainerResultOutput) ToLookupProtectionContainerResultOutput() LookupProtectionContainerResultOutput {
	return o
}

func (o LookupProtectionContainerResultOutput) ToLookupProtectionContainerResultOutputWithContext(ctx context.Context) LookupProtectionContainerResultOutput {
	return o
}

func (o LookupProtectionContainerResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProtectionContainerResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupProtectionContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectionContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProtectionContainerResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProtectionContainerResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupProtectionContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectionContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProtectionContainerResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupProtectionContainerResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupProtectionContainerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProtectionContainerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupProtectionContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectionContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProtectionContainerResultOutput{})
}
