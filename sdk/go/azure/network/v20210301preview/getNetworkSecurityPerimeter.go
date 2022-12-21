


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkSecurityPerimeter(ctx *pulumi.Context, args *LookupNetworkSecurityPerimeterArgs, opts ...pulumi.InvokeOption) (*LookupNetworkSecurityPerimeterResult, error) {
	var rv LookupNetworkSecurityPerimeterResult
	err := ctx.Invoke("azure-native:network/v20210301preview:getNetworkSecurityPerimeter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkSecurityPerimeterArgs struct {
	NetworkSecurityPerimeterName string `pulumi:"networkSecurityPerimeterName"`
	ResourceGroupName            string `pulumi:"resourceGroupName"`
}


type LookupNetworkSecurityPerimeterResult struct {
	Description       *string           `pulumi:"description"`
	DisplayName       *string           `pulumi:"displayName"`
	Etag              string            `pulumi:"etag"`
	Id                string            `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}

func LookupNetworkSecurityPerimeterOutput(ctx *pulumi.Context, args LookupNetworkSecurityPerimeterOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkSecurityPerimeterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkSecurityPerimeterResult, error) {
			args := v.(LookupNetworkSecurityPerimeterArgs)
			r, err := LookupNetworkSecurityPerimeter(ctx, &args, opts...)
			var s LookupNetworkSecurityPerimeterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkSecurityPerimeterResultOutput)
}

type LookupNetworkSecurityPerimeterOutputArgs struct {
	NetworkSecurityPerimeterName pulumi.StringInput `pulumi:"networkSecurityPerimeterName"`
	ResourceGroupName            pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkSecurityPerimeterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkSecurityPerimeterArgs)(nil)).Elem()
}


type LookupNetworkSecurityPerimeterResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkSecurityPerimeterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkSecurityPerimeterResult)(nil)).Elem()
}

func (o LookupNetworkSecurityPerimeterResultOutput) ToLookupNetworkSecurityPerimeterResultOutput() LookupNetworkSecurityPerimeterResultOutput {
	return o
}

func (o LookupNetworkSecurityPerimeterResultOutput) ToLookupNetworkSecurityPerimeterResultOutputWithContext(ctx context.Context) LookupNetworkSecurityPerimeterResultOutput {
	return o
}

func (o LookupNetworkSecurityPerimeterResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkSecurityPerimeterResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkSecurityPerimeterResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkSecurityPerimeterResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkSecurityPerimeterResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkSecurityPerimeterResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupNetworkSecurityPerimeterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkSecurityPerimeterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNetworkSecurityPerimeterResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkSecurityPerimeterResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkSecurityPerimeterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkSecurityPerimeterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNetworkSecurityPerimeterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkSecurityPerimeterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNetworkSecurityPerimeterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkSecurityPerimeterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNetworkSecurityPerimeterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkSecurityPerimeterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkSecurityPerimeterResultOutput{})
}
