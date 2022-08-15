


package customerinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Profile struct {
	pulumi.CustomResourceState

	ApiEntitySetName    pulumi.StringPtrOutput                `pulumi:"apiEntitySetName"`
	Attributes          pulumi.StringArrayMapOutput           `pulumi:"attributes"`
	Description         pulumi.StringMapOutput                `pulumi:"description"`
	DisplayName         pulumi.StringMapOutput                `pulumi:"displayName"`
	EntityType          pulumi.StringPtrOutput                `pulumi:"entityType"`
	Fields              PropertyDefinitionResponseArrayOutput `pulumi:"fields"`
	InstancesCount      pulumi.IntPtrOutput                   `pulumi:"instancesCount"`
	LargeImage          pulumi.StringPtrOutput                `pulumi:"largeImage"`
	LastChangedUtc      pulumi.StringOutput                   `pulumi:"lastChangedUtc"`
	LocalizedAttributes pulumi.StringMapMapOutput             `pulumi:"localizedAttributes"`
	MediumImage         pulumi.StringPtrOutput                `pulumi:"mediumImage"`
	Name                pulumi.StringOutput                   `pulumi:"name"`
	ProvisioningState   pulumi.StringOutput                   `pulumi:"provisioningState"`
	SchemaItemTypeLink  pulumi.StringPtrOutput                `pulumi:"schemaItemTypeLink"`
	SmallImage          pulumi.StringPtrOutput                `pulumi:"smallImage"`
	StrongIds           StrongIdResponseArrayOutput           `pulumi:"strongIds"`
	TenantId            pulumi.StringOutput                   `pulumi:"tenantId"`
	TimestampFieldName  pulumi.StringPtrOutput                `pulumi:"timestampFieldName"`
	Type                pulumi.StringOutput                   `pulumi:"type"`
	TypeName            pulumi.StringPtrOutput                `pulumi:"typeName"`
}


func NewProfile(ctx *pulumi.Context,
	name string, args *ProfileArgs, opts ...pulumi.ResourceOption) (*Profile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HubName == nil {
		return nil, errors.New("invalid value for required argument 'HubName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:customerinsights/v20170101:Profile"),
		},
		{
			Type: pulumi.String("azure-native:customerinsights/v20170426:Profile"),
		},
	})
	opts = append(opts, aliases)
	var resource Profile
	err := ctx.RegisterResource("azure-native:customerinsights:Profile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileState, opts ...pulumi.ResourceOption) (*Profile, error) {
	var resource Profile
	err := ctx.ReadResource("azure-native:customerinsights:Profile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type profileState struct {
}

type ProfileState struct {
}

func (ProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileState)(nil)).Elem()
}

type profileArgs struct {
	ApiEntitySetName    *string                      `pulumi:"apiEntitySetName"`
	Attributes          map[string][]string          `pulumi:"attributes"`
	Description         map[string]string            `pulumi:"description"`
	DisplayName         map[string]string            `pulumi:"displayName"`
	EntityType          *EntityTypes                 `pulumi:"entityType"`
	Fields              []PropertyDefinition         `pulumi:"fields"`
	HubName             string                       `pulumi:"hubName"`
	InstancesCount      *int                         `pulumi:"instancesCount"`
	LargeImage          *string                      `pulumi:"largeImage"`
	LocalizedAttributes map[string]map[string]string `pulumi:"localizedAttributes"`
	MediumImage         *string                      `pulumi:"mediumImage"`
	ProfileName         *string                      `pulumi:"profileName"`
	ResourceGroupName   string                       `pulumi:"resourceGroupName"`
	SchemaItemTypeLink  *string                      `pulumi:"schemaItemTypeLink"`
	SmallImage          *string                      `pulumi:"smallImage"`
	StrongIds           []StrongId                   `pulumi:"strongIds"`
	TimestampFieldName  *string                      `pulumi:"timestampFieldName"`
	TypeName            *string                      `pulumi:"typeName"`
}


type ProfileArgs struct {
	ApiEntitySetName    pulumi.StringPtrInput
	Attributes          pulumi.StringArrayMapInput
	Description         pulumi.StringMapInput
	DisplayName         pulumi.StringMapInput
	EntityType          EntityTypesPtrInput
	Fields              PropertyDefinitionArrayInput
	HubName             pulumi.StringInput
	InstancesCount      pulumi.IntPtrInput
	LargeImage          pulumi.StringPtrInput
	LocalizedAttributes pulumi.StringMapMapInput
	MediumImage         pulumi.StringPtrInput
	ProfileName         pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	SchemaItemTypeLink  pulumi.StringPtrInput
	SmallImage          pulumi.StringPtrInput
	StrongIds           StrongIdArrayInput
	TimestampFieldName  pulumi.StringPtrInput
	TypeName            pulumi.StringPtrInput
}

func (ProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileArgs)(nil)).Elem()
}

type ProfileInput interface {
	pulumi.Input

	ToProfileOutput() ProfileOutput
	ToProfileOutputWithContext(ctx context.Context) ProfileOutput
}

func (*Profile) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil)).Elem()
}

func (i *Profile) ToProfileOutput() ProfileOutput {
	return i.ToProfileOutputWithContext(context.Background())
}

func (i *Profile) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileOutput)
}

type ProfileOutput struct{ *pulumi.OutputState }

func (ProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil)).Elem()
}

func (o ProfileOutput) ToProfileOutput() ProfileOutput {
	return o
}

func (o ProfileOutput) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return o
}

func (o ProfileOutput) ApiEntitySetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.ApiEntitySetName }).(pulumi.StringPtrOutput)
}

func (o ProfileOutput) Attributes() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringArrayMapOutput { return v.Attributes }).(pulumi.StringArrayMapOutput)
}

func (o ProfileOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringMapOutput { return v.Description }).(pulumi.StringMapOutput)
}

func (o ProfileOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringMapOutput { return v.DisplayName }).(pulumi.StringMapOutput)
}

func (o ProfileOutput) EntityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.EntityType }).(pulumi.StringPtrOutput)
}

func (o ProfileOutput) Fields() PropertyDefinitionResponseArrayOutput {
	return o.ApplyT(func(v *Profile) PropertyDefinitionResponseArrayOutput { return v.Fields }).(PropertyDefinitionResponseArrayOutput)
}

func (o ProfileOutput) InstancesCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.IntPtrOutput { return v.InstancesCount }).(pulumi.IntPtrOutput)
}

func (o ProfileOutput) LargeImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.LargeImage }).(pulumi.StringPtrOutput)
}

func (o ProfileOutput) LastChangedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.LastChangedUtc }).(pulumi.StringOutput)
}

func (o ProfileOutput) LocalizedAttributes() pulumi.StringMapMapOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringMapMapOutput { return v.LocalizedAttributes }).(pulumi.StringMapMapOutput)
}

func (o ProfileOutput) MediumImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.MediumImage }).(pulumi.StringPtrOutput)
}

func (o ProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProfileOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ProfileOutput) SchemaItemTypeLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.SchemaItemTypeLink }).(pulumi.StringPtrOutput)
}

func (o ProfileOutput) SmallImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.SmallImage }).(pulumi.StringPtrOutput)
}

func (o ProfileOutput) StrongIds() StrongIdResponseArrayOutput {
	return o.ApplyT(func(v *Profile) StrongIdResponseArrayOutput { return v.StrongIds }).(StrongIdResponseArrayOutput)
}

func (o ProfileOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o ProfileOutput) TimestampFieldName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.TimestampFieldName }).(pulumi.StringPtrOutput)
}

func (o ProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ProfileOutput) TypeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.TypeName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ProfileOutput{})
}
