


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNspAssociation(ctx *pulumi.Context, args *LookupNspAssociationArgs, opts ...pulumi.InvokeOption) (*LookupNspAssociationResult, error) {
	var rv LookupNspAssociationResult
	err := ctx.Invoke("azure-native:network/v20210201preview:getNspAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNspAssociationArgs struct {
	AssociationName              string `pulumi:"associationName"`
	NetworkSecurityPerimeterName string `pulumi:"networkSecurityPerimeterName"`
	ResourceGroupName            string `pulumi:"resourceGroupName"`
}


type LookupNspAssociationResult struct {
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

func LookupNspAssociationOutput(ctx *pulumi.Context, args LookupNspAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupNspAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNspAssociationResult, error) {
			args := v.(LookupNspAssociationArgs)
			r, err := LookupNspAssociation(ctx, &args, opts...)
			var s LookupNspAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNspAssociationResultOutput)
}

type LookupNspAssociationOutputArgs struct {
	AssociationName              pulumi.StringInput `pulumi:"associationName"`
	NetworkSecurityPerimeterName pulumi.StringInput `pulumi:"networkSecurityPerimeterName"`
	ResourceGroupName            pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNspAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNspAssociationArgs)(nil)).Elem()
}


type LookupNspAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupNspAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNspAssociationResult)(nil)).Elem()
}

func (o LookupNspAssociationResultOutput) ToLookupNspAssociationResultOutput() LookupNspAssociationResultOutput {
	return o
}

func (o LookupNspAssociationResultOutput) ToLookupNspAssociationResultOutputWithContext(ctx context.Context) LookupNspAssociationResultOutput {
	return o
}

func (o LookupNspAssociationResultOutput) AccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNspAssociationResult) *string { return v.AccessMode }).(pulumi.StringPtrOutput)
}

func (o LookupNspAssociationResultOutput) HasProvisioningIssues() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAssociationResult) string { return v.HasProvisioningIssues }).(pulumi.StringOutput)
}

func (o LookupNspAssociationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAssociationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNspAssociationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNspAssociationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupNspAssociationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAssociationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNspAssociationResultOutput) PrivateLinkResource() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupNspAssociationResult) *SubResourceResponse { return v.PrivateLinkResource }).(SubResourceResponsePtrOutput)
}

func (o LookupNspAssociationResultOutput) Profile() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupNspAssociationResult) *SubResourceResponse { return v.Profile }).(SubResourceResponsePtrOutput)
}

func (o LookupNspAssociationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAssociationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNspAssociationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNspAssociationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNspAssociationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAssociationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNspAssociationResultOutput{})
}
