


package v20190301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ApplicationTypeVersion struct {
	pulumi.CustomResourceState

	AppPackageUrl        pulumi.StringOutput    `pulumi:"appPackageUrl"`
	DefaultParameterList pulumi.StringMapOutput `pulumi:"defaultParameterList"`
	Etag                 pulumi.StringOutput    `pulumi:"etag"`
	Location             pulumi.StringPtrOutput `pulumi:"location"`
	Name                 pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState    pulumi.StringOutput    `pulumi:"provisioningState"`
	Tags                 pulumi.StringMapOutput `pulumi:"tags"`
	Type                 pulumi.StringOutput    `pulumi:"type"`
}


func NewApplicationTypeVersion(ctx *pulumi.Context,
	name string, args *ApplicationTypeVersionArgs, opts ...pulumi.ResourceOption) (*ApplicationTypeVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppPackageUrl == nil {
		return nil, errors.New("invalid value for required argument 'AppPackageUrl'")
	}
	if args.ApplicationTypeName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationTypeName'")
	}
	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabric:ApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20170701preview:ApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190301:ApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190601preview:ApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20191101preview:ApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20200301:ApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20201201preview:ApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210601:ApplicationTypeVersion"),
		},
	})
	opts = append(opts, aliases)
	var resource ApplicationTypeVersion
	err := ctx.RegisterResource("azure-native:servicefabric/v20190301preview:ApplicationTypeVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplicationTypeVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationTypeVersionState, opts ...pulumi.ResourceOption) (*ApplicationTypeVersion, error) {
	var resource ApplicationTypeVersion
	err := ctx.ReadResource("azure-native:servicefabric/v20190301preview:ApplicationTypeVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type applicationTypeVersionState struct {
}

type ApplicationTypeVersionState struct {
}

func (ApplicationTypeVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationTypeVersionState)(nil)).Elem()
}

type applicationTypeVersionArgs struct {
	AppPackageUrl       string            `pulumi:"appPackageUrl"`
	ApplicationTypeName string            `pulumi:"applicationTypeName"`
	ClusterName         string            `pulumi:"clusterName"`
	Location            *string           `pulumi:"location"`
	ResourceGroupName   string            `pulumi:"resourceGroupName"`
	Tags                map[string]string `pulumi:"tags"`
	Version             *string           `pulumi:"version"`
}


type ApplicationTypeVersionArgs struct {
	AppPackageUrl       pulumi.StringInput
	ApplicationTypeName pulumi.StringInput
	ClusterName         pulumi.StringInput
	Location            pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	Tags                pulumi.StringMapInput
	Version             pulumi.StringPtrInput
}

func (ApplicationTypeVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationTypeVersionArgs)(nil)).Elem()
}

type ApplicationTypeVersionInput interface {
	pulumi.Input

	ToApplicationTypeVersionOutput() ApplicationTypeVersionOutput
	ToApplicationTypeVersionOutputWithContext(ctx context.Context) ApplicationTypeVersionOutput
}

func (*ApplicationTypeVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationTypeVersion)(nil)).Elem()
}

func (i *ApplicationTypeVersion) ToApplicationTypeVersionOutput() ApplicationTypeVersionOutput {
	return i.ToApplicationTypeVersionOutputWithContext(context.Background())
}

func (i *ApplicationTypeVersion) ToApplicationTypeVersionOutputWithContext(ctx context.Context) ApplicationTypeVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTypeVersionOutput)
}

type ApplicationTypeVersionOutput struct{ *pulumi.OutputState }

func (ApplicationTypeVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationTypeVersion)(nil)).Elem()
}

func (o ApplicationTypeVersionOutput) ToApplicationTypeVersionOutput() ApplicationTypeVersionOutput {
	return o
}

func (o ApplicationTypeVersionOutput) ToApplicationTypeVersionOutputWithContext(ctx context.Context) ApplicationTypeVersionOutput {
	return o
}

func (o ApplicationTypeVersionOutput) AppPackageUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationTypeVersion) pulumi.StringOutput { return v.AppPackageUrl }).(pulumi.StringOutput)
}

func (o ApplicationTypeVersionOutput) DefaultParameterList() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApplicationTypeVersion) pulumi.StringMapOutput { return v.DefaultParameterList }).(pulumi.StringMapOutput)
}

func (o ApplicationTypeVersionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationTypeVersion) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ApplicationTypeVersionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationTypeVersion) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ApplicationTypeVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationTypeVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationTypeVersionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationTypeVersion) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ApplicationTypeVersionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApplicationTypeVersion) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ApplicationTypeVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationTypeVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationTypeVersionOutput{})
}
