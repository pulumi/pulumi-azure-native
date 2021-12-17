


package v20180701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Application struct {
	pulumi.CustomResourceState

	DebugParams         pulumi.StringPtrOutput                        `pulumi:"debugParams"`
	Description         pulumi.StringPtrOutput                        `pulumi:"description"`
	Diagnostics         DiagnosticsDescriptionResponsePtrOutput       `pulumi:"diagnostics"`
	HealthState         pulumi.StringOutput                           `pulumi:"healthState"`
	Location            pulumi.StringOutput                           `pulumi:"location"`
	Name                pulumi.StringOutput                           `pulumi:"name"`
	ProvisioningState   pulumi.StringOutput                           `pulumi:"provisioningState"`
	ServiceNames        pulumi.StringArrayOutput                      `pulumi:"serviceNames"`
	Services            ServiceResourceDescriptionResponseArrayOutput `pulumi:"services"`
	Status              pulumi.StringOutput                           `pulumi:"status"`
	StatusDetails       pulumi.StringOutput                           `pulumi:"statusDetails"`
	Tags                pulumi.StringMapOutput                        `pulumi:"tags"`
	Type                pulumi.StringOutput                           `pulumi:"type"`
	UnhealthyEvaluation pulumi.StringOutput                           `pulumi:"unhealthyEvaluation"`
}


func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabricmesh:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabricmesh/v20180901preview:Application"),
		},
	})
	opts = append(opts, aliases)
	var resource Application
	err := ctx.RegisterResource("azure-native:servicefabricmesh/v20180701preview:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("azure-native:servicefabricmesh/v20180701preview:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type applicationState struct {
}

type ApplicationState struct {
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	ApplicationName   *string                      `pulumi:"applicationName"`
	DebugParams       *string                      `pulumi:"debugParams"`
	Description       *string                      `pulumi:"description"`
	Diagnostics       *DiagnosticsDescription      `pulumi:"diagnostics"`
	Location          *string                      `pulumi:"location"`
	ResourceGroupName string                       `pulumi:"resourceGroupName"`
	Services          []ServiceResourceDescription `pulumi:"services"`
	Tags              map[string]string            `pulumi:"tags"`
}


type ApplicationArgs struct {
	ApplicationName   pulumi.StringPtrInput
	DebugParams       pulumi.StringPtrInput
	Description       pulumi.StringPtrInput
	Diagnostics       DiagnosticsDescriptionPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Services          ServiceResourceDescriptionArrayInput
	Tags              pulumi.StringMapInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

type ApplicationOutput struct{ *pulumi.OutputState }

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApplicationOutput{})
}
