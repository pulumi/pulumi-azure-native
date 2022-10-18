


package v20210801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Job struct {
	pulumi.CustomResourceState

	CancellationReason        pulumi.StringOutput               `pulumi:"cancellationReason"`
	DeliveryInfo              JobDeliveryInfoResponsePtrOutput  `pulumi:"deliveryInfo"`
	DeliveryType              pulumi.StringPtrOutput            `pulumi:"deliveryType"`
	Details                   pulumi.AnyOutput                  `pulumi:"details"`
	Error                     CloudErrorResponseOutput          `pulumi:"error"`
	Identity                  ResourceIdentityResponsePtrOutput `pulumi:"identity"`
	IsCancellable             pulumi.BoolOutput                 `pulumi:"isCancellable"`
	IsCancellableWithoutFee   pulumi.BoolOutput                 `pulumi:"isCancellableWithoutFee"`
	IsDeletable               pulumi.BoolOutput                 `pulumi:"isDeletable"`
	IsPrepareToShipEnabled    pulumi.BoolOutput                 `pulumi:"isPrepareToShipEnabled"`
	IsShippingAddressEditable pulumi.BoolOutput                 `pulumi:"isShippingAddressEditable"`
	Location                  pulumi.StringOutput               `pulumi:"location"`
	Name                      pulumi.StringOutput               `pulumi:"name"`
	Sku                       SkuResponseOutput                 `pulumi:"sku"`
	StartTime                 pulumi.StringOutput               `pulumi:"startTime"`
	Status                    pulumi.StringOutput               `pulumi:"status"`
	SystemData                SystemDataResponseOutput          `pulumi:"systemData"`
	Tags                      pulumi.StringMapOutput            `pulumi:"tags"`
	TransferType              pulumi.StringOutput               `pulumi:"transferType"`
	Type                      pulumi.StringOutput               `pulumi:"type"`
}


func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.TransferType == nil {
		return nil, errors.New("invalid value for required argument 'TransferType'")
	}
	if isZero(args.DeliveryType) {
		args.DeliveryType = pulumi.StringPtr("NonScheduled")
	}
	if args.Identity != nil {
		args.Identity = args.Identity.ToResourceIdentityPtrOutput().ApplyT(func(v *ResourceIdentity) *ResourceIdentity { return v.Defaults() }).(ResourceIdentityPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databox:Job"),
		},
		{
			Type: pulumi.String("azure-native:databox/v20180101:Job"),
		},
		{
			Type: pulumi.String("azure-native:databox/v20190901:Job"),
		},
		{
			Type: pulumi.String("azure-native:databox/v20200401:Job"),
		},
		{
			Type: pulumi.String("azure-native:databox/v20201101:Job"),
		},
		{
			Type: pulumi.String("azure-native:databox/v20210301:Job"),
		},
		{
			Type: pulumi.String("azure-native:databox/v20210501:Job"),
		},
		{
			Type: pulumi.String("azure-native:databox/v20211201:Job"),
		},
		{
			Type: pulumi.String("azure-native:databox/v20220201:Job"),
		},
	})
	opts = append(opts, aliases)
	var resource Job
	err := ctx.RegisterResource("azure-native:databox/v20210801preview:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("azure-native:databox/v20210801preview:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type jobState struct {
}

type JobState struct {
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	DeliveryInfo      *JobDeliveryInfo  `pulumi:"deliveryInfo"`
	DeliveryType      *string           `pulumi:"deliveryType"`
	Details           interface{}       `pulumi:"details"`
	Identity          *ResourceIdentity `pulumi:"identity"`
	JobName           *string           `pulumi:"jobName"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Sku               Sku               `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
	TransferType      string            `pulumi:"transferType"`
}


type JobArgs struct {
	DeliveryInfo      JobDeliveryInfoPtrInput
	DeliveryType      pulumi.StringPtrInput
	Details           pulumi.Input
	Identity          ResourceIdentityPtrInput
	JobName           pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               SkuInput
	Tags              pulumi.StringMapInput
	TransferType      pulumi.StringInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

type JobInput interface {
	pulumi.Input

	ToJobOutput() JobOutput
	ToJobOutputWithContext(ctx context.Context) JobOutput
}

func (*Job) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

type JobOutput struct{ *pulumi.OutputState }

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

func (o JobOutput) CancellationReason() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.CancellationReason }).(pulumi.StringOutput)
}

func (o JobOutput) DeliveryInfo() JobDeliveryInfoResponsePtrOutput {
	return o.ApplyT(func(v *Job) JobDeliveryInfoResponsePtrOutput { return v.DeliveryInfo }).(JobDeliveryInfoResponsePtrOutput)
}

func (o JobOutput) DeliveryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.DeliveryType }).(pulumi.StringPtrOutput)
}

func (o JobOutput) Details() pulumi.AnyOutput {
	return o.ApplyT(func(v *Job) pulumi.AnyOutput { return v.Details }).(pulumi.AnyOutput)
}

func (o JobOutput) Error() CloudErrorResponseOutput {
	return o.ApplyT(func(v *Job) CloudErrorResponseOutput { return v.Error }).(CloudErrorResponseOutput)
}

func (o JobOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Job) ResourceIdentityResponsePtrOutput { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o JobOutput) IsCancellable() pulumi.BoolOutput {
	return o.ApplyT(func(v *Job) pulumi.BoolOutput { return v.IsCancellable }).(pulumi.BoolOutput)
}

func (o JobOutput) IsCancellableWithoutFee() pulumi.BoolOutput {
	return o.ApplyT(func(v *Job) pulumi.BoolOutput { return v.IsCancellableWithoutFee }).(pulumi.BoolOutput)
}

func (o JobOutput) IsDeletable() pulumi.BoolOutput {
	return o.ApplyT(func(v *Job) pulumi.BoolOutput { return v.IsDeletable }).(pulumi.BoolOutput)
}

func (o JobOutput) IsPrepareToShipEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Job) pulumi.BoolOutput { return v.IsPrepareToShipEnabled }).(pulumi.BoolOutput)
}

func (o JobOutput) IsShippingAddressEditable() pulumi.BoolOutput {
	return o.ApplyT(func(v *Job) pulumi.BoolOutput { return v.IsShippingAddressEditable }).(pulumi.BoolOutput)
}

func (o JobOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o JobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o JobOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *Job) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

func (o JobOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

func (o JobOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o JobOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Job) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o JobOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Job) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o JobOutput) TransferType() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.TransferType }).(pulumi.StringOutput)
}

func (o JobOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(JobOutput{})
}
