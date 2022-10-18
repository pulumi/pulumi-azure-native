


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Authorization struct {
	pulumi.CustomResourceState

	AuthorizationType pulumi.StringPtrOutput              `pulumi:"authorizationType"`
	Error             AuthorizationErrorResponsePtrOutput `pulumi:"error"`
	Name              pulumi.StringOutput                 `pulumi:"name"`
	OAuth2GrantType   pulumi.StringPtrOutput              `pulumi:"oAuth2GrantType"`
	Parameters        pulumi.StringMapOutput              `pulumi:"parameters"`
	Status            pulumi.StringPtrOutput              `pulumi:"status"`
	Type              pulumi.StringOutput                 `pulumi:"type"`
}


func NewAuthorization(ctx *pulumi.Context,
	name string, args *AuthorizationArgs, opts ...pulumi.ResourceOption) (*Authorization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthorizationProviderId == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizationProviderId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource Authorization
	err := ctx.RegisterResource("azure-native:apimanagement/v20220401preview:Authorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorizationState, opts ...pulumi.ResourceOption) (*Authorization, error) {
	var resource Authorization
	err := ctx.ReadResource("azure-native:apimanagement/v20220401preview:Authorization", name, id, state, &resource, opts...)
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
	AuthorizationId         *string             `pulumi:"authorizationId"`
	AuthorizationProviderId string              `pulumi:"authorizationProviderId"`
	AuthorizationType       *string             `pulumi:"authorizationType"`
	Error                   *AuthorizationError `pulumi:"error"`
	OAuth2GrantType         *string             `pulumi:"oAuth2GrantType"`
	Parameters              map[string]string   `pulumi:"parameters"`
	ResourceGroupName       string              `pulumi:"resourceGroupName"`
	ServiceName             string              `pulumi:"serviceName"`
	Status                  *string             `pulumi:"status"`
}


type AuthorizationArgs struct {
	AuthorizationId         pulumi.StringPtrInput
	AuthorizationProviderId pulumi.StringInput
	AuthorizationType       pulumi.StringPtrInput
	Error                   AuthorizationErrorPtrInput
	OAuth2GrantType         pulumi.StringPtrInput
	Parameters              pulumi.StringMapInput
	ResourceGroupName       pulumi.StringInput
	ServiceName             pulumi.StringInput
	Status                  pulumi.StringPtrInput
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
	return reflect.TypeOf((**Authorization)(nil)).Elem()
}

func (i *Authorization) ToAuthorizationOutput() AuthorizationOutput {
	return i.ToAuthorizationOutputWithContext(context.Background())
}

func (i *Authorization) ToAuthorizationOutputWithContext(ctx context.Context) AuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationOutput)
}

type AuthorizationOutput struct{ *pulumi.OutputState }

func (AuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Authorization)(nil)).Elem()
}

func (o AuthorizationOutput) ToAuthorizationOutput() AuthorizationOutput {
	return o
}

func (o AuthorizationOutput) ToAuthorizationOutputWithContext(ctx context.Context) AuthorizationOutput {
	return o
}

func (o AuthorizationOutput) AuthorizationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Authorization) pulumi.StringPtrOutput { return v.AuthorizationType }).(pulumi.StringPtrOutput)
}

func (o AuthorizationOutput) Error() AuthorizationErrorResponsePtrOutput {
	return o.ApplyT(func(v *Authorization) AuthorizationErrorResponsePtrOutput { return v.Error }).(AuthorizationErrorResponsePtrOutput)
}

func (o AuthorizationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Authorization) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AuthorizationOutput) OAuth2GrantType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Authorization) pulumi.StringPtrOutput { return v.OAuth2GrantType }).(pulumi.StringPtrOutput)
}

func (o AuthorizationOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Authorization) pulumi.StringMapOutput { return v.Parameters }).(pulumi.StringMapOutput)
}

func (o AuthorizationOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Authorization) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o AuthorizationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Authorization) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorizationOutput{})
}
