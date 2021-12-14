


package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CloudService struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                  `pulumi:"location"`
	Name       pulumi.StringOutput                  `pulumi:"name"`
	Properties CloudServicePropertiesResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput               `pulumi:"tags"`
	Type       pulumi.StringOutput                  `pulumi:"type"`
}


func NewCloudService(ctx *pulumi.Context,
	name string, args *CloudServiceArgs, opts ...pulumi.ResourceOption) (*CloudService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:CloudService"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201001preview:CloudService"),
		},
	})
	opts = append(opts, aliases)
	var resource CloudService
	err := ctx.RegisterResource("azure-native:compute/v20210301:CloudService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCloudService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudServiceState, opts ...pulumi.ResourceOption) (*CloudService, error) {
	var resource CloudService
	err := ctx.ReadResource("azure-native:compute/v20210301:CloudService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type cloudServiceState struct {
}

type CloudServiceState struct {
}

func (CloudServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudServiceState)(nil)).Elem()
}

type cloudServiceArgs struct {
	CloudServiceName  *string                 `pulumi:"cloudServiceName"`
	Location          *string                 `pulumi:"location"`
	Properties        *CloudServiceProperties `pulumi:"properties"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	Tags              map[string]string       `pulumi:"tags"`
}


type CloudServiceArgs struct {
	CloudServiceName  pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Properties        CloudServicePropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (CloudServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudServiceArgs)(nil)).Elem()
}

type CloudServiceInput interface {
	pulumi.Input

	ToCloudServiceOutput() CloudServiceOutput
	ToCloudServiceOutputWithContext(ctx context.Context) CloudServiceOutput
}

func (*CloudService) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudService)(nil))
}

func (i *CloudService) ToCloudServiceOutput() CloudServiceOutput {
	return i.ToCloudServiceOutputWithContext(context.Background())
}

func (i *CloudService) ToCloudServiceOutputWithContext(ctx context.Context) CloudServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceOutput)
}

type CloudServiceOutput struct{ *pulumi.OutputState }

func (CloudServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudService)(nil))
}

func (o CloudServiceOutput) ToCloudServiceOutput() CloudServiceOutput {
	return o
}

func (o CloudServiceOutput) ToCloudServiceOutputWithContext(ctx context.Context) CloudServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CloudServiceOutput{})
}
