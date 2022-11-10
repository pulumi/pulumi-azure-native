


package synapse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KustoPoolAttachedDatabaseConfiguration struct {
	pulumi.CustomResourceState

	AttachedDatabaseNames             pulumi.StringArrayOutput                     `pulumi:"attachedDatabaseNames"`
	DatabaseName                      pulumi.StringOutput                          `pulumi:"databaseName"`
	DefaultPrincipalsModificationKind pulumi.StringOutput                          `pulumi:"defaultPrincipalsModificationKind"`
	KustoPoolResourceId               pulumi.StringOutput                          `pulumi:"kustoPoolResourceId"`
	Location                          pulumi.StringPtrOutput                       `pulumi:"location"`
	Name                              pulumi.StringOutput                          `pulumi:"name"`
	ProvisioningState                 pulumi.StringOutput                          `pulumi:"provisioningState"`
	SystemData                        SystemDataResponseOutput                     `pulumi:"systemData"`
	TableLevelSharingProperties       TableLevelSharingPropertiesResponsePtrOutput `pulumi:"tableLevelSharingProperties"`
	Type                              pulumi.StringOutput                          `pulumi:"type"`
}


func NewKustoPoolAttachedDatabaseConfiguration(ctx *pulumi.Context,
	name string, args *KustoPoolAttachedDatabaseConfigurationArgs, opts ...pulumi.ResourceOption) (*KustoPoolAttachedDatabaseConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.DefaultPrincipalsModificationKind == nil {
		return nil, errors.New("invalid value for required argument 'DefaultPrincipalsModificationKind'")
	}
	if args.KustoPoolName == nil {
		return nil, errors.New("invalid value for required argument 'KustoPoolName'")
	}
	if args.KustoPoolResourceId == nil {
		return nil, errors.New("invalid value for required argument 'KustoPoolResourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:KustoPoolAttachedDatabaseConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource KustoPoolAttachedDatabaseConfiguration
	err := ctx.RegisterResource("azure-native:synapse:KustoPoolAttachedDatabaseConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetKustoPoolAttachedDatabaseConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KustoPoolAttachedDatabaseConfigurationState, opts ...pulumi.ResourceOption) (*KustoPoolAttachedDatabaseConfiguration, error) {
	var resource KustoPoolAttachedDatabaseConfiguration
	err := ctx.ReadResource("azure-native:synapse:KustoPoolAttachedDatabaseConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type kustoPoolAttachedDatabaseConfigurationState struct {
}

type KustoPoolAttachedDatabaseConfigurationState struct {
}

func (KustoPoolAttachedDatabaseConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoPoolAttachedDatabaseConfigurationState)(nil)).Elem()
}

type kustoPoolAttachedDatabaseConfigurationArgs struct {
	AttachedDatabaseConfigurationName *string                      `pulumi:"attachedDatabaseConfigurationName"`
	DatabaseName                      string                       `pulumi:"databaseName"`
	DefaultPrincipalsModificationKind string                       `pulumi:"defaultPrincipalsModificationKind"`
	KustoPoolName                     string                       `pulumi:"kustoPoolName"`
	KustoPoolResourceId               string                       `pulumi:"kustoPoolResourceId"`
	Location                          *string                      `pulumi:"location"`
	ResourceGroupName                 string                       `pulumi:"resourceGroupName"`
	TableLevelSharingProperties       *TableLevelSharingProperties `pulumi:"tableLevelSharingProperties"`
	WorkspaceName                     string                       `pulumi:"workspaceName"`
}


type KustoPoolAttachedDatabaseConfigurationArgs struct {
	AttachedDatabaseConfigurationName pulumi.StringPtrInput
	DatabaseName                      pulumi.StringInput
	DefaultPrincipalsModificationKind pulumi.StringInput
	KustoPoolName                     pulumi.StringInput
	KustoPoolResourceId               pulumi.StringInput
	Location                          pulumi.StringPtrInput
	ResourceGroupName                 pulumi.StringInput
	TableLevelSharingProperties       TableLevelSharingPropertiesPtrInput
	WorkspaceName                     pulumi.StringInput
}

func (KustoPoolAttachedDatabaseConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoPoolAttachedDatabaseConfigurationArgs)(nil)).Elem()
}

type KustoPoolAttachedDatabaseConfigurationInput interface {
	pulumi.Input

	ToKustoPoolAttachedDatabaseConfigurationOutput() KustoPoolAttachedDatabaseConfigurationOutput
	ToKustoPoolAttachedDatabaseConfigurationOutputWithContext(ctx context.Context) KustoPoolAttachedDatabaseConfigurationOutput
}

func (*KustoPoolAttachedDatabaseConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoPoolAttachedDatabaseConfiguration)(nil)).Elem()
}

func (i *KustoPoolAttachedDatabaseConfiguration) ToKustoPoolAttachedDatabaseConfigurationOutput() KustoPoolAttachedDatabaseConfigurationOutput {
	return i.ToKustoPoolAttachedDatabaseConfigurationOutputWithContext(context.Background())
}

func (i *KustoPoolAttachedDatabaseConfiguration) ToKustoPoolAttachedDatabaseConfigurationOutputWithContext(ctx context.Context) KustoPoolAttachedDatabaseConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustoPoolAttachedDatabaseConfigurationOutput)
}

type KustoPoolAttachedDatabaseConfigurationOutput struct{ *pulumi.OutputState }

func (KustoPoolAttachedDatabaseConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoPoolAttachedDatabaseConfiguration)(nil)).Elem()
}

func (o KustoPoolAttachedDatabaseConfigurationOutput) ToKustoPoolAttachedDatabaseConfigurationOutput() KustoPoolAttachedDatabaseConfigurationOutput {
	return o
}

func (o KustoPoolAttachedDatabaseConfigurationOutput) ToKustoPoolAttachedDatabaseConfigurationOutputWithContext(ctx context.Context) KustoPoolAttachedDatabaseConfigurationOutput {
	return o
}

func (o KustoPoolAttachedDatabaseConfigurationOutput) AttachedDatabaseNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KustoPoolAttachedDatabaseConfiguration) pulumi.StringArrayOutput {
		return v.AttachedDatabaseNames
	}).(pulumi.StringArrayOutput)
}

func (o KustoPoolAttachedDatabaseConfigurationOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolAttachedDatabaseConfiguration) pulumi.StringOutput { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o KustoPoolAttachedDatabaseConfigurationOutput) DefaultPrincipalsModificationKind() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolAttachedDatabaseConfiguration) pulumi.StringOutput {
		return v.DefaultPrincipalsModificationKind
	}).(pulumi.StringOutput)
}

func (o KustoPoolAttachedDatabaseConfigurationOutput) KustoPoolResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolAttachedDatabaseConfiguration) pulumi.StringOutput { return v.KustoPoolResourceId }).(pulumi.StringOutput)
}

func (o KustoPoolAttachedDatabaseConfigurationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KustoPoolAttachedDatabaseConfiguration) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o KustoPoolAttachedDatabaseConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolAttachedDatabaseConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o KustoPoolAttachedDatabaseConfigurationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolAttachedDatabaseConfiguration) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o KustoPoolAttachedDatabaseConfigurationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *KustoPoolAttachedDatabaseConfiguration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o KustoPoolAttachedDatabaseConfigurationOutput) TableLevelSharingProperties() TableLevelSharingPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *KustoPoolAttachedDatabaseConfiguration) TableLevelSharingPropertiesResponsePtrOutput {
		return v.TableLevelSharingProperties
	}).(TableLevelSharingPropertiesResponsePtrOutput)
}

func (o KustoPoolAttachedDatabaseConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolAttachedDatabaseConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(KustoPoolAttachedDatabaseConfigurationOutput{})
}
