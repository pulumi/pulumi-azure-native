


package v20180301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StorageAccountManagementPolicies struct {
	pulumi.CustomResourceState

	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	Name             pulumi.StringOutput `pulumi:"name"`
	Policy           pulumi.AnyOutput    `pulumi:"policy"`
	Type             pulumi.StringOutput `pulumi:"type"`
}


func NewStorageAccountManagementPolicies(ctx *pulumi.Context,
	name string, args *StorageAccountManagementPoliciesArgs, opts ...pulumi.ResourceOption) (*StorageAccountManagementPolicies, error) {
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
			Type: pulumi.String("azure-nextgen:storage/v20180301preview:StorageAccountManagementPolicies"),
		},
		{
			Type: pulumi.String("azure-native:storage:StorageAccountManagementPolicies"),
		},
		{
			Type: pulumi.String("azure-nextgen:storage:StorageAccountManagementPolicies"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20181101:StorageAccountManagementPolicies"),
		},
		{
			Type: pulumi.String("azure-nextgen:storage/v20181101:StorageAccountManagementPolicies"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190401:StorageAccountManagementPolicies"),
		},
		{
			Type: pulumi.String("azure-nextgen:storage/v20190401:StorageAccountManagementPolicies"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:StorageAccountManagementPolicies"),
		},
		{
			Type: pulumi.String("azure-nextgen:storage/v20190601:StorageAccountManagementPolicies"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:StorageAccountManagementPolicies"),
		},
		{
			Type: pulumi.String("azure-nextgen:storage/v20200801preview:StorageAccountManagementPolicies"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:StorageAccountManagementPolicies"),
		},
		{
			Type: pulumi.String("azure-nextgen:storage/v20210101:StorageAccountManagementPolicies"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:StorageAccountManagementPolicies"),
		},
		{
			Type: pulumi.String("azure-nextgen:storage/v20210201:StorageAccountManagementPolicies"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:StorageAccountManagementPolicies"),
		},
		{
			Type: pulumi.String("azure-nextgen:storage/v20210401:StorageAccountManagementPolicies"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:StorageAccountManagementPolicies"),
		},
		{
			Type: pulumi.String("azure-nextgen:storage/v20210601:StorageAccountManagementPolicies"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageAccountManagementPolicies
	err := ctx.RegisterResource("azure-native:storage/v20180301preview:StorageAccountManagementPolicies", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStorageAccountManagementPolicies(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageAccountManagementPoliciesState, opts ...pulumi.ResourceOption) (*StorageAccountManagementPolicies, error) {
	var resource StorageAccountManagementPolicies
	err := ctx.ReadResource("azure-native:storage/v20180301preview:StorageAccountManagementPolicies", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type storageAccountManagementPoliciesState struct {
}

type StorageAccountManagementPoliciesState struct {
}

func (StorageAccountManagementPoliciesState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountManagementPoliciesState)(nil)).Elem()
}

type storageAccountManagementPoliciesArgs struct {
	AccountName          string      `pulumi:"accountName"`
	ManagementPolicyName *string     `pulumi:"managementPolicyName"`
	Policy               interface{} `pulumi:"policy"`
	ResourceGroupName    string      `pulumi:"resourceGroupName"`
}


type StorageAccountManagementPoliciesArgs struct {
	AccountName          pulumi.StringInput
	ManagementPolicyName pulumi.StringPtrInput
	Policy               pulumi.Input
	ResourceGroupName    pulumi.StringInput
}

func (StorageAccountManagementPoliciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountManagementPoliciesArgs)(nil)).Elem()
}

type StorageAccountManagementPoliciesInput interface {
	pulumi.Input

	ToStorageAccountManagementPoliciesOutput() StorageAccountManagementPoliciesOutput
	ToStorageAccountManagementPoliciesOutputWithContext(ctx context.Context) StorageAccountManagementPoliciesOutput
}

func (*StorageAccountManagementPolicies) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountManagementPolicies)(nil))
}

func (i *StorageAccountManagementPolicies) ToStorageAccountManagementPoliciesOutput() StorageAccountManagementPoliciesOutput {
	return i.ToStorageAccountManagementPoliciesOutputWithContext(context.Background())
}

func (i *StorageAccountManagementPolicies) ToStorageAccountManagementPoliciesOutputWithContext(ctx context.Context) StorageAccountManagementPoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountManagementPoliciesOutput)
}

type StorageAccountManagementPoliciesOutput struct{ *pulumi.OutputState }

func (StorageAccountManagementPoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountManagementPolicies)(nil))
}

func (o StorageAccountManagementPoliciesOutput) ToStorageAccountManagementPoliciesOutput() StorageAccountManagementPoliciesOutput {
	return o
}

func (o StorageAccountManagementPoliciesOutput) ToStorageAccountManagementPoliciesOutputWithContext(ctx context.Context) StorageAccountManagementPoliciesOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StorageAccountManagementPoliciesOutput{})
}
