


package devtestlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServiceRunner struct {
	pulumi.CustomResourceState

	Identity IdentityPropertiesResponsePtrOutput `pulumi:"identity"`
	Location pulumi.StringPtrOutput              `pulumi:"location"`
	Name     pulumi.StringOutput                 `pulumi:"name"`
	Tags     pulumi.StringMapOutput              `pulumi:"tags"`
	Type     pulumi.StringOutput                 `pulumi:"type"`
}


func NewServiceRunner(ctx *pulumi.Context,
	name string, args *ServiceRunnerArgs, opts ...pulumi.ResourceOption) (*ServiceRunner, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:ServiceRunner"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:ServiceRunner"),
		},
	})
	opts = append(opts, aliases)
	var resource ServiceRunner
	err := ctx.RegisterResource("azure-native:devtestlab:ServiceRunner", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServiceRunner(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceRunnerState, opts ...pulumi.ResourceOption) (*ServiceRunner, error) {
	var resource ServiceRunner
	err := ctx.ReadResource("azure-native:devtestlab:ServiceRunner", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serviceRunnerState struct {
}

type ServiceRunnerState struct {
}

func (ServiceRunnerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceRunnerState)(nil)).Elem()
}

type serviceRunnerArgs struct {
	Identity          *IdentityProperties `pulumi:"identity"`
	LabName           string              `pulumi:"labName"`
	Location          *string             `pulumi:"location"`
	Name              *string             `pulumi:"name"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	Tags              map[string]string   `pulumi:"tags"`
}


type ServiceRunnerArgs struct {
	Identity          IdentityPropertiesPtrInput
	LabName           pulumi.StringInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (ServiceRunnerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceRunnerArgs)(nil)).Elem()
}

type ServiceRunnerInput interface {
	pulumi.Input

	ToServiceRunnerOutput() ServiceRunnerOutput
	ToServiceRunnerOutputWithContext(ctx context.Context) ServiceRunnerOutput
}

func (*ServiceRunner) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceRunner)(nil)).Elem()
}

func (i *ServiceRunner) ToServiceRunnerOutput() ServiceRunnerOutput {
	return i.ToServiceRunnerOutputWithContext(context.Background())
}

func (i *ServiceRunner) ToServiceRunnerOutputWithContext(ctx context.Context) ServiceRunnerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceRunnerOutput)
}

type ServiceRunnerOutput struct{ *pulumi.OutputState }

func (ServiceRunnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceRunner)(nil)).Elem()
}

func (o ServiceRunnerOutput) ToServiceRunnerOutput() ServiceRunnerOutput {
	return o
}

func (o ServiceRunnerOutput) ToServiceRunnerOutputWithContext(ctx context.Context) ServiceRunnerOutput {
	return o
}

func (o ServiceRunnerOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ServiceRunner) IdentityPropertiesResponsePtrOutput { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o ServiceRunnerOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceRunner) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ServiceRunnerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceRunner) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceRunnerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceRunner) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ServiceRunnerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceRunner) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceRunnerOutput{})
}
