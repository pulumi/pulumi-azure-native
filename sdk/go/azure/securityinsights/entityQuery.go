


package securityinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type EntityQuery struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput   `pulumi:"etag"`
	Kind       pulumi.StringOutput      `pulumi:"kind"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewEntityQuery(ctx *pulumi.Context,
	name string, args *EntityQueryArgs, opts ...pulumi.ResourceOption) (*EntityQuery, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("invalid value for required argument 'OperationalInsightsResourceProvider'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:EntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:EntityQuery"),
		},
	})
	opts = append(opts, aliases)
	var resource EntityQuery
	err := ctx.RegisterResource("azure-native:securityinsights:EntityQuery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEntityQuery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntityQueryState, opts ...pulumi.ResourceOption) (*EntityQuery, error) {
	var resource EntityQuery
	err := ctx.ReadResource("azure-native:securityinsights:EntityQuery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type entityQueryState struct {
}

type EntityQueryState struct {
}

func (EntityQueryState) ElementType() reflect.Type {
	return reflect.TypeOf((*entityQueryState)(nil)).Elem()
}

type entityQueryArgs struct {
	EntityQueryId                       *string `pulumi:"entityQueryId"`
	Etag                                *string `pulumi:"etag"`
	Kind                                string  `pulumi:"kind"`
	OperationalInsightsResourceProvider string  `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string  `pulumi:"resourceGroupName"`
	WorkspaceName                       string  `pulumi:"workspaceName"`
}


type EntityQueryArgs struct {
	EntityQueryId                       pulumi.StringPtrInput
	Etag                                pulumi.StringPtrInput
	Kind                                pulumi.StringInput
	OperationalInsightsResourceProvider pulumi.StringInput
	ResourceGroupName                   pulumi.StringInput
	WorkspaceName                       pulumi.StringInput
}

func (EntityQueryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entityQueryArgs)(nil)).Elem()
}

type EntityQueryInput interface {
	pulumi.Input

	ToEntityQueryOutput() EntityQueryOutput
	ToEntityQueryOutputWithContext(ctx context.Context) EntityQueryOutput
}

func (*EntityQuery) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityQuery)(nil))
}

func (i *EntityQuery) ToEntityQueryOutput() EntityQueryOutput {
	return i.ToEntityQueryOutputWithContext(context.Background())
}

func (i *EntityQuery) ToEntityQueryOutputWithContext(ctx context.Context) EntityQueryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityQueryOutput)
}

type EntityQueryOutput struct{ *pulumi.OutputState }

func (EntityQueryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityQuery)(nil))
}

func (o EntityQueryOutput) ToEntityQueryOutput() EntityQueryOutput {
	return o
}

func (o EntityQueryOutput) ToEntityQueryOutputWithContext(ctx context.Context) EntityQueryOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EntityQueryOutput{})
}
