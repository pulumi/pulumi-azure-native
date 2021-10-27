


package dataprotection

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceGuard struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput              `pulumi:"eTag"`
	Identity   DppIdentityDetailsResponsePtrOutput `pulumi:"identity"`
	Location   pulumi.StringPtrOutput              `pulumi:"location"`
	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties ResourceGuardResponseOutput         `pulumi:"properties"`
	SystemData SystemDataResponseOutput            `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput              `pulumi:"tags"`
	Type       pulumi.StringOutput                 `pulumi:"type"`
}


func NewResourceGuard(ctx *pulumi.Context,
	name string, args *ResourceGuardArgs, opts ...pulumi.ResourceOption) (*ResourceGuard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:dataprotection:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210701:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-nextgen:dataprotection/v20210701:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20211001preview:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-nextgen:dataprotection/v20211001preview:ResourceGuard"),
		},
	})
	opts = append(opts, aliases)
	var resource ResourceGuard
	err := ctx.RegisterResource("azure-native:dataprotection:ResourceGuard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetResourceGuard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceGuardState, opts ...pulumi.ResourceOption) (*ResourceGuard, error) {
	var resource ResourceGuard
	err := ctx.ReadResource("azure-native:dataprotection:ResourceGuard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type resourceGuardState struct {
}

type ResourceGuardState struct {
}

func (ResourceGuardState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGuardState)(nil)).Elem()
}

type resourceGuardArgs struct {
	ETag               *string             `pulumi:"eTag"`
	Identity           *DppIdentityDetails `pulumi:"identity"`
	Location           *string             `pulumi:"location"`
	ResourceGroupName  string              `pulumi:"resourceGroupName"`
	ResourceGuardsName *string             `pulumi:"resourceGuardsName"`
	Tags               map[string]string   `pulumi:"tags"`
}


type ResourceGuardArgs struct {
	ETag               pulumi.StringPtrInput
	Identity           DppIdentityDetailsPtrInput
	Location           pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	ResourceGuardsName pulumi.StringPtrInput
	Tags               pulumi.StringMapInput
}

func (ResourceGuardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGuardArgs)(nil)).Elem()
}

type ResourceGuardInput interface {
	pulumi.Input

	ToResourceGuardOutput() ResourceGuardOutput
	ToResourceGuardOutputWithContext(ctx context.Context) ResourceGuardOutput
}

func (*ResourceGuard) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGuard)(nil))
}

func (i *ResourceGuard) ToResourceGuardOutput() ResourceGuardOutput {
	return i.ToResourceGuardOutputWithContext(context.Background())
}

func (i *ResourceGuard) ToResourceGuardOutputWithContext(ctx context.Context) ResourceGuardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGuardOutput)
}

type ResourceGuardOutput struct{ *pulumi.OutputState }

func (ResourceGuardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGuard)(nil))
}

func (o ResourceGuardOutput) ToResourceGuardOutput() ResourceGuardOutput {
	return o
}

func (o ResourceGuardOutput) ToResourceGuardOutputWithContext(ctx context.Context) ResourceGuardOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceGuardInput)(nil)).Elem(), &ResourceGuard{})
	pulumi.RegisterOutputType(ResourceGuardOutput{})
}
