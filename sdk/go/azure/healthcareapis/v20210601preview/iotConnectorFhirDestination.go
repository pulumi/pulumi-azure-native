


package v20210601preview

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
			Type: pulumi.String("azure-nextgen:healthcareapis/v20210601preview:IotConnectorFhirDestination"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis:IotConnectorFhirDestination"),
		},
		{
			Type: pulumi.String("azure-nextgen:healthcareapis:IotConnectorFhirDestination"),
		},
	})
	opts = append(opts, aliases)
	var resource IotConnectorFhirDestination
	err := ctx.RegisterResource("azure-native:healthcareapis/v20210601preview:IotConnectorFhirDestination", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIotConnectorFhirDestination(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotConnectorFhirDestinationState, opts ...pulumi.ResourceOption) (*IotConnectorFhirDestination, error) {
	var resource IotConnectorFhirDestination
	err := ctx.ReadResource("azure-native:healthcareapis/v20210601preview:IotConnectorFhirDestination", name, id, state, &resource, opts...)
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
	Etag                           *string              `pulumi:"etag"`
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
	Etag                           pulumi.StringPtrInput
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
	return reflect.TypeOf((*IotConnectorFhirDestination)(nil))
}

func (i *IotConnectorFhirDestination) ToIotConnectorFhirDestinationOutput() IotConnectorFhirDestinationOutput {
	return i.ToIotConnectorFhirDestinationOutputWithContext(context.Background())
}

func (i *IotConnectorFhirDestination) ToIotConnectorFhirDestinationOutputWithContext(ctx context.Context) IotConnectorFhirDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotConnectorFhirDestinationOutput)
}

type IotConnectorFhirDestinationOutput struct{ *pulumi.OutputState }

func (IotConnectorFhirDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotConnectorFhirDestination)(nil))
}

func (o IotConnectorFhirDestinationOutput) ToIotConnectorFhirDestinationOutput() IotConnectorFhirDestinationOutput {
	return o
}

func (o IotConnectorFhirDestinationOutput) ToIotConnectorFhirDestinationOutputWithContext(ctx context.Context) IotConnectorFhirDestinationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IotConnectorFhirDestinationInput)(nil)).Elem(), &IotConnectorFhirDestination{})
	pulumi.RegisterOutputType(IotConnectorFhirDestinationOutput{})
}
