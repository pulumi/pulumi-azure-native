


package v20191101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Rollout struct {
	pulumi.CustomResourceState

	ArtifactSourceId        pulumi.StringPtrOutput       `pulumi:"artifactSourceId"`
	BuildVersion            pulumi.StringOutput          `pulumi:"buildVersion"`
	Identity                IdentityResponseOutput       `pulumi:"identity"`
	Location                pulumi.StringOutput          `pulumi:"location"`
	Name                    pulumi.StringOutput          `pulumi:"name"`
	StepGroups              StepGroupResponseArrayOutput `pulumi:"stepGroups"`
	Tags                    pulumi.StringMapOutput       `pulumi:"tags"`
	TargetServiceTopologyId pulumi.StringOutput          `pulumi:"targetServiceTopologyId"`
	Type                    pulumi.StringOutput          `pulumi:"type"`
}


func NewRollout(ctx *pulumi.Context,
	name string, args *RolloutArgs, opts ...pulumi.ResourceOption) (*Rollout, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BuildVersion == nil {
		return nil, errors.New("invalid value for required argument 'BuildVersion'")
	}
	if args.Identity == nil {
		return nil, errors.New("invalid value for required argument 'Identity'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StepGroups == nil {
		return nil, errors.New("invalid value for required argument 'StepGroups'")
	}
	if args.TargetServiceTopologyId == nil {
		return nil, errors.New("invalid value for required argument 'TargetServiceTopologyId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:deploymentmanager/v20191101preview:Rollout"),
		},
		{
			Type: pulumi.String("azure-native:deploymentmanager:Rollout"),
		},
		{
			Type: pulumi.String("azure-nextgen:deploymentmanager:Rollout"),
		},
		{
			Type: pulumi.String("azure-native:deploymentmanager/v20180901preview:Rollout"),
		},
		{
			Type: pulumi.String("azure-nextgen:deploymentmanager/v20180901preview:Rollout"),
		},
	})
	opts = append(opts, aliases)
	var resource Rollout
	err := ctx.RegisterResource("azure-native:deploymentmanager/v20191101preview:Rollout", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRollout(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RolloutState, opts ...pulumi.ResourceOption) (*Rollout, error) {
	var resource Rollout
	err := ctx.ReadResource("azure-native:deploymentmanager/v20191101preview:Rollout", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type rolloutState struct {
}

type RolloutState struct {
}

func (RolloutState) ElementType() reflect.Type {
	return reflect.TypeOf((*rolloutState)(nil)).Elem()
}

type rolloutArgs struct {
	ArtifactSourceId        *string           `pulumi:"artifactSourceId"`
	BuildVersion            string            `pulumi:"buildVersion"`
	Identity                Identity          `pulumi:"identity"`
	Location                *string           `pulumi:"location"`
	ResourceGroupName       string            `pulumi:"resourceGroupName"`
	RolloutName             *string           `pulumi:"rolloutName"`
	StepGroups              []StepGroup       `pulumi:"stepGroups"`
	Tags                    map[string]string `pulumi:"tags"`
	TargetServiceTopologyId string            `pulumi:"targetServiceTopologyId"`
}


type RolloutArgs struct {
	ArtifactSourceId        pulumi.StringPtrInput
	BuildVersion            pulumi.StringInput
	Identity                IdentityInput
	Location                pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	RolloutName             pulumi.StringPtrInput
	StepGroups              StepGroupArrayInput
	Tags                    pulumi.StringMapInput
	TargetServiceTopologyId pulumi.StringInput
}

func (RolloutArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rolloutArgs)(nil)).Elem()
}

type RolloutInput interface {
	pulumi.Input

	ToRolloutOutput() RolloutOutput
	ToRolloutOutputWithContext(ctx context.Context) RolloutOutput
}

func (*Rollout) ElementType() reflect.Type {
	return reflect.TypeOf((*Rollout)(nil))
}

func (i *Rollout) ToRolloutOutput() RolloutOutput {
	return i.ToRolloutOutputWithContext(context.Background())
}

func (i *Rollout) ToRolloutOutputWithContext(ctx context.Context) RolloutOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolloutOutput)
}

type RolloutOutput struct{ *pulumi.OutputState }

func (RolloutOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Rollout)(nil))
}

func (o RolloutOutput) ToRolloutOutput() RolloutOutput {
	return o
}

func (o RolloutOutput) ToRolloutOutputWithContext(ctx context.Context) RolloutOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RolloutInput)(nil)).Elem(), &Rollout{})
	pulumi.RegisterOutputType(RolloutOutput{})
}
