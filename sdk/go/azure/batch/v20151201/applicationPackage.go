


package v20151201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationPackage struct {
	pulumi.CustomResourceState

	Format             pulumi.StringPtrOutput `pulumi:"format"`
	LastActivationTime pulumi.StringPtrOutput `pulumi:"lastActivationTime"`
	State              pulumi.StringPtrOutput `pulumi:"state"`
	StorageUrl         pulumi.StringPtrOutput `pulumi:"storageUrl"`
	StorageUrlExpiry   pulumi.StringPtrOutput `pulumi:"storageUrlExpiry"`
	Version            pulumi.StringPtrOutput `pulumi:"version"`
}


func NewApplicationPackage(ctx *pulumi.Context,
	name string, args *ApplicationPackageArgs, opts ...pulumi.ResourceOption) (*ApplicationPackage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:batch:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170101:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170501:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170901:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20181201:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20190401:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20190801:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200301:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200501:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200901:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20210101:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20210601:ApplicationPackage"),
		},
	})
	opts = append(opts, aliases)
	var resource ApplicationPackage
	err := ctx.RegisterResource("azure-native:batch/v20151201:ApplicationPackage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplicationPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationPackageState, opts ...pulumi.ResourceOption) (*ApplicationPackage, error) {
	var resource ApplicationPackage
	err := ctx.ReadResource("azure-native:batch/v20151201:ApplicationPackage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type applicationPackageState struct {
}

type ApplicationPackageState struct {
}

func (ApplicationPackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationPackageState)(nil)).Elem()
}

type applicationPackageArgs struct {
	AccountName       string  `pulumi:"accountName"`
	ApplicationId     string  `pulumi:"applicationId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	Version           *string `pulumi:"version"`
}


type ApplicationPackageArgs struct {
	AccountName       pulumi.StringInput
	ApplicationId     pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Version           pulumi.StringPtrInput
}

func (ApplicationPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationPackageArgs)(nil)).Elem()
}

type ApplicationPackageInput interface {
	pulumi.Input

	ToApplicationPackageOutput() ApplicationPackageOutput
	ToApplicationPackageOutputWithContext(ctx context.Context) ApplicationPackageOutput
}

func (*ApplicationPackage) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPackage)(nil))
}

func (i *ApplicationPackage) ToApplicationPackageOutput() ApplicationPackageOutput {
	return i.ToApplicationPackageOutputWithContext(context.Background())
}

func (i *ApplicationPackage) ToApplicationPackageOutputWithContext(ctx context.Context) ApplicationPackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPackageOutput)
}

type ApplicationPackageOutput struct{ *pulumi.OutputState }

func (ApplicationPackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPackage)(nil))
}

func (o ApplicationPackageOutput) ToApplicationPackageOutput() ApplicationPackageOutput {
	return o
}

func (o ApplicationPackageOutput) ToApplicationPackageOutputWithContext(ctx context.Context) ApplicationPackageOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApplicationPackageOutput{})
}
