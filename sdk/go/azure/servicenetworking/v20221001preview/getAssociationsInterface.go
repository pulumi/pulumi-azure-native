


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAssociationsInterface(ctx *pulumi.Context, args *LookupAssociationsInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupAssociationsInterfaceResult, error) {
	var rv LookupAssociationsInterfaceResult
	err := ctx.Invoke("azure-native:servicenetworking/v20221001preview:getAssociationsInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssociationsInterfaceArgs struct {
	AssociationName       string `pulumi:"associationName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	TrafficControllerName string `pulumi:"trafficControllerName"`
}


type LookupAssociationsInterfaceResult struct {
	AssociationType   string                     `pulumi:"associationType"`
	Id                string                     `pulumi:"id"`
	Location          string                     `pulumi:"location"`
	Name              string                     `pulumi:"name"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	Subnet            *AssociationSubnetResponse `pulumi:"subnet"`
	SystemData        SystemDataResponse         `pulumi:"systemData"`
	Tags              map[string]string          `pulumi:"tags"`
	Type              string                     `pulumi:"type"`
}

func LookupAssociationsInterfaceOutput(ctx *pulumi.Context, args LookupAssociationsInterfaceOutputArgs, opts ...pulumi.InvokeOption) LookupAssociationsInterfaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAssociationsInterfaceResult, error) {
			args := v.(LookupAssociationsInterfaceArgs)
			r, err := LookupAssociationsInterface(ctx, &args, opts...)
			var s LookupAssociationsInterfaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAssociationsInterfaceResultOutput)
}

type LookupAssociationsInterfaceOutputArgs struct {
	AssociationName       pulumi.StringInput `pulumi:"associationName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	TrafficControllerName pulumi.StringInput `pulumi:"trafficControllerName"`
}

func (LookupAssociationsInterfaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssociationsInterfaceArgs)(nil)).Elem()
}


type LookupAssociationsInterfaceResultOutput struct{ *pulumi.OutputState }

func (LookupAssociationsInterfaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssociationsInterfaceResult)(nil)).Elem()
}

func (o LookupAssociationsInterfaceResultOutput) ToLookupAssociationsInterfaceResultOutput() LookupAssociationsInterfaceResultOutput {
	return o
}

func (o LookupAssociationsInterfaceResultOutput) ToLookupAssociationsInterfaceResultOutputWithContext(ctx context.Context) LookupAssociationsInterfaceResultOutput {
	return o
}

func (o LookupAssociationsInterfaceResultOutput) AssociationType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssociationsInterfaceResult) string { return v.AssociationType }).(pulumi.StringOutput)
}

func (o LookupAssociationsInterfaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssociationsInterfaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAssociationsInterfaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssociationsInterfaceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAssociationsInterfaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssociationsInterfaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAssociationsInterfaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssociationsInterfaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAssociationsInterfaceResultOutput) Subnet() AssociationSubnetResponsePtrOutput {
	return o.ApplyT(func(v LookupAssociationsInterfaceResult) *AssociationSubnetResponse { return v.Subnet }).(AssociationSubnetResponsePtrOutput)
}

func (o LookupAssociationsInterfaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAssociationsInterfaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAssociationsInterfaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAssociationsInterfaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAssociationsInterfaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssociationsInterfaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAssociationsInterfaceResultOutput{})
}
