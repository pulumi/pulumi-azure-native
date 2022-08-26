


package v20201001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type DeploymentScript struct {
	pulumi.CustomResourceState

	Identity   ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	Kind       pulumi.StringOutput                     `pulumi:"kind"`
	Location   pulumi.StringOutput                     `pulumi:"location"`
	Name       pulumi.StringOutput                     `pulumi:"name"`
	SystemData SystemDataResponseOutput                `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                  `pulumi:"tags"`
	Type       pulumi.StringOutput                     `pulumi:"type"`
}


func NewDeploymentScript(ctx *pulumi.Context,
	name string, args *DeploymentScriptArgs, opts ...pulumi.ResourceOption) (*DeploymentScript, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources:DeploymentScript"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20191001preview:DeploymentScript"),
		},
	})
	opts = append(opts, aliases)
	var resource DeploymentScript
	err := ctx.RegisterResource("azure-native:resources/v20201001:DeploymentScript", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDeploymentScript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentScriptState, opts ...pulumi.ResourceOption) (*DeploymentScript, error) {
	var resource DeploymentScript
	err := ctx.ReadResource("azure-native:resources/v20201001:DeploymentScript", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type deploymentScriptState struct {
}

type DeploymentScriptState struct {
}

func (DeploymentScriptState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentScriptState)(nil)).Elem()
}

type deploymentScriptArgs struct {
	Identity          *ManagedServiceIdentity `pulumi:"identity"`
	Kind              string                  `pulumi:"kind"`
	Location          *string                 `pulumi:"location"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	ScriptName        *string                 `pulumi:"scriptName"`
	Tags              map[string]string       `pulumi:"tags"`
}


type DeploymentScriptArgs struct {
	Identity          ManagedServiceIdentityPtrInput
	Kind              pulumi.StringInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ScriptName        pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (DeploymentScriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentScriptArgs)(nil)).Elem()
}

type DeploymentScriptInput interface {
	pulumi.Input

	ToDeploymentScriptOutput() DeploymentScriptOutput
	ToDeploymentScriptOutputWithContext(ctx context.Context) DeploymentScriptOutput
}

func (*DeploymentScript) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentScript)(nil)).Elem()
}

func (i *DeploymentScript) ToDeploymentScriptOutput() DeploymentScriptOutput {
	return i.ToDeploymentScriptOutputWithContext(context.Background())
}

func (i *DeploymentScript) ToDeploymentScriptOutputWithContext(ctx context.Context) DeploymentScriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentScriptOutput)
}

type DeploymentScriptOutput struct{ *pulumi.OutputState }

func (DeploymentScriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentScript)(nil)).Elem()
}

func (o DeploymentScriptOutput) ToDeploymentScriptOutput() DeploymentScriptOutput {
	return o
}

func (o DeploymentScriptOutput) ToDeploymentScriptOutputWithContext(ctx context.Context) DeploymentScriptOutput {
	return o
}

func (o DeploymentScriptOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentScript) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o DeploymentScriptOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentScript) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o DeploymentScriptOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentScript) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DeploymentScriptOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentScript) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DeploymentScriptOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DeploymentScript) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DeploymentScriptOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeploymentScript) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DeploymentScriptOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentScript) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentScriptOutput{})
}
