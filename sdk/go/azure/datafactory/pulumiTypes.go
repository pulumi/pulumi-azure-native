


package datafactory

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
	AdditionalColumns            interface{}           `pulumi:"additionalColumns"`
	DisableMetricsCollection     interface{}           `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections     interface{}           `pulumi:"maxConcurrentConnections"`
	PartitionOption              interface{}           `pulumi:"partitionOption"`
	PartitionSettings            *SqlPartitionSettings `pulumi:"partitionSettings"`
	ProduceAdditionalTypes       interface{}           `pulumi:"produceAdditionalTypes"`
	QueryTimeout                 interface{}           `pulumi:"queryTimeout"`
	SourceRetryCount             interface{}           `pulumi:"sourceRetryCount"`
	SourceRetryWait              interface{}           `pulumi:"sourceRetryWait"`
	SqlReaderQuery               interface{}           `pulumi:"sqlReaderQuery"`
	SqlReaderStoredProcedureName interface{}           `pulumi:"sqlReaderStoredProcedureName"`
	StoredProcedureParameters    interface{}           `pulumi:"storedProcedureParameters"`
	Type                         string                `pulumi:"type"`
}

type AmazonRdsForSqlServerSourceResponse struct {
	AdditionalColumns            interface{}                   `pulumi:"additionalColumns"`
	DisableMetricsCollection     interface{}                   `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections     interface{}                   `pulumi:"maxConcurrentConnections"`
	PartitionOption              interface{}                   `pulumi:"partitionOption"`
	PartitionSettings            *SqlPartitionSettingsResponse `pulumi:"partitionSettings"`
	ProduceAdditionalTypes       interface{}                   `pulumi:"produceAdditionalTypes"`
	QueryTimeout                 interface{}                   `pulumi:"queryTimeout"`
	SourceRetryCount             interface{}                   `pulumi:"sourceRetryCount"`
	SourceRetryWait              interface{}                   `pulumi:"sourceRetryWait"`
	SqlReaderQuery               interface{}                   `pulumi:"sqlReaderQuery"`
	SqlReaderStoredProcedureName interface{}                   `pulumi:"sqlReaderStoredProcedureName"`
	StoredProcedureParameters    interface{}                   `pulumi:"storedProcedureParameters"`
	Type                         string                        `pulumi:"type"`
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
	DisableMetricsCollection              interface{}        `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections              interface{}        `pulumi:"maxConcurrentConnections"`
	PreCopyScript                         interface{}        `pulumi:"preCopyScript"`
	SinkRetryCount                        interface{}        `pulumi:"sinkRetryCount"`
	SinkRetryWait                         interface{}        `pulumi:"sinkRetryWait"`
	SqlWriterStoredProcedureName          interface{}        `pulumi:"sqlWriterStoredProcedureName"`
	SqlWriterTableType                    interface{}        `pulumi:"sqlWriterTableType"`
	SqlWriterUseTableLock                 interface{}        `pulumi:"sqlWriterUseTableLock"`
	StoredProcedureParameters             interface{}        `pulumi:"storedProcedureParameters"`
	StoredProcedureTableTypeParameterName interface{}        `pulumi:"storedProcedureTableTypeParameterName"`
	TableOption                           interface{}        `pulumi:"tableOption"`
	Type                                  string             `pulumi:"type"`
	UpsertSettings                        *SqlUpsertSettings `pulumi:"upsertSettings"`
	WriteBatchSize                        interface{}        `pulumi:"writeBatchSize"`
	WriteBatchTimeout                     interface{}        `pulumi:"writeBatchTimeout"`
	WriteBehavior                         interface{}        `pulumi:"writeBehavior"`
}

type AzureSqlSinkResponse struct {
	DisableMetricsCollection              interface{}                `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections              interface{}                `pulumi:"maxConcurrentConnections"`
	PreCopyScript                         interface{}                `pulumi:"preCopyScript"`
	SinkRetryCount                        interface{}                `pulumi:"sinkRetryCount"`
	SinkRetryWait                         interface{}                `pulumi:"sinkRetryWait"`
	SqlWriterStoredProcedureName          interface{}                `pulumi:"sqlWriterStoredProcedureName"`
	SqlWriterTableType                    interface{}                `pulumi:"sqlWriterTableType"`
	SqlWriterUseTableLock                 interface{}                `pulumi:"sqlWriterUseTableLock"`
	StoredProcedureParameters             interface{}                `pulumi:"storedProcedureParameters"`
	StoredProcedureTableTypeParameterName interface{}                `pulumi:"storedProcedureTableTypeParameterName"`
	TableOption                           interface{}                `pulumi:"tableOption"`
	Type                                  string                     `pulumi:"type"`
	UpsertSettings                        *SqlUpsertSettingsResponse `pulumi:"upsertSettings"`
	WriteBatchSize                        interface{}                `pulumi:"writeBatchSize"`
	WriteBatchTimeout                     interface{}                `pulumi:"writeBatchTimeout"`
	WriteBehavior                         interface{}                `pulumi:"writeBehavior"`
}

type AzureSqlSource struct {
	AdditionalColumns            interface{}           `pulumi:"additionalColumns"`
	DisableMetricsCollection     interface{}           `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections     interface{}           `pulumi:"maxConcurrentConnections"`
	PartitionOption              interface{}           `pulumi:"partitionOption"`
	PartitionSettings            *SqlPartitionSettings `pulumi:"partitionSettings"`
	ProduceAdditionalTypes       interface{}           `pulumi:"produceAdditionalTypes"`
	QueryTimeout                 interface{}           `pulumi:"queryTimeout"`
	SourceRetryCount             interface{}           `pulumi:"sourceRetryCount"`
	SourceRetryWait              interface{}           `pulumi:"sourceRetryWait"`
	SqlReaderQuery               interface{}           `pulumi:"sqlReaderQuery"`
	SqlReaderStoredProcedureName interface{}           `pulumi:"sqlReaderStoredProcedureName"`
	StoredProcedureParameters    interface{}           `pulumi:"storedProcedureParameters"`
	Type                         string                `pulumi:"type"`
}

type AzureSqlSourceResponse struct {
	AdditionalColumns            interface{}                   `pulumi:"additionalColumns"`
	DisableMetricsCollection     interface{}                   `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections     interface{}                   `pulumi:"maxConcurrentConnections"`
	PartitionOption              interface{}                   `pulumi:"partitionOption"`
	PartitionSettings            *SqlPartitionSettingsResponse `pulumi:"partitionSettings"`
	ProduceAdditionalTypes       interface{}                   `pulumi:"produceAdditionalTypes"`
	QueryTimeout                 interface{}                   `pulumi:"queryTimeout"`
	SourceRetryCount             interface{}                   `pulumi:"sourceRetryCount"`
	SourceRetryWait              interface{}                   `pulumi:"sourceRetryWait"`
	SqlReaderQuery               interface{}                   `pulumi:"sqlReaderQuery"`
	SqlReaderStoredProcedureName interface{}                   `pulumi:"sqlReaderStoredProcedureName"`
	StoredProcedureParameters    interface{}                   `pulumi:"storedProcedureParameters"`
	Type                         string                        `pulumi:"type"`
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
	Annotations         []interface{}                     `pulumi:"annotations"`
	Authentication      interface{}                       `pulumi:"authentication"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	Endpoint            interface{}                       `pulumi:"endpoint"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Type                string                            `pulumi:"type"`
	WorkspaceResourceId interface{}                       `pulumi:"workspaceResourceId"`
}

type AzureSynapseArtifactsLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	Authentication      interface{}                               `pulumi:"authentication"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	Endpoint            interface{}                               `pulumi:"endpoint"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
	WorkspaceResourceId interface{}                               `pulumi:"workspaceResourceId"`
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
	DisablePublish      *bool               `pulumi:"disablePublish"`
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
	DisablePublish      *bool                       `pulumi:"disablePublish"`
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
	DisablePublish      *bool   `pulumi:"disablePublish"`
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
	DisablePublish      *bool   `pulumi:"disablePublish"`
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

type GoogleSheetsLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ApiToken            interface{}                       `pulumi:"apiToken"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Type                string                            `pulumi:"type"`
}

type GoogleSheetsLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ApiToken            interface{}                               `pulumi:"apiToken"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
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
}
