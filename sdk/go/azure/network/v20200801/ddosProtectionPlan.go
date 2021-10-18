


package v20200801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DdosProtectionPlan struct {
	pulumi.CustomResourceState

	Etag              pulumi.StringOutput            `pulumi:"etag"`
	Location          pulumi.StringPtrOutput         `pulumi:"location"`
	Name              pulumi.StringOutput            `pulumi:"name"`
	ProvisioningState pulumi.StringOutput            `pulumi:"provisioningState"`
	ResourceGuid      pulumi.StringOutput            `pulumi:"resourceGuid"`
	Tags              pulumi.StringMapOutput         `pulumi:"tags"`
	Type              pulumi.StringOutput            `pulumi:"type"`
	VirtualNetworks   SubResourceResponseArrayOutput `pulumi:"virtualNetworks"`
}


func NewDdosProtectionPlan(ctx *pulumi.Context,
	name string, args *DdosProtectionPlanArgs, opts ...pulumi.ResourceOption) (*DdosProtectionPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180201:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180401:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180601:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180701:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20201101:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210201:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210301:DdosProtectionPlan"),
		},
	})
	opts = append(opts, aliases)
	var resource DdosProtectionPlan
	err := ctx.RegisterResource("azure-native:network/v20200801:DdosProtectionPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDdosProtectionPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DdosProtectionPlanState, opts ...pulumi.ResourceOption) (*DdosProtectionPlan, error) {
	var resource DdosProtectionPlan
	err := ctx.ReadResource("azure-native:network/v20200801:DdosProtectionPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type ddosProtectionPlanState struct {
}

type DdosProtectionPlanState struct {
}

func (DdosProtectionPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosProtectionPlanState)(nil)).Elem()
}

type ddosProtectionPlanArgs struct {
	DdosProtectionPlanName *string           `pulumi:"ddosProtectionPlanName"`
	Location               *string           `pulumi:"location"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	Tags                   map[string]string `pulumi:"tags"`
}


type DdosProtectionPlanArgs struct {
	DdosProtectionPlanName pulumi.StringPtrInput
	Location               pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
}

func (DdosProtectionPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosProtectionPlanArgs)(nil)).Elem()
}

type DdosProtectionPlanInput interface {
	pulumi.Input

	ToDdosProtectionPlanOutput() DdosProtectionPlanOutput
	ToDdosProtectionPlanOutputWithContext(ctx context.Context) DdosProtectionPlanOutput
}

func (*DdosProtectionPlan) ElementType() reflect.Type {
	return reflect.TypeOf((*DdosProtectionPlan)(nil))
}

func (i *DdosProtectionPlan) ToDdosProtectionPlanOutput() DdosProtectionPlanOutput {
	return i.ToDdosProtectionPlanOutputWithContext(context.Background())
}

func (i *DdosProtectionPlan) ToDdosProtectionPlanOutputWithContext(ctx context.Context) DdosProtectionPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdosProtectionPlanOutput)
}

type DdosProtectionPlanOutput struct{ *pulumi.OutputState }

func (DdosProtectionPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DdosProtectionPlan)(nil))
}

func (o DdosProtectionPlanOutput) ToDdosProtectionPlanOutput() DdosProtectionPlanOutput {
	return o
}

func (o DdosProtectionPlanOutput) ToDdosProtectionPlanOutputWithContext(ctx context.Context) DdosProtectionPlanOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DdosProtectionPlanOutput{})
}
