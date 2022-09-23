


package v20220101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExpressRoutePortAuthorization struct {
	pulumi.CustomResourceState

	AuthorizationKey       pulumi.StringOutput    `pulumi:"authorizationKey"`
	AuthorizationUseStatus pulumi.StringOutput    `pulumi:"authorizationUseStatus"`
	CircuitResourceUri     pulumi.StringOutput    `pulumi:"circuitResourceUri"`
	Etag                   pulumi.StringOutput    `pulumi:"etag"`
	Name                   pulumi.StringPtrOutput `pulumi:"name"`
	ProvisioningState      pulumi.StringOutput    `pulumi:"provisioningState"`
	Type                   pulumi.StringOutput    `pulumi:"type"`
}


func NewExpressRoutePortAuthorization(ctx *pulumi.Context,
	name string, args *ExpressRoutePortAuthorizationArgs, opts ...pulumi.ResourceOption) (*ExpressRoutePortAuthorization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExpressRoutePortName == nil {
		return nil, errors.New("invalid value for required argument 'ExpressRoutePortName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ExpressRoutePortAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ExpressRoutePortAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ExpressRoutePortAuthorization"),
		},
	})
	opts = append(opts, aliases)
	var resource ExpressRoutePortAuthorization
	err := ctx.RegisterResource("azure-native:network/v20220101:ExpressRoutePortAuthorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExpressRoutePortAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRoutePortAuthorizationState, opts ...pulumi.ResourceOption) (*ExpressRoutePortAuthorization, error) {
	var resource ExpressRoutePortAuthorization
	err := ctx.ReadResource("azure-native:network/v20220101:ExpressRoutePortAuthorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type expressRoutePortAuthorizationState struct {
}

type ExpressRoutePortAuthorizationState struct {
}

func (ExpressRoutePortAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRoutePortAuthorizationState)(nil)).Elem()
}

type expressRoutePortAuthorizationArgs struct {
	AuthorizationName    *string `pulumi:"authorizationName"`
	ExpressRoutePortName string  `pulumi:"expressRoutePortName"`
	Id                   *string `pulumi:"id"`
	Name                 *string `pulumi:"name"`
	ResourceGroupName    string  `pulumi:"resourceGroupName"`
}


type ExpressRoutePortAuthorizationArgs struct {
	AuthorizationName    pulumi.StringPtrInput
	ExpressRoutePortName pulumi.StringInput
	Id                   pulumi.StringPtrInput
	Name                 pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
}

func (ExpressRoutePortAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRoutePortAuthorizationArgs)(nil)).Elem()
}

type ExpressRoutePortAuthorizationInput interface {
	pulumi.Input

	ToExpressRoutePortAuthorizationOutput() ExpressRoutePortAuthorizationOutput
	ToExpressRoutePortAuthorizationOutputWithContext(ctx context.Context) ExpressRoutePortAuthorizationOutput
}

func (*ExpressRoutePortAuthorization) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRoutePortAuthorization)(nil)).Elem()
}

func (i *ExpressRoutePortAuthorization) ToExpressRoutePortAuthorizationOutput() ExpressRoutePortAuthorizationOutput {
	return i.ToExpressRoutePortAuthorizationOutputWithContext(context.Background())
}

func (i *ExpressRoutePortAuthorization) ToExpressRoutePortAuthorizationOutputWithContext(ctx context.Context) ExpressRoutePortAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRoutePortAuthorizationOutput)
}

type ExpressRoutePortAuthorizationOutput struct{ *pulumi.OutputState }

func (ExpressRoutePortAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRoutePortAuthorization)(nil)).Elem()
}

func (o ExpressRoutePortAuthorizationOutput) ToExpressRoutePortAuthorizationOutput() ExpressRoutePortAuthorizationOutput {
	return o
}

func (o ExpressRoutePortAuthorizationOutput) ToExpressRoutePortAuthorizationOutputWithContext(ctx context.Context) ExpressRoutePortAuthorizationOutput {
	return o
}

func (o ExpressRoutePortAuthorizationOutput) AuthorizationKey() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePortAuthorization) pulumi.StringOutput { return v.AuthorizationKey }).(pulumi.StringOutput)
}

func (o ExpressRoutePortAuthorizationOutput) AuthorizationUseStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePortAuthorization) pulumi.StringOutput { return v.AuthorizationUseStatus }).(pulumi.StringOutput)
}

func (o ExpressRoutePortAuthorizationOutput) CircuitResourceUri() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePortAuthorization) pulumi.StringOutput { return v.CircuitResourceUri }).(pulumi.StringOutput)
}

func (o ExpressRoutePortAuthorizationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePortAuthorization) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ExpressRoutePortAuthorizationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRoutePortAuthorization) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExpressRoutePortAuthorizationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePortAuthorization) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ExpressRoutePortAuthorizationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePortAuthorization) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ExpressRoutePortAuthorizationOutput{})
}
