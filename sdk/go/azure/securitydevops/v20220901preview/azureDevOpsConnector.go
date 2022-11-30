


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureDevOpsConnector struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                          `pulumi:"location"`
	Name       pulumi.StringOutput                          `pulumi:"name"`
	Properties AzureDevOpsConnectorPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                       `pulumi:"tags"`
	Type       pulumi.StringOutput                          `pulumi:"type"`
}


func NewAzureDevOpsConnector(ctx *pulumi.Context,
	name string, args *AzureDevOpsConnectorArgs, opts ...pulumi.ResourceOption) (*AzureDevOpsConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securitydevops:AzureDevOpsConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource AzureDevOpsConnector
	err := ctx.RegisterResource("azure-native:securitydevops/v20220901preview:AzureDevOpsConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAzureDevOpsConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureDevOpsConnectorState, opts ...pulumi.ResourceOption) (*AzureDevOpsConnector, error) {
	var resource AzureDevOpsConnector
	err := ctx.ReadResource("azure-native:securitydevops/v20220901preview:AzureDevOpsConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type azureDevOpsConnectorState struct {
}

type AzureDevOpsConnectorState struct {
}

func (AzureDevOpsConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureDevOpsConnectorState)(nil)).Elem()
}

type azureDevOpsConnectorArgs struct {
	AzureDevOpsConnectorName *string                         `pulumi:"azureDevOpsConnectorName"`
	Location                 *string                         `pulumi:"location"`
	Properties               *AzureDevOpsConnectorProperties `pulumi:"properties"`
	ResourceGroupName        string                          `pulumi:"resourceGroupName"`
	Tags                     map[string]string               `pulumi:"tags"`
}


type AzureDevOpsConnectorArgs struct {
	AzureDevOpsConnectorName pulumi.StringPtrInput
	Location                 pulumi.StringPtrInput
	Properties               AzureDevOpsConnectorPropertiesPtrInput
	ResourceGroupName        pulumi.StringInput
	Tags                     pulumi.StringMapInput
}

func (AzureDevOpsConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureDevOpsConnectorArgs)(nil)).Elem()
}

type AzureDevOpsConnectorInput interface {
	pulumi.Input

	ToAzureDevOpsConnectorOutput() AzureDevOpsConnectorOutput
	ToAzureDevOpsConnectorOutputWithContext(ctx context.Context) AzureDevOpsConnectorOutput
}

func (*AzureDevOpsConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureDevOpsConnector)(nil)).Elem()
}

func (i *AzureDevOpsConnector) ToAzureDevOpsConnectorOutput() AzureDevOpsConnectorOutput {
	return i.ToAzureDevOpsConnectorOutputWithContext(context.Background())
}

func (i *AzureDevOpsConnector) ToAzureDevOpsConnectorOutputWithContext(ctx context.Context) AzureDevOpsConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureDevOpsConnectorOutput)
}

type AzureDevOpsConnectorOutput struct{ *pulumi.OutputState }

func (AzureDevOpsConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureDevOpsConnector)(nil)).Elem()
}

func (o AzureDevOpsConnectorOutput) ToAzureDevOpsConnectorOutput() AzureDevOpsConnectorOutput {
	return o
}

func (o AzureDevOpsConnectorOutput) ToAzureDevOpsConnectorOutputWithContext(ctx context.Context) AzureDevOpsConnectorOutput {
	return o
}

func (o AzureDevOpsConnectorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureDevOpsConnector) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AzureDevOpsConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureDevOpsConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AzureDevOpsConnectorOutput) Properties() AzureDevOpsConnectorPropertiesResponseOutput {
	return o.ApplyT(func(v *AzureDevOpsConnector) AzureDevOpsConnectorPropertiesResponseOutput { return v.Properties }).(AzureDevOpsConnectorPropertiesResponseOutput)
}

func (o AzureDevOpsConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AzureDevOpsConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AzureDevOpsConnectorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AzureDevOpsConnector) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AzureDevOpsConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureDevOpsConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureDevOpsConnectorOutput{})
}
