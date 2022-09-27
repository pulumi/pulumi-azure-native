


package v20180601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlMISource struct {
	AdditionalColumns            interface{}                         `pulumi:"additionalColumns"`
	DisableMetricsCollection     interface{}                         `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections     interface{}                         `pulumi:"maxConcurrentConnections"`
	PartitionOption              interface{}                         `pulumi:"partitionOption"`
	PartitionSettings            *SqlPartitionSettings               `pulumi:"partitionSettings"`
	ProduceAdditionalTypes       interface{}                         `pulumi:"produceAdditionalTypes"`
	QueryTimeout                 interface{}                         `pulumi:"queryTimeout"`
	SourceRetryCount             interface{}                         `pulumi:"sourceRetryCount"`
	SourceRetryWait              interface{}                         `pulumi:"sourceRetryWait"`
	SqlReaderQuery               interface{}                         `pulumi:"sqlReaderQuery"`
	SqlReaderStoredProcedureName interface{}                         `pulumi:"sqlReaderStoredProcedureName"`
	StoredProcedureParameters    map[string]StoredProcedureParameter `pulumi:"storedProcedureParameters"`
	Type                         string                              `pulumi:"type"`
}

type SqlMISourceResponse struct {
	AdditionalColumns            interface{}                                 `pulumi:"additionalColumns"`
	DisableMetricsCollection     interface{}                                 `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections     interface{}                                 `pulumi:"maxConcurrentConnections"`
	PartitionOption              interface{}                                 `pulumi:"partitionOption"`
	PartitionSettings            *SqlPartitionSettingsResponse               `pulumi:"partitionSettings"`
	ProduceAdditionalTypes       interface{}                                 `pulumi:"produceAdditionalTypes"`
	QueryTimeout                 interface{}                                 `pulumi:"queryTimeout"`
	SourceRetryCount             interface{}                                 `pulumi:"sourceRetryCount"`
	SourceRetryWait              interface{}                                 `pulumi:"sourceRetryWait"`
	SqlReaderQuery               interface{}                                 `pulumi:"sqlReaderQuery"`
	SqlReaderStoredProcedureName interface{}                                 `pulumi:"sqlReaderStoredProcedureName"`
	StoredProcedureParameters    map[string]StoredProcedureParameterResponse `pulumi:"storedProcedureParameters"`
	Type                         string                                      `pulumi:"type"`
}

type SqlPartitionSettings struct {
	PartitionColumnName interface{} `pulumi:"partitionColumnName"`
	PartitionLowerBound interface{} `pulumi:"partitionLowerBound"`
	PartitionUpperBound interface{} `pulumi:"partitionUpperBound"`
}

type SqlPartitionSettingsResponse struct {
	PartitionColumnName interface{} `pulumi:"partitionColumnName"`
	PartitionLowerBound interface{} `pulumi:"partitionLowerBound"`
	PartitionUpperBound interface{} `pulumi:"partitionUpperBound"`
}

type SqlServerLinkedService struct {
	AlwaysEncryptedSettings *SqlAlwaysEncryptedProperties     `pulumi:"alwaysEncryptedSettings"`
	Annotations             []interface{}                     `pulumi:"annotations"`
	ConnectVia              *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString        interface{}                       `pulumi:"connectionString"`
	Description             *string                           `pulumi:"description"`
	EncryptedCredential     interface{}                       `pulumi:"encryptedCredential"`
	Parameters              map[string]ParameterSpecification `pulumi:"parameters"`
	Password                interface{}                       `pulumi:"password"`
	Type                    string                            `pulumi:"type"`
	UserName                interface{}                       `pulumi:"userName"`
}

type SqlServerLinkedServiceResponse struct {
	AlwaysEncryptedSettings *SqlAlwaysEncryptedPropertiesResponse     `pulumi:"alwaysEncryptedSettings"`
	Annotations             []interface{}                             `pulumi:"annotations"`
	ConnectVia              *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString        interface{}                               `pulumi:"connectionString"`
	Description             *string                                   `pulumi:"description"`
	EncryptedCredential     interface{}                               `pulumi:"encryptedCredential"`
	Parameters              map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                interface{}                               `pulumi:"password"`
	Type                    string                                    `pulumi:"type"`
	UserName                interface{}                               `pulumi:"userName"`
}

type SqlServerSink struct {
	DisableMetricsCollection              interface{}                         `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections              interface{}                         `pulumi:"maxConcurrentConnections"`
	PreCopyScript                         interface{}                         `pulumi:"preCopyScript"`
	SinkRetryCount                        interface{}                         `pulumi:"sinkRetryCount"`
	SinkRetryWait                         interface{}                         `pulumi:"sinkRetryWait"`
	SqlWriterStoredProcedureName          interface{}                         `pulumi:"sqlWriterStoredProcedureName"`
	SqlWriterTableType                    interface{}                         `pulumi:"sqlWriterTableType"`
	SqlWriterUseTableLock                 interface{}                         `pulumi:"sqlWriterUseTableLock"`
	StoredProcedureParameters             map[string]StoredProcedureParameter `pulumi:"storedProcedureParameters"`
	StoredProcedureTableTypeParameterName interface{}                         `pulumi:"storedProcedureTableTypeParameterName"`
	TableOption                           interface{}                         `pulumi:"tableOption"`
	Type                                  string                              `pulumi:"type"`
	UpsertSettings                        *SqlUpsertSettings                  `pulumi:"upsertSettings"`
	WriteBatchSize                        interface{}                         `pulumi:"writeBatchSize"`
	WriteBatchTimeout                     interface{}                         `pulumi:"writeBatchTimeout"`
	WriteBehavior                         interface{}                         `pulumi:"writeBehavior"`
}

type SqlServerSinkResponse struct {
	DisableMetricsCollection              interface{}                                 `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections              interface{}                                 `pulumi:"maxConcurrentConnections"`
	PreCopyScript                         interface{}                                 `pulumi:"preCopyScript"`
	SinkRetryCount                        interface{}                                 `pulumi:"sinkRetryCount"`
	SinkRetryWait                         interface{}                                 `pulumi:"sinkRetryWait"`
	SqlWriterStoredProcedureName          interface{}                                 `pulumi:"sqlWriterStoredProcedureName"`
	SqlWriterTableType                    interface{}                                 `pulumi:"sqlWriterTableType"`
	SqlWriterUseTableLock                 interface{}                                 `pulumi:"sqlWriterUseTableLock"`
	StoredProcedureParameters             map[string]StoredProcedureParameterResponse `pulumi:"storedProcedureParameters"`
	StoredProcedureTableTypeParameterName interface{}                                 `pulumi:"storedProcedureTableTypeParameterName"`
	TableOption                           interface{}                                 `pulumi:"tableOption"`
	Type                                  string                                      `pulumi:"type"`
	UpsertSettings                        *SqlUpsertSettingsResponse                  `pulumi:"upsertSettings"`
	WriteBatchSize                        interface{}                                 `pulumi:"writeBatchSize"`
	WriteBatchTimeout                     interface{}                                 `pulumi:"writeBatchTimeout"`
	WriteBehavior                         interface{}                                 `pulumi:"writeBehavior"`
}

type SqlServerSource struct {
	AdditionalColumns            interface{}                         `pulumi:"additionalColumns"`
	DisableMetricsCollection     interface{}                         `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections     interface{}                         `pulumi:"maxConcurrentConnections"`
	PartitionOption              interface{}                         `pulumi:"partitionOption"`
	PartitionSettings            *SqlPartitionSettings               `pulumi:"partitionSettings"`
	ProduceAdditionalTypes       interface{}                         `pulumi:"produceAdditionalTypes"`
	QueryTimeout                 interface{}                         `pulumi:"queryTimeout"`
	SourceRetryCount             interface{}                         `pulumi:"sourceRetryCount"`
	SourceRetryWait              interface{}                         `pulumi:"sourceRetryWait"`
	SqlReaderQuery               interface{}                         `pulumi:"sqlReaderQuery"`
	SqlReaderStoredProcedureName interface{}                         `pulumi:"sqlReaderStoredProcedureName"`
	StoredProcedureParameters    map[string]StoredProcedureParameter `pulumi:"storedProcedureParameters"`
	Type                         string                              `pulumi:"type"`
}

type SqlServerSourceResponse struct {
	AdditionalColumns            interface{}                                 `pulumi:"additionalColumns"`
	DisableMetricsCollection     interface{}                                 `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections     interface{}                                 `pulumi:"maxConcurrentConnections"`
	PartitionOption              interface{}                                 `pulumi:"partitionOption"`
	PartitionSettings            *SqlPartitionSettingsResponse               `pulumi:"partitionSettings"`
	ProduceAdditionalTypes       interface{}                                 `pulumi:"produceAdditionalTypes"`
	QueryTimeout                 interface{}                                 `pulumi:"queryTimeout"`
	SourceRetryCount             interface{}                                 `pulumi:"sourceRetryCount"`
	SourceRetryWait              interface{}                                 `pulumi:"sourceRetryWait"`
	SqlReaderQuery               interface{}                                 `pulumi:"sqlReaderQuery"`
	SqlReaderStoredProcedureName interface{}                                 `pulumi:"sqlReaderStoredProcedureName"`
	StoredProcedureParameters    map[string]StoredProcedureParameterResponse `pulumi:"storedProcedureParameters"`
	Type                         string                                      `pulumi:"type"`
}

type SqlServerStoredProcedureActivity struct {
	DependsOn                 []ActivityDependency   `pulumi:"dependsOn"`
	Description               *string                `pulumi:"description"`
	LinkedServiceName         LinkedServiceReference `pulumi:"linkedServiceName"`
	Name                      string                 `pulumi:"name"`
	Policy                    *ActivityPolicy        `pulumi:"policy"`
	StoredProcedureName       interface{}            `pulumi:"storedProcedureName"`
	StoredProcedureParameters interface{}            `pulumi:"storedProcedureParameters"`
	Type                      string                 `pulumi:"type"`
	UserProperties            []UserProperty         `pulumi:"userProperties"`
}

type SqlServerStoredProcedureActivityResponse struct {
	DependsOn                 []ActivityDependencyResponse   `pulumi:"dependsOn"`
	Description               *string                        `pulumi:"description"`
	LinkedServiceName         LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	Name                      string                         `pulumi:"name"`
	Policy                    *ActivityPolicyResponse        `pulumi:"policy"`
	StoredProcedureName       interface{}                    `pulumi:"storedProcedureName"`
	StoredProcedureParameters interface{}                    `pulumi:"storedProcedureParameters"`
	Type                      string                         `pulumi:"type"`
	UserProperties            []UserPropertyResponse         `pulumi:"userProperties"`
}

type SqlServerTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Table             interface{}                       `pulumi:"table"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type SqlServerTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Table             interface{}                               `pulumi:"table"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type SqlSink struct {
	DisableMetricsCollection              interface{}                         `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections              interface{}                         `pulumi:"maxConcurrentConnections"`
	PreCopyScript                         interface{}                         `pulumi:"preCopyScript"`
	SinkRetryCount                        interface{}                         `pulumi:"sinkRetryCount"`
	SinkRetryWait                         interface{}                         `pulumi:"sinkRetryWait"`
	SqlWriterStoredProcedureName          interface{}                         `pulumi:"sqlWriterStoredProcedureName"`
	SqlWriterTableType                    interface{}                         `pulumi:"sqlWriterTableType"`
	SqlWriterUseTableLock                 interface{}                         `pulumi:"sqlWriterUseTableLock"`
	StoredProcedureParameters             map[string]StoredProcedureParameter `pulumi:"storedProcedureParameters"`
	StoredProcedureTableTypeParameterName interface{}                         `pulumi:"storedProcedureTableTypeParameterName"`
	TableOption                           interface{}                         `pulumi:"tableOption"`
	Type                                  string                              `pulumi:"type"`
	UpsertSettings                        *SqlUpsertSettings                  `pulumi:"upsertSettings"`
	WriteBatchSize                        interface{}                         `pulumi:"writeBatchSize"`
	WriteBatchTimeout                     interface{}                         `pulumi:"writeBatchTimeout"`
	WriteBehavior                         interface{}                         `pulumi:"writeBehavior"`
}

type SqlSinkResponse struct {
	DisableMetricsCollection              interface{}                                 `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections              interface{}                                 `pulumi:"maxConcurrentConnections"`
	PreCopyScript                         interface{}                                 `pulumi:"preCopyScript"`
	SinkRetryCount                        interface{}                                 `pulumi:"sinkRetryCount"`
	SinkRetryWait                         interface{}                                 `pulumi:"sinkRetryWait"`
	SqlWriterStoredProcedureName          interface{}                                 `pulumi:"sqlWriterStoredProcedureName"`
	SqlWriterTableType                    interface{}                                 `pulumi:"sqlWriterTableType"`
	SqlWriterUseTableLock                 interface{}                                 `pulumi:"sqlWriterUseTableLock"`
	StoredProcedureParameters             map[string]StoredProcedureParameterResponse `pulumi:"storedProcedureParameters"`
	StoredProcedureTableTypeParameterName interface{}                                 `pulumi:"storedProcedureTableTypeParameterName"`
	TableOption                           interface{}                                 `pulumi:"tableOption"`
	Type                                  string                                      `pulumi:"type"`
	UpsertSettings                        *SqlUpsertSettingsResponse                  `pulumi:"upsertSettings"`
	WriteBatchSize                        interface{}                                 `pulumi:"writeBatchSize"`
	WriteBatchTimeout                     interface{}                                 `pulumi:"writeBatchTimeout"`
	WriteBehavior                         interface{}                                 `pulumi:"writeBehavior"`
}

type SqlSource struct {
	AdditionalColumns            interface{}                         `pulumi:"additionalColumns"`
	DisableMetricsCollection     interface{}                         `pulumi:"disableMetricsCollection"`
	IsolationLevel               interface{}                         `pulumi:"isolationLevel"`
	MaxConcurrentConnections     interface{}                         `pulumi:"maxConcurrentConnections"`
	PartitionOption              interface{}                         `pulumi:"partitionOption"`
	PartitionSettings            *SqlPartitionSettings               `pulumi:"partitionSettings"`
	QueryTimeout                 interface{}                         `pulumi:"queryTimeout"`
	SourceRetryCount             interface{}                         `pulumi:"sourceRetryCount"`
	SourceRetryWait              interface{}                         `pulumi:"sourceRetryWait"`
	SqlReaderQuery               interface{}                         `pulumi:"sqlReaderQuery"`
	SqlReaderStoredProcedureName interface{}                         `pulumi:"sqlReaderStoredProcedureName"`
	StoredProcedureParameters    map[string]StoredProcedureParameter `pulumi:"storedProcedureParameters"`
	Type                         string                              `pulumi:"type"`
}

type SqlSourceResponse struct {
	AdditionalColumns            interface{}                                 `pulumi:"additionalColumns"`
	DisableMetricsCollection     interface{}                                 `pulumi:"disableMetricsCollection"`
	IsolationLevel               interface{}                                 `pulumi:"isolationLevel"`
	MaxConcurrentConnections     interface{}                                 `pulumi:"maxConcurrentConnections"`
	PartitionOption              interface{}                                 `pulumi:"partitionOption"`
	PartitionSettings            *SqlPartitionSettingsResponse               `pulumi:"partitionSettings"`
	QueryTimeout                 interface{}                                 `pulumi:"queryTimeout"`
	SourceRetryCount             interface{}                                 `pulumi:"sourceRetryCount"`
	SourceRetryWait              interface{}                                 `pulumi:"sourceRetryWait"`
	SqlReaderQuery               interface{}                                 `pulumi:"sqlReaderQuery"`
	SqlReaderStoredProcedureName interface{}                                 `pulumi:"sqlReaderStoredProcedureName"`
	StoredProcedureParameters    map[string]StoredProcedureParameterResponse `pulumi:"storedProcedureParameters"`
	Type                         string                                      `pulumi:"type"`
}

type SqlUpsertSettings struct {
	InterimSchemaName interface{} `pulumi:"interimSchemaName"`
	Keys              interface{} `pulumi:"keys"`
	UseTempDB         interface{} `pulumi:"useTempDB"`
}

type SqlUpsertSettingsResponse struct {
	InterimSchemaName interface{} `pulumi:"interimSchemaName"`
	Keys              interface{} `pulumi:"keys"`
	UseTempDB         interface{} `pulumi:"useTempDB"`
}

type SquareLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	ClientId              interface{}                       `pulumi:"clientId"`
	ClientSecret          interface{}                       `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionProperties  interface{}                       `pulumi:"connectionProperties"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Host                  interface{}                       `pulumi:"host"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	RedirectUri           interface{}                       `pulumi:"redirectUri"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
}

type SquareLinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	ClientId              interface{}                               `pulumi:"clientId"`
	ClientSecret          interface{}                               `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionProperties  interface{}                               `pulumi:"connectionProperties"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Host                  interface{}                               `pulumi:"host"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	RedirectUri           interface{}                               `pulumi:"redirectUri"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
}

type SquareObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type SquareObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type SquareSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SquareSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SsisEnvironmentReferenceResponse struct {
	EnvironmentFolderName *string  `pulumi:"environmentFolderName"`
	EnvironmentName       *string  `pulumi:"environmentName"`
	Id                    *float64 `pulumi:"id"`
	ReferenceType         *string  `pulumi:"referenceType"`
}

type SsisEnvironmentResponse struct {
	Description *string                `pulumi:"description"`
	FolderId    *float64               `pulumi:"folderId"`
	Id          *float64               `pulumi:"id"`
	Name        *string                `pulumi:"name"`
	Type        string                 `pulumi:"type"`
	Variables   []SsisVariableResponse `pulumi:"variables"`
}

type SsisFolderResponse struct {
	Description *string  `pulumi:"description"`
	Id          *float64 `pulumi:"id"`
	Name        *string  `pulumi:"name"`
	Type        string   `pulumi:"type"`
}

type SsisPackageResponse struct {
	Description    *string                 `pulumi:"description"`
	FolderId       *float64                `pulumi:"folderId"`
	Id             *float64                `pulumi:"id"`
	Name           *string                 `pulumi:"name"`
	Parameters     []SsisParameterResponse `pulumi:"parameters"`
	ProjectId      *float64                `pulumi:"projectId"`
	ProjectVersion *float64                `pulumi:"projectVersion"`
	Type           string                  `pulumi:"type"`
}

type SsisParameterResponse struct {
	DataType              *string  `pulumi:"dataType"`
	DefaultValue          *string  `pulumi:"defaultValue"`
	Description           *string  `pulumi:"description"`
	DesignDefaultValue    *string  `pulumi:"designDefaultValue"`
	Id                    *float64 `pulumi:"id"`
	Name                  *string  `pulumi:"name"`
	Required              *bool    `pulumi:"required"`
	Sensitive             *bool    `pulumi:"sensitive"`
	SensitiveDefaultValue *string  `pulumi:"sensitiveDefaultValue"`
	ValueSet              *bool    `pulumi:"valueSet"`
	ValueType             *string  `pulumi:"valueType"`
	Variable              *string  `pulumi:"variable"`
}

type SsisProjectResponse struct {
	Description     *string                            `pulumi:"description"`
	EnvironmentRefs []SsisEnvironmentReferenceResponse `pulumi:"environmentRefs"`
	FolderId        *float64                           `pulumi:"folderId"`
	Id              *float64                           `pulumi:"id"`
	Name            *string                            `pulumi:"name"`
	Parameters      []SsisParameterResponse            `pulumi:"parameters"`
	Type            string                             `pulumi:"type"`
	Version         *float64                           `pulumi:"version"`
}

type SsisVariableResponse struct {
	DataType       *string  `pulumi:"dataType"`
	Description    *string  `pulumi:"description"`
	Id             *float64 `pulumi:"id"`
	Name           *string  `pulumi:"name"`
	Sensitive      *bool    `pulumi:"sensitive"`
	SensitiveValue *string  `pulumi:"sensitiveValue"`
	Value          *string  `pulumi:"value"`
}

type StagingSettings struct {
	EnableCompression interface{}            `pulumi:"enableCompression"`
	LinkedServiceName LinkedServiceReference `pulumi:"linkedServiceName"`
	Path              interface{}            `pulumi:"path"`
}

type StagingSettingsResponse struct {
	EnableCompression interface{}                    `pulumi:"enableCompression"`
	LinkedServiceName LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	Path              interface{}                    `pulumi:"path"`
}

type StoredProcedureParameter struct {
	Type  *string     `pulumi:"type"`
	Value interface{} `pulumi:"value"`
}

type StoredProcedureParameterResponse struct {
	Type  *string     `pulumi:"type"`
	Value interface{} `pulumi:"value"`
}

type SwitchActivity struct {
	Cases             []SwitchCase         `pulumi:"cases"`
	DefaultActivities []interface{}        `pulumi:"defaultActivities"`
	DependsOn         []ActivityDependency `pulumi:"dependsOn"`
	Description       *string              `pulumi:"description"`
	Name              string               `pulumi:"name"`
	On                Expression           `pulumi:"on"`
	Type              string               `pulumi:"type"`
	UserProperties    []UserProperty       `pulumi:"userProperties"`
}

type SwitchActivityResponse struct {
	Cases             []SwitchCaseResponse         `pulumi:"cases"`
	DefaultActivities []interface{}                `pulumi:"defaultActivities"`
	DependsOn         []ActivityDependencyResponse `pulumi:"dependsOn"`
	Description       *string                      `pulumi:"description"`
	Name              string                       `pulumi:"name"`
	On                ExpressionResponse           `pulumi:"on"`
	Type              string                       `pulumi:"type"`
	UserProperties    []UserPropertyResponse       `pulumi:"userProperties"`
}

type SwitchCase struct {
	Activities []interface{} `pulumi:"activities"`
	Value      *string       `pulumi:"value"`
}

type SwitchCaseResponse struct {
	Activities []interface{} `pulumi:"activities"`
	Value      *string       `pulumi:"value"`
}

type SybaseLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	AuthenticationType  *string                           `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Database            interface{}                       `pulumi:"database"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Schema              interface{}                       `pulumi:"schema"`
	Server              interface{}                       `pulumi:"server"`
	Type                string                            `pulumi:"type"`
	Username            interface{}                       `pulumi:"username"`
}

type SybaseLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	AuthenticationType  *string                                   `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Database            interface{}                               `pulumi:"database"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Schema              interface{}                               `pulumi:"schema"`
	Server              interface{}                               `pulumi:"server"`
	Type                string                                    `pulumi:"type"`
	Username            interface{}                               `pulumi:"username"`
}

type SybaseSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SybaseSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SybaseTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type SybaseTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type SynapseNotebookActivity struct {
	Conf              interface{}                          `pulumi:"conf"`
	DependsOn         []ActivityDependency                 `pulumi:"dependsOn"`
	Description       *string                              `pulumi:"description"`
	DriverSize        interface{}                          `pulumi:"driverSize"`
	ExecutorSize      interface{}                          `pulumi:"executorSize"`
	LinkedServiceName *LinkedServiceReference              `pulumi:"linkedServiceName"`
	Name              string                               `pulumi:"name"`
	Notebook          SynapseNotebookReference             `pulumi:"notebook"`
	NumExecutors      *int                                 `pulumi:"numExecutors"`
	Parameters        map[string]NotebookParameter         `pulumi:"parameters"`
	Policy            *ActivityPolicy                      `pulumi:"policy"`
	SparkPool         *BigDataPoolParametrizationReference `pulumi:"sparkPool"`
	Type              string                               `pulumi:"type"`
	UserProperties    []UserProperty                       `pulumi:"userProperties"`
}

type SynapseNotebookActivityResponse struct {
	Conf              interface{}                                  `pulumi:"conf"`
	DependsOn         []ActivityDependencyResponse                 `pulumi:"dependsOn"`
	Description       *string                                      `pulumi:"description"`
	DriverSize        interface{}                                  `pulumi:"driverSize"`
	ExecutorSize      interface{}                                  `pulumi:"executorSize"`
	LinkedServiceName *LinkedServiceReferenceResponse              `pulumi:"linkedServiceName"`
	Name              string                                       `pulumi:"name"`
	Notebook          SynapseNotebookReferenceResponse             `pulumi:"notebook"`
	NumExecutors      *int                                         `pulumi:"numExecutors"`
	Parameters        map[string]NotebookParameterResponse         `pulumi:"parameters"`
	Policy            *ActivityPolicyResponse                      `pulumi:"policy"`
	SparkPool         *BigDataPoolParametrizationReferenceResponse `pulumi:"sparkPool"`
	Type              string                                       `pulumi:"type"`
	UserProperties    []UserPropertyResponse                       `pulumi:"userProperties"`
}

type SynapseNotebookReference struct {
	ReferenceName interface{} `pulumi:"referenceName"`
	Type          string      `pulumi:"type"`
}

type SynapseNotebookReferenceResponse struct {
	ReferenceName interface{} `pulumi:"referenceName"`
	Type          string      `pulumi:"type"`
}

type SynapseSparkJobDefinitionActivity struct {
	Arguments         []interface{}                        `pulumi:"arguments"`
	ClassName         interface{}                          `pulumi:"className"`
	Conf              interface{}                          `pulumi:"conf"`
	DependsOn         []ActivityDependency                 `pulumi:"dependsOn"`
	Description       *string                              `pulumi:"description"`
	DriverSize        interface{}                          `pulumi:"driverSize"`
	ExecutorSize      interface{}                          `pulumi:"executorSize"`
	File              interface{}                          `pulumi:"file"`
	Files             []interface{}                        `pulumi:"files"`
	LinkedServiceName *LinkedServiceReference              `pulumi:"linkedServiceName"`
	Name              string                               `pulumi:"name"`
	NumExecutors      *int                                 `pulumi:"numExecutors"`
	Policy            *ActivityPolicy                      `pulumi:"policy"`
	SparkJob          SynapseSparkJobReference             `pulumi:"sparkJob"`
	TargetBigDataPool *BigDataPoolParametrizationReference `pulumi:"targetBigDataPool"`
	Type              string                               `pulumi:"type"`
	UserProperties    []UserProperty                       `pulumi:"userProperties"`
}

type SynapseSparkJobDefinitionActivityResponse struct {
	Arguments         []interface{}                                `pulumi:"arguments"`
	ClassName         interface{}                                  `pulumi:"className"`
	Conf              interface{}                                  `pulumi:"conf"`
	DependsOn         []ActivityDependencyResponse                 `pulumi:"dependsOn"`
	Description       *string                                      `pulumi:"description"`
	DriverSize        interface{}                                  `pulumi:"driverSize"`
	ExecutorSize      interface{}                                  `pulumi:"executorSize"`
	File              interface{}                                  `pulumi:"file"`
	Files             []interface{}                                `pulumi:"files"`
	LinkedServiceName *LinkedServiceReferenceResponse              `pulumi:"linkedServiceName"`
	Name              string                                       `pulumi:"name"`
	NumExecutors      *int                                         `pulumi:"numExecutors"`
	Policy            *ActivityPolicyResponse                      `pulumi:"policy"`
	SparkJob          SynapseSparkJobReferenceResponse             `pulumi:"sparkJob"`
	TargetBigDataPool *BigDataPoolParametrizationReferenceResponse `pulumi:"targetBigDataPool"`
	Type              string                                       `pulumi:"type"`
	UserProperties    []UserPropertyResponse                       `pulumi:"userProperties"`
}

type SynapseSparkJobReference struct {
	ReferenceName string `pulumi:"referenceName"`
	Type          string `pulumi:"type"`
}

type SynapseSparkJobReferenceResponse struct {
	ReferenceName string `pulumi:"referenceName"`
	Type          string `pulumi:"type"`
}

type TabularSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type TabularSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type TarGZipReadSettings struct {
	PreserveCompressionFileNameAsFolder interface{} `pulumi:"preserveCompressionFileNameAsFolder"`
	Type                                string      `pulumi:"type"`
}

type TarGZipReadSettingsResponse struct {
	PreserveCompressionFileNameAsFolder interface{} `pulumi:"preserveCompressionFileNameAsFolder"`
	Type                                string      `pulumi:"type"`
}

type TarReadSettings struct {
	PreserveCompressionFileNameAsFolder interface{} `pulumi:"preserveCompressionFileNameAsFolder"`
	Type                                string      `pulumi:"type"`
}

type TarReadSettingsResponse struct {
	PreserveCompressionFileNameAsFolder interface{} `pulumi:"preserveCompressionFileNameAsFolder"`
	Type                                string      `pulumi:"type"`
}

type TeamDeskLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ApiToken            interface{}                       `pulumi:"apiToken"`
	AuthenticationType  string                            `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Type                string                            `pulumi:"type"`
	Url                 interface{}                       `pulumi:"url"`
	UserName            interface{}                       `pulumi:"userName"`
}

type TeamDeskLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ApiToken            interface{}                               `pulumi:"apiToken"`
	AuthenticationType  string                                    `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
	Url                 interface{}                               `pulumi:"url"`
	UserName            interface{}                               `pulumi:"userName"`
}

type TeradataLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	AuthenticationType  *string                           `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Server              interface{}                       `pulumi:"server"`
	Type                string                            `pulumi:"type"`
	Username            interface{}                       `pulumi:"username"`
}

type TeradataLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	AuthenticationType  *string                                   `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Server              interface{}                               `pulumi:"server"`
	Type                string                                    `pulumi:"type"`
	Username            interface{}                               `pulumi:"username"`
}

type TeradataPartitionSettings struct {
	PartitionColumnName interface{} `pulumi:"partitionColumnName"`
	PartitionLowerBound interface{} `pulumi:"partitionLowerBound"`
	PartitionUpperBound interface{} `pulumi:"partitionUpperBound"`
}

type TeradataPartitionSettingsResponse struct {
	PartitionColumnName interface{} `pulumi:"partitionColumnName"`
	PartitionLowerBound interface{} `pulumi:"partitionLowerBound"`
	PartitionUpperBound interface{} `pulumi:"partitionUpperBound"`
}

type TeradataSource struct {
	AdditionalColumns        interface{}                `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}                `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}                `pulumi:"maxConcurrentConnections"`
	PartitionOption          interface{}                `pulumi:"partitionOption"`
	PartitionSettings        *TeradataPartitionSettings `pulumi:"partitionSettings"`
	Query                    interface{}                `pulumi:"query"`
	QueryTimeout             interface{}                `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}                `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                `pulumi:"sourceRetryWait"`
	Type                     string                     `pulumi:"type"`
}

type TeradataSourceResponse struct {
	AdditionalColumns        interface{}                        `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}                        `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}                        `pulumi:"maxConcurrentConnections"`
	PartitionOption          interface{}                        `pulumi:"partitionOption"`
	PartitionSettings        *TeradataPartitionSettingsResponse `pulumi:"partitionSettings"`
	Query                    interface{}                        `pulumi:"query"`
	QueryTimeout             interface{}                        `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}                        `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                        `pulumi:"sourceRetryWait"`
	Type                     string                             `pulumi:"type"`
}

type TeradataTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Database          interface{}                       `pulumi:"database"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Table             interface{}                       `pulumi:"table"`
	Type              string                            `pulumi:"type"`
}

type TeradataTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Database          interface{}                               `pulumi:"database"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Table             interface{}                               `pulumi:"table"`
	Type              string                                    `pulumi:"type"`
}

type TextFormat struct {
	ColumnDelimiter  interface{} `pulumi:"columnDelimiter"`
	Deserializer     interface{} `pulumi:"deserializer"`
	EncodingName     interface{} `pulumi:"encodingName"`
	EscapeChar       interface{} `pulumi:"escapeChar"`
	FirstRowAsHeader interface{} `pulumi:"firstRowAsHeader"`
	NullValue        interface{} `pulumi:"nullValue"`
	QuoteChar        interface{} `pulumi:"quoteChar"`
	RowDelimiter     interface{} `pulumi:"rowDelimiter"`
	Serializer       interface{} `pulumi:"serializer"`
	SkipLineCount    interface{} `pulumi:"skipLineCount"`
	TreatEmptyAsNull interface{} `pulumi:"treatEmptyAsNull"`
	Type             string      `pulumi:"type"`
}

type TextFormatResponse struct {
	ColumnDelimiter  interface{} `pulumi:"columnDelimiter"`
	Deserializer     interface{} `pulumi:"deserializer"`
	EncodingName     interface{} `pulumi:"encodingName"`
	EscapeChar       interface{} `pulumi:"escapeChar"`
	FirstRowAsHeader interface{} `pulumi:"firstRowAsHeader"`
	NullValue        interface{} `pulumi:"nullValue"`
	QuoteChar        interface{} `pulumi:"quoteChar"`
	RowDelimiter     interface{} `pulumi:"rowDelimiter"`
	Serializer       interface{} `pulumi:"serializer"`
	SkipLineCount    interface{} `pulumi:"skipLineCount"`
	TreatEmptyAsNull interface{} `pulumi:"treatEmptyAsNull"`
	Type             string      `pulumi:"type"`
}

type Transformation struct {
	Dataset       *DatasetReference       `pulumi:"dataset"`
	Description   *string                 `pulumi:"description"`
	Flowlet       *DataFlowReference      `pulumi:"flowlet"`
	LinkedService *LinkedServiceReference `pulumi:"linkedService"`
	Name          string                  `pulumi:"name"`
}

type TransformationResponse struct {
	Dataset       *DatasetReferenceResponse       `pulumi:"dataset"`
	Description   *string                         `pulumi:"description"`
	Flowlet       *DataFlowReferenceResponse      `pulumi:"flowlet"`
	LinkedService *LinkedServiceReferenceResponse `pulumi:"linkedService"`
	Name          string                          `pulumi:"name"`
}

type TriggerDependencyReference struct {
	ReferenceTrigger TriggerReference `pulumi:"referenceTrigger"`
	Type             string           `pulumi:"type"`
}

type TriggerDependencyReferenceResponse struct {
	ReferenceTrigger TriggerReferenceResponse `pulumi:"referenceTrigger"`
	Type             string                   `pulumi:"type"`
}

type TriggerPipelineReference struct {
	Parameters        map[string]interface{} `pulumi:"parameters"`
	PipelineReference *PipelineReference     `pulumi:"pipelineReference"`
}

type TriggerPipelineReferenceResponse struct {
	Parameters        map[string]interface{}     `pulumi:"parameters"`
	PipelineReference *PipelineReferenceResponse `pulumi:"pipelineReference"`
}

type TriggerReference struct {
	ReferenceName string `pulumi:"referenceName"`
	Type          string `pulumi:"type"`
}

type TriggerReferenceResponse struct {
	ReferenceName string `pulumi:"referenceName"`
	Type          string `pulumi:"type"`
}

type TumblingWindowTrigger struct {
	Annotations    []interface{}            `pulumi:"annotations"`
	Delay          interface{}              `pulumi:"delay"`
	DependsOn      []interface{}            `pulumi:"dependsOn"`
	Description    *string                  `pulumi:"description"`
	EndTime        *string                  `pulumi:"endTime"`
	Frequency      string                   `pulumi:"frequency"`
	Interval       int                      `pulumi:"interval"`
	MaxConcurrency int                      `pulumi:"maxConcurrency"`
	Pipeline       TriggerPipelineReference `pulumi:"pipeline"`
	RetryPolicy    *RetryPolicy             `pulumi:"retryPolicy"`
	StartTime      string                   `pulumi:"startTime"`
	Type           string                   `pulumi:"type"`
}

type TumblingWindowTriggerDependencyReference struct {
	Offset           *string          `pulumi:"offset"`
	ReferenceTrigger TriggerReference `pulumi:"referenceTrigger"`
	Size             *string          `pulumi:"size"`
	Type             string           `pulumi:"type"`
}

type TumblingWindowTriggerDependencyReferenceResponse struct {
	Offset           *string                  `pulumi:"offset"`
	ReferenceTrigger TriggerReferenceResponse `pulumi:"referenceTrigger"`
	Size             *string                  `pulumi:"size"`
	Type             string                   `pulumi:"type"`
}

type TumblingWindowTriggerResponse struct {
	Annotations    []interface{}                    `pulumi:"annotations"`
	Delay          interface{}                      `pulumi:"delay"`
	DependsOn      []interface{}                    `pulumi:"dependsOn"`
	Description    *string                          `pulumi:"description"`
	EndTime        *string                          `pulumi:"endTime"`
	Frequency      string                           `pulumi:"frequency"`
	Interval       int                              `pulumi:"interval"`
	MaxConcurrency int                              `pulumi:"maxConcurrency"`
	Pipeline       TriggerPipelineReferenceResponse `pulumi:"pipeline"`
	RetryPolicy    *RetryPolicyResponse             `pulumi:"retryPolicy"`
	RuntimeState   string                           `pulumi:"runtimeState"`
	StartTime      string                           `pulumi:"startTime"`
	Type           string                           `pulumi:"type"`
}

type TwilioLinkedService struct {
	Annotations []interface{}                     `pulumi:"annotations"`
	ConnectVia  *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description *string                           `pulumi:"description"`
	Parameters  map[string]ParameterSpecification `pulumi:"parameters"`
	Password    interface{}                       `pulumi:"password"`
	Type        string                            `pulumi:"type"`
	UserName    interface{}                       `pulumi:"userName"`
}

type TwilioLinkedServiceResponse struct {
	Annotations []interface{}                             `pulumi:"annotations"`
	ConnectVia  *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description *string                                   `pulumi:"description"`
	Parameters  map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password    interface{}                               `pulumi:"password"`
	Type        string                                    `pulumi:"type"`
	UserName    interface{}                               `pulumi:"userName"`
}

type UntilActivity struct {
	Activities     []interface{}        `pulumi:"activities"`
	DependsOn      []ActivityDependency `pulumi:"dependsOn"`
	Description    *string              `pulumi:"description"`
	Expression     Expression           `pulumi:"expression"`
	Name           string               `pulumi:"name"`
	Timeout        interface{}          `pulumi:"timeout"`
	Type           string               `pulumi:"type"`
	UserProperties []UserProperty       `pulumi:"userProperties"`
}

type UntilActivityResponse struct {
	Activities     []interface{}                `pulumi:"activities"`
	DependsOn      []ActivityDependencyResponse `pulumi:"dependsOn"`
	Description    *string                      `pulumi:"description"`
	Expression     ExpressionResponse           `pulumi:"expression"`
	Name           string                       `pulumi:"name"`
	Timeout        interface{}                  `pulumi:"timeout"`
	Type           string                       `pulumi:"type"`
	UserProperties []UserPropertyResponse       `pulumi:"userProperties"`
}

type UserAccessPolicyResponse struct {
	AccessResourcePath *string `pulumi:"accessResourcePath"`
	ExpireTime         *string `pulumi:"expireTime"`
	Permissions        *string `pulumi:"permissions"`
	ProfileName        *string `pulumi:"profileName"`
	StartTime          *string `pulumi:"startTime"`
}

type UserAccessPolicyResponseOutput struct{ *pulumi.OutputState }

func (UserAccessPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAccessPolicyResponse)(nil)).Elem()
}

func (o UserAccessPolicyResponseOutput) ToUserAccessPolicyResponseOutput() UserAccessPolicyResponseOutput {
	return o
}

func (o UserAccessPolicyResponseOutput) ToUserAccessPolicyResponseOutputWithContext(ctx context.Context) UserAccessPolicyResponseOutput {
	return o
}

func (o UserAccessPolicyResponseOutput) AccessResourcePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAccessPolicyResponse) *string { return v.AccessResourcePath }).(pulumi.StringPtrOutput)
}

func (o UserAccessPolicyResponseOutput) ExpireTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAccessPolicyResponse) *string { return v.ExpireTime }).(pulumi.StringPtrOutput)
}

func (o UserAccessPolicyResponseOutput) Permissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAccessPolicyResponse) *string { return v.Permissions }).(pulumi.StringPtrOutput)
}

func (o UserAccessPolicyResponseOutput) ProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAccessPolicyResponse) *string { return v.ProfileName }).(pulumi.StringPtrOutput)
}

func (o UserAccessPolicyResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAccessPolicyResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type UserAccessPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (UserAccessPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAccessPolicyResponse)(nil)).Elem()
}

func (o UserAccessPolicyResponsePtrOutput) ToUserAccessPolicyResponsePtrOutput() UserAccessPolicyResponsePtrOutput {
	return o
}

func (o UserAccessPolicyResponsePtrOutput) ToUserAccessPolicyResponsePtrOutputWithContext(ctx context.Context) UserAccessPolicyResponsePtrOutput {
	return o
}

func (o UserAccessPolicyResponsePtrOutput) Elem() UserAccessPolicyResponseOutput {
	return o.ApplyT(func(v *UserAccessPolicyResponse) UserAccessPolicyResponse {
		if v != nil {
			return *v
		}
		var ret UserAccessPolicyResponse
		return ret
	}).(UserAccessPolicyResponseOutput)
}

func (o UserAccessPolicyResponsePtrOutput) AccessResourcePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAccessPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccessResourcePath
	}).(pulumi.StringPtrOutput)
}

func (o UserAccessPolicyResponsePtrOutput) ExpireTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAccessPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExpireTime
	}).(pulumi.StringPtrOutput)
}

func (o UserAccessPolicyResponsePtrOutput) Permissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAccessPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Permissions
	}).(pulumi.StringPtrOutput)
}

func (o UserAccessPolicyResponsePtrOutput) ProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAccessPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProfileName
	}).(pulumi.StringPtrOutput)
}

func (o UserAccessPolicyResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAccessPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

type UserProperty struct {
	Name  string      `pulumi:"name"`
	Value interface{} `pulumi:"value"`
}

type UserPropertyResponse struct {
	Name  string      `pulumi:"name"`
	Value interface{} `pulumi:"value"`
}

type ValidationActivity struct {
	ChildItems     interface{}          `pulumi:"childItems"`
	Dataset        DatasetReference     `pulumi:"dataset"`
	DependsOn      []ActivityDependency `pulumi:"dependsOn"`
	Description    *string              `pulumi:"description"`
	MinimumSize    interface{}          `pulumi:"minimumSize"`
	Name           string               `pulumi:"name"`
	Sleep          interface{}          `pulumi:"sleep"`
	Timeout        interface{}          `pulumi:"timeout"`
	Type           string               `pulumi:"type"`
	UserProperties []UserProperty       `pulumi:"userProperties"`
}

type ValidationActivityResponse struct {
	ChildItems     interface{}                  `pulumi:"childItems"`
	Dataset        DatasetReferenceResponse     `pulumi:"dataset"`
	DependsOn      []ActivityDependencyResponse `pulumi:"dependsOn"`
	Description    *string                      `pulumi:"description"`
	MinimumSize    interface{}                  `pulumi:"minimumSize"`
	Name           string                       `pulumi:"name"`
	Sleep          interface{}                  `pulumi:"sleep"`
	Timeout        interface{}                  `pulumi:"timeout"`
	Type           string                       `pulumi:"type"`
	UserProperties []UserPropertyResponse       `pulumi:"userProperties"`
}

type VariableSpecification struct {
	DefaultValue interface{} `pulumi:"defaultValue"`
	Type         string      `pulumi:"type"`
}





type VariableSpecificationInput interface {
	pulumi.Input

	ToVariableSpecificationOutput() VariableSpecificationOutput
	ToVariableSpecificationOutputWithContext(context.Context) VariableSpecificationOutput
}

type VariableSpecificationArgs struct {
	DefaultValue pulumi.Input       `pulumi:"defaultValue"`
	Type         pulumi.StringInput `pulumi:"type"`
}

func (VariableSpecificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VariableSpecification)(nil)).Elem()
}

func (i VariableSpecificationArgs) ToVariableSpecificationOutput() VariableSpecificationOutput {
	return i.ToVariableSpecificationOutputWithContext(context.Background())
}

func (i VariableSpecificationArgs) ToVariableSpecificationOutputWithContext(ctx context.Context) VariableSpecificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableSpecificationOutput)
}





type VariableSpecificationMapInput interface {
	pulumi.Input

	ToVariableSpecificationMapOutput() VariableSpecificationMapOutput
	ToVariableSpecificationMapOutputWithContext(context.Context) VariableSpecificationMapOutput
}

type VariableSpecificationMap map[string]VariableSpecificationInput

func (VariableSpecificationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VariableSpecification)(nil)).Elem()
}

func (i VariableSpecificationMap) ToVariableSpecificationMapOutput() VariableSpecificationMapOutput {
	return i.ToVariableSpecificationMapOutputWithContext(context.Background())
}

func (i VariableSpecificationMap) ToVariableSpecificationMapOutputWithContext(ctx context.Context) VariableSpecificationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableSpecificationMapOutput)
}

type VariableSpecificationOutput struct{ *pulumi.OutputState }

func (VariableSpecificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VariableSpecification)(nil)).Elem()
}

func (o VariableSpecificationOutput) ToVariableSpecificationOutput() VariableSpecificationOutput {
	return o
}

func (o VariableSpecificationOutput) ToVariableSpecificationOutputWithContext(ctx context.Context) VariableSpecificationOutput {
	return o
}

func (o VariableSpecificationOutput) DefaultValue() pulumi.AnyOutput {
	return o.ApplyT(func(v VariableSpecification) interface{} { return v.DefaultValue }).(pulumi.AnyOutput)
}

func (o VariableSpecificationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VariableSpecification) string { return v.Type }).(pulumi.StringOutput)
}

type VariableSpecificationMapOutput struct{ *pulumi.OutputState }

func (VariableSpecificationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VariableSpecification)(nil)).Elem()
}

func (o VariableSpecificationMapOutput) ToVariableSpecificationMapOutput() VariableSpecificationMapOutput {
	return o
}

func (o VariableSpecificationMapOutput) ToVariableSpecificationMapOutputWithContext(ctx context.Context) VariableSpecificationMapOutput {
	return o
}

func (o VariableSpecificationMapOutput) MapIndex(k pulumi.StringInput) VariableSpecificationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VariableSpecification {
		return vs[0].(map[string]VariableSpecification)[vs[1].(string)]
	}).(VariableSpecificationOutput)
}

type VariableSpecificationResponse struct {
	DefaultValue interface{} `pulumi:"defaultValue"`
	Type         string      `pulumi:"type"`
}

type VariableSpecificationResponseOutput struct{ *pulumi.OutputState }

func (VariableSpecificationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VariableSpecificationResponse)(nil)).Elem()
}

func (o VariableSpecificationResponseOutput) ToVariableSpecificationResponseOutput() VariableSpecificationResponseOutput {
	return o
}

func (o VariableSpecificationResponseOutput) ToVariableSpecificationResponseOutputWithContext(ctx context.Context) VariableSpecificationResponseOutput {
	return o
}

func (o VariableSpecificationResponseOutput) DefaultValue() pulumi.AnyOutput {
	return o.ApplyT(func(v VariableSpecificationResponse) interface{} { return v.DefaultValue }).(pulumi.AnyOutput)
}

func (o VariableSpecificationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VariableSpecificationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type VariableSpecificationResponseMapOutput struct{ *pulumi.OutputState }

func (VariableSpecificationResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VariableSpecificationResponse)(nil)).Elem()
}

func (o VariableSpecificationResponseMapOutput) ToVariableSpecificationResponseMapOutput() VariableSpecificationResponseMapOutput {
	return o
}

func (o VariableSpecificationResponseMapOutput) ToVariableSpecificationResponseMapOutputWithContext(ctx context.Context) VariableSpecificationResponseMapOutput {
	return o
}

func (o VariableSpecificationResponseMapOutput) MapIndex(k pulumi.StringInput) VariableSpecificationResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VariableSpecificationResponse {
		return vs[0].(map[string]VariableSpecificationResponse)[vs[1].(string)]
	}).(VariableSpecificationResponseOutput)
}

type VerticaLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Pwd                 *AzureKeyVaultSecretReference     `pulumi:"pwd"`
	Type                string                            `pulumi:"type"`
}

type VerticaLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Pwd                 *AzureKeyVaultSecretReferenceResponse     `pulumi:"pwd"`
	Type                string                                    `pulumi:"type"`
}

type VerticaSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type VerticaSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type VerticaTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Table             interface{}                       `pulumi:"table"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type VerticaTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Table             interface{}                               `pulumi:"table"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type WaitActivity struct {
	DependsOn         []ActivityDependency `pulumi:"dependsOn"`
	Description       *string              `pulumi:"description"`
	Name              string               `pulumi:"name"`
	Type              string               `pulumi:"type"`
	UserProperties    []UserProperty       `pulumi:"userProperties"`
	WaitTimeInSeconds interface{}          `pulumi:"waitTimeInSeconds"`
}

type WaitActivityResponse struct {
	DependsOn         []ActivityDependencyResponse `pulumi:"dependsOn"`
	Description       *string                      `pulumi:"description"`
	Name              string                       `pulumi:"name"`
	Type              string                       `pulumi:"type"`
	UserProperties    []UserPropertyResponse       `pulumi:"userProperties"`
	WaitTimeInSeconds interface{}                  `pulumi:"waitTimeInSeconds"`
}

type WebActivity struct {
	Authentication        *WebActivityAuthentication   `pulumi:"authentication"`
	Body                  interface{}                  `pulumi:"body"`
	ConnectVia            *IntegrationRuntimeReference `pulumi:"connectVia"`
	Datasets              []DatasetReference           `pulumi:"datasets"`
	DependsOn             []ActivityDependency         `pulumi:"dependsOn"`
	Description           *string                      `pulumi:"description"`
	DisableCertValidation *bool                        `pulumi:"disableCertValidation"`
	Headers               interface{}                  `pulumi:"headers"`
	LinkedServiceName     *LinkedServiceReference      `pulumi:"linkedServiceName"`
	LinkedServices        []LinkedServiceReference     `pulumi:"linkedServices"`
	Method                string                       `pulumi:"method"`
	Name                  string                       `pulumi:"name"`
	Policy                *ActivityPolicy              `pulumi:"policy"`
	Type                  string                       `pulumi:"type"`
	Url                   interface{}                  `pulumi:"url"`
	UserProperties        []UserProperty               `pulumi:"userProperties"`
}

type WebActivityAuthentication struct {
	Credential *CredentialReference `pulumi:"credential"`
	Password   interface{}          `pulumi:"password"`
	Pfx        interface{}          `pulumi:"pfx"`
	Resource   interface{}          `pulumi:"resource"`
	Type       *string              `pulumi:"type"`
	UserTenant interface{}          `pulumi:"userTenant"`
	Username   interface{}          `pulumi:"username"`
}

type WebActivityAuthenticationResponse struct {
	Credential *CredentialReferenceResponse `pulumi:"credential"`
	Password   interface{}                  `pulumi:"password"`
	Pfx        interface{}                  `pulumi:"pfx"`
	Resource   interface{}                  `pulumi:"resource"`
	Type       *string                      `pulumi:"type"`
	UserTenant interface{}                  `pulumi:"userTenant"`
	Username   interface{}                  `pulumi:"username"`
}

type WebActivityResponse struct {
	Authentication        *WebActivityAuthenticationResponse   `pulumi:"authentication"`
	Body                  interface{}                          `pulumi:"body"`
	ConnectVia            *IntegrationRuntimeReferenceResponse `pulumi:"connectVia"`
	Datasets              []DatasetReferenceResponse           `pulumi:"datasets"`
	DependsOn             []ActivityDependencyResponse         `pulumi:"dependsOn"`
	Description           *string                              `pulumi:"description"`
	DisableCertValidation *bool                                `pulumi:"disableCertValidation"`
	Headers               interface{}                          `pulumi:"headers"`
	LinkedServiceName     *LinkedServiceReferenceResponse      `pulumi:"linkedServiceName"`
	LinkedServices        []LinkedServiceReferenceResponse     `pulumi:"linkedServices"`
	Method                string                               `pulumi:"method"`
	Name                  string                               `pulumi:"name"`
	Policy                *ActivityPolicyResponse              `pulumi:"policy"`
	Type                  string                               `pulumi:"type"`
	Url                   interface{}                          `pulumi:"url"`
	UserProperties        []UserPropertyResponse               `pulumi:"userProperties"`
}

type WebAnonymousAuthentication struct {
	AuthenticationType string      `pulumi:"authenticationType"`
	Url                interface{} `pulumi:"url"`
}

type WebAnonymousAuthenticationResponse struct {
	AuthenticationType string      `pulumi:"authenticationType"`
	Url                interface{} `pulumi:"url"`
}

type WebBasicAuthentication struct {
	AuthenticationType string      `pulumi:"authenticationType"`
	Password           interface{} `pulumi:"password"`
	Url                interface{} `pulumi:"url"`
	Username           interface{} `pulumi:"username"`
}

type WebBasicAuthenticationResponse struct {
	AuthenticationType string      `pulumi:"authenticationType"`
	Password           interface{} `pulumi:"password"`
	Url                interface{} `pulumi:"url"`
	Username           interface{} `pulumi:"username"`
}

type WebClientCertificateAuthentication struct {
	AuthenticationType string      `pulumi:"authenticationType"`
	Password           interface{} `pulumi:"password"`
	Pfx                interface{} `pulumi:"pfx"`
	Url                interface{} `pulumi:"url"`
}

type WebClientCertificateAuthenticationResponse struct {
	AuthenticationType string      `pulumi:"authenticationType"`
	Password           interface{} `pulumi:"password"`
	Pfx                interface{} `pulumi:"pfx"`
	Url                interface{} `pulumi:"url"`
}

type WebHookActivity struct {
	Authentication         *WebActivityAuthentication `pulumi:"authentication"`
	Body                   interface{}                `pulumi:"body"`
	DependsOn              []ActivityDependency       `pulumi:"dependsOn"`
	Description            *string                    `pulumi:"description"`
	Headers                interface{}                `pulumi:"headers"`
	Method                 string                     `pulumi:"method"`
	Name                   string                     `pulumi:"name"`
	ReportStatusOnCallBack interface{}                `pulumi:"reportStatusOnCallBack"`
	Timeout                *string                    `pulumi:"timeout"`
	Type                   string                     `pulumi:"type"`
	Url                    interface{}                `pulumi:"url"`
	UserProperties         []UserProperty             `pulumi:"userProperties"`
}

type WebHookActivityResponse struct {
	Authentication         *WebActivityAuthenticationResponse `pulumi:"authentication"`
	Body                   interface{}                        `pulumi:"body"`
	DependsOn              []ActivityDependencyResponse       `pulumi:"dependsOn"`
	Description            *string                            `pulumi:"description"`
	Headers                interface{}                        `pulumi:"headers"`
	Method                 string                             `pulumi:"method"`
	Name                   string                             `pulumi:"name"`
	ReportStatusOnCallBack interface{}                        `pulumi:"reportStatusOnCallBack"`
	Timeout                *string                            `pulumi:"timeout"`
	Type                   string                             `pulumi:"type"`
	Url                    interface{}                        `pulumi:"url"`
	UserProperties         []UserPropertyResponse             `pulumi:"userProperties"`
}

type WebLinkedService struct {
	Annotations    []interface{}                     `pulumi:"annotations"`
	ConnectVia     *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description    *string                           `pulumi:"description"`
	Parameters     map[string]ParameterSpecification `pulumi:"parameters"`
	Type           string                            `pulumi:"type"`
	TypeProperties interface{}                       `pulumi:"typeProperties"`
}

type WebLinkedServiceResponse struct {
	Annotations    []interface{}                             `pulumi:"annotations"`
	ConnectVia     *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description    *string                                   `pulumi:"description"`
	Parameters     map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type           string                                    `pulumi:"type"`
	TypeProperties interface{}                               `pulumi:"typeProperties"`
}

type WebSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type WebSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type WebTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	Index             interface{}                       `pulumi:"index"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Path              interface{}                       `pulumi:"path"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type WebTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	Index             interface{}                               `pulumi:"index"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Path              interface{}                               `pulumi:"path"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type WranglingDataFlow struct {
	Annotations    []interface{}      `pulumi:"annotations"`
	Description    *string            `pulumi:"description"`
	DocumentLocale *string            `pulumi:"documentLocale"`
	Folder         *DataFlowFolder    `pulumi:"folder"`
	Script         *string            `pulumi:"script"`
	Sources        []PowerQuerySource `pulumi:"sources"`
	Type           string             `pulumi:"type"`
}

type WranglingDataFlowResponse struct {
	Annotations    []interface{}              `pulumi:"annotations"`
	Description    *string                    `pulumi:"description"`
	DocumentLocale *string                    `pulumi:"documentLocale"`
	Folder         *DataFlowResponseFolder    `pulumi:"folder"`
	Script         *string                    `pulumi:"script"`
	Sources        []PowerQuerySourceResponse `pulumi:"sources"`
	Type           string                     `pulumi:"type"`
}

type XeroLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionProperties  interface{}                       `pulumi:"connectionProperties"`
	ConsumerKey           interface{}                       `pulumi:"consumerKey"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Host                  interface{}                       `pulumi:"host"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	PrivateKey            interface{}                       `pulumi:"privateKey"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
}

type XeroLinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionProperties  interface{}                               `pulumi:"connectionProperties"`
	ConsumerKey           interface{}                               `pulumi:"consumerKey"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Host                  interface{}                               `pulumi:"host"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	PrivateKey            interface{}                               `pulumi:"privateKey"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
}

type XeroObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type XeroObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type XeroSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type XeroSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type XmlDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Compression       *DatasetCompression               `pulumi:"compression"`
	Description       *string                           `pulumi:"description"`
	EncodingName      interface{}                       `pulumi:"encodingName"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Location          interface{}                       `pulumi:"location"`
	NullValue         interface{}                       `pulumi:"nullValue"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type XmlDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Compression       *DatasetCompressionResponse               `pulumi:"compression"`
	Description       *string                                   `pulumi:"description"`
	EncodingName      interface{}                               `pulumi:"encodingName"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Location          interface{}                               `pulumi:"location"`
	NullValue         interface{}                               `pulumi:"nullValue"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type XmlReadSettings struct {
	CompressionProperties interface{} `pulumi:"compressionProperties"`
	DetectDataType        interface{} `pulumi:"detectDataType"`
	NamespacePrefixes     interface{} `pulumi:"namespacePrefixes"`
	Namespaces            interface{} `pulumi:"namespaces"`
	Type                  string      `pulumi:"type"`
	ValidationMode        interface{} `pulumi:"validationMode"`
}

type XmlReadSettingsResponse struct {
	CompressionProperties interface{} `pulumi:"compressionProperties"`
	DetectDataType        interface{} `pulumi:"detectDataType"`
	NamespacePrefixes     interface{} `pulumi:"namespacePrefixes"`
	Namespaces            interface{} `pulumi:"namespaces"`
	Type                  string      `pulumi:"type"`
	ValidationMode        interface{} `pulumi:"validationMode"`
}

type XmlSource struct {
	AdditionalColumns        interface{}      `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}      `pulumi:"disableMetricsCollection"`
	FormatSettings           *XmlReadSettings `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}      `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{}      `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}      `pulumi:"sourceRetryWait"`
	StoreSettings            interface{}      `pulumi:"storeSettings"`
	Type                     string           `pulumi:"type"`
}

type XmlSourceResponse struct {
	AdditionalColumns        interface{}              `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}              `pulumi:"disableMetricsCollection"`
	FormatSettings           *XmlReadSettingsResponse `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}              `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{}              `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}              `pulumi:"sourceRetryWait"`
	StoreSettings            interface{}              `pulumi:"storeSettings"`
	Type                     string                   `pulumi:"type"`
}

type ZendeskLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ApiToken            interface{}                       `pulumi:"apiToken"`
	AuthenticationType  string                            `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Type                string                            `pulumi:"type"`
	Url                 interface{}                       `pulumi:"url"`
	UserName            interface{}                       `pulumi:"userName"`
}

type ZendeskLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ApiToken            interface{}                               `pulumi:"apiToken"`
	AuthenticationType  string                                    `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
	Url                 interface{}                               `pulumi:"url"`
	UserName            interface{}                               `pulumi:"userName"`
}

type ZipDeflateReadSettings struct {
	PreserveZipFileNameAsFolder interface{} `pulumi:"preserveZipFileNameAsFolder"`
	Type                        string      `pulumi:"type"`
}

type ZipDeflateReadSettingsResponse struct {
	PreserveZipFileNameAsFolder interface{} `pulumi:"preserveZipFileNameAsFolder"`
	Type                        string      `pulumi:"type"`
}

type ZohoLinkedService struct {
	AccessToken           interface{}                       `pulumi:"accessToken"`
	Annotations           []interface{}                     `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionProperties  interface{}                       `pulumi:"connectionProperties"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Endpoint              interface{}                       `pulumi:"endpoint"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
}

type ZohoLinkedServiceResponse struct {
	AccessToken           interface{}                               `pulumi:"accessToken"`
	Annotations           []interface{}                             `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionProperties  interface{}                               `pulumi:"connectionProperties"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Endpoint              interface{}                               `pulumi:"endpoint"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
}

type ZohoObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type ZohoObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type ZohoSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type ZohoSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

func init() {
	pulumi.RegisterOutputType(UserAccessPolicyResponseOutput{})
	pulumi.RegisterOutputType(UserAccessPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(VariableSpecificationOutput{})
	pulumi.RegisterOutputType(VariableSpecificationMapOutput{})
	pulumi.RegisterOutputType(VariableSpecificationResponseOutput{})
	pulumi.RegisterOutputType(VariableSpecificationResponseMapOutput{})
}
