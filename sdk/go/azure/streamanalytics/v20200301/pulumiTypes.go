


package v20200301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AggregateFunctionProperties struct {
	Binding interface{}         `pulumi:"binding"`
	Inputs  []FunctionInputType `pulumi:"inputs"`
	Output  *FunctionOutputType `pulumi:"output"`
	Type    string              `pulumi:"type"`
}

type AggregateFunctionPropertiesResponse struct {
	Binding interface{}             `pulumi:"binding"`
	Etag    string                  `pulumi:"etag"`
	Inputs  []FunctionInputResponse `pulumi:"inputs"`
	Output  *FunctionOutputResponse `pulumi:"output"`
	Type    string                  `pulumi:"type"`
}

type AvroSerialization struct {
	Type string `pulumi:"type"`
}

type AvroSerializationResponse struct {
	Type string `pulumi:"type"`
}

type AzureDataLakeStoreOutputDataSource struct {
	AccountName            *string `pulumi:"accountName"`
	AuthenticationMode     *string `pulumi:"authenticationMode"`
	DateFormat             *string `pulumi:"dateFormat"`
	FilePathPrefix         *string `pulumi:"filePathPrefix"`
	RefreshToken           *string `pulumi:"refreshToken"`
	TenantId               *string `pulumi:"tenantId"`
	TimeFormat             *string `pulumi:"timeFormat"`
	TokenUserDisplayName   *string `pulumi:"tokenUserDisplayName"`
	TokenUserPrincipalName *string `pulumi:"tokenUserPrincipalName"`
	Type                   string  `pulumi:"type"`
}

type AzureDataLakeStoreOutputDataSourceResponse struct {
	AccountName            *string `pulumi:"accountName"`
	AuthenticationMode     *string `pulumi:"authenticationMode"`
	DateFormat             *string `pulumi:"dateFormat"`
	FilePathPrefix         *string `pulumi:"filePathPrefix"`
	RefreshToken           *string `pulumi:"refreshToken"`
	TenantId               *string `pulumi:"tenantId"`
	TimeFormat             *string `pulumi:"timeFormat"`
	TokenUserDisplayName   *string `pulumi:"tokenUserDisplayName"`
	TokenUserPrincipalName *string `pulumi:"tokenUserPrincipalName"`
	Type                   string  `pulumi:"type"`
}

type AzureFunctionOutputDataSource struct {
	ApiKey          *string  `pulumi:"apiKey"`
	FunctionAppName *string  `pulumi:"functionAppName"`
	FunctionName    *string  `pulumi:"functionName"`
	MaxBatchCount   *float64 `pulumi:"maxBatchCount"`
	MaxBatchSize    *float64 `pulumi:"maxBatchSize"`
	Type            string   `pulumi:"type"`
}

type AzureFunctionOutputDataSourceResponse struct {
	ApiKey          *string  `pulumi:"apiKey"`
	FunctionAppName *string  `pulumi:"functionAppName"`
	FunctionName    *string  `pulumi:"functionName"`
	MaxBatchCount   *float64 `pulumi:"maxBatchCount"`
	MaxBatchSize    *float64 `pulumi:"maxBatchSize"`
	Type            string   `pulumi:"type"`
}

type AzureMachineLearningWebServiceFunctionBinding struct {
	ApiKey    *string                                      `pulumi:"apiKey"`
	BatchSize *int                                         `pulumi:"batchSize"`
	Endpoint  *string                                      `pulumi:"endpoint"`
	Inputs    *AzureMachineLearningWebServiceInputs        `pulumi:"inputs"`
	Outputs   []AzureMachineLearningWebServiceOutputColumn `pulumi:"outputs"`
	Type      string                                       `pulumi:"type"`
}

type AzureMachineLearningWebServiceFunctionBindingResponse struct {
	ApiKey    *string                                              `pulumi:"apiKey"`
	BatchSize *int                                                 `pulumi:"batchSize"`
	Endpoint  *string                                              `pulumi:"endpoint"`
	Inputs    *AzureMachineLearningWebServiceInputsResponse        `pulumi:"inputs"`
	Outputs   []AzureMachineLearningWebServiceOutputColumnResponse `pulumi:"outputs"`
	Type      string                                               `pulumi:"type"`
}

type AzureMachineLearningWebServiceInputColumn struct {
	DataType *string `pulumi:"dataType"`
	MapTo    *int    `pulumi:"mapTo"`
	Name     *string `pulumi:"name"`
}

type AzureMachineLearningWebServiceInputColumnResponse struct {
	DataType *string `pulumi:"dataType"`
	MapTo    *int    `pulumi:"mapTo"`
	Name     *string `pulumi:"name"`
}

type AzureMachineLearningWebServiceInputs struct {
	ColumnNames []AzureMachineLearningWebServiceInputColumn `pulumi:"columnNames"`
	Name        *string                                     `pulumi:"name"`
}

type AzureMachineLearningWebServiceInputsResponse struct {
	ColumnNames []AzureMachineLearningWebServiceInputColumnResponse `pulumi:"columnNames"`
	Name        *string                                             `pulumi:"name"`
}

type AzureMachineLearningWebServiceOutputColumn struct {
	DataType *string `pulumi:"dataType"`
	Name     *string `pulumi:"name"`
}

type AzureMachineLearningWebServiceOutputColumnResponse struct {
	DataType *string `pulumi:"dataType"`
	Name     *string `pulumi:"name"`
}

type AzureSqlDatabaseOutputDataSource struct {
	AuthenticationMode *string  `pulumi:"authenticationMode"`
	Database           *string  `pulumi:"database"`
	MaxBatchCount      *float64 `pulumi:"maxBatchCount"`
	MaxWriterCount     *float64 `pulumi:"maxWriterCount"`
	Password           *string  `pulumi:"password"`
	Server             *string  `pulumi:"server"`
	Table              *string  `pulumi:"table"`
	Type               string   `pulumi:"type"`
	User               *string  `pulumi:"user"`
}

type AzureSqlDatabaseOutputDataSourceResponse struct {
	AuthenticationMode *string  `pulumi:"authenticationMode"`
	Database           *string  `pulumi:"database"`
	MaxBatchCount      *float64 `pulumi:"maxBatchCount"`
	MaxWriterCount     *float64 `pulumi:"maxWriterCount"`
	Password           *string  `pulumi:"password"`
	Server             *string  `pulumi:"server"`
	Table              *string  `pulumi:"table"`
	Type               string   `pulumi:"type"`
	User               *string  `pulumi:"user"`
}

type AzureSqlReferenceInputDataSource struct {
	Database           *string `pulumi:"database"`
	DeltaSnapshotQuery *string `pulumi:"deltaSnapshotQuery"`
	FullSnapshotQuery  *string `pulumi:"fullSnapshotQuery"`
	Password           *string `pulumi:"password"`
	RefreshRate        *string `pulumi:"refreshRate"`
	RefreshType        *string `pulumi:"refreshType"`
	Server             *string `pulumi:"server"`
	Table              *string `pulumi:"table"`
	Type               string  `pulumi:"type"`
	User               *string `pulumi:"user"`
}

type AzureSqlReferenceInputDataSourceResponse struct {
	Database           *string `pulumi:"database"`
	DeltaSnapshotQuery *string `pulumi:"deltaSnapshotQuery"`
	FullSnapshotQuery  *string `pulumi:"fullSnapshotQuery"`
	Password           *string `pulumi:"password"`
	RefreshRate        *string `pulumi:"refreshRate"`
	RefreshType        *string `pulumi:"refreshType"`
	Server             *string `pulumi:"server"`
	Table              *string `pulumi:"table"`
	Type               string  `pulumi:"type"`
	User               *string `pulumi:"user"`
}

type AzureSynapseOutputDataSource struct {
	Database *string `pulumi:"database"`
	Password *string `pulumi:"password"`
	Server   *string `pulumi:"server"`
	Table    *string `pulumi:"table"`
	Type     string  `pulumi:"type"`
	User     *string `pulumi:"user"`
}

type AzureSynapseOutputDataSourceResponse struct {
	Database *string `pulumi:"database"`
	Password *string `pulumi:"password"`
	Server   *string `pulumi:"server"`
	Table    *string `pulumi:"table"`
	Type     string  `pulumi:"type"`
	User     *string `pulumi:"user"`
}

type AzureTableOutputDataSource struct {
	AccountKey      *string  `pulumi:"accountKey"`
	AccountName     *string  `pulumi:"accountName"`
	BatchSize       *int     `pulumi:"batchSize"`
	ColumnsToRemove []string `pulumi:"columnsToRemove"`
	PartitionKey    *string  `pulumi:"partitionKey"`
	RowKey          *string  `pulumi:"rowKey"`
	Table           *string  `pulumi:"table"`
	Type            string   `pulumi:"type"`
}

type AzureTableOutputDataSourceResponse struct {
	AccountKey      *string  `pulumi:"accountKey"`
	AccountName     *string  `pulumi:"accountName"`
	BatchSize       *int     `pulumi:"batchSize"`
	ColumnsToRemove []string `pulumi:"columnsToRemove"`
	PartitionKey    *string  `pulumi:"partitionKey"`
	RowKey          *string  `pulumi:"rowKey"`
	Table           *string  `pulumi:"table"`
	Type            string   `pulumi:"type"`
}

type BlobOutputDataSource struct {
	AuthenticationMode *string          `pulumi:"authenticationMode"`
	BlobPathPrefix     *string          `pulumi:"blobPathPrefix"`
	Container          *string          `pulumi:"container"`
	DateFormat         *string          `pulumi:"dateFormat"`
	PathPattern        *string          `pulumi:"pathPattern"`
	StorageAccounts    []StorageAccount `pulumi:"storageAccounts"`
	TimeFormat         *string          `pulumi:"timeFormat"`
	Type               string           `pulumi:"type"`
}

type BlobOutputDataSourceResponse struct {
	AuthenticationMode *string                  `pulumi:"authenticationMode"`
	BlobPathPrefix     *string                  `pulumi:"blobPathPrefix"`
	Container          *string                  `pulumi:"container"`
	DateFormat         *string                  `pulumi:"dateFormat"`
	PathPattern        *string                  `pulumi:"pathPattern"`
	StorageAccounts    []StorageAccountResponse `pulumi:"storageAccounts"`
	TimeFormat         *string                  `pulumi:"timeFormat"`
	Type               string                   `pulumi:"type"`
}

type BlobReferenceInputDataSource struct {
	AuthenticationMode *string          `pulumi:"authenticationMode"`
	Container          *string          `pulumi:"container"`
	DateFormat         *string          `pulumi:"dateFormat"`
	PathPattern        *string          `pulumi:"pathPattern"`
	StorageAccounts    []StorageAccount `pulumi:"storageAccounts"`
	TimeFormat         *string          `pulumi:"timeFormat"`
	Type               string           `pulumi:"type"`
}

type BlobReferenceInputDataSourceResponse struct {
	AuthenticationMode *string                  `pulumi:"authenticationMode"`
	Container          *string                  `pulumi:"container"`
	DateFormat         *string                  `pulumi:"dateFormat"`
	PathPattern        *string                  `pulumi:"pathPattern"`
	StorageAccounts    []StorageAccountResponse `pulumi:"storageAccounts"`
	TimeFormat         *string                  `pulumi:"timeFormat"`
	Type               string                   `pulumi:"type"`
}

type BlobStreamInputDataSource struct {
	AuthenticationMode   *string          `pulumi:"authenticationMode"`
	Container            *string          `pulumi:"container"`
	DateFormat           *string          `pulumi:"dateFormat"`
	PathPattern          *string          `pulumi:"pathPattern"`
	SourcePartitionCount *int             `pulumi:"sourcePartitionCount"`
	StorageAccounts      []StorageAccount `pulumi:"storageAccounts"`
	TimeFormat           *string          `pulumi:"timeFormat"`
	Type                 string           `pulumi:"type"`
}

type BlobStreamInputDataSourceResponse struct {
	AuthenticationMode   *string                  `pulumi:"authenticationMode"`
	Container            *string                  `pulumi:"container"`
	DateFormat           *string                  `pulumi:"dateFormat"`
	PathPattern          *string                  `pulumi:"pathPattern"`
	SourcePartitionCount *int                     `pulumi:"sourcePartitionCount"`
	StorageAccounts      []StorageAccountResponse `pulumi:"storageAccounts"`
	TimeFormat           *string                  `pulumi:"timeFormat"`
	Type                 string                   `pulumi:"type"`
}

type ClusterInfo struct {
	Id *string `pulumi:"id"`
}





type ClusterInfoInput interface {
	pulumi.Input

	ToClusterInfoOutput() ClusterInfoOutput
	ToClusterInfoOutputWithContext(context.Context) ClusterInfoOutput
}

type ClusterInfoArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ClusterInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterInfo)(nil)).Elem()
}

func (i ClusterInfoArgs) ToClusterInfoOutput() ClusterInfoOutput {
	return i.ToClusterInfoOutputWithContext(context.Background())
}

func (i ClusterInfoArgs) ToClusterInfoOutputWithContext(ctx context.Context) ClusterInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterInfoOutput)
}

func (i ClusterInfoArgs) ToClusterInfoPtrOutput() ClusterInfoPtrOutput {
	return i.ToClusterInfoPtrOutputWithContext(context.Background())
}

func (i ClusterInfoArgs) ToClusterInfoPtrOutputWithContext(ctx context.Context) ClusterInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterInfoOutput).ToClusterInfoPtrOutputWithContext(ctx)
}









type ClusterInfoPtrInput interface {
	pulumi.Input

	ToClusterInfoPtrOutput() ClusterInfoPtrOutput
	ToClusterInfoPtrOutputWithContext(context.Context) ClusterInfoPtrOutput
}

type clusterInfoPtrType ClusterInfoArgs

func ClusterInfoPtr(v *ClusterInfoArgs) ClusterInfoPtrInput {
	return (*clusterInfoPtrType)(v)
}

func (*clusterInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterInfo)(nil)).Elem()
}

func (i *clusterInfoPtrType) ToClusterInfoPtrOutput() ClusterInfoPtrOutput {
	return i.ToClusterInfoPtrOutputWithContext(context.Background())
}

func (i *clusterInfoPtrType) ToClusterInfoPtrOutputWithContext(ctx context.Context) ClusterInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterInfoPtrOutput)
}

type ClusterInfoOutput struct{ *pulumi.OutputState }

func (ClusterInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterInfo)(nil)).Elem()
}

func (o ClusterInfoOutput) ToClusterInfoOutput() ClusterInfoOutput {
	return o
}

func (o ClusterInfoOutput) ToClusterInfoOutputWithContext(ctx context.Context) ClusterInfoOutput {
	return o
}

func (o ClusterInfoOutput) ToClusterInfoPtrOutput() ClusterInfoPtrOutput {
	return o.ToClusterInfoPtrOutputWithContext(context.Background())
}

func (o ClusterInfoOutput) ToClusterInfoPtrOutputWithContext(ctx context.Context) ClusterInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterInfo) *ClusterInfo {
		return &v
	}).(ClusterInfoPtrOutput)
}

func (o ClusterInfoOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterInfo) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ClusterInfoPtrOutput struct{ *pulumi.OutputState }

func (ClusterInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterInfo)(nil)).Elem()
}

func (o ClusterInfoPtrOutput) ToClusterInfoPtrOutput() ClusterInfoPtrOutput {
	return o
}

func (o ClusterInfoPtrOutput) ToClusterInfoPtrOutputWithContext(ctx context.Context) ClusterInfoPtrOutput {
	return o
}

func (o ClusterInfoPtrOutput) Elem() ClusterInfoOutput {
	return o.ApplyT(func(v *ClusterInfo) ClusterInfo {
		if v != nil {
			return *v
		}
		var ret ClusterInfo
		return ret
	}).(ClusterInfoOutput)
}

func (o ClusterInfoPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterInfo) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type ClusterInfoResponse struct {
	Id *string `pulumi:"id"`
}

type ClusterInfoResponseOutput struct{ *pulumi.OutputState }

func (ClusterInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterInfoResponse)(nil)).Elem()
}

func (o ClusterInfoResponseOutput) ToClusterInfoResponseOutput() ClusterInfoResponseOutput {
	return o
}

func (o ClusterInfoResponseOutput) ToClusterInfoResponseOutputWithContext(ctx context.Context) ClusterInfoResponseOutput {
	return o
}

func (o ClusterInfoResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterInfoResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ClusterInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ClusterInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterInfoResponse)(nil)).Elem()
}

func (o ClusterInfoResponsePtrOutput) ToClusterInfoResponsePtrOutput() ClusterInfoResponsePtrOutput {
	return o
}

func (o ClusterInfoResponsePtrOutput) ToClusterInfoResponsePtrOutputWithContext(ctx context.Context) ClusterInfoResponsePtrOutput {
	return o
}

func (o ClusterInfoResponsePtrOutput) Elem() ClusterInfoResponseOutput {
	return o.ApplyT(func(v *ClusterInfoResponse) ClusterInfoResponse {
		if v != nil {
			return *v
		}
		var ret ClusterInfoResponse
		return ret
	}).(ClusterInfoResponseOutput)
}

func (o ClusterInfoResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type ClusterJobResponse struct {
	Id             string `pulumi:"id"`
	JobState       string `pulumi:"jobState"`
	StreamingUnits int    `pulumi:"streamingUnits"`
}

type ClusterJobResponseOutput struct{ *pulumi.OutputState }

func (ClusterJobResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterJobResponse)(nil)).Elem()
}

func (o ClusterJobResponseOutput) ToClusterJobResponseOutput() ClusterJobResponseOutput {
	return o
}

func (o ClusterJobResponseOutput) ToClusterJobResponseOutputWithContext(ctx context.Context) ClusterJobResponseOutput {
	return o
}

func (o ClusterJobResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterJobResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ClusterJobResponseOutput) JobState() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterJobResponse) string { return v.JobState }).(pulumi.StringOutput)
}

func (o ClusterJobResponseOutput) StreamingUnits() pulumi.IntOutput {
	return o.ApplyT(func(v ClusterJobResponse) int { return v.StreamingUnits }).(pulumi.IntOutput)
}

type ClusterJobResponseArrayOutput struct{ *pulumi.OutputState }

func (ClusterJobResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterJobResponse)(nil)).Elem()
}

func (o ClusterJobResponseArrayOutput) ToClusterJobResponseArrayOutput() ClusterJobResponseArrayOutput {
	return o
}

func (o ClusterJobResponseArrayOutput) ToClusterJobResponseArrayOutputWithContext(ctx context.Context) ClusterJobResponseArrayOutput {
	return o
}

func (o ClusterJobResponseArrayOutput) Index(i pulumi.IntInput) ClusterJobResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterJobResponse {
		return vs[0].([]ClusterJobResponse)[vs[1].(int)]
	}).(ClusterJobResponseOutput)
}

type ClusterSku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
}





type ClusterSkuInput interface {
	pulumi.Input

	ToClusterSkuOutput() ClusterSkuOutput
	ToClusterSkuOutputWithContext(context.Context) ClusterSkuOutput
}

type ClusterSkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
}

func (ClusterSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSku)(nil)).Elem()
}

func (i ClusterSkuArgs) ToClusterSkuOutput() ClusterSkuOutput {
	return i.ToClusterSkuOutputWithContext(context.Background())
}

func (i ClusterSkuArgs) ToClusterSkuOutputWithContext(ctx context.Context) ClusterSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuOutput)
}

func (i ClusterSkuArgs) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return i.ToClusterSkuPtrOutputWithContext(context.Background())
}

func (i ClusterSkuArgs) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuOutput).ToClusterSkuPtrOutputWithContext(ctx)
}









type ClusterSkuPtrInput interface {
	pulumi.Input

	ToClusterSkuPtrOutput() ClusterSkuPtrOutput
	ToClusterSkuPtrOutputWithContext(context.Context) ClusterSkuPtrOutput
}

type clusterSkuPtrType ClusterSkuArgs

func ClusterSkuPtr(v *ClusterSkuArgs) ClusterSkuPtrInput {
	return (*clusterSkuPtrType)(v)
}

func (*clusterSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSku)(nil)).Elem()
}

func (i *clusterSkuPtrType) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return i.ToClusterSkuPtrOutputWithContext(context.Background())
}

func (i *clusterSkuPtrType) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuPtrOutput)
}

type ClusterSkuOutput struct{ *pulumi.OutputState }

func (ClusterSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSku)(nil)).Elem()
}

func (o ClusterSkuOutput) ToClusterSkuOutput() ClusterSkuOutput {
	return o
}

func (o ClusterSkuOutput) ToClusterSkuOutputWithContext(ctx context.Context) ClusterSkuOutput {
	return o
}

func (o ClusterSkuOutput) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return o.ToClusterSkuPtrOutputWithContext(context.Background())
}

func (o ClusterSkuOutput) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterSku) *ClusterSku {
		return &v
	}).(ClusterSkuPtrOutput)
}

func (o ClusterSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ClusterSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ClusterSkuPtrOutput struct{ *pulumi.OutputState }

func (ClusterSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSku)(nil)).Elem()
}

func (o ClusterSkuPtrOutput) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return o
}

func (o ClusterSkuPtrOutput) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return o
}

func (o ClusterSkuPtrOutput) Elem() ClusterSkuOutput {
	return o.ApplyT(func(v *ClusterSku) ClusterSku {
		if v != nil {
			return *v
		}
		var ret ClusterSku
		return ret
	}).(ClusterSkuOutput)
}

func (o ClusterSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ClusterSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type ClusterSkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
}

type ClusterSkuResponseOutput struct{ *pulumi.OutputState }

func (ClusterSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSkuResponse)(nil)).Elem()
}

func (o ClusterSkuResponseOutput) ToClusterSkuResponseOutput() ClusterSkuResponseOutput {
	return o
}

func (o ClusterSkuResponseOutput) ToClusterSkuResponseOutputWithContext(ctx context.Context) ClusterSkuResponseOutput {
	return o
}

func (o ClusterSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ClusterSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ClusterSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ClusterSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSkuResponse)(nil)).Elem()
}

func (o ClusterSkuResponsePtrOutput) ToClusterSkuResponsePtrOutput() ClusterSkuResponsePtrOutput {
	return o
}

func (o ClusterSkuResponsePtrOutput) ToClusterSkuResponsePtrOutputWithContext(ctx context.Context) ClusterSkuResponsePtrOutput {
	return o
}

func (o ClusterSkuResponsePtrOutput) Elem() ClusterSkuResponseOutput {
	return o.ApplyT(func(v *ClusterSkuResponse) ClusterSkuResponse {
		if v != nil {
			return *v
		}
		var ret ClusterSkuResponse
		return ret
	}).(ClusterSkuResponseOutput)
}

func (o ClusterSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ClusterSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type Compression struct {
	Type string `pulumi:"type"`
}

type CompressionResponse struct {
	Type string `pulumi:"type"`
}

type CsvSerialization struct {
	Encoding       *string `pulumi:"encoding"`
	FieldDelimiter *string `pulumi:"fieldDelimiter"`
	Type           string  `pulumi:"type"`
}

type CsvSerializationResponse struct {
	Encoding       *string `pulumi:"encoding"`
	FieldDelimiter *string `pulumi:"fieldDelimiter"`
	Type           string  `pulumi:"type"`
}

type DiagnosticConditionResponse struct {
	Code    string `pulumi:"code"`
	Message string `pulumi:"message"`
	Since   string `pulumi:"since"`
}

type DiagnosticConditionResponseOutput struct{ *pulumi.OutputState }

func (DiagnosticConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticConditionResponse)(nil)).Elem()
}

func (o DiagnosticConditionResponseOutput) ToDiagnosticConditionResponseOutput() DiagnosticConditionResponseOutput {
	return o
}

func (o DiagnosticConditionResponseOutput) ToDiagnosticConditionResponseOutputWithContext(ctx context.Context) DiagnosticConditionResponseOutput {
	return o
}

func (o DiagnosticConditionResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticConditionResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o DiagnosticConditionResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticConditionResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o DiagnosticConditionResponseOutput) Since() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticConditionResponse) string { return v.Since }).(pulumi.StringOutput)
}

type DiagnosticConditionResponseArrayOutput struct{ *pulumi.OutputState }

func (DiagnosticConditionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiagnosticConditionResponse)(nil)).Elem()
}

func (o DiagnosticConditionResponseArrayOutput) ToDiagnosticConditionResponseArrayOutput() DiagnosticConditionResponseArrayOutput {
	return o
}

func (o DiagnosticConditionResponseArrayOutput) ToDiagnosticConditionResponseArrayOutputWithContext(ctx context.Context) DiagnosticConditionResponseArrayOutput {
	return o
}

func (o DiagnosticConditionResponseArrayOutput) Index(i pulumi.IntInput) DiagnosticConditionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiagnosticConditionResponse {
		return vs[0].([]DiagnosticConditionResponse)[vs[1].(int)]
	}).(DiagnosticConditionResponseOutput)
}

type DiagnosticsResponse struct {
	Conditions []DiagnosticConditionResponse `pulumi:"conditions"`
}

type DiagnosticsResponseOutput struct{ *pulumi.OutputState }

func (DiagnosticsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsResponse)(nil)).Elem()
}

func (o DiagnosticsResponseOutput) ToDiagnosticsResponseOutput() DiagnosticsResponseOutput {
	return o
}

func (o DiagnosticsResponseOutput) ToDiagnosticsResponseOutputWithContext(ctx context.Context) DiagnosticsResponseOutput {
	return o
}

func (o DiagnosticsResponseOutput) Conditions() DiagnosticConditionResponseArrayOutput {
	return o.ApplyT(func(v DiagnosticsResponse) []DiagnosticConditionResponse { return v.Conditions }).(DiagnosticConditionResponseArrayOutput)
}

type DocumentDbOutputDataSource struct {
	AccountId             *string `pulumi:"accountId"`
	AccountKey            *string `pulumi:"accountKey"`
	CollectionNamePattern *string `pulumi:"collectionNamePattern"`
	Database              *string `pulumi:"database"`
	DocumentId            *string `pulumi:"documentId"`
	PartitionKey          *string `pulumi:"partitionKey"`
	Type                  string  `pulumi:"type"`
}

type DocumentDbOutputDataSourceResponse struct {
	AccountId             *string `pulumi:"accountId"`
	AccountKey            *string `pulumi:"accountKey"`
	CollectionNamePattern *string `pulumi:"collectionNamePattern"`
	Database              *string `pulumi:"database"`
	DocumentId            *string `pulumi:"documentId"`
	PartitionKey          *string `pulumi:"partitionKey"`
	Type                  string  `pulumi:"type"`
}

type EventHubOutputDataSource struct {
	AuthenticationMode     *string  `pulumi:"authenticationMode"`
	EventHubName           *string  `pulumi:"eventHubName"`
	PartitionKey           *string  `pulumi:"partitionKey"`
	PropertyColumns        []string `pulumi:"propertyColumns"`
	ServiceBusNamespace    *string  `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string  `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string  `pulumi:"sharedAccessPolicyName"`
	Type                   string   `pulumi:"type"`
}

type EventHubOutputDataSourceResponse struct {
	AuthenticationMode     *string  `pulumi:"authenticationMode"`
	EventHubName           *string  `pulumi:"eventHubName"`
	PartitionKey           *string  `pulumi:"partitionKey"`
	PropertyColumns        []string `pulumi:"propertyColumns"`
	ServiceBusNamespace    *string  `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string  `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string  `pulumi:"sharedAccessPolicyName"`
	Type                   string   `pulumi:"type"`
}

type EventHubStreamInputDataSource struct {
	AuthenticationMode     *string `pulumi:"authenticationMode"`
	ConsumerGroupName      *string `pulumi:"consumerGroupName"`
	EventHubName           *string `pulumi:"eventHubName"`
	ServiceBusNamespace    *string `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	Type                   string  `pulumi:"type"`
}

type EventHubStreamInputDataSourceResponse struct {
	AuthenticationMode     *string `pulumi:"authenticationMode"`
	ConsumerGroupName      *string `pulumi:"consumerGroupName"`
	EventHubName           *string `pulumi:"eventHubName"`
	ServiceBusNamespace    *string `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	Type                   string  `pulumi:"type"`
}

type EventHubV2OutputDataSource struct {
	AuthenticationMode     *string  `pulumi:"authenticationMode"`
	EventHubName           *string  `pulumi:"eventHubName"`
	PartitionKey           *string  `pulumi:"partitionKey"`
	PropertyColumns        []string `pulumi:"propertyColumns"`
	ServiceBusNamespace    *string  `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string  `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string  `pulumi:"sharedAccessPolicyName"`
	Type                   string   `pulumi:"type"`
}

type EventHubV2OutputDataSourceResponse struct {
	AuthenticationMode     *string  `pulumi:"authenticationMode"`
	EventHubName           *string  `pulumi:"eventHubName"`
	PartitionKey           *string  `pulumi:"partitionKey"`
	PropertyColumns        []string `pulumi:"propertyColumns"`
	ServiceBusNamespace    *string  `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string  `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string  `pulumi:"sharedAccessPolicyName"`
	Type                   string   `pulumi:"type"`
}

type EventHubV2StreamInputDataSource struct {
	AuthenticationMode     *string `pulumi:"authenticationMode"`
	ConsumerGroupName      *string `pulumi:"consumerGroupName"`
	EventHubName           *string `pulumi:"eventHubName"`
	ServiceBusNamespace    *string `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	Type                   string  `pulumi:"type"`
}

type EventHubV2StreamInputDataSourceResponse struct {
	AuthenticationMode     *string `pulumi:"authenticationMode"`
	ConsumerGroupName      *string `pulumi:"consumerGroupName"`
	EventHubName           *string `pulumi:"eventHubName"`
	ServiceBusNamespace    *string `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	Type                   string  `pulumi:"type"`
}

type FileReferenceInputDataSource struct {
	Path *string `pulumi:"path"`
	Type string  `pulumi:"type"`
}

type FileReferenceInputDataSourceResponse struct {
	Path *string `pulumi:"path"`
	Type string  `pulumi:"type"`
}

type FunctionType struct {
	Name       *string     `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
}





type FunctionTypeInput interface {
	pulumi.Input

	ToFunctionTypeOutput() FunctionTypeOutput
	ToFunctionTypeOutputWithContext(context.Context) FunctionTypeOutput
}

type FunctionTypeArgs struct {
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Properties pulumi.Input          `pulumi:"properties"`
}

func (FunctionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionType)(nil)).Elem()
}

func (i FunctionTypeArgs) ToFunctionTypeOutput() FunctionTypeOutput {
	return i.ToFunctionTypeOutputWithContext(context.Background())
}

func (i FunctionTypeArgs) ToFunctionTypeOutputWithContext(ctx context.Context) FunctionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionTypeOutput)
}





type FunctionTypeArrayInput interface {
	pulumi.Input

	ToFunctionTypeArrayOutput() FunctionTypeArrayOutput
	ToFunctionTypeArrayOutputWithContext(context.Context) FunctionTypeArrayOutput
}

type FunctionTypeArray []FunctionTypeInput

func (FunctionTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionType)(nil)).Elem()
}

func (i FunctionTypeArray) ToFunctionTypeArrayOutput() FunctionTypeArrayOutput {
	return i.ToFunctionTypeArrayOutputWithContext(context.Background())
}

func (i FunctionTypeArray) ToFunctionTypeArrayOutputWithContext(ctx context.Context) FunctionTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionTypeArrayOutput)
}

type FunctionTypeOutput struct{ *pulumi.OutputState }

func (FunctionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionType)(nil)).Elem()
}

func (o FunctionTypeOutput) ToFunctionTypeOutput() FunctionTypeOutput {
	return o
}

func (o FunctionTypeOutput) ToFunctionTypeOutputWithContext(ctx context.Context) FunctionTypeOutput {
	return o
}

func (o FunctionTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FunctionType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FunctionTypeOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v FunctionType) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

type FunctionTypeArrayOutput struct{ *pulumi.OutputState }

func (FunctionTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionType)(nil)).Elem()
}

func (o FunctionTypeArrayOutput) ToFunctionTypeArrayOutput() FunctionTypeArrayOutput {
	return o
}

func (o FunctionTypeArrayOutput) ToFunctionTypeArrayOutputWithContext(ctx context.Context) FunctionTypeArrayOutput {
	return o
}

func (o FunctionTypeArrayOutput) Index(i pulumi.IntInput) FunctionTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FunctionType {
		return vs[0].([]FunctionType)[vs[1].(int)]
	}).(FunctionTypeOutput)
}

type FunctionInputType struct {
	DataType                 *string `pulumi:"dataType"`
	IsConfigurationParameter *bool   `pulumi:"isConfigurationParameter"`
}

type FunctionInputResponse struct {
	DataType                 *string `pulumi:"dataType"`
	IsConfigurationParameter *bool   `pulumi:"isConfigurationParameter"`
}

type FunctionOutputType struct {
	DataType *string `pulumi:"dataType"`
}

type FunctionOutputResponse struct {
	DataType *string `pulumi:"dataType"`
}

type FunctionResponse struct {
	Id         string      `pulumi:"id"`
	Name       *string     `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}

type FunctionResponseOutput struct{ *pulumi.OutputState }

func (FunctionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionResponse)(nil)).Elem()
}

func (o FunctionResponseOutput) ToFunctionResponseOutput() FunctionResponseOutput {
	return o
}

func (o FunctionResponseOutput) ToFunctionResponseOutputWithContext(ctx context.Context) FunctionResponseOutput {
	return o
}

func (o FunctionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v FunctionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o FunctionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FunctionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FunctionResponseOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v FunctionResponse) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o FunctionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v FunctionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type FunctionResponseArrayOutput struct{ *pulumi.OutputState }

func (FunctionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionResponse)(nil)).Elem()
}

func (o FunctionResponseArrayOutput) ToFunctionResponseArrayOutput() FunctionResponseArrayOutput {
	return o
}

func (o FunctionResponseArrayOutput) ToFunctionResponseArrayOutputWithContext(ctx context.Context) FunctionResponseArrayOutput {
	return o
}

func (o FunctionResponseArrayOutput) Index(i pulumi.IntInput) FunctionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FunctionResponse {
		return vs[0].([]FunctionResponse)[vs[1].(int)]
	}).(FunctionResponseOutput)
}

type GatewayMessageBusOutputDataSource struct {
	Topic *string `pulumi:"topic"`
	Type  string  `pulumi:"type"`
}

type GatewayMessageBusOutputDataSourceResponse struct {
	Topic *string `pulumi:"topic"`
	Type  string  `pulumi:"type"`
}

type GatewayMessageBusStreamInputDataSource struct {
	Topic *string `pulumi:"topic"`
	Type  string  `pulumi:"type"`
}

type GatewayMessageBusStreamInputDataSourceResponse struct {
	Topic *string `pulumi:"topic"`
	Type  string  `pulumi:"type"`
}

type Identity struct {
	Type *string `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Identity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type IdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type InputType struct {
	Name       *string     `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
}





type InputTypeInput interface {
	pulumi.Input

	ToInputTypeOutput() InputTypeOutput
	ToInputTypeOutputWithContext(context.Context) InputTypeOutput
}

type InputTypeArgs struct {
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Properties pulumi.Input          `pulumi:"properties"`
}

func (InputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InputType)(nil)).Elem()
}

func (i InputTypeArgs) ToInputTypeOutput() InputTypeOutput {
	return i.ToInputTypeOutputWithContext(context.Background())
}

func (i InputTypeArgs) ToInputTypeOutputWithContext(ctx context.Context) InputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputTypeOutput)
}





type InputTypeArrayInput interface {
	pulumi.Input

	ToInputTypeArrayOutput() InputTypeArrayOutput
	ToInputTypeArrayOutputWithContext(context.Context) InputTypeArrayOutput
}

type InputTypeArray []InputTypeInput

func (InputTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InputType)(nil)).Elem()
}

func (i InputTypeArray) ToInputTypeArrayOutput() InputTypeArrayOutput {
	return i.ToInputTypeArrayOutputWithContext(context.Background())
}

func (i InputTypeArray) ToInputTypeArrayOutputWithContext(ctx context.Context) InputTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputTypeArrayOutput)
}

type InputTypeOutput struct{ *pulumi.OutputState }

func (InputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputType)(nil)).Elem()
}

func (o InputTypeOutput) ToInputTypeOutput() InputTypeOutput {
	return o
}

func (o InputTypeOutput) ToInputTypeOutputWithContext(ctx context.Context) InputTypeOutput {
	return o
}

func (o InputTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InputType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o InputTypeOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v InputType) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

type InputTypeArrayOutput struct{ *pulumi.OutputState }

func (InputTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InputType)(nil)).Elem()
}

func (o InputTypeArrayOutput) ToInputTypeArrayOutput() InputTypeArrayOutput {
	return o
}

func (o InputTypeArrayOutput) ToInputTypeArrayOutputWithContext(ctx context.Context) InputTypeArrayOutput {
	return o
}

func (o InputTypeArrayOutput) Index(i pulumi.IntInput) InputTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InputType {
		return vs[0].([]InputType)[vs[1].(int)]
	}).(InputTypeOutput)
}

type InputResponse struct {
	Id         string      `pulumi:"id"`
	Name       *string     `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}

type InputResponseOutput struct{ *pulumi.OutputState }

func (InputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputResponse)(nil)).Elem()
}

func (o InputResponseOutput) ToInputResponseOutput() InputResponseOutput {
	return o
}

func (o InputResponseOutput) ToInputResponseOutputWithContext(ctx context.Context) InputResponseOutput {
	return o
}

func (o InputResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v InputResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o InputResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InputResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o InputResponseOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v InputResponse) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o InputResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v InputResponse) string { return v.Type }).(pulumi.StringOutput)
}

type InputResponseArrayOutput struct{ *pulumi.OutputState }

func (InputResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InputResponse)(nil)).Elem()
}

func (o InputResponseArrayOutput) ToInputResponseArrayOutput() InputResponseArrayOutput {
	return o
}

func (o InputResponseArrayOutput) ToInputResponseArrayOutputWithContext(ctx context.Context) InputResponseArrayOutput {
	return o
}

func (o InputResponseArrayOutput) Index(i pulumi.IntInput) InputResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InputResponse {
		return vs[0].([]InputResponse)[vs[1].(int)]
	}).(InputResponseOutput)
}

type IoTHubStreamInputDataSource struct {
	ConsumerGroupName      *string `pulumi:"consumerGroupName"`
	Endpoint               *string `pulumi:"endpoint"`
	IotHubNamespace        *string `pulumi:"iotHubNamespace"`
	SharedAccessPolicyKey  *string `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	Type                   string  `pulumi:"type"`
}

type IoTHubStreamInputDataSourceResponse struct {
	ConsumerGroupName      *string `pulumi:"consumerGroupName"`
	Endpoint               *string `pulumi:"endpoint"`
	IotHubNamespace        *string `pulumi:"iotHubNamespace"`
	SharedAccessPolicyKey  *string `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	Type                   string  `pulumi:"type"`
}

type JavaScriptFunctionBinding struct {
	Script *string `pulumi:"script"`
	Type   string  `pulumi:"type"`
}

type JavaScriptFunctionBindingResponse struct {
	Script *string `pulumi:"script"`
	Type   string  `pulumi:"type"`
}

type JobStorageAccount struct {
	AccountKey         *string `pulumi:"accountKey"`
	AccountName        *string `pulumi:"accountName"`
	AuthenticationMode *string `pulumi:"authenticationMode"`
}





type JobStorageAccountInput interface {
	pulumi.Input

	ToJobStorageAccountOutput() JobStorageAccountOutput
	ToJobStorageAccountOutputWithContext(context.Context) JobStorageAccountOutput
}

type JobStorageAccountArgs struct {
	AccountKey         pulumi.StringPtrInput `pulumi:"accountKey"`
	AccountName        pulumi.StringPtrInput `pulumi:"accountName"`
	AuthenticationMode pulumi.StringPtrInput `pulumi:"authenticationMode"`
}

func (JobStorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStorageAccount)(nil)).Elem()
}

func (i JobStorageAccountArgs) ToJobStorageAccountOutput() JobStorageAccountOutput {
	return i.ToJobStorageAccountOutputWithContext(context.Background())
}

func (i JobStorageAccountArgs) ToJobStorageAccountOutputWithContext(ctx context.Context) JobStorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStorageAccountOutput)
}

func (i JobStorageAccountArgs) ToJobStorageAccountPtrOutput() JobStorageAccountPtrOutput {
	return i.ToJobStorageAccountPtrOutputWithContext(context.Background())
}

func (i JobStorageAccountArgs) ToJobStorageAccountPtrOutputWithContext(ctx context.Context) JobStorageAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStorageAccountOutput).ToJobStorageAccountPtrOutputWithContext(ctx)
}









type JobStorageAccountPtrInput interface {
	pulumi.Input

	ToJobStorageAccountPtrOutput() JobStorageAccountPtrOutput
	ToJobStorageAccountPtrOutputWithContext(context.Context) JobStorageAccountPtrOutput
}

type jobStorageAccountPtrType JobStorageAccountArgs

func JobStorageAccountPtr(v *JobStorageAccountArgs) JobStorageAccountPtrInput {
	return (*jobStorageAccountPtrType)(v)
}

func (*jobStorageAccountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStorageAccount)(nil)).Elem()
}

func (i *jobStorageAccountPtrType) ToJobStorageAccountPtrOutput() JobStorageAccountPtrOutput {
	return i.ToJobStorageAccountPtrOutputWithContext(context.Background())
}

func (i *jobStorageAccountPtrType) ToJobStorageAccountPtrOutputWithContext(ctx context.Context) JobStorageAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStorageAccountPtrOutput)
}

type JobStorageAccountOutput struct{ *pulumi.OutputState }

func (JobStorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStorageAccount)(nil)).Elem()
}

func (o JobStorageAccountOutput) ToJobStorageAccountOutput() JobStorageAccountOutput {
	return o
}

func (o JobStorageAccountOutput) ToJobStorageAccountOutputWithContext(ctx context.Context) JobStorageAccountOutput {
	return o
}

func (o JobStorageAccountOutput) ToJobStorageAccountPtrOutput() JobStorageAccountPtrOutput {
	return o.ToJobStorageAccountPtrOutputWithContext(context.Background())
}

func (o JobStorageAccountOutput) ToJobStorageAccountPtrOutputWithContext(ctx context.Context) JobStorageAccountPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStorageAccount) *JobStorageAccount {
		return &v
	}).(JobStorageAccountPtrOutput)
}

func (o JobStorageAccountOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStorageAccount) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o JobStorageAccountOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStorageAccount) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o JobStorageAccountOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStorageAccount) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

type JobStorageAccountPtrOutput struct{ *pulumi.OutputState }

func (JobStorageAccountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStorageAccount)(nil)).Elem()
}

func (o JobStorageAccountPtrOutput) ToJobStorageAccountPtrOutput() JobStorageAccountPtrOutput {
	return o
}

func (o JobStorageAccountPtrOutput) ToJobStorageAccountPtrOutputWithContext(ctx context.Context) JobStorageAccountPtrOutput {
	return o
}

func (o JobStorageAccountPtrOutput) Elem() JobStorageAccountOutput {
	return o.ApplyT(func(v *JobStorageAccount) JobStorageAccount {
		if v != nil {
			return *v
		}
		var ret JobStorageAccount
		return ret
	}).(JobStorageAccountOutput)
}

func (o JobStorageAccountPtrOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStorageAccount) *string {
		if v == nil {
			return nil
		}
		return v.AccountKey
	}).(pulumi.StringPtrOutput)
}

func (o JobStorageAccountPtrOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStorageAccount) *string {
		if v == nil {
			return nil
		}
		return v.AccountName
	}).(pulumi.StringPtrOutput)
}

func (o JobStorageAccountPtrOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStorageAccount) *string {
		if v == nil {
			return nil
		}
		return v.AuthenticationMode
	}).(pulumi.StringPtrOutput)
}

type JobStorageAccountResponse struct {
	AccountKey         *string `pulumi:"accountKey"`
	AccountName        *string `pulumi:"accountName"`
	AuthenticationMode *string `pulumi:"authenticationMode"`
}

type JobStorageAccountResponseOutput struct{ *pulumi.OutputState }

func (JobStorageAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStorageAccountResponse)(nil)).Elem()
}

func (o JobStorageAccountResponseOutput) ToJobStorageAccountResponseOutput() JobStorageAccountResponseOutput {
	return o
}

func (o JobStorageAccountResponseOutput) ToJobStorageAccountResponseOutputWithContext(ctx context.Context) JobStorageAccountResponseOutput {
	return o
}

func (o JobStorageAccountResponseOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStorageAccountResponse) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o JobStorageAccountResponseOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStorageAccountResponse) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o JobStorageAccountResponseOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStorageAccountResponse) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

type JobStorageAccountResponsePtrOutput struct{ *pulumi.OutputState }

func (JobStorageAccountResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStorageAccountResponse)(nil)).Elem()
}

func (o JobStorageAccountResponsePtrOutput) ToJobStorageAccountResponsePtrOutput() JobStorageAccountResponsePtrOutput {
	return o
}

func (o JobStorageAccountResponsePtrOutput) ToJobStorageAccountResponsePtrOutputWithContext(ctx context.Context) JobStorageAccountResponsePtrOutput {
	return o
}

func (o JobStorageAccountResponsePtrOutput) Elem() JobStorageAccountResponseOutput {
	return o.ApplyT(func(v *JobStorageAccountResponse) JobStorageAccountResponse {
		if v != nil {
			return *v
		}
		var ret JobStorageAccountResponse
		return ret
	}).(JobStorageAccountResponseOutput)
}

func (o JobStorageAccountResponsePtrOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStorageAccountResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccountKey
	}).(pulumi.StringPtrOutput)
}

func (o JobStorageAccountResponsePtrOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStorageAccountResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccountName
	}).(pulumi.StringPtrOutput)
}

func (o JobStorageAccountResponsePtrOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStorageAccountResponse) *string {
		if v == nil {
			return nil
		}
		return v.AuthenticationMode
	}).(pulumi.StringPtrOutput)
}

type JsonSerialization struct {
	Encoding *string `pulumi:"encoding"`
	Format   *string `pulumi:"format"`
	Type     string  `pulumi:"type"`
}

type JsonSerializationResponse struct {
	Encoding *string `pulumi:"encoding"`
	Format   *string `pulumi:"format"`
	Type     string  `pulumi:"type"`
}

type OutputType struct {
	Datasource    interface{} `pulumi:"datasource"`
	Name          *string     `pulumi:"name"`
	Serialization interface{} `pulumi:"serialization"`
	SizeWindow    *float64    `pulumi:"sizeWindow"`
	TimeWindow    *string     `pulumi:"timeWindow"`
}





type OutputTypeInput interface {
	pulumi.Input

	ToOutputTypeOutput() OutputTypeOutput
	ToOutputTypeOutputWithContext(context.Context) OutputTypeOutput
}

type OutputTypeArgs struct {
	Datasource    pulumi.Input           `pulumi:"datasource"`
	Name          pulumi.StringPtrInput  `pulumi:"name"`
	Serialization pulumi.Input           `pulumi:"serialization"`
	SizeWindow    pulumi.Float64PtrInput `pulumi:"sizeWindow"`
	TimeWindow    pulumi.StringPtrInput  `pulumi:"timeWindow"`
}

func (OutputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputType)(nil)).Elem()
}

func (i OutputTypeArgs) ToOutputTypeOutput() OutputTypeOutput {
	return i.ToOutputTypeOutputWithContext(context.Background())
}

func (i OutputTypeArgs) ToOutputTypeOutputWithContext(ctx context.Context) OutputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputTypeOutput)
}





type OutputTypeArrayInput interface {
	pulumi.Input

	ToOutputTypeArrayOutput() OutputTypeArrayOutput
	ToOutputTypeArrayOutputWithContext(context.Context) OutputTypeArrayOutput
}

type OutputTypeArray []OutputTypeInput

func (OutputTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutputType)(nil)).Elem()
}

func (i OutputTypeArray) ToOutputTypeArrayOutput() OutputTypeArrayOutput {
	return i.ToOutputTypeArrayOutputWithContext(context.Background())
}

func (i OutputTypeArray) ToOutputTypeArrayOutputWithContext(ctx context.Context) OutputTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputTypeArrayOutput)
}

type OutputTypeOutput struct{ *pulumi.OutputState }

func (OutputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputType)(nil)).Elem()
}

func (o OutputTypeOutput) ToOutputTypeOutput() OutputTypeOutput {
	return o
}

func (o OutputTypeOutput) ToOutputTypeOutputWithContext(ctx context.Context) OutputTypeOutput {
	return o
}

func (o OutputTypeOutput) Datasource() pulumi.AnyOutput {
	return o.ApplyT(func(v OutputType) interface{} { return v.Datasource }).(pulumi.AnyOutput)
}

func (o OutputTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutputType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OutputTypeOutput) Serialization() pulumi.AnyOutput {
	return o.ApplyT(func(v OutputType) interface{} { return v.Serialization }).(pulumi.AnyOutput)
}

func (o OutputTypeOutput) SizeWindow() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v OutputType) *float64 { return v.SizeWindow }).(pulumi.Float64PtrOutput)
}

func (o OutputTypeOutput) TimeWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutputType) *string { return v.TimeWindow }).(pulumi.StringPtrOutput)
}

type OutputTypeArrayOutput struct{ *pulumi.OutputState }

func (OutputTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutputType)(nil)).Elem()
}

func (o OutputTypeArrayOutput) ToOutputTypeArrayOutput() OutputTypeArrayOutput {
	return o
}

func (o OutputTypeArrayOutput) ToOutputTypeArrayOutputWithContext(ctx context.Context) OutputTypeArrayOutput {
	return o
}

func (o OutputTypeArrayOutput) Index(i pulumi.IntInput) OutputTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OutputType {
		return vs[0].([]OutputType)[vs[1].(int)]
	}).(OutputTypeOutput)
}

type OutputResponse struct {
	Datasource    interface{}         `pulumi:"datasource"`
	Diagnostics   DiagnosticsResponse `pulumi:"diagnostics"`
	Etag          string              `pulumi:"etag"`
	Id            string              `pulumi:"id"`
	Name          *string             `pulumi:"name"`
	Serialization interface{}         `pulumi:"serialization"`
	SizeWindow    *float64            `pulumi:"sizeWindow"`
	TimeWindow    *string             `pulumi:"timeWindow"`
	Type          string              `pulumi:"type"`
}

type OutputResponseOutput struct{ *pulumi.OutputState }

func (OutputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputResponse)(nil)).Elem()
}

func (o OutputResponseOutput) ToOutputResponseOutput() OutputResponseOutput {
	return o
}

func (o OutputResponseOutput) ToOutputResponseOutputWithContext(ctx context.Context) OutputResponseOutput {
	return o
}

func (o OutputResponseOutput) Datasource() pulumi.AnyOutput {
	return o.ApplyT(func(v OutputResponse) interface{} { return v.Datasource }).(pulumi.AnyOutput)
}

func (o OutputResponseOutput) Diagnostics() DiagnosticsResponseOutput {
	return o.ApplyT(func(v OutputResponse) DiagnosticsResponse { return v.Diagnostics }).(DiagnosticsResponseOutput)
}

func (o OutputResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v OutputResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o OutputResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v OutputResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o OutputResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutputResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OutputResponseOutput) Serialization() pulumi.AnyOutput {
	return o.ApplyT(func(v OutputResponse) interface{} { return v.Serialization }).(pulumi.AnyOutput)
}

func (o OutputResponseOutput) SizeWindow() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v OutputResponse) *float64 { return v.SizeWindow }).(pulumi.Float64PtrOutput)
}

func (o OutputResponseOutput) TimeWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutputResponse) *string { return v.TimeWindow }).(pulumi.StringPtrOutput)
}

func (o OutputResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v OutputResponse) string { return v.Type }).(pulumi.StringOutput)
}

type OutputResponseArrayOutput struct{ *pulumi.OutputState }

func (OutputResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutputResponse)(nil)).Elem()
}

func (o OutputResponseArrayOutput) ToOutputResponseArrayOutput() OutputResponseArrayOutput {
	return o
}

func (o OutputResponseArrayOutput) ToOutputResponseArrayOutputWithContext(ctx context.Context) OutputResponseArrayOutput {
	return o
}

func (o OutputResponseArrayOutput) Index(i pulumi.IntInput) OutputResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OutputResponse {
		return vs[0].([]OutputResponse)[vs[1].(int)]
	}).(OutputResponseOutput)
}

type ParquetSerialization struct {
	Type string `pulumi:"type"`
}

type ParquetSerializationResponse struct {
	Type string `pulumi:"type"`
}

type PowerBIOutputDataSource struct {
	AuthenticationMode     *string `pulumi:"authenticationMode"`
	Dataset                *string `pulumi:"dataset"`
	GroupId                *string `pulumi:"groupId"`
	GroupName              *string `pulumi:"groupName"`
	RefreshToken           *string `pulumi:"refreshToken"`
	Table                  *string `pulumi:"table"`
	TokenUserDisplayName   *string `pulumi:"tokenUserDisplayName"`
	TokenUserPrincipalName *string `pulumi:"tokenUserPrincipalName"`
	Type                   string  `pulumi:"type"`
}

type PowerBIOutputDataSourceResponse struct {
	AuthenticationMode     *string `pulumi:"authenticationMode"`
	Dataset                *string `pulumi:"dataset"`
	GroupId                *string `pulumi:"groupId"`
	GroupName              *string `pulumi:"groupName"`
	RefreshToken           *string `pulumi:"refreshToken"`
	Table                  *string `pulumi:"table"`
	TokenUserDisplayName   *string `pulumi:"tokenUserDisplayName"`
	TokenUserPrincipalName *string `pulumi:"tokenUserPrincipalName"`
	Type                   string  `pulumi:"type"`
}

type PrivateLinkConnectionStateResponse struct {
	ActionsRequired string `pulumi:"actionsRequired"`
	Description     string `pulumi:"description"`
	Status          string `pulumi:"status"`
}

type PrivateLinkConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkConnectionStateResponseOutput) ToPrivateLinkConnectionStateResponseOutput() PrivateLinkConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkConnectionStateResponseOutput) ToPrivateLinkConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkConnectionStateResponseOutput) ActionsRequired() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkConnectionStateResponse) string { return v.ActionsRequired }).(pulumi.StringOutput)
}

func (o PrivateLinkConnectionStateResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkConnectionStateResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o PrivateLinkConnectionStateResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkConnectionStateResponse) string { return v.Status }).(pulumi.StringOutput)
}

type PrivateLinkConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkConnectionStateResponsePtrOutput) ToPrivateLinkConnectionStateResponsePtrOutput() PrivateLinkConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkConnectionStateResponsePtrOutput) ToPrivateLinkConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkConnectionStateResponsePtrOutput) Elem() PrivateLinkConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) PrivateLinkConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkConnectionStateResponse
		return ret
	}).(PrivateLinkConnectionStateResponseOutput)
}

func (o PrivateLinkConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnection struct {
	GroupIds             []string `pulumi:"groupIds"`
	PrivateLinkServiceId *string  `pulumi:"privateLinkServiceId"`
}





type PrivateLinkServiceConnectionInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionOutput() PrivateLinkServiceConnectionOutput
	ToPrivateLinkServiceConnectionOutputWithContext(context.Context) PrivateLinkServiceConnectionOutput
}

type PrivateLinkServiceConnectionArgs struct {
	GroupIds             pulumi.StringArrayInput `pulumi:"groupIds"`
	PrivateLinkServiceId pulumi.StringPtrInput   `pulumi:"privateLinkServiceId"`
}

func (PrivateLinkServiceConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnection)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionArgs) ToPrivateLinkServiceConnectionOutput() PrivateLinkServiceConnectionOutput {
	return i.ToPrivateLinkServiceConnectionOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionArgs) ToPrivateLinkServiceConnectionOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionOutput)
}





type PrivateLinkServiceConnectionArrayInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionArrayOutput() PrivateLinkServiceConnectionArrayOutput
	ToPrivateLinkServiceConnectionArrayOutputWithContext(context.Context) PrivateLinkServiceConnectionArrayOutput
}

type PrivateLinkServiceConnectionArray []PrivateLinkServiceConnectionInput

func (PrivateLinkServiceConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkServiceConnection)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionArray) ToPrivateLinkServiceConnectionArrayOutput() PrivateLinkServiceConnectionArrayOutput {
	return i.ToPrivateLinkServiceConnectionArrayOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionArray) ToPrivateLinkServiceConnectionArrayOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionArrayOutput)
}

type PrivateLinkServiceConnectionOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnection)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionOutput) ToPrivateLinkServiceConnectionOutput() PrivateLinkServiceConnectionOutput {
	return o
}

func (o PrivateLinkServiceConnectionOutput) ToPrivateLinkServiceConnectionOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionOutput {
	return o
}

func (o PrivateLinkServiceConnectionOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnection) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o PrivateLinkServiceConnectionOutput) PrivateLinkServiceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnection) *string { return v.PrivateLinkServiceId }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionArrayOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkServiceConnection)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionArrayOutput) ToPrivateLinkServiceConnectionArrayOutput() PrivateLinkServiceConnectionArrayOutput {
	return o
}

func (o PrivateLinkServiceConnectionArrayOutput) ToPrivateLinkServiceConnectionArrayOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionArrayOutput {
	return o
}

func (o PrivateLinkServiceConnectionArrayOutput) Index(i pulumi.IntInput) PrivateLinkServiceConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateLinkServiceConnection {
		return vs[0].([]PrivateLinkServiceConnection)[vs[1].(int)]
	}).(PrivateLinkServiceConnectionOutput)
}

type PrivateLinkServiceConnectionResponse struct {
	GroupIds                          []string                            `pulumi:"groupIds"`
	PrivateLinkServiceConnectionState *PrivateLinkConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	PrivateLinkServiceId              *string                             `pulumi:"privateLinkServiceId"`
	RequestMessage                    string                              `pulumi:"requestMessage"`
}

type PrivateLinkServiceConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionResponseOutput) ToPrivateLinkServiceConnectionResponseOutput() PrivateLinkServiceConnectionResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionResponseOutput) ToPrivateLinkServiceConnectionResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o PrivateLinkServiceConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionResponse) *PrivateLinkConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStateResponsePtrOutput)
}

func (o PrivateLinkServiceConnectionResponseOutput) PrivateLinkServiceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionResponse) *string { return v.PrivateLinkServiceId }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionResponseOutput) RequestMessage() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionResponse) string { return v.RequestMessage }).(pulumi.StringOutput)
}

type PrivateLinkServiceConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkServiceConnectionResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionResponseArrayOutput) ToPrivateLinkServiceConnectionResponseArrayOutput() PrivateLinkServiceConnectionResponseArrayOutput {
	return o
}

func (o PrivateLinkServiceConnectionResponseArrayOutput) ToPrivateLinkServiceConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionResponseArrayOutput {
	return o
}

func (o PrivateLinkServiceConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateLinkServiceConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateLinkServiceConnectionResponse {
		return vs[0].([]PrivateLinkServiceConnectionResponse)[vs[1].(int)]
	}).(PrivateLinkServiceConnectionResponseOutput)
}

type ReferenceInputProperties struct {
	Compression   *Compression `pulumi:"compression"`
	Datasource    interface{}  `pulumi:"datasource"`
	PartitionKey  *string      `pulumi:"partitionKey"`
	Serialization interface{}  `pulumi:"serialization"`
	Type          string       `pulumi:"type"`
}

type ReferenceInputPropertiesResponse struct {
	Compression   *CompressionResponse `pulumi:"compression"`
	Datasource    interface{}          `pulumi:"datasource"`
	Diagnostics   DiagnosticsResponse  `pulumi:"diagnostics"`
	Etag          string               `pulumi:"etag"`
	PartitionKey  *string              `pulumi:"partitionKey"`
	Serialization interface{}          `pulumi:"serialization"`
	Type          string               `pulumi:"type"`
}

type ScalarFunctionProperties struct {
	Binding interface{}         `pulumi:"binding"`
	Inputs  []FunctionInputType `pulumi:"inputs"`
	Output  *FunctionOutputType `pulumi:"output"`
	Type    string              `pulumi:"type"`
}

type ScalarFunctionPropertiesResponse struct {
	Binding interface{}             `pulumi:"binding"`
	Etag    string                  `pulumi:"etag"`
	Inputs  []FunctionInputResponse `pulumi:"inputs"`
	Output  *FunctionOutputResponse `pulumi:"output"`
	Type    string                  `pulumi:"type"`
}

type ServiceBusQueueOutputDataSource struct {
	AuthenticationMode     *string     `pulumi:"authenticationMode"`
	PropertyColumns        []string    `pulumi:"propertyColumns"`
	QueueName              *string     `pulumi:"queueName"`
	ServiceBusNamespace    *string     `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string     `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string     `pulumi:"sharedAccessPolicyName"`
	SystemPropertyColumns  interface{} `pulumi:"systemPropertyColumns"`
	Type                   string      `pulumi:"type"`
}

type ServiceBusQueueOutputDataSourceResponse struct {
	AuthenticationMode     *string     `pulumi:"authenticationMode"`
	PropertyColumns        []string    `pulumi:"propertyColumns"`
	QueueName              *string     `pulumi:"queueName"`
	ServiceBusNamespace    *string     `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string     `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string     `pulumi:"sharedAccessPolicyName"`
	SystemPropertyColumns  interface{} `pulumi:"systemPropertyColumns"`
	Type                   string      `pulumi:"type"`
}

type ServiceBusTopicOutputDataSource struct {
	AuthenticationMode     *string           `pulumi:"authenticationMode"`
	PropertyColumns        []string          `pulumi:"propertyColumns"`
	ServiceBusNamespace    *string           `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string           `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string           `pulumi:"sharedAccessPolicyName"`
	SystemPropertyColumns  map[string]string `pulumi:"systemPropertyColumns"`
	TopicName              *string           `pulumi:"topicName"`
	Type                   string            `pulumi:"type"`
}

type ServiceBusTopicOutputDataSourceResponse struct {
	AuthenticationMode     *string           `pulumi:"authenticationMode"`
	PropertyColumns        []string          `pulumi:"propertyColumns"`
	ServiceBusNamespace    *string           `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string           `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string           `pulumi:"sharedAccessPolicyName"`
	SystemPropertyColumns  map[string]string `pulumi:"systemPropertyColumns"`
	TopicName              *string           `pulumi:"topicName"`
	Type                   string            `pulumi:"type"`
}

type Sku struct {
	Name *string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name *string `pulumi:"name"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type StorageAccount struct {
	AccountKey  *string `pulumi:"accountKey"`
	AccountName *string `pulumi:"accountName"`
}

type StorageAccountResponse struct {
	AccountKey  *string `pulumi:"accountKey"`
	AccountName *string `pulumi:"accountName"`
}

type StreamInputProperties struct {
	Compression   *Compression `pulumi:"compression"`
	Datasource    interface{}  `pulumi:"datasource"`
	PartitionKey  *string      `pulumi:"partitionKey"`
	Serialization interface{}  `pulumi:"serialization"`
	Type          string       `pulumi:"type"`
}

type StreamInputPropertiesResponse struct {
	Compression   *CompressionResponse `pulumi:"compression"`
	Datasource    interface{}          `pulumi:"datasource"`
	Diagnostics   DiagnosticsResponse  `pulumi:"diagnostics"`
	Etag          string               `pulumi:"etag"`
	PartitionKey  *string              `pulumi:"partitionKey"`
	Serialization interface{}          `pulumi:"serialization"`
	Type          string               `pulumi:"type"`
}

type Transformation struct {
	Name                *string `pulumi:"name"`
	Query               *string `pulumi:"query"`
	StreamingUnits      *int    `pulumi:"streamingUnits"`
	ValidStreamingUnits []int   `pulumi:"validStreamingUnits"`
}


func (val *Transformation) Defaults() *Transformation {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.StreamingUnits) {
		streamingUnits_ := 3
		tmp.StreamingUnits = &streamingUnits_
	}
	return &tmp
}





type TransformationInput interface {
	pulumi.Input

	ToTransformationOutput() TransformationOutput
	ToTransformationOutputWithContext(context.Context) TransformationOutput
}

type TransformationArgs struct {
	Name                pulumi.StringPtrInput `pulumi:"name"`
	Query               pulumi.StringPtrInput `pulumi:"query"`
	StreamingUnits      pulumi.IntPtrInput    `pulumi:"streamingUnits"`
	ValidStreamingUnits pulumi.IntArrayInput  `pulumi:"validStreamingUnits"`
}


func (val *TransformationArgs) Defaults() *TransformationArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.StreamingUnits) {
		tmp.StreamingUnits = pulumi.IntPtr(3)
	}
	return &tmp
}
func (TransformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Transformation)(nil)).Elem()
}

func (i TransformationArgs) ToTransformationOutput() TransformationOutput {
	return i.ToTransformationOutputWithContext(context.Background())
}

func (i TransformationArgs) ToTransformationOutputWithContext(ctx context.Context) TransformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformationOutput)
}

func (i TransformationArgs) ToTransformationPtrOutput() TransformationPtrOutput {
	return i.ToTransformationPtrOutputWithContext(context.Background())
}

func (i TransformationArgs) ToTransformationPtrOutputWithContext(ctx context.Context) TransformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformationOutput).ToTransformationPtrOutputWithContext(ctx)
}









type TransformationPtrInput interface {
	pulumi.Input

	ToTransformationPtrOutput() TransformationPtrOutput
	ToTransformationPtrOutputWithContext(context.Context) TransformationPtrOutput
}

type transformationPtrType TransformationArgs

func TransformationPtr(v *TransformationArgs) TransformationPtrInput {
	return (*transformationPtrType)(v)
}

func (*transformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Transformation)(nil)).Elem()
}

func (i *transformationPtrType) ToTransformationPtrOutput() TransformationPtrOutput {
	return i.ToTransformationPtrOutputWithContext(context.Background())
}

func (i *transformationPtrType) ToTransformationPtrOutputWithContext(ctx context.Context) TransformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformationPtrOutput)
}

type TransformationOutput struct{ *pulumi.OutputState }

func (TransformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Transformation)(nil)).Elem()
}

func (o TransformationOutput) ToTransformationOutput() TransformationOutput {
	return o
}

func (o TransformationOutput) ToTransformationOutputWithContext(ctx context.Context) TransformationOutput {
	return o
}

func (o TransformationOutput) ToTransformationPtrOutput() TransformationPtrOutput {
	return o.ToTransformationPtrOutputWithContext(context.Background())
}

func (o TransformationOutput) ToTransformationPtrOutputWithContext(ctx context.Context) TransformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Transformation) *Transformation {
		return &v
	}).(TransformationPtrOutput)
}

func (o TransformationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Transformation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TransformationOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Transformation) *string { return v.Query }).(pulumi.StringPtrOutput)
}

func (o TransformationOutput) StreamingUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Transformation) *int { return v.StreamingUnits }).(pulumi.IntPtrOutput)
}

func (o TransformationOutput) ValidStreamingUnits() pulumi.IntArrayOutput {
	return o.ApplyT(func(v Transformation) []int { return v.ValidStreamingUnits }).(pulumi.IntArrayOutput)
}

type TransformationPtrOutput struct{ *pulumi.OutputState }

func (TransformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Transformation)(nil)).Elem()
}

func (o TransformationPtrOutput) ToTransformationPtrOutput() TransformationPtrOutput {
	return o
}

func (o TransformationPtrOutput) ToTransformationPtrOutputWithContext(ctx context.Context) TransformationPtrOutput {
	return o
}

func (o TransformationPtrOutput) Elem() TransformationOutput {
	return o.ApplyT(func(v *Transformation) Transformation {
		if v != nil {
			return *v
		}
		var ret Transformation
		return ret
	}).(TransformationOutput)
}

func (o TransformationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Transformation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o TransformationPtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Transformation) *string {
		if v == nil {
			return nil
		}
		return v.Query
	}).(pulumi.StringPtrOutput)
}

func (o TransformationPtrOutput) StreamingUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Transformation) *int {
		if v == nil {
			return nil
		}
		return v.StreamingUnits
	}).(pulumi.IntPtrOutput)
}

func (o TransformationPtrOutput) ValidStreamingUnits() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Transformation) []int {
		if v == nil {
			return nil
		}
		return v.ValidStreamingUnits
	}).(pulumi.IntArrayOutput)
}

type TransformationResponse struct {
	Etag                string  `pulumi:"etag"`
	Id                  string  `pulumi:"id"`
	Name                *string `pulumi:"name"`
	Query               *string `pulumi:"query"`
	StreamingUnits      *int    `pulumi:"streamingUnits"`
	Type                string  `pulumi:"type"`
	ValidStreamingUnits []int   `pulumi:"validStreamingUnits"`
}


func (val *TransformationResponse) Defaults() *TransformationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.StreamingUnits) {
		streamingUnits_ := 3
		tmp.StreamingUnits = &streamingUnits_
	}
	return &tmp
}

type TransformationResponseOutput struct{ *pulumi.OutputState }

func (TransformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransformationResponse)(nil)).Elem()
}

func (o TransformationResponseOutput) ToTransformationResponseOutput() TransformationResponseOutput {
	return o
}

func (o TransformationResponseOutput) ToTransformationResponseOutputWithContext(ctx context.Context) TransformationResponseOutput {
	return o
}

func (o TransformationResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v TransformationResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o TransformationResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v TransformationResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o TransformationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TransformationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TransformationResponseOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TransformationResponse) *string { return v.Query }).(pulumi.StringPtrOutput)
}

func (o TransformationResponseOutput) StreamingUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TransformationResponse) *int { return v.StreamingUnits }).(pulumi.IntPtrOutput)
}

func (o TransformationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v TransformationResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o TransformationResponseOutput) ValidStreamingUnits() pulumi.IntArrayOutput {
	return o.ApplyT(func(v TransformationResponse) []int { return v.ValidStreamingUnits }).(pulumi.IntArrayOutput)
}

type TransformationResponsePtrOutput struct{ *pulumi.OutputState }

func (TransformationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransformationResponse)(nil)).Elem()
}

func (o TransformationResponsePtrOutput) ToTransformationResponsePtrOutput() TransformationResponsePtrOutput {
	return o
}

func (o TransformationResponsePtrOutput) ToTransformationResponsePtrOutputWithContext(ctx context.Context) TransformationResponsePtrOutput {
	return o
}

func (o TransformationResponsePtrOutput) Elem() TransformationResponseOutput {
	return o.ApplyT(func(v *TransformationResponse) TransformationResponse {
		if v != nil {
			return *v
		}
		var ret TransformationResponse
		return ret
	}).(TransformationResponseOutput)
}

func (o TransformationResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o TransformationResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o TransformationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o TransformationResponsePtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Query
	}).(pulumi.StringPtrOutput)
}

func (o TransformationResponsePtrOutput) StreamingUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TransformationResponse) *int {
		if v == nil {
			return nil
		}
		return v.StreamingUnits
	}).(pulumi.IntPtrOutput)
}

func (o TransformationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o TransformationResponsePtrOutput) ValidStreamingUnits() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *TransformationResponse) []int {
		if v == nil {
			return nil
		}
		return v.ValidStreamingUnits
	}).(pulumi.IntArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterInfoOutput{})
	pulumi.RegisterOutputType(ClusterInfoPtrOutput{})
	pulumi.RegisterOutputType(ClusterInfoResponseOutput{})
	pulumi.RegisterOutputType(ClusterInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ClusterJobResponseOutput{})
	pulumi.RegisterOutputType(ClusterJobResponseArrayOutput{})
	pulumi.RegisterOutputType(ClusterSkuOutput{})
	pulumi.RegisterOutputType(ClusterSkuPtrOutput{})
	pulumi.RegisterOutputType(ClusterSkuResponseOutput{})
	pulumi.RegisterOutputType(ClusterSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(DiagnosticConditionResponseOutput{})
	pulumi.RegisterOutputType(DiagnosticConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(DiagnosticsResponseOutput{})
	pulumi.RegisterOutputType(FunctionTypeOutput{})
	pulumi.RegisterOutputType(FunctionTypeArrayOutput{})
	pulumi.RegisterOutputType(FunctionResponseOutput{})
	pulumi.RegisterOutputType(FunctionResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(InputTypeOutput{})
	pulumi.RegisterOutputType(InputTypeArrayOutput{})
	pulumi.RegisterOutputType(InputResponseOutput{})
	pulumi.RegisterOutputType(InputResponseArrayOutput{})
	pulumi.RegisterOutputType(JobStorageAccountOutput{})
	pulumi.RegisterOutputType(JobStorageAccountPtrOutput{})
	pulumi.RegisterOutputType(JobStorageAccountResponseOutput{})
	pulumi.RegisterOutputType(JobStorageAccountResponsePtrOutput{})
	pulumi.RegisterOutputType(OutputTypeOutput{})
	pulumi.RegisterOutputType(OutputTypeArrayOutput{})
	pulumi.RegisterOutputType(OutputResponseOutput{})
	pulumi.RegisterOutputType(OutputResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionArrayOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(TransformationOutput{})
	pulumi.RegisterOutputType(TransformationPtrOutput{})
	pulumi.RegisterOutputType(TransformationResponseOutput{})
	pulumi.RegisterOutputType(TransformationResponsePtrOutput{})
}
