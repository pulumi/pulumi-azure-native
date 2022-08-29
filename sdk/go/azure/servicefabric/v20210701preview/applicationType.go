


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationType struct {
	pulumi.CustomResourceState

	Location          pulumi.StringPtrOutput   `pulumi:"location"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput   `pulumi:"tags"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewApplicationType(ctx *pulumi.Context,
	name string, args *ApplicationTypeArgs, opts ...pulumi.ResourceOption) (*ApplicationType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabric/v20210101preview:ApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210501:ApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210901privatepreview:ApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20211101preview:ApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220101:ApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220201preview:ApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220601preview:ApplicationType"),
		},
	})
	opts = append(opts, aliases)
	var resource ApplicationType
	err := ctx.RegisterResource("azure-native:servicefabric/v20210701preview:ApplicationType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplicationType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationTypeState, opts ...pulumi.ResourceOption) (*ApplicationType, error) {
	var resource ApplicationType
	err := ctx.ReadResource("azure-native:servicefabric/v20210701preview:ApplicationType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type applicationTypeState struct {
}

type ApplicationTypeState struct {
}

func (ApplicationTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationTypeState)(nil)).Elem()
}

type applicationTypeArgs struct {
	ApplicationTypeName *string           `pulumi:"applicationTypeName"`
	ClusterName         string            `pulumi:"clusterName"`
	Location            *string           `pulumi:"location"`
	ResourceGroupName   string            `pulumi:"resourceGroupName"`
	Tags                map[string]string `pulumi:"tags"`
}


type ApplicationTypeArgs struct {
	ApplicationTypeName pulumi.StringPtrInput
	ClusterName         pulumi.StringInput
	Location            pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	Tags                pulumi.StringMapInput
}

func (ApplicationTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationTypeArgs)(nil)).Elem()
}

type ApplicationTypeInput interface {
	pulumi.Input

	ToApplicationTypeOutput() ApplicationTypeOutput
	ToApplicationTypeOutputWithContext(ctx context.Context) ApplicationTypeOutput
}

func (*ApplicationType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationType)(nil)).Elem()
}

func (i *ApplicationType) ToApplicationTypeOutput() ApplicationTypeOutput {
	return i.ToApplicationTypeOutputWithContext(context.Background())
}

func (i *ApplicationType) ToApplicationTypeOutputWithContext(ctx context.Context) ApplicationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTypeOutput)
}

type ApplicationTypeOutput struct{ *pulumi.OutputState }

func (ApplicationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationType)(nil)).Elem()
}

func (o ApplicationTypeOutput) ToApplicationTypeOutput() ApplicationTypeOutput {
	return o
}

func (o ApplicationTypeOutput) ToApplicationTypeOutputWithContext(ctx context.Context) ApplicationTypeOutput {
	return o
}

func (o ApplicationTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationType) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ApplicationTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationType) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationTypeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationType) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ApplicationTypeOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ApplicationType) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ApplicationTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApplicationType) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ApplicationTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationType) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationTypeOutput{})
}
