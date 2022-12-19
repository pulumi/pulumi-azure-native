


package v20220301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StaticSiteLinkedBackendForBuild struct {
	pulumi.CustomResourceState

	BackendResourceId pulumi.StringPtrOutput `pulumi:"backendResourceId"`
	CreatedOn         pulumi.StringOutput    `pulumi:"createdOn"`
	Kind              pulumi.StringPtrOutput `pulumi:"kind"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	Region            pulumi.StringPtrOutput `pulumi:"region"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewStaticSiteLinkedBackendForBuild(ctx *pulumi.Context,
	name string, args *StaticSiteLinkedBackendForBuildArgs, opts ...pulumi.ResourceOption) (*StaticSiteLinkedBackendForBuild, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentName'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:StaticSiteLinkedBackendForBuild"),
		},
	})
	opts = append(opts, aliases)
	var resource StaticSiteLinkedBackendForBuild
	err := ctx.RegisterResource("azure-native:web/v20220301:StaticSiteLinkedBackendForBuild", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStaticSiteLinkedBackendForBuild(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticSiteLinkedBackendForBuildState, opts ...pulumi.ResourceOption) (*StaticSiteLinkedBackendForBuild, error) {
	var resource StaticSiteLinkedBackendForBuild
	err := ctx.ReadResource("azure-native:web/v20220301:StaticSiteLinkedBackendForBuild", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type staticSiteLinkedBackendForBuildState struct {
}

type StaticSiteLinkedBackendForBuildState struct {
}

func (StaticSiteLinkedBackendForBuildState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteLinkedBackendForBuildState)(nil)).Elem()
}

type staticSiteLinkedBackendForBuildArgs struct {
	BackendResourceId *string `pulumi:"backendResourceId"`
	EnvironmentName   string  `pulumi:"environmentName"`
	Kind              *string `pulumi:"kind"`
	LinkedBackendName *string `pulumi:"linkedBackendName"`
	Name              string  `pulumi:"name"`
	Region            *string `pulumi:"region"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type StaticSiteLinkedBackendForBuildArgs struct {
	BackendResourceId pulumi.StringPtrInput
	EnvironmentName   pulumi.StringInput
	Kind              pulumi.StringPtrInput
	LinkedBackendName pulumi.StringPtrInput
	Name              pulumi.StringInput
	Region            pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (StaticSiteLinkedBackendForBuildArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteLinkedBackendForBuildArgs)(nil)).Elem()
}

type StaticSiteLinkedBackendForBuildInput interface {
	pulumi.Input

	ToStaticSiteLinkedBackendForBuildOutput() StaticSiteLinkedBackendForBuildOutput
	ToStaticSiteLinkedBackendForBuildOutputWithContext(ctx context.Context) StaticSiteLinkedBackendForBuildOutput
}

func (*StaticSiteLinkedBackendForBuild) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteLinkedBackendForBuild)(nil)).Elem()
}

func (i *StaticSiteLinkedBackendForBuild) ToStaticSiteLinkedBackendForBuildOutput() StaticSiteLinkedBackendForBuildOutput {
	return i.ToStaticSiteLinkedBackendForBuildOutputWithContext(context.Background())
}

func (i *StaticSiteLinkedBackendForBuild) ToStaticSiteLinkedBackendForBuildOutputWithContext(ctx context.Context) StaticSiteLinkedBackendForBuildOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteLinkedBackendForBuildOutput)
}

type StaticSiteLinkedBackendForBuildOutput struct{ *pulumi.OutputState }

func (StaticSiteLinkedBackendForBuildOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteLinkedBackendForBuild)(nil)).Elem()
}

func (o StaticSiteLinkedBackendForBuildOutput) ToStaticSiteLinkedBackendForBuildOutput() StaticSiteLinkedBackendForBuildOutput {
	return o
}

func (o StaticSiteLinkedBackendForBuildOutput) ToStaticSiteLinkedBackendForBuildOutputWithContext(ctx context.Context) StaticSiteLinkedBackendForBuildOutput {
	return o
}

func (o StaticSiteLinkedBackendForBuildOutput) BackendResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteLinkedBackendForBuild) pulumi.StringPtrOutput { return v.BackendResourceId }).(pulumi.StringPtrOutput)
}

func (o StaticSiteLinkedBackendForBuildOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSiteLinkedBackendForBuild) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o StaticSiteLinkedBackendForBuildOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteLinkedBackendForBuild) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o StaticSiteLinkedBackendForBuildOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSiteLinkedBackendForBuild) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StaticSiteLinkedBackendForBuildOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSiteLinkedBackendForBuild) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o StaticSiteLinkedBackendForBuildOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteLinkedBackendForBuild) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

func (o StaticSiteLinkedBackendForBuildOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSiteLinkedBackendForBuild) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StaticSiteLinkedBackendForBuildOutput{})
}
