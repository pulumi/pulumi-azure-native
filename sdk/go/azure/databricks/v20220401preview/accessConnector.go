


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessConnector struct {
	pulumi.CustomResourceState

	Identity   IdentityDataResponsePtrOutput           `pulumi:"identity"`
	Location   pulumi.StringOutput                     `pulumi:"location"`
	Name       pulumi.StringOutput                     `pulumi:"name"`
	Properties AccessConnectorPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                  `pulumi:"tags"`
	Type       pulumi.StringOutput                     `pulumi:"type"`
}


func NewAccessConnector(ctx *pulumi.Context,
	name string, args *AccessConnectorArgs, opts ...pulumi.ResourceOption) (*AccessConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databricks:AccessConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource AccessConnector
	err := ctx.RegisterResource("azure-native:databricks/v20220401preview:AccessConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAccessConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessConnectorState, opts ...pulumi.ResourceOption) (*AccessConnector, error) {
	var resource AccessConnector
	err := ctx.ReadResource("azure-native:databricks/v20220401preview:AccessConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type accessConnectorState struct {
}

type AccessConnectorState struct {
}

func (AccessConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessConnectorState)(nil)).Elem()
}

type accessConnectorArgs struct {
	ConnectorName     *string           `pulumi:"connectorName"`
	Identity          *IdentityData     `pulumi:"identity"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type AccessConnectorArgs struct {
	ConnectorName     pulumi.StringPtrInput
	Identity          IdentityDataPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (AccessConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessConnectorArgs)(nil)).Elem()
}

type AccessConnectorInput interface {
	pulumi.Input

	ToAccessConnectorOutput() AccessConnectorOutput
	ToAccessConnectorOutputWithContext(ctx context.Context) AccessConnectorOutput
}

func (*AccessConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessConnector)(nil)).Elem()
}

func (i *AccessConnector) ToAccessConnectorOutput() AccessConnectorOutput {
	return i.ToAccessConnectorOutputWithContext(context.Background())
}

func (i *AccessConnector) ToAccessConnectorOutputWithContext(ctx context.Context) AccessConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessConnectorOutput)
}

type AccessConnectorOutput struct{ *pulumi.OutputState }

func (AccessConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessConnector)(nil)).Elem()
}

func (o AccessConnectorOutput) ToAccessConnectorOutput() AccessConnectorOutput {
	return o
}

func (o AccessConnectorOutput) ToAccessConnectorOutputWithContext(ctx context.Context) AccessConnectorOutput {
	return o
}

func (o AccessConnectorOutput) Identity() IdentityDataResponsePtrOutput {
	return o.ApplyT(func(v *AccessConnector) IdentityDataResponsePtrOutput { return v.Identity }).(IdentityDataResponsePtrOutput)
}

func (o AccessConnectorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessConnector) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AccessConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AccessConnectorOutput) Properties() AccessConnectorPropertiesResponseOutput {
	return o.ApplyT(func(v *AccessConnector) AccessConnectorPropertiesResponseOutput { return v.Properties }).(AccessConnectorPropertiesResponseOutput)
}

func (o AccessConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AccessConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AccessConnectorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AccessConnector) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AccessConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessConnectorOutput{})
}
