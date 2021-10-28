


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Authorization struct {
	pulumi.CustomResourceState

	ExpressRouteAuthorizationId  pulumi.StringOutput `pulumi:"expressRouteAuthorizationId"`
	ExpressRouteAuthorizationKey pulumi.StringOutput `pulumi:"expressRouteAuthorizationKey"`
	Name                         pulumi.StringOutput `pulumi:"name"`
	ProvisioningState            pulumi.StringOutput `pulumi:"provisioningState"`
	Type                         pulumi.StringOutput `pulumi:"type"`
}


func NewAuthorization(ctx *pulumi.Context,
	name string, args *AuthorizationArgs, opts ...pulumi.ResourceOption) (*Authorization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:avs/v20210101preview:Authorization"),
		},
		{
			Type: pulumi.String("azure-native:avs:Authorization"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs:Authorization"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200320:Authorization"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20200320:Authorization"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:Authorization"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20200717preview:Authorization"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:Authorization"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20210601:Authorization"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:Authorization"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20211201:Authorization"),
		},
	})
	opts = append(opts, aliases)
	var resource Authorization
	err := ctx.RegisterResource("azure-native:avs/v20210101preview:Authorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorizationState, opts ...pulumi.ResourceOption) (*Authorization, error) {
	var resource Authorization
	err := ctx.ReadResource("azure-native:avs/v20210101preview:Authorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type authorizationState struct {
}

type AuthorizationState struct {
}

func (AuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationState)(nil)).Elem()
}

type authorizationArgs struct {
	AuthorizationName *string `pulumi:"authorizationName"`
	PrivateCloudName  string  `pulumi:"privateCloudName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type AuthorizationArgs struct {
	AuthorizationName pulumi.StringPtrInput
	PrivateCloudName  pulumi.StringInput
	ResourceGroupName pulumi.StringInput
}

func (AuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationArgs)(nil)).Elem()
}

type AuthorizationInput interface {
	pulumi.Input

	ToAuthorizationOutput() AuthorizationOutput
	ToAuthorizationOutputWithContext(ctx context.Context) AuthorizationOutput
}

func (*Authorization) ElementType() reflect.Type {
	return reflect.TypeOf((*Authorization)(nil))
}

func (i *Authorization) ToAuthorizationOutput() AuthorizationOutput {
	return i.ToAuthorizationOutputWithContext(context.Background())
}

func (i *Authorization) ToAuthorizationOutputWithContext(ctx context.Context) AuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationOutput)
}

type AuthorizationOutput struct{ *pulumi.OutputState }

func (AuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Authorization)(nil))
}

func (o AuthorizationOutput) ToAuthorizationOutput() AuthorizationOutput {
	return o
}

func (o AuthorizationOutput) ToAuthorizationOutputWithContext(ctx context.Context) AuthorizationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AuthorizationOutput{})
}
