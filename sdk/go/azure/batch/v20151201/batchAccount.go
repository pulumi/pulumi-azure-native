


package v20151201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type BatchAccount struct {
	pulumi.CustomResourceState

	AccountEndpoint              pulumi.StringOutput                    `pulumi:"accountEndpoint"`
	ActiveJobAndJobScheduleQuota pulumi.IntOutput                       `pulumi:"activeJobAndJobScheduleQuota"`
	AutoStorage                  AutoStoragePropertiesResponsePtrOutput `pulumi:"autoStorage"`
	CoreQuota                    pulumi.IntOutput                       `pulumi:"coreQuota"`
	Location                     pulumi.StringPtrOutput                 `pulumi:"location"`
	Name                         pulumi.StringOutput                    `pulumi:"name"`
	PoolQuota                    pulumi.IntOutput                       `pulumi:"poolQuota"`
	ProvisioningState            pulumi.StringPtrOutput                 `pulumi:"provisioningState"`
	Tags                         pulumi.StringMapOutput                 `pulumi:"tags"`
	Type                         pulumi.StringOutput                    `pulumi:"type"`
}


func NewBatchAccount(ctx *pulumi.Context,
	name string, args *BatchAccountArgs, opts ...pulumi.ResourceOption) (*BatchAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:batch:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170101:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170501:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170901:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20181201:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20190401:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20190801:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200301:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200501:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200901:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20210101:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20210601:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20220101:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20220601:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20221001:BatchAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource BatchAccount
	err := ctx.RegisterResource("azure-native:batch/v20151201:BatchAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBatchAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BatchAccountState, opts ...pulumi.ResourceOption) (*BatchAccount, error) {
	var resource BatchAccount
	err := ctx.ReadResource("azure-native:batch/v20151201:BatchAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type batchAccountState struct {
}

type BatchAccountState struct {
}

func (BatchAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*batchAccountState)(nil)).Elem()
}

type batchAccountArgs struct {
	AccountName       *string                    `pulumi:"accountName"`
	AutoStorage       *AutoStorageBaseProperties `pulumi:"autoStorage"`
	Location          *string                    `pulumi:"location"`
	ResourceGroupName string                     `pulumi:"resourceGroupName"`
	Tags              map[string]string          `pulumi:"tags"`
}


type BatchAccountArgs struct {
	AccountName       pulumi.StringPtrInput
	AutoStorage       AutoStorageBasePropertiesPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (BatchAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*batchAccountArgs)(nil)).Elem()
}

type BatchAccountInput interface {
	pulumi.Input

	ToBatchAccountOutput() BatchAccountOutput
	ToBatchAccountOutputWithContext(ctx context.Context) BatchAccountOutput
}

func (*BatchAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchAccount)(nil)).Elem()
}

func (i *BatchAccount) ToBatchAccountOutput() BatchAccountOutput {
	return i.ToBatchAccountOutputWithContext(context.Background())
}

func (i *BatchAccount) ToBatchAccountOutputWithContext(ctx context.Context) BatchAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchAccountOutput)
}

type BatchAccountOutput struct{ *pulumi.OutputState }

func (BatchAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchAccount)(nil)).Elem()
}

func (o BatchAccountOutput) ToBatchAccountOutput() BatchAccountOutput {
	return o
}

func (o BatchAccountOutput) ToBatchAccountOutputWithContext(ctx context.Context) BatchAccountOutput {
	return o
}

func (o BatchAccountOutput) AccountEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.StringOutput { return v.AccountEndpoint }).(pulumi.StringOutput)
}

func (o BatchAccountOutput) ActiveJobAndJobScheduleQuota() pulumi.IntOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.IntOutput { return v.ActiveJobAndJobScheduleQuota }).(pulumi.IntOutput)
}

func (o BatchAccountOutput) AutoStorage() AutoStoragePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *BatchAccount) AutoStoragePropertiesResponsePtrOutput { return v.AutoStorage }).(AutoStoragePropertiesResponsePtrOutput)
}

func (o BatchAccountOutput) CoreQuota() pulumi.IntOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.IntOutput { return v.CoreQuota }).(pulumi.IntOutput)
}

func (o BatchAccountOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o BatchAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BatchAccountOutput) PoolQuota() pulumi.IntOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.IntOutput { return v.PoolQuota }).(pulumi.IntOutput)
}

func (o BatchAccountOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o BatchAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o BatchAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BatchAccountOutput{})
}
