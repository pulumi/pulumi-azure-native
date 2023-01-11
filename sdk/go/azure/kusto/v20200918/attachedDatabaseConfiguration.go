


package v20200918

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AttachedDatabaseConfiguration struct {
	pulumi.CustomResourceState

	AttachedDatabaseNames             pulumi.StringArrayOutput `pulumi:"attachedDatabaseNames"`
	ClusterResourceId                 pulumi.StringOutput      `pulumi:"clusterResourceId"`
	DatabaseName                      pulumi.StringOutput      `pulumi:"databaseName"`
	DefaultPrincipalsModificationKind pulumi.StringOutput      `pulumi:"defaultPrincipalsModificationKind"`
	Location                          pulumi.StringPtrOutput   `pulumi:"location"`
	Name                              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState                 pulumi.StringOutput      `pulumi:"provisioningState"`
	Type                              pulumi.StringOutput      `pulumi:"type"`
}


func NewAttachedDatabaseConfiguration(ctx *pulumi.Context,
	name string, args *AttachedDatabaseConfigurationArgs, opts ...pulumi.ResourceOption) (*AttachedDatabaseConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ClusterResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterResourceId'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.DefaultPrincipalsModificationKind == nil {
		return nil, errors.New("invalid value for required argument 'DefaultPrincipalsModificationKind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190907:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20191109:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200215:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200614:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210827:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220201:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220707:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20221111:AttachedDatabaseConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource AttachedDatabaseConfiguration
	err := ctx.RegisterResource("azure-native:kusto/v20200918:AttachedDatabaseConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAttachedDatabaseConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttachedDatabaseConfigurationState, opts ...pulumi.ResourceOption) (*AttachedDatabaseConfiguration, error) {
	var resource AttachedDatabaseConfiguration
	err := ctx.ReadResource("azure-native:kusto/v20200918:AttachedDatabaseConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type attachedDatabaseConfigurationState struct {
}

type AttachedDatabaseConfigurationState struct {
}

func (AttachedDatabaseConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedDatabaseConfigurationState)(nil)).Elem()
}

type attachedDatabaseConfigurationArgs struct {
	AttachedDatabaseConfigurationName *string `pulumi:"attachedDatabaseConfigurationName"`
	ClusterName                       string  `pulumi:"clusterName"`
	ClusterResourceId                 string  `pulumi:"clusterResourceId"`
	DatabaseName                      string  `pulumi:"databaseName"`
	DefaultPrincipalsModificationKind string  `pulumi:"defaultPrincipalsModificationKind"`
	Location                          *string `pulumi:"location"`
	ResourceGroupName                 string  `pulumi:"resourceGroupName"`
}


type AttachedDatabaseConfigurationArgs struct {
	AttachedDatabaseConfigurationName pulumi.StringPtrInput
	ClusterName                       pulumi.StringInput
	ClusterResourceId                 pulumi.StringInput
	DatabaseName                      pulumi.StringInput
	DefaultPrincipalsModificationKind pulumi.StringInput
	Location                          pulumi.StringPtrInput
	ResourceGroupName                 pulumi.StringInput
}

func (AttachedDatabaseConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedDatabaseConfigurationArgs)(nil)).Elem()
}

type AttachedDatabaseConfigurationInput interface {
	pulumi.Input

	ToAttachedDatabaseConfigurationOutput() AttachedDatabaseConfigurationOutput
	ToAttachedDatabaseConfigurationOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationOutput
}

func (*AttachedDatabaseConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedDatabaseConfiguration)(nil)).Elem()
}

func (i *AttachedDatabaseConfiguration) ToAttachedDatabaseConfigurationOutput() AttachedDatabaseConfigurationOutput {
	return i.ToAttachedDatabaseConfigurationOutputWithContext(context.Background())
}

func (i *AttachedDatabaseConfiguration) ToAttachedDatabaseConfigurationOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedDatabaseConfigurationOutput)
}

type AttachedDatabaseConfigurationOutput struct{ *pulumi.OutputState }

func (AttachedDatabaseConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedDatabaseConfiguration)(nil)).Elem()
}

func (o AttachedDatabaseConfigurationOutput) ToAttachedDatabaseConfigurationOutput() AttachedDatabaseConfigurationOutput {
	return o
}

func (o AttachedDatabaseConfigurationOutput) ToAttachedDatabaseConfigurationOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationOutput {
	return o
}

func (o AttachedDatabaseConfigurationOutput) AttachedDatabaseNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AttachedDatabaseConfiguration) pulumi.StringArrayOutput { return v.AttachedDatabaseNames }).(pulumi.StringArrayOutput)
}

func (o AttachedDatabaseConfigurationOutput) ClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDatabaseConfiguration) pulumi.StringOutput { return v.ClusterResourceId }).(pulumi.StringOutput)
}

func (o AttachedDatabaseConfigurationOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDatabaseConfiguration) pulumi.StringOutput { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o AttachedDatabaseConfigurationOutput) DefaultPrincipalsModificationKind() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDatabaseConfiguration) pulumi.StringOutput { return v.DefaultPrincipalsModificationKind }).(pulumi.StringOutput)
}

func (o AttachedDatabaseConfigurationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttachedDatabaseConfiguration) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o AttachedDatabaseConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDatabaseConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AttachedDatabaseConfigurationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDatabaseConfiguration) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AttachedDatabaseConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDatabaseConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AttachedDatabaseConfigurationOutput{})
}
