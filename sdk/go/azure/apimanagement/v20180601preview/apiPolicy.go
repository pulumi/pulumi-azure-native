


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiPolicy struct {
	pulumi.CustomResourceState

	ContentFormat pulumi.StringPtrOutput `pulumi:"contentFormat"`
	Name          pulumi.StringOutput    `pulumi:"name"`
	PolicyContent pulumi.StringOutput    `pulumi:"policyContent"`
	Type          pulumi.StringOutput    `pulumi:"type"`
}


func NewApiPolicy(ctx *pulumi.Context,
	name string, args *ApiPolicyArgs, opts ...pulumi.ResourceOption) (*ApiPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.PolicyContent == nil {
		return nil, errors.New("invalid value for required argument 'PolicyContent'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if isZero(args.ContentFormat) {
		args.ContentFormat = pulumi.StringPtr("xml")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:ApiPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ApiPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ApiPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ApiPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ApiPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ApiPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ApiPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ApiPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ApiPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ApiPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiPolicy
	err := ctx.RegisterResource("azure-native:apimanagement/v20180601preview:ApiPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApiPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiPolicyState, opts ...pulumi.ResourceOption) (*ApiPolicy, error) {
	var resource ApiPolicy
	err := ctx.ReadResource("azure-native:apimanagement/v20180601preview:ApiPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type apiPolicyState struct {
}

type ApiPolicyState struct {
}

func (ApiPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiPolicyState)(nil)).Elem()
}

type apiPolicyArgs struct {
	ApiId             string  `pulumi:"apiId"`
	ContentFormat     *string `pulumi:"contentFormat"`
	PolicyContent     string  `pulumi:"policyContent"`
	PolicyId          *string `pulumi:"policyId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type ApiPolicyArgs struct {
	ApiId             pulumi.StringInput
	ContentFormat     pulumi.StringPtrInput
	PolicyContent     pulumi.StringInput
	PolicyId          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (ApiPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiPolicyArgs)(nil)).Elem()
}

type ApiPolicyInput interface {
	pulumi.Input

	ToApiPolicyOutput() ApiPolicyOutput
	ToApiPolicyOutputWithContext(ctx context.Context) ApiPolicyOutput
}

func (*ApiPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPolicy)(nil)).Elem()
}

func (i *ApiPolicy) ToApiPolicyOutput() ApiPolicyOutput {
	return i.ToApiPolicyOutputWithContext(context.Background())
}

func (i *ApiPolicy) ToApiPolicyOutputWithContext(ctx context.Context) ApiPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPolicyOutput)
}

type ApiPolicyOutput struct{ *pulumi.OutputState }

func (ApiPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPolicy)(nil)).Elem()
}

func (o ApiPolicyOutput) ToApiPolicyOutput() ApiPolicyOutput {
	return o
}

func (o ApiPolicyOutput) ToApiPolicyOutputWithContext(ctx context.Context) ApiPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApiPolicyOutput{})
}
