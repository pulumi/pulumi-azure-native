


package mobilenetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSimPolicy(ctx *pulumi.Context, args *LookupSimPolicyArgs, opts ...pulumi.InvokeOption) (*LookupSimPolicyResult, error) {
	var rv LookupSimPolicyResult
	err := ctx.Invoke("azure-native:mobilenetwork:getSimPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSimPolicyArgs struct {
	MobileNetworkName string `pulumi:"mobileNetworkName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SimPolicyName     string `pulumi:"simPolicyName"`
}


type LookupSimPolicyResult struct {
	CreatedAt           *string                      `pulumi:"createdAt"`
	CreatedBy           *string                      `pulumi:"createdBy"`
	CreatedByType       *string                      `pulumi:"createdByType"`
	DefaultSlice        SliceResourceIdResponse      `pulumi:"defaultSlice"`
	Id                  string                       `pulumi:"id"`
	LastModifiedAt      *string                      `pulumi:"lastModifiedAt"`
	LastModifiedBy      *string                      `pulumi:"lastModifiedBy"`
	LastModifiedByType  *string                      `pulumi:"lastModifiedByType"`
	Location            string                       `pulumi:"location"`
	Name                string                       `pulumi:"name"`
	ProvisioningState   string                       `pulumi:"provisioningState"`
	RegistrationTimer   *int                         `pulumi:"registrationTimer"`
	RfspIndex           *int                         `pulumi:"rfspIndex"`
	SliceConfigurations []SliceConfigurationResponse `pulumi:"sliceConfigurations"`
	SystemData          SystemDataResponse           `pulumi:"systemData"`
	Tags                map[string]string            `pulumi:"tags"`
	Type                string                       `pulumi:"type"`
	UeAmbr              AmbrResponse                 `pulumi:"ueAmbr"`
}


func (val *LookupSimPolicyResult) Defaults() *LookupSimPolicyResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.RegistrationTimer) {
		registrationTimer_ := 3240
		tmp.RegistrationTimer = &registrationTimer_
	}
	return &tmp
}

func LookupSimPolicyOutput(ctx *pulumi.Context, args LookupSimPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupSimPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSimPolicyResult, error) {
			args := v.(LookupSimPolicyArgs)
			r, err := LookupSimPolicy(ctx, &args, opts...)
			var s LookupSimPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSimPolicyResultOutput)
}

type LookupSimPolicyOutputArgs struct {
	MobileNetworkName pulumi.StringInput `pulumi:"mobileNetworkName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SimPolicyName     pulumi.StringInput `pulumi:"simPolicyName"`
}

func (LookupSimPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSimPolicyArgs)(nil)).Elem()
}


type LookupSimPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupSimPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSimPolicyResult)(nil)).Elem()
}

func (o LookupSimPolicyResultOutput) ToLookupSimPolicyResultOutput() LookupSimPolicyResultOutput {
	return o
}

func (o LookupSimPolicyResultOutput) ToLookupSimPolicyResultOutputWithContext(ctx context.Context) LookupSimPolicyResultOutput {
	return o
}

func (o LookupSimPolicyResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupSimPolicyResultOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o LookupSimPolicyResultOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o LookupSimPolicyResultOutput) DefaultSlice() SliceResourceIdResponseOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) SliceResourceIdResponse { return v.DefaultSlice }).(SliceResourceIdResponseOutput)
}

func (o LookupSimPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSimPolicyResultOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o LookupSimPolicyResultOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o LookupSimPolicyResultOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o LookupSimPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSimPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSimPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSimPolicyResultOutput) RegistrationTimer() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) *int { return v.RegistrationTimer }).(pulumi.IntPtrOutput)
}

func (o LookupSimPolicyResultOutput) RfspIndex() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) *int { return v.RfspIndex }).(pulumi.IntPtrOutput)
}

func (o LookupSimPolicyResultOutput) SliceConfigurations() SliceConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) []SliceConfigurationResponse { return v.SliceConfigurations }).(SliceConfigurationResponseArrayOutput)
}

func (o LookupSimPolicyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSimPolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSimPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSimPolicyResultOutput) UeAmbr() AmbrResponseOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) AmbrResponse { return v.UeAmbr }).(AmbrResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSimPolicyResultOutput{})
}
