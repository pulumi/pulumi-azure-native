


package v20201216preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Package struct {
	pulumi.CustomResourceState

	ApplicationName   pulumi.StringOutput                        `pulumi:"applicationName"`
	BlobPath          pulumi.StringOutput                        `pulumi:"blobPath"`
	Etag              pulumi.StringOutput                        `pulumi:"etag"`
	FlightingRing     pulumi.StringOutput                        `pulumi:"flightingRing"`
	IsEnabled         pulumi.BoolOutput                          `pulumi:"isEnabled"`
	LastModifiedTime  pulumi.StringOutput                        `pulumi:"lastModifiedTime"`
	Location          pulumi.StringOutput                        `pulumi:"location"`
	Name              pulumi.StringOutput                        `pulumi:"name"`
	PackageStatus     pulumi.StringOutput                        `pulumi:"packageStatus"`
	ProvisioningState pulumi.StringOutput                        `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput                   `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                     `pulumi:"tags"`
	TargetOSList      TargetOSInfoResponseArrayOutput            `pulumi:"targetOSList"`
	TestTypes         pulumi.StringArrayOutput                   `pulumi:"testTypes"`
	Tests             TestResponseArrayOutput                    `pulumi:"tests"`
	Type              pulumi.StringOutput                        `pulumi:"type"`
	ValidationResults PackageValidationResultResponseArrayOutput `pulumi:"validationResults"`
	Version           pulumi.StringOutput                        `pulumi:"version"`
}


func NewPackage(ctx *pulumi.Context,
	name string, args *PackageArgs, opts ...pulumi.ResourceOption) (*Package, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationName'")
	}
	if args.BlobPath == nil {
		return nil, errors.New("invalid value for required argument 'BlobPath'")
	}
	if args.FlightingRing == nil {
		return nil, errors.New("invalid value for required argument 'FlightingRing'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TargetOSList == nil {
		return nil, errors.New("invalid value for required argument 'TargetOSList'")
	}
	if args.TestBaseAccountName == nil {
		return nil, errors.New("invalid value for required argument 'TestBaseAccountName'")
	}
	if args.Tests == nil {
		return nil, errors.New("invalid value for required argument 'Tests'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:testbase:Package"),
		},
		{
			Type: pulumi.String("azure-native:testbase/v20220401preview:Package"),
		},
	})
	opts = append(opts, aliases)
	var resource Package
	err := ctx.RegisterResource("azure-native:testbase/v20201216preview:Package", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PackageState, opts ...pulumi.ResourceOption) (*Package, error) {
	var resource Package
	err := ctx.ReadResource("azure-native:testbase/v20201216preview:Package", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type packageState struct {
}

type PackageState struct {
}

func (PackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*packageState)(nil)).Elem()
}

type packageArgs struct {
	ApplicationName     string            `pulumi:"applicationName"`
	BlobPath            string            `pulumi:"blobPath"`
	FlightingRing       string            `pulumi:"flightingRing"`
	Location            *string           `pulumi:"location"`
	PackageName         *string           `pulumi:"packageName"`
	ResourceGroupName   string            `pulumi:"resourceGroupName"`
	Tags                map[string]string `pulumi:"tags"`
	TargetOSList        []TargetOSInfo    `pulumi:"targetOSList"`
	TestBaseAccountName string            `pulumi:"testBaseAccountName"`
	Tests               []Test            `pulumi:"tests"`
	Version             string            `pulumi:"version"`
}


type PackageArgs struct {
	ApplicationName     pulumi.StringInput
	BlobPath            pulumi.StringInput
	FlightingRing       pulumi.StringInput
	Location            pulumi.StringPtrInput
	PackageName         pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	Tags                pulumi.StringMapInput
	TargetOSList        TargetOSInfoArrayInput
	TestBaseAccountName pulumi.StringInput
	Tests               TestArrayInput
	Version             pulumi.StringInput
}

func (PackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*packageArgs)(nil)).Elem()
}

type PackageInput interface {
	pulumi.Input

	ToPackageOutput() PackageOutput
	ToPackageOutputWithContext(ctx context.Context) PackageOutput
}

func (*Package) ElementType() reflect.Type {
	return reflect.TypeOf((**Package)(nil)).Elem()
}

func (i *Package) ToPackageOutput() PackageOutput {
	return i.ToPackageOutputWithContext(context.Background())
}

func (i *Package) ToPackageOutputWithContext(ctx context.Context) PackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageOutput)
}

type PackageOutput struct{ *pulumi.OutputState }

func (PackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Package)(nil)).Elem()
}

func (o PackageOutput) ToPackageOutput() PackageOutput {
	return o
}

func (o PackageOutput) ToPackageOutputWithContext(ctx context.Context) PackageOutput {
	return o
}

func (o PackageOutput) ApplicationName() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.ApplicationName }).(pulumi.StringOutput)
}

func (o PackageOutput) BlobPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.BlobPath }).(pulumi.StringOutput)
}

func (o PackageOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o PackageOutput) FlightingRing() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.FlightingRing }).(pulumi.StringOutput)
}

func (o PackageOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Package) pulumi.BoolOutput { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o PackageOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o PackageOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PackageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PackageOutput) PackageStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.PackageStatus }).(pulumi.StringOutput)
}

func (o PackageOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PackageOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Package) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PackageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Package) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PackageOutput) TargetOSList() TargetOSInfoResponseArrayOutput {
	return o.ApplyT(func(v *Package) TargetOSInfoResponseArrayOutput { return v.TargetOSList }).(TargetOSInfoResponseArrayOutput)
}

func (o PackageOutput) TestTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Package) pulumi.StringArrayOutput { return v.TestTypes }).(pulumi.StringArrayOutput)
}

func (o PackageOutput) Tests() TestResponseArrayOutput {
	return o.ApplyT(func(v *Package) TestResponseArrayOutput { return v.Tests }).(TestResponseArrayOutput)
}

func (o PackageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o PackageOutput) ValidationResults() PackageValidationResultResponseArrayOutput {
	return o.ApplyT(func(v *Package) PackageValidationResultResponseArrayOutput { return v.ValidationResults }).(PackageValidationResultResponseArrayOutput)
}

func (o PackageOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PackageOutput{})
}
