


package v20190601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ManagementPolicy struct {
	pulumi.CustomResourceState

	LastModifiedTime pulumi.StringOutput                  `pulumi:"lastModifiedTime"`
	Name             pulumi.StringOutput                  `pulumi:"name"`
	Policy           ManagementPolicySchemaResponseOutput `pulumi:"policy"`
	Type             pulumi.StringOutput                  `pulumi:"type"`
}


func NewManagementPolicy(ctx *pulumi.Context,
	name string, args *ManagementPolicyArgs, opts ...pulumi.ResourceOption) (*ManagementPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180301preview:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20181101:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190401:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220901:ManagementPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagementPolicy
	err := ctx.RegisterResource("azure-native:storage/v20190601:ManagementPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagementPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementPolicyState, opts ...pulumi.ResourceOption) (*ManagementPolicy, error) {
	var resource ManagementPolicy
	err := ctx.ReadResource("azure-native:storage/v20190601:ManagementPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managementPolicyState struct {
}

type ManagementPolicyState struct {
}

func (ManagementPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementPolicyState)(nil)).Elem()
}

type managementPolicyArgs struct {
	AccountName          string                 `pulumi:"accountName"`
	ManagementPolicyName *string                `pulumi:"managementPolicyName"`
	Policy               ManagementPolicySchema `pulumi:"policy"`
	ResourceGroupName    string                 `pulumi:"resourceGroupName"`
}


type ManagementPolicyArgs struct {
	AccountName          pulumi.StringInput
	ManagementPolicyName pulumi.StringPtrInput
	Policy               ManagementPolicySchemaInput
	ResourceGroupName    pulumi.StringInput
}

func (ManagementPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementPolicyArgs)(nil)).Elem()
}

type ManagementPolicyInput interface {
	pulumi.Input

	ToManagementPolicyOutput() ManagementPolicyOutput
	ToManagementPolicyOutputWithContext(ctx context.Context) ManagementPolicyOutput
}

func (*ManagementPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicy)(nil)).Elem()
}

func (i *ManagementPolicy) ToManagementPolicyOutput() ManagementPolicyOutput {
	return i.ToManagementPolicyOutputWithContext(context.Background())
}

func (i *ManagementPolicy) ToManagementPolicyOutputWithContext(ctx context.Context) ManagementPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyOutput)
}

type ManagementPolicyOutput struct{ *pulumi.OutputState }

func (ManagementPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicy)(nil)).Elem()
}

func (o ManagementPolicyOutput) ToManagementPolicyOutput() ManagementPolicyOutput {
	return o
}

func (o ManagementPolicyOutput) ToManagementPolicyOutputWithContext(ctx context.Context) ManagementPolicyOutput {
	return o
}

func (o ManagementPolicyOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementPolicy) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o ManagementPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagementPolicyOutput) Policy() ManagementPolicySchemaResponseOutput {
	return o.ApplyT(func(v *ManagementPolicy) ManagementPolicySchemaResponseOutput { return v.Policy }).(ManagementPolicySchemaResponseOutput)
}

func (o ManagementPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagementPolicyOutput{})
}
