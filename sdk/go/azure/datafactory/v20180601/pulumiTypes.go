


package v20180601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActivityDependency struct {
	Activity             string   `pulumi:"activity"`
	DependencyConditions []string `pulumi:"dependencyConditions"`
}

type ActivityDependencyResponse struct {
	Activity             string   `pulumi:"activity"`
	DependencyConditions []string `pulumi:"dependencyConditions"`
}

type ActivityPolicy struct {
	Retry                  interface{} `pulumi:"retry"`
	RetryIntervalInSeconds *int        `pulumi:"retryIntervalInSeconds"`
	SecureInput            *bool       `pulumi:"secureInput"`
	SecureOutput           *bool       `pulumi:"secureOutput"`
	Timeout                interface{} `pulumi:"timeout"`
}

type ActivityPolicyResponse struct {
	Retry                  interface{} `pulumi:"retry"`
	RetryIntervalInSeconds *int        `pulumi:"retryIntervalInSeconds"`
	SecureInput            *bool       `pulumi:"secureInput"`
	SecureOutput           *bool       `pulumi:"secureOutput"`
	Timeout                interface{} `pulumi:"timeout"`
}

type AmazonMWSLinkedService struct {
	AccessKeyId           interface{}                       `pulumi:"accessKeyId"`
	Annotations           []interface{}                     `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Endpoint              interface{}                       `pulumi:"endpoint"`
	MarketplaceID         interface{}                       `pulumi:"marketplaceID"`
	MwsAuthToken          interface{}                       `pulumi:"mwsAuthToken"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	SecretKey             interface{}                       `pulumi:"secretKey"`
	SellerID              interface{}                       `pulumi:"sellerID"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
}

type AmazonMWSLinkedServiceResponse struct {
	AccessKeyId           interface{}                               `pulumi:"accessKeyId"`
	Annotations           []interface{}                             `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Endpoint              interface{}                               `pulumi:"endpoint"`
	MarketplaceID         interface{}                               `pulumi:"marketplaceID"`
	MwsAuthToken          interface{}                               `pulumi:"mwsAuthToken"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	SecretKey             interface{}                               `pulumi:"secretKey"`
	SellerID              interface{}                               `pulumi:"sellerID"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
}

type AmazonMWSObjectDataset struct {
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

type AmazonMWSObjectDatasetResponse struct {
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

type AmazonMWSSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type AmazonMWSSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type AmazonRdsForOracleLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Type                string                            `pulumi:"type"`
}

type AmazonRdsForOracleLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
}

type AmazonRdsForOraclePartitionSettings struct {
	PartitionColumnName interface{}   `pulumi:"partitionColumnName"`
	PartitionLowerBound interface{}   `pulumi:"partitionLowerBound"`
	PartitionNames      []interface{} `pulumi:"partitionNames"`
	PartitionUpperBound interface{}   `pulumi:"partitionUpperBound"`
}

type AmazonRdsForOraclePartitionSettingsResponse struct {
	PartitionColumnName interface{}   `pulumi:"partitionColumnName"`
	PartitionLowerBound interface{}   `pulumi:"partitionLowerBound"`
	PartitionNames      []interface{} `pulumi:"partitionNames"`
	PartitionUpperBound interface{}   `pulumi:"partitionUpperBound"`
}

type AmazonRdsForOracleSource struct {
	AdditionalColumns        interface{}                          `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}                          `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}                          `pulumi:"maxConcurrentConnections"`
	OracleReaderQuery        interface{}                          `pulumi:"oracleReaderQuery"`
	PartitionOption          interface{}                          `pulumi:"partitionOption"`
	PartitionSettings        *AmazonRdsForOraclePartitionSettings `pulumi:"partitionSettings"`
	QueryTimeout             interface{}                          `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}                          `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                          `pulumi:"sourceRetryWait"`
	Type                     string                               `pulumi:"type"`
}

type AmazonRdsForOracleSourceResponse struct {
	AdditionalColumns        interface{}                                  `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}                                  `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}                                  `pulumi:"maxConcurrentConnections"`
	OracleReaderQuery        interface{}                                  `pulumi:"oracleReaderQuery"`
	PartitionOption          interface{}                                  `pulumi:"partitionOption"`
	PartitionSettings        *AmazonRdsForOraclePartitionSettingsResponse `pulumi:"partitionSettings"`
	QueryTimeout             interface{}                                  `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}                                  `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                                  `pulumi:"sourceRetryWait"`
	Type                     string                                       `pulumi:"type"`
}

type AmazonRdsForOracleTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Table             interface{}                       `pulumi:"table"`
	Type              string                            `pulumi:"type"`
}

type AmazonRdsForOracleTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Table             interface{}                               `pulumi:"table"`
	Type              string                                    `pulumi:"type"`
}

type AmazonRdsForSqlServerLinkedService struct {
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

type AmazonRdsForSqlServerLinkedServiceResponse struct {
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

type AmazonRdsForSqlServerSource struct {
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

type AmazonRdsForSqlServerSourceResponse struct {
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

type AmazonRdsForSqlServerTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Table             interface{}                       `pulumi:"table"`
	Type              string                            `pulumi:"type"`
}

type AmazonRdsForSqlServerTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Table             interface{}                               `pulumi:"table"`
	Type              string                                    `pulumi:"type"`
}

type AmazonRedshiftLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Database            interface{}                       `pulumi:"database"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Port                interface{}                       `pulumi:"port"`
	Server              interface{}                       `pulumi:"server"`
	Type                string                            `pulumi:"type"`
	Username            interface{}                       `pulumi:"username"`
}

type AmazonRedshiftLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Database            interface{}                               `pulumi:"database"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Port                interface{}                               `pulumi:"port"`
	Server              interface{}                               `pulumi:"server"`
	Type                string                                    `pulumi:"type"`
	Username            interface{}                               `pulumi:"username"`
}

type AmazonRedshiftSource struct {
	AdditionalColumns        interface{}             `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}             `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}             `pulumi:"maxConcurrentConnections"`
	Query                    interface{}             `pulumi:"query"`
	QueryTimeout             interface{}             `pulumi:"queryTimeout"`
	RedshiftUnloadSettings   *RedshiftUnloadSettings `pulumi:"redshiftUnloadSettings"`
	SourceRetryCount         interface{}             `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}             `pulumi:"sourceRetryWait"`
	Type                     string                  `pulumi:"type"`
}

type AmazonRedshiftSourceResponse struct {
	AdditionalColumns        interface{}                     `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}                     `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}                     `pulumi:"maxConcurrentConnections"`
	Query                    interface{}                     `pulumi:"query"`
	QueryTimeout             interface{}                     `pulumi:"queryTimeout"`
	RedshiftUnloadSettings   *RedshiftUnloadSettingsResponse `pulumi:"redshiftUnloadSettings"`
	SourceRetryCount         interface{}                     `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                     `pulumi:"sourceRetryWait"`
	Type                     string                          `pulumi:"type"`
}

type AmazonRedshiftTableDataset struct {
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

type AmazonRedshiftTableDatasetResponse struct {
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

type AmazonS3CompatibleLinkedService struct {
	AccessKeyId         interface{}                       `pulumi:"accessKeyId"`
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	ForcePathStyle      interface{}                       `pulumi:"forcePathStyle"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	SecretAccessKey     interface{}                       `pulumi:"secretAccessKey"`
	ServiceUrl          interface{}                       `pulumi:"serviceUrl"`
	Type                string                            `pulumi:"type"`
}

type AmazonS3CompatibleLinkedServiceResponse struct {
	AccessKeyId         interface{}                               `pulumi:"accessKeyId"`
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	ForcePathStyle      interface{}                               `pulumi:"forcePathStyle"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	SecretAccessKey     interface{}                               `pulumi:"secretAccessKey"`
	ServiceUrl          interface{}                               `pulumi:"serviceUrl"`
	Type                string                                    `pulumi:"type"`
}

type AmazonS3CompatibleLocation struct {
	BucketName interface{} `pulumi:"bucketName"`
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
	Version    interface{} `pulumi:"version"`
}

type AmazonS3CompatibleLocationResponse struct {
	BucketName interface{} `pulumi:"bucketName"`
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
	Version    interface{} `pulumi:"version"`
}

type AmazonS3CompatibleReadSettings struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Prefix                     interface{} `pulumi:"prefix"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type AmazonS3CompatibleReadSettingsResponse struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Prefix                     interface{} `pulumi:"prefix"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type AmazonS3Dataset struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	BucketName            interface{}                       `pulumi:"bucketName"`
	Compression           *DatasetCompression               `pulumi:"compression"`
	Description           *string                           `pulumi:"description"`
	Folder                *DatasetFolder                    `pulumi:"folder"`
	Format                interface{}                       `pulumi:"format"`
	Key                   interface{}                       `pulumi:"key"`
	LinkedServiceName     LinkedServiceReference            `pulumi:"linkedServiceName"`
	ModifiedDatetimeEnd   interface{}                       `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart interface{}                       `pulumi:"modifiedDatetimeStart"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Prefix                interface{}                       `pulumi:"prefix"`
	Schema                interface{}                       `pulumi:"schema"`
	Structure             interface{}                       `pulumi:"structure"`
	Type                  string                            `pulumi:"type"`
	Version               interface{}                       `pulumi:"version"`
}

type AmazonS3DatasetResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	BucketName            interface{}                               `pulumi:"bucketName"`
	Compression           *DatasetCompressionResponse               `pulumi:"compression"`
	Description           *string                                   `pulumi:"description"`
	Folder                *DatasetResponseFolder                    `pulumi:"folder"`
	Format                interface{}                               `pulumi:"format"`
	Key                   interface{}                               `pulumi:"key"`
	LinkedServiceName     LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	ModifiedDatetimeEnd   interface{}                               `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart interface{}                               `pulumi:"modifiedDatetimeStart"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Prefix                interface{}                               `pulumi:"prefix"`
	Schema                interface{}                               `pulumi:"schema"`
	Structure             interface{}                               `pulumi:"structure"`
	Type                  string                                    `pulumi:"type"`
	Version               interface{}                               `pulumi:"version"`
}

type AmazonS3LinkedService struct {
	AccessKeyId         interface{}                       `pulumi:"accessKeyId"`
	Annotations         []interface{}                     `pulumi:"annotations"`
	AuthenticationType  interface{}                       `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	SecretAccessKey     interface{}                       `pulumi:"secretAccessKey"`
	ServiceUrl          interface{}                       `pulumi:"serviceUrl"`
	SessionToken        interface{}                       `pulumi:"sessionToken"`
	Type                string                            `pulumi:"type"`
}

type AmazonS3LinkedServiceResponse struct {
	AccessKeyId         interface{}                               `pulumi:"accessKeyId"`
	Annotations         []interface{}                             `pulumi:"annotations"`
	AuthenticationType  interface{}                               `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	SecretAccessKey     interface{}                               `pulumi:"secretAccessKey"`
	ServiceUrl          interface{}                               `pulumi:"serviceUrl"`
	SessionToken        interface{}                               `pulumi:"sessionToken"`
	Type                string                                    `pulumi:"type"`
}

type AmazonS3Location struct {
	BucketName interface{} `pulumi:"bucketName"`
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
	Version    interface{} `pulumi:"version"`
}

type AmazonS3LocationResponse struct {
	BucketName interface{} `pulumi:"bucketName"`
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
	Version    interface{} `pulumi:"version"`
}

type AmazonS3ReadSettings struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Prefix                     interface{} `pulumi:"prefix"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type AmazonS3ReadSettingsResponse struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Prefix                     interface{} `pulumi:"prefix"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type AppFiguresLinkedService struct {
	Annotations []interface{}                     `pulumi:"annotations"`
	ClientKey   interface{}                       `pulumi:"clientKey"`
	ConnectVia  *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description *string                           `pulumi:"description"`
	Parameters  map[string]ParameterSpecification `pulumi:"parameters"`
	Password    interface{}                       `pulumi:"password"`
	Type        string                            `pulumi:"type"`
	UserName    interface{}                       `pulumi:"userName"`
}

type AppFiguresLinkedServiceResponse struct {
	Annotations []interface{}                             `pulumi:"annotations"`
	ClientKey   interface{}                               `pulumi:"clientKey"`
	ConnectVia  *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description *string                                   `pulumi:"description"`
	Parameters  map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password    interface{}                               `pulumi:"password"`
	Type        string                                    `pulumi:"type"`
	UserName    interface{}                               `pulumi:"userName"`
}

type AppendVariableActivity struct {
	DependsOn      []ActivityDependency `pulumi:"dependsOn"`
	Description    *string              `pulumi:"description"`
	Name           string               `pulumi:"name"`
	Type           string               `pulumi:"type"`
	UserProperties []UserProperty       `pulumi:"userProperties"`
	Value          interface{}          `pulumi:"value"`
	VariableName   *string              `pulumi:"variableName"`
}

type AppendVariableActivityResponse struct {
	DependsOn      []ActivityDependencyResponse `pulumi:"dependsOn"`
	Description    *string                      `pulumi:"description"`
	Name           string                       `pulumi:"name"`
	Type           string                       `pulumi:"type"`
	UserProperties []UserPropertyResponse       `pulumi:"userProperties"`
	Value          interface{}                  `pulumi:"value"`
	VariableName   *string                      `pulumi:"variableName"`
}

type ArmIdWrapperResponse struct {
	Id string `pulumi:"id"`
}

type ArmIdWrapperResponseOutput struct{ *pulumi.OutputState }

func (ArmIdWrapperResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmIdWrapperResponse)(nil)).Elem()
}

func (o ArmIdWrapperResponseOutput) ToArmIdWrapperResponseOutput() ArmIdWrapperResponseOutput {
	return o
}

func (o ArmIdWrapperResponseOutput) ToArmIdWrapperResponseOutputWithContext(ctx context.Context) ArmIdWrapperResponseOutput {
	return o
}

func (o ArmIdWrapperResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ArmIdWrapperResponse) string { return v.Id }).(pulumi.StringOutput)
}

type ArmIdWrapperResponsePtrOutput struct{ *pulumi.OutputState }

func (ArmIdWrapperResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmIdWrapperResponse)(nil)).Elem()
}

func (o ArmIdWrapperResponsePtrOutput) ToArmIdWrapperResponsePtrOutput() ArmIdWrapperResponsePtrOutput {
	return o
}

func (o ArmIdWrapperResponsePtrOutput) ToArmIdWrapperResponsePtrOutputWithContext(ctx context.Context) ArmIdWrapperResponsePtrOutput {
	return o
}

func (o ArmIdWrapperResponsePtrOutput) Elem() ArmIdWrapperResponseOutput {
	return o.ApplyT(func(v *ArmIdWrapperResponse) ArmIdWrapperResponse {
		if v != nil {
			return *v
		}
		var ret ArmIdWrapperResponse
		return ret
	}).(ArmIdWrapperResponseOutput)
}

func (o ArmIdWrapperResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmIdWrapperResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type AsanaLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ApiToken            interface{}                       `pulumi:"apiToken"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Type                string                            `pulumi:"type"`
}

type AsanaLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ApiToken            interface{}                               `pulumi:"apiToken"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
}

type AvroDataset struct {
	Annotations          []interface{}                     `pulumi:"annotations"`
	AvroCompressionCodec interface{}                       `pulumi:"avroCompressionCodec"`
	AvroCompressionLevel *int                              `pulumi:"avroCompressionLevel"`
	Description          *string                           `pulumi:"description"`
	Folder               *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName    LinkedServiceReference            `pulumi:"linkedServiceName"`
	Location             interface{}                       `pulumi:"location"`
	Parameters           map[string]ParameterSpecification `pulumi:"parameters"`
	Schema               interface{}                       `pulumi:"schema"`
	Structure            interface{}                       `pulumi:"structure"`
	Type                 string                            `pulumi:"type"`
}

type AvroDatasetResponse struct {
	Annotations          []interface{}                             `pulumi:"annotations"`
	AvroCompressionCodec interface{}                               `pulumi:"avroCompressionCodec"`
	AvroCompressionLevel *int                                      `pulumi:"avroCompressionLevel"`
	Description          *string                                   `pulumi:"description"`
	Folder               *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName    LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Location             interface{}                               `pulumi:"location"`
	Parameters           map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema               interface{}                               `pulumi:"schema"`
	Structure            interface{}                               `pulumi:"structure"`
	Type                 string                                    `pulumi:"type"`
}

type AvroFormat struct {
	Deserializer interface{} `pulumi:"deserializer"`
	Serializer   interface{} `pulumi:"serializer"`
	Type         string      `pulumi:"type"`
}

type AvroFormatResponse struct {
	Deserializer interface{} `pulumi:"deserializer"`
	Serializer   interface{} `pulumi:"serializer"`
	Type         string      `pulumi:"type"`
}

type AvroSink struct {
	DisableMetricsCollection interface{}        `pulumi:"disableMetricsCollection"`
	FormatSettings           *AvroWriteSettings `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}        `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{}        `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}        `pulumi:"sinkRetryWait"`
	StoreSettings            interface{}        `pulumi:"storeSettings"`
	Type                     string             `pulumi:"type"`
	WriteBatchSize           interface{}        `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}        `pulumi:"writeBatchTimeout"`
}

type AvroSinkResponse struct {
	DisableMetricsCollection interface{}                `pulumi:"disableMetricsCollection"`
	FormatSettings           *AvroWriteSettingsResponse `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}                `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{}                `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}                `pulumi:"sinkRetryWait"`
	StoreSettings            interface{}                `pulumi:"storeSettings"`
	Type                     string                     `pulumi:"type"`
	WriteBatchSize           interface{}                `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}                `pulumi:"writeBatchTimeout"`
}

type AvroSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	StoreSettings            interface{} `pulumi:"storeSettings"`
	Type                     string      `pulumi:"type"`
}

type AvroSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	StoreSettings            interface{} `pulumi:"storeSettings"`
	Type                     string      `pulumi:"type"`
}

type AvroWriteSettings struct {
	FileNamePrefix  interface{} `pulumi:"fileNamePrefix"`
	MaxRowsPerFile  interface{} `pulumi:"maxRowsPerFile"`
	RecordName      *string     `pulumi:"recordName"`
	RecordNamespace *string     `pulumi:"recordNamespace"`
	Type            string      `pulumi:"type"`
}

type AvroWriteSettingsResponse struct {
	FileNamePrefix  interface{} `pulumi:"fileNamePrefix"`
	MaxRowsPerFile  interface{} `pulumi:"maxRowsPerFile"`
	RecordName      *string     `pulumi:"recordName"`
	RecordNamespace *string     `pulumi:"recordNamespace"`
	Type            string      `pulumi:"type"`
}

type AzPowerShellSetup struct {
	Type    string `pulumi:"type"`
	Version string `pulumi:"version"`
}

type AzPowerShellSetupResponse struct {
	Type    string `pulumi:"type"`
	Version string `pulumi:"version"`
}

type AzureBatchLinkedService struct {
	AccessKey           interface{}                       `pulumi:"accessKey"`
	AccountName         interface{}                       `pulumi:"accountName"`
	Annotations         []interface{}                     `pulumi:"annotations"`
	BatchUri            interface{}                       `pulumi:"batchUri"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Credential          *CredentialReference              `pulumi:"credential"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	LinkedServiceName   LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	PoolName            interface{}                       `pulumi:"poolName"`
	Type                string                            `pulumi:"type"`
}

type AzureBatchLinkedServiceResponse struct {
	AccessKey           interface{}                               `pulumi:"accessKey"`
	AccountName         interface{}                               `pulumi:"accountName"`
	Annotations         []interface{}                             `pulumi:"annotations"`
	BatchUri            interface{}                               `pulumi:"batchUri"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Credential          *CredentialReferenceResponse              `pulumi:"credential"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	LinkedServiceName   LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	PoolName            interface{}                               `pulumi:"poolName"`
	Type                string                                    `pulumi:"type"`
}

type AzureBlobDataset struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	Compression           *DatasetCompression               `pulumi:"compression"`
	Description           *string                           `pulumi:"description"`
	FileName              interface{}                       `pulumi:"fileName"`
	Folder                *DatasetFolder                    `pulumi:"folder"`
	FolderPath            interface{}                       `pulumi:"folderPath"`
	Format                interface{}                       `pulumi:"format"`
	LinkedServiceName     LinkedServiceReference            `pulumi:"linkedServiceName"`
	ModifiedDatetimeEnd   interface{}                       `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart interface{}                       `pulumi:"modifiedDatetimeStart"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Schema                interface{}                       `pulumi:"schema"`
	Structure             interface{}                       `pulumi:"structure"`
	TableRootLocation     interface{}                       `pulumi:"tableRootLocation"`
	Type                  string                            `pulumi:"type"`
}

type AzureBlobDatasetResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	Compression           *DatasetCompressionResponse               `pulumi:"compression"`
	Description           *string                                   `pulumi:"description"`
	FileName              interface{}                               `pulumi:"fileName"`
	Folder                *DatasetResponseFolder                    `pulumi:"folder"`
	FolderPath            interface{}                               `pulumi:"folderPath"`
	Format                interface{}                               `pulumi:"format"`
	LinkedServiceName     LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	ModifiedDatetimeEnd   interface{}                               `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart interface{}                               `pulumi:"modifiedDatetimeStart"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema                interface{}                               `pulumi:"schema"`
	Structure             interface{}                               `pulumi:"structure"`
	TableRootLocation     interface{}                               `pulumi:"tableRootLocation"`
	Type                  string                                    `pulumi:"type"`
}

type AzureBlobFSDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Compression       *DatasetCompression               `pulumi:"compression"`
	Description       *string                           `pulumi:"description"`
	FileName          interface{}                       `pulumi:"fileName"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	FolderPath        interface{}                       `pulumi:"folderPath"`
	Format            interface{}                       `pulumi:"format"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type AzureBlobFSDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Compression       *DatasetCompressionResponse               `pulumi:"compression"`
	Description       *string                                   `pulumi:"description"`
	FileName          interface{}                               `pulumi:"fileName"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	FolderPath        interface{}                               `pulumi:"folderPath"`
	Format            interface{}                               `pulumi:"format"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type AzureBlobFSLinkedService struct {
	AccountKey                     interface{}                       `pulumi:"accountKey"`
	Annotations                    []interface{}                     `pulumi:"annotations"`
	AzureCloudType                 interface{}                       `pulumi:"azureCloudType"`
	ConnectVia                     *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Credential                     *CredentialReference              `pulumi:"credential"`
	Description                    *string                           `pulumi:"description"`
	EncryptedCredential            interface{}                       `pulumi:"encryptedCredential"`
	Parameters                     map[string]ParameterSpecification `pulumi:"parameters"`
	ServicePrincipalCredential     interface{}                       `pulumi:"servicePrincipalCredential"`
	ServicePrincipalCredentialType interface{}                       `pulumi:"servicePrincipalCredentialType"`
	ServicePrincipalId             interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey            interface{}                       `pulumi:"servicePrincipalKey"`
	Tenant                         interface{}                       `pulumi:"tenant"`
	Type                           string                            `pulumi:"type"`
	Url                            interface{}                       `pulumi:"url"`
}

type AzureBlobFSLinkedServiceResponse struct {
	AccountKey                     interface{}                               `pulumi:"accountKey"`
	Annotations                    []interface{}                             `pulumi:"annotations"`
	AzureCloudType                 interface{}                               `pulumi:"azureCloudType"`
	ConnectVia                     *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Credential                     *CredentialReferenceResponse              `pulumi:"credential"`
	Description                    *string                                   `pulumi:"description"`
	EncryptedCredential            interface{}                               `pulumi:"encryptedCredential"`
	Parameters                     map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	ServicePrincipalCredential     interface{}                               `pulumi:"servicePrincipalCredential"`
	ServicePrincipalCredentialType interface{}                               `pulumi:"servicePrincipalCredentialType"`
	ServicePrincipalId             interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey            interface{}                               `pulumi:"servicePrincipalKey"`
	Tenant                         interface{}                               `pulumi:"tenant"`
	Type                           string                                    `pulumi:"type"`
	Url                            interface{}                               `pulumi:"url"`
}

type AzureBlobFSLocation struct {
	FileName   interface{} `pulumi:"fileName"`
	FileSystem interface{} `pulumi:"fileSystem"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
}

type AzureBlobFSLocationResponse struct {
	FileName   interface{} `pulumi:"fileName"`
	FileSystem interface{} `pulumi:"fileSystem"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
}

type AzureBlobFSReadSettings struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type AzureBlobFSReadSettingsResponse struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type AzureBlobFSSink struct {
	CopyBehavior             interface{}    `pulumi:"copyBehavior"`
	DisableMetricsCollection interface{}    `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}    `pulumi:"maxConcurrentConnections"`
	Metadata                 []MetadataItem `pulumi:"metadata"`
	SinkRetryCount           interface{}    `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}    `pulumi:"sinkRetryWait"`
	Type                     string         `pulumi:"type"`
	WriteBatchSize           interface{}    `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}    `pulumi:"writeBatchTimeout"`
}

type AzureBlobFSSinkResponse struct {
	CopyBehavior             interface{}            `pulumi:"copyBehavior"`
	DisableMetricsCollection interface{}            `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}            `pulumi:"maxConcurrentConnections"`
	Metadata                 []MetadataItemResponse `pulumi:"metadata"`
	SinkRetryCount           interface{}            `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}            `pulumi:"sinkRetryWait"`
	Type                     string                 `pulumi:"type"`
	WriteBatchSize           interface{}            `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}            `pulumi:"writeBatchTimeout"`
}

type AzureBlobFSSource struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Recursive                interface{} `pulumi:"recursive"`
	SkipHeaderLineCount      interface{} `pulumi:"skipHeaderLineCount"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	TreatEmptyAsNull         interface{} `pulumi:"treatEmptyAsNull"`
	Type                     string      `pulumi:"type"`
}

type AzureBlobFSSourceResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Recursive                interface{} `pulumi:"recursive"`
	SkipHeaderLineCount      interface{} `pulumi:"skipHeaderLineCount"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	TreatEmptyAsNull         interface{} `pulumi:"treatEmptyAsNull"`
	Type                     string      `pulumi:"type"`
}

type AzureBlobFSWriteSettings struct {
	BlockSizeInMB            interface{} `pulumi:"blockSizeInMB"`
	CopyBehavior             interface{} `pulumi:"copyBehavior"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Type                     string      `pulumi:"type"`
}

type AzureBlobFSWriteSettingsResponse struct {
	BlockSizeInMB            interface{} `pulumi:"blockSizeInMB"`
	CopyBehavior             interface{} `pulumi:"copyBehavior"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Type                     string      `pulumi:"type"`
}

type AzureBlobStorageLinkedService struct {
	AccountKey          *AzureKeyVaultSecretReference     `pulumi:"accountKey"`
	AccountKind         *string                           `pulumi:"accountKind"`
	Annotations         []interface{}                     `pulumi:"annotations"`
	AzureCloudType      interface{}                       `pulumi:"azureCloudType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Credential          *CredentialReference              `pulumi:"credential"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential *string                           `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	SasToken            *AzureKeyVaultSecretReference     `pulumi:"sasToken"`
	SasUri              interface{}                       `pulumi:"sasUri"`
	ServiceEndpoint     *string                           `pulumi:"serviceEndpoint"`
	ServicePrincipalId  interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey interface{}                       `pulumi:"servicePrincipalKey"`
	Tenant              interface{}                       `pulumi:"tenant"`
	Type                string                            `pulumi:"type"`
}

type AzureBlobStorageLinkedServiceResponse struct {
	AccountKey          *AzureKeyVaultSecretReferenceResponse     `pulumi:"accountKey"`
	AccountKind         *string                                   `pulumi:"accountKind"`
	Annotations         []interface{}                             `pulumi:"annotations"`
	AzureCloudType      interface{}                               `pulumi:"azureCloudType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Credential          *CredentialReferenceResponse              `pulumi:"credential"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential *string                                   `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	SasToken            *AzureKeyVaultSecretReferenceResponse     `pulumi:"sasToken"`
	SasUri              interface{}                               `pulumi:"sasUri"`
	ServiceEndpoint     *string                                   `pulumi:"serviceEndpoint"`
	ServicePrincipalId  interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey interface{}                               `pulumi:"servicePrincipalKey"`
	Tenant              interface{}                               `pulumi:"tenant"`
	Type                string                                    `pulumi:"type"`
}

type AzureBlobStorageLocation struct {
	Container  interface{} `pulumi:"container"`
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
}

type AzureBlobStorageLocationResponse struct {
	Container  interface{} `pulumi:"container"`
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
}

type AzureBlobStorageReadSettings struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Prefix                     interface{} `pulumi:"prefix"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type AzureBlobStorageReadSettingsResponse struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Prefix                     interface{} `pulumi:"prefix"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type AzureBlobStorageWriteSettings struct {
	BlockSizeInMB            interface{} `pulumi:"blockSizeInMB"`
	CopyBehavior             interface{} `pulumi:"copyBehavior"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Type                     string      `pulumi:"type"`
}

type AzureBlobStorageWriteSettingsResponse struct {
	BlockSizeInMB            interface{} `pulumi:"blockSizeInMB"`
	CopyBehavior             interface{} `pulumi:"copyBehavior"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Type                     string      `pulumi:"type"`
}

type AzureDataExplorerCommandActivity struct {
	Command           interface{}             `pulumi:"command"`
	CommandTimeout    interface{}             `pulumi:"commandTimeout"`
	DependsOn         []ActivityDependency    `pulumi:"dependsOn"`
	Description       *string                 `pulumi:"description"`
	LinkedServiceName *LinkedServiceReference `pulumi:"linkedServiceName"`
	Name              string                  `pulumi:"name"`
	Policy            *ActivityPolicy         `pulumi:"policy"`
	Type              string                  `pulumi:"type"`
	UserProperties    []UserProperty          `pulumi:"userProperties"`
}

type AzureDataExplorerCommandActivityResponse struct {
	Command           interface{}                     `pulumi:"command"`
	CommandTimeout    interface{}                     `pulumi:"commandTimeout"`
	DependsOn         []ActivityDependencyResponse    `pulumi:"dependsOn"`
	Description       *string                         `pulumi:"description"`
	LinkedServiceName *LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	Name              string                          `pulumi:"name"`
	Policy            *ActivityPolicyResponse         `pulumi:"policy"`
	Type              string                          `pulumi:"type"`
	UserProperties    []UserPropertyResponse          `pulumi:"userProperties"`
}

type AzureDataExplorerLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Credential          *CredentialReference              `pulumi:"credential"`
	Database            interface{}                       `pulumi:"database"`
	Description         *string                           `pulumi:"description"`
	Endpoint            interface{}                       `pulumi:"endpoint"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	ServicePrincipalId  interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey interface{}                       `pulumi:"servicePrincipalKey"`
	Tenant              interface{}                       `pulumi:"tenant"`
	Type                string                            `pulumi:"type"`
}

type AzureDataExplorerLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Credential          *CredentialReferenceResponse              `pulumi:"credential"`
	Database            interface{}                               `pulumi:"database"`
	Description         *string                                   `pulumi:"description"`
	Endpoint            interface{}                               `pulumi:"endpoint"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	ServicePrincipalId  interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey interface{}                               `pulumi:"servicePrincipalKey"`
	Tenant              interface{}                               `pulumi:"tenant"`
	Type                string                                    `pulumi:"type"`
}

type AzureDataExplorerSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	FlushImmediately         interface{} `pulumi:"flushImmediately"`
	IngestionMappingAsJson   interface{} `pulumi:"ingestionMappingAsJson"`
	IngestionMappingName     interface{} `pulumi:"ingestionMappingName"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type AzureDataExplorerSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	FlushImmediately         interface{} `pulumi:"flushImmediately"`
	IngestionMappingAsJson   interface{} `pulumi:"ingestionMappingAsJson"`
	IngestionMappingName     interface{} `pulumi:"ingestionMappingName"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type AzureDataExplorerSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	NoTruncation             interface{} `pulumi:"noTruncation"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type AzureDataExplorerSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	NoTruncation             interface{} `pulumi:"noTruncation"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type AzureDataExplorerTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Table             interface{}                       `pulumi:"table"`
	Type              string                            `pulumi:"type"`
}

type AzureDataExplorerTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Table             interface{}                               `pulumi:"table"`
	Type              string                                    `pulumi:"type"`
}

type AzureDataLakeAnalyticsLinkedService struct {
	AccountName          interface{}                       `pulumi:"accountName"`
	Annotations          []interface{}                     `pulumi:"annotations"`
	ConnectVia           *IntegrationRuntimeReference      `pulumi:"connectVia"`
	DataLakeAnalyticsUri interface{}                       `pulumi:"dataLakeAnalyticsUri"`
	Description          *string                           `pulumi:"description"`
	EncryptedCredential  interface{}                       `pulumi:"encryptedCredential"`
	Parameters           map[string]ParameterSpecification `pulumi:"parameters"`
	ResourceGroupName    interface{}                       `pulumi:"resourceGroupName"`
	ServicePrincipalId   interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey  interface{}                       `pulumi:"servicePrincipalKey"`
	SubscriptionId       interface{}                       `pulumi:"subscriptionId"`
	Tenant               interface{}                       `pulumi:"tenant"`
	Type                 string                            `pulumi:"type"`
}

type AzureDataLakeAnalyticsLinkedServiceResponse struct {
	AccountName          interface{}                               `pulumi:"accountName"`
	Annotations          []interface{}                             `pulumi:"annotations"`
	ConnectVia           *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	DataLakeAnalyticsUri interface{}                               `pulumi:"dataLakeAnalyticsUri"`
	Description          *string                                   `pulumi:"description"`
	EncryptedCredential  interface{}                               `pulumi:"encryptedCredential"`
	Parameters           map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	ResourceGroupName    interface{}                               `pulumi:"resourceGroupName"`
	ServicePrincipalId   interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey  interface{}                               `pulumi:"servicePrincipalKey"`
	SubscriptionId       interface{}                               `pulumi:"subscriptionId"`
	Tenant               interface{}                               `pulumi:"tenant"`
	Type                 string                                    `pulumi:"type"`
}

type AzureDataLakeStoreDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Compression       *DatasetCompression               `pulumi:"compression"`
	Description       *string                           `pulumi:"description"`
	FileName          interface{}                       `pulumi:"fileName"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	FolderPath        interface{}                       `pulumi:"folderPath"`
	Format            interface{}                       `pulumi:"format"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type AzureDataLakeStoreDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Compression       *DatasetCompressionResponse               `pulumi:"compression"`
	Description       *string                                   `pulumi:"description"`
	FileName          interface{}                               `pulumi:"fileName"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	FolderPath        interface{}                               `pulumi:"folderPath"`
	Format            interface{}                               `pulumi:"format"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type AzureDataLakeStoreLinkedService struct {
	AccountName         interface{}                       `pulumi:"accountName"`
	Annotations         []interface{}                     `pulumi:"annotations"`
	AzureCloudType      interface{}                       `pulumi:"azureCloudType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Credential          *CredentialReference              `pulumi:"credential"`
	DataLakeStoreUri    interface{}                       `pulumi:"dataLakeStoreUri"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	ResourceGroupName   interface{}                       `pulumi:"resourceGroupName"`
	ServicePrincipalId  interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey interface{}                       `pulumi:"servicePrincipalKey"`
	SubscriptionId      interface{}                       `pulumi:"subscriptionId"`
	Tenant              interface{}                       `pulumi:"tenant"`
	Type                string                            `pulumi:"type"`
}

type AzureDataLakeStoreLinkedServiceResponse struct {
	AccountName         interface{}                               `pulumi:"accountName"`
	Annotations         []interface{}                             `pulumi:"annotations"`
	AzureCloudType      interface{}                               `pulumi:"azureCloudType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Credential          *CredentialReferenceResponse              `pulumi:"credential"`
	DataLakeStoreUri    interface{}                               `pulumi:"dataLakeStoreUri"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	ResourceGroupName   interface{}                               `pulumi:"resourceGroupName"`
	ServicePrincipalId  interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey interface{}                               `pulumi:"servicePrincipalKey"`
	SubscriptionId      interface{}                               `pulumi:"subscriptionId"`
	Tenant              interface{}                               `pulumi:"tenant"`
	Type                string                                    `pulumi:"type"`
}

type AzureDataLakeStoreLocation struct {
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
}

type AzureDataLakeStoreLocationResponse struct {
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
}

type AzureDataLakeStoreReadSettings struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	ListAfter                  interface{} `pulumi:"listAfter"`
	ListBefore                 interface{} `pulumi:"listBefore"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type AzureDataLakeStoreReadSettingsResponse struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	ListAfter                  interface{} `pulumi:"listAfter"`
	ListBefore                 interface{} `pulumi:"listBefore"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type AzureDataLakeStoreSink struct {
	CopyBehavior                 interface{} `pulumi:"copyBehavior"`
	DisableMetricsCollection     interface{} `pulumi:"disableMetricsCollection"`
	EnableAdlsSingleFileParallel interface{} `pulumi:"enableAdlsSingleFileParallel"`
	MaxConcurrentConnections     interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount               interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait                interface{} `pulumi:"sinkRetryWait"`
	Type                         string      `pulumi:"type"`
	WriteBatchSize               interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout            interface{} `pulumi:"writeBatchTimeout"`
}

type AzureDataLakeStoreSinkResponse struct {
	CopyBehavior                 interface{} `pulumi:"copyBehavior"`
	DisableMetricsCollection     interface{} `pulumi:"disableMetricsCollection"`
	EnableAdlsSingleFileParallel interface{} `pulumi:"enableAdlsSingleFileParallel"`
	MaxConcurrentConnections     interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount               interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait                interface{} `pulumi:"sinkRetryWait"`
	Type                         string      `pulumi:"type"`
	WriteBatchSize               interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout            interface{} `pulumi:"writeBatchTimeout"`
}

type AzureDataLakeStoreSource struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Recursive                interface{} `pulumi:"recursive"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type AzureDataLakeStoreSourceResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Recursive                interface{} `pulumi:"recursive"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type AzureDataLakeStoreWriteSettings struct {
	CopyBehavior             interface{} `pulumi:"copyBehavior"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	ExpiryDateTime           interface{} `pulumi:"expiryDateTime"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Type                     string      `pulumi:"type"`
}

type AzureDataLakeStoreWriteSettingsResponse struct {
	CopyBehavior             interface{} `pulumi:"copyBehavior"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	ExpiryDateTime           interface{} `pulumi:"expiryDateTime"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Type                     string      `pulumi:"type"`
}

type AzureDatabricksDeltaLakeDataset struct {
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

type AzureDatabricksDeltaLakeDatasetResponse struct {
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

type AzureDatabricksDeltaLakeExportCommand struct {
	DateFormat      interface{} `pulumi:"dateFormat"`
	TimestampFormat interface{} `pulumi:"timestampFormat"`
	Type            string      `pulumi:"type"`
}

type AzureDatabricksDeltaLakeExportCommandResponse struct {
	DateFormat      interface{} `pulumi:"dateFormat"`
	TimestampFormat interface{} `pulumi:"timestampFormat"`
	Type            string      `pulumi:"type"`
}

type AzureDatabricksDeltaLakeImportCommand struct {
	DateFormat      interface{} `pulumi:"dateFormat"`
	TimestampFormat interface{} `pulumi:"timestampFormat"`
	Type            string      `pulumi:"type"`
}

type AzureDatabricksDeltaLakeImportCommandResponse struct {
	DateFormat      interface{} `pulumi:"dateFormat"`
	TimestampFormat interface{} `pulumi:"timestampFormat"`
	Type            string      `pulumi:"type"`
}

type AzureDatabricksDeltaLakeLinkedService struct {
	AccessToken         interface{}                       `pulumi:"accessToken"`
	Annotations         []interface{}                     `pulumi:"annotations"`
	ClusterId           interface{}                       `pulumi:"clusterId"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Credential          *CredentialReference              `pulumi:"credential"`
	Description         *string                           `pulumi:"description"`
	Domain              interface{}                       `pulumi:"domain"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Type                string                            `pulumi:"type"`
	WorkspaceResourceId interface{}                       `pulumi:"workspaceResourceId"`
}

type AzureDatabricksDeltaLakeLinkedServiceResponse struct {
	AccessToken         interface{}                               `pulumi:"accessToken"`
	Annotations         []interface{}                             `pulumi:"annotations"`
	ClusterId           interface{}                               `pulumi:"clusterId"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Credential          *CredentialReferenceResponse              `pulumi:"credential"`
	Description         *string                                   `pulumi:"description"`
	Domain              interface{}                               `pulumi:"domain"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
	WorkspaceResourceId interface{}                               `pulumi:"workspaceResourceId"`
}

type AzureDatabricksDeltaLakeSink struct {
	DisableMetricsCollection interface{}                            `pulumi:"disableMetricsCollection"`
	ImportSettings           *AzureDatabricksDeltaLakeImportCommand `pulumi:"importSettings"`
	MaxConcurrentConnections interface{}                            `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{}                            `pulumi:"preCopyScript"`
	SinkRetryCount           interface{}                            `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}                            `pulumi:"sinkRetryWait"`
	Type                     string                                 `pulumi:"type"`
	WriteBatchSize           interface{}                            `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}                            `pulumi:"writeBatchTimeout"`
}

type AzureDatabricksDeltaLakeSinkResponse struct {
	DisableMetricsCollection interface{}                                    `pulumi:"disableMetricsCollection"`
	ImportSettings           *AzureDatabricksDeltaLakeImportCommandResponse `pulumi:"importSettings"`
	MaxConcurrentConnections interface{}                                    `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{}                                    `pulumi:"preCopyScript"`
	SinkRetryCount           interface{}                                    `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}                                    `pulumi:"sinkRetryWait"`
	Type                     string                                         `pulumi:"type"`
	WriteBatchSize           interface{}                                    `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}                                    `pulumi:"writeBatchTimeout"`
}

type AzureDatabricksDeltaLakeSource struct {
	DisableMetricsCollection interface{}                            `pulumi:"disableMetricsCollection"`
	ExportSettings           *AzureDatabricksDeltaLakeExportCommand `pulumi:"exportSettings"`
	MaxConcurrentConnections interface{}                            `pulumi:"maxConcurrentConnections"`
	Query                    interface{}                            `pulumi:"query"`
	SourceRetryCount         interface{}                            `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                            `pulumi:"sourceRetryWait"`
	Type                     string                                 `pulumi:"type"`
}

type AzureDatabricksDeltaLakeSourceResponse struct {
	DisableMetricsCollection interface{}                                    `pulumi:"disableMetricsCollection"`
	ExportSettings           *AzureDatabricksDeltaLakeExportCommandResponse `pulumi:"exportSettings"`
	MaxConcurrentConnections interface{}                                    `pulumi:"maxConcurrentConnections"`
	Query                    interface{}                                    `pulumi:"query"`
	SourceRetryCount         interface{}                                    `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                                    `pulumi:"sourceRetryWait"`
	Type                     string                                         `pulumi:"type"`
}

type AzureDatabricksLinkedService struct {
	AccessToken                 interface{}                       `pulumi:"accessToken"`
	Annotations                 []interface{}                     `pulumi:"annotations"`
	Authentication              interface{}                       `pulumi:"authentication"`
	ConnectVia                  *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Credential                  *CredentialReference              `pulumi:"credential"`
	Description                 *string                           `pulumi:"description"`
	Domain                      interface{}                       `pulumi:"domain"`
	EncryptedCredential         interface{}                       `pulumi:"encryptedCredential"`
	ExistingClusterId           interface{}                       `pulumi:"existingClusterId"`
	InstancePoolId              interface{}                       `pulumi:"instancePoolId"`
	NewClusterCustomTags        map[string]interface{}            `pulumi:"newClusterCustomTags"`
	NewClusterDriverNodeType    interface{}                       `pulumi:"newClusterDriverNodeType"`
	NewClusterEnableElasticDisk interface{}                       `pulumi:"newClusterEnableElasticDisk"`
	NewClusterInitScripts       interface{}                       `pulumi:"newClusterInitScripts"`
	NewClusterLogDestination    interface{}                       `pulumi:"newClusterLogDestination"`
	NewClusterNodeType          interface{}                       `pulumi:"newClusterNodeType"`
	NewClusterNumOfWorker       interface{}                       `pulumi:"newClusterNumOfWorker"`
	NewClusterSparkConf         map[string]interface{}            `pulumi:"newClusterSparkConf"`
	NewClusterSparkEnvVars      map[string]interface{}            `pulumi:"newClusterSparkEnvVars"`
	NewClusterVersion           interface{}                       `pulumi:"newClusterVersion"`
	Parameters                  map[string]ParameterSpecification `pulumi:"parameters"`
	PolicyId                    interface{}                       `pulumi:"policyId"`
	Type                        string                            `pulumi:"type"`
	WorkspaceResourceId         interface{}                       `pulumi:"workspaceResourceId"`
}

type AzureDatabricksLinkedServiceResponse struct {
	AccessToken                 interface{}                               `pulumi:"accessToken"`
	Annotations                 []interface{}                             `pulumi:"annotations"`
	Authentication              interface{}                               `pulumi:"authentication"`
	ConnectVia                  *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Credential                  *CredentialReferenceResponse              `pulumi:"credential"`
	Description                 *string                                   `pulumi:"description"`
	Domain                      interface{}                               `pulumi:"domain"`
	EncryptedCredential         interface{}                               `pulumi:"encryptedCredential"`
	ExistingClusterId           interface{}                               `pulumi:"existingClusterId"`
	InstancePoolId              interface{}                               `pulumi:"instancePoolId"`
	NewClusterCustomTags        map[string]interface{}                    `pulumi:"newClusterCustomTags"`
	NewClusterDriverNodeType    interface{}                               `pulumi:"newClusterDriverNodeType"`
	NewClusterEnableElasticDisk interface{}                               `pulumi:"newClusterEnableElasticDisk"`
	NewClusterInitScripts       interface{}                               `pulumi:"newClusterInitScripts"`
	NewClusterLogDestination    interface{}                               `pulumi:"newClusterLogDestination"`
	NewClusterNodeType          interface{}                               `pulumi:"newClusterNodeType"`
	NewClusterNumOfWorker       interface{}                               `pulumi:"newClusterNumOfWorker"`
	NewClusterSparkConf         map[string]interface{}                    `pulumi:"newClusterSparkConf"`
	NewClusterSparkEnvVars      map[string]interface{}                    `pulumi:"newClusterSparkEnvVars"`
	NewClusterVersion           interface{}                               `pulumi:"newClusterVersion"`
	Parameters                  map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	PolicyId                    interface{}                               `pulumi:"policyId"`
	Type                        string                                    `pulumi:"type"`
	WorkspaceResourceId         interface{}                               `pulumi:"workspaceResourceId"`
}

type AzureFileStorageLinkedService struct {
	AccountKey          *AzureKeyVaultSecretReference     `pulumi:"accountKey"`
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	FileShare           interface{}                       `pulumi:"fileShare"`
	Host                interface{}                       `pulumi:"host"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	SasToken            *AzureKeyVaultSecretReference     `pulumi:"sasToken"`
	SasUri              interface{}                       `pulumi:"sasUri"`
	Snapshot            interface{}                       `pulumi:"snapshot"`
	Type                string                            `pulumi:"type"`
	UserId              interface{}                       `pulumi:"userId"`
}

type AzureFileStorageLinkedServiceResponse struct {
	AccountKey          *AzureKeyVaultSecretReferenceResponse     `pulumi:"accountKey"`
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	FileShare           interface{}                               `pulumi:"fileShare"`
	Host                interface{}                               `pulumi:"host"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	SasToken            *AzureKeyVaultSecretReferenceResponse     `pulumi:"sasToken"`
	SasUri              interface{}                               `pulumi:"sasUri"`
	Snapshot            interface{}                               `pulumi:"snapshot"`
	Type                string                                    `pulumi:"type"`
	UserId              interface{}                               `pulumi:"userId"`
}

type AzureFileStorageLocation struct {
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
}

type AzureFileStorageLocationResponse struct {
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
}

type AzureFileStorageReadSettings struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Prefix                     interface{} `pulumi:"prefix"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type AzureFileStorageReadSettingsResponse struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Prefix                     interface{} `pulumi:"prefix"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type AzureFileStorageWriteSettings struct {
	CopyBehavior             interface{} `pulumi:"copyBehavior"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Type                     string      `pulumi:"type"`
}

type AzureFileStorageWriteSettingsResponse struct {
	CopyBehavior             interface{} `pulumi:"copyBehavior"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Type                     string      `pulumi:"type"`
}

type AzureFunctionActivity struct {
	Body              interface{}             `pulumi:"body"`
	DependsOn         []ActivityDependency    `pulumi:"dependsOn"`
	Description       *string                 `pulumi:"description"`
	FunctionName      interface{}             `pulumi:"functionName"`
	Headers           interface{}             `pulumi:"headers"`
	LinkedServiceName *LinkedServiceReference `pulumi:"linkedServiceName"`
	Method            string                  `pulumi:"method"`
	Name              string                  `pulumi:"name"`
	Policy            *ActivityPolicy         `pulumi:"policy"`
	Type              string                  `pulumi:"type"`
	UserProperties    []UserProperty          `pulumi:"userProperties"`
}

type AzureFunctionActivityResponse struct {
	Body              interface{}                     `pulumi:"body"`
	DependsOn         []ActivityDependencyResponse    `pulumi:"dependsOn"`
	Description       *string                         `pulumi:"description"`
	FunctionName      interface{}                     `pulumi:"functionName"`
	Headers           interface{}                     `pulumi:"headers"`
	LinkedServiceName *LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	Method            string                          `pulumi:"method"`
	Name              string                          `pulumi:"name"`
	Policy            *ActivityPolicyResponse         `pulumi:"policy"`
	Type              string                          `pulumi:"type"`
	UserProperties    []UserPropertyResponse          `pulumi:"userProperties"`
}

type AzureFunctionLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	Authentication      interface{}                       `pulumi:"authentication"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Credential          *CredentialReference              `pulumi:"credential"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	FunctionAppUrl      interface{}                       `pulumi:"functionAppUrl"`
	FunctionKey         interface{}                       `pulumi:"functionKey"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	ResourceId          interface{}                       `pulumi:"resourceId"`
	Type                string                            `pulumi:"type"`
}

type AzureFunctionLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	Authentication      interface{}                               `pulumi:"authentication"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Credential          *CredentialReferenceResponse              `pulumi:"credential"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	FunctionAppUrl      interface{}                               `pulumi:"functionAppUrl"`
	FunctionKey         interface{}                               `pulumi:"functionKey"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	ResourceId          interface{}                               `pulumi:"resourceId"`
	Type                string                                    `pulumi:"type"`
}

type AzureKeyVaultLinkedService struct {
	Annotations []interface{}                     `pulumi:"annotations"`
	BaseUrl     interface{}                       `pulumi:"baseUrl"`
	ConnectVia  *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Credential  *CredentialReference              `pulumi:"credential"`
	Description *string                           `pulumi:"description"`
	Parameters  map[string]ParameterSpecification `pulumi:"parameters"`
	Type        string                            `pulumi:"type"`
}

type AzureKeyVaultLinkedServiceResponse struct {
	Annotations []interface{}                             `pulumi:"annotations"`
	BaseUrl     interface{}                               `pulumi:"baseUrl"`
	ConnectVia  *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Credential  *CredentialReferenceResponse              `pulumi:"credential"`
	Description *string                                   `pulumi:"description"`
	Parameters  map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type        string                                    `pulumi:"type"`
}

type AzureKeyVaultSecretReference struct {
	SecretName    interface{}            `pulumi:"secretName"`
	SecretVersion interface{}            `pulumi:"secretVersion"`
	Store         LinkedServiceReference `pulumi:"store"`
	Type          string                 `pulumi:"type"`
}

type AzureKeyVaultSecretReferenceResponse struct {
	SecretName    interface{}                    `pulumi:"secretName"`
	SecretVersion interface{}                    `pulumi:"secretVersion"`
	Store         LinkedServiceReferenceResponse `pulumi:"store"`
	Type          string                         `pulumi:"type"`
}

type AzureMLBatchExecutionActivity struct {
	DependsOn         []ActivityDependency             `pulumi:"dependsOn"`
	Description       *string                          `pulumi:"description"`
	GlobalParameters  map[string]interface{}           `pulumi:"globalParameters"`
	LinkedServiceName *LinkedServiceReference          `pulumi:"linkedServiceName"`
	Name              string                           `pulumi:"name"`
	Policy            *ActivityPolicy                  `pulumi:"policy"`
	Type              string                           `pulumi:"type"`
	UserProperties    []UserProperty                   `pulumi:"userProperties"`
	WebServiceInputs  map[string]AzureMLWebServiceFile `pulumi:"webServiceInputs"`
	WebServiceOutputs map[string]AzureMLWebServiceFile `pulumi:"webServiceOutputs"`
}

type AzureMLBatchExecutionActivityResponse struct {
	DependsOn         []ActivityDependencyResponse             `pulumi:"dependsOn"`
	Description       *string                                  `pulumi:"description"`
	GlobalParameters  map[string]interface{}                   `pulumi:"globalParameters"`
	LinkedServiceName *LinkedServiceReferenceResponse          `pulumi:"linkedServiceName"`
	Name              string                                   `pulumi:"name"`
	Policy            *ActivityPolicyResponse                  `pulumi:"policy"`
	Type              string                                   `pulumi:"type"`
	UserProperties    []UserPropertyResponse                   `pulumi:"userProperties"`
	WebServiceInputs  map[string]AzureMLWebServiceFileResponse `pulumi:"webServiceInputs"`
	WebServiceOutputs map[string]AzureMLWebServiceFileResponse `pulumi:"webServiceOutputs"`
}

type AzureMLExecutePipelineActivity struct {
	ContinueOnStepFailure interface{}             `pulumi:"continueOnStepFailure"`
	DataPathAssignments   interface{}             `pulumi:"dataPathAssignments"`
	DependsOn             []ActivityDependency    `pulumi:"dependsOn"`
	Description           *string                 `pulumi:"description"`
	ExperimentName        interface{}             `pulumi:"experimentName"`
	LinkedServiceName     *LinkedServiceReference `pulumi:"linkedServiceName"`
	MlParentRunId         interface{}             `pulumi:"mlParentRunId"`
	MlPipelineEndpointId  interface{}             `pulumi:"mlPipelineEndpointId"`
	MlPipelineId          interface{}             `pulumi:"mlPipelineId"`
	MlPipelineParameters  interface{}             `pulumi:"mlPipelineParameters"`
	Name                  string                  `pulumi:"name"`
	Policy                *ActivityPolicy         `pulumi:"policy"`
	Type                  string                  `pulumi:"type"`
	UserProperties        []UserProperty          `pulumi:"userProperties"`
	Version               interface{}             `pulumi:"version"`
}

type AzureMLExecutePipelineActivityResponse struct {
	ContinueOnStepFailure interface{}                     `pulumi:"continueOnStepFailure"`
	DataPathAssignments   interface{}                     `pulumi:"dataPathAssignments"`
	DependsOn             []ActivityDependencyResponse    `pulumi:"dependsOn"`
	Description           *string                         `pulumi:"description"`
	ExperimentName        interface{}                     `pulumi:"experimentName"`
	LinkedServiceName     *LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	MlParentRunId         interface{}                     `pulumi:"mlParentRunId"`
	MlPipelineEndpointId  interface{}                     `pulumi:"mlPipelineEndpointId"`
	MlPipelineId          interface{}                     `pulumi:"mlPipelineId"`
	MlPipelineParameters  interface{}                     `pulumi:"mlPipelineParameters"`
	Name                  string                          `pulumi:"name"`
	Policy                *ActivityPolicyResponse         `pulumi:"policy"`
	Type                  string                          `pulumi:"type"`
	UserProperties        []UserPropertyResponse          `pulumi:"userProperties"`
	Version               interface{}                     `pulumi:"version"`
}

type AzureMLLinkedService struct {
	Annotations            []interface{}                     `pulumi:"annotations"`
	ApiKey                 interface{}                       `pulumi:"apiKey"`
	Authentication         interface{}                       `pulumi:"authentication"`
	ConnectVia             *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description            *string                           `pulumi:"description"`
	EncryptedCredential    interface{}                       `pulumi:"encryptedCredential"`
	MlEndpoint             interface{}                       `pulumi:"mlEndpoint"`
	Parameters             map[string]ParameterSpecification `pulumi:"parameters"`
	ServicePrincipalId     interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey    interface{}                       `pulumi:"servicePrincipalKey"`
	Tenant                 interface{}                       `pulumi:"tenant"`
	Type                   string                            `pulumi:"type"`
	UpdateResourceEndpoint interface{}                       `pulumi:"updateResourceEndpoint"`
}

type AzureMLLinkedServiceResponse struct {
	Annotations            []interface{}                             `pulumi:"annotations"`
	ApiKey                 interface{}                               `pulumi:"apiKey"`
	Authentication         interface{}                               `pulumi:"authentication"`
	ConnectVia             *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description            *string                                   `pulumi:"description"`
	EncryptedCredential    interface{}                               `pulumi:"encryptedCredential"`
	MlEndpoint             interface{}                               `pulumi:"mlEndpoint"`
	Parameters             map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	ServicePrincipalId     interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey    interface{}                               `pulumi:"servicePrincipalKey"`
	Tenant                 interface{}                               `pulumi:"tenant"`
	Type                   string                                    `pulumi:"type"`
	UpdateResourceEndpoint interface{}                               `pulumi:"updateResourceEndpoint"`
}

type AzureMLServiceLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	MlWorkspaceName     interface{}                       `pulumi:"mlWorkspaceName"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	ResourceGroupName   interface{}                       `pulumi:"resourceGroupName"`
	ServicePrincipalId  interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey interface{}                       `pulumi:"servicePrincipalKey"`
	SubscriptionId      interface{}                       `pulumi:"subscriptionId"`
	Tenant              interface{}                       `pulumi:"tenant"`
	Type                string                            `pulumi:"type"`
}

type AzureMLServiceLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	MlWorkspaceName     interface{}                               `pulumi:"mlWorkspaceName"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	ResourceGroupName   interface{}                               `pulumi:"resourceGroupName"`
	ServicePrincipalId  interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey interface{}                               `pulumi:"servicePrincipalKey"`
	SubscriptionId      interface{}                               `pulumi:"subscriptionId"`
	Tenant              interface{}                               `pulumi:"tenant"`
	Type                string                                    `pulumi:"type"`
}

type AzureMLUpdateResourceActivity struct {
	DependsOn                     []ActivityDependency    `pulumi:"dependsOn"`
	Description                   *string                 `pulumi:"description"`
	LinkedServiceName             *LinkedServiceReference `pulumi:"linkedServiceName"`
	Name                          string                  `pulumi:"name"`
	Policy                        *ActivityPolicy         `pulumi:"policy"`
	TrainedModelFilePath          interface{}             `pulumi:"trainedModelFilePath"`
	TrainedModelLinkedServiceName LinkedServiceReference  `pulumi:"trainedModelLinkedServiceName"`
	TrainedModelName              interface{}             `pulumi:"trainedModelName"`
	Type                          string                  `pulumi:"type"`
	UserProperties                []UserProperty          `pulumi:"userProperties"`
}

type AzureMLUpdateResourceActivityResponse struct {
	DependsOn                     []ActivityDependencyResponse    `pulumi:"dependsOn"`
	Description                   *string                         `pulumi:"description"`
	LinkedServiceName             *LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	Name                          string                          `pulumi:"name"`
	Policy                        *ActivityPolicyResponse         `pulumi:"policy"`
	TrainedModelFilePath          interface{}                     `pulumi:"trainedModelFilePath"`
	TrainedModelLinkedServiceName LinkedServiceReferenceResponse  `pulumi:"trainedModelLinkedServiceName"`
	TrainedModelName              interface{}                     `pulumi:"trainedModelName"`
	Type                          string                          `pulumi:"type"`
	UserProperties                []UserPropertyResponse          `pulumi:"userProperties"`
}

type AzureMLWebServiceFile struct {
	FilePath          interface{}            `pulumi:"filePath"`
	LinkedServiceName LinkedServiceReference `pulumi:"linkedServiceName"`
}

type AzureMLWebServiceFileResponse struct {
	FilePath          interface{}                    `pulumi:"filePath"`
	LinkedServiceName LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
}

type AzureMariaDBLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Pwd                 *AzureKeyVaultSecretReference     `pulumi:"pwd"`
	Type                string                            `pulumi:"type"`
}

type AzureMariaDBLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Pwd                 *AzureKeyVaultSecretReferenceResponse     `pulumi:"pwd"`
	Type                string                                    `pulumi:"type"`
}

type AzureMariaDBSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type AzureMariaDBSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type AzureMariaDBTableDataset struct {
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

type AzureMariaDBTableDatasetResponse struct {
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

type AzureMySqlLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReference     `pulumi:"password"`
	Type                string                            `pulumi:"type"`
}

type AzureMySqlLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReferenceResponse     `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
}

type AzureMySqlSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{} `pulumi:"preCopyScript"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type AzureMySqlSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{} `pulumi:"preCopyScript"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type AzureMySqlSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type AzureMySqlSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type AzureMySqlTableDataset struct {
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

type AzureMySqlTableDatasetResponse struct {
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

type AzurePostgreSqlLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReference     `pulumi:"password"`
	Type                string                            `pulumi:"type"`
}

type AzurePostgreSqlLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReferenceResponse     `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
}

type AzurePostgreSqlSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{} `pulumi:"preCopyScript"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type AzurePostgreSqlSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{} `pulumi:"preCopyScript"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type AzurePostgreSqlSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type AzurePostgreSqlSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type AzurePostgreSqlTableDataset struct {
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

type AzurePostgreSqlTableDatasetResponse struct {
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

type AzureQueueSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type AzureQueueSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type AzureSearchIndexDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	IndexName         interface{}                       `pulumi:"indexName"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type AzureSearchIndexDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	IndexName         interface{}                               `pulumi:"indexName"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type AzureSearchIndexSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            *string     `pulumi:"writeBehavior"`
}

type AzureSearchIndexSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            *string     `pulumi:"writeBehavior"`
}

type AzureSearchLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Key                 interface{}                       `pulumi:"key"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Type                string                            `pulumi:"type"`
	Url                 interface{}                       `pulumi:"url"`
}

type AzureSearchLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Key                 interface{}                               `pulumi:"key"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
	Url                 interface{}                               `pulumi:"url"`
}

type AzureSqlDWLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	AzureCloudType      interface{}                       `pulumi:"azureCloudType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Credential          *CredentialReference              `pulumi:"credential"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReference     `pulumi:"password"`
	ServicePrincipalId  interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey interface{}                       `pulumi:"servicePrincipalKey"`
	Tenant              interface{}                       `pulumi:"tenant"`
	Type                string                            `pulumi:"type"`
}

type AzureSqlDWLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	AzureCloudType      interface{}                               `pulumi:"azureCloudType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Credential          *CredentialReferenceResponse              `pulumi:"credential"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReferenceResponse     `pulumi:"password"`
	ServicePrincipalId  interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey interface{}                               `pulumi:"servicePrincipalKey"`
	Tenant              interface{}                               `pulumi:"tenant"`
	Type                string                                    `pulumi:"type"`
}

type AzureSqlDWTableDataset struct {
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

type AzureSqlDWTableDatasetResponse struct {
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

type AzureSqlDatabaseLinkedService struct {
	AlwaysEncryptedSettings *SqlAlwaysEncryptedProperties     `pulumi:"alwaysEncryptedSettings"`
	Annotations             []interface{}                     `pulumi:"annotations"`
	AzureCloudType          interface{}                       `pulumi:"azureCloudType"`
	ConnectVia              *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString        interface{}                       `pulumi:"connectionString"`
	Credential              *CredentialReference              `pulumi:"credential"`
	Description             *string                           `pulumi:"description"`
	EncryptedCredential     interface{}                       `pulumi:"encryptedCredential"`
	Parameters              map[string]ParameterSpecification `pulumi:"parameters"`
	Password                *AzureKeyVaultSecretReference     `pulumi:"password"`
	ServicePrincipalId      interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey     interface{}                       `pulumi:"servicePrincipalKey"`
	Tenant                  interface{}                       `pulumi:"tenant"`
	Type                    string                            `pulumi:"type"`
}

type AzureSqlDatabaseLinkedServiceResponse struct {
	AlwaysEncryptedSettings *SqlAlwaysEncryptedPropertiesResponse     `pulumi:"alwaysEncryptedSettings"`
	Annotations             []interface{}                             `pulumi:"annotations"`
	AzureCloudType          interface{}                               `pulumi:"azureCloudType"`
	ConnectVia              *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString        interface{}                               `pulumi:"connectionString"`
	Credential              *CredentialReferenceResponse              `pulumi:"credential"`
	Description             *string                                   `pulumi:"description"`
	EncryptedCredential     interface{}                               `pulumi:"encryptedCredential"`
	Parameters              map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                *AzureKeyVaultSecretReferenceResponse     `pulumi:"password"`
	ServicePrincipalId      interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey     interface{}                               `pulumi:"servicePrincipalKey"`
	Tenant                  interface{}                               `pulumi:"tenant"`
	Type                    string                                    `pulumi:"type"`
}

type AzureSqlMILinkedService struct {
	AlwaysEncryptedSettings *SqlAlwaysEncryptedProperties     `pulumi:"alwaysEncryptedSettings"`
	Annotations             []interface{}                     `pulumi:"annotations"`
	AzureCloudType          interface{}                       `pulumi:"azureCloudType"`
	ConnectVia              *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString        interface{}                       `pulumi:"connectionString"`
	Credential              *CredentialReference              `pulumi:"credential"`
	Description             *string                           `pulumi:"description"`
	EncryptedCredential     interface{}                       `pulumi:"encryptedCredential"`
	Parameters              map[string]ParameterSpecification `pulumi:"parameters"`
	Password                *AzureKeyVaultSecretReference     `pulumi:"password"`
	ServicePrincipalId      interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey     interface{}                       `pulumi:"servicePrincipalKey"`
	Tenant                  interface{}                       `pulumi:"tenant"`
	Type                    string                            `pulumi:"type"`
}

type AzureSqlMILinkedServiceResponse struct {
	AlwaysEncryptedSettings *SqlAlwaysEncryptedPropertiesResponse     `pulumi:"alwaysEncryptedSettings"`
	Annotations             []interface{}                             `pulumi:"annotations"`
	AzureCloudType          interface{}                               `pulumi:"azureCloudType"`
	ConnectVia              *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString        interface{}                               `pulumi:"connectionString"`
	Credential              *CredentialReferenceResponse              `pulumi:"credential"`
	Description             *string                                   `pulumi:"description"`
	EncryptedCredential     interface{}                               `pulumi:"encryptedCredential"`
	Parameters              map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                *AzureKeyVaultSecretReferenceResponse     `pulumi:"password"`
	ServicePrincipalId      interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey     interface{}                               `pulumi:"servicePrincipalKey"`
	Tenant                  interface{}                               `pulumi:"tenant"`
	Type                    string                                    `pulumi:"type"`
}

type AzureSqlMITableDataset struct {
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

type AzureSqlMITableDatasetResponse struct {
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

type AzureSqlSink struct {
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

type AzureSqlSinkResponse struct {
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

type AzureSqlSource struct {
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

type AzureSqlSourceResponse struct {
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

type AzureSqlTableDataset struct {
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

type AzureSqlTableDatasetResponse struct {
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

type AzureStorageLinkedService struct {
	AccountKey          *AzureKeyVaultSecretReference     `pulumi:"accountKey"`
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential *string                           `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	SasToken            *AzureKeyVaultSecretReference     `pulumi:"sasToken"`
	SasUri              interface{}                       `pulumi:"sasUri"`
	Type                string                            `pulumi:"type"`
}

type AzureStorageLinkedServiceResponse struct {
	AccountKey          *AzureKeyVaultSecretReferenceResponse     `pulumi:"accountKey"`
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential *string                                   `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	SasToken            *AzureKeyVaultSecretReferenceResponse     `pulumi:"sasToken"`
	SasUri              interface{}                               `pulumi:"sasUri"`
	Type                string                                    `pulumi:"type"`
}

type AzureSynapseArtifactsLinkedService struct {
	Annotations    []interface{}                     `pulumi:"annotations"`
	Authentication interface{}                       `pulumi:"authentication"`
	ConnectVia     *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description    *string                           `pulumi:"description"`
	Endpoint       interface{}                       `pulumi:"endpoint"`
	Parameters     map[string]ParameterSpecification `pulumi:"parameters"`
	Type           string                            `pulumi:"type"`
}

type AzureSynapseArtifactsLinkedServiceResponse struct {
	Annotations    []interface{}                             `pulumi:"annotations"`
	Authentication interface{}                               `pulumi:"authentication"`
	ConnectVia     *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description    *string                                   `pulumi:"description"`
	Endpoint       interface{}                               `pulumi:"endpoint"`
	Parameters     map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type           string                                    `pulumi:"type"`
}

type AzureTableDataset struct {
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

type AzureTableDatasetResponse struct {
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

type AzureTableSink struct {
	AzureTableDefaultPartitionKeyValue interface{} `pulumi:"azureTableDefaultPartitionKeyValue"`
	AzureTableInsertType               interface{} `pulumi:"azureTableInsertType"`
	AzureTablePartitionKeyName         interface{} `pulumi:"azureTablePartitionKeyName"`
	AzureTableRowKeyName               interface{} `pulumi:"azureTableRowKeyName"`
	DisableMetricsCollection           interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections           interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount                     interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait                      interface{} `pulumi:"sinkRetryWait"`
	Type                               string      `pulumi:"type"`
	WriteBatchSize                     interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout                  interface{} `pulumi:"writeBatchTimeout"`
}

type AzureTableSinkResponse struct {
	AzureTableDefaultPartitionKeyValue interface{} `pulumi:"azureTableDefaultPartitionKeyValue"`
	AzureTableInsertType               interface{} `pulumi:"azureTableInsertType"`
	AzureTablePartitionKeyName         interface{} `pulumi:"azureTablePartitionKeyName"`
	AzureTableRowKeyName               interface{} `pulumi:"azureTableRowKeyName"`
	DisableMetricsCollection           interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections           interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount                     interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait                      interface{} `pulumi:"sinkRetryWait"`
	Type                               string      `pulumi:"type"`
	WriteBatchSize                     interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout                  interface{} `pulumi:"writeBatchTimeout"`
}

type AzureTableSource struct {
	AdditionalColumns                   interface{} `pulumi:"additionalColumns"`
	AzureTableSourceIgnoreTableNotFound interface{} `pulumi:"azureTableSourceIgnoreTableNotFound"`
	AzureTableSourceQuery               interface{} `pulumi:"azureTableSourceQuery"`
	DisableMetricsCollection            interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections            interface{} `pulumi:"maxConcurrentConnections"`
	QueryTimeout                        interface{} `pulumi:"queryTimeout"`
	SourceRetryCount                    interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait                     interface{} `pulumi:"sourceRetryWait"`
	Type                                string      `pulumi:"type"`
}

type AzureTableSourceResponse struct {
	AdditionalColumns                   interface{} `pulumi:"additionalColumns"`
	AzureTableSourceIgnoreTableNotFound interface{} `pulumi:"azureTableSourceIgnoreTableNotFound"`
	AzureTableSourceQuery               interface{} `pulumi:"azureTableSourceQuery"`
	DisableMetricsCollection            interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections            interface{} `pulumi:"maxConcurrentConnections"`
	QueryTimeout                        interface{} `pulumi:"queryTimeout"`
	SourceRetryCount                    interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait                     interface{} `pulumi:"sourceRetryWait"`
	Type                                string      `pulumi:"type"`
}

type AzureTableStorageLinkedService struct {
	AccountKey          *AzureKeyVaultSecretReference     `pulumi:"accountKey"`
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential *string                           `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	SasToken            *AzureKeyVaultSecretReference     `pulumi:"sasToken"`
	SasUri              interface{}                       `pulumi:"sasUri"`
	Type                string                            `pulumi:"type"`
}

type AzureTableStorageLinkedServiceResponse struct {
	AccountKey          *AzureKeyVaultSecretReferenceResponse     `pulumi:"accountKey"`
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential *string                                   `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	SasToken            *AzureKeyVaultSecretReferenceResponse     `pulumi:"sasToken"`
	SasUri              interface{}                               `pulumi:"sasUri"`
	Type                string                                    `pulumi:"type"`
}

type BigDataPoolParametrizationReference struct {
	ReferenceName interface{} `pulumi:"referenceName"`
	Type          string      `pulumi:"type"`
}

type BigDataPoolParametrizationReferenceResponse struct {
	ReferenceName interface{} `pulumi:"referenceName"`
	Type          string      `pulumi:"type"`
}

type BinaryDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Compression       *DatasetCompression               `pulumi:"compression"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Location          interface{}                       `pulumi:"location"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type BinaryDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Compression       *DatasetCompressionResponse               `pulumi:"compression"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Location          interface{}                               `pulumi:"location"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type BinaryReadSettings struct {
	CompressionProperties interface{} `pulumi:"compressionProperties"`
	Type                  string      `pulumi:"type"`
}

type BinaryReadSettingsResponse struct {
	CompressionProperties interface{} `pulumi:"compressionProperties"`
	Type                  string      `pulumi:"type"`
}

type BinarySink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	StoreSettings            interface{} `pulumi:"storeSettings"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type BinarySinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	StoreSettings            interface{} `pulumi:"storeSettings"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type BinarySource struct {
	DisableMetricsCollection interface{}         `pulumi:"disableMetricsCollection"`
	FormatSettings           *BinaryReadSettings `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}         `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{}         `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}         `pulumi:"sourceRetryWait"`
	StoreSettings            interface{}         `pulumi:"storeSettings"`
	Type                     string              `pulumi:"type"`
}

type BinarySourceResponse struct {
	DisableMetricsCollection interface{}                 `pulumi:"disableMetricsCollection"`
	FormatSettings           *BinaryReadSettingsResponse `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}                 `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{}                 `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                 `pulumi:"sourceRetryWait"`
	StoreSettings            interface{}                 `pulumi:"storeSettings"`
	Type                     string                      `pulumi:"type"`
}

type BlobEventsTrigger struct {
	Annotations        []interface{}              `pulumi:"annotations"`
	BlobPathBeginsWith *string                    `pulumi:"blobPathBeginsWith"`
	BlobPathEndsWith   *string                    `pulumi:"blobPathEndsWith"`
	Description        *string                    `pulumi:"description"`
	Events             []string                   `pulumi:"events"`
	IgnoreEmptyBlobs   *bool                      `pulumi:"ignoreEmptyBlobs"`
	Pipelines          []TriggerPipelineReference `pulumi:"pipelines"`
	Scope              string                     `pulumi:"scope"`
	Type               string                     `pulumi:"type"`
}

type BlobEventsTriggerResponse struct {
	Annotations        []interface{}                      `pulumi:"annotations"`
	BlobPathBeginsWith *string                            `pulumi:"blobPathBeginsWith"`
	BlobPathEndsWith   *string                            `pulumi:"blobPathEndsWith"`
	Description        *string                            `pulumi:"description"`
	Events             []string                           `pulumi:"events"`
	IgnoreEmptyBlobs   *bool                              `pulumi:"ignoreEmptyBlobs"`
	Pipelines          []TriggerPipelineReferenceResponse `pulumi:"pipelines"`
	RuntimeState       string                             `pulumi:"runtimeState"`
	Scope              string                             `pulumi:"scope"`
	Type               string                             `pulumi:"type"`
}

type BlobSink struct {
	BlobWriterAddHeader      interface{}    `pulumi:"blobWriterAddHeader"`
	BlobWriterDateTimeFormat interface{}    `pulumi:"blobWriterDateTimeFormat"`
	BlobWriterOverwriteFiles interface{}    `pulumi:"blobWriterOverwriteFiles"`
	CopyBehavior             interface{}    `pulumi:"copyBehavior"`
	DisableMetricsCollection interface{}    `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}    `pulumi:"maxConcurrentConnections"`
	Metadata                 []MetadataItem `pulumi:"metadata"`
	SinkRetryCount           interface{}    `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}    `pulumi:"sinkRetryWait"`
	Type                     string         `pulumi:"type"`
	WriteBatchSize           interface{}    `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}    `pulumi:"writeBatchTimeout"`
}

type BlobSinkResponse struct {
	BlobWriterAddHeader      interface{}            `pulumi:"blobWriterAddHeader"`
	BlobWriterDateTimeFormat interface{}            `pulumi:"blobWriterDateTimeFormat"`
	BlobWriterOverwriteFiles interface{}            `pulumi:"blobWriterOverwriteFiles"`
	CopyBehavior             interface{}            `pulumi:"copyBehavior"`
	DisableMetricsCollection interface{}            `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}            `pulumi:"maxConcurrentConnections"`
	Metadata                 []MetadataItemResponse `pulumi:"metadata"`
	SinkRetryCount           interface{}            `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}            `pulumi:"sinkRetryWait"`
	Type                     string                 `pulumi:"type"`
	WriteBatchSize           interface{}            `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}            `pulumi:"writeBatchTimeout"`
}

type BlobSource struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Recursive                interface{} `pulumi:"recursive"`
	SkipHeaderLineCount      interface{} `pulumi:"skipHeaderLineCount"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	TreatEmptyAsNull         interface{} `pulumi:"treatEmptyAsNull"`
	Type                     string      `pulumi:"type"`
}

type BlobSourceResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Recursive                interface{} `pulumi:"recursive"`
	SkipHeaderLineCount      interface{} `pulumi:"skipHeaderLineCount"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	TreatEmptyAsNull         interface{} `pulumi:"treatEmptyAsNull"`
	Type                     string      `pulumi:"type"`
}

type BlobTrigger struct {
	Annotations    []interface{}              `pulumi:"annotations"`
	Description    *string                    `pulumi:"description"`
	FolderPath     string                     `pulumi:"folderPath"`
	LinkedService  LinkedServiceReference     `pulumi:"linkedService"`
	MaxConcurrency int                        `pulumi:"maxConcurrency"`
	Pipelines      []TriggerPipelineReference `pulumi:"pipelines"`
	Type           string                     `pulumi:"type"`
}

type BlobTriggerResponse struct {
	Annotations    []interface{}                      `pulumi:"annotations"`
	Description    *string                            `pulumi:"description"`
	FolderPath     string                             `pulumi:"folderPath"`
	LinkedService  LinkedServiceReferenceResponse     `pulumi:"linkedService"`
	MaxConcurrency int                                `pulumi:"maxConcurrency"`
	Pipelines      []TriggerPipelineReferenceResponse `pulumi:"pipelines"`
	RuntimeState   string                             `pulumi:"runtimeState"`
	Type           string                             `pulumi:"type"`
}

type CMKIdentityDefinition struct {
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}





type CMKIdentityDefinitionInput interface {
	pulumi.Input

	ToCMKIdentityDefinitionOutput() CMKIdentityDefinitionOutput
	ToCMKIdentityDefinitionOutputWithContext(context.Context) CMKIdentityDefinitionOutput
}

type CMKIdentityDefinitionArgs struct {
	UserAssignedIdentity pulumi.StringPtrInput `pulumi:"userAssignedIdentity"`
}

func (CMKIdentityDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CMKIdentityDefinition)(nil)).Elem()
}

func (i CMKIdentityDefinitionArgs) ToCMKIdentityDefinitionOutput() CMKIdentityDefinitionOutput {
	return i.ToCMKIdentityDefinitionOutputWithContext(context.Background())
}

func (i CMKIdentityDefinitionArgs) ToCMKIdentityDefinitionOutputWithContext(ctx context.Context) CMKIdentityDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CMKIdentityDefinitionOutput)
}

func (i CMKIdentityDefinitionArgs) ToCMKIdentityDefinitionPtrOutput() CMKIdentityDefinitionPtrOutput {
	return i.ToCMKIdentityDefinitionPtrOutputWithContext(context.Background())
}

func (i CMKIdentityDefinitionArgs) ToCMKIdentityDefinitionPtrOutputWithContext(ctx context.Context) CMKIdentityDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CMKIdentityDefinitionOutput).ToCMKIdentityDefinitionPtrOutputWithContext(ctx)
}









type CMKIdentityDefinitionPtrInput interface {
	pulumi.Input

	ToCMKIdentityDefinitionPtrOutput() CMKIdentityDefinitionPtrOutput
	ToCMKIdentityDefinitionPtrOutputWithContext(context.Context) CMKIdentityDefinitionPtrOutput
}

type cmkidentityDefinitionPtrType CMKIdentityDefinitionArgs

func CMKIdentityDefinitionPtr(v *CMKIdentityDefinitionArgs) CMKIdentityDefinitionPtrInput {
	return (*cmkidentityDefinitionPtrType)(v)
}

func (*cmkidentityDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CMKIdentityDefinition)(nil)).Elem()
}

func (i *cmkidentityDefinitionPtrType) ToCMKIdentityDefinitionPtrOutput() CMKIdentityDefinitionPtrOutput {
	return i.ToCMKIdentityDefinitionPtrOutputWithContext(context.Background())
}

func (i *cmkidentityDefinitionPtrType) ToCMKIdentityDefinitionPtrOutputWithContext(ctx context.Context) CMKIdentityDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CMKIdentityDefinitionPtrOutput)
}

type CMKIdentityDefinitionOutput struct{ *pulumi.OutputState }

func (CMKIdentityDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CMKIdentityDefinition)(nil)).Elem()
}

func (o CMKIdentityDefinitionOutput) ToCMKIdentityDefinitionOutput() CMKIdentityDefinitionOutput {
	return o
}

func (o CMKIdentityDefinitionOutput) ToCMKIdentityDefinitionOutputWithContext(ctx context.Context) CMKIdentityDefinitionOutput {
	return o
}

func (o CMKIdentityDefinitionOutput) ToCMKIdentityDefinitionPtrOutput() CMKIdentityDefinitionPtrOutput {
	return o.ToCMKIdentityDefinitionPtrOutputWithContext(context.Background())
}

func (o CMKIdentityDefinitionOutput) ToCMKIdentityDefinitionPtrOutputWithContext(ctx context.Context) CMKIdentityDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CMKIdentityDefinition) *CMKIdentityDefinition {
		return &v
	}).(CMKIdentityDefinitionPtrOutput)
}

func (o CMKIdentityDefinitionOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CMKIdentityDefinition) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type CMKIdentityDefinitionPtrOutput struct{ *pulumi.OutputState }

func (CMKIdentityDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CMKIdentityDefinition)(nil)).Elem()
}

func (o CMKIdentityDefinitionPtrOutput) ToCMKIdentityDefinitionPtrOutput() CMKIdentityDefinitionPtrOutput {
	return o
}

func (o CMKIdentityDefinitionPtrOutput) ToCMKIdentityDefinitionPtrOutputWithContext(ctx context.Context) CMKIdentityDefinitionPtrOutput {
	return o
}

func (o CMKIdentityDefinitionPtrOutput) Elem() CMKIdentityDefinitionOutput {
	return o.ApplyT(func(v *CMKIdentityDefinition) CMKIdentityDefinition {
		if v != nil {
			return *v
		}
		var ret CMKIdentityDefinition
		return ret
	}).(CMKIdentityDefinitionOutput)
}

func (o CMKIdentityDefinitionPtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CMKIdentityDefinition) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type CMKIdentityDefinitionResponse struct {
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}

type CMKIdentityDefinitionResponseOutput struct{ *pulumi.OutputState }

func (CMKIdentityDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CMKIdentityDefinitionResponse)(nil)).Elem()
}

func (o CMKIdentityDefinitionResponseOutput) ToCMKIdentityDefinitionResponseOutput() CMKIdentityDefinitionResponseOutput {
	return o
}

func (o CMKIdentityDefinitionResponseOutput) ToCMKIdentityDefinitionResponseOutputWithContext(ctx context.Context) CMKIdentityDefinitionResponseOutput {
	return o
}

func (o CMKIdentityDefinitionResponseOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CMKIdentityDefinitionResponse) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type CMKIdentityDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (CMKIdentityDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CMKIdentityDefinitionResponse)(nil)).Elem()
}

func (o CMKIdentityDefinitionResponsePtrOutput) ToCMKIdentityDefinitionResponsePtrOutput() CMKIdentityDefinitionResponsePtrOutput {
	return o
}

func (o CMKIdentityDefinitionResponsePtrOutput) ToCMKIdentityDefinitionResponsePtrOutputWithContext(ctx context.Context) CMKIdentityDefinitionResponsePtrOutput {
	return o
}

func (o CMKIdentityDefinitionResponsePtrOutput) Elem() CMKIdentityDefinitionResponseOutput {
	return o.ApplyT(func(v *CMKIdentityDefinitionResponse) CMKIdentityDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret CMKIdentityDefinitionResponse
		return ret
	}).(CMKIdentityDefinitionResponseOutput)
}

func (o CMKIdentityDefinitionResponsePtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CMKIdentityDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type CassandraLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	AuthenticationType  interface{}                       `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Host                interface{}                       `pulumi:"host"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Port                interface{}                       `pulumi:"port"`
	Type                string                            `pulumi:"type"`
	Username            interface{}                       `pulumi:"username"`
}

type CassandraLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	AuthenticationType  interface{}                               `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Host                interface{}                               `pulumi:"host"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Port                interface{}                               `pulumi:"port"`
	Type                string                                    `pulumi:"type"`
	Username            interface{}                               `pulumi:"username"`
}

type CassandraSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	ConsistencyLevel         *string     `pulumi:"consistencyLevel"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type CassandraSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	ConsistencyLevel         *string     `pulumi:"consistencyLevel"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type CassandraTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	Keyspace          interface{}                       `pulumi:"keyspace"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type CassandraTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	Keyspace          interface{}                               `pulumi:"keyspace"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type ChainingTrigger struct {
	Annotations  []interface{}            `pulumi:"annotations"`
	DependsOn    []PipelineReference      `pulumi:"dependsOn"`
	Description  *string                  `pulumi:"description"`
	Pipeline     TriggerPipelineReference `pulumi:"pipeline"`
	RunDimension string                   `pulumi:"runDimension"`
	Type         string                   `pulumi:"type"`
}

type ChainingTriggerResponse struct {
	Annotations  []interface{}                    `pulumi:"annotations"`
	DependsOn    []PipelineReferenceResponse      `pulumi:"dependsOn"`
	Description  *string                          `pulumi:"description"`
	Pipeline     TriggerPipelineReferenceResponse `pulumi:"pipeline"`
	RunDimension string                           `pulumi:"runDimension"`
	RuntimeState string                           `pulumi:"runtimeState"`
	Type         string                           `pulumi:"type"`
}

type CmdkeySetup struct {
	Password   interface{} `pulumi:"password"`
	TargetName interface{} `pulumi:"targetName"`
	Type       string      `pulumi:"type"`
	UserName   interface{} `pulumi:"userName"`
}

type CmdkeySetupResponse struct {
	Password   interface{} `pulumi:"password"`
	TargetName interface{} `pulumi:"targetName"`
	Type       string      `pulumi:"type"`
	UserName   interface{} `pulumi:"userName"`
}

type CommonDataServiceForAppsEntityDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	EntityName        interface{}                       `pulumi:"entityName"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type CommonDataServiceForAppsEntityDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	EntityName        interface{}                               `pulumi:"entityName"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type CommonDataServiceForAppsLinkedService struct {
	Annotations                    []interface{}                     `pulumi:"annotations"`
	AuthenticationType             interface{}                       `pulumi:"authenticationType"`
	ConnectVia                     *IntegrationRuntimeReference      `pulumi:"connectVia"`
	DeploymentType                 interface{}                       `pulumi:"deploymentType"`
	Description                    *string                           `pulumi:"description"`
	EncryptedCredential            interface{}                       `pulumi:"encryptedCredential"`
	HostName                       interface{}                       `pulumi:"hostName"`
	OrganizationName               interface{}                       `pulumi:"organizationName"`
	Parameters                     map[string]ParameterSpecification `pulumi:"parameters"`
	Password                       interface{}                       `pulumi:"password"`
	Port                           interface{}                       `pulumi:"port"`
	ServicePrincipalCredential     interface{}                       `pulumi:"servicePrincipalCredential"`
	ServicePrincipalCredentialType interface{}                       `pulumi:"servicePrincipalCredentialType"`
	ServicePrincipalId             interface{}                       `pulumi:"servicePrincipalId"`
	ServiceUri                     interface{}                       `pulumi:"serviceUri"`
	Type                           string                            `pulumi:"type"`
	Username                       interface{}                       `pulumi:"username"`
}

type CommonDataServiceForAppsLinkedServiceResponse struct {
	Annotations                    []interface{}                             `pulumi:"annotations"`
	AuthenticationType             interface{}                               `pulumi:"authenticationType"`
	ConnectVia                     *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	DeploymentType                 interface{}                               `pulumi:"deploymentType"`
	Description                    *string                                   `pulumi:"description"`
	EncryptedCredential            interface{}                               `pulumi:"encryptedCredential"`
	HostName                       interface{}                               `pulumi:"hostName"`
	OrganizationName               interface{}                               `pulumi:"organizationName"`
	Parameters                     map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                       interface{}                               `pulumi:"password"`
	Port                           interface{}                               `pulumi:"port"`
	ServicePrincipalCredential     interface{}                               `pulumi:"servicePrincipalCredential"`
	ServicePrincipalCredentialType interface{}                               `pulumi:"servicePrincipalCredentialType"`
	ServicePrincipalId             interface{}                               `pulumi:"servicePrincipalId"`
	ServiceUri                     interface{}                               `pulumi:"serviceUri"`
	Type                           string                                    `pulumi:"type"`
	Username                       interface{}                               `pulumi:"username"`
}

type CommonDataServiceForAppsSink struct {
	AlternateKeyName         interface{} `pulumi:"alternateKeyName"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	IgnoreNullValues         interface{} `pulumi:"ignoreNullValues"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            string      `pulumi:"writeBehavior"`
}

type CommonDataServiceForAppsSinkResponse struct {
	AlternateKeyName         interface{} `pulumi:"alternateKeyName"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	IgnoreNullValues         interface{} `pulumi:"ignoreNullValues"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            string      `pulumi:"writeBehavior"`
}

type CommonDataServiceForAppsSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type CommonDataServiceForAppsSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type ComponentSetup struct {
	ComponentName string      `pulumi:"componentName"`
	LicenseKey    interface{} `pulumi:"licenseKey"`
	Type          string      `pulumi:"type"`
}

type ComponentSetupResponse struct {
	ComponentName string      `pulumi:"componentName"`
	LicenseKey    interface{} `pulumi:"licenseKey"`
	Type          string      `pulumi:"type"`
}

type ConcurLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	ClientId              interface{}                       `pulumi:"clientId"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionProperties  interface{}                       `pulumi:"connectionProperties"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Password              interface{}                       `pulumi:"password"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
	Username              interface{}                       `pulumi:"username"`
}

type ConcurLinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	ClientId              interface{}                               `pulumi:"clientId"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionProperties  interface{}                               `pulumi:"connectionProperties"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password              interface{}                               `pulumi:"password"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
	Username              interface{}                               `pulumi:"username"`
}

type ConcurObjectDataset struct {
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

type ConcurObjectDatasetResponse struct {
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

type ConcurSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type ConcurSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type ConnectionStatePropertiesResponse struct {
	ActionsRequired string `pulumi:"actionsRequired"`
	Description     string `pulumi:"description"`
	Status          string `pulumi:"status"`
}

type ConnectionStatePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ConnectionStatePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStatePropertiesResponse)(nil)).Elem()
}

func (o ConnectionStatePropertiesResponseOutput) ToConnectionStatePropertiesResponseOutput() ConnectionStatePropertiesResponseOutput {
	return o
}

func (o ConnectionStatePropertiesResponseOutput) ToConnectionStatePropertiesResponseOutputWithContext(ctx context.Context) ConnectionStatePropertiesResponseOutput {
	return o
}

func (o ConnectionStatePropertiesResponseOutput) ActionsRequired() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionStatePropertiesResponse) string { return v.ActionsRequired }).(pulumi.StringOutput)
}

func (o ConnectionStatePropertiesResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionStatePropertiesResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o ConnectionStatePropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionStatePropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

type ConnectionStatePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectionStatePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionStatePropertiesResponse)(nil)).Elem()
}

func (o ConnectionStatePropertiesResponsePtrOutput) ToConnectionStatePropertiesResponsePtrOutput() ConnectionStatePropertiesResponsePtrOutput {
	return o
}

func (o ConnectionStatePropertiesResponsePtrOutput) ToConnectionStatePropertiesResponsePtrOutputWithContext(ctx context.Context) ConnectionStatePropertiesResponsePtrOutput {
	return o
}

func (o ConnectionStatePropertiesResponsePtrOutput) Elem() ConnectionStatePropertiesResponseOutput {
	return o.ApplyT(func(v *ConnectionStatePropertiesResponse) ConnectionStatePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ConnectionStatePropertiesResponse
		return ret
	}).(ConnectionStatePropertiesResponseOutput)
}

func (o ConnectionStatePropertiesResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionStatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionStatePropertiesResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionStatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionStatePropertiesResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionStatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type ControlActivity struct {
	DependsOn      []ActivityDependency `pulumi:"dependsOn"`
	Description    *string              `pulumi:"description"`
	Name           string               `pulumi:"name"`
	Type           string               `pulumi:"type"`
	UserProperties []UserProperty       `pulumi:"userProperties"`
}

type ControlActivityResponse struct {
	DependsOn      []ActivityDependencyResponse `pulumi:"dependsOn"`
	Description    *string                      `pulumi:"description"`
	Name           string                       `pulumi:"name"`
	Type           string                       `pulumi:"type"`
	UserProperties []UserPropertyResponse       `pulumi:"userProperties"`
}

type CopyActivity struct {
	DataIntegrationUnits            interface{}                      `pulumi:"dataIntegrationUnits"`
	DependsOn                       []ActivityDependency             `pulumi:"dependsOn"`
	Description                     *string                          `pulumi:"description"`
	EnableSkipIncompatibleRow       interface{}                      `pulumi:"enableSkipIncompatibleRow"`
	EnableStaging                   interface{}                      `pulumi:"enableStaging"`
	Inputs                          []DatasetReference               `pulumi:"inputs"`
	LinkedServiceName               *LinkedServiceReference          `pulumi:"linkedServiceName"`
	LogSettings                     *LogSettings                     `pulumi:"logSettings"`
	LogStorageSettings              *LogStorageSettings              `pulumi:"logStorageSettings"`
	Name                            string                           `pulumi:"name"`
	Outputs                         []DatasetReference               `pulumi:"outputs"`
	ParallelCopies                  interface{}                      `pulumi:"parallelCopies"`
	Policy                          *ActivityPolicy                  `pulumi:"policy"`
	Preserve                        []interface{}                    `pulumi:"preserve"`
	PreserveRules                   []interface{}                    `pulumi:"preserveRules"`
	RedirectIncompatibleRowSettings *RedirectIncompatibleRowSettings `pulumi:"redirectIncompatibleRowSettings"`
	Sink                            interface{}                      `pulumi:"sink"`
	SkipErrorFile                   *SkipErrorFile                   `pulumi:"skipErrorFile"`
	Source                          interface{}                      `pulumi:"source"`
	StagingSettings                 *StagingSettings                 `pulumi:"stagingSettings"`
	Translator                      interface{}                      `pulumi:"translator"`
	Type                            string                           `pulumi:"type"`
	UserProperties                  []UserProperty                   `pulumi:"userProperties"`
	ValidateDataConsistency         interface{}                      `pulumi:"validateDataConsistency"`
}

type CopyActivityLogSettings struct {
	EnableReliableLogging interface{} `pulumi:"enableReliableLogging"`
	LogLevel              interface{} `pulumi:"logLevel"`
}

type CopyActivityLogSettingsResponse struct {
	EnableReliableLogging interface{} `pulumi:"enableReliableLogging"`
	LogLevel              interface{} `pulumi:"logLevel"`
}

type CopyActivityResponse struct {
	DataIntegrationUnits            interface{}                              `pulumi:"dataIntegrationUnits"`
	DependsOn                       []ActivityDependencyResponse             `pulumi:"dependsOn"`
	Description                     *string                                  `pulumi:"description"`
	EnableSkipIncompatibleRow       interface{}                              `pulumi:"enableSkipIncompatibleRow"`
	EnableStaging                   interface{}                              `pulumi:"enableStaging"`
	Inputs                          []DatasetReferenceResponse               `pulumi:"inputs"`
	LinkedServiceName               *LinkedServiceReferenceResponse          `pulumi:"linkedServiceName"`
	LogSettings                     *LogSettingsResponse                     `pulumi:"logSettings"`
	LogStorageSettings              *LogStorageSettingsResponse              `pulumi:"logStorageSettings"`
	Name                            string                                   `pulumi:"name"`
	Outputs                         []DatasetReferenceResponse               `pulumi:"outputs"`
	ParallelCopies                  interface{}                              `pulumi:"parallelCopies"`
	Policy                          *ActivityPolicyResponse                  `pulumi:"policy"`
	Preserve                        []interface{}                            `pulumi:"preserve"`
	PreserveRules                   []interface{}                            `pulumi:"preserveRules"`
	RedirectIncompatibleRowSettings *RedirectIncompatibleRowSettingsResponse `pulumi:"redirectIncompatibleRowSettings"`
	Sink                            interface{}                              `pulumi:"sink"`
	SkipErrorFile                   *SkipErrorFileResponse                   `pulumi:"skipErrorFile"`
	Source                          interface{}                              `pulumi:"source"`
	StagingSettings                 *StagingSettingsResponse                 `pulumi:"stagingSettings"`
	Translator                      interface{}                              `pulumi:"translator"`
	Type                            string                                   `pulumi:"type"`
	UserProperties                  []UserPropertyResponse                   `pulumi:"userProperties"`
	ValidateDataConsistency         interface{}                              `pulumi:"validateDataConsistency"`
}

type CosmosDbLinkedService struct {
	AccountEndpoint                interface{}                       `pulumi:"accountEndpoint"`
	AccountKey                     interface{}                       `pulumi:"accountKey"`
	Annotations                    []interface{}                     `pulumi:"annotations"`
	AzureCloudType                 interface{}                       `pulumi:"azureCloudType"`
	ConnectVia                     *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionMode                 *string                           `pulumi:"connectionMode"`
	ConnectionString               interface{}                       `pulumi:"connectionString"`
	Credential                     *CredentialReference              `pulumi:"credential"`
	Database                       interface{}                       `pulumi:"database"`
	Description                    *string                           `pulumi:"description"`
	EncryptedCredential            interface{}                       `pulumi:"encryptedCredential"`
	Parameters                     map[string]ParameterSpecification `pulumi:"parameters"`
	ServicePrincipalCredential     interface{}                       `pulumi:"servicePrincipalCredential"`
	ServicePrincipalCredentialType *string                           `pulumi:"servicePrincipalCredentialType"`
	ServicePrincipalId             interface{}                       `pulumi:"servicePrincipalId"`
	Tenant                         interface{}                       `pulumi:"tenant"`
	Type                           string                            `pulumi:"type"`
}

type CosmosDbLinkedServiceResponse struct {
	AccountEndpoint                interface{}                               `pulumi:"accountEndpoint"`
	AccountKey                     interface{}                               `pulumi:"accountKey"`
	Annotations                    []interface{}                             `pulumi:"annotations"`
	AzureCloudType                 interface{}                               `pulumi:"azureCloudType"`
	ConnectVia                     *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionMode                 *string                                   `pulumi:"connectionMode"`
	ConnectionString               interface{}                               `pulumi:"connectionString"`
	Credential                     *CredentialReferenceResponse              `pulumi:"credential"`
	Database                       interface{}                               `pulumi:"database"`
	Description                    *string                                   `pulumi:"description"`
	EncryptedCredential            interface{}                               `pulumi:"encryptedCredential"`
	Parameters                     map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	ServicePrincipalCredential     interface{}                               `pulumi:"servicePrincipalCredential"`
	ServicePrincipalCredentialType *string                                   `pulumi:"servicePrincipalCredentialType"`
	ServicePrincipalId             interface{}                               `pulumi:"servicePrincipalId"`
	Tenant                         interface{}                               `pulumi:"tenant"`
	Type                           string                                    `pulumi:"type"`
}

type CosmosDbMongoDbApiCollectionDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Collection        interface{}                       `pulumi:"collection"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type CosmosDbMongoDbApiCollectionDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Collection        interface{}                               `pulumi:"collection"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type CosmosDbMongoDbApiLinkedService struct {
	Annotations            []interface{}                     `pulumi:"annotations"`
	ConnectVia             *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString       interface{}                       `pulumi:"connectionString"`
	Database               interface{}                       `pulumi:"database"`
	Description            *string                           `pulumi:"description"`
	IsServerVersionAbove32 interface{}                       `pulumi:"isServerVersionAbove32"`
	Parameters             map[string]ParameterSpecification `pulumi:"parameters"`
	Type                   string                            `pulumi:"type"`
}

type CosmosDbMongoDbApiLinkedServiceResponse struct {
	Annotations            []interface{}                             `pulumi:"annotations"`
	ConnectVia             *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString       interface{}                               `pulumi:"connectionString"`
	Database               interface{}                               `pulumi:"database"`
	Description            *string                                   `pulumi:"description"`
	IsServerVersionAbove32 interface{}                               `pulumi:"isServerVersionAbove32"`
	Parameters             map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                   string                                    `pulumi:"type"`
}

type CosmosDbMongoDbApiSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            interface{} `pulumi:"writeBehavior"`
}

type CosmosDbMongoDbApiSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            interface{} `pulumi:"writeBehavior"`
}

type CosmosDbMongoDbApiSource struct {
	AdditionalColumns        interface{}                     `pulumi:"additionalColumns"`
	BatchSize                interface{}                     `pulumi:"batchSize"`
	CursorMethods            *MongoDbCursorMethodsProperties `pulumi:"cursorMethods"`
	DisableMetricsCollection interface{}                     `pulumi:"disableMetricsCollection"`
	Filter                   interface{}                     `pulumi:"filter"`
	MaxConcurrentConnections interface{}                     `pulumi:"maxConcurrentConnections"`
	QueryTimeout             interface{}                     `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}                     `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                     `pulumi:"sourceRetryWait"`
	Type                     string                          `pulumi:"type"`
}

type CosmosDbMongoDbApiSourceResponse struct {
	AdditionalColumns        interface{}                             `pulumi:"additionalColumns"`
	BatchSize                interface{}                             `pulumi:"batchSize"`
	CursorMethods            *MongoDbCursorMethodsPropertiesResponse `pulumi:"cursorMethods"`
	DisableMetricsCollection interface{}                             `pulumi:"disableMetricsCollection"`
	Filter                   interface{}                             `pulumi:"filter"`
	MaxConcurrentConnections interface{}                             `pulumi:"maxConcurrentConnections"`
	QueryTimeout             interface{}                             `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}                             `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                             `pulumi:"sourceRetryWait"`
	Type                     string                                  `pulumi:"type"`
}

type CosmosDbSqlApiCollectionDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	CollectionName    interface{}                       `pulumi:"collectionName"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type CosmosDbSqlApiCollectionDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	CollectionName    interface{}                               `pulumi:"collectionName"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type CosmosDbSqlApiSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            interface{} `pulumi:"writeBehavior"`
}

type CosmosDbSqlApiSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            interface{} `pulumi:"writeBehavior"`
}

type CosmosDbSqlApiSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DetectDatetime           interface{} `pulumi:"detectDatetime"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PageSize                 interface{} `pulumi:"pageSize"`
	PreferredRegions         interface{} `pulumi:"preferredRegions"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type CosmosDbSqlApiSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DetectDatetime           interface{} `pulumi:"detectDatetime"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PageSize                 interface{} `pulumi:"pageSize"`
	PreferredRegions         interface{} `pulumi:"preferredRegions"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type CouchbaseLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	CredString          *AzureKeyVaultSecretReference     `pulumi:"credString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Type                string                            `pulumi:"type"`
}

type CouchbaseLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	CredString          *AzureKeyVaultSecretReferenceResponse     `pulumi:"credString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
}

type CouchbaseSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type CouchbaseSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type CouchbaseTableDataset struct {
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

type CouchbaseTableDatasetResponse struct {
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

type CredentialReference struct {
	ReferenceName string `pulumi:"referenceName"`
	Type          string `pulumi:"type"`
}

type CredentialReferenceResponse struct {
	ReferenceName string `pulumi:"referenceName"`
	Type          string `pulumi:"type"`
}

type CustomActivity struct {
	AutoUserSpecification interface{}                    `pulumi:"autoUserSpecification"`
	Command               interface{}                    `pulumi:"command"`
	DependsOn             []ActivityDependency           `pulumi:"dependsOn"`
	Description           *string                        `pulumi:"description"`
	ExtendedProperties    map[string]interface{}         `pulumi:"extendedProperties"`
	FolderPath            interface{}                    `pulumi:"folderPath"`
	LinkedServiceName     *LinkedServiceReference        `pulumi:"linkedServiceName"`
	Name                  string                         `pulumi:"name"`
	Policy                *ActivityPolicy                `pulumi:"policy"`
	ReferenceObjects      *CustomActivityReferenceObject `pulumi:"referenceObjects"`
	ResourceLinkedService *LinkedServiceReference        `pulumi:"resourceLinkedService"`
	RetentionTimeInDays   interface{}                    `pulumi:"retentionTimeInDays"`
	Type                  string                         `pulumi:"type"`
	UserProperties        []UserProperty                 `pulumi:"userProperties"`
}

type CustomActivityReferenceObject struct {
	Datasets       []DatasetReference       `pulumi:"datasets"`
	LinkedServices []LinkedServiceReference `pulumi:"linkedServices"`
}

type CustomActivityReferenceObjectResponse struct {
	Datasets       []DatasetReferenceResponse       `pulumi:"datasets"`
	LinkedServices []LinkedServiceReferenceResponse `pulumi:"linkedServices"`
}

type CustomActivityResponse struct {
	AutoUserSpecification interface{}                            `pulumi:"autoUserSpecification"`
	Command               interface{}                            `pulumi:"command"`
	DependsOn             []ActivityDependencyResponse           `pulumi:"dependsOn"`
	Description           *string                                `pulumi:"description"`
	ExtendedProperties    map[string]interface{}                 `pulumi:"extendedProperties"`
	FolderPath            interface{}                            `pulumi:"folderPath"`
	LinkedServiceName     *LinkedServiceReferenceResponse        `pulumi:"linkedServiceName"`
	Name                  string                                 `pulumi:"name"`
	Policy                *ActivityPolicyResponse                `pulumi:"policy"`
	ReferenceObjects      *CustomActivityReferenceObjectResponse `pulumi:"referenceObjects"`
	ResourceLinkedService *LinkedServiceReferenceResponse        `pulumi:"resourceLinkedService"`
	RetentionTimeInDays   interface{}                            `pulumi:"retentionTimeInDays"`
	Type                  string                                 `pulumi:"type"`
	UserProperties        []UserPropertyResponse                 `pulumi:"userProperties"`
}

type CustomDataSourceLinkedService struct {
	Annotations []interface{}                     `pulumi:"annotations"`
	ConnectVia  *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description *string                           `pulumi:"description"`
	Parameters  map[string]ParameterSpecification `pulumi:"parameters"`
	Type        string                            `pulumi:"type"`
}

type CustomDataSourceLinkedServiceResponse struct {
	Annotations []interface{}                             `pulumi:"annotations"`
	ConnectVia  *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description *string                                   `pulumi:"description"`
	Parameters  map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type        string                                    `pulumi:"type"`
}

type CustomDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type CustomDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type CustomEventsTrigger struct {
	Annotations       []interface{}              `pulumi:"annotations"`
	Description       *string                    `pulumi:"description"`
	Events            []interface{}              `pulumi:"events"`
	Pipelines         []TriggerPipelineReference `pulumi:"pipelines"`
	Scope             string                     `pulumi:"scope"`
	SubjectBeginsWith *string                    `pulumi:"subjectBeginsWith"`
	SubjectEndsWith   *string                    `pulumi:"subjectEndsWith"`
	Type              string                     `pulumi:"type"`
}

type CustomEventsTriggerResponse struct {
	Annotations       []interface{}                      `pulumi:"annotations"`
	Description       *string                            `pulumi:"description"`
	Events            []interface{}                      `pulumi:"events"`
	Pipelines         []TriggerPipelineReferenceResponse `pulumi:"pipelines"`
	RuntimeState      string                             `pulumi:"runtimeState"`
	Scope             string                             `pulumi:"scope"`
	SubjectBeginsWith *string                            `pulumi:"subjectBeginsWith"`
	SubjectEndsWith   *string                            `pulumi:"subjectEndsWith"`
	Type              string                             `pulumi:"type"`
}

type DWCopyCommandDefaultValue struct {
	ColumnName   interface{} `pulumi:"columnName"`
	DefaultValue interface{} `pulumi:"defaultValue"`
}

type DWCopyCommandDefaultValueResponse struct {
	ColumnName   interface{} `pulumi:"columnName"`
	DefaultValue interface{} `pulumi:"defaultValue"`
}

type DWCopyCommandSettings struct {
	AdditionalOptions map[string]string           `pulumi:"additionalOptions"`
	DefaultValues     []DWCopyCommandDefaultValue `pulumi:"defaultValues"`
}

type DWCopyCommandSettingsResponse struct {
	AdditionalOptions map[string]string                   `pulumi:"additionalOptions"`
	DefaultValues     []DWCopyCommandDefaultValueResponse `pulumi:"defaultValues"`
}

type DataFlowFolder struct {
	Name *string `pulumi:"name"`
}

type DataFlowReference struct {
	DatasetParameters interface{}            `pulumi:"datasetParameters"`
	Parameters        map[string]interface{} `pulumi:"parameters"`
	ReferenceName     string                 `pulumi:"referenceName"`
	Type              string                 `pulumi:"type"`
}

type DataFlowReferenceResponse struct {
	DatasetParameters interface{}            `pulumi:"datasetParameters"`
	Parameters        map[string]interface{} `pulumi:"parameters"`
	ReferenceName     string                 `pulumi:"referenceName"`
	Type              string                 `pulumi:"type"`
}

type DataFlowResponseFolder struct {
	Name *string `pulumi:"name"`
}

type DataFlowSink struct {
	Dataset                   *DatasetReference       `pulumi:"dataset"`
	Description               *string                 `pulumi:"description"`
	Flowlet                   *DataFlowReference      `pulumi:"flowlet"`
	LinkedService             *LinkedServiceReference `pulumi:"linkedService"`
	Name                      string                  `pulumi:"name"`
	RejectedDataLinkedService *LinkedServiceReference `pulumi:"rejectedDataLinkedService"`
	SchemaLinkedService       *LinkedServiceReference `pulumi:"schemaLinkedService"`
}

type DataFlowSinkResponse struct {
	Dataset                   *DatasetReferenceResponse       `pulumi:"dataset"`
	Description               *string                         `pulumi:"description"`
	Flowlet                   *DataFlowReferenceResponse      `pulumi:"flowlet"`
	LinkedService             *LinkedServiceReferenceResponse `pulumi:"linkedService"`
	Name                      string                          `pulumi:"name"`
	RejectedDataLinkedService *LinkedServiceReferenceResponse `pulumi:"rejectedDataLinkedService"`
	SchemaLinkedService       *LinkedServiceReferenceResponse `pulumi:"schemaLinkedService"`
}

type DataFlowSource struct {
	Dataset             *DatasetReference       `pulumi:"dataset"`
	Description         *string                 `pulumi:"description"`
	Flowlet             *DataFlowReference      `pulumi:"flowlet"`
	LinkedService       *LinkedServiceReference `pulumi:"linkedService"`
	Name                string                  `pulumi:"name"`
	SchemaLinkedService *LinkedServiceReference `pulumi:"schemaLinkedService"`
}

type DataFlowSourceResponse struct {
	Dataset             *DatasetReferenceResponse       `pulumi:"dataset"`
	Description         *string                         `pulumi:"description"`
	Flowlet             *DataFlowReferenceResponse      `pulumi:"flowlet"`
	LinkedService       *LinkedServiceReferenceResponse `pulumi:"linkedService"`
	Name                string                          `pulumi:"name"`
	SchemaLinkedService *LinkedServiceReferenceResponse `pulumi:"schemaLinkedService"`
}

type DataFlowStagingInfo struct {
	FolderPath    interface{}             `pulumi:"folderPath"`
	LinkedService *LinkedServiceReference `pulumi:"linkedService"`
}

type DataFlowStagingInfoResponse struct {
	FolderPath    interface{}                     `pulumi:"folderPath"`
	LinkedService *LinkedServiceReferenceResponse `pulumi:"linkedService"`
}

type DataLakeAnalyticsUSQLActivity struct {
	CompilationMode     interface{}             `pulumi:"compilationMode"`
	DegreeOfParallelism interface{}             `pulumi:"degreeOfParallelism"`
	DependsOn           []ActivityDependency    `pulumi:"dependsOn"`
	Description         *string                 `pulumi:"description"`
	LinkedServiceName   *LinkedServiceReference `pulumi:"linkedServiceName"`
	Name                string                  `pulumi:"name"`
	Parameters          map[string]interface{}  `pulumi:"parameters"`
	Policy              *ActivityPolicy         `pulumi:"policy"`
	Priority            interface{}             `pulumi:"priority"`
	RuntimeVersion      interface{}             `pulumi:"runtimeVersion"`
	ScriptLinkedService LinkedServiceReference  `pulumi:"scriptLinkedService"`
	ScriptPath          interface{}             `pulumi:"scriptPath"`
	Type                string                  `pulumi:"type"`
	UserProperties      []UserProperty          `pulumi:"userProperties"`
}

type DataLakeAnalyticsUSQLActivityResponse struct {
	CompilationMode     interface{}                     `pulumi:"compilationMode"`
	DegreeOfParallelism interface{}                     `pulumi:"degreeOfParallelism"`
	DependsOn           []ActivityDependencyResponse    `pulumi:"dependsOn"`
	Description         *string                         `pulumi:"description"`
	LinkedServiceName   *LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	Name                string                          `pulumi:"name"`
	Parameters          map[string]interface{}          `pulumi:"parameters"`
	Policy              *ActivityPolicyResponse         `pulumi:"policy"`
	Priority            interface{}                     `pulumi:"priority"`
	RuntimeVersion      interface{}                     `pulumi:"runtimeVersion"`
	ScriptLinkedService LinkedServiceReferenceResponse  `pulumi:"scriptLinkedService"`
	ScriptPath          interface{}                     `pulumi:"scriptPath"`
	Type                string                          `pulumi:"type"`
	UserProperties      []UserPropertyResponse          `pulumi:"userProperties"`
}

type DatabricksNotebookActivity struct {
	BaseParameters    map[string]interface{}   `pulumi:"baseParameters"`
	DependsOn         []ActivityDependency     `pulumi:"dependsOn"`
	Description       *string                  `pulumi:"description"`
	Libraries         []map[string]interface{} `pulumi:"libraries"`
	LinkedServiceName *LinkedServiceReference  `pulumi:"linkedServiceName"`
	Name              string                   `pulumi:"name"`
	NotebookPath      interface{}              `pulumi:"notebookPath"`
	Policy            *ActivityPolicy          `pulumi:"policy"`
	Type              string                   `pulumi:"type"`
	UserProperties    []UserProperty           `pulumi:"userProperties"`
}

type DatabricksNotebookActivityResponse struct {
	BaseParameters    map[string]interface{}          `pulumi:"baseParameters"`
	DependsOn         []ActivityDependencyResponse    `pulumi:"dependsOn"`
	Description       *string                         `pulumi:"description"`
	Libraries         []map[string]interface{}        `pulumi:"libraries"`
	LinkedServiceName *LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	Name              string                          `pulumi:"name"`
	NotebookPath      interface{}                     `pulumi:"notebookPath"`
	Policy            *ActivityPolicyResponse         `pulumi:"policy"`
	Type              string                          `pulumi:"type"`
	UserProperties    []UserPropertyResponse          `pulumi:"userProperties"`
}

type DatabricksSparkJarActivity struct {
	DependsOn         []ActivityDependency     `pulumi:"dependsOn"`
	Description       *string                  `pulumi:"description"`
	Libraries         []map[string]interface{} `pulumi:"libraries"`
	LinkedServiceName *LinkedServiceReference  `pulumi:"linkedServiceName"`
	MainClassName     interface{}              `pulumi:"mainClassName"`
	Name              string                   `pulumi:"name"`
	Parameters        []interface{}            `pulumi:"parameters"`
	Policy            *ActivityPolicy          `pulumi:"policy"`
	Type              string                   `pulumi:"type"`
	UserProperties    []UserProperty           `pulumi:"userProperties"`
}

type DatabricksSparkJarActivityResponse struct {
	DependsOn         []ActivityDependencyResponse    `pulumi:"dependsOn"`
	Description       *string                         `pulumi:"description"`
	Libraries         []map[string]interface{}        `pulumi:"libraries"`
	LinkedServiceName *LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	MainClassName     interface{}                     `pulumi:"mainClassName"`
	Name              string                          `pulumi:"name"`
	Parameters        []interface{}                   `pulumi:"parameters"`
	Policy            *ActivityPolicyResponse         `pulumi:"policy"`
	Type              string                          `pulumi:"type"`
	UserProperties    []UserPropertyResponse          `pulumi:"userProperties"`
}

type DatabricksSparkPythonActivity struct {
	DependsOn         []ActivityDependency     `pulumi:"dependsOn"`
	Description       *string                  `pulumi:"description"`
	Libraries         []map[string]interface{} `pulumi:"libraries"`
	LinkedServiceName *LinkedServiceReference  `pulumi:"linkedServiceName"`
	Name              string                   `pulumi:"name"`
	Parameters        []interface{}            `pulumi:"parameters"`
	Policy            *ActivityPolicy          `pulumi:"policy"`
	PythonFile        interface{}              `pulumi:"pythonFile"`
	Type              string                   `pulumi:"type"`
	UserProperties    []UserProperty           `pulumi:"userProperties"`
}

type DatabricksSparkPythonActivityResponse struct {
	DependsOn         []ActivityDependencyResponse    `pulumi:"dependsOn"`
	Description       *string                         `pulumi:"description"`
	Libraries         []map[string]interface{}        `pulumi:"libraries"`
	LinkedServiceName *LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	Name              string                          `pulumi:"name"`
	Parameters        []interface{}                   `pulumi:"parameters"`
	Policy            *ActivityPolicyResponse         `pulumi:"policy"`
	PythonFile        interface{}                     `pulumi:"pythonFile"`
	Type              string                          `pulumi:"type"`
	UserProperties    []UserPropertyResponse          `pulumi:"userProperties"`
}

type DatasetCompression struct {
	Level interface{} `pulumi:"level"`
	Type  interface{} `pulumi:"type"`
}

type DatasetCompressionResponse struct {
	Level interface{} `pulumi:"level"`
	Type  interface{} `pulumi:"type"`
}

type DatasetFolder struct {
	Name *string `pulumi:"name"`
}

type DatasetReference struct {
	Parameters    map[string]interface{} `pulumi:"parameters"`
	ReferenceName string                 `pulumi:"referenceName"`
	Type          string                 `pulumi:"type"`
}

type DatasetReferenceResponse struct {
	Parameters    map[string]interface{} `pulumi:"parameters"`
	ReferenceName string                 `pulumi:"referenceName"`
	Type          string                 `pulumi:"type"`
}

type DatasetResponseFolder struct {
	Name *string `pulumi:"name"`
}

type DataworldLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ApiToken            interface{}                       `pulumi:"apiToken"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Type                string                            `pulumi:"type"`
}

type DataworldLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ApiToken            interface{}                               `pulumi:"apiToken"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
}

type Db2LinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	AuthenticationType    *string                           `pulumi:"authenticationType"`
	CertificateCommonName interface{}                       `pulumi:"certificateCommonName"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString      interface{}                       `pulumi:"connectionString"`
	Database              interface{}                       `pulumi:"database"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	PackageCollection     interface{}                       `pulumi:"packageCollection"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Password              interface{}                       `pulumi:"password"`
	Server                interface{}                       `pulumi:"server"`
	Type                  string                            `pulumi:"type"`
	Username              interface{}                       `pulumi:"username"`
}

type Db2LinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	AuthenticationType    *string                                   `pulumi:"authenticationType"`
	CertificateCommonName interface{}                               `pulumi:"certificateCommonName"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString      interface{}                               `pulumi:"connectionString"`
	Database              interface{}                               `pulumi:"database"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	PackageCollection     interface{}                               `pulumi:"packageCollection"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password              interface{}                               `pulumi:"password"`
	Server                interface{}                               `pulumi:"server"`
	Type                  string                                    `pulumi:"type"`
	Username              interface{}                               `pulumi:"username"`
}

type Db2Source struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type Db2SourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type Db2TableDataset struct {
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

type Db2TableDatasetResponse struct {
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

type DeleteActivity struct {
	Dataset                  DatasetReference        `pulumi:"dataset"`
	DependsOn                []ActivityDependency    `pulumi:"dependsOn"`
	Description              *string                 `pulumi:"description"`
	EnableLogging            interface{}             `pulumi:"enableLogging"`
	LinkedServiceName        *LinkedServiceReference `pulumi:"linkedServiceName"`
	LogStorageSettings       *LogStorageSettings     `pulumi:"logStorageSettings"`
	MaxConcurrentConnections *int                    `pulumi:"maxConcurrentConnections"`
	Name                     string                  `pulumi:"name"`
	Policy                   *ActivityPolicy         `pulumi:"policy"`
	Recursive                interface{}             `pulumi:"recursive"`
	StoreSettings            interface{}             `pulumi:"storeSettings"`
	Type                     string                  `pulumi:"type"`
	UserProperties           []UserProperty          `pulumi:"userProperties"`
}

type DeleteActivityResponse struct {
	Dataset                  DatasetReferenceResponse        `pulumi:"dataset"`
	DependsOn                []ActivityDependencyResponse    `pulumi:"dependsOn"`
	Description              *string                         `pulumi:"description"`
	EnableLogging            interface{}                     `pulumi:"enableLogging"`
	LinkedServiceName        *LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	LogStorageSettings       *LogStorageSettingsResponse     `pulumi:"logStorageSettings"`
	MaxConcurrentConnections *int                            `pulumi:"maxConcurrentConnections"`
	Name                     string                          `pulumi:"name"`
	Policy                   *ActivityPolicyResponse         `pulumi:"policy"`
	Recursive                interface{}                     `pulumi:"recursive"`
	StoreSettings            interface{}                     `pulumi:"storeSettings"`
	Type                     string                          `pulumi:"type"`
	UserProperties           []UserPropertyResponse          `pulumi:"userProperties"`
}

type DelimitedTextDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	ColumnDelimiter   interface{}                       `pulumi:"columnDelimiter"`
	CompressionCodec  interface{}                       `pulumi:"compressionCodec"`
	CompressionLevel  interface{}                       `pulumi:"compressionLevel"`
	Description       *string                           `pulumi:"description"`
	EncodingName      interface{}                       `pulumi:"encodingName"`
	EscapeChar        interface{}                       `pulumi:"escapeChar"`
	FirstRowAsHeader  interface{}                       `pulumi:"firstRowAsHeader"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Location          interface{}                       `pulumi:"location"`
	NullValue         interface{}                       `pulumi:"nullValue"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	QuoteChar         interface{}                       `pulumi:"quoteChar"`
	RowDelimiter      interface{}                       `pulumi:"rowDelimiter"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type DelimitedTextDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	ColumnDelimiter   interface{}                               `pulumi:"columnDelimiter"`
	CompressionCodec  interface{}                               `pulumi:"compressionCodec"`
	CompressionLevel  interface{}                               `pulumi:"compressionLevel"`
	Description       *string                                   `pulumi:"description"`
	EncodingName      interface{}                               `pulumi:"encodingName"`
	EscapeChar        interface{}                               `pulumi:"escapeChar"`
	FirstRowAsHeader  interface{}                               `pulumi:"firstRowAsHeader"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Location          interface{}                               `pulumi:"location"`
	NullValue         interface{}                               `pulumi:"nullValue"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	QuoteChar         interface{}                               `pulumi:"quoteChar"`
	RowDelimiter      interface{}                               `pulumi:"rowDelimiter"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type DelimitedTextReadSettings struct {
	CompressionProperties interface{} `pulumi:"compressionProperties"`
	SkipLineCount         interface{} `pulumi:"skipLineCount"`
	Type                  string      `pulumi:"type"`
}

type DelimitedTextReadSettingsResponse struct {
	CompressionProperties interface{} `pulumi:"compressionProperties"`
	SkipLineCount         interface{} `pulumi:"skipLineCount"`
	Type                  string      `pulumi:"type"`
}

type DelimitedTextSink struct {
	DisableMetricsCollection interface{}                 `pulumi:"disableMetricsCollection"`
	FormatSettings           *DelimitedTextWriteSettings `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}                 `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{}                 `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}                 `pulumi:"sinkRetryWait"`
	StoreSettings            interface{}                 `pulumi:"storeSettings"`
	Type                     string                      `pulumi:"type"`
	WriteBatchSize           interface{}                 `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}                 `pulumi:"writeBatchTimeout"`
}

type DelimitedTextSinkResponse struct {
	DisableMetricsCollection interface{}                         `pulumi:"disableMetricsCollection"`
	FormatSettings           *DelimitedTextWriteSettingsResponse `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}                         `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{}                         `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}                         `pulumi:"sinkRetryWait"`
	StoreSettings            interface{}                         `pulumi:"storeSettings"`
	Type                     string                              `pulumi:"type"`
	WriteBatchSize           interface{}                         `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}                         `pulumi:"writeBatchTimeout"`
}

type DelimitedTextSource struct {
	AdditionalColumns        interface{}                `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}                `pulumi:"disableMetricsCollection"`
	FormatSettings           *DelimitedTextReadSettings `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}                `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{}                `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                `pulumi:"sourceRetryWait"`
	StoreSettings            interface{}                `pulumi:"storeSettings"`
	Type                     string                     `pulumi:"type"`
}

type DelimitedTextSourceResponse struct {
	AdditionalColumns        interface{}                        `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}                        `pulumi:"disableMetricsCollection"`
	FormatSettings           *DelimitedTextReadSettingsResponse `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}                        `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{}                        `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                        `pulumi:"sourceRetryWait"`
	StoreSettings            interface{}                        `pulumi:"storeSettings"`
	Type                     string                             `pulumi:"type"`
}

type DelimitedTextWriteSettings struct {
	FileExtension  interface{} `pulumi:"fileExtension"`
	FileNamePrefix interface{} `pulumi:"fileNamePrefix"`
	MaxRowsPerFile interface{} `pulumi:"maxRowsPerFile"`
	QuoteAllText   interface{} `pulumi:"quoteAllText"`
	Type           string      `pulumi:"type"`
}

type DelimitedTextWriteSettingsResponse struct {
	FileExtension  interface{} `pulumi:"fileExtension"`
	FileNamePrefix interface{} `pulumi:"fileNamePrefix"`
	MaxRowsPerFile interface{} `pulumi:"maxRowsPerFile"`
	QuoteAllText   interface{} `pulumi:"quoteAllText"`
	Type           string      `pulumi:"type"`
}

type DistcpSettings struct {
	DistcpOptions           interface{} `pulumi:"distcpOptions"`
	ResourceManagerEndpoint interface{} `pulumi:"resourceManagerEndpoint"`
	TempScriptPath          interface{} `pulumi:"tempScriptPath"`
}

type DistcpSettingsResponse struct {
	DistcpOptions           interface{} `pulumi:"distcpOptions"`
	ResourceManagerEndpoint interface{} `pulumi:"resourceManagerEndpoint"`
	TempScriptPath          interface{} `pulumi:"tempScriptPath"`
}

type DocumentDbCollectionDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	CollectionName    interface{}                       `pulumi:"collectionName"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type DocumentDbCollectionDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	CollectionName    interface{}                               `pulumi:"collectionName"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type DocumentDbCollectionSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	NestingSeparator         interface{} `pulumi:"nestingSeparator"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            interface{} `pulumi:"writeBehavior"`
}

type DocumentDbCollectionSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	NestingSeparator         interface{} `pulumi:"nestingSeparator"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            interface{} `pulumi:"writeBehavior"`
}

type DocumentDbCollectionSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	NestingSeparator         interface{} `pulumi:"nestingSeparator"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type DocumentDbCollectionSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	NestingSeparator         interface{} `pulumi:"nestingSeparator"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type DrillLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Pwd                 *AzureKeyVaultSecretReference     `pulumi:"pwd"`
	Type                string                            `pulumi:"type"`
}

type DrillLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Pwd                 *AzureKeyVaultSecretReferenceResponse     `pulumi:"pwd"`
	Type                string                                    `pulumi:"type"`
}

type DrillSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type DrillSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type DrillTableDataset struct {
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

type DrillTableDatasetResponse struct {
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

type DynamicsAXLinkedService struct {
	AadResourceId       interface{}                       `pulumi:"aadResourceId"`
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	ServicePrincipalId  interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey interface{}                       `pulumi:"servicePrincipalKey"`
	Tenant              interface{}                       `pulumi:"tenant"`
	Type                string                            `pulumi:"type"`
	Url                 interface{}                       `pulumi:"url"`
}

type DynamicsAXLinkedServiceResponse struct {
	AadResourceId       interface{}                               `pulumi:"aadResourceId"`
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	ServicePrincipalId  interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey interface{}                               `pulumi:"servicePrincipalKey"`
	Tenant              interface{}                               `pulumi:"tenant"`
	Type                string                                    `pulumi:"type"`
	Url                 interface{}                               `pulumi:"url"`
}

type DynamicsAXResourceDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Path              interface{}                       `pulumi:"path"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type DynamicsAXResourceDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Path              interface{}                               `pulumi:"path"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type DynamicsAXSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type DynamicsAXSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type DynamicsCrmEntityDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	EntityName        interface{}                       `pulumi:"entityName"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type DynamicsCrmEntityDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	EntityName        interface{}                               `pulumi:"entityName"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type DynamicsCrmLinkedService struct {
	Annotations                    []interface{}                     `pulumi:"annotations"`
	AuthenticationType             interface{}                       `pulumi:"authenticationType"`
	ConnectVia                     *IntegrationRuntimeReference      `pulumi:"connectVia"`
	DeploymentType                 interface{}                       `pulumi:"deploymentType"`
	Description                    *string                           `pulumi:"description"`
	EncryptedCredential            interface{}                       `pulumi:"encryptedCredential"`
	HostName                       interface{}                       `pulumi:"hostName"`
	OrganizationName               interface{}                       `pulumi:"organizationName"`
	Parameters                     map[string]ParameterSpecification `pulumi:"parameters"`
	Password                       interface{}                       `pulumi:"password"`
	Port                           interface{}                       `pulumi:"port"`
	ServicePrincipalCredential     interface{}                       `pulumi:"servicePrincipalCredential"`
	ServicePrincipalCredentialType interface{}                       `pulumi:"servicePrincipalCredentialType"`
	ServicePrincipalId             interface{}                       `pulumi:"servicePrincipalId"`
	ServiceUri                     interface{}                       `pulumi:"serviceUri"`
	Type                           string                            `pulumi:"type"`
	Username                       interface{}                       `pulumi:"username"`
}

type DynamicsCrmLinkedServiceResponse struct {
	Annotations                    []interface{}                             `pulumi:"annotations"`
	AuthenticationType             interface{}                               `pulumi:"authenticationType"`
	ConnectVia                     *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	DeploymentType                 interface{}                               `pulumi:"deploymentType"`
	Description                    *string                                   `pulumi:"description"`
	EncryptedCredential            interface{}                               `pulumi:"encryptedCredential"`
	HostName                       interface{}                               `pulumi:"hostName"`
	OrganizationName               interface{}                               `pulumi:"organizationName"`
	Parameters                     map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                       interface{}                               `pulumi:"password"`
	Port                           interface{}                               `pulumi:"port"`
	ServicePrincipalCredential     interface{}                               `pulumi:"servicePrincipalCredential"`
	ServicePrincipalCredentialType interface{}                               `pulumi:"servicePrincipalCredentialType"`
	ServicePrincipalId             interface{}                               `pulumi:"servicePrincipalId"`
	ServiceUri                     interface{}                               `pulumi:"serviceUri"`
	Type                           string                                    `pulumi:"type"`
	Username                       interface{}                               `pulumi:"username"`
}

type DynamicsCrmSink struct {
	AlternateKeyName         interface{} `pulumi:"alternateKeyName"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	IgnoreNullValues         interface{} `pulumi:"ignoreNullValues"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            string      `pulumi:"writeBehavior"`
}

type DynamicsCrmSinkResponse struct {
	AlternateKeyName         interface{} `pulumi:"alternateKeyName"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	IgnoreNullValues         interface{} `pulumi:"ignoreNullValues"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            string      `pulumi:"writeBehavior"`
}

type DynamicsCrmSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type DynamicsCrmSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type DynamicsEntityDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	EntityName        interface{}                       `pulumi:"entityName"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type DynamicsEntityDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	EntityName        interface{}                               `pulumi:"entityName"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type DynamicsLinkedService struct {
	Annotations                    []interface{}                     `pulumi:"annotations"`
	AuthenticationType             interface{}                       `pulumi:"authenticationType"`
	ConnectVia                     *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Credential                     *CredentialReference              `pulumi:"credential"`
	DeploymentType                 interface{}                       `pulumi:"deploymentType"`
	Description                    *string                           `pulumi:"description"`
	EncryptedCredential            interface{}                       `pulumi:"encryptedCredential"`
	HostName                       interface{}                       `pulumi:"hostName"`
	OrganizationName               interface{}                       `pulumi:"organizationName"`
	Parameters                     map[string]ParameterSpecification `pulumi:"parameters"`
	Password                       interface{}                       `pulumi:"password"`
	Port                           interface{}                       `pulumi:"port"`
	ServicePrincipalCredential     interface{}                       `pulumi:"servicePrincipalCredential"`
	ServicePrincipalCredentialType interface{}                       `pulumi:"servicePrincipalCredentialType"`
	ServicePrincipalId             interface{}                       `pulumi:"servicePrincipalId"`
	ServiceUri                     interface{}                       `pulumi:"serviceUri"`
	Type                           string                            `pulumi:"type"`
	Username                       interface{}                       `pulumi:"username"`
}

type DynamicsLinkedServiceResponse struct {
	Annotations                    []interface{}                             `pulumi:"annotations"`
	AuthenticationType             interface{}                               `pulumi:"authenticationType"`
	ConnectVia                     *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Credential                     *CredentialReferenceResponse              `pulumi:"credential"`
	DeploymentType                 interface{}                               `pulumi:"deploymentType"`
	Description                    *string                                   `pulumi:"description"`
	EncryptedCredential            interface{}                               `pulumi:"encryptedCredential"`
	HostName                       interface{}                               `pulumi:"hostName"`
	OrganizationName               interface{}                               `pulumi:"organizationName"`
	Parameters                     map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                       interface{}                               `pulumi:"password"`
	Port                           interface{}                               `pulumi:"port"`
	ServicePrincipalCredential     interface{}                               `pulumi:"servicePrincipalCredential"`
	ServicePrincipalCredentialType interface{}                               `pulumi:"servicePrincipalCredentialType"`
	ServicePrincipalId             interface{}                               `pulumi:"servicePrincipalId"`
	ServiceUri                     interface{}                               `pulumi:"serviceUri"`
	Type                           string                                    `pulumi:"type"`
	Username                       interface{}                               `pulumi:"username"`
}

type DynamicsSink struct {
	AlternateKeyName         interface{} `pulumi:"alternateKeyName"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	IgnoreNullValues         interface{} `pulumi:"ignoreNullValues"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            string      `pulumi:"writeBehavior"`
}

type DynamicsSinkResponse struct {
	AlternateKeyName         interface{} `pulumi:"alternateKeyName"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	IgnoreNullValues         interface{} `pulumi:"ignoreNullValues"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            string      `pulumi:"writeBehavior"`
}

type DynamicsSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type DynamicsSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type EloquaLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Endpoint              interface{}                       `pulumi:"endpoint"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Password              interface{}                       `pulumi:"password"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
	Username              interface{}                       `pulumi:"username"`
}

type EloquaLinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Endpoint              interface{}                               `pulumi:"endpoint"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password              interface{}                               `pulumi:"password"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
	Username              interface{}                               `pulumi:"username"`
}

type EloquaObjectDataset struct {
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

type EloquaObjectDatasetResponse struct {
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

type EloquaSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type EloquaSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type EncryptionConfiguration struct {
	Identity     *CMKIdentityDefinition `pulumi:"identity"`
	KeyName      string                 `pulumi:"keyName"`
	KeyVersion   *string                `pulumi:"keyVersion"`
	VaultBaseUrl string                 `pulumi:"vaultBaseUrl"`
}





type EncryptionConfigurationInput interface {
	pulumi.Input

	ToEncryptionConfigurationOutput() EncryptionConfigurationOutput
	ToEncryptionConfigurationOutputWithContext(context.Context) EncryptionConfigurationOutput
}

type EncryptionConfigurationArgs struct {
	Identity     CMKIdentityDefinitionPtrInput `pulumi:"identity"`
	KeyName      pulumi.StringInput            `pulumi:"keyName"`
	KeyVersion   pulumi.StringPtrInput         `pulumi:"keyVersion"`
	VaultBaseUrl pulumi.StringInput            `pulumi:"vaultBaseUrl"`
}

func (EncryptionConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionConfiguration)(nil)).Elem()
}

func (i EncryptionConfigurationArgs) ToEncryptionConfigurationOutput() EncryptionConfigurationOutput {
	return i.ToEncryptionConfigurationOutputWithContext(context.Background())
}

func (i EncryptionConfigurationArgs) ToEncryptionConfigurationOutputWithContext(ctx context.Context) EncryptionConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionConfigurationOutput)
}

func (i EncryptionConfigurationArgs) ToEncryptionConfigurationPtrOutput() EncryptionConfigurationPtrOutput {
	return i.ToEncryptionConfigurationPtrOutputWithContext(context.Background())
}

func (i EncryptionConfigurationArgs) ToEncryptionConfigurationPtrOutputWithContext(ctx context.Context) EncryptionConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionConfigurationOutput).ToEncryptionConfigurationPtrOutputWithContext(ctx)
}









type EncryptionConfigurationPtrInput interface {
	pulumi.Input

	ToEncryptionConfigurationPtrOutput() EncryptionConfigurationPtrOutput
	ToEncryptionConfigurationPtrOutputWithContext(context.Context) EncryptionConfigurationPtrOutput
}

type encryptionConfigurationPtrType EncryptionConfigurationArgs

func EncryptionConfigurationPtr(v *EncryptionConfigurationArgs) EncryptionConfigurationPtrInput {
	return (*encryptionConfigurationPtrType)(v)
}

func (*encryptionConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionConfiguration)(nil)).Elem()
}

func (i *encryptionConfigurationPtrType) ToEncryptionConfigurationPtrOutput() EncryptionConfigurationPtrOutput {
	return i.ToEncryptionConfigurationPtrOutputWithContext(context.Background())
}

func (i *encryptionConfigurationPtrType) ToEncryptionConfigurationPtrOutputWithContext(ctx context.Context) EncryptionConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionConfigurationPtrOutput)
}

type EncryptionConfigurationOutput struct{ *pulumi.OutputState }

func (EncryptionConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionConfiguration)(nil)).Elem()
}

func (o EncryptionConfigurationOutput) ToEncryptionConfigurationOutput() EncryptionConfigurationOutput {
	return o
}

func (o EncryptionConfigurationOutput) ToEncryptionConfigurationOutputWithContext(ctx context.Context) EncryptionConfigurationOutput {
	return o
}

func (o EncryptionConfigurationOutput) ToEncryptionConfigurationPtrOutput() EncryptionConfigurationPtrOutput {
	return o.ToEncryptionConfigurationPtrOutputWithContext(context.Background())
}

func (o EncryptionConfigurationOutput) ToEncryptionConfigurationPtrOutputWithContext(ctx context.Context) EncryptionConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionConfiguration) *EncryptionConfiguration {
		return &v
	}).(EncryptionConfigurationPtrOutput)
}

func (o EncryptionConfigurationOutput) Identity() CMKIdentityDefinitionPtrOutput {
	return o.ApplyT(func(v EncryptionConfiguration) *CMKIdentityDefinition { return v.Identity }).(CMKIdentityDefinitionPtrOutput)
}

func (o EncryptionConfigurationOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionConfiguration) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o EncryptionConfigurationOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionConfiguration) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

func (o EncryptionConfigurationOutput) VaultBaseUrl() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionConfiguration) string { return v.VaultBaseUrl }).(pulumi.StringOutput)
}

type EncryptionConfigurationPtrOutput struct{ *pulumi.OutputState }

func (EncryptionConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionConfiguration)(nil)).Elem()
}

func (o EncryptionConfigurationPtrOutput) ToEncryptionConfigurationPtrOutput() EncryptionConfigurationPtrOutput {
	return o
}

func (o EncryptionConfigurationPtrOutput) ToEncryptionConfigurationPtrOutputWithContext(ctx context.Context) EncryptionConfigurationPtrOutput {
	return o
}

func (o EncryptionConfigurationPtrOutput) Elem() EncryptionConfigurationOutput {
	return o.ApplyT(func(v *EncryptionConfiguration) EncryptionConfiguration {
		if v != nil {
			return *v
		}
		var ret EncryptionConfiguration
		return ret
	}).(EncryptionConfigurationOutput)
}

func (o EncryptionConfigurationPtrOutput) Identity() CMKIdentityDefinitionPtrOutput {
	return o.ApplyT(func(v *EncryptionConfiguration) *CMKIdentityDefinition {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(CMKIdentityDefinitionPtrOutput)
}

func (o EncryptionConfigurationPtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionConfigurationPtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionConfigurationPtrOutput) VaultBaseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.VaultBaseUrl
	}).(pulumi.StringPtrOutput)
}

type EncryptionConfigurationResponse struct {
	Identity     *CMKIdentityDefinitionResponse `pulumi:"identity"`
	KeyName      string                         `pulumi:"keyName"`
	KeyVersion   *string                        `pulumi:"keyVersion"`
	VaultBaseUrl string                         `pulumi:"vaultBaseUrl"`
}

type EncryptionConfigurationResponseOutput struct{ *pulumi.OutputState }

func (EncryptionConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionConfigurationResponse)(nil)).Elem()
}

func (o EncryptionConfigurationResponseOutput) ToEncryptionConfigurationResponseOutput() EncryptionConfigurationResponseOutput {
	return o
}

func (o EncryptionConfigurationResponseOutput) ToEncryptionConfigurationResponseOutputWithContext(ctx context.Context) EncryptionConfigurationResponseOutput {
	return o
}

func (o EncryptionConfigurationResponseOutput) Identity() CMKIdentityDefinitionResponsePtrOutput {
	return o.ApplyT(func(v EncryptionConfigurationResponse) *CMKIdentityDefinitionResponse { return v.Identity }).(CMKIdentityDefinitionResponsePtrOutput)
}

func (o EncryptionConfigurationResponseOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionConfigurationResponse) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o EncryptionConfigurationResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionConfigurationResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

func (o EncryptionConfigurationResponseOutput) VaultBaseUrl() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionConfigurationResponse) string { return v.VaultBaseUrl }).(pulumi.StringOutput)
}

type EncryptionConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionConfigurationResponse)(nil)).Elem()
}

func (o EncryptionConfigurationResponsePtrOutput) ToEncryptionConfigurationResponsePtrOutput() EncryptionConfigurationResponsePtrOutput {
	return o
}

func (o EncryptionConfigurationResponsePtrOutput) ToEncryptionConfigurationResponsePtrOutputWithContext(ctx context.Context) EncryptionConfigurationResponsePtrOutput {
	return o
}

func (o EncryptionConfigurationResponsePtrOutput) Elem() EncryptionConfigurationResponseOutput {
	return o.ApplyT(func(v *EncryptionConfigurationResponse) EncryptionConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionConfigurationResponse
		return ret
	}).(EncryptionConfigurationResponseOutput)
}

func (o EncryptionConfigurationResponsePtrOutput) Identity() CMKIdentityDefinitionResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionConfigurationResponse) *CMKIdentityDefinitionResponse {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(CMKIdentityDefinitionResponsePtrOutput)
}

func (o EncryptionConfigurationResponsePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionConfigurationResponsePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionConfigurationResponsePtrOutput) VaultBaseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.VaultBaseUrl
	}).(pulumi.StringPtrOutput)
}

type EntityReference struct {
	ReferenceName *string `pulumi:"referenceName"`
	Type          *string `pulumi:"type"`
}

type EntityReferenceResponse struct {
	ReferenceName *string `pulumi:"referenceName"`
	Type          *string `pulumi:"type"`
}

type EnvironmentVariableSetup struct {
	Type          string `pulumi:"type"`
	VariableName  string `pulumi:"variableName"`
	VariableValue string `pulumi:"variableValue"`
}

type EnvironmentVariableSetupResponse struct {
	Type          string `pulumi:"type"`
	VariableName  string `pulumi:"variableName"`
	VariableValue string `pulumi:"variableValue"`
}

type ExcelDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Compression       *DatasetCompression               `pulumi:"compression"`
	Description       *string                           `pulumi:"description"`
	FirstRowAsHeader  interface{}                       `pulumi:"firstRowAsHeader"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Location          interface{}                       `pulumi:"location"`
	NullValue         interface{}                       `pulumi:"nullValue"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Range             interface{}                       `pulumi:"range"`
	Schema            interface{}                       `pulumi:"schema"`
	SheetIndex        interface{}                       `pulumi:"sheetIndex"`
	SheetName         interface{}                       `pulumi:"sheetName"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type ExcelDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Compression       *DatasetCompressionResponse               `pulumi:"compression"`
	Description       *string                                   `pulumi:"description"`
	FirstRowAsHeader  interface{}                               `pulumi:"firstRowAsHeader"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Location          interface{}                               `pulumi:"location"`
	NullValue         interface{}                               `pulumi:"nullValue"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Range             interface{}                               `pulumi:"range"`
	Schema            interface{}                               `pulumi:"schema"`
	SheetIndex        interface{}                               `pulumi:"sheetIndex"`
	SheetName         interface{}                               `pulumi:"sheetName"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type ExcelSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	StoreSettings            interface{} `pulumi:"storeSettings"`
	Type                     string      `pulumi:"type"`
}

type ExcelSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	StoreSettings            interface{} `pulumi:"storeSettings"`
	Type                     string      `pulumi:"type"`
}

type ExecuteDataFlowActivity struct {
	Compute                  *ExecuteDataFlowActivityTypePropertiesCompute `pulumi:"compute"`
	ContinueOnError          interface{}                                   `pulumi:"continueOnError"`
	DataFlow                 DataFlowReference                             `pulumi:"dataFlow"`
	DependsOn                []ActivityDependency                          `pulumi:"dependsOn"`
	Description              *string                                       `pulumi:"description"`
	IntegrationRuntime       *IntegrationRuntimeReference                  `pulumi:"integrationRuntime"`
	LinkedServiceName        *LinkedServiceReference                       `pulumi:"linkedServiceName"`
	Name                     string                                        `pulumi:"name"`
	Policy                   *ActivityPolicy                               `pulumi:"policy"`
	RunConcurrently          interface{}                                   `pulumi:"runConcurrently"`
	SourceStagingConcurrency interface{}                                   `pulumi:"sourceStagingConcurrency"`
	Staging                  *DataFlowStagingInfo                          `pulumi:"staging"`
	TraceLevel               interface{}                                   `pulumi:"traceLevel"`
	Type                     string                                        `pulumi:"type"`
	UserProperties           []UserProperty                                `pulumi:"userProperties"`
}

type ExecuteDataFlowActivityResponse struct {
	Compute                  *ExecuteDataFlowActivityTypePropertiesResponseCompute `pulumi:"compute"`
	ContinueOnError          interface{}                                           `pulumi:"continueOnError"`
	DataFlow                 DataFlowReferenceResponse                             `pulumi:"dataFlow"`
	DependsOn                []ActivityDependencyResponse                          `pulumi:"dependsOn"`
	Description              *string                                               `pulumi:"description"`
	IntegrationRuntime       *IntegrationRuntimeReferenceResponse                  `pulumi:"integrationRuntime"`
	LinkedServiceName        *LinkedServiceReferenceResponse                       `pulumi:"linkedServiceName"`
	Name                     string                                                `pulumi:"name"`
	Policy                   *ActivityPolicyResponse                               `pulumi:"policy"`
	RunConcurrently          interface{}                                           `pulumi:"runConcurrently"`
	SourceStagingConcurrency interface{}                                           `pulumi:"sourceStagingConcurrency"`
	Staging                  *DataFlowStagingInfoResponse                          `pulumi:"staging"`
	TraceLevel               interface{}                                           `pulumi:"traceLevel"`
	Type                     string                                                `pulumi:"type"`
	UserProperties           []UserPropertyResponse                                `pulumi:"userProperties"`
}

type ExecuteDataFlowActivityTypePropertiesCompute struct {
	ComputeType interface{} `pulumi:"computeType"`
	CoreCount   interface{} `pulumi:"coreCount"`
}

type ExecuteDataFlowActivityTypePropertiesResponseCompute struct {
	ComputeType interface{} `pulumi:"computeType"`
	CoreCount   interface{} `pulumi:"coreCount"`
}

type ExecutePipelineActivity struct {
	DependsOn        []ActivityDependency           `pulumi:"dependsOn"`
	Description      *string                        `pulumi:"description"`
	Name             string                         `pulumi:"name"`
	Parameters       map[string]interface{}         `pulumi:"parameters"`
	Pipeline         PipelineReference              `pulumi:"pipeline"`
	Policy           *ExecutePipelineActivityPolicy `pulumi:"policy"`
	Type             string                         `pulumi:"type"`
	UserProperties   []UserProperty                 `pulumi:"userProperties"`
	WaitOnCompletion *bool                          `pulumi:"waitOnCompletion"`
}

type ExecutePipelineActivityPolicy struct {
	SecureInput *bool `pulumi:"secureInput"`
}

type ExecutePipelineActivityPolicyResponse struct {
	SecureInput *bool `pulumi:"secureInput"`
}

type ExecutePipelineActivityResponse struct {
	DependsOn        []ActivityDependencyResponse           `pulumi:"dependsOn"`
	Description      *string                                `pulumi:"description"`
	Name             string                                 `pulumi:"name"`
	Parameters       map[string]interface{}                 `pulumi:"parameters"`
	Pipeline         PipelineReferenceResponse              `pulumi:"pipeline"`
	Policy           *ExecutePipelineActivityPolicyResponse `pulumi:"policy"`
	Type             string                                 `pulumi:"type"`
	UserProperties   []UserPropertyResponse                 `pulumi:"userProperties"`
	WaitOnCompletion *bool                                  `pulumi:"waitOnCompletion"`
}

type ExecuteSSISPackageActivity struct {
	ConnectVia                IntegrationRuntimeReference                  `pulumi:"connectVia"`
	DependsOn                 []ActivityDependency                         `pulumi:"dependsOn"`
	Description               *string                                      `pulumi:"description"`
	EnvironmentPath           interface{}                                  `pulumi:"environmentPath"`
	ExecutionCredential       *SSISExecutionCredential                     `pulumi:"executionCredential"`
	LinkedServiceName         *LinkedServiceReference                      `pulumi:"linkedServiceName"`
	LogLocation               *SSISLogLocation                             `pulumi:"logLocation"`
	LoggingLevel              interface{}                                  `pulumi:"loggingLevel"`
	Name                      string                                       `pulumi:"name"`
	PackageConnectionManagers map[string]map[string]SSISExecutionParameter `pulumi:"packageConnectionManagers"`
	PackageLocation           SSISPackageLocation                          `pulumi:"packageLocation"`
	PackageParameters         map[string]SSISExecutionParameter            `pulumi:"packageParameters"`
	Policy                    *ActivityPolicy                              `pulumi:"policy"`
	ProjectConnectionManagers map[string]map[string]SSISExecutionParameter `pulumi:"projectConnectionManagers"`
	ProjectParameters         map[string]SSISExecutionParameter            `pulumi:"projectParameters"`
	PropertyOverrides         map[string]SSISPropertyOverride              `pulumi:"propertyOverrides"`
	Runtime                   interface{}                                  `pulumi:"runtime"`
	Type                      string                                       `pulumi:"type"`
	UserProperties            []UserProperty                               `pulumi:"userProperties"`
}

type ExecuteSSISPackageActivityResponse struct {
	ConnectVia                IntegrationRuntimeReferenceResponse                  `pulumi:"connectVia"`
	DependsOn                 []ActivityDependencyResponse                         `pulumi:"dependsOn"`
	Description               *string                                              `pulumi:"description"`
	EnvironmentPath           interface{}                                          `pulumi:"environmentPath"`
	ExecutionCredential       *SSISExecutionCredentialResponse                     `pulumi:"executionCredential"`
	LinkedServiceName         *LinkedServiceReferenceResponse                      `pulumi:"linkedServiceName"`
	LogLocation               *SSISLogLocationResponse                             `pulumi:"logLocation"`
	LoggingLevel              interface{}                                          `pulumi:"loggingLevel"`
	Name                      string                                               `pulumi:"name"`
	PackageConnectionManagers map[string]map[string]SSISExecutionParameterResponse `pulumi:"packageConnectionManagers"`
	PackageLocation           SSISPackageLocationResponse                          `pulumi:"packageLocation"`
	PackageParameters         map[string]SSISExecutionParameterResponse            `pulumi:"packageParameters"`
	Policy                    *ActivityPolicyResponse                              `pulumi:"policy"`
	ProjectConnectionManagers map[string]map[string]SSISExecutionParameterResponse `pulumi:"projectConnectionManagers"`
	ProjectParameters         map[string]SSISExecutionParameterResponse            `pulumi:"projectParameters"`
	PropertyOverrides         map[string]SSISPropertyOverrideResponse              `pulumi:"propertyOverrides"`
	Runtime                   interface{}                                          `pulumi:"runtime"`
	Type                      string                                               `pulumi:"type"`
	UserProperties            []UserPropertyResponse                               `pulumi:"userProperties"`
}

type ExecuteWranglingDataflowActivity struct {
	Compute                  *ExecuteDataFlowActivityTypePropertiesCompute `pulumi:"compute"`
	ContinueOnError          interface{}                                   `pulumi:"continueOnError"`
	DataFlow                 DataFlowReference                             `pulumi:"dataFlow"`
	DependsOn                []ActivityDependency                          `pulumi:"dependsOn"`
	Description              *string                                       `pulumi:"description"`
	IntegrationRuntime       *IntegrationRuntimeReference                  `pulumi:"integrationRuntime"`
	Name                     string                                        `pulumi:"name"`
	Policy                   *ActivityPolicy                               `pulumi:"policy"`
	Queries                  []PowerQuerySinkMapping                       `pulumi:"queries"`
	RunConcurrently          interface{}                                   `pulumi:"runConcurrently"`
	Sinks                    map[string]PowerQuerySink                     `pulumi:"sinks"`
	SourceStagingConcurrency interface{}                                   `pulumi:"sourceStagingConcurrency"`
	Staging                  *DataFlowStagingInfo                          `pulumi:"staging"`
	TraceLevel               interface{}                                   `pulumi:"traceLevel"`
	Type                     string                                        `pulumi:"type"`
	UserProperties           []UserProperty                                `pulumi:"userProperties"`
}

type ExecuteWranglingDataflowActivityResponse struct {
	Compute                  *ExecuteDataFlowActivityTypePropertiesResponseCompute `pulumi:"compute"`
	ContinueOnError          interface{}                                           `pulumi:"continueOnError"`
	DataFlow                 DataFlowReferenceResponse                             `pulumi:"dataFlow"`
	DependsOn                []ActivityDependencyResponse                          `pulumi:"dependsOn"`
	Description              *string                                               `pulumi:"description"`
	IntegrationRuntime       *IntegrationRuntimeReferenceResponse                  `pulumi:"integrationRuntime"`
	Name                     string                                                `pulumi:"name"`
	Policy                   *ActivityPolicyResponse                               `pulumi:"policy"`
	Queries                  []PowerQuerySinkMappingResponse                       `pulumi:"queries"`
	RunConcurrently          interface{}                                           `pulumi:"runConcurrently"`
	Sinks                    map[string]PowerQuerySinkResponse                     `pulumi:"sinks"`
	SourceStagingConcurrency interface{}                                           `pulumi:"sourceStagingConcurrency"`
	Staging                  *DataFlowStagingInfoResponse                          `pulumi:"staging"`
	TraceLevel               interface{}                                           `pulumi:"traceLevel"`
	Type                     string                                                `pulumi:"type"`
	UserProperties           []UserPropertyResponse                                `pulumi:"userProperties"`
}

type ExecutionActivity struct {
	DependsOn         []ActivityDependency    `pulumi:"dependsOn"`
	Description       *string                 `pulumi:"description"`
	LinkedServiceName *LinkedServiceReference `pulumi:"linkedServiceName"`
	Name              string                  `pulumi:"name"`
	Policy            *ActivityPolicy         `pulumi:"policy"`
	Type              string                  `pulumi:"type"`
	UserProperties    []UserProperty          `pulumi:"userProperties"`
}

type ExecutionActivityResponse struct {
	DependsOn         []ActivityDependencyResponse    `pulumi:"dependsOn"`
	Description       *string                         `pulumi:"description"`
	LinkedServiceName *LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	Name              string                          `pulumi:"name"`
	Policy            *ActivityPolicyResponse         `pulumi:"policy"`
	Type              string                          `pulumi:"type"`
	UserProperties    []UserPropertyResponse          `pulumi:"userProperties"`
}

type Expression struct {
	Type  string `pulumi:"type"`
	Value string `pulumi:"value"`
}

type ExpressionResponse struct {
	Type  string `pulumi:"type"`
	Value string `pulumi:"value"`
}

type FactoryGitHubConfiguration struct {
	AccountName         string              `pulumi:"accountName"`
	ClientId            *string             `pulumi:"clientId"`
	ClientSecret        *GitHubClientSecret `pulumi:"clientSecret"`
	CollaborationBranch string              `pulumi:"collaborationBranch"`
	HostName            *string             `pulumi:"hostName"`
	LastCommitId        *string             `pulumi:"lastCommitId"`
	RepositoryName      string              `pulumi:"repositoryName"`
	RootFolder          string              `pulumi:"rootFolder"`
	Type                string              `pulumi:"type"`
}

type FactoryGitHubConfigurationResponse struct {
	AccountName         string                      `pulumi:"accountName"`
	ClientId            *string                     `pulumi:"clientId"`
	ClientSecret        *GitHubClientSecretResponse `pulumi:"clientSecret"`
	CollaborationBranch string                      `pulumi:"collaborationBranch"`
	HostName            *string                     `pulumi:"hostName"`
	LastCommitId        *string                     `pulumi:"lastCommitId"`
	RepositoryName      string                      `pulumi:"repositoryName"`
	RootFolder          string                      `pulumi:"rootFolder"`
	Type                string                      `pulumi:"type"`
}

type FactoryIdentity struct {
	Type                   string                 `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type FactoryIdentityInput interface {
	pulumi.Input

	ToFactoryIdentityOutput() FactoryIdentityOutput
	ToFactoryIdentityOutputWithContext(context.Context) FactoryIdentityOutput
}

type FactoryIdentityArgs struct {
	Type                   pulumi.StringInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput    `pulumi:"userAssignedIdentities"`
}

func (FactoryIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FactoryIdentity)(nil)).Elem()
}

func (i FactoryIdentityArgs) ToFactoryIdentityOutput() FactoryIdentityOutput {
	return i.ToFactoryIdentityOutputWithContext(context.Background())
}

func (i FactoryIdentityArgs) ToFactoryIdentityOutputWithContext(ctx context.Context) FactoryIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FactoryIdentityOutput)
}

func (i FactoryIdentityArgs) ToFactoryIdentityPtrOutput() FactoryIdentityPtrOutput {
	return i.ToFactoryIdentityPtrOutputWithContext(context.Background())
}

func (i FactoryIdentityArgs) ToFactoryIdentityPtrOutputWithContext(ctx context.Context) FactoryIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FactoryIdentityOutput).ToFactoryIdentityPtrOutputWithContext(ctx)
}









type FactoryIdentityPtrInput interface {
	pulumi.Input

	ToFactoryIdentityPtrOutput() FactoryIdentityPtrOutput
	ToFactoryIdentityPtrOutputWithContext(context.Context) FactoryIdentityPtrOutput
}

type factoryIdentityPtrType FactoryIdentityArgs

func FactoryIdentityPtr(v *FactoryIdentityArgs) FactoryIdentityPtrInput {
	return (*factoryIdentityPtrType)(v)
}

func (*factoryIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FactoryIdentity)(nil)).Elem()
}

func (i *factoryIdentityPtrType) ToFactoryIdentityPtrOutput() FactoryIdentityPtrOutput {
	return i.ToFactoryIdentityPtrOutputWithContext(context.Background())
}

func (i *factoryIdentityPtrType) ToFactoryIdentityPtrOutputWithContext(ctx context.Context) FactoryIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FactoryIdentityPtrOutput)
}

type FactoryIdentityOutput struct{ *pulumi.OutputState }

func (FactoryIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FactoryIdentity)(nil)).Elem()
}

func (o FactoryIdentityOutput) ToFactoryIdentityOutput() FactoryIdentityOutput {
	return o
}

func (o FactoryIdentityOutput) ToFactoryIdentityOutputWithContext(ctx context.Context) FactoryIdentityOutput {
	return o
}

func (o FactoryIdentityOutput) ToFactoryIdentityPtrOutput() FactoryIdentityPtrOutput {
	return o.ToFactoryIdentityPtrOutputWithContext(context.Background())
}

func (o FactoryIdentityOutput) ToFactoryIdentityPtrOutputWithContext(ctx context.Context) FactoryIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FactoryIdentity) *FactoryIdentity {
		return &v
	}).(FactoryIdentityPtrOutput)
}

func (o FactoryIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v FactoryIdentity) string { return v.Type }).(pulumi.StringOutput)
}

func (o FactoryIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v FactoryIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type FactoryIdentityPtrOutput struct{ *pulumi.OutputState }

func (FactoryIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FactoryIdentity)(nil)).Elem()
}

func (o FactoryIdentityPtrOutput) ToFactoryIdentityPtrOutput() FactoryIdentityPtrOutput {
	return o
}

func (o FactoryIdentityPtrOutput) ToFactoryIdentityPtrOutputWithContext(ctx context.Context) FactoryIdentityPtrOutput {
	return o
}

func (o FactoryIdentityPtrOutput) Elem() FactoryIdentityOutput {
	return o.ApplyT(func(v *FactoryIdentity) FactoryIdentity {
		if v != nil {
			return *v
		}
		var ret FactoryIdentity
		return ret
	}).(FactoryIdentityOutput)
}

func (o FactoryIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FactoryIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o FactoryIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *FactoryIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type FactoryIdentityResponse struct {
	PrincipalId            string                 `pulumi:"principalId"`
	TenantId               string                 `pulumi:"tenantId"`
	Type                   string                 `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}

type FactoryIdentityResponseOutput struct{ *pulumi.OutputState }

func (FactoryIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FactoryIdentityResponse)(nil)).Elem()
}

func (o FactoryIdentityResponseOutput) ToFactoryIdentityResponseOutput() FactoryIdentityResponseOutput {
	return o
}

func (o FactoryIdentityResponseOutput) ToFactoryIdentityResponseOutputWithContext(ctx context.Context) FactoryIdentityResponseOutput {
	return o
}

func (o FactoryIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v FactoryIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o FactoryIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v FactoryIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o FactoryIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v FactoryIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o FactoryIdentityResponseOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v FactoryIdentityResponse) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type FactoryIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (FactoryIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FactoryIdentityResponse)(nil)).Elem()
}

func (o FactoryIdentityResponsePtrOutput) ToFactoryIdentityResponsePtrOutput() FactoryIdentityResponsePtrOutput {
	return o
}

func (o FactoryIdentityResponsePtrOutput) ToFactoryIdentityResponsePtrOutputWithContext(ctx context.Context) FactoryIdentityResponsePtrOutput {
	return o
}

func (o FactoryIdentityResponsePtrOutput) Elem() FactoryIdentityResponseOutput {
	return o.ApplyT(func(v *FactoryIdentityResponse) FactoryIdentityResponse {
		if v != nil {
			return *v
		}
		var ret FactoryIdentityResponse
		return ret
	}).(FactoryIdentityResponseOutput)
}

func (o FactoryIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FactoryIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o FactoryIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FactoryIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o FactoryIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FactoryIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o FactoryIdentityResponsePtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *FactoryIdentityResponse) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type FactoryVSTSConfiguration struct {
	AccountName         string  `pulumi:"accountName"`
	CollaborationBranch string  `pulumi:"collaborationBranch"`
	LastCommitId        *string `pulumi:"lastCommitId"`
	ProjectName         string  `pulumi:"projectName"`
	RepositoryName      string  `pulumi:"repositoryName"`
	RootFolder          string  `pulumi:"rootFolder"`
	TenantId            *string `pulumi:"tenantId"`
	Type                string  `pulumi:"type"`
}

type FactoryVSTSConfigurationResponse struct {
	AccountName         string  `pulumi:"accountName"`
	CollaborationBranch string  `pulumi:"collaborationBranch"`
	LastCommitId        *string `pulumi:"lastCommitId"`
	ProjectName         string  `pulumi:"projectName"`
	RepositoryName      string  `pulumi:"repositoryName"`
	RootFolder          string  `pulumi:"rootFolder"`
	TenantId            *string `pulumi:"tenantId"`
	Type                string  `pulumi:"type"`
}

type FailActivity struct {
	DependsOn      []ActivityDependency `pulumi:"dependsOn"`
	Description    *string              `pulumi:"description"`
	ErrorCode      interface{}          `pulumi:"errorCode"`
	Message        interface{}          `pulumi:"message"`
	Name           string               `pulumi:"name"`
	Type           string               `pulumi:"type"`
	UserProperties []UserProperty       `pulumi:"userProperties"`
}

type FailActivityResponse struct {
	DependsOn      []ActivityDependencyResponse `pulumi:"dependsOn"`
	Description    *string                      `pulumi:"description"`
	ErrorCode      interface{}                  `pulumi:"errorCode"`
	Message        interface{}                  `pulumi:"message"`
	Name           string                       `pulumi:"name"`
	Type           string                       `pulumi:"type"`
	UserProperties []UserPropertyResponse       `pulumi:"userProperties"`
}

type FileServerLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Host                interface{}                       `pulumi:"host"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Type                string                            `pulumi:"type"`
	UserId              interface{}                       `pulumi:"userId"`
}

type FileServerLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Host                interface{}                               `pulumi:"host"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
	UserId              interface{}                               `pulumi:"userId"`
}

type FileServerLocation struct {
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
}

type FileServerLocationResponse struct {
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
}

type FileServerReadSettings struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileFilter                 interface{} `pulumi:"fileFilter"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type FileServerReadSettingsResponse struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileFilter                 interface{} `pulumi:"fileFilter"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type FileServerWriteSettings struct {
	CopyBehavior             interface{} `pulumi:"copyBehavior"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Type                     string      `pulumi:"type"`
}

type FileServerWriteSettingsResponse struct {
	CopyBehavior             interface{} `pulumi:"copyBehavior"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Type                     string      `pulumi:"type"`
}

type FileShareDataset struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	Compression           *DatasetCompression               `pulumi:"compression"`
	Description           *string                           `pulumi:"description"`
	FileFilter            interface{}                       `pulumi:"fileFilter"`
	FileName              interface{}                       `pulumi:"fileName"`
	Folder                *DatasetFolder                    `pulumi:"folder"`
	FolderPath            interface{}                       `pulumi:"folderPath"`
	Format                interface{}                       `pulumi:"format"`
	LinkedServiceName     LinkedServiceReference            `pulumi:"linkedServiceName"`
	ModifiedDatetimeEnd   interface{}                       `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart interface{}                       `pulumi:"modifiedDatetimeStart"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Schema                interface{}                       `pulumi:"schema"`
	Structure             interface{}                       `pulumi:"structure"`
	Type                  string                            `pulumi:"type"`
}

type FileShareDatasetResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	Compression           *DatasetCompressionResponse               `pulumi:"compression"`
	Description           *string                                   `pulumi:"description"`
	FileFilter            interface{}                               `pulumi:"fileFilter"`
	FileName              interface{}                               `pulumi:"fileName"`
	Folder                *DatasetResponseFolder                    `pulumi:"folder"`
	FolderPath            interface{}                               `pulumi:"folderPath"`
	Format                interface{}                               `pulumi:"format"`
	LinkedServiceName     LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	ModifiedDatetimeEnd   interface{}                               `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart interface{}                               `pulumi:"modifiedDatetimeStart"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema                interface{}                               `pulumi:"schema"`
	Structure             interface{}                               `pulumi:"structure"`
	Type                  string                                    `pulumi:"type"`
}

type FileSystemSink struct {
	CopyBehavior             interface{} `pulumi:"copyBehavior"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type FileSystemSinkResponse struct {
	CopyBehavior             interface{} `pulumi:"copyBehavior"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type FileSystemSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Recursive                interface{} `pulumi:"recursive"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type FileSystemSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Recursive                interface{} `pulumi:"recursive"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type FilterActivity struct {
	Condition      Expression           `pulumi:"condition"`
	DependsOn      []ActivityDependency `pulumi:"dependsOn"`
	Description    *string              `pulumi:"description"`
	Items          Expression           `pulumi:"items"`
	Name           string               `pulumi:"name"`
	Type           string               `pulumi:"type"`
	UserProperties []UserProperty       `pulumi:"userProperties"`
}

type FilterActivityResponse struct {
	Condition      ExpressionResponse           `pulumi:"condition"`
	DependsOn      []ActivityDependencyResponse `pulumi:"dependsOn"`
	Description    *string                      `pulumi:"description"`
	Items          ExpressionResponse           `pulumi:"items"`
	Name           string                       `pulumi:"name"`
	Type           string                       `pulumi:"type"`
	UserProperties []UserPropertyResponse       `pulumi:"userProperties"`
}

type Flowlet struct {
	Annotations     []interface{}    `pulumi:"annotations"`
	Description     *string          `pulumi:"description"`
	Folder          *DataFlowFolder  `pulumi:"folder"`
	Script          *string          `pulumi:"script"`
	ScriptLines     []string         `pulumi:"scriptLines"`
	Sinks           []DataFlowSink   `pulumi:"sinks"`
	Sources         []DataFlowSource `pulumi:"sources"`
	Transformations []Transformation `pulumi:"transformations"`
	Type            string           `pulumi:"type"`
}

type FlowletResponse struct {
	Annotations     []interface{}            `pulumi:"annotations"`
	Description     *string                  `pulumi:"description"`
	Folder          *DataFlowResponseFolder  `pulumi:"folder"`
	Script          *string                  `pulumi:"script"`
	ScriptLines     []string                 `pulumi:"scriptLines"`
	Sinks           []DataFlowSinkResponse   `pulumi:"sinks"`
	Sources         []DataFlowSourceResponse `pulumi:"sources"`
	Transformations []TransformationResponse `pulumi:"transformations"`
	Type            string                   `pulumi:"type"`
}

type ForEachActivity struct {
	Activities     []interface{}        `pulumi:"activities"`
	BatchCount     *int                 `pulumi:"batchCount"`
	DependsOn      []ActivityDependency `pulumi:"dependsOn"`
	Description    *string              `pulumi:"description"`
	IsSequential   *bool                `pulumi:"isSequential"`
	Items          Expression           `pulumi:"items"`
	Name           string               `pulumi:"name"`
	Type           string               `pulumi:"type"`
	UserProperties []UserProperty       `pulumi:"userProperties"`
}

type ForEachActivityResponse struct {
	Activities     []interface{}                `pulumi:"activities"`
	BatchCount     *int                         `pulumi:"batchCount"`
	DependsOn      []ActivityDependencyResponse `pulumi:"dependsOn"`
	Description    *string                      `pulumi:"description"`
	IsSequential   *bool                        `pulumi:"isSequential"`
	Items          ExpressionResponse           `pulumi:"items"`
	Name           string                       `pulumi:"name"`
	Type           string                       `pulumi:"type"`
	UserProperties []UserPropertyResponse       `pulumi:"userProperties"`
}

type FtpReadSettings struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableChunking            interface{} `pulumi:"disableChunking"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	UseBinaryTransfer          *bool       `pulumi:"useBinaryTransfer"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type FtpReadSettingsResponse struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableChunking            interface{} `pulumi:"disableChunking"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	UseBinaryTransfer          *bool       `pulumi:"useBinaryTransfer"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type FtpServerLinkedService struct {
	Annotations                       []interface{}                     `pulumi:"annotations"`
	AuthenticationType                *string                           `pulumi:"authenticationType"`
	ConnectVia                        *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description                       *string                           `pulumi:"description"`
	EnableServerCertificateValidation interface{}                       `pulumi:"enableServerCertificateValidation"`
	EnableSsl                         interface{}                       `pulumi:"enableSsl"`
	EncryptedCredential               interface{}                       `pulumi:"encryptedCredential"`
	Host                              interface{}                       `pulumi:"host"`
	Parameters                        map[string]ParameterSpecification `pulumi:"parameters"`
	Password                          interface{}                       `pulumi:"password"`
	Port                              interface{}                       `pulumi:"port"`
	Type                              string                            `pulumi:"type"`
	UserName                          interface{}                       `pulumi:"userName"`
}

type FtpServerLinkedServiceResponse struct {
	Annotations                       []interface{}                             `pulumi:"annotations"`
	AuthenticationType                *string                                   `pulumi:"authenticationType"`
	ConnectVia                        *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description                       *string                                   `pulumi:"description"`
	EnableServerCertificateValidation interface{}                               `pulumi:"enableServerCertificateValidation"`
	EnableSsl                         interface{}                               `pulumi:"enableSsl"`
	EncryptedCredential               interface{}                               `pulumi:"encryptedCredential"`
	Host                              interface{}                               `pulumi:"host"`
	Parameters                        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                          interface{}                               `pulumi:"password"`
	Port                              interface{}                               `pulumi:"port"`
	Type                              string                                    `pulumi:"type"`
	UserName                          interface{}                               `pulumi:"userName"`
}

type FtpServerLocation struct {
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
}

type FtpServerLocationResponse struct {
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
}

type GetMetadataActivity struct {
	Dataset           DatasetReference        `pulumi:"dataset"`
	DependsOn         []ActivityDependency    `pulumi:"dependsOn"`
	Description       *string                 `pulumi:"description"`
	FieldList         []interface{}           `pulumi:"fieldList"`
	FormatSettings    interface{}             `pulumi:"formatSettings"`
	LinkedServiceName *LinkedServiceReference `pulumi:"linkedServiceName"`
	Name              string                  `pulumi:"name"`
	Policy            *ActivityPolicy         `pulumi:"policy"`
	StoreSettings     interface{}             `pulumi:"storeSettings"`
	Type              string                  `pulumi:"type"`
	UserProperties    []UserProperty          `pulumi:"userProperties"`
}

type GetMetadataActivityResponse struct {
	Dataset           DatasetReferenceResponse        `pulumi:"dataset"`
	DependsOn         []ActivityDependencyResponse    `pulumi:"dependsOn"`
	Description       *string                         `pulumi:"description"`
	FieldList         []interface{}                   `pulumi:"fieldList"`
	FormatSettings    interface{}                     `pulumi:"formatSettings"`
	LinkedServiceName *LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	Name              string                          `pulumi:"name"`
	Policy            *ActivityPolicyResponse         `pulumi:"policy"`
	StoreSettings     interface{}                     `pulumi:"storeSettings"`
	Type              string                          `pulumi:"type"`
	UserProperties    []UserPropertyResponse          `pulumi:"userProperties"`
}

type GitHubClientSecret struct {
	ByoaSecretAkvUrl *string `pulumi:"byoaSecretAkvUrl"`
	ByoaSecretName   *string `pulumi:"byoaSecretName"`
}





type GitHubClientSecretInput interface {
	pulumi.Input

	ToGitHubClientSecretOutput() GitHubClientSecretOutput
	ToGitHubClientSecretOutputWithContext(context.Context) GitHubClientSecretOutput
}

type GitHubClientSecretArgs struct {
	ByoaSecretAkvUrl pulumi.StringPtrInput `pulumi:"byoaSecretAkvUrl"`
	ByoaSecretName   pulumi.StringPtrInput `pulumi:"byoaSecretName"`
}

func (GitHubClientSecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubClientSecret)(nil)).Elem()
}

func (i GitHubClientSecretArgs) ToGitHubClientSecretOutput() GitHubClientSecretOutput {
	return i.ToGitHubClientSecretOutputWithContext(context.Background())
}

func (i GitHubClientSecretArgs) ToGitHubClientSecretOutputWithContext(ctx context.Context) GitHubClientSecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubClientSecretOutput)
}

func (i GitHubClientSecretArgs) ToGitHubClientSecretPtrOutput() GitHubClientSecretPtrOutput {
	return i.ToGitHubClientSecretPtrOutputWithContext(context.Background())
}

func (i GitHubClientSecretArgs) ToGitHubClientSecretPtrOutputWithContext(ctx context.Context) GitHubClientSecretPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubClientSecretOutput).ToGitHubClientSecretPtrOutputWithContext(ctx)
}









type GitHubClientSecretPtrInput interface {
	pulumi.Input

	ToGitHubClientSecretPtrOutput() GitHubClientSecretPtrOutput
	ToGitHubClientSecretPtrOutputWithContext(context.Context) GitHubClientSecretPtrOutput
}

type gitHubClientSecretPtrType GitHubClientSecretArgs

func GitHubClientSecretPtr(v *GitHubClientSecretArgs) GitHubClientSecretPtrInput {
	return (*gitHubClientSecretPtrType)(v)
}

func (*gitHubClientSecretPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubClientSecret)(nil)).Elem()
}

func (i *gitHubClientSecretPtrType) ToGitHubClientSecretPtrOutput() GitHubClientSecretPtrOutput {
	return i.ToGitHubClientSecretPtrOutputWithContext(context.Background())
}

func (i *gitHubClientSecretPtrType) ToGitHubClientSecretPtrOutputWithContext(ctx context.Context) GitHubClientSecretPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubClientSecretPtrOutput)
}

type GitHubClientSecretOutput struct{ *pulumi.OutputState }

func (GitHubClientSecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubClientSecret)(nil)).Elem()
}

func (o GitHubClientSecretOutput) ToGitHubClientSecretOutput() GitHubClientSecretOutput {
	return o
}

func (o GitHubClientSecretOutput) ToGitHubClientSecretOutputWithContext(ctx context.Context) GitHubClientSecretOutput {
	return o
}

func (o GitHubClientSecretOutput) ToGitHubClientSecretPtrOutput() GitHubClientSecretPtrOutput {
	return o.ToGitHubClientSecretPtrOutputWithContext(context.Background())
}

func (o GitHubClientSecretOutput) ToGitHubClientSecretPtrOutputWithContext(ctx context.Context) GitHubClientSecretPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GitHubClientSecret) *GitHubClientSecret {
		return &v
	}).(GitHubClientSecretPtrOutput)
}

func (o GitHubClientSecretOutput) ByoaSecretAkvUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubClientSecret) *string { return v.ByoaSecretAkvUrl }).(pulumi.StringPtrOutput)
}

func (o GitHubClientSecretOutput) ByoaSecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubClientSecret) *string { return v.ByoaSecretName }).(pulumi.StringPtrOutput)
}

type GitHubClientSecretPtrOutput struct{ *pulumi.OutputState }

func (GitHubClientSecretPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubClientSecret)(nil)).Elem()
}

func (o GitHubClientSecretPtrOutput) ToGitHubClientSecretPtrOutput() GitHubClientSecretPtrOutput {
	return o
}

func (o GitHubClientSecretPtrOutput) ToGitHubClientSecretPtrOutputWithContext(ctx context.Context) GitHubClientSecretPtrOutput {
	return o
}

func (o GitHubClientSecretPtrOutput) Elem() GitHubClientSecretOutput {
	return o.ApplyT(func(v *GitHubClientSecret) GitHubClientSecret {
		if v != nil {
			return *v
		}
		var ret GitHubClientSecret
		return ret
	}).(GitHubClientSecretOutput)
}

func (o GitHubClientSecretPtrOutput) ByoaSecretAkvUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubClientSecret) *string {
		if v == nil {
			return nil
		}
		return v.ByoaSecretAkvUrl
	}).(pulumi.StringPtrOutput)
}

func (o GitHubClientSecretPtrOutput) ByoaSecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubClientSecret) *string {
		if v == nil {
			return nil
		}
		return v.ByoaSecretName
	}).(pulumi.StringPtrOutput)
}

type GitHubClientSecretResponse struct {
	ByoaSecretAkvUrl *string `pulumi:"byoaSecretAkvUrl"`
	ByoaSecretName   *string `pulumi:"byoaSecretName"`
}

type GlobalParameterSpecification struct {
	Type  string      `pulumi:"type"`
	Value interface{} `pulumi:"value"`
}





type GlobalParameterSpecificationInput interface {
	pulumi.Input

	ToGlobalParameterSpecificationOutput() GlobalParameterSpecificationOutput
	ToGlobalParameterSpecificationOutputWithContext(context.Context) GlobalParameterSpecificationOutput
}

type GlobalParameterSpecificationArgs struct {
	Type  pulumi.StringInput `pulumi:"type"`
	Value pulumi.Input       `pulumi:"value"`
}

func (GlobalParameterSpecificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalParameterSpecification)(nil)).Elem()
}

func (i GlobalParameterSpecificationArgs) ToGlobalParameterSpecificationOutput() GlobalParameterSpecificationOutput {
	return i.ToGlobalParameterSpecificationOutputWithContext(context.Background())
}

func (i GlobalParameterSpecificationArgs) ToGlobalParameterSpecificationOutputWithContext(ctx context.Context) GlobalParameterSpecificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalParameterSpecificationOutput)
}





type GlobalParameterSpecificationMapInput interface {
	pulumi.Input

	ToGlobalParameterSpecificationMapOutput() GlobalParameterSpecificationMapOutput
	ToGlobalParameterSpecificationMapOutputWithContext(context.Context) GlobalParameterSpecificationMapOutput
}

type GlobalParameterSpecificationMap map[string]GlobalParameterSpecificationInput

func (GlobalParameterSpecificationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GlobalParameterSpecification)(nil)).Elem()
}

func (i GlobalParameterSpecificationMap) ToGlobalParameterSpecificationMapOutput() GlobalParameterSpecificationMapOutput {
	return i.ToGlobalParameterSpecificationMapOutputWithContext(context.Background())
}

func (i GlobalParameterSpecificationMap) ToGlobalParameterSpecificationMapOutputWithContext(ctx context.Context) GlobalParameterSpecificationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalParameterSpecificationMapOutput)
}

type GlobalParameterSpecificationOutput struct{ *pulumi.OutputState }

func (GlobalParameterSpecificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalParameterSpecification)(nil)).Elem()
}

func (o GlobalParameterSpecificationOutput) ToGlobalParameterSpecificationOutput() GlobalParameterSpecificationOutput {
	return o
}

func (o GlobalParameterSpecificationOutput) ToGlobalParameterSpecificationOutputWithContext(ctx context.Context) GlobalParameterSpecificationOutput {
	return o
}

func (o GlobalParameterSpecificationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GlobalParameterSpecification) string { return v.Type }).(pulumi.StringOutput)
}

func (o GlobalParameterSpecificationOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v GlobalParameterSpecification) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type GlobalParameterSpecificationMapOutput struct{ *pulumi.OutputState }

func (GlobalParameterSpecificationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GlobalParameterSpecification)(nil)).Elem()
}

func (o GlobalParameterSpecificationMapOutput) ToGlobalParameterSpecificationMapOutput() GlobalParameterSpecificationMapOutput {
	return o
}

func (o GlobalParameterSpecificationMapOutput) ToGlobalParameterSpecificationMapOutputWithContext(ctx context.Context) GlobalParameterSpecificationMapOutput {
	return o
}

func (o GlobalParameterSpecificationMapOutput) MapIndex(k pulumi.StringInput) GlobalParameterSpecificationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GlobalParameterSpecification {
		return vs[0].(map[string]GlobalParameterSpecification)[vs[1].(string)]
	}).(GlobalParameterSpecificationOutput)
}

type GlobalParameterSpecificationResponse struct {
	Type  string      `pulumi:"type"`
	Value interface{} `pulumi:"value"`
}

type GlobalParameterSpecificationResponseOutput struct{ *pulumi.OutputState }

func (GlobalParameterSpecificationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalParameterSpecificationResponse)(nil)).Elem()
}

func (o GlobalParameterSpecificationResponseOutput) ToGlobalParameterSpecificationResponseOutput() GlobalParameterSpecificationResponseOutput {
	return o
}

func (o GlobalParameterSpecificationResponseOutput) ToGlobalParameterSpecificationResponseOutputWithContext(ctx context.Context) GlobalParameterSpecificationResponseOutput {
	return o
}

func (o GlobalParameterSpecificationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GlobalParameterSpecificationResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o GlobalParameterSpecificationResponseOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v GlobalParameterSpecificationResponse) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type GlobalParameterSpecificationResponseMapOutput struct{ *pulumi.OutputState }

func (GlobalParameterSpecificationResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GlobalParameterSpecificationResponse)(nil)).Elem()
}

func (o GlobalParameterSpecificationResponseMapOutput) ToGlobalParameterSpecificationResponseMapOutput() GlobalParameterSpecificationResponseMapOutput {
	return o
}

func (o GlobalParameterSpecificationResponseMapOutput) ToGlobalParameterSpecificationResponseMapOutputWithContext(ctx context.Context) GlobalParameterSpecificationResponseMapOutput {
	return o
}

func (o GlobalParameterSpecificationResponseMapOutput) MapIndex(k pulumi.StringInput) GlobalParameterSpecificationResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GlobalParameterSpecificationResponse {
		return vs[0].(map[string]GlobalParameterSpecificationResponse)[vs[1].(string)]
	}).(GlobalParameterSpecificationResponseOutput)
}

type GoogleAdWordsLinkedService struct {
	Annotations          []interface{}                     `pulumi:"annotations"`
	AuthenticationType   *string                           `pulumi:"authenticationType"`
	ClientCustomerID     interface{}                       `pulumi:"clientCustomerID"`
	ClientId             interface{}                       `pulumi:"clientId"`
	ClientSecret         interface{}                       `pulumi:"clientSecret"`
	ConnectVia           *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionProperties interface{}                       `pulumi:"connectionProperties"`
	Description          *string                           `pulumi:"description"`
	DeveloperToken       interface{}                       `pulumi:"developerToken"`
	Email                interface{}                       `pulumi:"email"`
	EncryptedCredential  interface{}                       `pulumi:"encryptedCredential"`
	KeyFilePath          interface{}                       `pulumi:"keyFilePath"`
	Parameters           map[string]ParameterSpecification `pulumi:"parameters"`
	RefreshToken         interface{}                       `pulumi:"refreshToken"`
	TrustedCertPath      interface{}                       `pulumi:"trustedCertPath"`
	Type                 string                            `pulumi:"type"`
	UseSystemTrustStore  interface{}                       `pulumi:"useSystemTrustStore"`
}

type GoogleAdWordsLinkedServiceResponse struct {
	Annotations          []interface{}                             `pulumi:"annotations"`
	AuthenticationType   *string                                   `pulumi:"authenticationType"`
	ClientCustomerID     interface{}                               `pulumi:"clientCustomerID"`
	ClientId             interface{}                               `pulumi:"clientId"`
	ClientSecret         interface{}                               `pulumi:"clientSecret"`
	ConnectVia           *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionProperties interface{}                               `pulumi:"connectionProperties"`
	Description          *string                                   `pulumi:"description"`
	DeveloperToken       interface{}                               `pulumi:"developerToken"`
	Email                interface{}                               `pulumi:"email"`
	EncryptedCredential  interface{}                               `pulumi:"encryptedCredential"`
	KeyFilePath          interface{}                               `pulumi:"keyFilePath"`
	Parameters           map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	RefreshToken         interface{}                               `pulumi:"refreshToken"`
	TrustedCertPath      interface{}                               `pulumi:"trustedCertPath"`
	Type                 string                                    `pulumi:"type"`
	UseSystemTrustStore  interface{}                               `pulumi:"useSystemTrustStore"`
}

type GoogleAdWordsObjectDataset struct {
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

type GoogleAdWordsObjectDatasetResponse struct {
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

type GoogleAdWordsSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type GoogleAdWordsSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type GoogleBigQueryLinkedService struct {
	AdditionalProjects      interface{}                       `pulumi:"additionalProjects"`
	Annotations             []interface{}                     `pulumi:"annotations"`
	AuthenticationType      string                            `pulumi:"authenticationType"`
	ClientId                interface{}                       `pulumi:"clientId"`
	ClientSecret            interface{}                       `pulumi:"clientSecret"`
	ConnectVia              *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description             *string                           `pulumi:"description"`
	Email                   interface{}                       `pulumi:"email"`
	EncryptedCredential     interface{}                       `pulumi:"encryptedCredential"`
	KeyFilePath             interface{}                       `pulumi:"keyFilePath"`
	Parameters              map[string]ParameterSpecification `pulumi:"parameters"`
	Project                 interface{}                       `pulumi:"project"`
	RefreshToken            interface{}                       `pulumi:"refreshToken"`
	RequestGoogleDriveScope interface{}                       `pulumi:"requestGoogleDriveScope"`
	TrustedCertPath         interface{}                       `pulumi:"trustedCertPath"`
	Type                    string                            `pulumi:"type"`
	UseSystemTrustStore     interface{}                       `pulumi:"useSystemTrustStore"`
}

type GoogleBigQueryLinkedServiceResponse struct {
	AdditionalProjects      interface{}                               `pulumi:"additionalProjects"`
	Annotations             []interface{}                             `pulumi:"annotations"`
	AuthenticationType      string                                    `pulumi:"authenticationType"`
	ClientId                interface{}                               `pulumi:"clientId"`
	ClientSecret            interface{}                               `pulumi:"clientSecret"`
	ConnectVia              *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description             *string                                   `pulumi:"description"`
	Email                   interface{}                               `pulumi:"email"`
	EncryptedCredential     interface{}                               `pulumi:"encryptedCredential"`
	KeyFilePath             interface{}                               `pulumi:"keyFilePath"`
	Parameters              map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Project                 interface{}                               `pulumi:"project"`
	RefreshToken            interface{}                               `pulumi:"refreshToken"`
	RequestGoogleDriveScope interface{}                               `pulumi:"requestGoogleDriveScope"`
	TrustedCertPath         interface{}                               `pulumi:"trustedCertPath"`
	Type                    string                                    `pulumi:"type"`
	UseSystemTrustStore     interface{}                               `pulumi:"useSystemTrustStore"`
}

type GoogleBigQueryObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Dataset           interface{}                       `pulumi:"dataset"`
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

type GoogleBigQueryObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Dataset           interface{}                               `pulumi:"dataset"`
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

type GoogleBigQuerySource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type GoogleBigQuerySourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type GoogleCloudStorageLinkedService struct {
	AccessKeyId         interface{}                       `pulumi:"accessKeyId"`
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	SecretAccessKey     interface{}                       `pulumi:"secretAccessKey"`
	ServiceUrl          interface{}                       `pulumi:"serviceUrl"`
	Type                string                            `pulumi:"type"`
}

type GoogleCloudStorageLinkedServiceResponse struct {
	AccessKeyId         interface{}                               `pulumi:"accessKeyId"`
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	SecretAccessKey     interface{}                               `pulumi:"secretAccessKey"`
	ServiceUrl          interface{}                               `pulumi:"serviceUrl"`
	Type                string                                    `pulumi:"type"`
}

type GoogleCloudStorageLocation struct {
	BucketName interface{} `pulumi:"bucketName"`
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
	Version    interface{} `pulumi:"version"`
}

type GoogleCloudStorageLocationResponse struct {
	BucketName interface{} `pulumi:"bucketName"`
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
	Version    interface{} `pulumi:"version"`
}

type GoogleCloudStorageReadSettings struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Prefix                     interface{} `pulumi:"prefix"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type GoogleCloudStorageReadSettingsResponse struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Prefix                     interface{} `pulumi:"prefix"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type GreenplumLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Pwd                 *AzureKeyVaultSecretReference     `pulumi:"pwd"`
	Type                string                            `pulumi:"type"`
}

type GreenplumLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Pwd                 *AzureKeyVaultSecretReferenceResponse     `pulumi:"pwd"`
	Type                string                                    `pulumi:"type"`
}

type GreenplumSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type GreenplumSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type GreenplumTableDataset struct {
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

type GreenplumTableDatasetResponse struct {
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

type HBaseLinkedService struct {
	AllowHostNameCNMismatch   interface{}                       `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                       `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                     `pulumi:"annotations"`
	AuthenticationType        string                            `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description               *string                           `pulumi:"description"`
	EnableSsl                 interface{}                       `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                       `pulumi:"encryptedCredential"`
	Host                      interface{}                       `pulumi:"host"`
	HttpPath                  interface{}                       `pulumi:"httpPath"`
	Parameters                map[string]ParameterSpecification `pulumi:"parameters"`
	Password                  interface{}                       `pulumi:"password"`
	Port                      interface{}                       `pulumi:"port"`
	TrustedCertPath           interface{}                       `pulumi:"trustedCertPath"`
	Type                      string                            `pulumi:"type"`
	Username                  interface{}                       `pulumi:"username"`
}

type HBaseLinkedServiceResponse struct {
	AllowHostNameCNMismatch   interface{}                               `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                               `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                             `pulumi:"annotations"`
	AuthenticationType        string                                    `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description               *string                                   `pulumi:"description"`
	EnableSsl                 interface{}                               `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                               `pulumi:"encryptedCredential"`
	Host                      interface{}                               `pulumi:"host"`
	HttpPath                  interface{}                               `pulumi:"httpPath"`
	Parameters                map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                  interface{}                               `pulumi:"password"`
	Port                      interface{}                               `pulumi:"port"`
	TrustedCertPath           interface{}                               `pulumi:"trustedCertPath"`
	Type                      string                                    `pulumi:"type"`
	Username                  interface{}                               `pulumi:"username"`
}

type HBaseObjectDataset struct {
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

type HBaseObjectDatasetResponse struct {
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

type HBaseSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type HBaseSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type HDInsightHiveActivity struct {
	Arguments             []interface{}            `pulumi:"arguments"`
	Defines               map[string]interface{}   `pulumi:"defines"`
	DependsOn             []ActivityDependency     `pulumi:"dependsOn"`
	Description           *string                  `pulumi:"description"`
	GetDebugInfo          *string                  `pulumi:"getDebugInfo"`
	LinkedServiceName     *LinkedServiceReference  `pulumi:"linkedServiceName"`
	Name                  string                   `pulumi:"name"`
	Policy                *ActivityPolicy          `pulumi:"policy"`
	QueryTimeout          *int                     `pulumi:"queryTimeout"`
	ScriptLinkedService   *LinkedServiceReference  `pulumi:"scriptLinkedService"`
	ScriptPath            interface{}              `pulumi:"scriptPath"`
	StorageLinkedServices []LinkedServiceReference `pulumi:"storageLinkedServices"`
	Type                  string                   `pulumi:"type"`
	UserProperties        []UserProperty           `pulumi:"userProperties"`
	Variables             []interface{}            `pulumi:"variables"`
}

type HDInsightHiveActivityResponse struct {
	Arguments             []interface{}                    `pulumi:"arguments"`
	Defines               map[string]interface{}           `pulumi:"defines"`
	DependsOn             []ActivityDependencyResponse     `pulumi:"dependsOn"`
	Description           *string                          `pulumi:"description"`
	GetDebugInfo          *string                          `pulumi:"getDebugInfo"`
	LinkedServiceName     *LinkedServiceReferenceResponse  `pulumi:"linkedServiceName"`
	Name                  string                           `pulumi:"name"`
	Policy                *ActivityPolicyResponse          `pulumi:"policy"`
	QueryTimeout          *int                             `pulumi:"queryTimeout"`
	ScriptLinkedService   *LinkedServiceReferenceResponse  `pulumi:"scriptLinkedService"`
	ScriptPath            interface{}                      `pulumi:"scriptPath"`
	StorageLinkedServices []LinkedServiceReferenceResponse `pulumi:"storageLinkedServices"`
	Type                  string                           `pulumi:"type"`
	UserProperties        []UserPropertyResponse           `pulumi:"userProperties"`
	Variables             []interface{}                    `pulumi:"variables"`
}

type HDInsightLinkedService struct {
	Annotations               []interface{}                     `pulumi:"annotations"`
	ClusterUri                interface{}                       `pulumi:"clusterUri"`
	ConnectVia                *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description               *string                           `pulumi:"description"`
	EncryptedCredential       interface{}                       `pulumi:"encryptedCredential"`
	FileSystem                interface{}                       `pulumi:"fileSystem"`
	HcatalogLinkedServiceName *LinkedServiceReference           `pulumi:"hcatalogLinkedServiceName"`
	IsEspEnabled              interface{}                       `pulumi:"isEspEnabled"`
	LinkedServiceName         *LinkedServiceReference           `pulumi:"linkedServiceName"`
	Parameters                map[string]ParameterSpecification `pulumi:"parameters"`
	Password                  interface{}                       `pulumi:"password"`
	Type                      string                            `pulumi:"type"`
	UserName                  interface{}                       `pulumi:"userName"`
}

type HDInsightLinkedServiceResponse struct {
	Annotations               []interface{}                             `pulumi:"annotations"`
	ClusterUri                interface{}                               `pulumi:"clusterUri"`
	ConnectVia                *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description               *string                                   `pulumi:"description"`
	EncryptedCredential       interface{}                               `pulumi:"encryptedCredential"`
	FileSystem                interface{}                               `pulumi:"fileSystem"`
	HcatalogLinkedServiceName *LinkedServiceReferenceResponse           `pulumi:"hcatalogLinkedServiceName"`
	IsEspEnabled              interface{}                               `pulumi:"isEspEnabled"`
	LinkedServiceName         *LinkedServiceReferenceResponse           `pulumi:"linkedServiceName"`
	Parameters                map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                  interface{}                               `pulumi:"password"`
	Type                      string                                    `pulumi:"type"`
	UserName                  interface{}                               `pulumi:"userName"`
}

type HDInsightMapReduceActivity struct {
	Arguments             []interface{}            `pulumi:"arguments"`
	ClassName             interface{}              `pulumi:"className"`
	Defines               map[string]interface{}   `pulumi:"defines"`
	DependsOn             []ActivityDependency     `pulumi:"dependsOn"`
	Description           *string                  `pulumi:"description"`
	GetDebugInfo          *string                  `pulumi:"getDebugInfo"`
	JarFilePath           interface{}              `pulumi:"jarFilePath"`
	JarLibs               []interface{}            `pulumi:"jarLibs"`
	JarLinkedService      *LinkedServiceReference  `pulumi:"jarLinkedService"`
	LinkedServiceName     *LinkedServiceReference  `pulumi:"linkedServiceName"`
	Name                  string                   `pulumi:"name"`
	Policy                *ActivityPolicy          `pulumi:"policy"`
	StorageLinkedServices []LinkedServiceReference `pulumi:"storageLinkedServices"`
	Type                  string                   `pulumi:"type"`
	UserProperties        []UserProperty           `pulumi:"userProperties"`
}

type HDInsightMapReduceActivityResponse struct {
	Arguments             []interface{}                    `pulumi:"arguments"`
	ClassName             interface{}                      `pulumi:"className"`
	Defines               map[string]interface{}           `pulumi:"defines"`
	DependsOn             []ActivityDependencyResponse     `pulumi:"dependsOn"`
	Description           *string                          `pulumi:"description"`
	GetDebugInfo          *string                          `pulumi:"getDebugInfo"`
	JarFilePath           interface{}                      `pulumi:"jarFilePath"`
	JarLibs               []interface{}                    `pulumi:"jarLibs"`
	JarLinkedService      *LinkedServiceReferenceResponse  `pulumi:"jarLinkedService"`
	LinkedServiceName     *LinkedServiceReferenceResponse  `pulumi:"linkedServiceName"`
	Name                  string                           `pulumi:"name"`
	Policy                *ActivityPolicyResponse          `pulumi:"policy"`
	StorageLinkedServices []LinkedServiceReferenceResponse `pulumi:"storageLinkedServices"`
	Type                  string                           `pulumi:"type"`
	UserProperties        []UserPropertyResponse           `pulumi:"userProperties"`
}

type HDInsightOnDemandLinkedService struct {
	AdditionalLinkedServiceNames []LinkedServiceReference          `pulumi:"additionalLinkedServiceNames"`
	Annotations                  []interface{}                     `pulumi:"annotations"`
	ClusterNamePrefix            interface{}                       `pulumi:"clusterNamePrefix"`
	ClusterPassword              interface{}                       `pulumi:"clusterPassword"`
	ClusterResourceGroup         interface{}                       `pulumi:"clusterResourceGroup"`
	ClusterSize                  interface{}                       `pulumi:"clusterSize"`
	ClusterSshPassword           interface{}                       `pulumi:"clusterSshPassword"`
	ClusterSshUserName           interface{}                       `pulumi:"clusterSshUserName"`
	ClusterType                  interface{}                       `pulumi:"clusterType"`
	ClusterUserName              interface{}                       `pulumi:"clusterUserName"`
	ConnectVia                   *IntegrationRuntimeReference      `pulumi:"connectVia"`
	CoreConfiguration            interface{}                       `pulumi:"coreConfiguration"`
	Credential                   *CredentialReference              `pulumi:"credential"`
	DataNodeSize                 interface{}                       `pulumi:"dataNodeSize"`
	Description                  *string                           `pulumi:"description"`
	EncryptedCredential          interface{}                       `pulumi:"encryptedCredential"`
	HBaseConfiguration           interface{}                       `pulumi:"hBaseConfiguration"`
	HcatalogLinkedServiceName    *LinkedServiceReference           `pulumi:"hcatalogLinkedServiceName"`
	HdfsConfiguration            interface{}                       `pulumi:"hdfsConfiguration"`
	HeadNodeSize                 interface{}                       `pulumi:"headNodeSize"`
	HiveConfiguration            interface{}                       `pulumi:"hiveConfiguration"`
	HostSubscriptionId           interface{}                       `pulumi:"hostSubscriptionId"`
	LinkedServiceName            LinkedServiceReference            `pulumi:"linkedServiceName"`
	MapReduceConfiguration       interface{}                       `pulumi:"mapReduceConfiguration"`
	OozieConfiguration           interface{}                       `pulumi:"oozieConfiguration"`
	Parameters                   map[string]ParameterSpecification `pulumi:"parameters"`
	ScriptActions                []ScriptAction                    `pulumi:"scriptActions"`
	ServicePrincipalId           interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey          interface{}                       `pulumi:"servicePrincipalKey"`
	SparkVersion                 interface{}                       `pulumi:"sparkVersion"`
	StormConfiguration           interface{}                       `pulumi:"stormConfiguration"`
	SubnetName                   interface{}                       `pulumi:"subnetName"`
	Tenant                       interface{}                       `pulumi:"tenant"`
	TimeToLive                   interface{}                       `pulumi:"timeToLive"`
	Type                         string                            `pulumi:"type"`
	Version                      interface{}                       `pulumi:"version"`
	VirtualNetworkId             interface{}                       `pulumi:"virtualNetworkId"`
	YarnConfiguration            interface{}                       `pulumi:"yarnConfiguration"`
	ZookeeperNodeSize            interface{}                       `pulumi:"zookeeperNodeSize"`
}

type HDInsightOnDemandLinkedServiceResponse struct {
	AdditionalLinkedServiceNames []LinkedServiceReferenceResponse          `pulumi:"additionalLinkedServiceNames"`
	Annotations                  []interface{}                             `pulumi:"annotations"`
	ClusterNamePrefix            interface{}                               `pulumi:"clusterNamePrefix"`
	ClusterPassword              interface{}                               `pulumi:"clusterPassword"`
	ClusterResourceGroup         interface{}                               `pulumi:"clusterResourceGroup"`
	ClusterSize                  interface{}                               `pulumi:"clusterSize"`
	ClusterSshPassword           interface{}                               `pulumi:"clusterSshPassword"`
	ClusterSshUserName           interface{}                               `pulumi:"clusterSshUserName"`
	ClusterType                  interface{}                               `pulumi:"clusterType"`
	ClusterUserName              interface{}                               `pulumi:"clusterUserName"`
	ConnectVia                   *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	CoreConfiguration            interface{}                               `pulumi:"coreConfiguration"`
	Credential                   *CredentialReferenceResponse              `pulumi:"credential"`
	DataNodeSize                 interface{}                               `pulumi:"dataNodeSize"`
	Description                  *string                                   `pulumi:"description"`
	EncryptedCredential          interface{}                               `pulumi:"encryptedCredential"`
	HBaseConfiguration           interface{}                               `pulumi:"hBaseConfiguration"`
	HcatalogLinkedServiceName    *LinkedServiceReferenceResponse           `pulumi:"hcatalogLinkedServiceName"`
	HdfsConfiguration            interface{}                               `pulumi:"hdfsConfiguration"`
	HeadNodeSize                 interface{}                               `pulumi:"headNodeSize"`
	HiveConfiguration            interface{}                               `pulumi:"hiveConfiguration"`
	HostSubscriptionId           interface{}                               `pulumi:"hostSubscriptionId"`
	LinkedServiceName            LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	MapReduceConfiguration       interface{}                               `pulumi:"mapReduceConfiguration"`
	OozieConfiguration           interface{}                               `pulumi:"oozieConfiguration"`
	Parameters                   map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	ScriptActions                []ScriptActionResponse                    `pulumi:"scriptActions"`
	ServicePrincipalId           interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey          interface{}                               `pulumi:"servicePrincipalKey"`
	SparkVersion                 interface{}                               `pulumi:"sparkVersion"`
	StormConfiguration           interface{}                               `pulumi:"stormConfiguration"`
	SubnetName                   interface{}                               `pulumi:"subnetName"`
	Tenant                       interface{}                               `pulumi:"tenant"`
	TimeToLive                   interface{}                               `pulumi:"timeToLive"`
	Type                         string                                    `pulumi:"type"`
	Version                      interface{}                               `pulumi:"version"`
	VirtualNetworkId             interface{}                               `pulumi:"virtualNetworkId"`
	YarnConfiguration            interface{}                               `pulumi:"yarnConfiguration"`
	ZookeeperNodeSize            interface{}                               `pulumi:"zookeeperNodeSize"`
}

type HDInsightPigActivity struct {
	Arguments             interface{}              `pulumi:"arguments"`
	Defines               map[string]interface{}   `pulumi:"defines"`
	DependsOn             []ActivityDependency     `pulumi:"dependsOn"`
	Description           *string                  `pulumi:"description"`
	GetDebugInfo          *string                  `pulumi:"getDebugInfo"`
	LinkedServiceName     *LinkedServiceReference  `pulumi:"linkedServiceName"`
	Name                  string                   `pulumi:"name"`
	Policy                *ActivityPolicy          `pulumi:"policy"`
	ScriptLinkedService   *LinkedServiceReference  `pulumi:"scriptLinkedService"`
	ScriptPath            interface{}              `pulumi:"scriptPath"`
	StorageLinkedServices []LinkedServiceReference `pulumi:"storageLinkedServices"`
	Type                  string                   `pulumi:"type"`
	UserProperties        []UserProperty           `pulumi:"userProperties"`
}

type HDInsightPigActivityResponse struct {
	Arguments             interface{}                      `pulumi:"arguments"`
	Defines               map[string]interface{}           `pulumi:"defines"`
	DependsOn             []ActivityDependencyResponse     `pulumi:"dependsOn"`
	Description           *string                          `pulumi:"description"`
	GetDebugInfo          *string                          `pulumi:"getDebugInfo"`
	LinkedServiceName     *LinkedServiceReferenceResponse  `pulumi:"linkedServiceName"`
	Name                  string                           `pulumi:"name"`
	Policy                *ActivityPolicyResponse          `pulumi:"policy"`
	ScriptLinkedService   *LinkedServiceReferenceResponse  `pulumi:"scriptLinkedService"`
	ScriptPath            interface{}                      `pulumi:"scriptPath"`
	StorageLinkedServices []LinkedServiceReferenceResponse `pulumi:"storageLinkedServices"`
	Type                  string                           `pulumi:"type"`
	UserProperties        []UserPropertyResponse           `pulumi:"userProperties"`
}

type HDInsightSparkActivity struct {
	Arguments             []interface{}           `pulumi:"arguments"`
	ClassName             *string                 `pulumi:"className"`
	DependsOn             []ActivityDependency    `pulumi:"dependsOn"`
	Description           *string                 `pulumi:"description"`
	EntryFilePath         interface{}             `pulumi:"entryFilePath"`
	GetDebugInfo          *string                 `pulumi:"getDebugInfo"`
	LinkedServiceName     *LinkedServiceReference `pulumi:"linkedServiceName"`
	Name                  string                  `pulumi:"name"`
	Policy                *ActivityPolicy         `pulumi:"policy"`
	ProxyUser             interface{}             `pulumi:"proxyUser"`
	RootPath              interface{}             `pulumi:"rootPath"`
	SparkConfig           map[string]interface{}  `pulumi:"sparkConfig"`
	SparkJobLinkedService *LinkedServiceReference `pulumi:"sparkJobLinkedService"`
	Type                  string                  `pulumi:"type"`
	UserProperties        []UserProperty          `pulumi:"userProperties"`
}

type HDInsightSparkActivityResponse struct {
	Arguments             []interface{}                   `pulumi:"arguments"`
	ClassName             *string                         `pulumi:"className"`
	DependsOn             []ActivityDependencyResponse    `pulumi:"dependsOn"`
	Description           *string                         `pulumi:"description"`
	EntryFilePath         interface{}                     `pulumi:"entryFilePath"`
	GetDebugInfo          *string                         `pulumi:"getDebugInfo"`
	LinkedServiceName     *LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	Name                  string                          `pulumi:"name"`
	Policy                *ActivityPolicyResponse         `pulumi:"policy"`
	ProxyUser             interface{}                     `pulumi:"proxyUser"`
	RootPath              interface{}                     `pulumi:"rootPath"`
	SparkConfig           map[string]interface{}          `pulumi:"sparkConfig"`
	SparkJobLinkedService *LinkedServiceReferenceResponse `pulumi:"sparkJobLinkedService"`
	Type                  string                          `pulumi:"type"`
	UserProperties        []UserPropertyResponse          `pulumi:"userProperties"`
}

type HDInsightStreamingActivity struct {
	Arguments             []interface{}            `pulumi:"arguments"`
	Combiner              interface{}              `pulumi:"combiner"`
	CommandEnvironment    []interface{}            `pulumi:"commandEnvironment"`
	Defines               map[string]interface{}   `pulumi:"defines"`
	DependsOn             []ActivityDependency     `pulumi:"dependsOn"`
	Description           *string                  `pulumi:"description"`
	FileLinkedService     *LinkedServiceReference  `pulumi:"fileLinkedService"`
	FilePaths             []interface{}            `pulumi:"filePaths"`
	GetDebugInfo          *string                  `pulumi:"getDebugInfo"`
	Input                 interface{}              `pulumi:"input"`
	LinkedServiceName     *LinkedServiceReference  `pulumi:"linkedServiceName"`
	Mapper                interface{}              `pulumi:"mapper"`
	Name                  string                   `pulumi:"name"`
	Output                interface{}              `pulumi:"output"`
	Policy                *ActivityPolicy          `pulumi:"policy"`
	Reducer               interface{}              `pulumi:"reducer"`
	StorageLinkedServices []LinkedServiceReference `pulumi:"storageLinkedServices"`
	Type                  string                   `pulumi:"type"`
	UserProperties        []UserProperty           `pulumi:"userProperties"`
}

type HDInsightStreamingActivityResponse struct {
	Arguments             []interface{}                    `pulumi:"arguments"`
	Combiner              interface{}                      `pulumi:"combiner"`
	CommandEnvironment    []interface{}                    `pulumi:"commandEnvironment"`
	Defines               map[string]interface{}           `pulumi:"defines"`
	DependsOn             []ActivityDependencyResponse     `pulumi:"dependsOn"`
	Description           *string                          `pulumi:"description"`
	FileLinkedService     *LinkedServiceReferenceResponse  `pulumi:"fileLinkedService"`
	FilePaths             []interface{}                    `pulumi:"filePaths"`
	GetDebugInfo          *string                          `pulumi:"getDebugInfo"`
	Input                 interface{}                      `pulumi:"input"`
	LinkedServiceName     *LinkedServiceReferenceResponse  `pulumi:"linkedServiceName"`
	Mapper                interface{}                      `pulumi:"mapper"`
	Name                  string                           `pulumi:"name"`
	Output                interface{}                      `pulumi:"output"`
	Policy                *ActivityPolicyResponse          `pulumi:"policy"`
	Reducer               interface{}                      `pulumi:"reducer"`
	StorageLinkedServices []LinkedServiceReferenceResponse `pulumi:"storageLinkedServices"`
	Type                  string                           `pulumi:"type"`
	UserProperties        []UserPropertyResponse           `pulumi:"userProperties"`
}

type HdfsLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	AuthenticationType  interface{}                       `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Type                string                            `pulumi:"type"`
	Url                 interface{}                       `pulumi:"url"`
	UserName            interface{}                       `pulumi:"userName"`
}

type HdfsLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	AuthenticationType  interface{}                               `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
	Url                 interface{}                               `pulumi:"url"`
	UserName            interface{}                               `pulumi:"userName"`
}

type HdfsLocation struct {
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
}

type HdfsLocationResponse struct {
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
}

type HdfsReadSettings struct {
	DeleteFilesAfterCompletion interface{}     `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{}     `pulumi:"disableMetricsCollection"`
	DistcpSettings             *DistcpSettings `pulumi:"distcpSettings"`
	EnablePartitionDiscovery   *bool           `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{}     `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{}     `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{}     `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{}     `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{}     `pulumi:"partitionRootPath"`
	Recursive                  interface{}     `pulumi:"recursive"`
	Type                       string          `pulumi:"type"`
	WildcardFileName           interface{}     `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{}     `pulumi:"wildcardFolderPath"`
}

type HdfsReadSettingsResponse struct {
	DeleteFilesAfterCompletion interface{}             `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{}             `pulumi:"disableMetricsCollection"`
	DistcpSettings             *DistcpSettingsResponse `pulumi:"distcpSettings"`
	EnablePartitionDiscovery   *bool                   `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{}             `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{}             `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{}             `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{}             `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{}             `pulumi:"partitionRootPath"`
	Recursive                  interface{}             `pulumi:"recursive"`
	Type                       string                  `pulumi:"type"`
	WildcardFileName           interface{}             `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{}             `pulumi:"wildcardFolderPath"`
}

type HdfsSource struct {
	DisableMetricsCollection interface{}     `pulumi:"disableMetricsCollection"`
	DistcpSettings           *DistcpSettings `pulumi:"distcpSettings"`
	MaxConcurrentConnections interface{}     `pulumi:"maxConcurrentConnections"`
	Recursive                interface{}     `pulumi:"recursive"`
	SourceRetryCount         interface{}     `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}     `pulumi:"sourceRetryWait"`
	Type                     string          `pulumi:"type"`
}

type HdfsSourceResponse struct {
	DisableMetricsCollection interface{}             `pulumi:"disableMetricsCollection"`
	DistcpSettings           *DistcpSettingsResponse `pulumi:"distcpSettings"`
	MaxConcurrentConnections interface{}             `pulumi:"maxConcurrentConnections"`
	Recursive                interface{}             `pulumi:"recursive"`
	SourceRetryCount         interface{}             `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}             `pulumi:"sourceRetryWait"`
	Type                     string                  `pulumi:"type"`
}

type HiveLinkedService struct {
	AllowHostNameCNMismatch   interface{}                       `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                       `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                     `pulumi:"annotations"`
	AuthenticationType        string                            `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description               *string                           `pulumi:"description"`
	EnableSsl                 interface{}                       `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                       `pulumi:"encryptedCredential"`
	Host                      interface{}                       `pulumi:"host"`
	HttpPath                  interface{}                       `pulumi:"httpPath"`
	Parameters                map[string]ParameterSpecification `pulumi:"parameters"`
	Password                  interface{}                       `pulumi:"password"`
	Port                      interface{}                       `pulumi:"port"`
	ServerType                *string                           `pulumi:"serverType"`
	ServiceDiscoveryMode      interface{}                       `pulumi:"serviceDiscoveryMode"`
	ThriftTransportProtocol   *string                           `pulumi:"thriftTransportProtocol"`
	TrustedCertPath           interface{}                       `pulumi:"trustedCertPath"`
	Type                      string                            `pulumi:"type"`
	UseNativeQuery            interface{}                       `pulumi:"useNativeQuery"`
	UseSystemTrustStore       interface{}                       `pulumi:"useSystemTrustStore"`
	Username                  interface{}                       `pulumi:"username"`
	ZooKeeperNameSpace        interface{}                       `pulumi:"zooKeeperNameSpace"`
}

type HiveLinkedServiceResponse struct {
	AllowHostNameCNMismatch   interface{}                               `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                               `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                             `pulumi:"annotations"`
	AuthenticationType        string                                    `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description               *string                                   `pulumi:"description"`
	EnableSsl                 interface{}                               `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                               `pulumi:"encryptedCredential"`
	Host                      interface{}                               `pulumi:"host"`
	HttpPath                  interface{}                               `pulumi:"httpPath"`
	Parameters                map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                  interface{}                               `pulumi:"password"`
	Port                      interface{}                               `pulumi:"port"`
	ServerType                *string                                   `pulumi:"serverType"`
	ServiceDiscoveryMode      interface{}                               `pulumi:"serviceDiscoveryMode"`
	ThriftTransportProtocol   *string                                   `pulumi:"thriftTransportProtocol"`
	TrustedCertPath           interface{}                               `pulumi:"trustedCertPath"`
	Type                      string                                    `pulumi:"type"`
	UseNativeQuery            interface{}                               `pulumi:"useNativeQuery"`
	UseSystemTrustStore       interface{}                               `pulumi:"useSystemTrustStore"`
	Username                  interface{}                               `pulumi:"username"`
	ZooKeeperNameSpace        interface{}                               `pulumi:"zooKeeperNameSpace"`
}

type HiveObjectDataset struct {
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

type HiveObjectDatasetResponse struct {
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

type HiveSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type HiveSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type HttpDataset struct {
	AdditionalHeaders interface{}                       `pulumi:"additionalHeaders"`
	Annotations       []interface{}                     `pulumi:"annotations"`
	Compression       *DatasetCompression               `pulumi:"compression"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	Format            interface{}                       `pulumi:"format"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	RelativeUrl       interface{}                       `pulumi:"relativeUrl"`
	RequestBody       interface{}                       `pulumi:"requestBody"`
	RequestMethod     interface{}                       `pulumi:"requestMethod"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type HttpDatasetResponse struct {
	AdditionalHeaders interface{}                               `pulumi:"additionalHeaders"`
	Annotations       []interface{}                             `pulumi:"annotations"`
	Compression       *DatasetCompressionResponse               `pulumi:"compression"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	Format            interface{}                               `pulumi:"format"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	RelativeUrl       interface{}                               `pulumi:"relativeUrl"`
	RequestBody       interface{}                               `pulumi:"requestBody"`
	RequestMethod     interface{}                               `pulumi:"requestMethod"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type HttpLinkedService struct {
	Annotations                       []interface{}                     `pulumi:"annotations"`
	AuthHeaders                       interface{}                       `pulumi:"authHeaders"`
	AuthenticationType                *string                           `pulumi:"authenticationType"`
	CertThumbprint                    interface{}                       `pulumi:"certThumbprint"`
	ConnectVia                        *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description                       *string                           `pulumi:"description"`
	EmbeddedCertData                  interface{}                       `pulumi:"embeddedCertData"`
	EnableServerCertificateValidation interface{}                       `pulumi:"enableServerCertificateValidation"`
	EncryptedCredential               interface{}                       `pulumi:"encryptedCredential"`
	Parameters                        map[string]ParameterSpecification `pulumi:"parameters"`
	Password                          interface{}                       `pulumi:"password"`
	Type                              string                            `pulumi:"type"`
	Url                               interface{}                       `pulumi:"url"`
	UserName                          interface{}                       `pulumi:"userName"`
}

type HttpLinkedServiceResponse struct {
	Annotations                       []interface{}                             `pulumi:"annotations"`
	AuthHeaders                       interface{}                               `pulumi:"authHeaders"`
	AuthenticationType                *string                                   `pulumi:"authenticationType"`
	CertThumbprint                    interface{}                               `pulumi:"certThumbprint"`
	ConnectVia                        *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description                       *string                                   `pulumi:"description"`
	EmbeddedCertData                  interface{}                               `pulumi:"embeddedCertData"`
	EnableServerCertificateValidation interface{}                               `pulumi:"enableServerCertificateValidation"`
	EncryptedCredential               interface{}                               `pulumi:"encryptedCredential"`
	Parameters                        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                          interface{}                               `pulumi:"password"`
	Type                              string                                    `pulumi:"type"`
	Url                               interface{}                               `pulumi:"url"`
	UserName                          interface{}                               `pulumi:"userName"`
}

type HttpReadSettings struct {
	AdditionalHeaders        interface{} `pulumi:"additionalHeaders"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery *bool       `pulumi:"enablePartitionDiscovery"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PartitionRootPath        interface{} `pulumi:"partitionRootPath"`
	RequestBody              interface{} `pulumi:"requestBody"`
	RequestMethod            interface{} `pulumi:"requestMethod"`
	RequestTimeout           interface{} `pulumi:"requestTimeout"`
	Type                     string      `pulumi:"type"`
}

type HttpReadSettingsResponse struct {
	AdditionalHeaders        interface{} `pulumi:"additionalHeaders"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery *bool       `pulumi:"enablePartitionDiscovery"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PartitionRootPath        interface{} `pulumi:"partitionRootPath"`
	RequestBody              interface{} `pulumi:"requestBody"`
	RequestMethod            interface{} `pulumi:"requestMethod"`
	RequestTimeout           interface{} `pulumi:"requestTimeout"`
	Type                     string      `pulumi:"type"`
}

type HttpServerLocation struct {
	FileName    interface{} `pulumi:"fileName"`
	FolderPath  interface{} `pulumi:"folderPath"`
	RelativeUrl interface{} `pulumi:"relativeUrl"`
	Type        string      `pulumi:"type"`
}

type HttpServerLocationResponse struct {
	FileName    interface{} `pulumi:"fileName"`
	FolderPath  interface{} `pulumi:"folderPath"`
	RelativeUrl interface{} `pulumi:"relativeUrl"`
	Type        string      `pulumi:"type"`
}

type HttpSource struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type HttpSourceResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type HubspotLinkedService struct {
	AccessToken           interface{}                       `pulumi:"accessToken"`
	Annotations           []interface{}                     `pulumi:"annotations"`
	ClientId              interface{}                       `pulumi:"clientId"`
	ClientSecret          interface{}                       `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	RefreshToken          interface{}                       `pulumi:"refreshToken"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
}

type HubspotLinkedServiceResponse struct {
	AccessToken           interface{}                               `pulumi:"accessToken"`
	Annotations           []interface{}                             `pulumi:"annotations"`
	ClientId              interface{}                               `pulumi:"clientId"`
	ClientSecret          interface{}                               `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	RefreshToken          interface{}                               `pulumi:"refreshToken"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
}

type HubspotObjectDataset struct {
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

type HubspotObjectDatasetResponse struct {
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

type HubspotSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type HubspotSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type IfConditionActivity struct {
	DependsOn         []ActivityDependency `pulumi:"dependsOn"`
	Description       *string              `pulumi:"description"`
	Expression        Expression           `pulumi:"expression"`
	IfFalseActivities []interface{}        `pulumi:"ifFalseActivities"`
	IfTrueActivities  []interface{}        `pulumi:"ifTrueActivities"`
	Name              string               `pulumi:"name"`
	Type              string               `pulumi:"type"`
	UserProperties    []UserProperty       `pulumi:"userProperties"`
}

type IfConditionActivityResponse struct {
	DependsOn         []ActivityDependencyResponse `pulumi:"dependsOn"`
	Description       *string                      `pulumi:"description"`
	Expression        ExpressionResponse           `pulumi:"expression"`
	IfFalseActivities []interface{}                `pulumi:"ifFalseActivities"`
	IfTrueActivities  []interface{}                `pulumi:"ifTrueActivities"`
	Name              string                       `pulumi:"name"`
	Type              string                       `pulumi:"type"`
	UserProperties    []UserPropertyResponse       `pulumi:"userProperties"`
}

type ImpalaLinkedService struct {
	AllowHostNameCNMismatch   interface{}                       `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                       `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                     `pulumi:"annotations"`
	AuthenticationType        string                            `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description               *string                           `pulumi:"description"`
	EnableSsl                 interface{}                       `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                       `pulumi:"encryptedCredential"`
	Host                      interface{}                       `pulumi:"host"`
	Parameters                map[string]ParameterSpecification `pulumi:"parameters"`
	Password                  interface{}                       `pulumi:"password"`
	Port                      interface{}                       `pulumi:"port"`
	TrustedCertPath           interface{}                       `pulumi:"trustedCertPath"`
	Type                      string                            `pulumi:"type"`
	UseSystemTrustStore       interface{}                       `pulumi:"useSystemTrustStore"`
	Username                  interface{}                       `pulumi:"username"`
}

type ImpalaLinkedServiceResponse struct {
	AllowHostNameCNMismatch   interface{}                               `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                               `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                             `pulumi:"annotations"`
	AuthenticationType        string                                    `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description               *string                                   `pulumi:"description"`
	EnableSsl                 interface{}                               `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                               `pulumi:"encryptedCredential"`
	Host                      interface{}                               `pulumi:"host"`
	Parameters                map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                  interface{}                               `pulumi:"password"`
	Port                      interface{}                               `pulumi:"port"`
	TrustedCertPath           interface{}                               `pulumi:"trustedCertPath"`
	Type                      string                                    `pulumi:"type"`
	UseSystemTrustStore       interface{}                               `pulumi:"useSystemTrustStore"`
	Username                  interface{}                               `pulumi:"username"`
}

type ImpalaObjectDataset struct {
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

type ImpalaObjectDatasetResponse struct {
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

type ImpalaSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type ImpalaSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type InformixLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	AuthenticationType  interface{}                       `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Credential          interface{}                       `pulumi:"credential"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Type                string                            `pulumi:"type"`
	UserName            interface{}                       `pulumi:"userName"`
}

type InformixLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	AuthenticationType  interface{}                               `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Credential          interface{}                               `pulumi:"credential"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
	UserName            interface{}                               `pulumi:"userName"`
}

type InformixSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{} `pulumi:"preCopyScript"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type InformixSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{} `pulumi:"preCopyScript"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type InformixSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type InformixSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type InformixTableDataset struct {
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

type InformixTableDatasetResponse struct {
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

type IntegrationRuntimeComputeProperties struct {
	DataFlowProperties           *IntegrationRuntimeDataFlowProperties `pulumi:"dataFlowProperties"`
	Location                     *string                               `pulumi:"location"`
	MaxParallelExecutionsPerNode *int                                  `pulumi:"maxParallelExecutionsPerNode"`
	NodeSize                     *string                               `pulumi:"nodeSize"`
	NumberOfNodes                *int                                  `pulumi:"numberOfNodes"`
	VNetProperties               *IntegrationRuntimeVNetProperties     `pulumi:"vNetProperties"`
}

type IntegrationRuntimeComputePropertiesResponse struct {
	DataFlowProperties           *IntegrationRuntimeDataFlowPropertiesResponse `pulumi:"dataFlowProperties"`
	Location                     *string                                       `pulumi:"location"`
	MaxParallelExecutionsPerNode *int                                          `pulumi:"maxParallelExecutionsPerNode"`
	NodeSize                     *string                                       `pulumi:"nodeSize"`
	NumberOfNodes                *int                                          `pulumi:"numberOfNodes"`
	VNetProperties               *IntegrationRuntimeVNetPropertiesResponse     `pulumi:"vNetProperties"`
}

type IntegrationRuntimeCustomSetupScriptProperties struct {
	BlobContainerUri *string       `pulumi:"blobContainerUri"`
	SasToken         *SecureString `pulumi:"sasToken"`
}

type IntegrationRuntimeCustomSetupScriptPropertiesResponse struct {
	BlobContainerUri *string               `pulumi:"blobContainerUri"`
	SasToken         *SecureStringResponse `pulumi:"sasToken"`
}

type IntegrationRuntimeCustomerVirtualNetwork struct {
	SubnetId *string `pulumi:"subnetId"`
}

type IntegrationRuntimeCustomerVirtualNetworkResponse struct {
	SubnetId *string `pulumi:"subnetId"`
}

type IntegrationRuntimeDataFlowProperties struct {
	Cleanup     *bool   `pulumi:"cleanup"`
	ComputeType *string `pulumi:"computeType"`
	CoreCount   *int    `pulumi:"coreCount"`
	TimeToLive  *int    `pulumi:"timeToLive"`
}

type IntegrationRuntimeDataFlowPropertiesResponse struct {
	Cleanup     *bool   `pulumi:"cleanup"`
	ComputeType *string `pulumi:"computeType"`
	CoreCount   *int    `pulumi:"coreCount"`
	TimeToLive  *int    `pulumi:"timeToLive"`
}

type IntegrationRuntimeDataProxyProperties struct {
	ConnectVia           *EntityReference `pulumi:"connectVia"`
	Path                 *string          `pulumi:"path"`
	StagingLinkedService *EntityReference `pulumi:"stagingLinkedService"`
}

type IntegrationRuntimeDataProxyPropertiesResponse struct {
	ConnectVia           *EntityReferenceResponse `pulumi:"connectVia"`
	Path                 *string                  `pulumi:"path"`
	StagingLinkedService *EntityReferenceResponse `pulumi:"stagingLinkedService"`
}

type IntegrationRuntimeReference struct {
	Parameters    map[string]interface{} `pulumi:"parameters"`
	ReferenceName string                 `pulumi:"referenceName"`
	Type          string                 `pulumi:"type"`
}

type IntegrationRuntimeReferenceResponse struct {
	Parameters    map[string]interface{} `pulumi:"parameters"`
	ReferenceName string                 `pulumi:"referenceName"`
	Type          string                 `pulumi:"type"`
}

type IntegrationRuntimeSsisCatalogInfo struct {
	CatalogAdminPassword  *SecureString `pulumi:"catalogAdminPassword"`
	CatalogAdminUserName  *string       `pulumi:"catalogAdminUserName"`
	CatalogPricingTier    *string       `pulumi:"catalogPricingTier"`
	CatalogServerEndpoint *string       `pulumi:"catalogServerEndpoint"`
	DualStandbyPairName   *string       `pulumi:"dualStandbyPairName"`
}

type IntegrationRuntimeSsisCatalogInfoResponse struct {
	CatalogAdminPassword  *SecureStringResponse `pulumi:"catalogAdminPassword"`
	CatalogAdminUserName  *string               `pulumi:"catalogAdminUserName"`
	CatalogPricingTier    *string               `pulumi:"catalogPricingTier"`
	CatalogServerEndpoint *string               `pulumi:"catalogServerEndpoint"`
	DualStandbyPairName   *string               `pulumi:"dualStandbyPairName"`
}

type IntegrationRuntimeSsisProperties struct {
	CatalogInfo                  *IntegrationRuntimeSsisCatalogInfo             `pulumi:"catalogInfo"`
	Credential                   *CredentialReference                           `pulumi:"credential"`
	CustomSetupScriptProperties  *IntegrationRuntimeCustomSetupScriptProperties `pulumi:"customSetupScriptProperties"`
	DataProxyProperties          *IntegrationRuntimeDataProxyProperties         `pulumi:"dataProxyProperties"`
	Edition                      *string                                        `pulumi:"edition"`
	ExpressCustomSetupProperties []interface{}                                  `pulumi:"expressCustomSetupProperties"`
	LicenseType                  *string                                        `pulumi:"licenseType"`
	PackageStores                []PackageStore                                 `pulumi:"packageStores"`
}

type IntegrationRuntimeSsisPropertiesResponse struct {
	CatalogInfo                  *IntegrationRuntimeSsisCatalogInfoResponse             `pulumi:"catalogInfo"`
	Credential                   *CredentialReferenceResponse                           `pulumi:"credential"`
	CustomSetupScriptProperties  *IntegrationRuntimeCustomSetupScriptPropertiesResponse `pulumi:"customSetupScriptProperties"`
	DataProxyProperties          *IntegrationRuntimeDataProxyPropertiesResponse         `pulumi:"dataProxyProperties"`
	Edition                      *string                                                `pulumi:"edition"`
	ExpressCustomSetupProperties []interface{}                                          `pulumi:"expressCustomSetupProperties"`
	LicenseType                  *string                                                `pulumi:"licenseType"`
	PackageStores                []PackageStoreResponse                                 `pulumi:"packageStores"`
}

type IntegrationRuntimeVNetProperties struct {
	PublicIPs []string `pulumi:"publicIPs"`
	Subnet    *string  `pulumi:"subnet"`
	SubnetId  *string  `pulumi:"subnetId"`
	VNetId    *string  `pulumi:"vNetId"`
}

type IntegrationRuntimeVNetPropertiesResponse struct {
	PublicIPs []string `pulumi:"publicIPs"`
	Subnet    *string  `pulumi:"subnet"`
	SubnetId  *string  `pulumi:"subnetId"`
	VNetId    *string  `pulumi:"vNetId"`
}

type JiraLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Host                  interface{}                       `pulumi:"host"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Password              interface{}                       `pulumi:"password"`
	Port                  interface{}                       `pulumi:"port"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
	Username              interface{}                       `pulumi:"username"`
}

type JiraLinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Host                  interface{}                               `pulumi:"host"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password              interface{}                               `pulumi:"password"`
	Port                  interface{}                               `pulumi:"port"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
	Username              interface{}                               `pulumi:"username"`
}

type JiraObjectDataset struct {
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

type JiraObjectDatasetResponse struct {
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

type JiraSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type JiraSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type JsonDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Compression       *DatasetCompression               `pulumi:"compression"`
	Description       *string                           `pulumi:"description"`
	EncodingName      interface{}                       `pulumi:"encodingName"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Location          interface{}                       `pulumi:"location"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type JsonDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Compression       *DatasetCompressionResponse               `pulumi:"compression"`
	Description       *string                                   `pulumi:"description"`
	EncodingName      interface{}                               `pulumi:"encodingName"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Location          interface{}                               `pulumi:"location"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type JsonFormat struct {
	Deserializer       interface{} `pulumi:"deserializer"`
	EncodingName       interface{} `pulumi:"encodingName"`
	FilePattern        interface{} `pulumi:"filePattern"`
	JsonNodeReference  interface{} `pulumi:"jsonNodeReference"`
	JsonPathDefinition interface{} `pulumi:"jsonPathDefinition"`
	NestingSeparator   interface{} `pulumi:"nestingSeparator"`
	Serializer         interface{} `pulumi:"serializer"`
	Type               string      `pulumi:"type"`
}

type JsonFormatResponse struct {
	Deserializer       interface{} `pulumi:"deserializer"`
	EncodingName       interface{} `pulumi:"encodingName"`
	FilePattern        interface{} `pulumi:"filePattern"`
	JsonNodeReference  interface{} `pulumi:"jsonNodeReference"`
	JsonPathDefinition interface{} `pulumi:"jsonPathDefinition"`
	NestingSeparator   interface{} `pulumi:"nestingSeparator"`
	Serializer         interface{} `pulumi:"serializer"`
	Type               string      `pulumi:"type"`
}

type JsonReadSettings struct {
	CompressionProperties interface{} `pulumi:"compressionProperties"`
	Type                  string      `pulumi:"type"`
}

type JsonReadSettingsResponse struct {
	CompressionProperties interface{} `pulumi:"compressionProperties"`
	Type                  string      `pulumi:"type"`
}

type JsonSink struct {
	DisableMetricsCollection interface{}        `pulumi:"disableMetricsCollection"`
	FormatSettings           *JsonWriteSettings `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}        `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{}        `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}        `pulumi:"sinkRetryWait"`
	StoreSettings            interface{}        `pulumi:"storeSettings"`
	Type                     string             `pulumi:"type"`
	WriteBatchSize           interface{}        `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}        `pulumi:"writeBatchTimeout"`
}

type JsonSinkResponse struct {
	DisableMetricsCollection interface{}                `pulumi:"disableMetricsCollection"`
	FormatSettings           *JsonWriteSettingsResponse `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}                `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{}                `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}                `pulumi:"sinkRetryWait"`
	StoreSettings            interface{}                `pulumi:"storeSettings"`
	Type                     string                     `pulumi:"type"`
	WriteBatchSize           interface{}                `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}                `pulumi:"writeBatchTimeout"`
}

type JsonSource struct {
	AdditionalColumns        interface{}       `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}       `pulumi:"disableMetricsCollection"`
	FormatSettings           *JsonReadSettings `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}       `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{}       `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}       `pulumi:"sourceRetryWait"`
	StoreSettings            interface{}       `pulumi:"storeSettings"`
	Type                     string            `pulumi:"type"`
}

type JsonSourceResponse struct {
	AdditionalColumns        interface{}               `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}               `pulumi:"disableMetricsCollection"`
	FormatSettings           *JsonReadSettingsResponse `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}               `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{}               `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}               `pulumi:"sourceRetryWait"`
	StoreSettings            interface{}               `pulumi:"storeSettings"`
	Type                     string                    `pulumi:"type"`
}

type JsonWriteSettings struct {
	FilePattern interface{} `pulumi:"filePattern"`
	Type        string      `pulumi:"type"`
}

type JsonWriteSettingsResponse struct {
	FilePattern interface{} `pulumi:"filePattern"`
	Type        string      `pulumi:"type"`
}

type LinkedIntegrationRuntimeKeyAuthorization struct {
	AuthorizationType string       `pulumi:"authorizationType"`
	Key               SecureString `pulumi:"key"`
}

type LinkedIntegrationRuntimeKeyAuthorizationResponse struct {
	AuthorizationType string               `pulumi:"authorizationType"`
	Key               SecureStringResponse `pulumi:"key"`
}

type LinkedIntegrationRuntimeRbacAuthorization struct {
	AuthorizationType string               `pulumi:"authorizationType"`
	Credential        *CredentialReference `pulumi:"credential"`
	ResourceId        string               `pulumi:"resourceId"`
}

type LinkedIntegrationRuntimeRbacAuthorizationResponse struct {
	AuthorizationType string                       `pulumi:"authorizationType"`
	Credential        *CredentialReferenceResponse `pulumi:"credential"`
	ResourceId        string                       `pulumi:"resourceId"`
}

type LinkedIntegrationRuntimeResponse struct {
	CreateTime          string `pulumi:"createTime"`
	DataFactoryLocation string `pulumi:"dataFactoryLocation"`
	DataFactoryName     string `pulumi:"dataFactoryName"`
	Name                string `pulumi:"name"`
	SubscriptionId      string `pulumi:"subscriptionId"`
}

type LinkedServiceReference struct {
	Parameters    map[string]interface{} `pulumi:"parameters"`
	ReferenceName string                 `pulumi:"referenceName"`
	Type          string                 `pulumi:"type"`
}

type LinkedServiceReferenceResponse struct {
	Parameters    map[string]interface{} `pulumi:"parameters"`
	ReferenceName string                 `pulumi:"referenceName"`
	Type          string                 `pulumi:"type"`
}

type LogLocationSettings struct {
	LinkedServiceName LinkedServiceReference `pulumi:"linkedServiceName"`
	Path              interface{}            `pulumi:"path"`
}

type LogLocationSettingsResponse struct {
	LinkedServiceName LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	Path              interface{}                    `pulumi:"path"`
}

type LogSettings struct {
	CopyActivityLogSettings *CopyActivityLogSettings `pulumi:"copyActivityLogSettings"`
	EnableCopyActivityLog   interface{}              `pulumi:"enableCopyActivityLog"`
	LogLocationSettings     LogLocationSettings      `pulumi:"logLocationSettings"`
}

type LogSettingsResponse struct {
	CopyActivityLogSettings *CopyActivityLogSettingsResponse `pulumi:"copyActivityLogSettings"`
	EnableCopyActivityLog   interface{}                      `pulumi:"enableCopyActivityLog"`
	LogLocationSettings     LogLocationSettingsResponse      `pulumi:"logLocationSettings"`
}

type LogStorageSettings struct {
	EnableReliableLogging interface{}            `pulumi:"enableReliableLogging"`
	LinkedServiceName     LinkedServiceReference `pulumi:"linkedServiceName"`
	LogLevel              interface{}            `pulumi:"logLevel"`
	Path                  interface{}            `pulumi:"path"`
}

type LogStorageSettingsResponse struct {
	EnableReliableLogging interface{}                    `pulumi:"enableReliableLogging"`
	LinkedServiceName     LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	LogLevel              interface{}                    `pulumi:"logLevel"`
	Path                  interface{}                    `pulumi:"path"`
}

type LookupActivity struct {
	Dataset           DatasetReference        `pulumi:"dataset"`
	DependsOn         []ActivityDependency    `pulumi:"dependsOn"`
	Description       *string                 `pulumi:"description"`
	FirstRowOnly      interface{}             `pulumi:"firstRowOnly"`
	LinkedServiceName *LinkedServiceReference `pulumi:"linkedServiceName"`
	Name              string                  `pulumi:"name"`
	Policy            *ActivityPolicy         `pulumi:"policy"`
	Source            interface{}             `pulumi:"source"`
	Type              string                  `pulumi:"type"`
	UserProperties    []UserProperty          `pulumi:"userProperties"`
}

type LookupActivityResponse struct {
	Dataset           DatasetReferenceResponse        `pulumi:"dataset"`
	DependsOn         []ActivityDependencyResponse    `pulumi:"dependsOn"`
	Description       *string                         `pulumi:"description"`
	FirstRowOnly      interface{}                     `pulumi:"firstRowOnly"`
	LinkedServiceName *LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	Name              string                          `pulumi:"name"`
	Policy            *ActivityPolicyResponse         `pulumi:"policy"`
	Source            interface{}                     `pulumi:"source"`
	Type              string                          `pulumi:"type"`
	UserProperties    []UserPropertyResponse          `pulumi:"userProperties"`
}

type MagentoLinkedService struct {
	AccessToken           interface{}                       `pulumi:"accessToken"`
	Annotations           []interface{}                     `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Host                  interface{}                       `pulumi:"host"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
}

type MagentoLinkedServiceResponse struct {
	AccessToken           interface{}                               `pulumi:"accessToken"`
	Annotations           []interface{}                             `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Host                  interface{}                               `pulumi:"host"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
}

type MagentoObjectDataset struct {
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

type MagentoObjectDatasetResponse struct {
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

type MagentoSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MagentoSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type ManagedIntegrationRuntime struct {
	ComputeProperties      *IntegrationRuntimeComputeProperties      `pulumi:"computeProperties"`
	CustomerVirtualNetwork *IntegrationRuntimeCustomerVirtualNetwork `pulumi:"customerVirtualNetwork"`
	Description            *string                                   `pulumi:"description"`
	ManagedVirtualNetwork  *ManagedVirtualNetworkReference           `pulumi:"managedVirtualNetwork"`
	SsisProperties         *IntegrationRuntimeSsisProperties         `pulumi:"ssisProperties"`
	Type                   string                                    `pulumi:"type"`
}

type ManagedIntegrationRuntimeErrorResponse struct {
	Code       string   `pulumi:"code"`
	Message    string   `pulumi:"message"`
	Parameters []string `pulumi:"parameters"`
	Time       string   `pulumi:"time"`
}

type ManagedIntegrationRuntimeNodeResponse struct {
	Errors []ManagedIntegrationRuntimeErrorResponse `pulumi:"errors"`
	NodeId string                                   `pulumi:"nodeId"`
	Status string                                   `pulumi:"status"`
}

type ManagedIntegrationRuntimeOperationResultResponse struct {
	ActivityId string   `pulumi:"activityId"`
	ErrorCode  string   `pulumi:"errorCode"`
	Parameters []string `pulumi:"parameters"`
	Result     string   `pulumi:"result"`
	StartTime  string   `pulumi:"startTime"`
	Type       string   `pulumi:"type"`
}

type ManagedIntegrationRuntimeResponse struct {
	ComputeProperties      *IntegrationRuntimeComputePropertiesResponse      `pulumi:"computeProperties"`
	CustomerVirtualNetwork *IntegrationRuntimeCustomerVirtualNetworkResponse `pulumi:"customerVirtualNetwork"`
	Description            *string                                           `pulumi:"description"`
	ManagedVirtualNetwork  *ManagedVirtualNetworkReferenceResponse           `pulumi:"managedVirtualNetwork"`
	SsisProperties         *IntegrationRuntimeSsisPropertiesResponse         `pulumi:"ssisProperties"`
	State                  string                                            `pulumi:"state"`
	Type                   string                                            `pulumi:"type"`
}

type ManagedIntegrationRuntimeStatusResponse struct {
	CreateTime      string                                           `pulumi:"createTime"`
	DataFactoryName string                                           `pulumi:"dataFactoryName"`
	LastOperation   ManagedIntegrationRuntimeOperationResultResponse `pulumi:"lastOperation"`
	Nodes           []ManagedIntegrationRuntimeNodeResponse          `pulumi:"nodes"`
	OtherErrors     []ManagedIntegrationRuntimeErrorResponse         `pulumi:"otherErrors"`
	State           string                                           `pulumi:"state"`
	Type            string                                           `pulumi:"type"`
}

type ManagedPrivateEndpointType struct {
	Fqdns                 []string `pulumi:"fqdns"`
	GroupId               *string  `pulumi:"groupId"`
	PrivateLinkResourceId *string  `pulumi:"privateLinkResourceId"`
}





type ManagedPrivateEndpointTypeInput interface {
	pulumi.Input

	ToManagedPrivateEndpointTypeOutput() ManagedPrivateEndpointTypeOutput
	ToManagedPrivateEndpointTypeOutputWithContext(context.Context) ManagedPrivateEndpointTypeOutput
}

type ManagedPrivateEndpointTypeArgs struct {
	Fqdns                 pulumi.StringArrayInput `pulumi:"fqdns"`
	GroupId               pulumi.StringPtrInput   `pulumi:"groupId"`
	PrivateLinkResourceId pulumi.StringPtrInput   `pulumi:"privateLinkResourceId"`
}

func (ManagedPrivateEndpointTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPrivateEndpointType)(nil)).Elem()
}

func (i ManagedPrivateEndpointTypeArgs) ToManagedPrivateEndpointTypeOutput() ManagedPrivateEndpointTypeOutput {
	return i.ToManagedPrivateEndpointTypeOutputWithContext(context.Background())
}

func (i ManagedPrivateEndpointTypeArgs) ToManagedPrivateEndpointTypeOutputWithContext(ctx context.Context) ManagedPrivateEndpointTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPrivateEndpointTypeOutput)
}

type ManagedPrivateEndpointTypeOutput struct{ *pulumi.OutputState }

func (ManagedPrivateEndpointTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPrivateEndpointType)(nil)).Elem()
}

func (o ManagedPrivateEndpointTypeOutput) ToManagedPrivateEndpointTypeOutput() ManagedPrivateEndpointTypeOutput {
	return o
}

func (o ManagedPrivateEndpointTypeOutput) ToManagedPrivateEndpointTypeOutputWithContext(ctx context.Context) ManagedPrivateEndpointTypeOutput {
	return o
}

func (o ManagedPrivateEndpointTypeOutput) Fqdns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedPrivateEndpointType) []string { return v.Fqdns }).(pulumi.StringArrayOutput)
}

func (o ManagedPrivateEndpointTypeOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedPrivateEndpointType) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o ManagedPrivateEndpointTypeOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedPrivateEndpointType) *string { return v.PrivateLinkResourceId }).(pulumi.StringPtrOutput)
}

type ManagedPrivateEndpointResponse struct {
	ConnectionState       *ConnectionStatePropertiesResponse `pulumi:"connectionState"`
	Fqdns                 []string                           `pulumi:"fqdns"`
	GroupId               *string                            `pulumi:"groupId"`
	IsReserved            bool                               `pulumi:"isReserved"`
	PrivateLinkResourceId *string                            `pulumi:"privateLinkResourceId"`
	ProvisioningState     string                             `pulumi:"provisioningState"`
}

type ManagedPrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (ManagedPrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPrivateEndpointResponse)(nil)).Elem()
}

func (o ManagedPrivateEndpointResponseOutput) ToManagedPrivateEndpointResponseOutput() ManagedPrivateEndpointResponseOutput {
	return o
}

func (o ManagedPrivateEndpointResponseOutput) ToManagedPrivateEndpointResponseOutputWithContext(ctx context.Context) ManagedPrivateEndpointResponseOutput {
	return o
}

func (o ManagedPrivateEndpointResponseOutput) ConnectionState() ConnectionStatePropertiesResponsePtrOutput {
	return o.ApplyT(func(v ManagedPrivateEndpointResponse) *ConnectionStatePropertiesResponse { return v.ConnectionState }).(ConnectionStatePropertiesResponsePtrOutput)
}

func (o ManagedPrivateEndpointResponseOutput) Fqdns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedPrivateEndpointResponse) []string { return v.Fqdns }).(pulumi.StringArrayOutput)
}

func (o ManagedPrivateEndpointResponseOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedPrivateEndpointResponse) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o ManagedPrivateEndpointResponseOutput) IsReserved() pulumi.BoolOutput {
	return o.ApplyT(func(v ManagedPrivateEndpointResponse) bool { return v.IsReserved }).(pulumi.BoolOutput)
}

func (o ManagedPrivateEndpointResponseOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedPrivateEndpointResponse) *string { return v.PrivateLinkResourceId }).(pulumi.StringPtrOutput)
}

func (o ManagedPrivateEndpointResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedPrivateEndpointResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type ManagedVirtualNetworkReference struct {
	ReferenceName string `pulumi:"referenceName"`
	Type          string `pulumi:"type"`
}

type ManagedVirtualNetworkReferenceResponse struct {
	ReferenceName string `pulumi:"referenceName"`
	Type          string `pulumi:"type"`
}

type MappingDataFlow struct {
	Annotations     []interface{}    `pulumi:"annotations"`
	Description     *string          `pulumi:"description"`
	Folder          *DataFlowFolder  `pulumi:"folder"`
	Script          *string          `pulumi:"script"`
	ScriptLines     []string         `pulumi:"scriptLines"`
	Sinks           []DataFlowSink   `pulumi:"sinks"`
	Sources         []DataFlowSource `pulumi:"sources"`
	Transformations []Transformation `pulumi:"transformations"`
	Type            string           `pulumi:"type"`
}

type MappingDataFlowResponse struct {
	Annotations     []interface{}            `pulumi:"annotations"`
	Description     *string                  `pulumi:"description"`
	Folder          *DataFlowResponseFolder  `pulumi:"folder"`
	Script          *string                  `pulumi:"script"`
	ScriptLines     []string                 `pulumi:"scriptLines"`
	Sinks           []DataFlowSinkResponse   `pulumi:"sinks"`
	Sources         []DataFlowSourceResponse `pulumi:"sources"`
	Transformations []TransformationResponse `pulumi:"transformations"`
	Type            string                   `pulumi:"type"`
}

type MariaDBLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Pwd                 *AzureKeyVaultSecretReference     `pulumi:"pwd"`
	Type                string                            `pulumi:"type"`
}

type MariaDBLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Pwd                 *AzureKeyVaultSecretReferenceResponse     `pulumi:"pwd"`
	Type                string                                    `pulumi:"type"`
}

type MariaDBSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MariaDBSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MariaDBTableDataset struct {
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

type MariaDBTableDatasetResponse struct {
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

type MarketoLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	ClientId              interface{}                       `pulumi:"clientId"`
	ClientSecret          interface{}                       `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Endpoint              interface{}                       `pulumi:"endpoint"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
}

type MarketoLinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	ClientId              interface{}                               `pulumi:"clientId"`
	ClientSecret          interface{}                               `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Endpoint              interface{}                               `pulumi:"endpoint"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
}

type MarketoObjectDataset struct {
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

type MarketoObjectDatasetResponse struct {
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

type MarketoSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MarketoSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MetadataItem struct {
	Name  interface{} `pulumi:"name"`
	Value interface{} `pulumi:"value"`
}

type MetadataItemResponse struct {
	Name  interface{} `pulumi:"name"`
	Value interface{} `pulumi:"value"`
}

type MicrosoftAccessLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	AuthenticationType  interface{}                       `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Credential          interface{}                       `pulumi:"credential"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Type                string                            `pulumi:"type"`
	UserName            interface{}                       `pulumi:"userName"`
}

type MicrosoftAccessLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	AuthenticationType  interface{}                               `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Credential          interface{}                               `pulumi:"credential"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
	UserName            interface{}                               `pulumi:"userName"`
}

type MicrosoftAccessSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{} `pulumi:"preCopyScript"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type MicrosoftAccessSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{} `pulumi:"preCopyScript"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type MicrosoftAccessSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MicrosoftAccessSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MicrosoftAccessTableDataset struct {
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

type MicrosoftAccessTableDatasetResponse struct {
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

type MongoDbAtlasCollectionDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Collection        interface{}                       `pulumi:"collection"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type MongoDbAtlasCollectionDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Collection        interface{}                               `pulumi:"collection"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type MongoDbAtlasLinkedService struct {
	Annotations      []interface{}                     `pulumi:"annotations"`
	ConnectVia       *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString interface{}                       `pulumi:"connectionString"`
	Database         interface{}                       `pulumi:"database"`
	Description      *string                           `pulumi:"description"`
	Parameters       map[string]ParameterSpecification `pulumi:"parameters"`
	Type             string                            `pulumi:"type"`
}

type MongoDbAtlasLinkedServiceResponse struct {
	Annotations      []interface{}                             `pulumi:"annotations"`
	ConnectVia       *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString interface{}                               `pulumi:"connectionString"`
	Database         interface{}                               `pulumi:"database"`
	Description      *string                                   `pulumi:"description"`
	Parameters       map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type             string                                    `pulumi:"type"`
}

type MongoDbAtlasSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            interface{} `pulumi:"writeBehavior"`
}

type MongoDbAtlasSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            interface{} `pulumi:"writeBehavior"`
}

type MongoDbAtlasSource struct {
	AdditionalColumns        interface{}                     `pulumi:"additionalColumns"`
	BatchSize                interface{}                     `pulumi:"batchSize"`
	CursorMethods            *MongoDbCursorMethodsProperties `pulumi:"cursorMethods"`
	DisableMetricsCollection interface{}                     `pulumi:"disableMetricsCollection"`
	Filter                   interface{}                     `pulumi:"filter"`
	MaxConcurrentConnections interface{}                     `pulumi:"maxConcurrentConnections"`
	QueryTimeout             interface{}                     `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}                     `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                     `pulumi:"sourceRetryWait"`
	Type                     string                          `pulumi:"type"`
}

type MongoDbAtlasSourceResponse struct {
	AdditionalColumns        interface{}                             `pulumi:"additionalColumns"`
	BatchSize                interface{}                             `pulumi:"batchSize"`
	CursorMethods            *MongoDbCursorMethodsPropertiesResponse `pulumi:"cursorMethods"`
	DisableMetricsCollection interface{}                             `pulumi:"disableMetricsCollection"`
	Filter                   interface{}                             `pulumi:"filter"`
	MaxConcurrentConnections interface{}                             `pulumi:"maxConcurrentConnections"`
	QueryTimeout             interface{}                             `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}                             `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                             `pulumi:"sourceRetryWait"`
	Type                     string                                  `pulumi:"type"`
}

type MongoDbCollectionDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	CollectionName    interface{}                       `pulumi:"collectionName"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type MongoDbCollectionDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	CollectionName    interface{}                               `pulumi:"collectionName"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type MongoDbCursorMethodsProperties struct {
	Limit   interface{} `pulumi:"limit"`
	Project interface{} `pulumi:"project"`
	Skip    interface{} `pulumi:"skip"`
	Sort    interface{} `pulumi:"sort"`
}

type MongoDbCursorMethodsPropertiesResponse struct {
	Limit   interface{} `pulumi:"limit"`
	Project interface{} `pulumi:"project"`
	Skip    interface{} `pulumi:"skip"`
	Sort    interface{} `pulumi:"sort"`
}

type MongoDbLinkedService struct {
	AllowSelfSignedServerCert interface{}                       `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                     `pulumi:"annotations"`
	AuthSource                interface{}                       `pulumi:"authSource"`
	AuthenticationType        *string                           `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReference      `pulumi:"connectVia"`
	DatabaseName              interface{}                       `pulumi:"databaseName"`
	Description               *string                           `pulumi:"description"`
	EnableSsl                 interface{}                       `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                       `pulumi:"encryptedCredential"`
	Parameters                map[string]ParameterSpecification `pulumi:"parameters"`
	Password                  interface{}                       `pulumi:"password"`
	Port                      interface{}                       `pulumi:"port"`
	Server                    interface{}                       `pulumi:"server"`
	Type                      string                            `pulumi:"type"`
	Username                  interface{}                       `pulumi:"username"`
}

type MongoDbLinkedServiceResponse struct {
	AllowSelfSignedServerCert interface{}                               `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                             `pulumi:"annotations"`
	AuthSource                interface{}                               `pulumi:"authSource"`
	AuthenticationType        *string                                   `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	DatabaseName              interface{}                               `pulumi:"databaseName"`
	Description               *string                                   `pulumi:"description"`
	EnableSsl                 interface{}                               `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                               `pulumi:"encryptedCredential"`
	Parameters                map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                  interface{}                               `pulumi:"password"`
	Port                      interface{}                               `pulumi:"port"`
	Server                    interface{}                               `pulumi:"server"`
	Type                      string                                    `pulumi:"type"`
	Username                  interface{}                               `pulumi:"username"`
}

type MongoDbSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MongoDbSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MongoDbV2CollectionDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Collection        interface{}                       `pulumi:"collection"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type MongoDbV2CollectionDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Collection        interface{}                               `pulumi:"collection"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type MongoDbV2LinkedService struct {
	Annotations      []interface{}                     `pulumi:"annotations"`
	ConnectVia       *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString interface{}                       `pulumi:"connectionString"`
	Database         interface{}                       `pulumi:"database"`
	Description      *string                           `pulumi:"description"`
	Parameters       map[string]ParameterSpecification `pulumi:"parameters"`
	Type             string                            `pulumi:"type"`
}

type MongoDbV2LinkedServiceResponse struct {
	Annotations      []interface{}                             `pulumi:"annotations"`
	ConnectVia       *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString interface{}                               `pulumi:"connectionString"`
	Database         interface{}                               `pulumi:"database"`
	Description      *string                                   `pulumi:"description"`
	Parameters       map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type             string                                    `pulumi:"type"`
}

type MongoDbV2Sink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            interface{} `pulumi:"writeBehavior"`
}

type MongoDbV2SinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            interface{} `pulumi:"writeBehavior"`
}

type MongoDbV2Source struct {
	AdditionalColumns        interface{}                     `pulumi:"additionalColumns"`
	BatchSize                interface{}                     `pulumi:"batchSize"`
	CursorMethods            *MongoDbCursorMethodsProperties `pulumi:"cursorMethods"`
	DisableMetricsCollection interface{}                     `pulumi:"disableMetricsCollection"`
	Filter                   interface{}                     `pulumi:"filter"`
	MaxConcurrentConnections interface{}                     `pulumi:"maxConcurrentConnections"`
	QueryTimeout             interface{}                     `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}                     `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                     `pulumi:"sourceRetryWait"`
	Type                     string                          `pulumi:"type"`
}

type MongoDbV2SourceResponse struct {
	AdditionalColumns        interface{}                             `pulumi:"additionalColumns"`
	BatchSize                interface{}                             `pulumi:"batchSize"`
	CursorMethods            *MongoDbCursorMethodsPropertiesResponse `pulumi:"cursorMethods"`
	DisableMetricsCollection interface{}                             `pulumi:"disableMetricsCollection"`
	Filter                   interface{}                             `pulumi:"filter"`
	MaxConcurrentConnections interface{}                             `pulumi:"maxConcurrentConnections"`
	QueryTimeout             interface{}                             `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}                             `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                             `pulumi:"sourceRetryWait"`
	Type                     string                                  `pulumi:"type"`
}

type MultiplePipelineTrigger struct {
	Annotations []interface{}              `pulumi:"annotations"`
	Description *string                    `pulumi:"description"`
	Pipelines   []TriggerPipelineReference `pulumi:"pipelines"`
	Type        string                     `pulumi:"type"`
}

type MultiplePipelineTriggerResponse struct {
	Annotations  []interface{}                      `pulumi:"annotations"`
	Description  *string                            `pulumi:"description"`
	Pipelines    []TriggerPipelineReferenceResponse `pulumi:"pipelines"`
	RuntimeState string                             `pulumi:"runtimeState"`
	Type         string                             `pulumi:"type"`
}

type MySqlLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReference     `pulumi:"password"`
	Type                string                            `pulumi:"type"`
}

type MySqlLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReferenceResponse     `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
}

type MySqlSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MySqlSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MySqlTableDataset struct {
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

type MySqlTableDatasetResponse struct {
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

type NetezzaLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Pwd                 *AzureKeyVaultSecretReference     `pulumi:"pwd"`
	Type                string                            `pulumi:"type"`
}

type NetezzaLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Pwd                 *AzureKeyVaultSecretReferenceResponse     `pulumi:"pwd"`
	Type                string                                    `pulumi:"type"`
}

type NetezzaPartitionSettings struct {
	PartitionColumnName interface{} `pulumi:"partitionColumnName"`
	PartitionLowerBound interface{} `pulumi:"partitionLowerBound"`
	PartitionUpperBound interface{} `pulumi:"partitionUpperBound"`
}

type NetezzaPartitionSettingsResponse struct {
	PartitionColumnName interface{} `pulumi:"partitionColumnName"`
	PartitionLowerBound interface{} `pulumi:"partitionLowerBound"`
	PartitionUpperBound interface{} `pulumi:"partitionUpperBound"`
}

type NetezzaSource struct {
	AdditionalColumns        interface{}               `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}               `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}               `pulumi:"maxConcurrentConnections"`
	PartitionOption          interface{}               `pulumi:"partitionOption"`
	PartitionSettings        *NetezzaPartitionSettings `pulumi:"partitionSettings"`
	Query                    interface{}               `pulumi:"query"`
	QueryTimeout             interface{}               `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}               `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}               `pulumi:"sourceRetryWait"`
	Type                     string                    `pulumi:"type"`
}

type NetezzaSourceResponse struct {
	AdditionalColumns        interface{}                       `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}                       `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}                       `pulumi:"maxConcurrentConnections"`
	PartitionOption          interface{}                       `pulumi:"partitionOption"`
	PartitionSettings        *NetezzaPartitionSettingsResponse `pulumi:"partitionSettings"`
	Query                    interface{}                       `pulumi:"query"`
	QueryTimeout             interface{}                       `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}                       `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                       `pulumi:"sourceRetryWait"`
	Type                     string                            `pulumi:"type"`
}

type NetezzaTableDataset struct {
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

type NetezzaTableDatasetResponse struct {
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

type NotebookParameter struct {
	Type  *string     `pulumi:"type"`
	Value interface{} `pulumi:"value"`
}

type NotebookParameterResponse struct {
	Type  *string     `pulumi:"type"`
	Value interface{} `pulumi:"value"`
}

type ODataLinkedService struct {
	AadResourceId                        interface{}                       `pulumi:"aadResourceId"`
	AadServicePrincipalCredentialType    *string                           `pulumi:"aadServicePrincipalCredentialType"`
	Annotations                          []interface{}                     `pulumi:"annotations"`
	AuthHeaders                          interface{}                       `pulumi:"authHeaders"`
	AuthenticationType                   *string                           `pulumi:"authenticationType"`
	AzureCloudType                       interface{}                       `pulumi:"azureCloudType"`
	ConnectVia                           *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description                          *string                           `pulumi:"description"`
	EncryptedCredential                  interface{}                       `pulumi:"encryptedCredential"`
	Parameters                           map[string]ParameterSpecification `pulumi:"parameters"`
	Password                             interface{}                       `pulumi:"password"`
	ServicePrincipalEmbeddedCert         interface{}                       `pulumi:"servicePrincipalEmbeddedCert"`
	ServicePrincipalEmbeddedCertPassword interface{}                       `pulumi:"servicePrincipalEmbeddedCertPassword"`
	ServicePrincipalId                   interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey                  interface{}                       `pulumi:"servicePrincipalKey"`
	Tenant                               interface{}                       `pulumi:"tenant"`
	Type                                 string                            `pulumi:"type"`
	Url                                  interface{}                       `pulumi:"url"`
	UserName                             interface{}                       `pulumi:"userName"`
}

type ODataLinkedServiceResponse struct {
	AadResourceId                        interface{}                               `pulumi:"aadResourceId"`
	AadServicePrincipalCredentialType    *string                                   `pulumi:"aadServicePrincipalCredentialType"`
	Annotations                          []interface{}                             `pulumi:"annotations"`
	AuthHeaders                          interface{}                               `pulumi:"authHeaders"`
	AuthenticationType                   *string                                   `pulumi:"authenticationType"`
	AzureCloudType                       interface{}                               `pulumi:"azureCloudType"`
	ConnectVia                           *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description                          *string                                   `pulumi:"description"`
	EncryptedCredential                  interface{}                               `pulumi:"encryptedCredential"`
	Parameters                           map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                             interface{}                               `pulumi:"password"`
	ServicePrincipalEmbeddedCert         interface{}                               `pulumi:"servicePrincipalEmbeddedCert"`
	ServicePrincipalEmbeddedCertPassword interface{}                               `pulumi:"servicePrincipalEmbeddedCertPassword"`
	ServicePrincipalId                   interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey                  interface{}                               `pulumi:"servicePrincipalKey"`
	Tenant                               interface{}                               `pulumi:"tenant"`
	Type                                 string                                    `pulumi:"type"`
	Url                                  interface{}                               `pulumi:"url"`
	UserName                             interface{}                               `pulumi:"userName"`
}

type ODataResourceDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Path              interface{}                       `pulumi:"path"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type ODataResourceDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Path              interface{}                               `pulumi:"path"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type ODataSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type ODataSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type OdbcLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	AuthenticationType  interface{}                       `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Credential          interface{}                       `pulumi:"credential"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Type                string                            `pulumi:"type"`
	UserName            interface{}                       `pulumi:"userName"`
}

type OdbcLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	AuthenticationType  interface{}                               `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Credential          interface{}                               `pulumi:"credential"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
	UserName            interface{}                               `pulumi:"userName"`
}

type OdbcSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{} `pulumi:"preCopyScript"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type OdbcSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{} `pulumi:"preCopyScript"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type OdbcSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type OdbcSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type OdbcTableDataset struct {
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

type OdbcTableDatasetResponse struct {
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

type Office365Dataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Predicate         interface{}                       `pulumi:"predicate"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type Office365DatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Predicate         interface{}                               `pulumi:"predicate"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type Office365LinkedService struct {
	Annotations              []interface{}                     `pulumi:"annotations"`
	ConnectVia               *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description              *string                           `pulumi:"description"`
	EncryptedCredential      interface{}                       `pulumi:"encryptedCredential"`
	Office365TenantId        interface{}                       `pulumi:"office365TenantId"`
	Parameters               map[string]ParameterSpecification `pulumi:"parameters"`
	ServicePrincipalId       interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey      interface{}                       `pulumi:"servicePrincipalKey"`
	ServicePrincipalTenantId interface{}                       `pulumi:"servicePrincipalTenantId"`
	Type                     string                            `pulumi:"type"`
}

type Office365LinkedServiceResponse struct {
	Annotations              []interface{}                             `pulumi:"annotations"`
	ConnectVia               *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description              *string                                   `pulumi:"description"`
	EncryptedCredential      interface{}                               `pulumi:"encryptedCredential"`
	Office365TenantId        interface{}                               `pulumi:"office365TenantId"`
	Parameters               map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	ServicePrincipalId       interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey      interface{}                               `pulumi:"servicePrincipalKey"`
	ServicePrincipalTenantId interface{}                               `pulumi:"servicePrincipalTenantId"`
	Type                     string                                    `pulumi:"type"`
}

type Office365Source struct {
	AllowedGroups            interface{} `pulumi:"allowedGroups"`
	DateFilterColumn         interface{} `pulumi:"dateFilterColumn"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	EndTime                  interface{} `pulumi:"endTime"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	OutputColumns            interface{} `pulumi:"outputColumns"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	StartTime                interface{} `pulumi:"startTime"`
	Type                     string      `pulumi:"type"`
	UserScopeFilterUri       interface{} `pulumi:"userScopeFilterUri"`
}

type Office365SourceResponse struct {
	AllowedGroups            interface{} `pulumi:"allowedGroups"`
	DateFilterColumn         interface{} `pulumi:"dateFilterColumn"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	EndTime                  interface{} `pulumi:"endTime"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	OutputColumns            interface{} `pulumi:"outputColumns"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	StartTime                interface{} `pulumi:"startTime"`
	Type                     string      `pulumi:"type"`
	UserScopeFilterUri       interface{} `pulumi:"userScopeFilterUri"`
}

type OracleCloudStorageLinkedService struct {
	AccessKeyId         interface{}                       `pulumi:"accessKeyId"`
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	SecretAccessKey     interface{}                       `pulumi:"secretAccessKey"`
	ServiceUrl          interface{}                       `pulumi:"serviceUrl"`
	Type                string                            `pulumi:"type"`
}

type OracleCloudStorageLinkedServiceResponse struct {
	AccessKeyId         interface{}                               `pulumi:"accessKeyId"`
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	SecretAccessKey     interface{}                               `pulumi:"secretAccessKey"`
	ServiceUrl          interface{}                               `pulumi:"serviceUrl"`
	Type                string                                    `pulumi:"type"`
}

type OracleCloudStorageLocation struct {
	BucketName interface{} `pulumi:"bucketName"`
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
	Version    interface{} `pulumi:"version"`
}

type OracleCloudStorageLocationResponse struct {
	BucketName interface{} `pulumi:"bucketName"`
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
	Version    interface{} `pulumi:"version"`
}

type OracleCloudStorageReadSettings struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Prefix                     interface{} `pulumi:"prefix"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type OracleCloudStorageReadSettingsResponse struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Prefix                     interface{} `pulumi:"prefix"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type OracleLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReference     `pulumi:"password"`
	Type                string                            `pulumi:"type"`
}

type OracleLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReferenceResponse     `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
}

type OraclePartitionSettings struct {
	PartitionColumnName interface{}   `pulumi:"partitionColumnName"`
	PartitionLowerBound interface{}   `pulumi:"partitionLowerBound"`
	PartitionNames      []interface{} `pulumi:"partitionNames"`
	PartitionUpperBound interface{}   `pulumi:"partitionUpperBound"`
}

type OraclePartitionSettingsResponse struct {
	PartitionColumnName interface{}   `pulumi:"partitionColumnName"`
	PartitionLowerBound interface{}   `pulumi:"partitionLowerBound"`
	PartitionNames      []interface{} `pulumi:"partitionNames"`
	PartitionUpperBound interface{}   `pulumi:"partitionUpperBound"`
}

type OracleServiceCloudLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Host                  interface{}                       `pulumi:"host"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Password              interface{}                       `pulumi:"password"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
	Username              interface{}                       `pulumi:"username"`
}

type OracleServiceCloudLinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Host                  interface{}                               `pulumi:"host"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password              interface{}                               `pulumi:"password"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
	Username              interface{}                               `pulumi:"username"`
}

type OracleServiceCloudObjectDataset struct {
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

type OracleServiceCloudObjectDatasetResponse struct {
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

type OracleServiceCloudSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type OracleServiceCloudSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type OracleSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{} `pulumi:"preCopyScript"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type OracleSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{} `pulumi:"preCopyScript"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type OracleSource struct {
	AdditionalColumns        interface{}              `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}              `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}              `pulumi:"maxConcurrentConnections"`
	OracleReaderQuery        interface{}              `pulumi:"oracleReaderQuery"`
	PartitionOption          interface{}              `pulumi:"partitionOption"`
	PartitionSettings        *OraclePartitionSettings `pulumi:"partitionSettings"`
	QueryTimeout             interface{}              `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}              `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}              `pulumi:"sourceRetryWait"`
	Type                     string                   `pulumi:"type"`
}

type OracleSourceResponse struct {
	AdditionalColumns        interface{}                      `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}                      `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}                      `pulumi:"maxConcurrentConnections"`
	OracleReaderQuery        interface{}                      `pulumi:"oracleReaderQuery"`
	PartitionOption          interface{}                      `pulumi:"partitionOption"`
	PartitionSettings        *OraclePartitionSettingsResponse `pulumi:"partitionSettings"`
	QueryTimeout             interface{}                      `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}                      `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                      `pulumi:"sourceRetryWait"`
	Type                     string                           `pulumi:"type"`
}

type OracleTableDataset struct {
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

type OracleTableDatasetResponse struct {
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

type OrcDataset struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	Description         *string                           `pulumi:"description"`
	Folder              *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName   LinkedServiceReference            `pulumi:"linkedServiceName"`
	Location            interface{}                       `pulumi:"location"`
	OrcCompressionCodec interface{}                       `pulumi:"orcCompressionCodec"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Schema              interface{}                       `pulumi:"schema"`
	Structure           interface{}                       `pulumi:"structure"`
	Type                string                            `pulumi:"type"`
}

type OrcDatasetResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	Description         *string                                   `pulumi:"description"`
	Folder              *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName   LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Location            interface{}                               `pulumi:"location"`
	OrcCompressionCodec interface{}                               `pulumi:"orcCompressionCodec"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema              interface{}                               `pulumi:"schema"`
	Structure           interface{}                               `pulumi:"structure"`
	Type                string                                    `pulumi:"type"`
}

type OrcFormat struct {
	Deserializer interface{} `pulumi:"deserializer"`
	Serializer   interface{} `pulumi:"serializer"`
	Type         string      `pulumi:"type"`
}

type OrcFormatResponse struct {
	Deserializer interface{} `pulumi:"deserializer"`
	Serializer   interface{} `pulumi:"serializer"`
	Type         string      `pulumi:"type"`
}

type OrcSink struct {
	DisableMetricsCollection interface{}       `pulumi:"disableMetricsCollection"`
	FormatSettings           *OrcWriteSettings `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}       `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{}       `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}       `pulumi:"sinkRetryWait"`
	StoreSettings            interface{}       `pulumi:"storeSettings"`
	Type                     string            `pulumi:"type"`
	WriteBatchSize           interface{}       `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}       `pulumi:"writeBatchTimeout"`
}

type OrcSinkResponse struct {
	DisableMetricsCollection interface{}               `pulumi:"disableMetricsCollection"`
	FormatSettings           *OrcWriteSettingsResponse `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}               `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{}               `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}               `pulumi:"sinkRetryWait"`
	StoreSettings            interface{}               `pulumi:"storeSettings"`
	Type                     string                    `pulumi:"type"`
	WriteBatchSize           interface{}               `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}               `pulumi:"writeBatchTimeout"`
}

type OrcSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	StoreSettings            interface{} `pulumi:"storeSettings"`
	Type                     string      `pulumi:"type"`
}

type OrcSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	StoreSettings            interface{} `pulumi:"storeSettings"`
	Type                     string      `pulumi:"type"`
}

type OrcWriteSettings struct {
	FileNamePrefix interface{} `pulumi:"fileNamePrefix"`
	MaxRowsPerFile interface{} `pulumi:"maxRowsPerFile"`
	Type           string      `pulumi:"type"`
}

type OrcWriteSettingsResponse struct {
	FileNamePrefix interface{} `pulumi:"fileNamePrefix"`
	MaxRowsPerFile interface{} `pulumi:"maxRowsPerFile"`
	Type           string      `pulumi:"type"`
}

type PackageStore struct {
	Name                      string          `pulumi:"name"`
	PackageStoreLinkedService EntityReference `pulumi:"packageStoreLinkedService"`
}

type PackageStoreResponse struct {
	Name                      string                  `pulumi:"name"`
	PackageStoreLinkedService EntityReferenceResponse `pulumi:"packageStoreLinkedService"`
}

type ParameterSpecification struct {
	DefaultValue interface{} `pulumi:"defaultValue"`
	Type         string      `pulumi:"type"`
}





type ParameterSpecificationInput interface {
	pulumi.Input

	ToParameterSpecificationOutput() ParameterSpecificationOutput
	ToParameterSpecificationOutputWithContext(context.Context) ParameterSpecificationOutput
}

type ParameterSpecificationArgs struct {
	DefaultValue pulumi.Input       `pulumi:"defaultValue"`
	Type         pulumi.StringInput `pulumi:"type"`
}

func (ParameterSpecificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterSpecification)(nil)).Elem()
}

func (i ParameterSpecificationArgs) ToParameterSpecificationOutput() ParameterSpecificationOutput {
	return i.ToParameterSpecificationOutputWithContext(context.Background())
}

func (i ParameterSpecificationArgs) ToParameterSpecificationOutputWithContext(ctx context.Context) ParameterSpecificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterSpecificationOutput)
}





type ParameterSpecificationMapInput interface {
	pulumi.Input

	ToParameterSpecificationMapOutput() ParameterSpecificationMapOutput
	ToParameterSpecificationMapOutputWithContext(context.Context) ParameterSpecificationMapOutput
}

type ParameterSpecificationMap map[string]ParameterSpecificationInput

func (ParameterSpecificationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterSpecification)(nil)).Elem()
}

func (i ParameterSpecificationMap) ToParameterSpecificationMapOutput() ParameterSpecificationMapOutput {
	return i.ToParameterSpecificationMapOutputWithContext(context.Background())
}

func (i ParameterSpecificationMap) ToParameterSpecificationMapOutputWithContext(ctx context.Context) ParameterSpecificationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterSpecificationMapOutput)
}

type ParameterSpecificationOutput struct{ *pulumi.OutputState }

func (ParameterSpecificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterSpecification)(nil)).Elem()
}

func (o ParameterSpecificationOutput) ToParameterSpecificationOutput() ParameterSpecificationOutput {
	return o
}

func (o ParameterSpecificationOutput) ToParameterSpecificationOutputWithContext(ctx context.Context) ParameterSpecificationOutput {
	return o
}

func (o ParameterSpecificationOutput) DefaultValue() pulumi.AnyOutput {
	return o.ApplyT(func(v ParameterSpecification) interface{} { return v.DefaultValue }).(pulumi.AnyOutput)
}

func (o ParameterSpecificationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterSpecification) string { return v.Type }).(pulumi.StringOutput)
}

type ParameterSpecificationMapOutput struct{ *pulumi.OutputState }

func (ParameterSpecificationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterSpecification)(nil)).Elem()
}

func (o ParameterSpecificationMapOutput) ToParameterSpecificationMapOutput() ParameterSpecificationMapOutput {
	return o
}

func (o ParameterSpecificationMapOutput) ToParameterSpecificationMapOutputWithContext(ctx context.Context) ParameterSpecificationMapOutput {
	return o
}

func (o ParameterSpecificationMapOutput) MapIndex(k pulumi.StringInput) ParameterSpecificationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterSpecification {
		return vs[0].(map[string]ParameterSpecification)[vs[1].(string)]
	}).(ParameterSpecificationOutput)
}

type ParameterSpecificationResponse struct {
	DefaultValue interface{} `pulumi:"defaultValue"`
	Type         string      `pulumi:"type"`
}

type ParameterSpecificationResponseOutput struct{ *pulumi.OutputState }

func (ParameterSpecificationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterSpecificationResponse)(nil)).Elem()
}

func (o ParameterSpecificationResponseOutput) ToParameterSpecificationResponseOutput() ParameterSpecificationResponseOutput {
	return o
}

func (o ParameterSpecificationResponseOutput) ToParameterSpecificationResponseOutputWithContext(ctx context.Context) ParameterSpecificationResponseOutput {
	return o
}

func (o ParameterSpecificationResponseOutput) DefaultValue() pulumi.AnyOutput {
	return o.ApplyT(func(v ParameterSpecificationResponse) interface{} { return v.DefaultValue }).(pulumi.AnyOutput)
}

func (o ParameterSpecificationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterSpecificationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ParameterSpecificationResponseMapOutput struct{ *pulumi.OutputState }

func (ParameterSpecificationResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterSpecificationResponse)(nil)).Elem()
}

func (o ParameterSpecificationResponseMapOutput) ToParameterSpecificationResponseMapOutput() ParameterSpecificationResponseMapOutput {
	return o
}

func (o ParameterSpecificationResponseMapOutput) ToParameterSpecificationResponseMapOutputWithContext(ctx context.Context) ParameterSpecificationResponseMapOutput {
	return o
}

func (o ParameterSpecificationResponseMapOutput) MapIndex(k pulumi.StringInput) ParameterSpecificationResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterSpecificationResponse {
		return vs[0].(map[string]ParameterSpecificationResponse)[vs[1].(string)]
	}).(ParameterSpecificationResponseOutput)
}

type ParquetDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	CompressionCodec  interface{}                       `pulumi:"compressionCodec"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Location          interface{}                       `pulumi:"location"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type ParquetDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	CompressionCodec  interface{}                               `pulumi:"compressionCodec"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Location          interface{}                               `pulumi:"location"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type ParquetFormat struct {
	Deserializer interface{} `pulumi:"deserializer"`
	Serializer   interface{} `pulumi:"serializer"`
	Type         string      `pulumi:"type"`
}

type ParquetFormatResponse struct {
	Deserializer interface{} `pulumi:"deserializer"`
	Serializer   interface{} `pulumi:"serializer"`
	Type         string      `pulumi:"type"`
}

type ParquetSink struct {
	DisableMetricsCollection interface{}           `pulumi:"disableMetricsCollection"`
	FormatSettings           *ParquetWriteSettings `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}           `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{}           `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}           `pulumi:"sinkRetryWait"`
	StoreSettings            interface{}           `pulumi:"storeSettings"`
	Type                     string                `pulumi:"type"`
	WriteBatchSize           interface{}           `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}           `pulumi:"writeBatchTimeout"`
}

type ParquetSinkResponse struct {
	DisableMetricsCollection interface{}                   `pulumi:"disableMetricsCollection"`
	FormatSettings           *ParquetWriteSettingsResponse `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}                   `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{}                   `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}                   `pulumi:"sinkRetryWait"`
	StoreSettings            interface{}                   `pulumi:"storeSettings"`
	Type                     string                        `pulumi:"type"`
	WriteBatchSize           interface{}                   `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}                   `pulumi:"writeBatchTimeout"`
}

type ParquetSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	StoreSettings            interface{} `pulumi:"storeSettings"`
	Type                     string      `pulumi:"type"`
}

type ParquetSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	StoreSettings            interface{} `pulumi:"storeSettings"`
	Type                     string      `pulumi:"type"`
}

type ParquetWriteSettings struct {
	FileNamePrefix interface{} `pulumi:"fileNamePrefix"`
	MaxRowsPerFile interface{} `pulumi:"maxRowsPerFile"`
	Type           string      `pulumi:"type"`
}

type ParquetWriteSettingsResponse struct {
	FileNamePrefix interface{} `pulumi:"fileNamePrefix"`
	MaxRowsPerFile interface{} `pulumi:"maxRowsPerFile"`
	Type           string      `pulumi:"type"`
}

type PaypalLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	ClientId              interface{}                       `pulumi:"clientId"`
	ClientSecret          interface{}                       `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Host                  interface{}                       `pulumi:"host"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
}

type PaypalLinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	ClientId              interface{}                               `pulumi:"clientId"`
	ClientSecret          interface{}                               `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Host                  interface{}                               `pulumi:"host"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
}

type PaypalObjectDataset struct {
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

type PaypalObjectDatasetResponse struct {
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

type PaypalSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type PaypalSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type PhoenixLinkedService struct {
	AllowHostNameCNMismatch   interface{}                       `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                       `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                     `pulumi:"annotations"`
	AuthenticationType        string                            `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description               *string                           `pulumi:"description"`
	EnableSsl                 interface{}                       `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                       `pulumi:"encryptedCredential"`
	Host                      interface{}                       `pulumi:"host"`
	HttpPath                  interface{}                       `pulumi:"httpPath"`
	Parameters                map[string]ParameterSpecification `pulumi:"parameters"`
	Password                  interface{}                       `pulumi:"password"`
	Port                      interface{}                       `pulumi:"port"`
	TrustedCertPath           interface{}                       `pulumi:"trustedCertPath"`
	Type                      string                            `pulumi:"type"`
	UseSystemTrustStore       interface{}                       `pulumi:"useSystemTrustStore"`
	Username                  interface{}                       `pulumi:"username"`
}

type PhoenixLinkedServiceResponse struct {
	AllowHostNameCNMismatch   interface{}                               `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                               `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                             `pulumi:"annotations"`
	AuthenticationType        string                                    `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description               *string                                   `pulumi:"description"`
	EnableSsl                 interface{}                               `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                               `pulumi:"encryptedCredential"`
	Host                      interface{}                               `pulumi:"host"`
	HttpPath                  interface{}                               `pulumi:"httpPath"`
	Parameters                map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                  interface{}                               `pulumi:"password"`
	Port                      interface{}                               `pulumi:"port"`
	TrustedCertPath           interface{}                               `pulumi:"trustedCertPath"`
	Type                      string                                    `pulumi:"type"`
	UseSystemTrustStore       interface{}                               `pulumi:"useSystemTrustStore"`
	Username                  interface{}                               `pulumi:"username"`
}

type PhoenixObjectDataset struct {
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

type PhoenixObjectDatasetResponse struct {
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

type PhoenixSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type PhoenixSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type PipelineElapsedTimeMetricPolicy struct {
	Duration interface{} `pulumi:"duration"`
}





type PipelineElapsedTimeMetricPolicyInput interface {
	pulumi.Input

	ToPipelineElapsedTimeMetricPolicyOutput() PipelineElapsedTimeMetricPolicyOutput
	ToPipelineElapsedTimeMetricPolicyOutputWithContext(context.Context) PipelineElapsedTimeMetricPolicyOutput
}

type PipelineElapsedTimeMetricPolicyArgs struct {
	Duration pulumi.Input `pulumi:"duration"`
}

func (PipelineElapsedTimeMetricPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineElapsedTimeMetricPolicy)(nil)).Elem()
}

func (i PipelineElapsedTimeMetricPolicyArgs) ToPipelineElapsedTimeMetricPolicyOutput() PipelineElapsedTimeMetricPolicyOutput {
	return i.ToPipelineElapsedTimeMetricPolicyOutputWithContext(context.Background())
}

func (i PipelineElapsedTimeMetricPolicyArgs) ToPipelineElapsedTimeMetricPolicyOutputWithContext(ctx context.Context) PipelineElapsedTimeMetricPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineElapsedTimeMetricPolicyOutput)
}

func (i PipelineElapsedTimeMetricPolicyArgs) ToPipelineElapsedTimeMetricPolicyPtrOutput() PipelineElapsedTimeMetricPolicyPtrOutput {
	return i.ToPipelineElapsedTimeMetricPolicyPtrOutputWithContext(context.Background())
}

func (i PipelineElapsedTimeMetricPolicyArgs) ToPipelineElapsedTimeMetricPolicyPtrOutputWithContext(ctx context.Context) PipelineElapsedTimeMetricPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineElapsedTimeMetricPolicyOutput).ToPipelineElapsedTimeMetricPolicyPtrOutputWithContext(ctx)
}









type PipelineElapsedTimeMetricPolicyPtrInput interface {
	pulumi.Input

	ToPipelineElapsedTimeMetricPolicyPtrOutput() PipelineElapsedTimeMetricPolicyPtrOutput
	ToPipelineElapsedTimeMetricPolicyPtrOutputWithContext(context.Context) PipelineElapsedTimeMetricPolicyPtrOutput
}

type pipelineElapsedTimeMetricPolicyPtrType PipelineElapsedTimeMetricPolicyArgs

func PipelineElapsedTimeMetricPolicyPtr(v *PipelineElapsedTimeMetricPolicyArgs) PipelineElapsedTimeMetricPolicyPtrInput {
	return (*pipelineElapsedTimeMetricPolicyPtrType)(v)
}

func (*pipelineElapsedTimeMetricPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineElapsedTimeMetricPolicy)(nil)).Elem()
}

func (i *pipelineElapsedTimeMetricPolicyPtrType) ToPipelineElapsedTimeMetricPolicyPtrOutput() PipelineElapsedTimeMetricPolicyPtrOutput {
	return i.ToPipelineElapsedTimeMetricPolicyPtrOutputWithContext(context.Background())
}

func (i *pipelineElapsedTimeMetricPolicyPtrType) ToPipelineElapsedTimeMetricPolicyPtrOutputWithContext(ctx context.Context) PipelineElapsedTimeMetricPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineElapsedTimeMetricPolicyPtrOutput)
}

type PipelineElapsedTimeMetricPolicyOutput struct{ *pulumi.OutputState }

func (PipelineElapsedTimeMetricPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineElapsedTimeMetricPolicy)(nil)).Elem()
}

func (o PipelineElapsedTimeMetricPolicyOutput) ToPipelineElapsedTimeMetricPolicyOutput() PipelineElapsedTimeMetricPolicyOutput {
	return o
}

func (o PipelineElapsedTimeMetricPolicyOutput) ToPipelineElapsedTimeMetricPolicyOutputWithContext(ctx context.Context) PipelineElapsedTimeMetricPolicyOutput {
	return o
}

func (o PipelineElapsedTimeMetricPolicyOutput) ToPipelineElapsedTimeMetricPolicyPtrOutput() PipelineElapsedTimeMetricPolicyPtrOutput {
	return o.ToPipelineElapsedTimeMetricPolicyPtrOutputWithContext(context.Background())
}

func (o PipelineElapsedTimeMetricPolicyOutput) ToPipelineElapsedTimeMetricPolicyPtrOutputWithContext(ctx context.Context) PipelineElapsedTimeMetricPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineElapsedTimeMetricPolicy) *PipelineElapsedTimeMetricPolicy {
		return &v
	}).(PipelineElapsedTimeMetricPolicyPtrOutput)
}

func (o PipelineElapsedTimeMetricPolicyOutput) Duration() pulumi.AnyOutput {
	return o.ApplyT(func(v PipelineElapsedTimeMetricPolicy) interface{} { return v.Duration }).(pulumi.AnyOutput)
}

type PipelineElapsedTimeMetricPolicyPtrOutput struct{ *pulumi.OutputState }

func (PipelineElapsedTimeMetricPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineElapsedTimeMetricPolicy)(nil)).Elem()
}

func (o PipelineElapsedTimeMetricPolicyPtrOutput) ToPipelineElapsedTimeMetricPolicyPtrOutput() PipelineElapsedTimeMetricPolicyPtrOutput {
	return o
}

func (o PipelineElapsedTimeMetricPolicyPtrOutput) ToPipelineElapsedTimeMetricPolicyPtrOutputWithContext(ctx context.Context) PipelineElapsedTimeMetricPolicyPtrOutput {
	return o
}

func (o PipelineElapsedTimeMetricPolicyPtrOutput) Elem() PipelineElapsedTimeMetricPolicyOutput {
	return o.ApplyT(func(v *PipelineElapsedTimeMetricPolicy) PipelineElapsedTimeMetricPolicy {
		if v != nil {
			return *v
		}
		var ret PipelineElapsedTimeMetricPolicy
		return ret
	}).(PipelineElapsedTimeMetricPolicyOutput)
}

func (o PipelineElapsedTimeMetricPolicyPtrOutput) Duration() pulumi.AnyOutput {
	return o.ApplyT(func(v *PipelineElapsedTimeMetricPolicy) interface{} {
		if v == nil {
			return nil
		}
		return v.Duration
	}).(pulumi.AnyOutput)
}

type PipelineElapsedTimeMetricPolicyResponse struct {
	Duration interface{} `pulumi:"duration"`
}

type PipelineElapsedTimeMetricPolicyResponseOutput struct{ *pulumi.OutputState }

func (PipelineElapsedTimeMetricPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineElapsedTimeMetricPolicyResponse)(nil)).Elem()
}

func (o PipelineElapsedTimeMetricPolicyResponseOutput) ToPipelineElapsedTimeMetricPolicyResponseOutput() PipelineElapsedTimeMetricPolicyResponseOutput {
	return o
}

func (o PipelineElapsedTimeMetricPolicyResponseOutput) ToPipelineElapsedTimeMetricPolicyResponseOutputWithContext(ctx context.Context) PipelineElapsedTimeMetricPolicyResponseOutput {
	return o
}

func (o PipelineElapsedTimeMetricPolicyResponseOutput) Duration() pulumi.AnyOutput {
	return o.ApplyT(func(v PipelineElapsedTimeMetricPolicyResponse) interface{} { return v.Duration }).(pulumi.AnyOutput)
}

type PipelineElapsedTimeMetricPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelineElapsedTimeMetricPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineElapsedTimeMetricPolicyResponse)(nil)).Elem()
}

func (o PipelineElapsedTimeMetricPolicyResponsePtrOutput) ToPipelineElapsedTimeMetricPolicyResponsePtrOutput() PipelineElapsedTimeMetricPolicyResponsePtrOutput {
	return o
}

func (o PipelineElapsedTimeMetricPolicyResponsePtrOutput) ToPipelineElapsedTimeMetricPolicyResponsePtrOutputWithContext(ctx context.Context) PipelineElapsedTimeMetricPolicyResponsePtrOutput {
	return o
}

func (o PipelineElapsedTimeMetricPolicyResponsePtrOutput) Elem() PipelineElapsedTimeMetricPolicyResponseOutput {
	return o.ApplyT(func(v *PipelineElapsedTimeMetricPolicyResponse) PipelineElapsedTimeMetricPolicyResponse {
		if v != nil {
			return *v
		}
		var ret PipelineElapsedTimeMetricPolicyResponse
		return ret
	}).(PipelineElapsedTimeMetricPolicyResponseOutput)
}

func (o PipelineElapsedTimeMetricPolicyResponsePtrOutput) Duration() pulumi.AnyOutput {
	return o.ApplyT(func(v *PipelineElapsedTimeMetricPolicyResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Duration
	}).(pulumi.AnyOutput)
}

type PipelineFolder struct {
	Name *string `pulumi:"name"`
}





type PipelineFolderInput interface {
	pulumi.Input

	ToPipelineFolderOutput() PipelineFolderOutput
	ToPipelineFolderOutputWithContext(context.Context) PipelineFolderOutput
}

type PipelineFolderArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (PipelineFolderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineFolder)(nil)).Elem()
}

func (i PipelineFolderArgs) ToPipelineFolderOutput() PipelineFolderOutput {
	return i.ToPipelineFolderOutputWithContext(context.Background())
}

func (i PipelineFolderArgs) ToPipelineFolderOutputWithContext(ctx context.Context) PipelineFolderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineFolderOutput)
}

func (i PipelineFolderArgs) ToPipelineFolderPtrOutput() PipelineFolderPtrOutput {
	return i.ToPipelineFolderPtrOutputWithContext(context.Background())
}

func (i PipelineFolderArgs) ToPipelineFolderPtrOutputWithContext(ctx context.Context) PipelineFolderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineFolderOutput).ToPipelineFolderPtrOutputWithContext(ctx)
}









type PipelineFolderPtrInput interface {
	pulumi.Input

	ToPipelineFolderPtrOutput() PipelineFolderPtrOutput
	ToPipelineFolderPtrOutputWithContext(context.Context) PipelineFolderPtrOutput
}

type pipelineFolderPtrType PipelineFolderArgs

func PipelineFolderPtr(v *PipelineFolderArgs) PipelineFolderPtrInput {
	return (*pipelineFolderPtrType)(v)
}

func (*pipelineFolderPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineFolder)(nil)).Elem()
}

func (i *pipelineFolderPtrType) ToPipelineFolderPtrOutput() PipelineFolderPtrOutput {
	return i.ToPipelineFolderPtrOutputWithContext(context.Background())
}

func (i *pipelineFolderPtrType) ToPipelineFolderPtrOutputWithContext(ctx context.Context) PipelineFolderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineFolderPtrOutput)
}

type PipelineFolderOutput struct{ *pulumi.OutputState }

func (PipelineFolderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineFolder)(nil)).Elem()
}

func (o PipelineFolderOutput) ToPipelineFolderOutput() PipelineFolderOutput {
	return o
}

func (o PipelineFolderOutput) ToPipelineFolderOutputWithContext(ctx context.Context) PipelineFolderOutput {
	return o
}

func (o PipelineFolderOutput) ToPipelineFolderPtrOutput() PipelineFolderPtrOutput {
	return o.ToPipelineFolderPtrOutputWithContext(context.Background())
}

func (o PipelineFolderOutput) ToPipelineFolderPtrOutputWithContext(ctx context.Context) PipelineFolderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineFolder) *PipelineFolder {
		return &v
	}).(PipelineFolderPtrOutput)
}

func (o PipelineFolderOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineFolder) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type PipelineFolderPtrOutput struct{ *pulumi.OutputState }

func (PipelineFolderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineFolder)(nil)).Elem()
}

func (o PipelineFolderPtrOutput) ToPipelineFolderPtrOutput() PipelineFolderPtrOutput {
	return o
}

func (o PipelineFolderPtrOutput) ToPipelineFolderPtrOutputWithContext(ctx context.Context) PipelineFolderPtrOutput {
	return o
}

func (o PipelineFolderPtrOutput) Elem() PipelineFolderOutput {
	return o.ApplyT(func(v *PipelineFolder) PipelineFolder {
		if v != nil {
			return *v
		}
		var ret PipelineFolder
		return ret
	}).(PipelineFolderOutput)
}

func (o PipelineFolderPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineFolder) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type PipelinePolicy struct {
	ElapsedTimeMetric *PipelineElapsedTimeMetricPolicy `pulumi:"elapsedTimeMetric"`
}





type PipelinePolicyInput interface {
	pulumi.Input

	ToPipelinePolicyOutput() PipelinePolicyOutput
	ToPipelinePolicyOutputWithContext(context.Context) PipelinePolicyOutput
}

type PipelinePolicyArgs struct {
	ElapsedTimeMetric PipelineElapsedTimeMetricPolicyPtrInput `pulumi:"elapsedTimeMetric"`
}

func (PipelinePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelinePolicy)(nil)).Elem()
}

func (i PipelinePolicyArgs) ToPipelinePolicyOutput() PipelinePolicyOutput {
	return i.ToPipelinePolicyOutputWithContext(context.Background())
}

func (i PipelinePolicyArgs) ToPipelinePolicyOutputWithContext(ctx context.Context) PipelinePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelinePolicyOutput)
}

func (i PipelinePolicyArgs) ToPipelinePolicyPtrOutput() PipelinePolicyPtrOutput {
	return i.ToPipelinePolicyPtrOutputWithContext(context.Background())
}

func (i PipelinePolicyArgs) ToPipelinePolicyPtrOutputWithContext(ctx context.Context) PipelinePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelinePolicyOutput).ToPipelinePolicyPtrOutputWithContext(ctx)
}









type PipelinePolicyPtrInput interface {
	pulumi.Input

	ToPipelinePolicyPtrOutput() PipelinePolicyPtrOutput
	ToPipelinePolicyPtrOutputWithContext(context.Context) PipelinePolicyPtrOutput
}

type pipelinePolicyPtrType PipelinePolicyArgs

func PipelinePolicyPtr(v *PipelinePolicyArgs) PipelinePolicyPtrInput {
	return (*pipelinePolicyPtrType)(v)
}

func (*pipelinePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelinePolicy)(nil)).Elem()
}

func (i *pipelinePolicyPtrType) ToPipelinePolicyPtrOutput() PipelinePolicyPtrOutput {
	return i.ToPipelinePolicyPtrOutputWithContext(context.Background())
}

func (i *pipelinePolicyPtrType) ToPipelinePolicyPtrOutputWithContext(ctx context.Context) PipelinePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelinePolicyPtrOutput)
}

type PipelinePolicyOutput struct{ *pulumi.OutputState }

func (PipelinePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelinePolicy)(nil)).Elem()
}

func (o PipelinePolicyOutput) ToPipelinePolicyOutput() PipelinePolicyOutput {
	return o
}

func (o PipelinePolicyOutput) ToPipelinePolicyOutputWithContext(ctx context.Context) PipelinePolicyOutput {
	return o
}

func (o PipelinePolicyOutput) ToPipelinePolicyPtrOutput() PipelinePolicyPtrOutput {
	return o.ToPipelinePolicyPtrOutputWithContext(context.Background())
}

func (o PipelinePolicyOutput) ToPipelinePolicyPtrOutputWithContext(ctx context.Context) PipelinePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelinePolicy) *PipelinePolicy {
		return &v
	}).(PipelinePolicyPtrOutput)
}

func (o PipelinePolicyOutput) ElapsedTimeMetric() PipelineElapsedTimeMetricPolicyPtrOutput {
	return o.ApplyT(func(v PipelinePolicy) *PipelineElapsedTimeMetricPolicy { return v.ElapsedTimeMetric }).(PipelineElapsedTimeMetricPolicyPtrOutput)
}

type PipelinePolicyPtrOutput struct{ *pulumi.OutputState }

func (PipelinePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelinePolicy)(nil)).Elem()
}

func (o PipelinePolicyPtrOutput) ToPipelinePolicyPtrOutput() PipelinePolicyPtrOutput {
	return o
}

func (o PipelinePolicyPtrOutput) ToPipelinePolicyPtrOutputWithContext(ctx context.Context) PipelinePolicyPtrOutput {
	return o
}

func (o PipelinePolicyPtrOutput) Elem() PipelinePolicyOutput {
	return o.ApplyT(func(v *PipelinePolicy) PipelinePolicy {
		if v != nil {
			return *v
		}
		var ret PipelinePolicy
		return ret
	}).(PipelinePolicyOutput)
}

func (o PipelinePolicyPtrOutput) ElapsedTimeMetric() PipelineElapsedTimeMetricPolicyPtrOutput {
	return o.ApplyT(func(v *PipelinePolicy) *PipelineElapsedTimeMetricPolicy {
		if v == nil {
			return nil
		}
		return v.ElapsedTimeMetric
	}).(PipelineElapsedTimeMetricPolicyPtrOutput)
}

type PipelinePolicyResponse struct {
	ElapsedTimeMetric *PipelineElapsedTimeMetricPolicyResponse `pulumi:"elapsedTimeMetric"`
}

type PipelinePolicyResponseOutput struct{ *pulumi.OutputState }

func (PipelinePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelinePolicyResponse)(nil)).Elem()
}

func (o PipelinePolicyResponseOutput) ToPipelinePolicyResponseOutput() PipelinePolicyResponseOutput {
	return o
}

func (o PipelinePolicyResponseOutput) ToPipelinePolicyResponseOutputWithContext(ctx context.Context) PipelinePolicyResponseOutput {
	return o
}

func (o PipelinePolicyResponseOutput) ElapsedTimeMetric() PipelineElapsedTimeMetricPolicyResponsePtrOutput {
	return o.ApplyT(func(v PipelinePolicyResponse) *PipelineElapsedTimeMetricPolicyResponse { return v.ElapsedTimeMetric }).(PipelineElapsedTimeMetricPolicyResponsePtrOutput)
}

type PipelinePolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelinePolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelinePolicyResponse)(nil)).Elem()
}

func (o PipelinePolicyResponsePtrOutput) ToPipelinePolicyResponsePtrOutput() PipelinePolicyResponsePtrOutput {
	return o
}

func (o PipelinePolicyResponsePtrOutput) ToPipelinePolicyResponsePtrOutputWithContext(ctx context.Context) PipelinePolicyResponsePtrOutput {
	return o
}

func (o PipelinePolicyResponsePtrOutput) Elem() PipelinePolicyResponseOutput {
	return o.ApplyT(func(v *PipelinePolicyResponse) PipelinePolicyResponse {
		if v != nil {
			return *v
		}
		var ret PipelinePolicyResponse
		return ret
	}).(PipelinePolicyResponseOutput)
}

func (o PipelinePolicyResponsePtrOutput) ElapsedTimeMetric() PipelineElapsedTimeMetricPolicyResponsePtrOutput {
	return o.ApplyT(func(v *PipelinePolicyResponse) *PipelineElapsedTimeMetricPolicyResponse {
		if v == nil {
			return nil
		}
		return v.ElapsedTimeMetric
	}).(PipelineElapsedTimeMetricPolicyResponsePtrOutput)
}

type PipelineReference struct {
	Name          *string `pulumi:"name"`
	ReferenceName string  `pulumi:"referenceName"`
	Type          string  `pulumi:"type"`
}

type PipelineReferenceResponse struct {
	Name          *string `pulumi:"name"`
	ReferenceName string  `pulumi:"referenceName"`
	Type          string  `pulumi:"type"`
}

type PipelineResponseFolder struct {
	Name *string `pulumi:"name"`
}

type PipelineResponseFolderOutput struct{ *pulumi.OutputState }

func (PipelineResponseFolderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineResponseFolder)(nil)).Elem()
}

func (o PipelineResponseFolderOutput) ToPipelineResponseFolderOutput() PipelineResponseFolderOutput {
	return o
}

func (o PipelineResponseFolderOutput) ToPipelineResponseFolderOutputWithContext(ctx context.Context) PipelineResponseFolderOutput {
	return o
}

func (o PipelineResponseFolderOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineResponseFolder) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type PipelineResponseFolderPtrOutput struct{ *pulumi.OutputState }

func (PipelineResponseFolderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineResponseFolder)(nil)).Elem()
}

func (o PipelineResponseFolderPtrOutput) ToPipelineResponseFolderPtrOutput() PipelineResponseFolderPtrOutput {
	return o
}

func (o PipelineResponseFolderPtrOutput) ToPipelineResponseFolderPtrOutputWithContext(ctx context.Context) PipelineResponseFolderPtrOutput {
	return o
}

func (o PipelineResponseFolderPtrOutput) Elem() PipelineResponseFolderOutput {
	return o.ApplyT(func(v *PipelineResponseFolder) PipelineResponseFolder {
		if v != nil {
			return *v
		}
		var ret PipelineResponseFolder
		return ret
	}).(PipelineResponseFolderOutput)
}

func (o PipelineResponseFolderPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineResponseFolder) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type PolybaseSettings struct {
	RejectSampleValue interface{} `pulumi:"rejectSampleValue"`
	RejectType        *string     `pulumi:"rejectType"`
	RejectValue       interface{} `pulumi:"rejectValue"`
	UseTypeDefault    interface{} `pulumi:"useTypeDefault"`
}

type PolybaseSettingsResponse struct {
	RejectSampleValue interface{} `pulumi:"rejectSampleValue"`
	RejectType        *string     `pulumi:"rejectType"`
	RejectValue       interface{} `pulumi:"rejectValue"`
	UseTypeDefault    interface{} `pulumi:"useTypeDefault"`
}

type PostgreSqlLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReference     `pulumi:"password"`
	Type                string                            `pulumi:"type"`
}

type PostgreSqlLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReferenceResponse     `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
}

type PostgreSqlSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type PostgreSqlSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type PostgreSqlTableDataset struct {
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

type PostgreSqlTableDatasetResponse struct {
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

type PowerQuerySink struct {
	Dataset                   *DatasetReference       `pulumi:"dataset"`
	Description               *string                 `pulumi:"description"`
	Flowlet                   *DataFlowReference      `pulumi:"flowlet"`
	LinkedService             *LinkedServiceReference `pulumi:"linkedService"`
	Name                      string                  `pulumi:"name"`
	RejectedDataLinkedService *LinkedServiceReference `pulumi:"rejectedDataLinkedService"`
	SchemaLinkedService       *LinkedServiceReference `pulumi:"schemaLinkedService"`
	Script                    *string                 `pulumi:"script"`
}

type PowerQuerySinkMapping struct {
	DataflowSinks []PowerQuerySink `pulumi:"dataflowSinks"`
	QueryName     *string          `pulumi:"queryName"`
}

type PowerQuerySinkMappingResponse struct {
	DataflowSinks []PowerQuerySinkResponse `pulumi:"dataflowSinks"`
	QueryName     *string                  `pulumi:"queryName"`
}

type PowerQuerySinkResponse struct {
	Dataset                   *DatasetReferenceResponse       `pulumi:"dataset"`
	Description               *string                         `pulumi:"description"`
	Flowlet                   *DataFlowReferenceResponse      `pulumi:"flowlet"`
	LinkedService             *LinkedServiceReferenceResponse `pulumi:"linkedService"`
	Name                      string                          `pulumi:"name"`
	RejectedDataLinkedService *LinkedServiceReferenceResponse `pulumi:"rejectedDataLinkedService"`
	SchemaLinkedService       *LinkedServiceReferenceResponse `pulumi:"schemaLinkedService"`
	Script                    *string                         `pulumi:"script"`
}

type PowerQuerySource struct {
	Dataset             *DatasetReference       `pulumi:"dataset"`
	Description         *string                 `pulumi:"description"`
	Flowlet             *DataFlowReference      `pulumi:"flowlet"`
	LinkedService       *LinkedServiceReference `pulumi:"linkedService"`
	Name                string                  `pulumi:"name"`
	SchemaLinkedService *LinkedServiceReference `pulumi:"schemaLinkedService"`
	Script              *string                 `pulumi:"script"`
}

type PowerQuerySourceResponse struct {
	Dataset             *DatasetReferenceResponse       `pulumi:"dataset"`
	Description         *string                         `pulumi:"description"`
	Flowlet             *DataFlowReferenceResponse      `pulumi:"flowlet"`
	LinkedService       *LinkedServiceReferenceResponse `pulumi:"linkedService"`
	Name                string                          `pulumi:"name"`
	SchemaLinkedService *LinkedServiceReferenceResponse `pulumi:"schemaLinkedService"`
	Script              *string                         `pulumi:"script"`
}

type PrestoLinkedService struct {
	AllowHostNameCNMismatch   interface{}                       `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                       `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                     `pulumi:"annotations"`
	AuthenticationType        string                            `pulumi:"authenticationType"`
	Catalog                   interface{}                       `pulumi:"catalog"`
	ConnectVia                *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description               *string                           `pulumi:"description"`
	EnableSsl                 interface{}                       `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                       `pulumi:"encryptedCredential"`
	Host                      interface{}                       `pulumi:"host"`
	Parameters                map[string]ParameterSpecification `pulumi:"parameters"`
	Password                  interface{}                       `pulumi:"password"`
	Port                      interface{}                       `pulumi:"port"`
	ServerVersion             interface{}                       `pulumi:"serverVersion"`
	TimeZoneID                interface{}                       `pulumi:"timeZoneID"`
	TrustedCertPath           interface{}                       `pulumi:"trustedCertPath"`
	Type                      string                            `pulumi:"type"`
	UseSystemTrustStore       interface{}                       `pulumi:"useSystemTrustStore"`
	Username                  interface{}                       `pulumi:"username"`
}

type PrestoLinkedServiceResponse struct {
	AllowHostNameCNMismatch   interface{}                               `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                               `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                             `pulumi:"annotations"`
	AuthenticationType        string                                    `pulumi:"authenticationType"`
	Catalog                   interface{}                               `pulumi:"catalog"`
	ConnectVia                *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description               *string                                   `pulumi:"description"`
	EnableSsl                 interface{}                               `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                               `pulumi:"encryptedCredential"`
	Host                      interface{}                               `pulumi:"host"`
	Parameters                map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                  interface{}                               `pulumi:"password"`
	Port                      interface{}                               `pulumi:"port"`
	ServerVersion             interface{}                               `pulumi:"serverVersion"`
	TimeZoneID                interface{}                               `pulumi:"timeZoneID"`
	TrustedCertPath           interface{}                               `pulumi:"trustedCertPath"`
	Type                      string                                    `pulumi:"type"`
	UseSystemTrustStore       interface{}                               `pulumi:"useSystemTrustStore"`
	Username                  interface{}                               `pulumi:"username"`
}

type PrestoObjectDataset struct {
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

type PrestoObjectDatasetResponse struct {
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

type PrestoSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type PrestoSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type PrivateEndpoint struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointInput interface {
	pulumi.Input

	ToPrivateEndpointOutput() PrivateEndpointOutput
	ToPrivateEndpointOutputWithContext(context.Context) PrivateEndpointOutput
}

type PrivateEndpointArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return i.ToPrivateEndpointOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput)
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput).ToPrivateEndpointPtrOutputWithContext(ctx)
}









type PrivateEndpointPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput
	ToPrivateEndpointPtrOutputWithContext(context.Context) PrivateEndpointPtrOutput
}

type privateEndpointPtrType PrivateEndpointArgs

func PrivateEndpointPtr(v *PrivateEndpointArgs) PrivateEndpointPtrInput {
	return (*privateEndpointPtrType)(v)
}

func (*privateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPtrOutput)
}

type PrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpoint) *PrivateEndpoint {
		return &v
	}).(PrivateEndpointPtrOutput)
}

func (o PrivateEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) Elem() PrivateEndpointOutput {
	return o.ApplyT(func(v *PrivateEndpoint) PrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret PrivateEndpoint
		return ret
	}).(PrivateEndpointOutput)
}

func (o PrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkConnectionApprovalRequest struct {
	PrivateEndpoint                   *PrivateEndpoint            `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateLinkConnectionApprovalRequestInput interface {
	pulumi.Input

	ToPrivateLinkConnectionApprovalRequestOutput() PrivateLinkConnectionApprovalRequestOutput
	ToPrivateLinkConnectionApprovalRequestOutputWithContext(context.Context) PrivateLinkConnectionApprovalRequestOutput
}

type PrivateLinkConnectionApprovalRequestArgs struct {
	PrivateEndpoint                   PrivateEndpointPtrInput            `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkConnectionStatePtrInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateLinkConnectionApprovalRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConnectionApprovalRequest)(nil)).Elem()
}

func (i PrivateLinkConnectionApprovalRequestArgs) ToPrivateLinkConnectionApprovalRequestOutput() PrivateLinkConnectionApprovalRequestOutput {
	return i.ToPrivateLinkConnectionApprovalRequestOutputWithContext(context.Background())
}

func (i PrivateLinkConnectionApprovalRequestArgs) ToPrivateLinkConnectionApprovalRequestOutputWithContext(ctx context.Context) PrivateLinkConnectionApprovalRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionApprovalRequestOutput)
}

func (i PrivateLinkConnectionApprovalRequestArgs) ToPrivateLinkConnectionApprovalRequestPtrOutput() PrivateLinkConnectionApprovalRequestPtrOutput {
	return i.ToPrivateLinkConnectionApprovalRequestPtrOutputWithContext(context.Background())
}

func (i PrivateLinkConnectionApprovalRequestArgs) ToPrivateLinkConnectionApprovalRequestPtrOutputWithContext(ctx context.Context) PrivateLinkConnectionApprovalRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionApprovalRequestOutput).ToPrivateLinkConnectionApprovalRequestPtrOutputWithContext(ctx)
}









type PrivateLinkConnectionApprovalRequestPtrInput interface {
	pulumi.Input

	ToPrivateLinkConnectionApprovalRequestPtrOutput() PrivateLinkConnectionApprovalRequestPtrOutput
	ToPrivateLinkConnectionApprovalRequestPtrOutputWithContext(context.Context) PrivateLinkConnectionApprovalRequestPtrOutput
}

type privateLinkConnectionApprovalRequestPtrType PrivateLinkConnectionApprovalRequestArgs

func PrivateLinkConnectionApprovalRequestPtr(v *PrivateLinkConnectionApprovalRequestArgs) PrivateLinkConnectionApprovalRequestPtrInput {
	return (*privateLinkConnectionApprovalRequestPtrType)(v)
}

func (*privateLinkConnectionApprovalRequestPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkConnectionApprovalRequest)(nil)).Elem()
}

func (i *privateLinkConnectionApprovalRequestPtrType) ToPrivateLinkConnectionApprovalRequestPtrOutput() PrivateLinkConnectionApprovalRequestPtrOutput {
	return i.ToPrivateLinkConnectionApprovalRequestPtrOutputWithContext(context.Background())
}

func (i *privateLinkConnectionApprovalRequestPtrType) ToPrivateLinkConnectionApprovalRequestPtrOutputWithContext(ctx context.Context) PrivateLinkConnectionApprovalRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionApprovalRequestPtrOutput)
}

type PrivateLinkConnectionApprovalRequestOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionApprovalRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConnectionApprovalRequest)(nil)).Elem()
}

func (o PrivateLinkConnectionApprovalRequestOutput) ToPrivateLinkConnectionApprovalRequestOutput() PrivateLinkConnectionApprovalRequestOutput {
	return o
}

func (o PrivateLinkConnectionApprovalRequestOutput) ToPrivateLinkConnectionApprovalRequestOutputWithContext(ctx context.Context) PrivateLinkConnectionApprovalRequestOutput {
	return o
}

func (o PrivateLinkConnectionApprovalRequestOutput) ToPrivateLinkConnectionApprovalRequestPtrOutput() PrivateLinkConnectionApprovalRequestPtrOutput {
	return o.ToPrivateLinkConnectionApprovalRequestPtrOutputWithContext(context.Background())
}

func (o PrivateLinkConnectionApprovalRequestOutput) ToPrivateLinkConnectionApprovalRequestPtrOutputWithContext(ctx context.Context) PrivateLinkConnectionApprovalRequestPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkConnectionApprovalRequest) *PrivateLinkConnectionApprovalRequest {
		return &v
	}).(PrivateLinkConnectionApprovalRequestPtrOutput)
}

func (o PrivateLinkConnectionApprovalRequestOutput) PrivateEndpoint() PrivateEndpointPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionApprovalRequest) *PrivateEndpoint { return v.PrivateEndpoint }).(PrivateEndpointPtrOutput)
}

func (o PrivateLinkConnectionApprovalRequestOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStatePtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionApprovalRequest) *PrivateLinkConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStatePtrOutput)
}

type PrivateLinkConnectionApprovalRequestPtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionApprovalRequestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkConnectionApprovalRequest)(nil)).Elem()
}

func (o PrivateLinkConnectionApprovalRequestPtrOutput) ToPrivateLinkConnectionApprovalRequestPtrOutput() PrivateLinkConnectionApprovalRequestPtrOutput {
	return o
}

func (o PrivateLinkConnectionApprovalRequestPtrOutput) ToPrivateLinkConnectionApprovalRequestPtrOutputWithContext(ctx context.Context) PrivateLinkConnectionApprovalRequestPtrOutput {
	return o
}

func (o PrivateLinkConnectionApprovalRequestPtrOutput) Elem() PrivateLinkConnectionApprovalRequestOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionApprovalRequest) PrivateLinkConnectionApprovalRequest {
		if v != nil {
			return *v
		}
		var ret PrivateLinkConnectionApprovalRequest
		return ret
	}).(PrivateLinkConnectionApprovalRequestOutput)
}

func (o PrivateLinkConnectionApprovalRequestPtrOutput) PrivateEndpoint() PrivateEndpointPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionApprovalRequest) *PrivateEndpoint {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoint
	}).(PrivateEndpointPtrOutput)
}

func (o PrivateLinkConnectionApprovalRequestPtrOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStatePtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionApprovalRequest) *PrivateLinkConnectionState {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStatePtrOutput)
}

type PrivateLinkConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkConnectionStateOutput() PrivateLinkConnectionStateOutput
	ToPrivateLinkConnectionStateOutputWithContext(context.Context) PrivateLinkConnectionStateOutput
}

type PrivateLinkConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConnectionState)(nil)).Elem()
}

func (i PrivateLinkConnectionStateArgs) ToPrivateLinkConnectionStateOutput() PrivateLinkConnectionStateOutput {
	return i.ToPrivateLinkConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkConnectionStateArgs) ToPrivateLinkConnectionStateOutputWithContext(ctx context.Context) PrivateLinkConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionStateOutput)
}

func (i PrivateLinkConnectionStateArgs) ToPrivateLinkConnectionStatePtrOutput() PrivateLinkConnectionStatePtrOutput {
	return i.ToPrivateLinkConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkConnectionStateArgs) ToPrivateLinkConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionStateOutput).ToPrivateLinkConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkConnectionStatePtrOutput() PrivateLinkConnectionStatePtrOutput
	ToPrivateLinkConnectionStatePtrOutputWithContext(context.Context) PrivateLinkConnectionStatePtrOutput
}

type privateLinkConnectionStatePtrType PrivateLinkConnectionStateArgs

func PrivateLinkConnectionStatePtr(v *PrivateLinkConnectionStateArgs) PrivateLinkConnectionStatePtrInput {
	return (*privateLinkConnectionStatePtrType)(v)
}

func (*privateLinkConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkConnectionState)(nil)).Elem()
}

func (i *privateLinkConnectionStatePtrType) ToPrivateLinkConnectionStatePtrOutput() PrivateLinkConnectionStatePtrOutput {
	return i.ToPrivateLinkConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkConnectionStatePtrType) ToPrivateLinkConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionStatePtrOutput)
}

type PrivateLinkConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConnectionState)(nil)).Elem()
}

func (o PrivateLinkConnectionStateOutput) ToPrivateLinkConnectionStateOutput() PrivateLinkConnectionStateOutput {
	return o
}

func (o PrivateLinkConnectionStateOutput) ToPrivateLinkConnectionStateOutputWithContext(ctx context.Context) PrivateLinkConnectionStateOutput {
	return o
}

func (o PrivateLinkConnectionStateOutput) ToPrivateLinkConnectionStatePtrOutput() PrivateLinkConnectionStatePtrOutput {
	return o.ToPrivateLinkConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkConnectionStateOutput) ToPrivateLinkConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkConnectionState) *PrivateLinkConnectionState {
		return &v
	}).(PrivateLinkConnectionStatePtrOutput)
}

func (o PrivateLinkConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkConnectionState)(nil)).Elem()
}

func (o PrivateLinkConnectionStatePtrOutput) ToPrivateLinkConnectionStatePtrOutput() PrivateLinkConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkConnectionStatePtrOutput) ToPrivateLinkConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkConnectionStatePtrOutput) Elem() PrivateLinkConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionState) PrivateLinkConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkConnectionState
		return ret
	}).(PrivateLinkConnectionStateOutput)
}

func (o PrivateLinkConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
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

func (o PrivateLinkConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
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
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PurviewConfiguration struct {
	PurviewResourceId *string `pulumi:"purviewResourceId"`
}





type PurviewConfigurationInput interface {
	pulumi.Input

	ToPurviewConfigurationOutput() PurviewConfigurationOutput
	ToPurviewConfigurationOutputWithContext(context.Context) PurviewConfigurationOutput
}

type PurviewConfigurationArgs struct {
	PurviewResourceId pulumi.StringPtrInput `pulumi:"purviewResourceId"`
}

func (PurviewConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PurviewConfiguration)(nil)).Elem()
}

func (i PurviewConfigurationArgs) ToPurviewConfigurationOutput() PurviewConfigurationOutput {
	return i.ToPurviewConfigurationOutputWithContext(context.Background())
}

func (i PurviewConfigurationArgs) ToPurviewConfigurationOutputWithContext(ctx context.Context) PurviewConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PurviewConfigurationOutput)
}

func (i PurviewConfigurationArgs) ToPurviewConfigurationPtrOutput() PurviewConfigurationPtrOutput {
	return i.ToPurviewConfigurationPtrOutputWithContext(context.Background())
}

func (i PurviewConfigurationArgs) ToPurviewConfigurationPtrOutputWithContext(ctx context.Context) PurviewConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PurviewConfigurationOutput).ToPurviewConfigurationPtrOutputWithContext(ctx)
}









type PurviewConfigurationPtrInput interface {
	pulumi.Input

	ToPurviewConfigurationPtrOutput() PurviewConfigurationPtrOutput
	ToPurviewConfigurationPtrOutputWithContext(context.Context) PurviewConfigurationPtrOutput
}

type purviewConfigurationPtrType PurviewConfigurationArgs

func PurviewConfigurationPtr(v *PurviewConfigurationArgs) PurviewConfigurationPtrInput {
	return (*purviewConfigurationPtrType)(v)
}

func (*purviewConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PurviewConfiguration)(nil)).Elem()
}

func (i *purviewConfigurationPtrType) ToPurviewConfigurationPtrOutput() PurviewConfigurationPtrOutput {
	return i.ToPurviewConfigurationPtrOutputWithContext(context.Background())
}

func (i *purviewConfigurationPtrType) ToPurviewConfigurationPtrOutputWithContext(ctx context.Context) PurviewConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PurviewConfigurationPtrOutput)
}

type PurviewConfigurationOutput struct{ *pulumi.OutputState }

func (PurviewConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PurviewConfiguration)(nil)).Elem()
}

func (o PurviewConfigurationOutput) ToPurviewConfigurationOutput() PurviewConfigurationOutput {
	return o
}

func (o PurviewConfigurationOutput) ToPurviewConfigurationOutputWithContext(ctx context.Context) PurviewConfigurationOutput {
	return o
}

func (o PurviewConfigurationOutput) ToPurviewConfigurationPtrOutput() PurviewConfigurationPtrOutput {
	return o.ToPurviewConfigurationPtrOutputWithContext(context.Background())
}

func (o PurviewConfigurationOutput) ToPurviewConfigurationPtrOutputWithContext(ctx context.Context) PurviewConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PurviewConfiguration) *PurviewConfiguration {
		return &v
	}).(PurviewConfigurationPtrOutput)
}

func (o PurviewConfigurationOutput) PurviewResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PurviewConfiguration) *string { return v.PurviewResourceId }).(pulumi.StringPtrOutput)
}

type PurviewConfigurationPtrOutput struct{ *pulumi.OutputState }

func (PurviewConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PurviewConfiguration)(nil)).Elem()
}

func (o PurviewConfigurationPtrOutput) ToPurviewConfigurationPtrOutput() PurviewConfigurationPtrOutput {
	return o
}

func (o PurviewConfigurationPtrOutput) ToPurviewConfigurationPtrOutputWithContext(ctx context.Context) PurviewConfigurationPtrOutput {
	return o
}

func (o PurviewConfigurationPtrOutput) Elem() PurviewConfigurationOutput {
	return o.ApplyT(func(v *PurviewConfiguration) PurviewConfiguration {
		if v != nil {
			return *v
		}
		var ret PurviewConfiguration
		return ret
	}).(PurviewConfigurationOutput)
}

func (o PurviewConfigurationPtrOutput) PurviewResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurviewConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.PurviewResourceId
	}).(pulumi.StringPtrOutput)
}

type PurviewConfigurationResponse struct {
	PurviewResourceId *string `pulumi:"purviewResourceId"`
}

type PurviewConfigurationResponseOutput struct{ *pulumi.OutputState }

func (PurviewConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PurviewConfigurationResponse)(nil)).Elem()
}

func (o PurviewConfigurationResponseOutput) ToPurviewConfigurationResponseOutput() PurviewConfigurationResponseOutput {
	return o
}

func (o PurviewConfigurationResponseOutput) ToPurviewConfigurationResponseOutputWithContext(ctx context.Context) PurviewConfigurationResponseOutput {
	return o
}

func (o PurviewConfigurationResponseOutput) PurviewResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PurviewConfigurationResponse) *string { return v.PurviewResourceId }).(pulumi.StringPtrOutput)
}

type PurviewConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (PurviewConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PurviewConfigurationResponse)(nil)).Elem()
}

func (o PurviewConfigurationResponsePtrOutput) ToPurviewConfigurationResponsePtrOutput() PurviewConfigurationResponsePtrOutput {
	return o
}

func (o PurviewConfigurationResponsePtrOutput) ToPurviewConfigurationResponsePtrOutputWithContext(ctx context.Context) PurviewConfigurationResponsePtrOutput {
	return o
}

func (o PurviewConfigurationResponsePtrOutput) Elem() PurviewConfigurationResponseOutput {
	return o.ApplyT(func(v *PurviewConfigurationResponse) PurviewConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret PurviewConfigurationResponse
		return ret
	}).(PurviewConfigurationResponseOutput)
}

func (o PurviewConfigurationResponsePtrOutput) PurviewResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurviewConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PurviewResourceId
	}).(pulumi.StringPtrOutput)
}

type QuickBooksLinkedService struct {
	AccessToken           interface{}                       `pulumi:"accessToken"`
	AccessTokenSecret     interface{}                       `pulumi:"accessTokenSecret"`
	Annotations           []interface{}                     `pulumi:"annotations"`
	CompanyId             interface{}                       `pulumi:"companyId"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionProperties  interface{}                       `pulumi:"connectionProperties"`
	ConsumerKey           interface{}                       `pulumi:"consumerKey"`
	ConsumerSecret        interface{}                       `pulumi:"consumerSecret"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Endpoint              interface{}                       `pulumi:"endpoint"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
}

type QuickBooksLinkedServiceResponse struct {
	AccessToken           interface{}                               `pulumi:"accessToken"`
	AccessTokenSecret     interface{}                               `pulumi:"accessTokenSecret"`
	Annotations           []interface{}                             `pulumi:"annotations"`
	CompanyId             interface{}                               `pulumi:"companyId"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionProperties  interface{}                               `pulumi:"connectionProperties"`
	ConsumerKey           interface{}                               `pulumi:"consumerKey"`
	ConsumerSecret        interface{}                               `pulumi:"consumerSecret"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Endpoint              interface{}                               `pulumi:"endpoint"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
}

type QuickBooksObjectDataset struct {
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

type QuickBooksObjectDatasetResponse struct {
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

type QuickBooksSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type QuickBooksSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type QuickbaseLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Type                string                            `pulumi:"type"`
	Url                 interface{}                       `pulumi:"url"`
	UserToken           interface{}                       `pulumi:"userToken"`
}

type QuickbaseLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
	Url                 interface{}                               `pulumi:"url"`
	UserToken           interface{}                               `pulumi:"userToken"`
}

type RecurrenceSchedule struct {
	Hours              []int                          `pulumi:"hours"`
	Minutes            []int                          `pulumi:"minutes"`
	MonthDays          []int                          `pulumi:"monthDays"`
	MonthlyOccurrences []RecurrenceScheduleOccurrence `pulumi:"monthlyOccurrences"`
	WeekDays           []DaysOfWeek                   `pulumi:"weekDays"`
}

type RecurrenceScheduleOccurrence struct {
	Day        *DayOfWeek `pulumi:"day"`
	Occurrence *int       `pulumi:"occurrence"`
}

type RecurrenceScheduleOccurrenceResponse struct {
	Day        *string `pulumi:"day"`
	Occurrence *int    `pulumi:"occurrence"`
}

type RecurrenceScheduleResponse struct {
	Hours              []int                                  `pulumi:"hours"`
	Minutes            []int                                  `pulumi:"minutes"`
	MonthDays          []int                                  `pulumi:"monthDays"`
	MonthlyOccurrences []RecurrenceScheduleOccurrenceResponse `pulumi:"monthlyOccurrences"`
	WeekDays           []string                               `pulumi:"weekDays"`
}

type RedirectIncompatibleRowSettings struct {
	LinkedServiceName interface{} `pulumi:"linkedServiceName"`
	Path              interface{} `pulumi:"path"`
}

type RedirectIncompatibleRowSettingsResponse struct {
	LinkedServiceName interface{} `pulumi:"linkedServiceName"`
	Path              interface{} `pulumi:"path"`
}

type RedshiftUnloadSettings struct {
	BucketName          interface{}            `pulumi:"bucketName"`
	S3LinkedServiceName LinkedServiceReference `pulumi:"s3LinkedServiceName"`
}

type RedshiftUnloadSettingsResponse struct {
	BucketName          interface{}                    `pulumi:"bucketName"`
	S3LinkedServiceName LinkedServiceReferenceResponse `pulumi:"s3LinkedServiceName"`
}

type RelationalSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type RelationalSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type RelationalTableDataset struct {
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

type RelationalTableDatasetResponse struct {
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

type RemotePrivateEndpointConnectionResponse struct {
	PrivateEndpoint                   *ArmIdWrapperResponse               `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                              `pulumi:"provisioningState"`
}

type RemotePrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (RemotePrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemotePrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o RemotePrivateEndpointConnectionResponseOutput) ToRemotePrivateEndpointConnectionResponseOutput() RemotePrivateEndpointConnectionResponseOutput {
	return o
}

func (o RemotePrivateEndpointConnectionResponseOutput) ToRemotePrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) RemotePrivateEndpointConnectionResponseOutput {
	return o
}

func (o RemotePrivateEndpointConnectionResponseOutput) PrivateEndpoint() ArmIdWrapperResponsePtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) *ArmIdWrapperResponse { return v.PrivateEndpoint }).(ArmIdWrapperResponsePtrOutput)
}

func (o RemotePrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) *PrivateLinkConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStateResponsePtrOutput)
}

func (o RemotePrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type RerunTumblingWindowTrigger struct {
	Annotations        []interface{} `pulumi:"annotations"`
	Description        *string       `pulumi:"description"`
	ParentTrigger      interface{}   `pulumi:"parentTrigger"`
	RequestedEndTime   string        `pulumi:"requestedEndTime"`
	RequestedStartTime string        `pulumi:"requestedStartTime"`
	RerunConcurrency   int           `pulumi:"rerunConcurrency"`
	Type               string        `pulumi:"type"`
}

type RerunTumblingWindowTriggerResponse struct {
	Annotations        []interface{} `pulumi:"annotations"`
	Description        *string       `pulumi:"description"`
	ParentTrigger      interface{}   `pulumi:"parentTrigger"`
	RequestedEndTime   string        `pulumi:"requestedEndTime"`
	RequestedStartTime string        `pulumi:"requestedStartTime"`
	RerunConcurrency   int           `pulumi:"rerunConcurrency"`
	RuntimeState       string        `pulumi:"runtimeState"`
	Type               string        `pulumi:"type"`
}

type ResponsysLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	ClientId              interface{}                       `pulumi:"clientId"`
	ClientSecret          interface{}                       `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Endpoint              interface{}                       `pulumi:"endpoint"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
}

type ResponsysLinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	ClientId              interface{}                               `pulumi:"clientId"`
	ClientSecret          interface{}                               `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Endpoint              interface{}                               `pulumi:"endpoint"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
}

type ResponsysObjectDataset struct {
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

type ResponsysObjectDatasetResponse struct {
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

type ResponsysSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type ResponsysSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type RestResourceDataset struct {
	AdditionalHeaders interface{}                       `pulumi:"additionalHeaders"`
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	PaginationRules   interface{}                       `pulumi:"paginationRules"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	RelativeUrl       interface{}                       `pulumi:"relativeUrl"`
	RequestBody       interface{}                       `pulumi:"requestBody"`
	RequestMethod     interface{}                       `pulumi:"requestMethod"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type RestResourceDatasetResponse struct {
	AdditionalHeaders interface{}                               `pulumi:"additionalHeaders"`
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	PaginationRules   interface{}                               `pulumi:"paginationRules"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	RelativeUrl       interface{}                               `pulumi:"relativeUrl"`
	RequestBody       interface{}                               `pulumi:"requestBody"`
	RequestMethod     interface{}                               `pulumi:"requestMethod"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type RestServiceLinkedService struct {
	AadResourceId                     interface{}                       `pulumi:"aadResourceId"`
	Annotations                       []interface{}                     `pulumi:"annotations"`
	AuthHeaders                       interface{}                       `pulumi:"authHeaders"`
	AuthenticationType                string                            `pulumi:"authenticationType"`
	AzureCloudType                    interface{}                       `pulumi:"azureCloudType"`
	ClientId                          interface{}                       `pulumi:"clientId"`
	ClientSecret                      interface{}                       `pulumi:"clientSecret"`
	ConnectVia                        *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Credential                        *CredentialReference              `pulumi:"credential"`
	Description                       *string                           `pulumi:"description"`
	EnableServerCertificateValidation interface{}                       `pulumi:"enableServerCertificateValidation"`
	EncryptedCredential               interface{}                       `pulumi:"encryptedCredential"`
	Parameters                        map[string]ParameterSpecification `pulumi:"parameters"`
	Password                          interface{}                       `pulumi:"password"`
	Resource                          interface{}                       `pulumi:"resource"`
	Scope                             interface{}                       `pulumi:"scope"`
	ServicePrincipalId                interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey               interface{}                       `pulumi:"servicePrincipalKey"`
	Tenant                            interface{}                       `pulumi:"tenant"`
	TokenEndpoint                     interface{}                       `pulumi:"tokenEndpoint"`
	Type                              string                            `pulumi:"type"`
	Url                               interface{}                       `pulumi:"url"`
	UserName                          interface{}                       `pulumi:"userName"`
}

type RestServiceLinkedServiceResponse struct {
	AadResourceId                     interface{}                               `pulumi:"aadResourceId"`
	Annotations                       []interface{}                             `pulumi:"annotations"`
	AuthHeaders                       interface{}                               `pulumi:"authHeaders"`
	AuthenticationType                string                                    `pulumi:"authenticationType"`
	AzureCloudType                    interface{}                               `pulumi:"azureCloudType"`
	ClientId                          interface{}                               `pulumi:"clientId"`
	ClientSecret                      interface{}                               `pulumi:"clientSecret"`
	ConnectVia                        *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Credential                        *CredentialReferenceResponse              `pulumi:"credential"`
	Description                       *string                                   `pulumi:"description"`
	EnableServerCertificateValidation interface{}                               `pulumi:"enableServerCertificateValidation"`
	EncryptedCredential               interface{}                               `pulumi:"encryptedCredential"`
	Parameters                        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                          interface{}                               `pulumi:"password"`
	Resource                          interface{}                               `pulumi:"resource"`
	Scope                             interface{}                               `pulumi:"scope"`
	ServicePrincipalId                interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey               interface{}                               `pulumi:"servicePrincipalKey"`
	Tenant                            interface{}                               `pulumi:"tenant"`
	TokenEndpoint                     interface{}                               `pulumi:"tokenEndpoint"`
	Type                              string                                    `pulumi:"type"`
	Url                               interface{}                               `pulumi:"url"`
	UserName                          interface{}                               `pulumi:"userName"`
}

type RestSink struct {
	AdditionalHeaders        interface{} `pulumi:"additionalHeaders"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpCompressionType      interface{} `pulumi:"httpCompressionType"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	RequestInterval          interface{} `pulumi:"requestInterval"`
	RequestMethod            interface{} `pulumi:"requestMethod"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type RestSinkResponse struct {
	AdditionalHeaders        interface{} `pulumi:"additionalHeaders"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpCompressionType      interface{} `pulumi:"httpCompressionType"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	RequestInterval          interface{} `pulumi:"requestInterval"`
	RequestMethod            interface{} `pulumi:"requestMethod"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type RestSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	AdditionalHeaders        interface{} `pulumi:"additionalHeaders"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PaginationRules          interface{} `pulumi:"paginationRules"`
	RequestBody              interface{} `pulumi:"requestBody"`
	RequestInterval          interface{} `pulumi:"requestInterval"`
	RequestMethod            interface{} `pulumi:"requestMethod"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type RestSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	AdditionalHeaders        interface{} `pulumi:"additionalHeaders"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PaginationRules          interface{} `pulumi:"paginationRules"`
	RequestBody              interface{} `pulumi:"requestBody"`
	RequestInterval          interface{} `pulumi:"requestInterval"`
	RequestMethod            interface{} `pulumi:"requestMethod"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type RetryPolicy struct {
	Count             interface{} `pulumi:"count"`
	IntervalInSeconds *int        `pulumi:"intervalInSeconds"`
}

type RetryPolicyResponse struct {
	Count             interface{} `pulumi:"count"`
	IntervalInSeconds *int        `pulumi:"intervalInSeconds"`
}

type SSISAccessCredential struct {
	Domain   interface{} `pulumi:"domain"`
	Password interface{} `pulumi:"password"`
	UserName interface{} `pulumi:"userName"`
}

type SSISAccessCredentialResponse struct {
	Domain   interface{} `pulumi:"domain"`
	Password interface{} `pulumi:"password"`
	UserName interface{} `pulumi:"userName"`
}

type SSISChildPackage struct {
	PackageContent          interface{} `pulumi:"packageContent"`
	PackageLastModifiedDate *string     `pulumi:"packageLastModifiedDate"`
	PackageName             *string     `pulumi:"packageName"`
	PackagePath             interface{} `pulumi:"packagePath"`
}

type SSISChildPackageResponse struct {
	PackageContent          interface{} `pulumi:"packageContent"`
	PackageLastModifiedDate *string     `pulumi:"packageLastModifiedDate"`
	PackageName             *string     `pulumi:"packageName"`
	PackagePath             interface{} `pulumi:"packagePath"`
}

type SSISExecutionCredential struct {
	Domain   interface{}  `pulumi:"domain"`
	Password SecureString `pulumi:"password"`
	UserName interface{}  `pulumi:"userName"`
}

type SSISExecutionCredentialResponse struct {
	Domain   interface{}          `pulumi:"domain"`
	Password SecureStringResponse `pulumi:"password"`
	UserName interface{}          `pulumi:"userName"`
}

type SSISExecutionParameter struct {
	Value interface{} `pulumi:"value"`
}

type SSISExecutionParameterResponse struct {
	Value interface{} `pulumi:"value"`
}

type SSISLogLocation struct {
	AccessCredential   *SSISAccessCredential `pulumi:"accessCredential"`
	LogPath            interface{}           `pulumi:"logPath"`
	LogRefreshInterval interface{}           `pulumi:"logRefreshInterval"`
	Type               string                `pulumi:"type"`
}

type SSISLogLocationResponse struct {
	AccessCredential   *SSISAccessCredentialResponse `pulumi:"accessCredential"`
	LogPath            interface{}                   `pulumi:"logPath"`
	LogRefreshInterval interface{}                   `pulumi:"logRefreshInterval"`
	Type               string                        `pulumi:"type"`
}

type SSISPackageLocation struct {
	AccessCredential              *SSISAccessCredential `pulumi:"accessCredential"`
	ChildPackages                 []SSISChildPackage    `pulumi:"childPackages"`
	ConfigurationAccessCredential *SSISAccessCredential `pulumi:"configurationAccessCredential"`
	ConfigurationPath             interface{}           `pulumi:"configurationPath"`
	PackageContent                interface{}           `pulumi:"packageContent"`
	PackageLastModifiedDate       *string               `pulumi:"packageLastModifiedDate"`
	PackageName                   *string               `pulumi:"packageName"`
	PackagePassword               interface{}           `pulumi:"packagePassword"`
	PackagePath                   interface{}           `pulumi:"packagePath"`
	Type                          *string               `pulumi:"type"`
}

type SSISPackageLocationResponse struct {
	AccessCredential              *SSISAccessCredentialResponse `pulumi:"accessCredential"`
	ChildPackages                 []SSISChildPackageResponse    `pulumi:"childPackages"`
	ConfigurationAccessCredential *SSISAccessCredentialResponse `pulumi:"configurationAccessCredential"`
	ConfigurationPath             interface{}                   `pulumi:"configurationPath"`
	PackageContent                interface{}                   `pulumi:"packageContent"`
	PackageLastModifiedDate       *string                       `pulumi:"packageLastModifiedDate"`
	PackageName                   *string                       `pulumi:"packageName"`
	PackagePassword               interface{}                   `pulumi:"packagePassword"`
	PackagePath                   interface{}                   `pulumi:"packagePath"`
	Type                          *string                       `pulumi:"type"`
}

type SSISPropertyOverride struct {
	IsSensitive *bool       `pulumi:"isSensitive"`
	Value       interface{} `pulumi:"value"`
}

type SSISPropertyOverrideResponse struct {
	IsSensitive *bool       `pulumi:"isSensitive"`
	Value       interface{} `pulumi:"value"`
}

type SalesforceLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ApiVersion          interface{}                       `pulumi:"apiVersion"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	EnvironmentUrl      interface{}                       `pulumi:"environmentUrl"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	SecurityToken       interface{}                       `pulumi:"securityToken"`
	Type                string                            `pulumi:"type"`
	Username            interface{}                       `pulumi:"username"`
}

type SalesforceLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ApiVersion          interface{}                               `pulumi:"apiVersion"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	EnvironmentUrl      interface{}                               `pulumi:"environmentUrl"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	SecurityToken       interface{}                               `pulumi:"securityToken"`
	Type                string                                    `pulumi:"type"`
	Username            interface{}                               `pulumi:"username"`
}

type SalesforceMarketingCloudLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	ClientId              interface{}                       `pulumi:"clientId"`
	ClientSecret          interface{}                       `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionProperties  interface{}                       `pulumi:"connectionProperties"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
}

type SalesforceMarketingCloudLinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	ClientId              interface{}                               `pulumi:"clientId"`
	ClientSecret          interface{}                               `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionProperties  interface{}                               `pulumi:"connectionProperties"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
}

type SalesforceMarketingCloudObjectDataset struct {
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

type SalesforceMarketingCloudObjectDatasetResponse struct {
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

type SalesforceMarketingCloudSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SalesforceMarketingCloudSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SalesforceObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	ObjectApiName     interface{}                       `pulumi:"objectApiName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type SalesforceObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	ObjectApiName     interface{}                               `pulumi:"objectApiName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type SalesforceServiceCloudLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ApiVersion          interface{}                       `pulumi:"apiVersion"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	EnvironmentUrl      interface{}                       `pulumi:"environmentUrl"`
	ExtendedProperties  interface{}                       `pulumi:"extendedProperties"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	SecurityToken       interface{}                       `pulumi:"securityToken"`
	Type                string                            `pulumi:"type"`
	Username            interface{}                       `pulumi:"username"`
}

type SalesforceServiceCloudLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ApiVersion          interface{}                               `pulumi:"apiVersion"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	EnvironmentUrl      interface{}                               `pulumi:"environmentUrl"`
	ExtendedProperties  interface{}                               `pulumi:"extendedProperties"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	SecurityToken       interface{}                               `pulumi:"securityToken"`
	Type                string                                    `pulumi:"type"`
	Username            interface{}                               `pulumi:"username"`
}

type SalesforceServiceCloudObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	ObjectApiName     interface{}                       `pulumi:"objectApiName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type SalesforceServiceCloudObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	ObjectApiName     interface{}                               `pulumi:"objectApiName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type SalesforceServiceCloudSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	ExternalIdFieldName      interface{} `pulumi:"externalIdFieldName"`
	IgnoreNullValues         interface{} `pulumi:"ignoreNullValues"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            *string     `pulumi:"writeBehavior"`
}

type SalesforceServiceCloudSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	ExternalIdFieldName      interface{} `pulumi:"externalIdFieldName"`
	IgnoreNullValues         interface{} `pulumi:"ignoreNullValues"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            *string     `pulumi:"writeBehavior"`
}

type SalesforceServiceCloudSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	ReadBehavior             *string     `pulumi:"readBehavior"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SalesforceServiceCloudSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	ReadBehavior             *string     `pulumi:"readBehavior"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SalesforceSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	ExternalIdFieldName      interface{} `pulumi:"externalIdFieldName"`
	IgnoreNullValues         interface{} `pulumi:"ignoreNullValues"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            *string     `pulumi:"writeBehavior"`
}

type SalesforceSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	ExternalIdFieldName      interface{} `pulumi:"externalIdFieldName"`
	IgnoreNullValues         interface{} `pulumi:"ignoreNullValues"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            *string     `pulumi:"writeBehavior"`
}

type SalesforceSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	ReadBehavior             *string     `pulumi:"readBehavior"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SalesforceSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	ReadBehavior             *string     `pulumi:"readBehavior"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SapBWLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ClientId            interface{}                       `pulumi:"clientId"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Server              interface{}                       `pulumi:"server"`
	SystemNumber        interface{}                       `pulumi:"systemNumber"`
	Type                string                            `pulumi:"type"`
	UserName            interface{}                       `pulumi:"userName"`
}

type SapBWLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ClientId            interface{}                               `pulumi:"clientId"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Server              interface{}                               `pulumi:"server"`
	SystemNumber        interface{}                               `pulumi:"systemNumber"`
	Type                string                                    `pulumi:"type"`
	UserName            interface{}                               `pulumi:"userName"`
}

type SapBwCubeDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type SapBwCubeDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type SapBwSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SapBwSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SapCloudForCustomerLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Type                string                            `pulumi:"type"`
	Url                 interface{}                       `pulumi:"url"`
	Username            interface{}                       `pulumi:"username"`
}

type SapCloudForCustomerLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
	Url                 interface{}                               `pulumi:"url"`
	Username            interface{}                               `pulumi:"username"`
}

type SapCloudForCustomerResourceDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Path              interface{}                       `pulumi:"path"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type SapCloudForCustomerResourceDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Path              interface{}                               `pulumi:"path"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type SapCloudForCustomerSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            *string     `pulumi:"writeBehavior"`
}

type SapCloudForCustomerSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            *string     `pulumi:"writeBehavior"`
}

type SapCloudForCustomerSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SapCloudForCustomerSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SapEccLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential *string                           `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Type                string                            `pulumi:"type"`
	Url                 string                            `pulumi:"url"`
	Username            *string                           `pulumi:"username"`
}

type SapEccLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential *string                                   `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
	Url                 string                                    `pulumi:"url"`
	Username            *string                                   `pulumi:"username"`
}

type SapEccResourceDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Path              interface{}                       `pulumi:"path"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type SapEccResourceDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Path              interface{}                               `pulumi:"path"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type SapEccSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SapEccSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SapHanaLinkedService struct {
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
	UserName            interface{}                       `pulumi:"userName"`
}

type SapHanaLinkedServiceResponse struct {
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
	UserName            interface{}                               `pulumi:"userName"`
}

type SapHanaPartitionSettings struct {
	PartitionColumnName interface{} `pulumi:"partitionColumnName"`
}

type SapHanaPartitionSettingsResponse struct {
	PartitionColumnName interface{} `pulumi:"partitionColumnName"`
}

type SapHanaSource struct {
	AdditionalColumns        interface{}               `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}               `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}               `pulumi:"maxConcurrentConnections"`
	PacketSize               interface{}               `pulumi:"packetSize"`
	PartitionOption          interface{}               `pulumi:"partitionOption"`
	PartitionSettings        *SapHanaPartitionSettings `pulumi:"partitionSettings"`
	Query                    interface{}               `pulumi:"query"`
	QueryTimeout             interface{}               `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}               `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}               `pulumi:"sourceRetryWait"`
	Type                     string                    `pulumi:"type"`
}

type SapHanaSourceResponse struct {
	AdditionalColumns        interface{}                       `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}                       `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}                       `pulumi:"maxConcurrentConnections"`
	PacketSize               interface{}                       `pulumi:"packetSize"`
	PartitionOption          interface{}                       `pulumi:"partitionOption"`
	PartitionSettings        *SapHanaPartitionSettingsResponse `pulumi:"partitionSettings"`
	Query                    interface{}                       `pulumi:"query"`
	QueryTimeout             interface{}                       `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}                       `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                       `pulumi:"sourceRetryWait"`
	Type                     string                            `pulumi:"type"`
}

type SapHanaTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Table             interface{}                       `pulumi:"table"`
	Type              string                            `pulumi:"type"`
}

type SapHanaTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Table             interface{}                               `pulumi:"table"`
	Type              string                                    `pulumi:"type"`
}

type SapOdpLinkedService struct {
	Annotations          []interface{}                     `pulumi:"annotations"`
	ClientId             interface{}                       `pulumi:"clientId"`
	ConnectVia           *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description          *string                           `pulumi:"description"`
	EncryptedCredential  interface{}                       `pulumi:"encryptedCredential"`
	Language             interface{}                       `pulumi:"language"`
	LogonGroup           interface{}                       `pulumi:"logonGroup"`
	MessageServer        interface{}                       `pulumi:"messageServer"`
	MessageServerService interface{}                       `pulumi:"messageServerService"`
	Parameters           map[string]ParameterSpecification `pulumi:"parameters"`
	Password             interface{}                       `pulumi:"password"`
	Server               interface{}                       `pulumi:"server"`
	SncLibraryPath       interface{}                       `pulumi:"sncLibraryPath"`
	SncMode              interface{}                       `pulumi:"sncMode"`
	SncMyName            interface{}                       `pulumi:"sncMyName"`
	SncPartnerName       interface{}                       `pulumi:"sncPartnerName"`
	SncQop               interface{}                       `pulumi:"sncQop"`
	SubscriberName       interface{}                       `pulumi:"subscriberName"`
	SystemId             interface{}                       `pulumi:"systemId"`
	SystemNumber         interface{}                       `pulumi:"systemNumber"`
	Type                 string                            `pulumi:"type"`
	UserName             interface{}                       `pulumi:"userName"`
	X509CertificatePath  interface{}                       `pulumi:"x509CertificatePath"`
}

type SapOdpLinkedServiceResponse struct {
	Annotations          []interface{}                             `pulumi:"annotations"`
	ClientId             interface{}                               `pulumi:"clientId"`
	ConnectVia           *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description          *string                                   `pulumi:"description"`
	EncryptedCredential  interface{}                               `pulumi:"encryptedCredential"`
	Language             interface{}                               `pulumi:"language"`
	LogonGroup           interface{}                               `pulumi:"logonGroup"`
	MessageServer        interface{}                               `pulumi:"messageServer"`
	MessageServerService interface{}                               `pulumi:"messageServerService"`
	Parameters           map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password             interface{}                               `pulumi:"password"`
	Server               interface{}                               `pulumi:"server"`
	SncLibraryPath       interface{}                               `pulumi:"sncLibraryPath"`
	SncMode              interface{}                               `pulumi:"sncMode"`
	SncMyName            interface{}                               `pulumi:"sncMyName"`
	SncPartnerName       interface{}                               `pulumi:"sncPartnerName"`
	SncQop               interface{}                               `pulumi:"sncQop"`
	SubscriberName       interface{}                               `pulumi:"subscriberName"`
	SystemId             interface{}                               `pulumi:"systemId"`
	SystemNumber         interface{}                               `pulumi:"systemNumber"`
	Type                 string                                    `pulumi:"type"`
	UserName             interface{}                               `pulumi:"userName"`
	X509CertificatePath  interface{}                               `pulumi:"x509CertificatePath"`
}

type SapOdpResourceDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Context           interface{}                       `pulumi:"context"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	ObjectName        interface{}                       `pulumi:"objectName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type SapOdpResourceDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Context           interface{}                               `pulumi:"context"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	ObjectName        interface{}                               `pulumi:"objectName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type SapOdpSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	ExtractionMode           interface{} `pulumi:"extractionMode"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Projection               interface{} `pulumi:"projection"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	Selection                interface{} `pulumi:"selection"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	SubscriberProcess        interface{} `pulumi:"subscriberProcess"`
	Type                     string      `pulumi:"type"`
}

type SapOdpSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	ExtractionMode           interface{} `pulumi:"extractionMode"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Projection               interface{} `pulumi:"projection"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	Selection                interface{} `pulumi:"selection"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	SubscriberProcess        interface{} `pulumi:"subscriberProcess"`
	Type                     string      `pulumi:"type"`
}

type SapOpenHubLinkedService struct {
	Annotations          []interface{}                     `pulumi:"annotations"`
	ClientId             interface{}                       `pulumi:"clientId"`
	ConnectVia           *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description          *string                           `pulumi:"description"`
	EncryptedCredential  interface{}                       `pulumi:"encryptedCredential"`
	Language             interface{}                       `pulumi:"language"`
	LogonGroup           interface{}                       `pulumi:"logonGroup"`
	MessageServer        interface{}                       `pulumi:"messageServer"`
	MessageServerService interface{}                       `pulumi:"messageServerService"`
	Parameters           map[string]ParameterSpecification `pulumi:"parameters"`
	Password             interface{}                       `pulumi:"password"`
	Server               interface{}                       `pulumi:"server"`
	SystemId             interface{}                       `pulumi:"systemId"`
	SystemNumber         interface{}                       `pulumi:"systemNumber"`
	Type                 string                            `pulumi:"type"`
	UserName             interface{}                       `pulumi:"userName"`
}

type SapOpenHubLinkedServiceResponse struct {
	Annotations          []interface{}                             `pulumi:"annotations"`
	ClientId             interface{}                               `pulumi:"clientId"`
	ConnectVia           *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description          *string                                   `pulumi:"description"`
	EncryptedCredential  interface{}                               `pulumi:"encryptedCredential"`
	Language             interface{}                               `pulumi:"language"`
	LogonGroup           interface{}                               `pulumi:"logonGroup"`
	MessageServer        interface{}                               `pulumi:"messageServer"`
	MessageServerService interface{}                               `pulumi:"messageServerService"`
	Parameters           map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password             interface{}                               `pulumi:"password"`
	Server               interface{}                               `pulumi:"server"`
	SystemId             interface{}                               `pulumi:"systemId"`
	SystemNumber         interface{}                               `pulumi:"systemNumber"`
	Type                 string                                    `pulumi:"type"`
	UserName             interface{}                               `pulumi:"userName"`
}

type SapOpenHubSource struct {
	AdditionalColumns                interface{} `pulumi:"additionalColumns"`
	BaseRequestId                    interface{} `pulumi:"baseRequestId"`
	CustomRfcReadTableFunctionModule interface{} `pulumi:"customRfcReadTableFunctionModule"`
	DisableMetricsCollection         interface{} `pulumi:"disableMetricsCollection"`
	ExcludeLastRequest               interface{} `pulumi:"excludeLastRequest"`
	MaxConcurrentConnections         interface{} `pulumi:"maxConcurrentConnections"`
	QueryTimeout                     interface{} `pulumi:"queryTimeout"`
	SapDataColumnDelimiter           interface{} `pulumi:"sapDataColumnDelimiter"`
	SourceRetryCount                 interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait                  interface{} `pulumi:"sourceRetryWait"`
	Type                             string      `pulumi:"type"`
}

type SapOpenHubSourceResponse struct {
	AdditionalColumns                interface{} `pulumi:"additionalColumns"`
	BaseRequestId                    interface{} `pulumi:"baseRequestId"`
	CustomRfcReadTableFunctionModule interface{} `pulumi:"customRfcReadTableFunctionModule"`
	DisableMetricsCollection         interface{} `pulumi:"disableMetricsCollection"`
	ExcludeLastRequest               interface{} `pulumi:"excludeLastRequest"`
	MaxConcurrentConnections         interface{} `pulumi:"maxConcurrentConnections"`
	QueryTimeout                     interface{} `pulumi:"queryTimeout"`
	SapDataColumnDelimiter           interface{} `pulumi:"sapDataColumnDelimiter"`
	SourceRetryCount                 interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait                  interface{} `pulumi:"sourceRetryWait"`
	Type                             string      `pulumi:"type"`
}

type SapOpenHubTableDataset struct {
	Annotations            []interface{}                     `pulumi:"annotations"`
	BaseRequestId          interface{}                       `pulumi:"baseRequestId"`
	Description            *string                           `pulumi:"description"`
	ExcludeLastRequest     interface{}                       `pulumi:"excludeLastRequest"`
	Folder                 *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName      LinkedServiceReference            `pulumi:"linkedServiceName"`
	OpenHubDestinationName interface{}                       `pulumi:"openHubDestinationName"`
	Parameters             map[string]ParameterSpecification `pulumi:"parameters"`
	Schema                 interface{}                       `pulumi:"schema"`
	Structure              interface{}                       `pulumi:"structure"`
	Type                   string                            `pulumi:"type"`
}

type SapOpenHubTableDatasetResponse struct {
	Annotations            []interface{}                             `pulumi:"annotations"`
	BaseRequestId          interface{}                               `pulumi:"baseRequestId"`
	Description            *string                                   `pulumi:"description"`
	ExcludeLastRequest     interface{}                               `pulumi:"excludeLastRequest"`
	Folder                 *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName      LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	OpenHubDestinationName interface{}                               `pulumi:"openHubDestinationName"`
	Parameters             map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema                 interface{}                               `pulumi:"schema"`
	Structure              interface{}                               `pulumi:"structure"`
	Type                   string                                    `pulumi:"type"`
}

type SapTableLinkedService struct {
	Annotations          []interface{}                     `pulumi:"annotations"`
	ClientId             interface{}                       `pulumi:"clientId"`
	ConnectVia           *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description          *string                           `pulumi:"description"`
	EncryptedCredential  interface{}                       `pulumi:"encryptedCredential"`
	Language             interface{}                       `pulumi:"language"`
	LogonGroup           interface{}                       `pulumi:"logonGroup"`
	MessageServer        interface{}                       `pulumi:"messageServer"`
	MessageServerService interface{}                       `pulumi:"messageServerService"`
	Parameters           map[string]ParameterSpecification `pulumi:"parameters"`
	Password             interface{}                       `pulumi:"password"`
	Server               interface{}                       `pulumi:"server"`
	SncLibraryPath       interface{}                       `pulumi:"sncLibraryPath"`
	SncMode              interface{}                       `pulumi:"sncMode"`
	SncMyName            interface{}                       `pulumi:"sncMyName"`
	SncPartnerName       interface{}                       `pulumi:"sncPartnerName"`
	SncQop               interface{}                       `pulumi:"sncQop"`
	SystemId             interface{}                       `pulumi:"systemId"`
	SystemNumber         interface{}                       `pulumi:"systemNumber"`
	Type                 string                            `pulumi:"type"`
	UserName             interface{}                       `pulumi:"userName"`
}

type SapTableLinkedServiceResponse struct {
	Annotations          []interface{}                             `pulumi:"annotations"`
	ClientId             interface{}                               `pulumi:"clientId"`
	ConnectVia           *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description          *string                                   `pulumi:"description"`
	EncryptedCredential  interface{}                               `pulumi:"encryptedCredential"`
	Language             interface{}                               `pulumi:"language"`
	LogonGroup           interface{}                               `pulumi:"logonGroup"`
	MessageServer        interface{}                               `pulumi:"messageServer"`
	MessageServerService interface{}                               `pulumi:"messageServerService"`
	Parameters           map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password             interface{}                               `pulumi:"password"`
	Server               interface{}                               `pulumi:"server"`
	SncLibraryPath       interface{}                               `pulumi:"sncLibraryPath"`
	SncMode              interface{}                               `pulumi:"sncMode"`
	SncMyName            interface{}                               `pulumi:"sncMyName"`
	SncPartnerName       interface{}                               `pulumi:"sncPartnerName"`
	SncQop               interface{}                               `pulumi:"sncQop"`
	SystemId             interface{}                               `pulumi:"systemId"`
	SystemNumber         interface{}                               `pulumi:"systemNumber"`
	Type                 string                                    `pulumi:"type"`
	UserName             interface{}                               `pulumi:"userName"`
}

type SapTablePartitionSettings struct {
	MaxPartitionsNumber interface{} `pulumi:"maxPartitionsNumber"`
	PartitionColumnName interface{} `pulumi:"partitionColumnName"`
	PartitionLowerBound interface{} `pulumi:"partitionLowerBound"`
	PartitionUpperBound interface{} `pulumi:"partitionUpperBound"`
}

type SapTablePartitionSettingsResponse struct {
	MaxPartitionsNumber interface{} `pulumi:"maxPartitionsNumber"`
	PartitionColumnName interface{} `pulumi:"partitionColumnName"`
	PartitionLowerBound interface{} `pulumi:"partitionLowerBound"`
	PartitionUpperBound interface{} `pulumi:"partitionUpperBound"`
}

type SapTableResourceDataset struct {
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

type SapTableResourceDatasetResponse struct {
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

type SapTableSource struct {
	AdditionalColumns                interface{}                `pulumi:"additionalColumns"`
	BatchSize                        interface{}                `pulumi:"batchSize"`
	CustomRfcReadTableFunctionModule interface{}                `pulumi:"customRfcReadTableFunctionModule"`
	DisableMetricsCollection         interface{}                `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections         interface{}                `pulumi:"maxConcurrentConnections"`
	PartitionOption                  interface{}                `pulumi:"partitionOption"`
	PartitionSettings                *SapTablePartitionSettings `pulumi:"partitionSettings"`
	QueryTimeout                     interface{}                `pulumi:"queryTimeout"`
	RfcTableFields                   interface{}                `pulumi:"rfcTableFields"`
	RfcTableOptions                  interface{}                `pulumi:"rfcTableOptions"`
	RowCount                         interface{}                `pulumi:"rowCount"`
	RowSkips                         interface{}                `pulumi:"rowSkips"`
	SapDataColumnDelimiter           interface{}                `pulumi:"sapDataColumnDelimiter"`
	SourceRetryCount                 interface{}                `pulumi:"sourceRetryCount"`
	SourceRetryWait                  interface{}                `pulumi:"sourceRetryWait"`
	Type                             string                     `pulumi:"type"`
}

type SapTableSourceResponse struct {
	AdditionalColumns                interface{}                        `pulumi:"additionalColumns"`
	BatchSize                        interface{}                        `pulumi:"batchSize"`
	CustomRfcReadTableFunctionModule interface{}                        `pulumi:"customRfcReadTableFunctionModule"`
	DisableMetricsCollection         interface{}                        `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections         interface{}                        `pulumi:"maxConcurrentConnections"`
	PartitionOption                  interface{}                        `pulumi:"partitionOption"`
	PartitionSettings                *SapTablePartitionSettingsResponse `pulumi:"partitionSettings"`
	QueryTimeout                     interface{}                        `pulumi:"queryTimeout"`
	RfcTableFields                   interface{}                        `pulumi:"rfcTableFields"`
	RfcTableOptions                  interface{}                        `pulumi:"rfcTableOptions"`
	RowCount                         interface{}                        `pulumi:"rowCount"`
	RowSkips                         interface{}                        `pulumi:"rowSkips"`
	SapDataColumnDelimiter           interface{}                        `pulumi:"sapDataColumnDelimiter"`
	SourceRetryCount                 interface{}                        `pulumi:"sourceRetryCount"`
	SourceRetryWait                  interface{}                        `pulumi:"sourceRetryWait"`
	Type                             string                             `pulumi:"type"`
}

type ScheduleTrigger struct {
	Annotations []interface{}              `pulumi:"annotations"`
	Description *string                    `pulumi:"description"`
	Pipelines   []TriggerPipelineReference `pulumi:"pipelines"`
	Recurrence  ScheduleTriggerRecurrence  `pulumi:"recurrence"`
	Type        string                     `pulumi:"type"`
}

type ScheduleTriggerRecurrence struct {
	EndTime   *string             `pulumi:"endTime"`
	Frequency *string             `pulumi:"frequency"`
	Interval  *int                `pulumi:"interval"`
	Schedule  *RecurrenceSchedule `pulumi:"schedule"`
	StartTime *string             `pulumi:"startTime"`
	TimeZone  *string             `pulumi:"timeZone"`
}

type ScheduleTriggerRecurrenceResponse struct {
	EndTime   *string                     `pulumi:"endTime"`
	Frequency *string                     `pulumi:"frequency"`
	Interval  *int                        `pulumi:"interval"`
	Schedule  *RecurrenceScheduleResponse `pulumi:"schedule"`
	StartTime *string                     `pulumi:"startTime"`
	TimeZone  *string                     `pulumi:"timeZone"`
}

type ScheduleTriggerResponse struct {
	Annotations  []interface{}                      `pulumi:"annotations"`
	Description  *string                            `pulumi:"description"`
	Pipelines    []TriggerPipelineReferenceResponse `pulumi:"pipelines"`
	Recurrence   ScheduleTriggerRecurrenceResponse  `pulumi:"recurrence"`
	RuntimeState string                             `pulumi:"runtimeState"`
	Type         string                             `pulumi:"type"`
}

type ScriptAction struct {
	Name       string      `pulumi:"name"`
	Parameters *string     `pulumi:"parameters"`
	Roles      interface{} `pulumi:"roles"`
	Uri        string      `pulumi:"uri"`
}

type ScriptActionResponse struct {
	Name       string      `pulumi:"name"`
	Parameters *string     `pulumi:"parameters"`
	Roles      interface{} `pulumi:"roles"`
	Uri        string      `pulumi:"uri"`
}

type ScriptActivity struct {
	DependsOn         []ActivityDependency                     `pulumi:"dependsOn"`
	Description       *string                                  `pulumi:"description"`
	LinkedServiceName LinkedServiceReference                   `pulumi:"linkedServiceName"`
	LogSettings       *ScriptActivityTypePropertiesLogSettings `pulumi:"logSettings"`
	Name              string                                   `pulumi:"name"`
	Policy            *ActivityPolicy                          `pulumi:"policy"`
	Scripts           []ScriptActivityScriptBlock              `pulumi:"scripts"`
	Type              string                                   `pulumi:"type"`
	UserProperties    []UserProperty                           `pulumi:"userProperties"`
}

type ScriptActivityParameter struct {
	Direction *string     `pulumi:"direction"`
	Name      interface{} `pulumi:"name"`
	Size      *int        `pulumi:"size"`
	Type      *string     `pulumi:"type"`
	Value     interface{} `pulumi:"value"`
}

type ScriptActivityParameterResponse struct {
	Direction *string     `pulumi:"direction"`
	Name      interface{} `pulumi:"name"`
	Size      *int        `pulumi:"size"`
	Type      *string     `pulumi:"type"`
	Value     interface{} `pulumi:"value"`
}

type ScriptActivityResponse struct {
	DependsOn         []ActivityDependencyResponse                     `pulumi:"dependsOn"`
	Description       *string                                          `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse                   `pulumi:"linkedServiceName"`
	LogSettings       *ScriptActivityTypePropertiesResponseLogSettings `pulumi:"logSettings"`
	Name              string                                           `pulumi:"name"`
	Policy            *ActivityPolicyResponse                          `pulumi:"policy"`
	Scripts           []ScriptActivityScriptBlockResponse              `pulumi:"scripts"`
	Type              string                                           `pulumi:"type"`
	UserProperties    []UserPropertyResponse                           `pulumi:"userProperties"`
}

type ScriptActivityScriptBlock struct {
	Parameters []ScriptActivityParameter `pulumi:"parameters"`
	Text       interface{}               `pulumi:"text"`
	Type       string                    `pulumi:"type"`
}

type ScriptActivityScriptBlockResponse struct {
	Parameters []ScriptActivityParameterResponse `pulumi:"parameters"`
	Text       interface{}                       `pulumi:"text"`
	Type       string                            `pulumi:"type"`
}

type ScriptActivityTypePropertiesLogSettings struct {
	LogDestination      string               `pulumi:"logDestination"`
	LogLocationSettings *LogLocationSettings `pulumi:"logLocationSettings"`
}

type ScriptActivityTypePropertiesResponseLogSettings struct {
	LogDestination      string                       `pulumi:"logDestination"`
	LogLocationSettings *LogLocationSettingsResponse `pulumi:"logLocationSettings"`
}

type SecureString struct {
	Type  string `pulumi:"type"`
	Value string `pulumi:"value"`
}

type SecureStringResponse struct {
	Type  string `pulumi:"type"`
	Value string `pulumi:"value"`
}

type SelfDependencyTumblingWindowTriggerReference struct {
	Offset string  `pulumi:"offset"`
	Size   *string `pulumi:"size"`
	Type   string  `pulumi:"type"`
}

type SelfDependencyTumblingWindowTriggerReferenceResponse struct {
	Offset string  `pulumi:"offset"`
	Size   *string `pulumi:"size"`
	Type   string  `pulumi:"type"`
}

type SelfHostedIntegrationRuntime struct {
	Description *string     `pulumi:"description"`
	LinkedInfo  interface{} `pulumi:"linkedInfo"`
	Type        string      `pulumi:"type"`
}

type SelfHostedIntegrationRuntimeNodeResponse struct {
	Capabilities        map[string]string `pulumi:"capabilities"`
	ConcurrentJobsLimit int               `pulumi:"concurrentJobsLimit"`
	ExpiryTime          string            `pulumi:"expiryTime"`
	HostServiceUri      string            `pulumi:"hostServiceUri"`
	IsActiveDispatcher  bool              `pulumi:"isActiveDispatcher"`
	LastConnectTime     string            `pulumi:"lastConnectTime"`
	LastEndUpdateTime   string            `pulumi:"lastEndUpdateTime"`
	LastStartTime       string            `pulumi:"lastStartTime"`
	LastStartUpdateTime string            `pulumi:"lastStartUpdateTime"`
	LastStopTime        string            `pulumi:"lastStopTime"`
	LastUpdateResult    string            `pulumi:"lastUpdateResult"`
	MachineName         string            `pulumi:"machineName"`
	MaxConcurrentJobs   int               `pulumi:"maxConcurrentJobs"`
	NodeName            string            `pulumi:"nodeName"`
	RegisterTime        string            `pulumi:"registerTime"`
	Status              string            `pulumi:"status"`
	Version             string            `pulumi:"version"`
	VersionStatus       string            `pulumi:"versionStatus"`
}

type SelfHostedIntegrationRuntimeResponse struct {
	Description *string     `pulumi:"description"`
	LinkedInfo  interface{} `pulumi:"linkedInfo"`
	Type        string      `pulumi:"type"`
}

type SelfHostedIntegrationRuntimeStatusResponse struct {
	AutoUpdate                string                                     `pulumi:"autoUpdate"`
	AutoUpdateETA             string                                     `pulumi:"autoUpdateETA"`
	Capabilities              map[string]string                          `pulumi:"capabilities"`
	CreateTime                string                                     `pulumi:"createTime"`
	DataFactoryName           string                                     `pulumi:"dataFactoryName"`
	InternalChannelEncryption string                                     `pulumi:"internalChannelEncryption"`
	LatestVersion             string                                     `pulumi:"latestVersion"`
	Links                     []LinkedIntegrationRuntimeResponse         `pulumi:"links"`
	LocalTimeZoneOffset       string                                     `pulumi:"localTimeZoneOffset"`
	Nodes                     []SelfHostedIntegrationRuntimeNodeResponse `pulumi:"nodes"`
	PushedVersion             string                                     `pulumi:"pushedVersion"`
	ScheduledUpdateDate       string                                     `pulumi:"scheduledUpdateDate"`
	ServiceUrls               []string                                   `pulumi:"serviceUrls"`
	State                     string                                     `pulumi:"state"`
	TaskQueueId               string                                     `pulumi:"taskQueueId"`
	Type                      string                                     `pulumi:"type"`
	UpdateDelayOffset         string                                     `pulumi:"updateDelayOffset"`
	Version                   string                                     `pulumi:"version"`
	VersionStatus             string                                     `pulumi:"versionStatus"`
}

type ServiceNowLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	AuthenticationType    string                            `pulumi:"authenticationType"`
	ClientId              interface{}                       `pulumi:"clientId"`
	ClientSecret          interface{}                       `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Endpoint              interface{}                       `pulumi:"endpoint"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Password              interface{}                       `pulumi:"password"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
	Username              interface{}                       `pulumi:"username"`
}

type ServiceNowLinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	AuthenticationType    string                                    `pulumi:"authenticationType"`
	ClientId              interface{}                               `pulumi:"clientId"`
	ClientSecret          interface{}                               `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Endpoint              interface{}                               `pulumi:"endpoint"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password              interface{}                               `pulumi:"password"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
	Username              interface{}                               `pulumi:"username"`
}

type ServiceNowObjectDataset struct {
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

type ServiceNowObjectDatasetResponse struct {
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

type ServiceNowSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type ServiceNowSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SetVariableActivity struct {
	DependsOn      []ActivityDependency `pulumi:"dependsOn"`
	Description    *string              `pulumi:"description"`
	Name           string               `pulumi:"name"`
	Type           string               `pulumi:"type"`
	UserProperties []UserProperty       `pulumi:"userProperties"`
	Value          interface{}          `pulumi:"value"`
	VariableName   *string              `pulumi:"variableName"`
}

type SetVariableActivityResponse struct {
	DependsOn      []ActivityDependencyResponse `pulumi:"dependsOn"`
	Description    *string                      `pulumi:"description"`
	Name           string                       `pulumi:"name"`
	Type           string                       `pulumi:"type"`
	UserProperties []UserPropertyResponse       `pulumi:"userProperties"`
	Value          interface{}                  `pulumi:"value"`
	VariableName   *string                      `pulumi:"variableName"`
}

type SftpLocation struct {
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
}

type SftpLocationResponse struct {
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
}

type SftpReadSettings struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableChunking            interface{} `pulumi:"disableChunking"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type SftpReadSettingsResponse struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableChunking            interface{} `pulumi:"disableChunking"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type SftpServerLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	AuthenticationType    *string                           `pulumi:"authenticationType"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Host                  interface{}                       `pulumi:"host"`
	HostKeyFingerprint    interface{}                       `pulumi:"hostKeyFingerprint"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	PassPhrase            interface{}                       `pulumi:"passPhrase"`
	Password              interface{}                       `pulumi:"password"`
	Port                  interface{}                       `pulumi:"port"`
	PrivateKeyContent     interface{}                       `pulumi:"privateKeyContent"`
	PrivateKeyPath        interface{}                       `pulumi:"privateKeyPath"`
	SkipHostKeyValidation interface{}                       `pulumi:"skipHostKeyValidation"`
	Type                  string                            `pulumi:"type"`
	UserName              interface{}                       `pulumi:"userName"`
}

type SftpServerLinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	AuthenticationType    *string                                   `pulumi:"authenticationType"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Host                  interface{}                               `pulumi:"host"`
	HostKeyFingerprint    interface{}                               `pulumi:"hostKeyFingerprint"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	PassPhrase            interface{}                               `pulumi:"passPhrase"`
	Password              interface{}                               `pulumi:"password"`
	Port                  interface{}                               `pulumi:"port"`
	PrivateKeyContent     interface{}                               `pulumi:"privateKeyContent"`
	PrivateKeyPath        interface{}                               `pulumi:"privateKeyPath"`
	SkipHostKeyValidation interface{}                               `pulumi:"skipHostKeyValidation"`
	Type                  string                                    `pulumi:"type"`
	UserName              interface{}                               `pulumi:"userName"`
}

type SftpWriteSettings struct {
	CopyBehavior             interface{} `pulumi:"copyBehavior"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	OperationTimeout         interface{} `pulumi:"operationTimeout"`
	Type                     string      `pulumi:"type"`
	UseTempFileRename        interface{} `pulumi:"useTempFileRename"`
}

type SftpWriteSettingsResponse struct {
	CopyBehavior             interface{} `pulumi:"copyBehavior"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	OperationTimeout         interface{} `pulumi:"operationTimeout"`
	Type                     string      `pulumi:"type"`
	UseTempFileRename        interface{} `pulumi:"useTempFileRename"`
}

type SharePointOnlineListLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	ServicePrincipalId  interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey interface{}                       `pulumi:"servicePrincipalKey"`
	SiteUrl             interface{}                       `pulumi:"siteUrl"`
	TenantId            interface{}                       `pulumi:"tenantId"`
	Type                string                            `pulumi:"type"`
}

type SharePointOnlineListLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	ServicePrincipalId  interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey interface{}                               `pulumi:"servicePrincipalKey"`
	SiteUrl             interface{}                               `pulumi:"siteUrl"`
	TenantId            interface{}                               `pulumi:"tenantId"`
	Type                string                                    `pulumi:"type"`
}

type SharePointOnlineListResourceDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	ListName          interface{}                       `pulumi:"listName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type SharePointOnlineListResourceDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	ListName          interface{}                               `pulumi:"listName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type SharePointOnlineListSource struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SharePointOnlineListSourceResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type ShopifyLinkedService struct {
	AccessToken           interface{}                       `pulumi:"accessToken"`
	Annotations           []interface{}                     `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Host                  interface{}                       `pulumi:"host"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
}

type ShopifyLinkedServiceResponse struct {
	AccessToken           interface{}                               `pulumi:"accessToken"`
	Annotations           []interface{}                             `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Host                  interface{}                               `pulumi:"host"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
}

type ShopifyObjectDataset struct {
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

type ShopifyObjectDatasetResponse struct {
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

type ShopifySource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type ShopifySourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SkipErrorFile struct {
	DataInconsistency interface{} `pulumi:"dataInconsistency"`
	FileMissing       interface{} `pulumi:"fileMissing"`
}

type SkipErrorFileResponse struct {
	DataInconsistency interface{} `pulumi:"dataInconsistency"`
	FileMissing       interface{} `pulumi:"fileMissing"`
}

type SmartsheetLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ApiToken            interface{}                       `pulumi:"apiToken"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Type                string                            `pulumi:"type"`
}

type SmartsheetLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ApiToken            interface{}                               `pulumi:"apiToken"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
}

type SnowflakeDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Table             interface{}                       `pulumi:"table"`
	Type              string                            `pulumi:"type"`
}

type SnowflakeDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Table             interface{}                               `pulumi:"table"`
	Type              string                                    `pulumi:"type"`
}

type SnowflakeExportCopyCommand struct {
	AdditionalCopyOptions   map[string]interface{} `pulumi:"additionalCopyOptions"`
	AdditionalFormatOptions map[string]interface{} `pulumi:"additionalFormatOptions"`
	Type                    string                 `pulumi:"type"`
}

type SnowflakeExportCopyCommandResponse struct {
	AdditionalCopyOptions   map[string]interface{} `pulumi:"additionalCopyOptions"`
	AdditionalFormatOptions map[string]interface{} `pulumi:"additionalFormatOptions"`
	Type                    string                 `pulumi:"type"`
}

type SnowflakeImportCopyCommand struct {
	AdditionalCopyOptions   map[string]interface{} `pulumi:"additionalCopyOptions"`
	AdditionalFormatOptions map[string]interface{} `pulumi:"additionalFormatOptions"`
	Type                    string                 `pulumi:"type"`
}

type SnowflakeImportCopyCommandResponse struct {
	AdditionalCopyOptions   map[string]interface{} `pulumi:"additionalCopyOptions"`
	AdditionalFormatOptions map[string]interface{} `pulumi:"additionalFormatOptions"`
	Type                    string                 `pulumi:"type"`
}

type SnowflakeLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReference     `pulumi:"password"`
	Type                string                            `pulumi:"type"`
}

type SnowflakeLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReferenceResponse     `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
}

type SnowflakeSink struct {
	DisableMetricsCollection interface{}                 `pulumi:"disableMetricsCollection"`
	ImportSettings           *SnowflakeImportCopyCommand `pulumi:"importSettings"`
	MaxConcurrentConnections interface{}                 `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{}                 `pulumi:"preCopyScript"`
	SinkRetryCount           interface{}                 `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}                 `pulumi:"sinkRetryWait"`
	Type                     string                      `pulumi:"type"`
	WriteBatchSize           interface{}                 `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}                 `pulumi:"writeBatchTimeout"`
}

type SnowflakeSinkResponse struct {
	DisableMetricsCollection interface{}                         `pulumi:"disableMetricsCollection"`
	ImportSettings           *SnowflakeImportCopyCommandResponse `pulumi:"importSettings"`
	MaxConcurrentConnections interface{}                         `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{}                         `pulumi:"preCopyScript"`
	SinkRetryCount           interface{}                         `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}                         `pulumi:"sinkRetryWait"`
	Type                     string                              `pulumi:"type"`
	WriteBatchSize           interface{}                         `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}                         `pulumi:"writeBatchTimeout"`
}

type SnowflakeSource struct {
	DisableMetricsCollection interface{}                 `pulumi:"disableMetricsCollection"`
	ExportSettings           *SnowflakeExportCopyCommand `pulumi:"exportSettings"`
	MaxConcurrentConnections interface{}                 `pulumi:"maxConcurrentConnections"`
	Query                    interface{}                 `pulumi:"query"`
	SourceRetryCount         interface{}                 `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                 `pulumi:"sourceRetryWait"`
	Type                     string                      `pulumi:"type"`
}

type SnowflakeSourceResponse struct {
	DisableMetricsCollection interface{}                         `pulumi:"disableMetricsCollection"`
	ExportSettings           *SnowflakeExportCopyCommandResponse `pulumi:"exportSettings"`
	MaxConcurrentConnections interface{}                         `pulumi:"maxConcurrentConnections"`
	Query                    interface{}                         `pulumi:"query"`
	SourceRetryCount         interface{}                         `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                         `pulumi:"sourceRetryWait"`
	Type                     string                              `pulumi:"type"`
}

type SparkLinkedService struct {
	AllowHostNameCNMismatch   interface{}                       `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                       `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                     `pulumi:"annotations"`
	AuthenticationType        string                            `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description               *string                           `pulumi:"description"`
	EnableSsl                 interface{}                       `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                       `pulumi:"encryptedCredential"`
	Host                      interface{}                       `pulumi:"host"`
	HttpPath                  interface{}                       `pulumi:"httpPath"`
	Parameters                map[string]ParameterSpecification `pulumi:"parameters"`
	Password                  interface{}                       `pulumi:"password"`
	Port                      interface{}                       `pulumi:"port"`
	ServerType                *string                           `pulumi:"serverType"`
	ThriftTransportProtocol   *string                           `pulumi:"thriftTransportProtocol"`
	TrustedCertPath           interface{}                       `pulumi:"trustedCertPath"`
	Type                      string                            `pulumi:"type"`
	UseSystemTrustStore       interface{}                       `pulumi:"useSystemTrustStore"`
	Username                  interface{}                       `pulumi:"username"`
}

type SparkLinkedServiceResponse struct {
	AllowHostNameCNMismatch   interface{}                               `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                               `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                             `pulumi:"annotations"`
	AuthenticationType        string                                    `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description               *string                                   `pulumi:"description"`
	EnableSsl                 interface{}                               `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                               `pulumi:"encryptedCredential"`
	Host                      interface{}                               `pulumi:"host"`
	HttpPath                  interface{}                               `pulumi:"httpPath"`
	Parameters                map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                  interface{}                               `pulumi:"password"`
	Port                      interface{}                               `pulumi:"port"`
	ServerType                *string                                   `pulumi:"serverType"`
	ThriftTransportProtocol   *string                                   `pulumi:"thriftTransportProtocol"`
	TrustedCertPath           interface{}                               `pulumi:"trustedCertPath"`
	Type                      string                                    `pulumi:"type"`
	UseSystemTrustStore       interface{}                               `pulumi:"useSystemTrustStore"`
	Username                  interface{}                               `pulumi:"username"`
}

type SparkObjectDataset struct {
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

type SparkObjectDatasetResponse struct {
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

type SparkSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SparkSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SqlAlwaysEncryptedProperties struct {
	AlwaysEncryptedAkvAuthType string               `pulumi:"alwaysEncryptedAkvAuthType"`
	Credential                 *CredentialReference `pulumi:"credential"`
	ServicePrincipalId         interface{}          `pulumi:"servicePrincipalId"`
	ServicePrincipalKey        interface{}          `pulumi:"servicePrincipalKey"`
}

type SqlAlwaysEncryptedPropertiesResponse struct {
	AlwaysEncryptedAkvAuthType string                       `pulumi:"alwaysEncryptedAkvAuthType"`
	Credential                 *CredentialReferenceResponse `pulumi:"credential"`
	ServicePrincipalId         interface{}                  `pulumi:"servicePrincipalId"`
	ServicePrincipalKey        interface{}                  `pulumi:"servicePrincipalKey"`
}

type SqlDWSink struct {
	AllowCopyCommand         interface{}            `pulumi:"allowCopyCommand"`
	AllowPolyBase            interface{}            `pulumi:"allowPolyBase"`
	CopyCommandSettings      *DWCopyCommandSettings `pulumi:"copyCommandSettings"`
	DisableMetricsCollection interface{}            `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}            `pulumi:"maxConcurrentConnections"`
	PolyBaseSettings         *PolybaseSettings      `pulumi:"polyBaseSettings"`
	PreCopyScript            interface{}            `pulumi:"preCopyScript"`
	SinkRetryCount           interface{}            `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}            `pulumi:"sinkRetryWait"`
	SqlWriterUseTableLock    interface{}            `pulumi:"sqlWriterUseTableLock"`
	TableOption              interface{}            `pulumi:"tableOption"`
	Type                     string                 `pulumi:"type"`
	UpsertSettings           *SqlDWUpsertSettings   `pulumi:"upsertSettings"`
	WriteBatchSize           interface{}            `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}            `pulumi:"writeBatchTimeout"`
	WriteBehavior            interface{}            `pulumi:"writeBehavior"`
}

type SqlDWSinkResponse struct {
	AllowCopyCommand         interface{}                    `pulumi:"allowCopyCommand"`
	AllowPolyBase            interface{}                    `pulumi:"allowPolyBase"`
	CopyCommandSettings      *DWCopyCommandSettingsResponse `pulumi:"copyCommandSettings"`
	DisableMetricsCollection interface{}                    `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}                    `pulumi:"maxConcurrentConnections"`
	PolyBaseSettings         *PolybaseSettingsResponse      `pulumi:"polyBaseSettings"`
	PreCopyScript            interface{}                    `pulumi:"preCopyScript"`
	SinkRetryCount           interface{}                    `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}                    `pulumi:"sinkRetryWait"`
	SqlWriterUseTableLock    interface{}                    `pulumi:"sqlWriterUseTableLock"`
	TableOption              interface{}                    `pulumi:"tableOption"`
	Type                     string                         `pulumi:"type"`
	UpsertSettings           *SqlDWUpsertSettingsResponse   `pulumi:"upsertSettings"`
	WriteBatchSize           interface{}                    `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}                    `pulumi:"writeBatchTimeout"`
	WriteBehavior            interface{}                    `pulumi:"writeBehavior"`
}

type SqlDWSource struct {
	AdditionalColumns            interface{}           `pulumi:"additionalColumns"`
	DisableMetricsCollection     interface{}           `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections     interface{}           `pulumi:"maxConcurrentConnections"`
	PartitionOption              interface{}           `pulumi:"partitionOption"`
	PartitionSettings            *SqlPartitionSettings `pulumi:"partitionSettings"`
	QueryTimeout                 interface{}           `pulumi:"queryTimeout"`
	SourceRetryCount             interface{}           `pulumi:"sourceRetryCount"`
	SourceRetryWait              interface{}           `pulumi:"sourceRetryWait"`
	SqlReaderQuery               interface{}           `pulumi:"sqlReaderQuery"`
	SqlReaderStoredProcedureName interface{}           `pulumi:"sqlReaderStoredProcedureName"`
	StoredProcedureParameters    interface{}           `pulumi:"storedProcedureParameters"`
	Type                         string                `pulumi:"type"`
}

type SqlDWSourceResponse struct {
	AdditionalColumns            interface{}                   `pulumi:"additionalColumns"`
	DisableMetricsCollection     interface{}                   `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections     interface{}                   `pulumi:"maxConcurrentConnections"`
	PartitionOption              interface{}                   `pulumi:"partitionOption"`
	PartitionSettings            *SqlPartitionSettingsResponse `pulumi:"partitionSettings"`
	QueryTimeout                 interface{}                   `pulumi:"queryTimeout"`
	SourceRetryCount             interface{}                   `pulumi:"sourceRetryCount"`
	SourceRetryWait              interface{}                   `pulumi:"sourceRetryWait"`
	SqlReaderQuery               interface{}                   `pulumi:"sqlReaderQuery"`
	SqlReaderStoredProcedureName interface{}                   `pulumi:"sqlReaderStoredProcedureName"`
	StoredProcedureParameters    interface{}                   `pulumi:"storedProcedureParameters"`
	Type                         string                        `pulumi:"type"`
}

type SqlDWUpsertSettings struct {
	InterimSchemaName interface{} `pulumi:"interimSchemaName"`
	Keys              interface{} `pulumi:"keys"`
}

type SqlDWUpsertSettingsResponse struct {
	InterimSchemaName interface{} `pulumi:"interimSchemaName"`
	Keys              interface{} `pulumi:"keys"`
}

type SqlMISink struct {
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

type SqlMISinkResponse struct {
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
	pulumi.RegisterOutputType(ArmIdWrapperResponseOutput{})
	pulumi.RegisterOutputType(ArmIdWrapperResponsePtrOutput{})
	pulumi.RegisterOutputType(CMKIdentityDefinitionOutput{})
	pulumi.RegisterOutputType(CMKIdentityDefinitionPtrOutput{})
	pulumi.RegisterOutputType(CMKIdentityDefinitionResponseOutput{})
	pulumi.RegisterOutputType(CMKIdentityDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectionStatePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConnectionStatePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionConfigurationOutput{})
	pulumi.RegisterOutputType(EncryptionConfigurationPtrOutput{})
	pulumi.RegisterOutputType(EncryptionConfigurationResponseOutput{})
	pulumi.RegisterOutputType(EncryptionConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(FactoryIdentityOutput{})
	pulumi.RegisterOutputType(FactoryIdentityPtrOutput{})
	pulumi.RegisterOutputType(FactoryIdentityResponseOutput{})
	pulumi.RegisterOutputType(FactoryIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(GitHubClientSecretOutput{})
	pulumi.RegisterOutputType(GitHubClientSecretPtrOutput{})
	pulumi.RegisterOutputType(GlobalParameterSpecificationOutput{})
	pulumi.RegisterOutputType(GlobalParameterSpecificationMapOutput{})
	pulumi.RegisterOutputType(GlobalParameterSpecificationResponseOutput{})
	pulumi.RegisterOutputType(GlobalParameterSpecificationResponseMapOutput{})
	pulumi.RegisterOutputType(ManagedPrivateEndpointTypeOutput{})
	pulumi.RegisterOutputType(ManagedPrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(ParameterSpecificationOutput{})
	pulumi.RegisterOutputType(ParameterSpecificationMapOutput{})
	pulumi.RegisterOutputType(ParameterSpecificationResponseOutput{})
	pulumi.RegisterOutputType(ParameterSpecificationResponseMapOutput{})
	pulumi.RegisterOutputType(PipelineElapsedTimeMetricPolicyOutput{})
	pulumi.RegisterOutputType(PipelineElapsedTimeMetricPolicyPtrOutput{})
	pulumi.RegisterOutputType(PipelineElapsedTimeMetricPolicyResponseOutput{})
	pulumi.RegisterOutputType(PipelineElapsedTimeMetricPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(PipelineFolderOutput{})
	pulumi.RegisterOutputType(PipelineFolderPtrOutput{})
	pulumi.RegisterOutputType(PipelinePolicyOutput{})
	pulumi.RegisterOutputType(PipelinePolicyPtrOutput{})
	pulumi.RegisterOutputType(PipelinePolicyResponseOutput{})
	pulumi.RegisterOutputType(PipelinePolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(PipelineResponseFolderOutput{})
	pulumi.RegisterOutputType(PipelineResponseFolderPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionApprovalRequestOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionApprovalRequestPtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(PurviewConfigurationOutput{})
	pulumi.RegisterOutputType(PurviewConfigurationPtrOutput{})
	pulumi.RegisterOutputType(PurviewConfigurationResponseOutput{})
	pulumi.RegisterOutputType(PurviewConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(RemotePrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(UserAccessPolicyResponseOutput{})
	pulumi.RegisterOutputType(UserAccessPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(VariableSpecificationOutput{})
	pulumi.RegisterOutputType(VariableSpecificationMapOutput{})
	pulumi.RegisterOutputType(VariableSpecificationResponseOutput{})
	pulumi.RegisterOutputType(VariableSpecificationResponseMapOutput{})
}
