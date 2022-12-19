


package v20220301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActiveDirectoryConnector struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                              `pulumi:"name"`
	Properties ActiveDirectoryConnectorPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                         `pulumi:"systemData"`
	Type       pulumi.StringOutput                              `pulumi:"type"`
}


func NewActiveDirectoryConnector(ctx *pulumi.Context,
	name string, args *ActiveDirectoryConnectorArgs, opts ...pulumi.ResourceOption) (*ActiveDirectoryConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataControllerName == nil {
		return nil, errors.New("invalid value for required argument 'DataControllerName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.Properties = args.Properties.ToActiveDirectoryConnectorPropertiesOutput().ApplyT(func(v ActiveDirectoryConnectorProperties) ActiveDirectoryConnectorProperties { return *v.Defaults() }).(ActiveDirectoryConnectorPropertiesOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurearcdata:ActiveDirectoryConnector"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20220615preview:ActiveDirectoryConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource ActiveDirectoryConnector
	err := ctx.RegisterResource("azure-native:azurearcdata/v20220301preview:ActiveDirectoryConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetActiveDirectoryConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActiveDirectoryConnectorState, opts ...pulumi.ResourceOption) (*ActiveDirectoryConnector, error) {
	var resource ActiveDirectoryConnector
	err := ctx.ReadResource("azure-native:azurearcdata/v20220301preview:ActiveDirectoryConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type activeDirectoryConnectorState struct {
}

type ActiveDirectoryConnectorState struct {
}

func (ActiveDirectoryConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*activeDirectoryConnectorState)(nil)).Elem()
}

type activeDirectoryConnectorArgs struct {
	ActiveDirectoryConnectorName *string                            `pulumi:"activeDirectoryConnectorName"`
	DataControllerName           string                             `pulumi:"dataControllerName"`
	Properties                   ActiveDirectoryConnectorProperties `pulumi:"properties"`
	ResourceGroupName            string                             `pulumi:"resourceGroupName"`
}


type ActiveDirectoryConnectorArgs struct {
	ActiveDirectoryConnectorName pulumi.StringPtrInput
	DataControllerName           pulumi.StringInput
	Properties                   ActiveDirectoryConnectorPropertiesInput
	ResourceGroupName            pulumi.StringInput
}

func (ActiveDirectoryConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*activeDirectoryConnectorArgs)(nil)).Elem()
}

type ActiveDirectoryConnectorInput interface {
	pulumi.Input

	ToActiveDirectoryConnectorOutput() ActiveDirectoryConnectorOutput
	ToActiveDirectoryConnectorOutputWithContext(ctx context.Context) ActiveDirectoryConnectorOutput
}

func (*ActiveDirectoryConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryConnector)(nil)).Elem()
}

func (i *ActiveDirectoryConnector) ToActiveDirectoryConnectorOutput() ActiveDirectoryConnectorOutput {
	return i.ToActiveDirectoryConnectorOutputWithContext(context.Background())
}

func (i *ActiveDirectoryConnector) ToActiveDirectoryConnectorOutputWithContext(ctx context.Context) ActiveDirectoryConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryConnectorOutput)
}

type ActiveDirectoryConnectorOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryConnector)(nil)).Elem()
}

func (o ActiveDirectoryConnectorOutput) ToActiveDirectoryConnectorOutput() ActiveDirectoryConnectorOutput {
	return o
}

func (o ActiveDirectoryConnectorOutput) ToActiveDirectoryConnectorOutputWithContext(ctx context.Context) ActiveDirectoryConnectorOutput {
	return o
}

func (o ActiveDirectoryConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ActiveDirectoryConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ActiveDirectoryConnectorOutput) Properties() ActiveDirectoryConnectorPropertiesResponseOutput {
	return o.ApplyT(func(v *ActiveDirectoryConnector) ActiveDirectoryConnectorPropertiesResponseOutput {
		return v.Properties
	}).(ActiveDirectoryConnectorPropertiesResponseOutput)
}

func (o ActiveDirectoryConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ActiveDirectoryConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ActiveDirectoryConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ActiveDirectoryConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ActiveDirectoryConnectorOutput{})
}
