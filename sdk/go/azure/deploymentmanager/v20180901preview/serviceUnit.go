


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ServiceUnit struct {
	pulumi.CustomResourceState

	Artifacts           ServiceUnitArtifactsResponsePtrOutput `pulumi:"artifacts"`
	DeploymentMode      pulumi.StringOutput                   `pulumi:"deploymentMode"`
	Location            pulumi.StringOutput                   `pulumi:"location"`
	Name                pulumi.StringOutput                   `pulumi:"name"`
	Tags                pulumi.StringMapOutput                `pulumi:"tags"`
	TargetResourceGroup pulumi.StringOutput                   `pulumi:"targetResourceGroup"`
	Type                pulumi.StringOutput                   `pulumi:"type"`
}


func NewServiceUnit(ctx *pulumi.Context,
	name string, args *ServiceUnitArgs, opts ...pulumi.ResourceOption) (*ServiceUnit, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeploymentMode == nil {
		return nil, errors.New("invalid value for required argument 'DeploymentMode'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.ServiceTopologyName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceTopologyName'")
	}
	if args.TargetResourceGroup == nil {
		return nil, errors.New("invalid value for required argument 'TargetResourceGroup'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:deploymentmanager:ServiceUnit"),
		},
		{
			Type: pulumi.String("azure-native:deploymentmanager/v20191101preview:ServiceUnit"),
		},
	})
	opts = append(opts, aliases)
	var resource ServiceUnit
	err := ctx.RegisterResource("azure-native:deploymentmanager/v20180901preview:ServiceUnit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServiceUnit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceUnitState, opts ...pulumi.ResourceOption) (*ServiceUnit, error) {
	var resource ServiceUnit
	err := ctx.ReadResource("azure-native:deploymentmanager/v20180901preview:ServiceUnit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serviceUnitState struct {
}

type ServiceUnitState struct {
}

func (ServiceUnitState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceUnitState)(nil)).Elem()
}

type serviceUnitArgs struct {
	Artifacts           *ServiceUnitArtifacts `pulumi:"artifacts"`
	DeploymentMode      DeploymentMode        `pulumi:"deploymentMode"`
	Location            *string               `pulumi:"location"`
	ResourceGroupName   string                `pulumi:"resourceGroupName"`
	ServiceName         string                `pulumi:"serviceName"`
	ServiceTopologyName string                `pulumi:"serviceTopologyName"`
	ServiceUnitName     *string               `pulumi:"serviceUnitName"`
	Tags                map[string]string     `pulumi:"tags"`
	TargetResourceGroup string                `pulumi:"targetResourceGroup"`
}


type ServiceUnitArgs struct {
	Artifacts           ServiceUnitArtifactsPtrInput
	DeploymentMode      DeploymentModeInput
	Location            pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	ServiceName         pulumi.StringInput
	ServiceTopologyName pulumi.StringInput
	ServiceUnitName     pulumi.StringPtrInput
	Tags                pulumi.StringMapInput
	TargetResourceGroup pulumi.StringInput
}

func (ServiceUnitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceUnitArgs)(nil)).Elem()
}

type ServiceUnitInput interface {
	pulumi.Input

	ToServiceUnitOutput() ServiceUnitOutput
	ToServiceUnitOutputWithContext(ctx context.Context) ServiceUnitOutput
}

func (*ServiceUnit) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceUnit)(nil)).Elem()
}

func (i *ServiceUnit) ToServiceUnitOutput() ServiceUnitOutput {
	return i.ToServiceUnitOutputWithContext(context.Background())
}

func (i *ServiceUnit) ToServiceUnitOutputWithContext(ctx context.Context) ServiceUnitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceUnitOutput)
}

type ServiceUnitOutput struct{ *pulumi.OutputState }

func (ServiceUnitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceUnit)(nil)).Elem()
}

func (o ServiceUnitOutput) ToServiceUnitOutput() ServiceUnitOutput {
	return o
}

func (o ServiceUnitOutput) ToServiceUnitOutputWithContext(ctx context.Context) ServiceUnitOutput {
	return o
}

func (o ServiceUnitOutput) Artifacts() ServiceUnitArtifactsResponsePtrOutput {
	return o.ApplyT(func(v *ServiceUnit) ServiceUnitArtifactsResponsePtrOutput { return v.Artifacts }).(ServiceUnitArtifactsResponsePtrOutput)
}

func (o ServiceUnitOutput) DeploymentMode() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceUnit) pulumi.StringOutput { return v.DeploymentMode }).(pulumi.StringOutput)
}

func (o ServiceUnitOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceUnit) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ServiceUnitOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceUnit) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceUnitOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceUnit) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ServiceUnitOutput) TargetResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceUnit) pulumi.StringOutput { return v.TargetResourceGroup }).(pulumi.StringOutput)
}

func (o ServiceUnitOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceUnit) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceUnitOutput{})
}
