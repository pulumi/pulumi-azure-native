


package healthcareapis

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IotConnectorFhirDestination struct {
	pulumi.CustomResourceState

	Etag                           pulumi.StringPtrOutput             `pulumi:"etag"`
	FhirMapping                    IotMappingPropertiesResponseOutput `pulumi:"fhirMapping"`
	FhirServiceResourceId          pulumi.StringOutput                `pulumi:"fhirServiceResourceId"`
	Location                       pulumi.StringPtrOutput             `pulumi:"location"`
	Name                           pulumi.StringOutput                `pulumi:"name"`
	ResourceIdentityResolutionType pulumi.StringOutput                `pulumi:"resourceIdentityResolutionType"`
	SystemData                     SystemDataResponseOutput           `pulumi:"systemData"`
	Type                           pulumi.StringOutput                `pulumi:"type"`
}


func NewIotConnectorFhirDestination(ctx *pulumi.Context,
	name string, args *IotConnectorFhirDestinationArgs, opts ...pulumi.ResourceOption) (*IotConnectorFhirDestination, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FhirMapping == nil {
		return nil, errors.New("invalid value for required argument 'FhirMapping'")
	}
	if args.FhirServiceResourceId == nil {
		return nil, errors.New("invalid value for required argument 'FhirServiceResourceId'")
	}
	if args.IotConnectorName == nil {
		return nil, errors.New("invalid value for required argument 'IotConnectorName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceIdentityResolutionType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceIdentityResolutionType'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:healthcareapis/v20210601preview:IotConnectorFhirDestination"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20211101:IotConnectorFhirDestination"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220131preview:IotConnectorFhirDestination"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220515:IotConnectorFhirDestination"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220601:IotConnectorFhirDestination"),
		},
	})
	opts = append(opts, aliases)
	var resource IotConnectorFhirDestination
	err := ctx.RegisterResource("azure-native:healthcareapis:IotConnectorFhirDestination", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIotConnectorFhirDestination(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotConnectorFhirDestinationState, opts ...pulumi.ResourceOption) (*IotConnectorFhirDestination, error) {
	var resource IotConnectorFhirDestination
	err := ctx.ReadResource("azure-native:healthcareapis:IotConnectorFhirDestination", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type iotConnectorFhirDestinationState struct {
}

type IotConnectorFhirDestinationState struct {
}

func (IotConnectorFhirDestinationState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotConnectorFhirDestinationState)(nil)).Elem()
}

type iotConnectorFhirDestinationArgs struct {
	FhirDestinationName            *string              `pulumi:"fhirDestinationName"`
	FhirMapping                    IotMappingProperties `pulumi:"fhirMapping"`
	FhirServiceResourceId          string               `pulumi:"fhirServiceResourceId"`
	IotConnectorName               string               `pulumi:"iotConnectorName"`
	Location                       *string              `pulumi:"location"`
	ResourceGroupName              string               `pulumi:"resourceGroupName"`
	ResourceIdentityResolutionType string               `pulumi:"resourceIdentityResolutionType"`
	WorkspaceName                  string               `pulumi:"workspaceName"`
}


type IotConnectorFhirDestinationArgs struct {
	FhirDestinationName            pulumi.StringPtrInput
	FhirMapping                    IotMappingPropertiesInput
	FhirServiceResourceId          pulumi.StringInput
	IotConnectorName               pulumi.StringInput
	Location                       pulumi.StringPtrInput
	ResourceGroupName              pulumi.StringInput
	ResourceIdentityResolutionType pulumi.StringInput
	WorkspaceName                  pulumi.StringInput
}

func (IotConnectorFhirDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotConnectorFhirDestinationArgs)(nil)).Elem()
}

type IotConnectorFhirDestinationInput interface {
	pulumi.Input

	ToIotConnectorFhirDestinationOutput() IotConnectorFhirDestinationOutput
	ToIotConnectorFhirDestinationOutputWithContext(ctx context.Context) IotConnectorFhirDestinationOutput
}

func (*IotConnectorFhirDestination) ElementType() reflect.Type {
	return reflect.TypeOf((**IotConnectorFhirDestination)(nil)).Elem()
}

func (i *IotConnectorFhirDestination) ToIotConnectorFhirDestinationOutput() IotConnectorFhirDestinationOutput {
	return i.ToIotConnectorFhirDestinationOutputWithContext(context.Background())
}

func (i *IotConnectorFhirDestination) ToIotConnectorFhirDestinationOutputWithContext(ctx context.Context) IotConnectorFhirDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotConnectorFhirDestinationOutput)
}

type IotConnectorFhirDestinationOutput struct{ *pulumi.OutputState }

func (IotConnectorFhirDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotConnectorFhirDestination)(nil)).Elem()
}

func (o IotConnectorFhirDestinationOutput) ToIotConnectorFhirDestinationOutput() IotConnectorFhirDestinationOutput {
	return o
}

func (o IotConnectorFhirDestinationOutput) ToIotConnectorFhirDestinationOutputWithContext(ctx context.Context) IotConnectorFhirDestinationOutput {
	return o
}

func (o IotConnectorFhirDestinationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotConnectorFhirDestination) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o IotConnectorFhirDestinationOutput) FhirMapping() IotMappingPropertiesResponseOutput {
	return o.ApplyT(func(v *IotConnectorFhirDestination) IotMappingPropertiesResponseOutput { return v.FhirMapping }).(IotMappingPropertiesResponseOutput)
}

func (o IotConnectorFhirDestinationOutput) FhirServiceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *IotConnectorFhirDestination) pulumi.StringOutput { return v.FhirServiceResourceId }).(pulumi.StringOutput)
}

func (o IotConnectorFhirDestinationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotConnectorFhirDestination) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o IotConnectorFhirDestinationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IotConnectorFhirDestination) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IotConnectorFhirDestinationOutput) ResourceIdentityResolutionType() pulumi.StringOutput {
	return o.ApplyT(func(v *IotConnectorFhirDestination) pulumi.StringOutput { return v.ResourceIdentityResolutionType }).(pulumi.StringOutput)
}

func (o IotConnectorFhirDestinationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *IotConnectorFhirDestination) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o IotConnectorFhirDestinationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IotConnectorFhirDestination) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IotConnectorFhirDestinationOutput{})
}
