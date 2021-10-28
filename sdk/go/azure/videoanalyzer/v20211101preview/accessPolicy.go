


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessPolicy struct {
	pulumi.CustomResourceState

	Authentication JwtAuthenticationResponsePtrOutput `pulumi:"authentication"`
	Name           pulumi.StringOutput                `pulumi:"name"`
	Role           pulumi.StringPtrOutput             `pulumi:"role"`
	SystemData     SystemDataResponseOutput           `pulumi:"systemData"`
	Type           pulumi.StringOutput                `pulumi:"type"`
}


func NewAccessPolicy(ctx *pulumi.Context,
	name string, args *AccessPolicyArgs, opts ...pulumi.ResourceOption) (*AccessPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:videoanalyzer/v20211101preview:AccessPolicy"),
		},
		{
			Type: pulumi.String("azure-native:videoanalyzer:AccessPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:videoanalyzer:AccessPolicy"),
		},
		{
			Type: pulumi.String("azure-native:videoanalyzer/v20210501preview:AccessPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:videoanalyzer/v20210501preview:AccessPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource AccessPolicy
	err := ctx.RegisterResource("azure-native:videoanalyzer/v20211101preview:AccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessPolicyState, opts ...pulumi.ResourceOption) (*AccessPolicy, error) {
	var resource AccessPolicy
	err := ctx.ReadResource("azure-native:videoanalyzer/v20211101preview:AccessPolicy", name, id, state, &resource, opts...)
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
	AccessPolicyName  *string            `pulumi:"accessPolicyName"`
	AccountName       string             `pulumi:"accountName"`
	Authentication    *JwtAuthentication `pulumi:"authentication"`
	ResourceGroupName string             `pulumi:"resourceGroupName"`
	Role              *string            `pulumi:"role"`
}


type AccessPolicyArgs struct {
	AccessPolicyName  pulumi.StringPtrInput
	AccountName       pulumi.StringInput
	Authentication    JwtAuthenticationPtrInput
	ResourceGroupName pulumi.StringInput
	Role              pulumi.StringPtrInput
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
	return reflect.TypeOf((*AccessPolicy)(nil))
}

func (i *AccessPolicy) ToAccessPolicyOutput() AccessPolicyOutput {
	return i.ToAccessPolicyOutputWithContext(context.Background())
}

func (i *AccessPolicy) ToAccessPolicyOutputWithContext(ctx context.Context) AccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyOutput)
}

type AccessPolicyOutput struct{ *pulumi.OutputState }

func (AccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicy)(nil))
}

func (o AccessPolicyOutput) ToAccessPolicyOutput() AccessPolicyOutput {
	return o
}

func (o AccessPolicyOutput) ToAccessPolicyOutputWithContext(ctx context.Context) AccessPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPolicyInput)(nil)).Elem(), &AccessPolicy{})
	pulumi.RegisterOutputType(AccessPolicyOutput{})
}
