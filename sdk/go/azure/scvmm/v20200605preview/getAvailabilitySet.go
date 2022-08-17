


package v20200605preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAvailabilitySet(ctx *pulumi.Context, args *LookupAvailabilitySetArgs, opts ...pulumi.InvokeOption) (*LookupAvailabilitySetResult, error) {
	var rv LookupAvailabilitySetResult
	err := ctx.Invoke("azure-native:scvmm/v20200605preview:getAvailabilitySet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAvailabilitySetArgs struct {
	AvailabilitySetName string `pulumi:"availabilitySetName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupAvailabilitySetResult struct {
	AvailabilitySetName *string                   `pulumi:"availabilitySetName"`
	ExtendedLocation    *ExtendedLocationResponse `pulumi:"extendedLocation"`
	Id                  string                    `pulumi:"id"`
	Location            *string                   `pulumi:"location"`
	Name                string                    `pulumi:"name"`
	ProvisioningState   string                    `pulumi:"provisioningState"`
	SystemData          SystemDataResponse        `pulumi:"systemData"`
	Tags                map[string]string         `pulumi:"tags"`
	Type                string                    `pulumi:"type"`
	VmmServerId         *string                   `pulumi:"vmmServerId"`
}

func LookupAvailabilitySetOutput(ctx *pulumi.Context, args LookupAvailabilitySetOutputArgs, opts ...pulumi.InvokeOption) LookupAvailabilitySetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAvailabilitySetResult, error) {
			args := v.(LookupAvailabilitySetArgs)
			r, err := LookupAvailabilitySet(ctx, &args, opts...)
			var s LookupAvailabilitySetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAvailabilitySetResultOutput)
}

type LookupAvailabilitySetOutputArgs struct {
	AvailabilitySetName pulumi.StringInput `pulumi:"availabilitySetName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAvailabilitySetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAvailabilitySetArgs)(nil)).Elem()
}


type LookupAvailabilitySetResultOutput struct{ *pulumi.OutputState }

func (LookupAvailabilitySetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAvailabilitySetResult)(nil)).Elem()
}

func (o LookupAvailabilitySetResultOutput) ToLookupAvailabilitySetResultOutput() LookupAvailabilitySetResultOutput {
	return o
}

func (o LookupAvailabilitySetResultOutput) ToLookupAvailabilitySetResultOutputWithContext(ctx context.Context) LookupAvailabilitySetResultOutput {
	return o
}

func (o LookupAvailabilitySetResultOutput) AvailabilitySetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) *string { return v.AvailabilitySetName }).(pulumi.StringPtrOutput)
}

func (o LookupAvailabilitySetResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupAvailabilitySetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAvailabilitySetResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupAvailabilitySetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAvailabilitySetResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAvailabilitySetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAvailabilitySetResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAvailabilitySetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAvailabilitySetResultOutput) VmmServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) *string { return v.VmmServerId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAvailabilitySetResultOutput{})
}
