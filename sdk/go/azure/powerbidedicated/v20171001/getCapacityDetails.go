


package v20171001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupCapacityDetails(ctx *pulumi.Context, args *LookupCapacityDetailsArgs, opts ...pulumi.InvokeOption) (*LookupCapacityDetailsResult, error) {
	var rv LookupCapacityDetailsResult
	err := ctx.Invoke("azure-native:powerbidedicated/v20171001:getCapacityDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCapacityDetailsArgs struct {
	DedicatedCapacityName string `pulumi:"dedicatedCapacityName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupCapacityDetailsResult struct {
	Administration    *DedicatedCapacityAdministratorsResponse `pulumi:"administration"`
	FriendlyName      string                                   `pulumi:"friendlyName"`
	Id                string                                   `pulumi:"id"`
	Location          string                                   `pulumi:"location"`
	Mode              string                                   `pulumi:"mode"`
	Name              string                                   `pulumi:"name"`
	ProvisioningState string                                   `pulumi:"provisioningState"`
	Sku               ResourceSkuResponse                      `pulumi:"sku"`
	State             string                                   `pulumi:"state"`
	Tags              map[string]string                        `pulumi:"tags"`
	TenantId          string                                   `pulumi:"tenantId"`
	Type              string                                   `pulumi:"type"`
}

func LookupCapacityDetailsOutput(ctx *pulumi.Context, args LookupCapacityDetailsOutputArgs, opts ...pulumi.InvokeOption) LookupCapacityDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCapacityDetailsResult, error) {
			args := v.(LookupCapacityDetailsArgs)
			r, err := LookupCapacityDetails(ctx, &args, opts...)
			var s LookupCapacityDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCapacityDetailsResultOutput)
}

type LookupCapacityDetailsOutputArgs struct {
	DedicatedCapacityName pulumi.StringInput `pulumi:"dedicatedCapacityName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCapacityDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityDetailsArgs)(nil)).Elem()
}


type LookupCapacityDetailsResultOutput struct{ *pulumi.OutputState }

func (LookupCapacityDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityDetailsResult)(nil)).Elem()
}

func (o LookupCapacityDetailsResultOutput) ToLookupCapacityDetailsResultOutput() LookupCapacityDetailsResultOutput {
	return o
}

func (o LookupCapacityDetailsResultOutput) ToLookupCapacityDetailsResultOutputWithContext(ctx context.Context) LookupCapacityDetailsResultOutput {
	return o
}

func (o LookupCapacityDetailsResultOutput) Administration() DedicatedCapacityAdministratorsResponsePtrOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) *DedicatedCapacityAdministratorsResponse { return v.Administration }).(DedicatedCapacityAdministratorsResponsePtrOutput)
}

func (o LookupCapacityDetailsResultOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) string { return v.FriendlyName }).(pulumi.StringOutput)
}

func (o LookupCapacityDetailsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCapacityDetailsResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupCapacityDetailsResultOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) string { return v.Mode }).(pulumi.StringOutput)
}

func (o LookupCapacityDetailsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCapacityDetailsResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupCapacityDetailsResultOutput) Sku() ResourceSkuResponseOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) ResourceSkuResponse { return v.Sku }).(ResourceSkuResponseOutput)
}

func (o LookupCapacityDetailsResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupCapacityDetailsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCapacityDetailsResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupCapacityDetailsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCapacityDetailsResultOutput{})
}
