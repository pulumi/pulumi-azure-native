


package v20201015preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPartnerNamespace(ctx *pulumi.Context, args *LookupPartnerNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupPartnerNamespaceResult, error) {
	var rv LookupPartnerNamespaceResult
	err := ctx.Invoke("azure-native:eventgrid/v20201015preview:getPartnerNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPartnerNamespaceArgs struct {
	PartnerNamespaceName string `pulumi:"partnerNamespaceName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupPartnerNamespaceResult struct {
	Endpoint                            string             `pulumi:"endpoint"`
	Id                                  string             `pulumi:"id"`
	Location                            string             `pulumi:"location"`
	Name                                string             `pulumi:"name"`
	PartnerRegistrationFullyQualifiedId *string            `pulumi:"partnerRegistrationFullyQualifiedId"`
	ProvisioningState                   string             `pulumi:"provisioningState"`
	SystemData                          SystemDataResponse `pulumi:"systemData"`
	Tags                                map[string]string  `pulumi:"tags"`
	Type                                string             `pulumi:"type"`
}

func LookupPartnerNamespaceOutput(ctx *pulumi.Context, args LookupPartnerNamespaceOutputArgs, opts ...pulumi.InvokeOption) LookupPartnerNamespaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPartnerNamespaceResult, error) {
			args := v.(LookupPartnerNamespaceArgs)
			r, err := LookupPartnerNamespace(ctx, &args, opts...)
			var s LookupPartnerNamespaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPartnerNamespaceResultOutput)
}

type LookupPartnerNamespaceOutputArgs struct {
	PartnerNamespaceName pulumi.StringInput `pulumi:"partnerNamespaceName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPartnerNamespaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerNamespaceArgs)(nil)).Elem()
}


type LookupPartnerNamespaceResultOutput struct{ *pulumi.OutputState }

func (LookupPartnerNamespaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerNamespaceResult)(nil)).Elem()
}

func (o LookupPartnerNamespaceResultOutput) ToLookupPartnerNamespaceResultOutput() LookupPartnerNamespaceResultOutput {
	return o
}

func (o LookupPartnerNamespaceResultOutput) ToLookupPartnerNamespaceResultOutputWithContext(ctx context.Context) LookupPartnerNamespaceResultOutput {
	return o
}

func (o LookupPartnerNamespaceResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o LookupPartnerNamespaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPartnerNamespaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPartnerNamespaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPartnerNamespaceResultOutput) PartnerRegistrationFullyQualifiedId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) *string { return v.PartnerRegistrationFullyQualifiedId }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerNamespaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPartnerNamespaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPartnerNamespaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPartnerNamespaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPartnerNamespaceResultOutput{})
}
