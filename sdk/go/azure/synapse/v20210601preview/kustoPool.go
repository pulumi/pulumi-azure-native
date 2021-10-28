


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KustoPool struct {
	pulumi.CustomResourceState

	DataIngestionUri      pulumi.StringOutput                  `pulumi:"dataIngestionUri"`
	EnablePurge           pulumi.BoolPtrOutput                 `pulumi:"enablePurge"`
	EnableStreamingIngest pulumi.BoolPtrOutput                 `pulumi:"enableStreamingIngest"`
	Etag                  pulumi.StringOutput                  `pulumi:"etag"`
	LanguageExtensions    LanguageExtensionsListResponseOutput `pulumi:"languageExtensions"`
	Location              pulumi.StringOutput                  `pulumi:"location"`
	Name                  pulumi.StringOutput                  `pulumi:"name"`
	OptimizedAutoscale    OptimizedAutoscaleResponsePtrOutput  `pulumi:"optimizedAutoscale"`
	ProvisioningState     pulumi.StringOutput                  `pulumi:"provisioningState"`
	Sku                   AzureSkuResponseOutput               `pulumi:"sku"`
	State                 pulumi.StringOutput                  `pulumi:"state"`
	StateReason           pulumi.StringOutput                  `pulumi:"stateReason"`
	SystemData            SystemDataResponseOutput             `pulumi:"systemData"`
	Tags                  pulumi.StringMapOutput               `pulumi:"tags"`
	Type                  pulumi.StringOutput                  `pulumi:"type"`
	Uri                   pulumi.StringOutput                  `pulumi:"uri"`
	WorkspaceUID          pulumi.StringPtrOutput               `pulumi:"workspaceUID"`
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
	if args.EnablePurge == nil {
		args.EnablePurge = pulumi.BoolPtr(false)
	}
	if args.EnableStreamingIngest == nil {
		args.EnableStreamingIngest = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:synapse/v20210601preview:KustoPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse:KustoPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse:KustoPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:KustoPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20210401preview:KustoPool"),
		},
	})
	opts = append(opts, aliases)
	var resource KustoPool
	err := ctx.RegisterResource("azure-native:synapse/v20210601preview:KustoPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetKustoPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KustoPoolState, opts ...pulumi.ResourceOption) (*KustoPool, error) {
	var resource KustoPool
	err := ctx.ReadResource("azure-native:synapse/v20210601preview:KustoPool", name, id, state, &resource, opts...)
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
	EnablePurge           *bool               `pulumi:"enablePurge"`
	EnableStreamingIngest *bool               `pulumi:"enableStreamingIngest"`
	KustoPoolName         *string             `pulumi:"kustoPoolName"`
	Location              *string             `pulumi:"location"`
	OptimizedAutoscale    *OptimizedAutoscale `pulumi:"optimizedAutoscale"`
	ResourceGroupName     string              `pulumi:"resourceGroupName"`
	Sku                   AzureSku            `pulumi:"sku"`
	Tags                  map[string]string   `pulumi:"tags"`
	WorkspaceName         string              `pulumi:"workspaceName"`
	WorkspaceUID          *string             `pulumi:"workspaceUID"`
}


type KustoPoolArgs struct {
	EnablePurge           pulumi.BoolPtrInput
	EnableStreamingIngest pulumi.BoolPtrInput
	KustoPoolName         pulumi.StringPtrInput
	Location              pulumi.StringPtrInput
	OptimizedAutoscale    OptimizedAutoscalePtrInput
	ResourceGroupName     pulumi.StringInput
	Sku                   AzureSkuInput
	Tags                  pulumi.StringMapInput
	WorkspaceName         pulumi.StringInput
	WorkspaceUID          pulumi.StringPtrInput
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
	return reflect.TypeOf((*KustoPool)(nil))
}

func (i *KustoPool) ToKustoPoolOutput() KustoPoolOutput {
	return i.ToKustoPoolOutputWithContext(context.Background())
}

func (i *KustoPool) ToKustoPoolOutputWithContext(ctx context.Context) KustoPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustoPoolOutput)
}

type KustoPoolOutput struct{ *pulumi.OutputState }

func (KustoPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KustoPool)(nil))
}

func (o KustoPoolOutput) ToKustoPoolOutput() KustoPoolOutput {
	return o
}

func (o KustoPoolOutput) ToKustoPoolOutputWithContext(ctx context.Context) KustoPoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(KustoPoolOutput{})
}
