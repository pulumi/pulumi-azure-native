


package v20160901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ApplianceDefinition struct {
	pulumi.CustomResourceState

	Artifacts      ApplianceArtifactResponseArrayOutput              `pulumi:"artifacts"`
	Authorizations ApplianceProviderAuthorizationResponseArrayOutput `pulumi:"authorizations"`
	Description    pulumi.StringPtrOutput                            `pulumi:"description"`
	DisplayName    pulumi.StringPtrOutput                            `pulumi:"displayName"`
	Identity       IdentityResponsePtrOutput                         `pulumi:"identity"`
	Location       pulumi.StringPtrOutput                            `pulumi:"location"`
	LockLevel      pulumi.StringOutput                               `pulumi:"lockLevel"`
	ManagedBy      pulumi.StringPtrOutput                            `pulumi:"managedBy"`
	Name           pulumi.StringOutput                               `pulumi:"name"`
	PackageFileUri pulumi.StringOutput                               `pulumi:"packageFileUri"`
	Sku            SkuResponsePtrOutput                              `pulumi:"sku"`
	Tags           pulumi.StringMapOutput                            `pulumi:"tags"`
	Type           pulumi.StringOutput                               `pulumi:"type"`
}


func NewApplianceDefinition(ctx *pulumi.Context,
	name string, args *ApplianceDefinitionArgs, opts ...pulumi.ResourceOption) (*ApplianceDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Authorizations == nil {
		return nil, errors.New("invalid value for required argument 'Authorizations'")
	}
	if args.LockLevel == nil {
		return nil, errors.New("invalid value for required argument 'LockLevel'")
	}
	if args.PackageFileUri == nil {
		return nil, errors.New("invalid value for required argument 'PackageFileUri'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:solutions:ApplianceDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20170901:ApplianceDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20171201:ApplianceDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20180201:ApplianceDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20180301:ApplianceDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20180601:ApplianceDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20180901preview:ApplianceDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20190701:ApplianceDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20200821preview:ApplianceDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20210201preview:ApplianceDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20210701:ApplianceDefinition"),
		},
	})
	opts = append(opts, aliases)
	var resource ApplianceDefinition
	err := ctx.RegisterResource("azure-native:solutions/v20160901preview:ApplianceDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplianceDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplianceDefinitionState, opts ...pulumi.ResourceOption) (*ApplianceDefinition, error) {
	var resource ApplianceDefinition
	err := ctx.ReadResource("azure-native:solutions/v20160901preview:ApplianceDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type applianceDefinitionState struct {
}

type ApplianceDefinitionState struct {
}

func (ApplianceDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*applianceDefinitionState)(nil)).Elem()
}

type applianceDefinitionArgs struct {
	ApplianceDefinitionName *string                          `pulumi:"applianceDefinitionName"`
	Artifacts               []ApplianceArtifact              `pulumi:"artifacts"`
	Authorizations          []ApplianceProviderAuthorization `pulumi:"authorizations"`
	Description             *string                          `pulumi:"description"`
	DisplayName             *string                          `pulumi:"displayName"`
	Identity                *Identity                        `pulumi:"identity"`
	Location                *string                          `pulumi:"location"`
	LockLevel               ApplianceLockLevel               `pulumi:"lockLevel"`
	ManagedBy               *string                          `pulumi:"managedBy"`
	PackageFileUri          string                           `pulumi:"packageFileUri"`
	ResourceGroupName       string                           `pulumi:"resourceGroupName"`
	Sku                     *Sku                             `pulumi:"sku"`
	Tags                    map[string]string                `pulumi:"tags"`
}


type ApplianceDefinitionArgs struct {
	ApplianceDefinitionName pulumi.StringPtrInput
	Artifacts               ApplianceArtifactArrayInput
	Authorizations          ApplianceProviderAuthorizationArrayInput
	Description             pulumi.StringPtrInput
	DisplayName             pulumi.StringPtrInput
	Identity                IdentityPtrInput
	Location                pulumi.StringPtrInput
	LockLevel               ApplianceLockLevelInput
	ManagedBy               pulumi.StringPtrInput
	PackageFileUri          pulumi.StringInput
	ResourceGroupName       pulumi.StringInput
	Sku                     SkuPtrInput
	Tags                    pulumi.StringMapInput
}

func (ApplianceDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applianceDefinitionArgs)(nil)).Elem()
}

type ApplianceDefinitionInput interface {
	pulumi.Input

	ToApplianceDefinitionOutput() ApplianceDefinitionOutput
	ToApplianceDefinitionOutputWithContext(ctx context.Context) ApplianceDefinitionOutput
}

func (*ApplianceDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplianceDefinition)(nil)).Elem()
}

func (i *ApplianceDefinition) ToApplianceDefinitionOutput() ApplianceDefinitionOutput {
	return i.ToApplianceDefinitionOutputWithContext(context.Background())
}

func (i *ApplianceDefinition) ToApplianceDefinitionOutputWithContext(ctx context.Context) ApplianceDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplianceDefinitionOutput)
}

type ApplianceDefinitionOutput struct{ *pulumi.OutputState }

func (ApplianceDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplianceDefinition)(nil)).Elem()
}

func (o ApplianceDefinitionOutput) ToApplianceDefinitionOutput() ApplianceDefinitionOutput {
	return o
}

func (o ApplianceDefinitionOutput) ToApplianceDefinitionOutputWithContext(ctx context.Context) ApplianceDefinitionOutput {
	return o
}

func (o ApplianceDefinitionOutput) Artifacts() ApplianceArtifactResponseArrayOutput {
	return o.ApplyT(func(v *ApplianceDefinition) ApplianceArtifactResponseArrayOutput { return v.Artifacts }).(ApplianceArtifactResponseArrayOutput)
}

func (o ApplianceDefinitionOutput) Authorizations() ApplianceProviderAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v *ApplianceDefinition) ApplianceProviderAuthorizationResponseArrayOutput {
		return v.Authorizations
	}).(ApplianceProviderAuthorizationResponseArrayOutput)
}

func (o ApplianceDefinitionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplianceDefinition) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ApplianceDefinitionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplianceDefinition) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ApplianceDefinitionOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *ApplianceDefinition) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o ApplianceDefinitionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplianceDefinition) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ApplianceDefinitionOutput) LockLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplianceDefinition) pulumi.StringOutput { return v.LockLevel }).(pulumi.StringOutput)
}

func (o ApplianceDefinitionOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplianceDefinition) pulumi.StringPtrOutput { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

func (o ApplianceDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplianceDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplianceDefinitionOutput) PackageFileUri() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplianceDefinition) pulumi.StringOutput { return v.PackageFileUri }).(pulumi.StringOutput)
}

func (o ApplianceDefinitionOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *ApplianceDefinition) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o ApplianceDefinitionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApplianceDefinition) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ApplianceDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplianceDefinition) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplianceDefinitionOutput{})
}
