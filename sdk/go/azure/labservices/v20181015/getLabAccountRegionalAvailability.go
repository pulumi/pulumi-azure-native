


package v20181015

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetLabAccountRegionalAvailability(ctx *pulumi.Context, args *GetLabAccountRegionalAvailabilityArgs, opts ...pulumi.InvokeOption) (*GetLabAccountRegionalAvailabilityResult, error) {
	var rv GetLabAccountRegionalAvailabilityResult
	err := ctx.Invoke("azure-native:labservices/v20181015:getLabAccountRegionalAvailability", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetLabAccountRegionalAvailabilityArgs struct {
	LabAccountName    string `pulumi:"labAccountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetLabAccountRegionalAvailabilityResult struct {
	RegionalAvailability []RegionalAvailabilityResponse `pulumi:"regionalAvailability"`
}

func GetLabAccountRegionalAvailabilityOutput(ctx *pulumi.Context, args GetLabAccountRegionalAvailabilityOutputArgs, opts ...pulumi.InvokeOption) GetLabAccountRegionalAvailabilityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLabAccountRegionalAvailabilityResult, error) {
			args := v.(GetLabAccountRegionalAvailabilityArgs)
			r, err := GetLabAccountRegionalAvailability(ctx, &args, opts...)
			var s GetLabAccountRegionalAvailabilityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLabAccountRegionalAvailabilityResultOutput)
}

type GetLabAccountRegionalAvailabilityOutputArgs struct {
	LabAccountName    pulumi.StringInput `pulumi:"labAccountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetLabAccountRegionalAvailabilityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLabAccountRegionalAvailabilityArgs)(nil)).Elem()
}


type GetLabAccountRegionalAvailabilityResultOutput struct{ *pulumi.OutputState }

func (GetLabAccountRegionalAvailabilityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLabAccountRegionalAvailabilityResult)(nil)).Elem()
}

func (o GetLabAccountRegionalAvailabilityResultOutput) ToGetLabAccountRegionalAvailabilityResultOutput() GetLabAccountRegionalAvailabilityResultOutput {
	return o
}

func (o GetLabAccountRegionalAvailabilityResultOutput) ToGetLabAccountRegionalAvailabilityResultOutputWithContext(ctx context.Context) GetLabAccountRegionalAvailabilityResultOutput {
	return o
}

func (o GetLabAccountRegionalAvailabilityResultOutput) RegionalAvailability() RegionalAvailabilityResponseArrayOutput {
	return o.ApplyT(func(v GetLabAccountRegionalAvailabilityResult) []RegionalAvailabilityResponse {
		return v.RegionalAvailability
	}).(RegionalAvailabilityResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLabAccountRegionalAvailabilityResultOutput{})
}
