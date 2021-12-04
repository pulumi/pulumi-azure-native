


package loadtestservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LoadTest struct {
	pulumi.CustomResourceState

	DataPlaneURI      pulumi.StringOutput                            `pulumi:"dataPlaneURI"`
	Description       pulumi.StringPtrOutput                         `pulumi:"description"`
	Identity          SystemAssignedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	Location          pulumi.StringOutput                            `pulumi:"location"`
	Name              pulumi.StringOutput                            `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                            `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput                       `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                         `pulumi:"tags"`
	Type              pulumi.StringOutput                            `pulumi:"type"`
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
			Type: pulumi.String("azure-native:loadtestservice/v20211201preview:LoadTest"),
		},
	})
	opts = append(opts, aliases)
	var resource LoadTest
	err := ctx.RegisterResource("azure-native:loadtestservice:LoadTest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLoadTest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadTestState, opts ...pulumi.ResourceOption) (*LoadTest, error) {
	var resource LoadTest
	err := ctx.ReadResource("azure-native:loadtestservice:LoadTest", name, id, state, &resource, opts...)
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
	Description       *string                        `pulumi:"description"`
	Identity          *SystemAssignedServiceIdentity `pulumi:"identity"`
	LoadTestName      *string                        `pulumi:"loadTestName"`
	Location          *string                        `pulumi:"location"`
	ResourceGroupName string                         `pulumi:"resourceGroupName"`
	Tags              map[string]string              `pulumi:"tags"`
}


type LoadTestArgs struct {
	Description       pulumi.StringPtrInput
	Identity          SystemAssignedServiceIdentityPtrInput
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
	return reflect.TypeOf((*LoadTest)(nil))
}

func (i *LoadTest) ToLoadTestOutput() LoadTestOutput {
	return i.ToLoadTestOutputWithContext(context.Background())
}

func (i *LoadTest) ToLoadTestOutputWithContext(ctx context.Context) LoadTestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadTestOutput)
}

type LoadTestOutput struct{ *pulumi.OutputState }

func (LoadTestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadTest)(nil))
}

func (o LoadTestOutput) ToLoadTestOutput() LoadTestOutput {
	return o
}

func (o LoadTestOutput) ToLoadTestOutputWithContext(ctx context.Context) LoadTestOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LoadTestOutput{})
}
