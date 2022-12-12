


package v20221101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataNetwork(ctx *pulumi.Context, args *LookupDataNetworkArgs, opts ...pulumi.InvokeOption) (*LookupDataNetworkResult, error) {
	var rv LookupDataNetworkResult
	err := ctx.Invoke("azure-native:mobilenetwork/v20221101:getDataNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataNetworkArgs struct {
	DataNetworkName   string `pulumi:"dataNetworkName"`
	MobileNetworkName string `pulumi:"mobileNetworkName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDataNetworkResult struct {
	Description       *string            `pulumi:"description"`
	Id                string             `pulumi:"id"`
	Location          string             `pulumi:"location"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Tags              map[string]string  `pulumi:"tags"`
	Type              string             `pulumi:"type"`
}

func LookupDataNetworkOutput(ctx *pulumi.Context, args LookupDataNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupDataNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataNetworkResult, error) {
			args := v.(LookupDataNetworkArgs)
			r, err := LookupDataNetwork(ctx, &args, opts...)
			var s LookupDataNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataNetworkResultOutput)
}

type LookupDataNetworkOutputArgs struct {
	DataNetworkName   pulumi.StringInput `pulumi:"dataNetworkName"`
	MobileNetworkName pulumi.StringInput `pulumi:"mobileNetworkName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDataNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataNetworkArgs)(nil)).Elem()
}


type LookupDataNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupDataNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataNetworkResult)(nil)).Elem()
}

func (o LookupDataNetworkResultOutput) ToLookupDataNetworkResultOutput() LookupDataNetworkResultOutput {
	return o
}

func (o LookupDataNetworkResultOutput) ToLookupDataNetworkResultOutputWithContext(ctx context.Context) LookupDataNetworkResultOutput {
	return o
}

func (o LookupDataNetworkResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataNetworkResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupDataNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataNetworkResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataNetworkResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDataNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDataNetworkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataNetworkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDataNetworkResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDataNetworkResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDataNetworkResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDataNetworkResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDataNetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataNetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataNetworkResultOutput{})
}
