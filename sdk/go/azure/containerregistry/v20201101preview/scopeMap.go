


package v20201101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ScopeMap struct {
	pulumi.CustomResourceState

	Actions           pulumi.StringArrayOutput `pulumi:"actions"`
	CreationDate      pulumi.StringOutput      `pulumi:"creationDate"`
	Description       pulumi.StringPtrOutput   `pulumi:"description"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewScopeMap(ctx *pulumi.Context,
	name string, args *ScopeMapArgs, opts ...pulumi.ResourceOption) (*ScopeMap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerregistry:ScopeMap"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20190501preview:ScopeMap"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210601preview:ScopeMap"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210801preview:ScopeMap"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20211201preview:ScopeMap"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20220201preview:ScopeMap"),
		},
	})
	opts = append(opts, aliases)
	var resource ScopeMap
	err := ctx.RegisterResource("azure-native:containerregistry/v20201101preview:ScopeMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetScopeMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScopeMapState, opts ...pulumi.ResourceOption) (*ScopeMap, error) {
	var resource ScopeMap
	err := ctx.ReadResource("azure-native:containerregistry/v20201101preview:ScopeMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type scopeMapState struct {
}

type ScopeMapState struct {
}

func (ScopeMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*scopeMapState)(nil)).Elem()
}

type scopeMapArgs struct {
	Actions           []string `pulumi:"actions"`
	Description       *string  `pulumi:"description"`
	RegistryName      string   `pulumi:"registryName"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	ScopeMapName      *string  `pulumi:"scopeMapName"`
}


type ScopeMapArgs struct {
	Actions           pulumi.StringArrayInput
	Description       pulumi.StringPtrInput
	RegistryName      pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ScopeMapName      pulumi.StringPtrInput
}

func (ScopeMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scopeMapArgs)(nil)).Elem()
}

type ScopeMapInput interface {
	pulumi.Input

	ToScopeMapOutput() ScopeMapOutput
	ToScopeMapOutputWithContext(ctx context.Context) ScopeMapOutput
}

func (*ScopeMap) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeMap)(nil)).Elem()
}

func (i *ScopeMap) ToScopeMapOutput() ScopeMapOutput {
	return i.ToScopeMapOutputWithContext(context.Background())
}

func (i *ScopeMap) ToScopeMapOutputWithContext(ctx context.Context) ScopeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeMapOutput)
}

type ScopeMapOutput struct{ *pulumi.OutputState }

func (ScopeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeMap)(nil)).Elem()
}

func (o ScopeMapOutput) ToScopeMapOutput() ScopeMapOutput {
	return o
}

func (o ScopeMapOutput) ToScopeMapOutputWithContext(ctx context.Context) ScopeMapOutput {
	return o
}

func (o ScopeMapOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScopeMap) pulumi.StringArrayOutput { return v.Actions }).(pulumi.StringArrayOutput)
}

func (o ScopeMapOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeMap) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

func (o ScopeMapOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeMap) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ScopeMapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeMap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ScopeMapOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeMap) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ScopeMapOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ScopeMap) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ScopeMapOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeMap) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScopeMapOutput{})
}
