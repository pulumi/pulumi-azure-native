


package timeseriesinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessPolicy struct {
	pulumi.CustomResourceState

	Description       pulumi.StringPtrOutput   `pulumi:"description"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	PrincipalObjectId pulumi.StringPtrOutput   `pulumi:"principalObjectId"`
	Roles             pulumi.StringArrayOutput `pulumi:"roles"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewAccessPolicy(ctx *pulumi.Context,
	name string, args *AccessPolicyArgs, opts ...pulumi.ResourceOption) (*AccessPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20170228preview:AccessPolicy"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20171115:AccessPolicy"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20180815preview:AccessPolicy"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20200515:AccessPolicy"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210331preview:AccessPolicy"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210630preview:AccessPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource AccessPolicy
	err := ctx.RegisterResource("azure-native:timeseriesinsights:AccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessPolicyState, opts ...pulumi.ResourceOption) (*AccessPolicy, error) {
	var resource AccessPolicy
	err := ctx.ReadResource("azure-native:timeseriesinsights:AccessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type accessPolicyState struct {
}

type AccessPolicyState struct {
}

func (AccessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPolicyState)(nil)).Elem()
}

type accessPolicyArgs struct {
	AccessPolicyName  *string  `pulumi:"accessPolicyName"`
	Description       *string  `pulumi:"description"`
	EnvironmentName   string   `pulumi:"environmentName"`
	PrincipalObjectId *string  `pulumi:"principalObjectId"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	Roles             []string `pulumi:"roles"`
}


type AccessPolicyArgs struct {
	AccessPolicyName  pulumi.StringPtrInput
	Description       pulumi.StringPtrInput
	EnvironmentName   pulumi.StringInput
	PrincipalObjectId pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Roles             pulumi.StringArrayInput
}

func (AccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPolicyArgs)(nil)).Elem()
}

type AccessPolicyInput interface {
	pulumi.Input

	ToAccessPolicyOutput() AccessPolicyOutput
	ToAccessPolicyOutputWithContext(ctx context.Context) AccessPolicyOutput
}

func (*AccessPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPolicy)(nil)).Elem()
}

func (i *AccessPolicy) ToAccessPolicyOutput() AccessPolicyOutput {
	return i.ToAccessPolicyOutputWithContext(context.Background())
}

func (i *AccessPolicy) ToAccessPolicyOutputWithContext(ctx context.Context) AccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyOutput)
}

type AccessPolicyOutput struct{ *pulumi.OutputState }

func (AccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPolicy)(nil)).Elem()
}

func (o AccessPolicyOutput) ToAccessPolicyOutput() AccessPolicyOutput {
	return o
}

func (o AccessPolicyOutput) ToAccessPolicyOutputWithContext(ctx context.Context) AccessPolicyOutput {
	return o
}

func (o AccessPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AccessPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AccessPolicyOutput) PrincipalObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessPolicy) pulumi.StringPtrOutput { return v.PrincipalObjectId }).(pulumi.StringPtrOutput)
}

func (o AccessPolicyOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AccessPolicy) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o AccessPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessPolicyOutput{})
}
