


package v20201101

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
	if args.DeliveryType == nil {
		args.DeliveryType = pulumi.StringPtr("NonScheduled")
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
			Type: pulumi.String("azure-native:databox/v20210301:Job"),
		},
		{
			Type: pulumi.String("azure-native:databox/v20210501:Job"),
		},
		{
			Type: pulumi.String("azure-native:databox/v20210801preview:Job"),
		},
	})
	opts = append(opts, aliases)
	var resource Job
	err := ctx.RegisterResource("azure-native:databox/v20201101:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("azure-native:databox/v20201101:Job", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*Job)(nil))
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

type JobOutput struct{ *pulumi.OutputState }

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Job)(nil))
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(JobOutput{})
}
