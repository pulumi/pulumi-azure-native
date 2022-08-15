


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KustoPool struct {
	pulumi.CustomResourceState

	DataIngestionUri  pulumi.StringOutput      `pulumi:"dataIngestionUri"`
	EngineType        pulumi.StringPtrOutput   `pulumi:"engineType"`
	Etag              pulumi.StringOutput      `pulumi:"etag"`
	Location          pulumi.StringOutput      `pulumi:"location"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	Sku               AzureSkuResponseOutput   `pulumi:"sku"`
	State             pulumi.StringOutput      `pulumi:"state"`
	StateReason       pulumi.StringOutput      `pulumi:"stateReason"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput   `pulumi:"tags"`
	Type              pulumi.StringOutput      `pulumi:"type"`
	Uri               pulumi.StringOutput      `pulumi:"uri"`
	WorkspaceUid      pulumi.StringPtrOutput   `pulumi:"workspaceUid"`
}


func NewKustoPool(ctx *pulumi.Context,
	name string, args *KustoPoolArgs, opts ...pulumi.ResourceOption) (*KustoPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse:kustoPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:kustoPool"),
		},
	})
	opts = append(opts, aliases)
	var resource KustoPool
	err := ctx.RegisterResource("azure-native:synapse/v20210401preview:kustoPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetKustoPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KustoPoolState, opts ...pulumi.ResourceOption) (*KustoPool, error) {
	var resource KustoPool
	err := ctx.ReadResource("azure-native:synapse/v20210401preview:kustoPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type kustoPoolState struct {
}

type KustoPoolState struct {
}

func (KustoPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoPoolState)(nil)).Elem()
}

type kustoPoolArgs struct {
	EngineType        *string           `pulumi:"engineType"`
	KustoPoolName     *string           `pulumi:"kustoPoolName"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Sku               AzureSku          `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
	WorkspaceName     string            `pulumi:"workspaceName"`
	WorkspaceUid      *string           `pulumi:"workspaceUid"`
}


type KustoPoolArgs struct {
	EngineType        pulumi.StringPtrInput
	KustoPoolName     pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               AzureSkuInput
	Tags              pulumi.StringMapInput
	WorkspaceName     pulumi.StringInput
	WorkspaceUid      pulumi.StringPtrInput
}

func (KustoPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoPoolArgs)(nil)).Elem()
}

type KustoPoolInput interface {
	pulumi.Input

	ToKustoPoolOutput() KustoPoolOutput
	ToKustoPoolOutputWithContext(ctx context.Context) KustoPoolOutput
}

func (*KustoPool) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoPool)(nil)).Elem()
}

func (i *KustoPool) ToKustoPoolOutput() KustoPoolOutput {
	return i.ToKustoPoolOutputWithContext(context.Background())
}

func (i *KustoPool) ToKustoPoolOutputWithContext(ctx context.Context) KustoPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustoPoolOutput)
}

type KustoPoolOutput struct{ *pulumi.OutputState }

func (KustoPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoPool)(nil)).Elem()
}

func (o KustoPoolOutput) ToKustoPoolOutput() KustoPoolOutput {
	return o
}

func (o KustoPoolOutput) ToKustoPoolOutputWithContext(ctx context.Context) KustoPoolOutput {
	return o
}

func (o KustoPoolOutput) DataIngestionUri() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringOutput { return v.DataIngestionUri }).(pulumi.StringOutput)
}

func (o KustoPoolOutput) EngineType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringPtrOutput { return v.EngineType }).(pulumi.StringPtrOutput)
}

func (o KustoPoolOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o KustoPoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o KustoPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o KustoPoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o KustoPoolOutput) Sku() AzureSkuResponseOutput {
	return o.ApplyT(func(v *KustoPool) AzureSkuResponseOutput { return v.Sku }).(AzureSkuResponseOutput)
}

func (o KustoPoolOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o KustoPoolOutput) StateReason() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringOutput { return v.StateReason }).(pulumi.StringOutput)
}

func (o KustoPoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *KustoPool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o KustoPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o KustoPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o KustoPoolOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringOutput { return v.Uri }).(pulumi.StringOutput)
}

func (o KustoPoolOutput) WorkspaceUid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringPtrOutput { return v.WorkspaceUid }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(KustoPoolOutput{})
}
