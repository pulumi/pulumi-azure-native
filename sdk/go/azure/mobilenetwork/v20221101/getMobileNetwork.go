


package v20221101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMobileNetwork(ctx *pulumi.Context, args *LookupMobileNetworkArgs, opts ...pulumi.InvokeOption) (*LookupMobileNetworkResult, error) {
	var rv LookupMobileNetworkResult
	err := ctx.Invoke("azure-native:mobilenetwork/v20221101:getMobileNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMobileNetworkArgs struct {
	MobileNetworkName string `pulumi:"mobileNetworkName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupMobileNetworkResult struct {
	Id                                string             `pulumi:"id"`
	Location                          string             `pulumi:"location"`
	Name                              string             `pulumi:"name"`
	ProvisioningState                 string             `pulumi:"provisioningState"`
	PublicLandMobileNetworkIdentifier PlmnIdResponse     `pulumi:"publicLandMobileNetworkIdentifier"`
	ServiceKey                        string             `pulumi:"serviceKey"`
	SystemData                        SystemDataResponse `pulumi:"systemData"`
	Tags                              map[string]string  `pulumi:"tags"`
	Type                              string             `pulumi:"type"`
}

func LookupMobileNetworkOutput(ctx *pulumi.Context, args LookupMobileNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupMobileNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMobileNetworkResult, error) {
			args := v.(LookupMobileNetworkArgs)
			r, err := LookupMobileNetwork(ctx, &args, opts...)
			var s LookupMobileNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMobileNetworkResultOutput)
}

type LookupMobileNetworkOutputArgs struct {
	MobileNetworkName pulumi.StringInput `pulumi:"mobileNetworkName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMobileNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMobileNetworkArgs)(nil)).Elem()
}


type LookupMobileNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupMobileNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMobileNetworkResult)(nil)).Elem()
}

func (o LookupMobileNetworkResultOutput) ToLookupMobileNetworkResultOutput() LookupMobileNetworkResultOutput {
	return o
}

func (o LookupMobileNetworkResultOutput) ToLookupMobileNetworkResultOutputWithContext(ctx context.Context) LookupMobileNetworkResultOutput {
	return o
}

func (o LookupMobileNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobileNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMobileNetworkResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobileNetworkResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupMobileNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobileNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMobileNetworkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobileNetworkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupMobileNetworkResultOutput) PublicLandMobileNetworkIdentifier() PlmnIdResponseOutput {
	return o.ApplyT(func(v LookupMobileNetworkResult) PlmnIdResponse { return v.PublicLandMobileNetworkIdentifier }).(PlmnIdResponseOutput)
}

func (o LookupMobileNetworkResultOutput) ServiceKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobileNetworkResult) string { return v.ServiceKey }).(pulumi.StringOutput)
}

func (o LookupMobileNetworkResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMobileNetworkResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMobileNetworkResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMobileNetworkResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMobileNetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobileNetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMobileNetworkResultOutput{})
}
