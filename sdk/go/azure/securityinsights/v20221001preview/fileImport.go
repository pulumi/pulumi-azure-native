


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FileImport struct {
	pulumi.CustomResourceState

	ContentType             pulumi.StringOutput                `pulumi:"contentType"`
	CreatedTimeUTC          pulumi.StringOutput                `pulumi:"createdTimeUTC"`
	ErrorFile               FileMetadataResponseOutput         `pulumi:"errorFile"`
	ErrorsPreview           ValidationErrorResponseArrayOutput `pulumi:"errorsPreview"`
	FilesValidUntilTimeUTC  pulumi.StringOutput                `pulumi:"filesValidUntilTimeUTC"`
	ImportFile              FileMetadataResponseOutput         `pulumi:"importFile"`
	ImportValidUntilTimeUTC pulumi.StringOutput                `pulumi:"importValidUntilTimeUTC"`
	IngestedRecordCount     pulumi.IntOutput                   `pulumi:"ingestedRecordCount"`
	IngestionMode           pulumi.StringOutput                `pulumi:"ingestionMode"`
	Name                    pulumi.StringOutput                `pulumi:"name"`
	Source                  pulumi.StringOutput                `pulumi:"source"`
	State                   pulumi.StringOutput                `pulumi:"state"`
	SystemData              SystemDataResponseOutput           `pulumi:"systemData"`
	TotalRecordCount        pulumi.IntOutput                   `pulumi:"totalRecordCount"`
	Type                    pulumi.StringOutput                `pulumi:"type"`
	ValidRecordCount        pulumi.IntOutput                   `pulumi:"validRecordCount"`
}


func NewFileImport(ctx *pulumi.Context,
	name string, args *FileImportArgs, opts ...pulumi.ResourceOption) (*FileImport, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContentType == nil {
		return nil, errors.New("invalid value for required argument 'ContentType'")
	}
	if args.ImportFile == nil {
		return nil, errors.New("invalid value for required argument 'ImportFile'")
	}
	if args.IngestionMode == nil {
		return nil, errors.New("invalid value for required argument 'IngestionMode'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:FileImport"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:FileImport"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:FileImport"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:FileImport"),
		},
	})
	opts = append(opts, aliases)
	var resource FileImport
	err := ctx.RegisterResource("azure-native:securityinsights/v20221001preview:FileImport", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFileImport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileImportState, opts ...pulumi.ResourceOption) (*FileImport, error) {
	var resource FileImport
	err := ctx.ReadResource("azure-native:securityinsights/v20221001preview:FileImport", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type fileImportState struct {
}

type FileImportState struct {
}

func (FileImportState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileImportState)(nil)).Elem()
}

type fileImportArgs struct {
	ContentType       string       `pulumi:"contentType"`
	FileImportId      *string      `pulumi:"fileImportId"`
	ImportFile        FileMetadata `pulumi:"importFile"`
	IngestionMode     string       `pulumi:"ingestionMode"`
	ResourceGroupName string       `pulumi:"resourceGroupName"`
	Source            string       `pulumi:"source"`
	WorkspaceName     string       `pulumi:"workspaceName"`
}


type FileImportArgs struct {
	ContentType       pulumi.StringInput
	FileImportId      pulumi.StringPtrInput
	ImportFile        FileMetadataInput
	IngestionMode     pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Source            pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (FileImportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileImportArgs)(nil)).Elem()
}

type FileImportInput interface {
	pulumi.Input

	ToFileImportOutput() FileImportOutput
	ToFileImportOutputWithContext(ctx context.Context) FileImportOutput
}

func (*FileImport) ElementType() reflect.Type {
	return reflect.TypeOf((**FileImport)(nil)).Elem()
}

func (i *FileImport) ToFileImportOutput() FileImportOutput {
	return i.ToFileImportOutputWithContext(context.Background())
}

func (i *FileImport) ToFileImportOutputWithContext(ctx context.Context) FileImportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileImportOutput)
}

type FileImportOutput struct{ *pulumi.OutputState }

func (FileImportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileImport)(nil)).Elem()
}

func (o FileImportOutput) ToFileImportOutput() FileImportOutput {
	return o
}

func (o FileImportOutput) ToFileImportOutputWithContext(ctx context.Context) FileImportOutput {
	return o
}

func (o FileImportOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v *FileImport) pulumi.StringOutput { return v.ContentType }).(pulumi.StringOutput)
}

func (o FileImportOutput) CreatedTimeUTC() pulumi.StringOutput {
	return o.ApplyT(func(v *FileImport) pulumi.StringOutput { return v.CreatedTimeUTC }).(pulumi.StringOutput)
}

func (o FileImportOutput) ErrorFile() FileMetadataResponseOutput {
	return o.ApplyT(func(v *FileImport) FileMetadataResponseOutput { return v.ErrorFile }).(FileMetadataResponseOutput)
}

func (o FileImportOutput) ErrorsPreview() ValidationErrorResponseArrayOutput {
	return o.ApplyT(func(v *FileImport) ValidationErrorResponseArrayOutput { return v.ErrorsPreview }).(ValidationErrorResponseArrayOutput)
}

func (o FileImportOutput) FilesValidUntilTimeUTC() pulumi.StringOutput {
	return o.ApplyT(func(v *FileImport) pulumi.StringOutput { return v.FilesValidUntilTimeUTC }).(pulumi.StringOutput)
}

func (o FileImportOutput) ImportFile() FileMetadataResponseOutput {
	return o.ApplyT(func(v *FileImport) FileMetadataResponseOutput { return v.ImportFile }).(FileMetadataResponseOutput)
}

func (o FileImportOutput) ImportValidUntilTimeUTC() pulumi.StringOutput {
	return o.ApplyT(func(v *FileImport) pulumi.StringOutput { return v.ImportValidUntilTimeUTC }).(pulumi.StringOutput)
}

func (o FileImportOutput) IngestedRecordCount() pulumi.IntOutput {
	return o.ApplyT(func(v *FileImport) pulumi.IntOutput { return v.IngestedRecordCount }).(pulumi.IntOutput)
}

func (o FileImportOutput) IngestionMode() pulumi.StringOutput {
	return o.ApplyT(func(v *FileImport) pulumi.StringOutput { return v.IngestionMode }).(pulumi.StringOutput)
}

func (o FileImportOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FileImport) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FileImportOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *FileImport) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

func (o FileImportOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *FileImport) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o FileImportOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FileImport) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o FileImportOutput) TotalRecordCount() pulumi.IntOutput {
	return o.ApplyT(func(v *FileImport) pulumi.IntOutput { return v.TotalRecordCount }).(pulumi.IntOutput)
}

func (o FileImportOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FileImport) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o FileImportOutput) ValidRecordCount() pulumi.IntOutput {
	return o.ApplyT(func(v *FileImport) pulumi.IntOutput { return v.ValidRecordCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(FileImportOutput{})
}
