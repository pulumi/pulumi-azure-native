


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNspAssociationsProxy(ctx *pulumi.Context, args *LookupNspAssociationsProxyArgs, opts ...pulumi.InvokeOption) (*LookupNspAssociationsProxyResult, error) {
	var rv LookupNspAssociationsProxyResult
	err := ctx.Invoke("azure-native:network:getNspAssociationsProxy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNspAssociationsProxyArgs struct {
	AssociationName              string `pulumi:"associationName"`
	NetworkSecurityPerimeterName string `pulumi:"networkSecurityPerimeterName"`
	ResourceGroupName            string `pulumi:"resourceGroupName"`
}


type LookupNspAssociationsProxyResult struct {
	AccessMode            *string              `pulumi:"accessMode"`
	HasProvisioningIssues string               `pulumi:"hasProvisioningIssues"`
	Id                    string               `pulumi:"id"`
	Location              *string              `pulumi:"location"`
	Name                  string               `pulumi:"name"`
	PrivateLinkResource   *SubResourceResponse `pulumi:"privateLinkResource"`
	Profile               *SubResourceResponse `pulumi:"profile"`
	ProvisioningState     string               `pulumi:"provisioningState"`
	Tags                  map[string]string    `pulumi:"tags"`
	Type                  string               `pulumi:"type"`
}

func LookupNspAssociationsProxyOutput(ctx *pulumi.Context, args LookupNspAssociationsProxyOutputArgs, opts ...pulumi.InvokeOption) LookupNspAssociationsProxyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNspAssociationsProxyResult, error) {
			args := v.(LookupNspAssociationsProxyArgs)
			r, err := LookupNspAssociationsProxy(ctx, &args, opts...)
			var s LookupNspAssociationsProxyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNspAssociationsProxyResultOutput)
}

type LookupNspAssociationsProxyOutputArgs struct {
	AssociationName              pulumi.StringInput `pulumi:"associationName"`
	NetworkSecurityPerimeterName pulumi.StringInput `pulumi:"networkSecurityPerimeterName"`
	ResourceGroupName            pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNspAssociationsProxyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNspAssociationsProxyArgs)(nil)).Elem()
}


type LookupNspAssociationsProxyResultOutput struct{ *pulumi.OutputState }

func (LookupNspAssociationsProxyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNspAssociationsProxyResult)(nil)).Elem()
}

func (o LookupNspAssociationsProxyResultOutput) ToLookupNspAssociationsProxyResultOutput() LookupNspAssociationsProxyResultOutput {
	return o
}

func (o LookupNspAssociationsProxyResultOutput) ToLookupNspAssociationsProxyResultOutputWithContext(ctx context.Context) LookupNspAssociationsProxyResultOutput {
	return o
}

func (o LookupNspAssociationsProxyResultOutput) AccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNspAssociationsProxyResult) *string { return v.AccessMode }).(pulumi.StringPtrOutput)
}

func (o LookupNspAssociationsProxyResultOutput) HasProvisioningIssues() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAssociationsProxyResult) string { return v.HasProvisioningIssues }).(pulumi.StringOutput)
}

func (o LookupNspAssociationsProxyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAssociationsProxyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNspAssociationsProxyResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNspAssociationsProxyResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupNspAssociationsProxyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAssociationsProxyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNspAssociationsProxyResultOutput) PrivateLinkResource() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupNspAssociationsProxyResult) *SubResourceResponse { return v.PrivateLinkResource }).(SubResourceResponsePtrOutput)
}

func (o LookupNspAssociationsProxyResultOutput) Profile() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupNspAssociationsProxyResult) *SubResourceResponse { return v.Profile }).(SubResourceResponsePtrOutput)
}

func (o LookupNspAssociationsProxyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAssociationsProxyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNspAssociationsProxyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNspAssociationsProxyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNspAssociationsProxyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAssociationsProxyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNspAssociationsProxyResultOutput{})
}
