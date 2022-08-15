


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BlobInventoryPolicy struct {
	pulumi.CustomResourceState

	LastModifiedTime pulumi.StringOutput                     `pulumi:"lastModifiedTime"`
	Name             pulumi.StringOutput                     `pulumi:"name"`
	Policy           BlobInventoryPolicySchemaResponseOutput `pulumi:"policy"`
	SystemData       SystemDataResponseOutput                `pulumi:"systemData"`
	Type             pulumi.StringOutput                     `pulumi:"type"`
}


func NewBlobInventoryPolicy(ctx *pulumi.Context,
	name string, args *BlobInventoryPolicyArgs, opts ...pulumi.ResourceOption) (*BlobInventoryPolicy, error) {
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
			Type: pulumi.String("azure-native:storage:BlobInventoryPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:BlobInventoryPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:BlobInventoryPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:BlobInventoryPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:BlobInventoryPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:BlobInventoryPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:BlobInventoryPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:BlobInventoryPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:BlobInventoryPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource BlobInventoryPolicy
	err := ctx.RegisterResource("azure-native:storage/v20210601:BlobInventoryPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBlobInventoryPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobInventoryPolicyState, opts ...pulumi.ResourceOption) (*BlobInventoryPolicy, error) {
	var resource BlobInventoryPolicy
	err := ctx.ReadResource("azure-native:storage/v20210601:BlobInventoryPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type blobInventoryPolicyState struct {
}

type BlobInventoryPolicyState struct {
}

func (BlobInventoryPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobInventoryPolicyState)(nil)).Elem()
}

type blobInventoryPolicyArgs struct {
	AccountName             string                    `pulumi:"accountName"`
	BlobInventoryPolicyName *string                   `pulumi:"blobInventoryPolicyName"`
	Policy                  BlobInventoryPolicySchema `pulumi:"policy"`
	ResourceGroupName       string                    `pulumi:"resourceGroupName"`
}


type BlobInventoryPolicyArgs struct {
	AccountName             pulumi.StringInput
	BlobInventoryPolicyName pulumi.StringPtrInput
	Policy                  BlobInventoryPolicySchemaInput
	ResourceGroupName       pulumi.StringInput
}

func (BlobInventoryPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobInventoryPolicyArgs)(nil)).Elem()
}

type BlobInventoryPolicyInput interface {
	pulumi.Input

	ToBlobInventoryPolicyOutput() BlobInventoryPolicyOutput
	ToBlobInventoryPolicyOutputWithContext(ctx context.Context) BlobInventoryPolicyOutput
}

func (*BlobInventoryPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobInventoryPolicy)(nil)).Elem()
}

func (i *BlobInventoryPolicy) ToBlobInventoryPolicyOutput() BlobInventoryPolicyOutput {
	return i.ToBlobInventoryPolicyOutputWithContext(context.Background())
}

func (i *BlobInventoryPolicy) ToBlobInventoryPolicyOutputWithContext(ctx context.Context) BlobInventoryPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobInventoryPolicyOutput)
}

type BlobInventoryPolicyOutput struct{ *pulumi.OutputState }

func (BlobInventoryPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobInventoryPolicy)(nil)).Elem()
}

func (o BlobInventoryPolicyOutput) ToBlobInventoryPolicyOutput() BlobInventoryPolicyOutput {
	return o
}

func (o BlobInventoryPolicyOutput) ToBlobInventoryPolicyOutputWithContext(ctx context.Context) BlobInventoryPolicyOutput {
	return o
}

func (o BlobInventoryPolicyOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobInventoryPolicy) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o BlobInventoryPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobInventoryPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BlobInventoryPolicyOutput) Policy() BlobInventoryPolicySchemaResponseOutput {
	return o.ApplyT(func(v *BlobInventoryPolicy) BlobInventoryPolicySchemaResponseOutput { return v.Policy }).(BlobInventoryPolicySchemaResponseOutput)
}

func (o BlobInventoryPolicyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BlobInventoryPolicy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o BlobInventoryPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobInventoryPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BlobInventoryPolicyOutput{})
}
