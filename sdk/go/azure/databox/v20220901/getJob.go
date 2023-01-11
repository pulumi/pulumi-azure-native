


package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("azure-native:databox/v20220901:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupJobArgs struct {
	Expand            *string `pulumi:"expand"`
	JobName           string  `pulumi:"jobName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupJobResult struct {
	CancellationReason               string                    `pulumi:"cancellationReason"`
	DeliveryInfo                     *JobDeliveryInfoResponse  `pulumi:"deliveryInfo"`
	DeliveryType                     *string                   `pulumi:"deliveryType"`
	Details                          interface{}               `pulumi:"details"`
	Error                            CloudErrorResponse        `pulumi:"error"`
	Id                               string                    `pulumi:"id"`
	Identity                         *ResourceIdentityResponse `pulumi:"identity"`
	IsCancellable                    bool                      `pulumi:"isCancellable"`
	IsCancellableWithoutFee          bool                      `pulumi:"isCancellableWithoutFee"`
	IsDeletable                      bool                      `pulumi:"isDeletable"`
	IsPrepareToShipEnabled           bool                      `pulumi:"isPrepareToShipEnabled"`
	IsShippingAddressEditable        bool                      `pulumi:"isShippingAddressEditable"`
	Location                         string                    `pulumi:"location"`
	Name                             string                    `pulumi:"name"`
	ReverseShippingDetailsUpdate     string                    `pulumi:"reverseShippingDetailsUpdate"`
	ReverseTransportPreferenceUpdate string                    `pulumi:"reverseTransportPreferenceUpdate"`
	Sku                              SkuResponse               `pulumi:"sku"`
	StartTime                        string                    `pulumi:"startTime"`
	Status                           string                    `pulumi:"status"`
	SystemData                       SystemDataResponse        `pulumi:"systemData"`
	Tags                             map[string]string         `pulumi:"tags"`
	TransferType                     string                    `pulumi:"transferType"`
	Type                             string                    `pulumi:"type"`
}


func (val *LookupJobResult) Defaults() *LookupJobResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DeliveryType) {
		deliveryType_ := "NonScheduled"
		tmp.DeliveryType = &deliveryType_
	}
	tmp.Identity = tmp.Identity.Defaults()

	return &tmp
}

func LookupJobOutput(ctx *pulumi.Context, args LookupJobOutputArgs, opts ...pulumi.InvokeOption) LookupJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobResult, error) {
			args := v.(LookupJobArgs)
			r, err := LookupJob(ctx, &args, opts...)
			var s LookupJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobResultOutput)
}

type LookupJobOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	JobName           pulumi.StringInput    `pulumi:"jobName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobArgs)(nil)).Elem()
}


type LookupJobResultOutput struct{ *pulumi.OutputState }

func (LookupJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobResult)(nil)).Elem()
}

func (o LookupJobResultOutput) ToLookupJobResultOutput() LookupJobResultOutput {
	return o
}

func (o LookupJobResultOutput) ToLookupJobResultOutputWithContext(ctx context.Context) LookupJobResultOutput {
	return o
}

func (o LookupJobResultOutput) CancellationReason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.CancellationReason }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) DeliveryInfo() JobDeliveryInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupJobResult) *JobDeliveryInfoResponse { return v.DeliveryInfo }).(JobDeliveryInfoResponsePtrOutput)
}

func (o LookupJobResultOutput) DeliveryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *string { return v.DeliveryType }).(pulumi.StringPtrOutput)
}

func (o LookupJobResultOutput) Details() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupJobResult) interface{} { return v.Details }).(pulumi.AnyOutput)
}

func (o LookupJobResultOutput) Error() CloudErrorResponseOutput {
	return o.ApplyT(func(v LookupJobResult) CloudErrorResponse { return v.Error }).(CloudErrorResponseOutput)
}

func (o LookupJobResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupJobResult) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o LookupJobResultOutput) IsCancellable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupJobResult) bool { return v.IsCancellable }).(pulumi.BoolOutput)
}

func (o LookupJobResultOutput) IsCancellableWithoutFee() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupJobResult) bool { return v.IsCancellableWithoutFee }).(pulumi.BoolOutput)
}

func (o LookupJobResultOutput) IsDeletable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupJobResult) bool { return v.IsDeletable }).(pulumi.BoolOutput)
}

func (o LookupJobResultOutput) IsPrepareToShipEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupJobResult) bool { return v.IsPrepareToShipEnabled }).(pulumi.BoolOutput)
}

func (o LookupJobResultOutput) IsShippingAddressEditable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupJobResult) bool { return v.IsShippingAddressEditable }).(pulumi.BoolOutput)
}

func (o LookupJobResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) ReverseShippingDetailsUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.ReverseShippingDetailsUpdate }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) ReverseTransportPreferenceUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.ReverseTransportPreferenceUpdate }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupJobResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupJobResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupJobResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupJobResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupJobResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupJobResultOutput) TransferType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.TransferType }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobResultOutput{})
}
