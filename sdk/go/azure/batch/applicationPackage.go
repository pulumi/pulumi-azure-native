


package batch

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationPackage struct {
	pulumi.CustomResourceState

	Etag               pulumi.StringOutput `pulumi:"etag"`
	Format             pulumi.StringOutput `pulumi:"format"`
	LastActivationTime pulumi.StringOutput `pulumi:"lastActivationTime"`
	Name               pulumi.StringOutput `pulumi:"name"`
	State              pulumi.StringOutput `pulumi:"state"`
	StorageUrl         pulumi.StringOutput `pulumi:"storageUrl"`
	StorageUrlExpiry   pulumi.StringOutput `pulumi:"storageUrlExpiry"`
	Type               pulumi.StringOutput `pulumi:"type"`
}


func NewApplicationPackage(ctx *pulumi.Context,
	name string, args *ApplicationPackageArgs, opts ...pulumi.ResourceOption) (*ApplicationPackage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ApplicationName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:batch:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20151201:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20151201:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170101:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20170101:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170501:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20170501:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170901:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20170901:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20181201:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20181201:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20190401:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20190401:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20190801:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20190801:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200301:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20200301:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200501:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20200501:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200901:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20200901:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20210101:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20210101:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20210601:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20210601:ApplicationPackage"),
		},
	})
	opts = append(opts, aliases)
	var resource ApplicationPackage
	err := ctx.RegisterResource("azure-native:batch:ApplicationPackage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplicationPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationPackageState, opts ...pulumi.ResourceOption) (*ApplicationPackage, error) {
	var resource ApplicationPackage
	err := ctx.ReadResource("azure-native:batch:ApplicationPackage", name, id, state, &resource, opts...)
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
	ApplicationName   string  `pulumi:"applicationName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	VersionName       *string `pulumi:"versionName"`
}


type ApplicationPackageArgs struct {
	AccountName       pulumi.StringInput
	ApplicationName   pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	VersionName       pulumi.StringPtrInput
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
