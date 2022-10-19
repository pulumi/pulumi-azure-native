


package v20220415preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LoadTest struct {
	pulumi.CustomResourceState

	DataPlaneURI      pulumi.StringOutput                     `pulumi:"dataPlaneURI"`
	Description       pulumi.StringPtrOutput                  `pulumi:"description"`
	Encryption        EncryptionPropertiesResponsePtrOutput   `pulumi:"encryption"`
	Identity          ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	Location          pulumi.StringOutput                     `pulumi:"location"`
	Name              pulumi.StringOutput                     `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                     `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput                `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                  `pulumi:"tags"`
	Type              pulumi.StringOutput                     `pulumi:"type"`
}


func NewLoadTest(ctx *pulumi.Context,
	name string, args *LoadTestArgs, opts ...pulumi.ResourceOption) (*LoadTest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:loadtestservice:LoadTest"),
		},
		{
			Type: pulumi.String("azure-native:loadtestservice/v20211201preview:LoadTest"),
		},
		{
			Type: pulumi.String("azure-native:loadtestservice/v20221201:LoadTest"),
		},
	})
	opts = append(opts, aliases)
	var resource LoadTest
	err := ctx.RegisterResource("azure-native:loadtestservice/v20220415preview:LoadTest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLoadTest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadTestState, opts ...pulumi.ResourceOption) (*LoadTest, error) {
	var resource LoadTest
	err := ctx.ReadResource("azure-native:loadtestservice/v20220415preview:LoadTest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type loadTestState struct {
}

type LoadTestState struct {
}

func (LoadTestState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadTestState)(nil)).Elem()
}

type loadTestArgs struct {
	Description       *string                 `pulumi:"description"`
	Encryption        *EncryptionProperties   `pulumi:"encryption"`
	Identity          *ManagedServiceIdentity `pulumi:"identity"`
	LoadTestName      *string                 `pulumi:"loadTestName"`
	Location          *string                 `pulumi:"location"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	Tags              map[string]string       `pulumi:"tags"`
}


type LoadTestArgs struct {
	Description       pulumi.StringPtrInput
	Encryption        EncryptionPropertiesPtrInput
	Identity          ManagedServiceIdentityPtrInput
	LoadTestName      pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (LoadTestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadTestArgs)(nil)).Elem()
}

type LoadTestInput interface {
	pulumi.Input

	ToLoadTestOutput() LoadTestOutput
	ToLoadTestOutputWithContext(ctx context.Context) LoadTestOutput
}

func (*LoadTest) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadTest)(nil)).Elem()
}

func (i *LoadTest) ToLoadTestOutput() LoadTestOutput {
	return i.ToLoadTestOutputWithContext(context.Background())
}

func (i *LoadTest) ToLoadTestOutputWithContext(ctx context.Context) LoadTestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadTestOutput)
}

type LoadTestOutput struct{ *pulumi.OutputState }

func (LoadTestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadTest)(nil)).Elem()
}

func (o LoadTestOutput) ToLoadTestOutput() LoadTestOutput {
	return o
}

func (o LoadTestOutput) ToLoadTestOutputWithContext(ctx context.Context) LoadTestOutput {
	return o
}

func (o LoadTestOutput) DataPlaneURI() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadTest) pulumi.StringOutput { return v.DataPlaneURI }).(pulumi.StringOutput)
}

func (o LoadTestOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadTest) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LoadTestOutput) Encryption() EncryptionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *LoadTest) EncryptionPropertiesResponsePtrOutput { return v.Encryption }).(EncryptionPropertiesResponsePtrOutput)
}

func (o LoadTestOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *LoadTest) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LoadTestOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadTest) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o LoadTestOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadTest) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LoadTestOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadTest) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LoadTestOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *LoadTest) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LoadTestOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LoadTest) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LoadTestOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadTest) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LoadTestOutput{})
}
