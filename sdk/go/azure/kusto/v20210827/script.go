


package v20210827

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Script struct {
	pulumi.CustomResourceState

	ContinueOnErrors  pulumi.BoolPtrOutput     `pulumi:"continueOnErrors"`
	ForceUpdateTag    pulumi.StringPtrOutput   `pulumi:"forceUpdateTag"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	ScriptUrl         pulumi.StringOutput      `pulumi:"scriptUrl"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewScript(ctx *pulumi.Context,
	name string, args *ScriptArgs, opts ...pulumi.ResourceOption) (*Script, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScriptUrl == nil {
		return nil, errors.New("invalid value for required argument 'ScriptUrl'")
	}
	if args.ScriptUrlSasToken == nil {
		return nil, errors.New("invalid value for required argument 'ScriptUrlSasToken'")
	}
	if args.ContinueOnErrors == nil {
		args.ContinueOnErrors = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:kusto/v20210827:Script"),
		},
		{
			Type: pulumi.String("azure-native:kusto:Script"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto:Script"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:Script"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20210101:Script"),
		},
	})
	opts = append(opts, aliases)
	var resource Script
	err := ctx.RegisterResource("azure-native:kusto/v20210827:Script", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetScript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScriptState, opts ...pulumi.ResourceOption) (*Script, error) {
	var resource Script
	err := ctx.ReadResource("azure-native:kusto/v20210827:Script", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type scriptState struct {
}

type ScriptState struct {
}

func (ScriptState) ElementType() reflect.Type {
	return reflect.TypeOf((*scriptState)(nil)).Elem()
}

type scriptArgs struct {
	ClusterName       string  `pulumi:"clusterName"`
	ContinueOnErrors  *bool   `pulumi:"continueOnErrors"`
	DatabaseName      string  `pulumi:"databaseName"`
	ForceUpdateTag    *string `pulumi:"forceUpdateTag"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ScriptName        *string `pulumi:"scriptName"`
	ScriptUrl         string  `pulumi:"scriptUrl"`
	ScriptUrlSasToken string  `pulumi:"scriptUrlSasToken"`
}


type ScriptArgs struct {
	ClusterName       pulumi.StringInput
	ContinueOnErrors  pulumi.BoolPtrInput
	DatabaseName      pulumi.StringInput
	ForceUpdateTag    pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ScriptName        pulumi.StringPtrInput
	ScriptUrl         pulumi.StringInput
	ScriptUrlSasToken pulumi.StringInput
}

func (ScriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scriptArgs)(nil)).Elem()
}

type ScriptInput interface {
	pulumi.Input

	ToScriptOutput() ScriptOutput
	ToScriptOutputWithContext(ctx context.Context) ScriptOutput
}

func (*Script) ElementType() reflect.Type {
	return reflect.TypeOf((*Script)(nil))
}

func (i *Script) ToScriptOutput() ScriptOutput {
	return i.ToScriptOutputWithContext(context.Background())
}

func (i *Script) ToScriptOutputWithContext(ctx context.Context) ScriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptOutput)
}

type ScriptOutput struct{ *pulumi.OutputState }

func (ScriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Script)(nil))
}

func (o ScriptOutput) ToScriptOutput() ScriptOutput {
	return o
}

func (o ScriptOutput) ToScriptOutputWithContext(ctx context.Context) ScriptOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ScriptOutput{})
}
