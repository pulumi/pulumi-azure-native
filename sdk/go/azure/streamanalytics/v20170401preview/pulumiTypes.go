


package v20170401preview

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

type AzureMachineLearningServiceFunctionBinding struct {
	ApiKey                   *string                                   `pulumi:"apiKey"`
	BatchSize                *int                                      `pulumi:"batchSize"`
	Endpoint                 *string                                   `pulumi:"endpoint"`
	Inputs                   []AzureMachineLearningServiceInputColumn  `pulumi:"inputs"`
	NumberOfParallelRequests *int                                      `pulumi:"numberOfParallelRequests"`
	Outputs                  []AzureMachineLearningServiceOutputColumn `pulumi:"outputs"`
	Type                     string                                    `pulumi:"type"`
}

type AzureMachineLearningServiceFunctionBindingResponse struct {
	ApiKey                   *string                                           `pulumi:"apiKey"`
	BatchSize                *int                                              `pulumi:"batchSize"`
	Endpoint                 *string                                           `pulumi:"endpoint"`
	Inputs                   []AzureMachineLearningServiceInputColumnResponse  `pulumi:"inputs"`
	NumberOfParallelRequests *int                                              `pulumi:"numberOfParallelRequests"`
	Outputs                  []AzureMachineLearningServiceOutputColumnResponse `pulumi:"outputs"`
	Type                     string                                            `pulumi:"type"`
}

type AzureMachineLearningServiceInputColumn struct {
	DataType *string `pulumi:"dataType"`
	MapTo    *int    `pulumi:"mapTo"`
	Name     *string `pulumi:"name"`
}

type AzureMachineLearningServiceInputColumnResponse struct {
	DataType *string `pulumi:"dataType"`
	MapTo    *int    `pulumi:"mapTo"`
	Name     *string `pulumi:"name"`
}

type AzureMachineLearningServiceOutputColumn struct {
	DataType *string `pulumi:"dataType"`
	MapTo    *int    `pulumi:"mapTo"`
	Name     *string `pulumi:"name"`
}

type AzureMachineLearningServiceOutputColumnResponse struct {
	DataType *string `pulumi:"dataType"`
	MapTo    *int    `pulumi:"mapTo"`
	Name     *string `pulumi:"name"`
}

type AzureMachineLearningStudioFunctionBinding struct {
	ApiKey    *string                                  `pulumi:"apiKey"`
	BatchSize *int                                     `pulumi:"batchSize"`
	Endpoint  *string                                  `pulumi:"endpoint"`
	Inputs    *AzureMachineLearningStudioInputs        `pulumi:"inputs"`
	Outputs   []AzureMachineLearningStudioOutputColumn `pulumi:"outputs"`
	Type      string                                   `pulumi:"type"`
}

type AzureMachineLearningStudioFunctionBindingResponse struct {
	ApiKey    *string                                          `pulumi:"apiKey"`
	BatchSize *int                                             `pulumi:"batchSize"`
	Endpoint  *string                                          `pulumi:"endpoint"`
	Inputs    *AzureMachineLearningStudioInputsResponse        `pulumi:"inputs"`
	Outputs   []AzureMachineLearningStudioOutputColumnResponse `pulumi:"outputs"`
	Type      string                                           `pulumi:"type"`
}

type AzureMachineLearningStudioInputColumn struct {
	DataType *string `pulumi:"dataType"`
	MapTo    *int    `pulumi:"mapTo"`
	Name     *string `pulumi:"name"`
}

type AzureMachineLearningStudioInputColumnResponse struct {
	DataType *string `pulumi:"dataType"`
	MapTo    *int    `pulumi:"mapTo"`
	Name     *string `pulumi:"name"`
}

type AzureMachineLearningStudioInputs struct {
	ColumnNames []AzureMachineLearningStudioInputColumn `pulumi:"columnNames"`
	Name        *string                                 `pulumi:"name"`
}

type AzureMachineLearningStudioInputsResponse struct {
	ColumnNames []AzureMachineLearningStudioInputColumnResponse `pulumi:"columnNames"`
	Name        *string                                         `pulumi:"name"`
}

type AzureMachineLearningStudioOutputColumn struct {
	DataType *string `pulumi:"dataType"`
	Name     *string `pulumi:"name"`
}

type AzureMachineLearningStudioOutputColumnResponse struct {
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
	Properties *AzureSqlReferenceInputDataSourceProperties `pulumi:"properties"`
	Type       string                                      `pulumi:"type"`
}

type AzureSqlReferenceInputDataSourceProperties struct {
	Database           *string `pulumi:"database"`
	DeltaSnapshotQuery *string `pulumi:"deltaSnapshotQuery"`
	FullSnapshotQuery  *string `pulumi:"fullSnapshotQuery"`
	Password           *string `pulumi:"password"`
	RefreshRate        *string `pulumi:"refreshRate"`
	RefreshType        *string `pulumi:"refreshType"`
	Server             *string `pulumi:"server"`
	Table              *string `pulumi:"table"`
	User               *string `pulumi:"user"`
}

type AzureSqlReferenceInputDataSourcePropertiesResponse struct {
	Database           *string `pulumi:"database"`
	DeltaSnapshotQuery *string `pulumi:"deltaSnapshotQuery"`
	FullSnapshotQuery  *string `pulumi:"fullSnapshotQuery"`
	Password           *string `pulumi:"password"`
	RefreshRate        *string `pulumi:"refreshRate"`
	RefreshType        *string `pulumi:"refreshType"`
	Server             *string `pulumi:"server"`
	Table              *string `pulumi:"table"`
	User               *string `pulumi:"user"`
}

type AzureSqlReferenceInputDataSourceResponse struct {
	Properties *AzureSqlReferenceInputDataSourcePropertiesResponse `pulumi:"properties"`
	Type       string                                              `pulumi:"type"`
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
	Container          *string          `pulumi:"container"`
	DateFormat         *string          `pulumi:"dateFormat"`
	PathPattern        *string          `pulumi:"pathPattern"`
	StorageAccounts    []StorageAccount `pulumi:"storageAccounts"`
	TimeFormat         *string          `pulumi:"timeFormat"`
	Type               string           `pulumi:"type"`
}

type BlobOutputDataSourceResponse struct {
	AuthenticationMode *string                  `pulumi:"authenticationMode"`
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

type CSharpFunctionBinding struct {
	Class   *string `pulumi:"class"`
	DllPath *string `pulumi:"dllPath"`
	Method  *string `pulumi:"method"`
	Script  *string `pulumi:"script"`
	Type    string  `pulumi:"type"`
}

type CSharpFunctionBindingResponse struct {
	Class   *string `pulumi:"class"`
	DllPath *string `pulumi:"dllPath"`
	Method  *string `pulumi:"method"`
	Script  *string `pulumi:"script"`
	Type    string  `pulumi:"type"`
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

type CustomClrSerialization struct {
	SerializationClassName *string `pulumi:"serializationClassName"`
	SerializationDllPath   *string `pulumi:"serializationDllPath"`
	Type                   string  `pulumi:"type"`
}

type CustomClrSerializationResponse struct {
	SerializationClassName *string `pulumi:"serializationClassName"`
	SerializationDllPath   *string `pulumi:"serializationDllPath"`
	Type                   string  `pulumi:"type"`
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

type External struct {
	Container      *string         `pulumi:"container"`
	Path           *string         `pulumi:"path"`
	StorageAccount *StorageAccount `pulumi:"storageAccount"`
}





type ExternalInput interface {
	pulumi.Input

	ToExternalOutput() ExternalOutput
	ToExternalOutputWithContext(context.Context) ExternalOutput
}

type ExternalArgs struct {
	Container      pulumi.StringPtrInput  `pulumi:"container"`
	Path           pulumi.StringPtrInput  `pulumi:"path"`
	StorageAccount StorageAccountPtrInput `pulumi:"storageAccount"`
}

func (ExternalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*External)(nil)).Elem()
}

func (i ExternalArgs) ToExternalOutput() ExternalOutput {
	return i.ToExternalOutputWithContext(context.Background())
}

func (i ExternalArgs) ToExternalOutputWithContext(ctx context.Context) ExternalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalOutput)
}

func (i ExternalArgs) ToExternalPtrOutput() ExternalPtrOutput {
	return i.ToExternalPtrOutputWithContext(context.Background())
}

func (i ExternalArgs) ToExternalPtrOutputWithContext(ctx context.Context) ExternalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalOutput).ToExternalPtrOutputWithContext(ctx)
}









type ExternalPtrInput interface {
	pulumi.Input

	ToExternalPtrOutput() ExternalPtrOutput
	ToExternalPtrOutputWithContext(context.Context) ExternalPtrOutput
}

type externalPtrType ExternalArgs

func ExternalPtr(v *ExternalArgs) ExternalPtrInput {
	return (*externalPtrType)(v)
}

func (*externalPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**External)(nil)).Elem()
}

func (i *externalPtrType) ToExternalPtrOutput() ExternalPtrOutput {
	return i.ToExternalPtrOutputWithContext(context.Background())
}

func (i *externalPtrType) ToExternalPtrOutputWithContext(ctx context.Context) ExternalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalPtrOutput)
}

type ExternalOutput struct{ *pulumi.OutputState }

func (ExternalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*External)(nil)).Elem()
}

func (o ExternalOutput) ToExternalOutput() ExternalOutput {
	return o
}

func (o ExternalOutput) ToExternalOutputWithContext(ctx context.Context) ExternalOutput {
	return o
}

func (o ExternalOutput) ToExternalPtrOutput() ExternalPtrOutput {
	return o.ToExternalPtrOutputWithContext(context.Background())
}

func (o ExternalOutput) ToExternalPtrOutputWithContext(ctx context.Context) ExternalPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v External) *External {
		return &v
	}).(ExternalPtrOutput)
}

func (o ExternalOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v External) *string { return v.Container }).(pulumi.StringPtrOutput)
}

func (o ExternalOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v External) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ExternalOutput) StorageAccount() StorageAccountPtrOutput {
	return o.ApplyT(func(v External) *StorageAccount { return v.StorageAccount }).(StorageAccountPtrOutput)
}

type ExternalPtrOutput struct{ *pulumi.OutputState }

func (ExternalPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**External)(nil)).Elem()
}

func (o ExternalPtrOutput) ToExternalPtrOutput() ExternalPtrOutput {
	return o
}

func (o ExternalPtrOutput) ToExternalPtrOutputWithContext(ctx context.Context) ExternalPtrOutput {
	return o
}

func (o ExternalPtrOutput) Elem() ExternalOutput {
	return o.ApplyT(func(v *External) External {
		if v != nil {
			return *v
		}
		var ret External
		return ret
	}).(ExternalOutput)
}

func (o ExternalPtrOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *External) *string {
		if v == nil {
			return nil
		}
		return v.Container
	}).(pulumi.StringPtrOutput)
}

func (o ExternalPtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *External) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o ExternalPtrOutput) StorageAccount() StorageAccountPtrOutput {
	return o.ApplyT(func(v *External) *StorageAccount {
		if v == nil {
			return nil
		}
		return v.StorageAccount
	}).(StorageAccountPtrOutput)
}

type ExternalResponse struct {
	Container      *string                 `pulumi:"container"`
	Path           *string                 `pulumi:"path"`
	StorageAccount *StorageAccountResponse `pulumi:"storageAccount"`
}

type ExternalResponseOutput struct{ *pulumi.OutputState }

func (ExternalResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExternalResponse)(nil)).Elem()
}

func (o ExternalResponseOutput) ToExternalResponseOutput() ExternalResponseOutput {
	return o
}

func (o ExternalResponseOutput) ToExternalResponseOutputWithContext(ctx context.Context) ExternalResponseOutput {
	return o
}

func (o ExternalResponseOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExternalResponse) *string { return v.Container }).(pulumi.StringPtrOutput)
}

func (o ExternalResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExternalResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ExternalResponseOutput) StorageAccount() StorageAccountResponsePtrOutput {
	return o.ApplyT(func(v ExternalResponse) *StorageAccountResponse { return v.StorageAccount }).(StorageAccountResponsePtrOutput)
}

type ExternalResponsePtrOutput struct{ *pulumi.OutputState }

func (ExternalResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalResponse)(nil)).Elem()
}

func (o ExternalResponsePtrOutput) ToExternalResponsePtrOutput() ExternalResponsePtrOutput {
	return o
}

func (o ExternalResponsePtrOutput) ToExternalResponsePtrOutputWithContext(ctx context.Context) ExternalResponsePtrOutput {
	return o
}

func (o ExternalResponsePtrOutput) Elem() ExternalResponseOutput {
	return o.ApplyT(func(v *ExternalResponse) ExternalResponse {
		if v != nil {
			return *v
		}
		var ret ExternalResponse
		return ret
	}).(ExternalResponseOutput)
}

func (o ExternalResponsePtrOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExternalResponse) *string {
		if v == nil {
			return nil
		}
		return v.Container
	}).(pulumi.StringPtrOutput)
}

func (o ExternalResponsePtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExternalResponse) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o ExternalResponsePtrOutput) StorageAccount() StorageAccountResponsePtrOutput {
	return o.ApplyT(func(v *ExternalResponse) *StorageAccountResponse {
		if v == nil {
			return nil
		}
		return v.StorageAccount
	}).(StorageAccountResponsePtrOutput)
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

type Identity struct {
	PrincipalId *string `pulumi:"principalId"`
	TenantId    *string `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
	TenantId    pulumi.StringPtrInput `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
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

func (o IdentityOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Identity) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o IdentityOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Identity) *string { return v.TenantId }).(pulumi.StringPtrOutput)
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

func (o IdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identity) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identity) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
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
	PrincipalId *string `pulumi:"principalId"`
	TenantId    *string `pulumi:"tenantId"`
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

func (o IdentityResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
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
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
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
	SizeWindow    *int        `pulumi:"sizeWindow"`
	TimeWindow    *string     `pulumi:"timeWindow"`
}





type OutputTypeInput interface {
	pulumi.Input

	ToOutputTypeOutput() OutputTypeOutput
	ToOutputTypeOutputWithContext(context.Context) OutputTypeOutput
}

type OutputTypeArgs struct {
	Datasource    pulumi.Input          `pulumi:"datasource"`
	Name          pulumi.StringPtrInput `pulumi:"name"`
	Serialization pulumi.Input          `pulumi:"serialization"`
	SizeWindow    pulumi.IntPtrInput    `pulumi:"sizeWindow"`
	TimeWindow    pulumi.StringPtrInput `pulumi:"timeWindow"`
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

func (o OutputTypeOutput) SizeWindow() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OutputType) *int { return v.SizeWindow }).(pulumi.IntPtrOutput)
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
	SizeWindow    *int                `pulumi:"sizeWindow"`
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

func (o OutputResponseOutput) SizeWindow() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OutputResponse) *int { return v.SizeWindow }).(pulumi.IntPtrOutput)
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

type RawOutputDatasource struct {
	PayloadUri *string `pulumi:"payloadUri"`
	Type       string  `pulumi:"type"`
}

type RawOutputDatasourceResponse struct {
	PayloadUri *string `pulumi:"payloadUri"`
	Type       string  `pulumi:"type"`
}

type RawReferenceInputDataSource struct {
	Payload    *string `pulumi:"payload"`
	PayloadUri *string `pulumi:"payloadUri"`
	Type       string  `pulumi:"type"`
}

type RawReferenceInputDataSourceResponse struct {
	Payload    *string `pulumi:"payload"`
	PayloadUri *string `pulumi:"payloadUri"`
	Type       string  `pulumi:"type"`
}

type RawStreamInputDataSource struct {
	Payload    *string `pulumi:"payload"`
	PayloadUri *string `pulumi:"payloadUri"`
	Type       string  `pulumi:"type"`
}

type RawStreamInputDataSourceResponse struct {
	Payload    *string `pulumi:"payload"`
	PayloadUri *string `pulumi:"payloadUri"`
	Type       string  `pulumi:"type"`
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
	AuthenticationMode     *string           `pulumi:"authenticationMode"`
	PropertyColumns        []string          `pulumi:"propertyColumns"`
	QueueName              *string           `pulumi:"queueName"`
	ServiceBusNamespace    *string           `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string           `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string           `pulumi:"sharedAccessPolicyName"`
	SystemPropertyColumns  map[string]string `pulumi:"systemPropertyColumns"`
	Type                   string            `pulumi:"type"`
}

type ServiceBusQueueOutputDataSourceResponse struct {
	AuthenticationMode     *string           `pulumi:"authenticationMode"`
	PropertyColumns        []string          `pulumi:"propertyColumns"`
	QueueName              *string           `pulumi:"queueName"`
	ServiceBusNamespace    *string           `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string           `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string           `pulumi:"sharedAccessPolicyName"`
	SystemPropertyColumns  map[string]string `pulumi:"systemPropertyColumns"`
	Type                   string            `pulumi:"type"`
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

type StorageAccount struct {
	AccountKey  *string `pulumi:"accountKey"`
	AccountName *string `pulumi:"accountName"`
}





type StorageAccountInput interface {
	pulumi.Input

	ToStorageAccountOutput() StorageAccountOutput
	ToStorageAccountOutputWithContext(context.Context) StorageAccountOutput
}

type StorageAccountArgs struct {
	AccountKey  pulumi.StringPtrInput `pulumi:"accountKey"`
	AccountName pulumi.StringPtrInput `pulumi:"accountName"`
}

func (StorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (i StorageAccountArgs) ToStorageAccountOutput() StorageAccountOutput {
	return i.ToStorageAccountOutputWithContext(context.Background())
}

func (i StorageAccountArgs) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput)
}

func (i StorageAccountArgs) ToStorageAccountPtrOutput() StorageAccountPtrOutput {
	return i.ToStorageAccountPtrOutputWithContext(context.Background())
}

func (i StorageAccountArgs) ToStorageAccountPtrOutputWithContext(ctx context.Context) StorageAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput).ToStorageAccountPtrOutputWithContext(ctx)
}









type StorageAccountPtrInput interface {
	pulumi.Input

	ToStorageAccountPtrOutput() StorageAccountPtrOutput
	ToStorageAccountPtrOutputWithContext(context.Context) StorageAccountPtrOutput
}

type storageAccountPtrType StorageAccountArgs

func StorageAccountPtr(v *StorageAccountArgs) StorageAccountPtrInput {
	return (*storageAccountPtrType)(v)
}

func (*storageAccountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccount)(nil)).Elem()
}

func (i *storageAccountPtrType) ToStorageAccountPtrOutput() StorageAccountPtrOutput {
	return i.ToStorageAccountPtrOutputWithContext(context.Background())
}

func (i *storageAccountPtrType) ToStorageAccountPtrOutputWithContext(ctx context.Context) StorageAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPtrOutput)
}

type StorageAccountOutput struct{ *pulumi.OutputState }

func (StorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (o StorageAccountOutput) ToStorageAccountOutput() StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountPtrOutput() StorageAccountPtrOutput {
	return o.ToStorageAccountPtrOutputWithContext(context.Background())
}

func (o StorageAccountOutput) ToStorageAccountPtrOutputWithContext(ctx context.Context) StorageAccountPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccount) *StorageAccount {
		return &v
	}).(StorageAccountPtrOutput)
}

func (o StorageAccountOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccount) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccount) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

type StorageAccountPtrOutput struct{ *pulumi.OutputState }

func (StorageAccountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccount)(nil)).Elem()
}

func (o StorageAccountPtrOutput) ToStorageAccountPtrOutput() StorageAccountPtrOutput {
	return o
}

func (o StorageAccountPtrOutput) ToStorageAccountPtrOutputWithContext(ctx context.Context) StorageAccountPtrOutput {
	return o
}

func (o StorageAccountPtrOutput) Elem() StorageAccountOutput {
	return o.ApplyT(func(v *StorageAccount) StorageAccount {
		if v != nil {
			return *v
		}
		var ret StorageAccount
		return ret
	}).(StorageAccountOutput)
}

func (o StorageAccountPtrOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) *string {
		if v == nil {
			return nil
		}
		return v.AccountKey
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountPtrOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) *string {
		if v == nil {
			return nil
		}
		return v.AccountName
	}).(pulumi.StringPtrOutput)
}

type StorageAccountResponse struct {
	AccountKey  *string `pulumi:"accountKey"`
	AccountName *string `pulumi:"accountName"`
}

type StorageAccountResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutput() StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutputWithContext(ctx context.Context) StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o StorageAccountResponseOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

type StorageAccountResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageAccountResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponsePtrOutput) ToStorageAccountResponsePtrOutput() StorageAccountResponsePtrOutput {
	return o
}

func (o StorageAccountResponsePtrOutput) ToStorageAccountResponsePtrOutputWithContext(ctx context.Context) StorageAccountResponsePtrOutput {
	return o
}

func (o StorageAccountResponsePtrOutput) Elem() StorageAccountResponseOutput {
	return o.ApplyT(func(v *StorageAccountResponse) StorageAccountResponse {
		if v != nil {
			return *v
		}
		var ret StorageAccountResponse
		return ret
	}).(StorageAccountResponseOutput)
}

func (o StorageAccountResponsePtrOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccountKey
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountResponsePtrOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccountName
	}).(pulumi.StringPtrOutput)
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

type StreamingJobSku struct {
	Name *string `pulumi:"name"`
}





type StreamingJobSkuInput interface {
	pulumi.Input

	ToStreamingJobSkuOutput() StreamingJobSkuOutput
	ToStreamingJobSkuOutputWithContext(context.Context) StreamingJobSkuOutput
}

type StreamingJobSkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (StreamingJobSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingJobSku)(nil)).Elem()
}

func (i StreamingJobSkuArgs) ToStreamingJobSkuOutput() StreamingJobSkuOutput {
	return i.ToStreamingJobSkuOutputWithContext(context.Background())
}

func (i StreamingJobSkuArgs) ToStreamingJobSkuOutputWithContext(ctx context.Context) StreamingJobSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingJobSkuOutput)
}

func (i StreamingJobSkuArgs) ToStreamingJobSkuPtrOutput() StreamingJobSkuPtrOutput {
	return i.ToStreamingJobSkuPtrOutputWithContext(context.Background())
}

func (i StreamingJobSkuArgs) ToStreamingJobSkuPtrOutputWithContext(ctx context.Context) StreamingJobSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingJobSkuOutput).ToStreamingJobSkuPtrOutputWithContext(ctx)
}









type StreamingJobSkuPtrInput interface {
	pulumi.Input

	ToStreamingJobSkuPtrOutput() StreamingJobSkuPtrOutput
	ToStreamingJobSkuPtrOutputWithContext(context.Context) StreamingJobSkuPtrOutput
}

type streamingJobSkuPtrType StreamingJobSkuArgs

func StreamingJobSkuPtr(v *StreamingJobSkuArgs) StreamingJobSkuPtrInput {
	return (*streamingJobSkuPtrType)(v)
}

func (*streamingJobSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingJobSku)(nil)).Elem()
}

func (i *streamingJobSkuPtrType) ToStreamingJobSkuPtrOutput() StreamingJobSkuPtrOutput {
	return i.ToStreamingJobSkuPtrOutputWithContext(context.Background())
}

func (i *streamingJobSkuPtrType) ToStreamingJobSkuPtrOutputWithContext(ctx context.Context) StreamingJobSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingJobSkuPtrOutput)
}

type StreamingJobSkuOutput struct{ *pulumi.OutputState }

func (StreamingJobSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingJobSku)(nil)).Elem()
}

func (o StreamingJobSkuOutput) ToStreamingJobSkuOutput() StreamingJobSkuOutput {
	return o
}

func (o StreamingJobSkuOutput) ToStreamingJobSkuOutputWithContext(ctx context.Context) StreamingJobSkuOutput {
	return o
}

func (o StreamingJobSkuOutput) ToStreamingJobSkuPtrOutput() StreamingJobSkuPtrOutput {
	return o.ToStreamingJobSkuPtrOutputWithContext(context.Background())
}

func (o StreamingJobSkuOutput) ToStreamingJobSkuPtrOutputWithContext(ctx context.Context) StreamingJobSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamingJobSku) *StreamingJobSku {
		return &v
	}).(StreamingJobSkuPtrOutput)
}

func (o StreamingJobSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingJobSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type StreamingJobSkuPtrOutput struct{ *pulumi.OutputState }

func (StreamingJobSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingJobSku)(nil)).Elem()
}

func (o StreamingJobSkuPtrOutput) ToStreamingJobSkuPtrOutput() StreamingJobSkuPtrOutput {
	return o
}

func (o StreamingJobSkuPtrOutput) ToStreamingJobSkuPtrOutputWithContext(ctx context.Context) StreamingJobSkuPtrOutput {
	return o
}

func (o StreamingJobSkuPtrOutput) Elem() StreamingJobSkuOutput {
	return o.ApplyT(func(v *StreamingJobSku) StreamingJobSku {
		if v != nil {
			return *v
		}
		var ret StreamingJobSku
		return ret
	}).(StreamingJobSkuOutput)
}

func (o StreamingJobSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingJobSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type StreamingJobSkuResponse struct {
	Name *string `pulumi:"name"`
}

type StreamingJobSkuResponseOutput struct{ *pulumi.OutputState }

func (StreamingJobSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingJobSkuResponse)(nil)).Elem()
}

func (o StreamingJobSkuResponseOutput) ToStreamingJobSkuResponseOutput() StreamingJobSkuResponseOutput {
	return o
}

func (o StreamingJobSkuResponseOutput) ToStreamingJobSkuResponseOutputWithContext(ctx context.Context) StreamingJobSkuResponseOutput {
	return o
}

func (o StreamingJobSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingJobSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type StreamingJobSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (StreamingJobSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingJobSkuResponse)(nil)).Elem()
}

func (o StreamingJobSkuResponsePtrOutput) ToStreamingJobSkuResponsePtrOutput() StreamingJobSkuResponsePtrOutput {
	return o
}

func (o StreamingJobSkuResponsePtrOutput) ToStreamingJobSkuResponsePtrOutputWithContext(ctx context.Context) StreamingJobSkuResponsePtrOutput {
	return o
}

func (o StreamingJobSkuResponsePtrOutput) Elem() StreamingJobSkuResponseOutput {
	return o.ApplyT(func(v *StreamingJobSkuResponse) StreamingJobSkuResponse {
		if v != nil {
			return *v
		}
		var ret StreamingJobSkuResponse
		return ret
	}).(StreamingJobSkuResponseOutput)
}

func (o StreamingJobSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingJobSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type Transformation struct {
	Name           *string `pulumi:"name"`
	Query          *string `pulumi:"query"`
	StreamingUnits *int    `pulumi:"streamingUnits"`
}





type TransformationInput interface {
	pulumi.Input

	ToTransformationOutput() TransformationOutput
	ToTransformationOutputWithContext(context.Context) TransformationOutput
}

type TransformationArgs struct {
	Name           pulumi.StringPtrInput `pulumi:"name"`
	Query          pulumi.StringPtrInput `pulumi:"query"`
	StreamingUnits pulumi.IntPtrInput    `pulumi:"streamingUnits"`
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

type TransformationResponse struct {
	Etag           string  `pulumi:"etag"`
	Id             string  `pulumi:"id"`
	Name           *string `pulumi:"name"`
	Query          *string `pulumi:"query"`
	StreamingUnits *int    `pulumi:"streamingUnits"`
	Type           string  `pulumi:"type"`
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

func init() {
	pulumi.RegisterOutputType(ClusterInfoOutput{})
	pulumi.RegisterOutputType(ClusterInfoPtrOutput{})
	pulumi.RegisterOutputType(ClusterInfoResponseOutput{})
	pulumi.RegisterOutputType(ClusterInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(DiagnosticConditionResponseOutput{})
	pulumi.RegisterOutputType(DiagnosticConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(DiagnosticsResponseOutput{})
	pulumi.RegisterOutputType(ExternalOutput{})
	pulumi.RegisterOutputType(ExternalPtrOutput{})
	pulumi.RegisterOutputType(ExternalResponseOutput{})
	pulumi.RegisterOutputType(ExternalResponsePtrOutput{})
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
	pulumi.RegisterOutputType(StorageAccountOutput{})
	pulumi.RegisterOutputType(StorageAccountPtrOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountResponsePtrOutput{})
	pulumi.RegisterOutputType(StreamingJobSkuOutput{})
	pulumi.RegisterOutputType(StreamingJobSkuPtrOutput{})
	pulumi.RegisterOutputType(StreamingJobSkuResponseOutput{})
	pulumi.RegisterOutputType(StreamingJobSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(TransformationOutput{})
	pulumi.RegisterOutputType(TransformationPtrOutput{})
	pulumi.RegisterOutputType(TransformationResponseOutput{})
	pulumi.RegisterOutputType(TransformationResponsePtrOutput{})
}
