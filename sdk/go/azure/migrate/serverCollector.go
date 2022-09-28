


package migrate

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerCollector struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput            `pulumi:"eTag"`
	Name       pulumi.StringOutput               `pulumi:"name"`
	Properties CollectorPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput               `pulumi:"type"`
}


func NewServerCollector(ctx *pulumi.Context,
	name string, args *ServerCollectorArgs, opts ...pulumi.ResourceOption) (*ServerCollector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate/v20191001:ServerCollector"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerCollector
	err := ctx.RegisterResource("azure-native:migrate:ServerCollector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerCollector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerCollectorState, opts ...pulumi.ResourceOption) (*ServerCollector, error) {
	var resource ServerCollector
	err := ctx.ReadResource("azure-native:migrate:ServerCollector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serverCollectorState struct {
}

type ServerCollectorState struct {
}

func (ServerCollectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverCollectorState)(nil)).Elem()
}

type serverCollectorArgs struct {
	ETag                *string              `pulumi:"eTag"`
	ProjectName         string               `pulumi:"projectName"`
	Properties          *CollectorProperties `pulumi:"properties"`
	ResourceGroupName   string               `pulumi:"resourceGroupName"`
	ServerCollectorName *string              `pulumi:"serverCollectorName"`
}


type ServerCollectorArgs struct {
	ETag                pulumi.StringPtrInput
	ProjectName         pulumi.StringInput
	Properties          CollectorPropertiesPtrInput
	ResourceGroupName   pulumi.StringInput
	ServerCollectorName pulumi.StringPtrInput
}

func (ServerCollectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverCollectorArgs)(nil)).Elem()
}

type ServerCollectorInput interface {
	pulumi.Input

	ToServerCollectorOutput() ServerCollectorOutput
	ToServerCollectorOutputWithContext(ctx context.Context) ServerCollectorOutput
}

func (*ServerCollector) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerCollector)(nil)).Elem()
}

func (i *ServerCollector) ToServerCollectorOutput() ServerCollectorOutput {
	return i.ToServerCollectorOutputWithContext(context.Background())
}

func (i *ServerCollector) ToServerCollectorOutputWithContext(ctx context.Context) ServerCollectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCollectorOutput)
}

type ServerCollectorOutput struct{ *pulumi.OutputState }

func (ServerCollectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerCollector)(nil)).Elem()
}

func (o ServerCollectorOutput) ToServerCollectorOutput() ServerCollectorOutput {
	return o
}

func (o ServerCollectorOutput) ToServerCollectorOutputWithContext(ctx context.Context) ServerCollectorOutput {
	return o
}

func (o ServerCollectorOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerCollector) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o ServerCollectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerCollector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServerCollectorOutput) Properties() CollectorPropertiesResponseOutput {
	return o.ApplyT(func(v *ServerCollector) CollectorPropertiesResponseOutput { return v.Properties }).(CollectorPropertiesResponseOutput)
}

func (o ServerCollectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerCollector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerCollectorOutput{})
}
