


package devcenter

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AttachedNetworkByDevCenter struct {
	pulumi.CustomResourceState

	DomainJoinType            pulumi.StringOutput      `pulumi:"domainJoinType"`
	HealthCheckStatus         pulumi.StringOutput      `pulumi:"healthCheckStatus"`
	Name                      pulumi.StringOutput      `pulumi:"name"`
	NetworkConnectionId       pulumi.StringOutput      `pulumi:"networkConnectionId"`
	NetworkConnectionLocation pulumi.StringOutput      `pulumi:"networkConnectionLocation"`
	ProvisioningState         pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData                SystemDataResponseOutput `pulumi:"systemData"`
	Type                      pulumi.StringOutput      `pulumi:"type"`
}


func NewAttachedNetworkByDevCenter(ctx *pulumi.Context,
	name string, args *AttachedNetworkByDevCenterArgs, opts ...pulumi.ResourceOption) (*AttachedNetworkByDevCenter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DevCenterName == nil {
		return nil, errors.New("invalid value for required argument 'DevCenterName'")
	}
	if args.NetworkConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkConnectionId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter/v20220801preview:AttachedNetworkByDevCenter"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220901preview:AttachedNetworkByDevCenter"),
		},
	})
	opts = append(opts, aliases)
	var resource AttachedNetworkByDevCenter
	err := ctx.RegisterResource("azure-native:devcenter:AttachedNetworkByDevCenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAttachedNetworkByDevCenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttachedNetworkByDevCenterState, opts ...pulumi.ResourceOption) (*AttachedNetworkByDevCenter, error) {
	var resource AttachedNetworkByDevCenter
	err := ctx.ReadResource("azure-native:devcenter:AttachedNetworkByDevCenter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type attachedNetworkByDevCenterState struct {
}

type AttachedNetworkByDevCenterState struct {
}

func (AttachedNetworkByDevCenterState) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedNetworkByDevCenterState)(nil)).Elem()
}

type attachedNetworkByDevCenterArgs struct {
	AttachedNetworkConnectionName *string `pulumi:"attachedNetworkConnectionName"`
	DevCenterName                 string  `pulumi:"devCenterName"`
	NetworkConnectionId           string  `pulumi:"networkConnectionId"`
	ResourceGroupName             string  `pulumi:"resourceGroupName"`
}


type AttachedNetworkByDevCenterArgs struct {
	AttachedNetworkConnectionName pulumi.StringPtrInput
	DevCenterName                 pulumi.StringInput
	NetworkConnectionId           pulumi.StringInput
	ResourceGroupName             pulumi.StringInput
}

func (AttachedNetworkByDevCenterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedNetworkByDevCenterArgs)(nil)).Elem()
}

type AttachedNetworkByDevCenterInput interface {
	pulumi.Input

	ToAttachedNetworkByDevCenterOutput() AttachedNetworkByDevCenterOutput
	ToAttachedNetworkByDevCenterOutputWithContext(ctx context.Context) AttachedNetworkByDevCenterOutput
}

func (*AttachedNetworkByDevCenter) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedNetworkByDevCenter)(nil)).Elem()
}

func (i *AttachedNetworkByDevCenter) ToAttachedNetworkByDevCenterOutput() AttachedNetworkByDevCenterOutput {
	return i.ToAttachedNetworkByDevCenterOutputWithContext(context.Background())
}

func (i *AttachedNetworkByDevCenter) ToAttachedNetworkByDevCenterOutputWithContext(ctx context.Context) AttachedNetworkByDevCenterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedNetworkByDevCenterOutput)
}

type AttachedNetworkByDevCenterOutput struct{ *pulumi.OutputState }

func (AttachedNetworkByDevCenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedNetworkByDevCenter)(nil)).Elem()
}

func (o AttachedNetworkByDevCenterOutput) ToAttachedNetworkByDevCenterOutput() AttachedNetworkByDevCenterOutput {
	return o
}

func (o AttachedNetworkByDevCenterOutput) ToAttachedNetworkByDevCenterOutputWithContext(ctx context.Context) AttachedNetworkByDevCenterOutput {
	return o
}

func (o AttachedNetworkByDevCenterOutput) DomainJoinType() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedNetworkByDevCenter) pulumi.StringOutput { return v.DomainJoinType }).(pulumi.StringOutput)
}

func (o AttachedNetworkByDevCenterOutput) HealthCheckStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedNetworkByDevCenter) pulumi.StringOutput { return v.HealthCheckStatus }).(pulumi.StringOutput)
}

func (o AttachedNetworkByDevCenterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedNetworkByDevCenter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AttachedNetworkByDevCenterOutput) NetworkConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedNetworkByDevCenter) pulumi.StringOutput { return v.NetworkConnectionId }).(pulumi.StringOutput)
}

func (o AttachedNetworkByDevCenterOutput) NetworkConnectionLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedNetworkByDevCenter) pulumi.StringOutput { return v.NetworkConnectionLocation }).(pulumi.StringOutput)
}

func (o AttachedNetworkByDevCenterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedNetworkByDevCenter) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AttachedNetworkByDevCenterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AttachedNetworkByDevCenter) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AttachedNetworkByDevCenterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedNetworkByDevCenter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AttachedNetworkByDevCenterOutput{})
}
