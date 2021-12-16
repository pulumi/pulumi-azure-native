


package v20181120

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GuestConfigurationAssignment struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                               `pulumi:"location"`
	Name       pulumi.StringPtrOutput                               `pulumi:"name"`
	Properties GuestConfigurationAssignmentPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                                  `pulumi:"type"`
}


func NewGuestConfigurationAssignment(ctx *pulumi.Context,
	name string, args *GuestConfigurationAssignmentArgs, opts ...pulumi.ResourceOption) (*GuestConfigurationAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VmName == nil {
		return nil, errors.New("invalid value for required argument 'VmName'")
	}
	propertiesApplier := func(v GuestConfigurationAssignmentProperties) *GuestConfigurationAssignmentProperties {
		return v.Defaults()
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToGuestConfigurationAssignmentPropertiesPtrOutput().Elem().ApplyT(propertiesApplier).(GuestConfigurationAssignmentPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:guestconfiguration:GuestConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:guestconfiguration/v20180630preview:GuestConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:guestconfiguration/v20200625:GuestConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:guestconfiguration/v20210125:GuestConfigurationAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource GuestConfigurationAssignment
	err := ctx.RegisterResource("azure-native:guestconfiguration/v20181120:GuestConfigurationAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGuestConfigurationAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GuestConfigurationAssignmentState, opts ...pulumi.ResourceOption) (*GuestConfigurationAssignment, error) {
	var resource GuestConfigurationAssignment
	err := ctx.ReadResource("azure-native:guestconfiguration/v20181120:GuestConfigurationAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type guestConfigurationAssignmentState struct {
}

type GuestConfigurationAssignmentState struct {
}

func (GuestConfigurationAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*guestConfigurationAssignmentState)(nil)).Elem()
}

type guestConfigurationAssignmentArgs struct {
	GuestConfigurationAssignmentName *string                                 `pulumi:"guestConfigurationAssignmentName"`
	Location                         *string                                 `pulumi:"location"`
	Name                             *string                                 `pulumi:"name"`
	Properties                       *GuestConfigurationAssignmentProperties `pulumi:"properties"`
	ResourceGroupName                string                                  `pulumi:"resourceGroupName"`
	VmName                           string                                  `pulumi:"vmName"`
}


type GuestConfigurationAssignmentArgs struct {
	GuestConfigurationAssignmentName pulumi.StringPtrInput
	Location                         pulumi.StringPtrInput
	Name                             pulumi.StringPtrInput
	Properties                       GuestConfigurationAssignmentPropertiesPtrInput
	ResourceGroupName                pulumi.StringInput
	VmName                           pulumi.StringInput
}

func (GuestConfigurationAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*guestConfigurationAssignmentArgs)(nil)).Elem()
}

type GuestConfigurationAssignmentInput interface {
	pulumi.Input

	ToGuestConfigurationAssignmentOutput() GuestConfigurationAssignmentOutput
	ToGuestConfigurationAssignmentOutputWithContext(ctx context.Context) GuestConfigurationAssignmentOutput
}

func (*GuestConfigurationAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationAssignment)(nil)).Elem()
}

func (i *GuestConfigurationAssignment) ToGuestConfigurationAssignmentOutput() GuestConfigurationAssignmentOutput {
	return i.ToGuestConfigurationAssignmentOutputWithContext(context.Background())
}

func (i *GuestConfigurationAssignment) ToGuestConfigurationAssignmentOutputWithContext(ctx context.Context) GuestConfigurationAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationAssignmentOutput)
}

type GuestConfigurationAssignmentOutput struct{ *pulumi.OutputState }

func (GuestConfigurationAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationAssignment)(nil)).Elem()
}

func (o GuestConfigurationAssignmentOutput) ToGuestConfigurationAssignmentOutput() GuestConfigurationAssignmentOutput {
	return o
}

func (o GuestConfigurationAssignmentOutput) ToGuestConfigurationAssignmentOutputWithContext(ctx context.Context) GuestConfigurationAssignmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GuestConfigurationAssignmentOutput{})
}
