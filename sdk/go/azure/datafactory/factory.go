


package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Factory struct {
	pulumi.CustomResourceState

	CreateTime          pulumi.StringOutput                           `pulumi:"createTime"`
	ETag                pulumi.StringOutput                           `pulumi:"eTag"`
	Encryption          EncryptionConfigurationResponsePtrOutput      `pulumi:"encryption"`
	GlobalParameters    GlobalParameterSpecificationResponseMapOutput `pulumi:"globalParameters"`
	Identity            FactoryIdentityResponsePtrOutput              `pulumi:"identity"`
	Location            pulumi.StringPtrOutput                        `pulumi:"location"`
	Name                pulumi.StringOutput                           `pulumi:"name"`
	ProvisioningState   pulumi.StringOutput                           `pulumi:"provisioningState"`
	PublicNetworkAccess pulumi.StringPtrOutput                        `pulumi:"publicNetworkAccess"`
	RepoConfiguration   pulumi.AnyOutput                              `pulumi:"repoConfiguration"`
	Tags                pulumi.StringMapOutput                        `pulumi:"tags"`
	Type                pulumi.StringOutput                           `pulumi:"type"`
	Version             pulumi.StringOutput                           `pulumi:"version"`
}


func NewFactory(ctx *pulumi.Context,
	name string, args *FactoryArgs, opts ...pulumi.ResourceOption) (*Factory, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datafactory/v20170901preview:Factory"),
		},
		{
			Type: pulumi.String("azure-native:datafactory/v20180601:Factory"),
		},
	})
	opts = append(opts, aliases)
	var resource Factory
	err := ctx.RegisterResource("azure-native:datafactory:Factory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFactory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FactoryState, opts ...pulumi.ResourceOption) (*Factory, error) {
	var resource Factory
	err := ctx.ReadResource("azure-native:datafactory:Factory", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type factoryState struct {
}

type FactoryState struct {
}

func (FactoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*factoryState)(nil)).Elem()
}

type factoryArgs struct {
	Encryption          *EncryptionConfiguration                `pulumi:"encryption"`
	FactoryName         *string                                 `pulumi:"factoryName"`
	GlobalParameters    map[string]GlobalParameterSpecification `pulumi:"globalParameters"`
	Identity            *FactoryIdentity                        `pulumi:"identity"`
	Location            *string                                 `pulumi:"location"`
	PublicNetworkAccess *string                                 `pulumi:"publicNetworkAccess"`
	RepoConfiguration   interface{}                             `pulumi:"repoConfiguration"`
	ResourceGroupName   string                                  `pulumi:"resourceGroupName"`
	Tags                map[string]string                       `pulumi:"tags"`
}


type FactoryArgs struct {
	Encryption          EncryptionConfigurationPtrInput
	FactoryName         pulumi.StringPtrInput
	GlobalParameters    GlobalParameterSpecificationMapInput
	Identity            FactoryIdentityPtrInput
	Location            pulumi.StringPtrInput
	PublicNetworkAccess pulumi.StringPtrInput
	RepoConfiguration   pulumi.Input
	ResourceGroupName   pulumi.StringInput
	Tags                pulumi.StringMapInput
}

func (FactoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*factoryArgs)(nil)).Elem()
}

type FactoryInput interface {
	pulumi.Input

	ToFactoryOutput() FactoryOutput
	ToFactoryOutputWithContext(ctx context.Context) FactoryOutput
}

func (*Factory) ElementType() reflect.Type {
	return reflect.TypeOf((*Factory)(nil))
}

func (i *Factory) ToFactoryOutput() FactoryOutput {
	return i.ToFactoryOutputWithContext(context.Background())
}

func (i *Factory) ToFactoryOutputWithContext(ctx context.Context) FactoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FactoryOutput)
}

type FactoryOutput struct{ *pulumi.OutputState }

func (FactoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Factory)(nil))
}

func (o FactoryOutput) ToFactoryOutput() FactoryOutput {
	return o
}

func (o FactoryOutput) ToFactoryOutputWithContext(ctx context.Context) FactoryOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FactoryOutput{})
}
