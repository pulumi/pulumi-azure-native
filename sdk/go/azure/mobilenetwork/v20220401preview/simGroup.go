


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SimGroup struct {
	pulumi.CustomResourceState

	CreatedAt          pulumi.StringPtrOutput                   `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrOutput                   `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrOutput                   `pulumi:"createdByType"`
	EncryptionKey      KeyVaultKeyResponsePtrOutput             `pulumi:"encryptionKey"`
	Identity           ManagedServiceIdentityResponsePtrOutput  `pulumi:"identity"`
	LastModifiedAt     pulumi.StringPtrOutput                   `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrOutput                   `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrOutput                   `pulumi:"lastModifiedByType"`
	Location           pulumi.StringOutput                      `pulumi:"location"`
	MobileNetwork      MobileNetworkResourceIdResponsePtrOutput `pulumi:"mobileNetwork"`
	Name               pulumi.StringOutput                      `pulumi:"name"`
	ProvisioningState  pulumi.StringOutput                      `pulumi:"provisioningState"`
	SystemData         SystemDataResponseOutput                 `pulumi:"systemData"`
	Tags               pulumi.StringMapOutput                   `pulumi:"tags"`
	Type               pulumi.StringOutput                      `pulumi:"type"`
}


func NewSimGroup(ctx *pulumi.Context,
	name string, args *SimGroupArgs, opts ...pulumi.ResourceOption) (*SimGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork:SimGroup"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20221101:SimGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource SimGroup
	err := ctx.RegisterResource("azure-native:mobilenetwork/v20220401preview:SimGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSimGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SimGroupState, opts ...pulumi.ResourceOption) (*SimGroup, error) {
	var resource SimGroup
	err := ctx.ReadResource("azure-native:mobilenetwork/v20220401preview:SimGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type simGroupState struct {
}

type SimGroupState struct {
}

func (SimGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*simGroupState)(nil)).Elem()
}

type simGroupArgs struct {
	CreatedAt          *string                  `pulumi:"createdAt"`
	CreatedBy          *string                  `pulumi:"createdBy"`
	CreatedByType      *string                  `pulumi:"createdByType"`
	EncryptionKey      *KeyVaultKey             `pulumi:"encryptionKey"`
	Identity           *ManagedServiceIdentity  `pulumi:"identity"`
	LastModifiedAt     *string                  `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string                  `pulumi:"lastModifiedBy"`
	LastModifiedByType *string                  `pulumi:"lastModifiedByType"`
	Location           *string                  `pulumi:"location"`
	MobileNetwork      *MobileNetworkResourceId `pulumi:"mobileNetwork"`
	ResourceGroupName  string                   `pulumi:"resourceGroupName"`
	SimGroupName       *string                  `pulumi:"simGroupName"`
	Tags               map[string]string        `pulumi:"tags"`
}


type SimGroupArgs struct {
	CreatedAt          pulumi.StringPtrInput
	CreatedBy          pulumi.StringPtrInput
	CreatedByType      pulumi.StringPtrInput
	EncryptionKey      KeyVaultKeyPtrInput
	Identity           ManagedServiceIdentityPtrInput
	LastModifiedAt     pulumi.StringPtrInput
	LastModifiedBy     pulumi.StringPtrInput
	LastModifiedByType pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	MobileNetwork      MobileNetworkResourceIdPtrInput
	ResourceGroupName  pulumi.StringInput
	SimGroupName       pulumi.StringPtrInput
	Tags               pulumi.StringMapInput
}

func (SimGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*simGroupArgs)(nil)).Elem()
}

type SimGroupInput interface {
	pulumi.Input

	ToSimGroupOutput() SimGroupOutput
	ToSimGroupOutputWithContext(ctx context.Context) SimGroupOutput
}

func (*SimGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**SimGroup)(nil)).Elem()
}

func (i *SimGroup) ToSimGroupOutput() SimGroupOutput {
	return i.ToSimGroupOutputWithContext(context.Background())
}

func (i *SimGroup) ToSimGroupOutputWithContext(ctx context.Context) SimGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimGroupOutput)
}

type SimGroupOutput struct{ *pulumi.OutputState }

func (SimGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SimGroup)(nil)).Elem()
}

func (o SimGroupOutput) ToSimGroupOutput() SimGroupOutput {
	return o
}

func (o SimGroupOutput) ToSimGroupOutputWithContext(ctx context.Context) SimGroupOutput {
	return o
}

func (o SimGroupOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimGroup) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SimGroupOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimGroup) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SimGroupOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimGroup) pulumi.StringPtrOutput { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SimGroupOutput) EncryptionKey() KeyVaultKeyResponsePtrOutput {
	return o.ApplyT(func(v *SimGroup) KeyVaultKeyResponsePtrOutput { return v.EncryptionKey }).(KeyVaultKeyResponsePtrOutput)
}

func (o SimGroupOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *SimGroup) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o SimGroupOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimGroup) pulumi.StringPtrOutput { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SimGroupOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimGroup) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SimGroupOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimGroup) pulumi.StringPtrOutput { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o SimGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SimGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SimGroupOutput) MobileNetwork() MobileNetworkResourceIdResponsePtrOutput {
	return o.ApplyT(func(v *SimGroup) MobileNetworkResourceIdResponsePtrOutput { return v.MobileNetwork }).(MobileNetworkResourceIdResponsePtrOutput)
}

func (o SimGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SimGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SimGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SimGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SimGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SimGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SimGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SimGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SimGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SimGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SimGroupOutput{})
}
