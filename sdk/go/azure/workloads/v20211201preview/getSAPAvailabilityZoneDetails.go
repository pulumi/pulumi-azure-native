


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSAPAvailabilityZoneDetails(ctx *pulumi.Context, args *GetSAPAvailabilityZoneDetailsArgs, opts ...pulumi.InvokeOption) (*GetSAPAvailabilityZoneDetailsResult, error) {
	var rv GetSAPAvailabilityZoneDetailsResult
	err := ctx.Invoke("azure-native:workloads/v20211201preview:getSAPAvailabilityZoneDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSAPAvailabilityZoneDetailsArgs struct {
	AppLocation  string `pulumi:"appLocation"`
	DatabaseType string `pulumi:"databaseType"`
	Location     string `pulumi:"location"`
	SapProduct   string `pulumi:"sapProduct"`
}


type GetSAPAvailabilityZoneDetailsResult struct {
	AvailabilityZonePairs []SAPAvailabilityZonePairResponse `pulumi:"availabilityZonePairs"`
}

func GetSAPAvailabilityZoneDetailsOutput(ctx *pulumi.Context, args GetSAPAvailabilityZoneDetailsOutputArgs, opts ...pulumi.InvokeOption) GetSAPAvailabilityZoneDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSAPAvailabilityZoneDetailsResult, error) {
			args := v.(GetSAPAvailabilityZoneDetailsArgs)
			r, err := GetSAPAvailabilityZoneDetails(ctx, &args, opts...)
			var s GetSAPAvailabilityZoneDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSAPAvailabilityZoneDetailsResultOutput)
}

type GetSAPAvailabilityZoneDetailsOutputArgs struct {
	AppLocation  pulumi.StringInput `pulumi:"appLocation"`
	DatabaseType pulumi.StringInput `pulumi:"databaseType"`
	Location     pulumi.StringInput `pulumi:"location"`
	SapProduct   pulumi.StringInput `pulumi:"sapProduct"`
}

func (GetSAPAvailabilityZoneDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSAPAvailabilityZoneDetailsArgs)(nil)).Elem()
}


type GetSAPAvailabilityZoneDetailsResultOutput struct{ *pulumi.OutputState }

func (GetSAPAvailabilityZoneDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSAPAvailabilityZoneDetailsResult)(nil)).Elem()
}

func (o GetSAPAvailabilityZoneDetailsResultOutput) ToGetSAPAvailabilityZoneDetailsResultOutput() GetSAPAvailabilityZoneDetailsResultOutput {
	return o
}

func (o GetSAPAvailabilityZoneDetailsResultOutput) ToGetSAPAvailabilityZoneDetailsResultOutputWithContext(ctx context.Context) GetSAPAvailabilityZoneDetailsResultOutput {
	return o
}

func (o GetSAPAvailabilityZoneDetailsResultOutput) AvailabilityZonePairs() SAPAvailabilityZonePairResponseArrayOutput {
	return o.ApplyT(func(v GetSAPAvailabilityZoneDetailsResult) []SAPAvailabilityZonePairResponse {
		return v.AvailabilityZonePairs
	}).(SAPAvailabilityZonePairResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSAPAvailabilityZoneDetailsResultOutput{})
}
