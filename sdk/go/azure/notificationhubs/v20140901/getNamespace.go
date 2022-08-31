


package v20140901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupNamespace(ctx *pulumi.Context, args *LookupNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceResult, error) {
	var rv LookupNamespaceResult
	err := ctx.Invoke("azure-native:notificationhubs/v20140901:getNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceArgs struct {
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupNamespaceResult struct {
	Id         *string                     `pulumi:"id"`
	Location   *string                     `pulumi:"location"`
	Name       *string                     `pulumi:"name"`
	Properties NamespacePropertiesResponse `pulumi:"properties"`
	Tags       map[string]string           `pulumi:"tags"`
	Type       *string                     `pulumi:"type"`
}

func LookupNamespaceOutput(ctx *pulumi.Context, args LookupNamespaceOutputArgs, opts ...pulumi.InvokeOption) LookupNamespaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNamespaceResult, error) {
			args := v.(LookupNamespaceArgs)
			r, err := LookupNamespace(ctx, &args, opts...)
			var s LookupNamespaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNamespaceResultOutput)
}

type LookupNamespaceOutputArgs struct {
	NamespaceName     pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNamespaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceArgs)(nil)).Elem()
}


type LookupNamespaceResultOutput struct{ *pulumi.OutputState }

func (LookupNamespaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceResult)(nil)).Elem()
}

func (o LookupNamespaceResultOutput) ToLookupNamespaceResultOutput() LookupNamespaceResultOutput {
	return o
}

func (o LookupNamespaceResultOutput) ToLookupNamespaceResultOutputWithContext(ctx context.Context) LookupNamespaceResultOutput {
	return o
}

func (o LookupNamespaceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupNamespaceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupNamespaceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupNamespaceResultOutput) Properties() NamespacePropertiesResponseOutput {
	return o.ApplyT(func(v LookupNamespaceResult) NamespacePropertiesResponse { return v.Properties }).(NamespacePropertiesResponseOutput)
}

func (o LookupNamespaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNamespaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNamespaceResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamespaceResultOutput{})
}
