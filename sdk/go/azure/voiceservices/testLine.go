


package voiceservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TestLine struct {
	pulumi.CustomResourceState

	Location          pulumi.StringOutput      `pulumi:"location"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	PhoneNumber       pulumi.StringOutput      `pulumi:"phoneNumber"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	Purpose           pulumi.StringOutput      `pulumi:"purpose"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput   `pulumi:"tags"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewTestLine(ctx *pulumi.Context,
	name string, args *TestLineArgs, opts ...pulumi.ResourceOption) (*TestLine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CommunicationsGatewayName == nil {
		return nil, errors.New("invalid value for required argument 'CommunicationsGatewayName'")
	}
	if args.PhoneNumber == nil {
		return nil, errors.New("invalid value for required argument 'PhoneNumber'")
	}
	if args.Purpose == nil {
		return nil, errors.New("invalid value for required argument 'Purpose'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:voiceservices/v20221201preview:TestLine"),
		},
		{
			Type: pulumi.String("azure-native:voiceservices/v20230131:TestLine"),
		},
	})
	opts = append(opts, aliases)
	var resource TestLine
	err := ctx.RegisterResource("azure-native:voiceservices:TestLine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTestLine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TestLineState, opts ...pulumi.ResourceOption) (*TestLine, error) {
	var resource TestLine
	err := ctx.ReadResource("azure-native:voiceservices:TestLine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type testLineState struct {
}

type TestLineState struct {
}

func (TestLineState) ElementType() reflect.Type {
	return reflect.TypeOf((*testLineState)(nil)).Elem()
}

type testLineArgs struct {
	CommunicationsGatewayName string            `pulumi:"communicationsGatewayName"`
	Location                  *string           `pulumi:"location"`
	PhoneNumber               string            `pulumi:"phoneNumber"`
	Purpose                   TestLinePurpose   `pulumi:"purpose"`
	ResourceGroupName         string            `pulumi:"resourceGroupName"`
	Tags                      map[string]string `pulumi:"tags"`
	TestLineName              *string           `pulumi:"testLineName"`
}


type TestLineArgs struct {
	CommunicationsGatewayName pulumi.StringInput
	Location                  pulumi.StringPtrInput
	PhoneNumber               pulumi.StringInput
	Purpose                   TestLinePurposeInput
	ResourceGroupName         pulumi.StringInput
	Tags                      pulumi.StringMapInput
	TestLineName              pulumi.StringPtrInput
}

func (TestLineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*testLineArgs)(nil)).Elem()
}

type TestLineInput interface {
	pulumi.Input

	ToTestLineOutput() TestLineOutput
	ToTestLineOutputWithContext(ctx context.Context) TestLineOutput
}

func (*TestLine) ElementType() reflect.Type {
	return reflect.TypeOf((**TestLine)(nil)).Elem()
}

func (i *TestLine) ToTestLineOutput() TestLineOutput {
	return i.ToTestLineOutputWithContext(context.Background())
}

func (i *TestLine) ToTestLineOutputWithContext(ctx context.Context) TestLineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TestLineOutput)
}

type TestLineOutput struct{ *pulumi.OutputState }

func (TestLineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TestLine)(nil)).Elem()
}

func (o TestLineOutput) ToTestLineOutput() TestLineOutput {
	return o
}

func (o TestLineOutput) ToTestLineOutputWithContext(ctx context.Context) TestLineOutput {
	return o
}

func (o TestLineOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *TestLine) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o TestLineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TestLine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TestLineOutput) PhoneNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *TestLine) pulumi.StringOutput { return v.PhoneNumber }).(pulumi.StringOutput)
}

func (o TestLineOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *TestLine) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o TestLineOutput) Purpose() pulumi.StringOutput {
	return o.ApplyT(func(v *TestLine) pulumi.StringOutput { return v.Purpose }).(pulumi.StringOutput)
}

func (o TestLineOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *TestLine) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o TestLineOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TestLine) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o TestLineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TestLine) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TestLineOutput{})
}
