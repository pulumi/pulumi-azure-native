


package v20170901preview

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
	SecureOutput           *bool       `pulumi:"secureOutput"`
	Timeout                interface{} `pulumi:"timeout"`
}

type ActivityPolicyResponse struct {
	Retry                  interface{} `pulumi:"retry"`
	RetryIntervalInSeconds *int        `pulumi:"retryIntervalInSeconds"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type AmazonMWSObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type AmazonMWSSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type AmazonMWSSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
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
	Query                  interface{}             `pulumi:"query"`
	RedshiftUnloadSettings *RedshiftUnloadSettings `pulumi:"redshiftUnloadSettings"`
	SourceRetryCount       interface{}             `pulumi:"sourceRetryCount"`
	SourceRetryWait        interface{}             `pulumi:"sourceRetryWait"`
	Type                   string                  `pulumi:"type"`
}

type AmazonRedshiftSourceResponse struct {
	Query                  interface{}                     `pulumi:"query"`
	RedshiftUnloadSettings *RedshiftUnloadSettingsResponse `pulumi:"redshiftUnloadSettings"`
	SourceRetryCount       interface{}                     `pulumi:"sourceRetryCount"`
	SourceRetryWait        interface{}                     `pulumi:"sourceRetryWait"`
	Type                   string                          `pulumi:"type"`
}

type AmazonS3Dataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	BucketName        interface{}                       `pulumi:"bucketName"`
	Compression       interface{}                       `pulumi:"compression"`
	Description       *string                           `pulumi:"description"`
	Format            interface{}                       `pulumi:"format"`
	Key               interface{}                       `pulumi:"key"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Prefix            interface{}                       `pulumi:"prefix"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
	Version           interface{}                       `pulumi:"version"`
}

type AmazonS3DatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	BucketName        interface{}                               `pulumi:"bucketName"`
	Compression       interface{}                               `pulumi:"compression"`
	Description       *string                                   `pulumi:"description"`
	Format            interface{}                               `pulumi:"format"`
	Key               interface{}                               `pulumi:"key"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Prefix            interface{}                               `pulumi:"prefix"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
	Version           interface{}                               `pulumi:"version"`
}

type AmazonS3LinkedService struct {
	AccessKeyId         interface{}                       `pulumi:"accessKeyId"`
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	SecretAccessKey     interface{}                       `pulumi:"secretAccessKey"`
	Type                string                            `pulumi:"type"`
}

type AmazonS3LinkedServiceResponse struct {
	AccessKeyId         interface{}                               `pulumi:"accessKeyId"`
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	SecretAccessKey     interface{}                               `pulumi:"secretAccessKey"`
	Type                string                                    `pulumi:"type"`
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

type AzureBatchLinkedService struct {
	AccessKey           interface{}                       `pulumi:"accessKey"`
	AccountName         interface{}                       `pulumi:"accountName"`
	Annotations         []interface{}                     `pulumi:"annotations"`
	BatchUri            interface{}                       `pulumi:"batchUri"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
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
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	LinkedServiceName   LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	PoolName            interface{}                               `pulumi:"poolName"`
	Type                string                                    `pulumi:"type"`
}

type AzureBlobDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Compression       interface{}                       `pulumi:"compression"`
	Description       *string                           `pulumi:"description"`
	FileName          interface{}                       `pulumi:"fileName"`
	FolderPath        interface{}                       `pulumi:"folderPath"`
	Format            interface{}                       `pulumi:"format"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	TableRootLocation interface{}                       `pulumi:"tableRootLocation"`
	Type              string                            `pulumi:"type"`
}

type AzureBlobDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Compression       interface{}                               `pulumi:"compression"`
	Description       *string                                   `pulumi:"description"`
	FileName          interface{}                               `pulumi:"fileName"`
	FolderPath        interface{}                               `pulumi:"folderPath"`
	Format            interface{}                               `pulumi:"format"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	TableRootLocation interface{}                               `pulumi:"tableRootLocation"`
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
	Compression       interface{}                       `pulumi:"compression"`
	Description       *string                           `pulumi:"description"`
	FileName          interface{}                       `pulumi:"fileName"`
	FolderPath        interface{}                       `pulumi:"folderPath"`
	Format            interface{}                       `pulumi:"format"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type AzureDataLakeStoreDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Compression       interface{}                               `pulumi:"compression"`
	Description       *string                                   `pulumi:"description"`
	FileName          interface{}                               `pulumi:"fileName"`
	FolderPath        interface{}                               `pulumi:"folderPath"`
	Format            interface{}                               `pulumi:"format"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type AzureDataLakeStoreLinkedService struct {
	AccountName         interface{}                       `pulumi:"accountName"`
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
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
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
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

type AzureDataLakeStoreSink struct {
	CopyBehavior      interface{} `pulumi:"copyBehavior"`
	SinkRetryCount    interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait     interface{} `pulumi:"sinkRetryWait"`
	Type              string      `pulumi:"type"`
	WriteBatchSize    interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout interface{} `pulumi:"writeBatchTimeout"`
}

type AzureDataLakeStoreSinkResponse struct {
	CopyBehavior      interface{} `pulumi:"copyBehavior"`
	SinkRetryCount    interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait     interface{} `pulumi:"sinkRetryWait"`
	Type              string      `pulumi:"type"`
	WriteBatchSize    interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout interface{} `pulumi:"writeBatchTimeout"`
}

type AzureDataLakeStoreSource struct {
	Recursive        interface{} `pulumi:"recursive"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type AzureDataLakeStoreSourceResponse struct {
	Recursive        interface{} `pulumi:"recursive"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type AzureDatabricksLinkedService struct {
	AccessToken           interface{}                       `pulumi:"accessToken"`
	Annotations           []interface{}                     `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	Domain                interface{}                       `pulumi:"domain"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	ExistingClusterId     interface{}                       `pulumi:"existingClusterId"`
	NewClusterNodeType    interface{}                       `pulumi:"newClusterNodeType"`
	NewClusterNumOfWorker interface{}                       `pulumi:"newClusterNumOfWorker"`
	NewClusterSparkConf   map[string]interface{}            `pulumi:"newClusterSparkConf"`
	NewClusterVersion     interface{}                       `pulumi:"newClusterVersion"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Type                  string                            `pulumi:"type"`
}

type AzureDatabricksLinkedServiceResponse struct {
	AccessToken           interface{}                               `pulumi:"accessToken"`
	Annotations           []interface{}                             `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	Domain                interface{}                               `pulumi:"domain"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	ExistingClusterId     interface{}                               `pulumi:"existingClusterId"`
	NewClusterNodeType    interface{}                               `pulumi:"newClusterNodeType"`
	NewClusterNumOfWorker interface{}                               `pulumi:"newClusterNumOfWorker"`
	NewClusterSparkConf   map[string]interface{}                    `pulumi:"newClusterSparkConf"`
	NewClusterVersion     interface{}                               `pulumi:"newClusterVersion"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                  string                                    `pulumi:"type"`
}

type AzureKeyVaultLinkedService struct {
	Annotations []interface{}                     `pulumi:"annotations"`
	BaseUrl     interface{}                       `pulumi:"baseUrl"`
	ConnectVia  *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description *string                           `pulumi:"description"`
	Parameters  map[string]ParameterSpecification `pulumi:"parameters"`
	Type        string                            `pulumi:"type"`
}

type AzureKeyVaultLinkedServiceResponse struct {
	Annotations []interface{}                             `pulumi:"annotations"`
	BaseUrl     interface{}                               `pulumi:"baseUrl"`
	ConnectVia  *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
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
	WebServiceInputs  map[string]AzureMLWebServiceFileResponse `pulumi:"webServiceInputs"`
	WebServiceOutputs map[string]AzureMLWebServiceFileResponse `pulumi:"webServiceOutputs"`
}

type AzureMLLinkedService struct {
	Annotations            []interface{}                     `pulumi:"annotations"`
	ApiKey                 interface{}                       `pulumi:"apiKey"`
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
}

type AzureMLWebServiceFile struct {
	FilePath          interface{}            `pulumi:"filePath"`
	LinkedServiceName LinkedServiceReference `pulumi:"linkedServiceName"`
}

type AzureMLWebServiceFileResponse struct {
	FilePath          interface{}                    `pulumi:"filePath"`
	LinkedServiceName LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
}

type AzureMySqlLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Type                string                            `pulumi:"type"`
}

type AzureMySqlLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
}

type AzureMySqlSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type AzureMySqlSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type AzureMySqlTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type AzureMySqlTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
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
	Type                string                            `pulumi:"type"`
}

type AzurePostgreSqlLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
}

type AzurePostgreSqlSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type AzurePostgreSqlSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type AzurePostgreSqlTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type AzurePostgreSqlTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type AzureQueueSink struct {
	SinkRetryCount    interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait     interface{} `pulumi:"sinkRetryWait"`
	Type              string      `pulumi:"type"`
	WriteBatchSize    interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout interface{} `pulumi:"writeBatchTimeout"`
}

type AzureQueueSinkResponse struct {
	SinkRetryCount    interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait     interface{} `pulumi:"sinkRetryWait"`
	Type              string      `pulumi:"type"`
	WriteBatchSize    interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout interface{} `pulumi:"writeBatchTimeout"`
}

type AzureSearchIndexDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	IndexName         interface{}                       `pulumi:"indexName"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type AzureSearchIndexDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	IndexName         interface{}                               `pulumi:"indexName"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type AzureSearchIndexSink struct {
	SinkRetryCount    interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait     interface{} `pulumi:"sinkRetryWait"`
	Type              string      `pulumi:"type"`
	WriteBatchSize    interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior     *string     `pulumi:"writeBehavior"`
}

type AzureSearchIndexSinkResponse struct {
	SinkRetryCount    interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait     interface{} `pulumi:"sinkRetryWait"`
	Type              string      `pulumi:"type"`
	WriteBatchSize    interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior     *string     `pulumi:"writeBehavior"`
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
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	ServicePrincipalId  interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey interface{}                       `pulumi:"servicePrincipalKey"`
	Tenant              interface{}                       `pulumi:"tenant"`
	Type                string                            `pulumi:"type"`
}

type AzureSqlDWLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	ServicePrincipalId  interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey interface{}                               `pulumi:"servicePrincipalKey"`
	Tenant              interface{}                               `pulumi:"tenant"`
	Type                string                                    `pulumi:"type"`
}

type AzureSqlDWTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type AzureSqlDWTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type AzureSqlDatabaseLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	ServicePrincipalId  interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey interface{}                       `pulumi:"servicePrincipalKey"`
	Tenant              interface{}                       `pulumi:"tenant"`
	Type                string                            `pulumi:"type"`
}

type AzureSqlDatabaseLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	ServicePrincipalId  interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey interface{}                               `pulumi:"servicePrincipalKey"`
	Tenant              interface{}                               `pulumi:"tenant"`
	Type                string                                    `pulumi:"type"`
}

type AzureSqlTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type AzureSqlTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type AzureStorageLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	SasUri              interface{}                       `pulumi:"sasUri"`
	Type                string                            `pulumi:"type"`
}

type AzureStorageLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	SasUri              interface{}                               `pulumi:"sasUri"`
	Type                string                                    `pulumi:"type"`
}

type AzureTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type AzureTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type AzureTableSink struct {
	AzureTableDefaultPartitionKeyValue interface{} `pulumi:"azureTableDefaultPartitionKeyValue"`
	AzureTableInsertType               interface{} `pulumi:"azureTableInsertType"`
	AzureTablePartitionKeyName         interface{} `pulumi:"azureTablePartitionKeyName"`
	AzureTableRowKeyName               interface{} `pulumi:"azureTableRowKeyName"`
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
	SinkRetryCount                     interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait                      interface{} `pulumi:"sinkRetryWait"`
	Type                               string      `pulumi:"type"`
	WriteBatchSize                     interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout                  interface{} `pulumi:"writeBatchTimeout"`
}

type AzureTableSource struct {
	AzureTableSourceIgnoreTableNotFound interface{} `pulumi:"azureTableSourceIgnoreTableNotFound"`
	AzureTableSourceQuery               interface{} `pulumi:"azureTableSourceQuery"`
	SourceRetryCount                    interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait                     interface{} `pulumi:"sourceRetryWait"`
	Type                                string      `pulumi:"type"`
}

type AzureTableSourceResponse struct {
	AzureTableSourceIgnoreTableNotFound interface{} `pulumi:"azureTableSourceIgnoreTableNotFound"`
	AzureTableSourceQuery               interface{} `pulumi:"azureTableSourceQuery"`
	SourceRetryCount                    interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait                     interface{} `pulumi:"sourceRetryWait"`
	Type                                string      `pulumi:"type"`
}

type BlobEventsTrigger struct {
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
	BlobWriterAddHeader      interface{} `pulumi:"blobWriterAddHeader"`
	BlobWriterDateTimeFormat interface{} `pulumi:"blobWriterDateTimeFormat"`
	BlobWriterOverwriteFiles interface{} `pulumi:"blobWriterOverwriteFiles"`
	CopyBehavior             interface{} `pulumi:"copyBehavior"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type BlobSinkResponse struct {
	BlobWriterAddHeader      interface{} `pulumi:"blobWriterAddHeader"`
	BlobWriterDateTimeFormat interface{} `pulumi:"blobWriterDateTimeFormat"`
	BlobWriterOverwriteFiles interface{} `pulumi:"blobWriterOverwriteFiles"`
	CopyBehavior             interface{} `pulumi:"copyBehavior"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type BlobSource struct {
	Recursive           interface{} `pulumi:"recursive"`
	SkipHeaderLineCount interface{} `pulumi:"skipHeaderLineCount"`
	SourceRetryCount    interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait     interface{} `pulumi:"sourceRetryWait"`
	TreatEmptyAsNull    interface{} `pulumi:"treatEmptyAsNull"`
	Type                string      `pulumi:"type"`
}

type BlobSourceResponse struct {
	Recursive           interface{} `pulumi:"recursive"`
	SkipHeaderLineCount interface{} `pulumi:"skipHeaderLineCount"`
	SourceRetryCount    interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait     interface{} `pulumi:"sourceRetryWait"`
	TreatEmptyAsNull    interface{} `pulumi:"treatEmptyAsNull"`
	Type                string      `pulumi:"type"`
}

type BlobTrigger struct {
	Description    *string                    `pulumi:"description"`
	FolderPath     string                     `pulumi:"folderPath"`
	LinkedService  LinkedServiceReference     `pulumi:"linkedService"`
	MaxConcurrency int                        `pulumi:"maxConcurrency"`
	Pipelines      []TriggerPipelineReference `pulumi:"pipelines"`
	Type           string                     `pulumi:"type"`
}

type BlobTriggerResponse struct {
	Description    *string                            `pulumi:"description"`
	FolderPath     string                             `pulumi:"folderPath"`
	LinkedService  LinkedServiceReferenceResponse     `pulumi:"linkedService"`
	MaxConcurrency int                                `pulumi:"maxConcurrency"`
	Pipelines      []TriggerPipelineReferenceResponse `pulumi:"pipelines"`
	RuntimeState   string                             `pulumi:"runtimeState"`
	Type           string                             `pulumi:"type"`
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
	ConsistencyLevel *string     `pulumi:"consistencyLevel"`
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type CassandraSourceResponse struct {
	ConsistencyLevel *string     `pulumi:"consistencyLevel"`
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type CassandraTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Keyspace          interface{}                       `pulumi:"keyspace"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type CassandraTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Keyspace          interface{}                               `pulumi:"keyspace"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type ConcurLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	ClientId              interface{}                       `pulumi:"clientId"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type ConcurObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type ConcurSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type ConcurSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type ControlActivity struct {
	DependsOn   []ActivityDependency `pulumi:"dependsOn"`
	Description *string              `pulumi:"description"`
	Name        string               `pulumi:"name"`
	Type        string               `pulumi:"type"`
}

type ControlActivityResponse struct {
	DependsOn   []ActivityDependencyResponse `pulumi:"dependsOn"`
	Description *string                      `pulumi:"description"`
	Name        string                       `pulumi:"name"`
	Type        string                       `pulumi:"type"`
}

type CopyActivity struct {
	CloudDataMovementUnits          interface{}                      `pulumi:"cloudDataMovementUnits"`
	DependsOn                       []ActivityDependency             `pulumi:"dependsOn"`
	Description                     *string                          `pulumi:"description"`
	EnableSkipIncompatibleRow       interface{}                      `pulumi:"enableSkipIncompatibleRow"`
	EnableStaging                   interface{}                      `pulumi:"enableStaging"`
	Inputs                          []DatasetReference               `pulumi:"inputs"`
	LinkedServiceName               *LinkedServiceReference          `pulumi:"linkedServiceName"`
	Name                            string                           `pulumi:"name"`
	Outputs                         []DatasetReference               `pulumi:"outputs"`
	ParallelCopies                  interface{}                      `pulumi:"parallelCopies"`
	Policy                          *ActivityPolicy                  `pulumi:"policy"`
	RedirectIncompatibleRowSettings *RedirectIncompatibleRowSettings `pulumi:"redirectIncompatibleRowSettings"`
	Sink                            interface{}                      `pulumi:"sink"`
	Source                          interface{}                      `pulumi:"source"`
	StagingSettings                 *StagingSettings                 `pulumi:"stagingSettings"`
	Translator                      interface{}                      `pulumi:"translator"`
	Type                            string                           `pulumi:"type"`
}

type CopyActivityResponse struct {
	CloudDataMovementUnits          interface{}                              `pulumi:"cloudDataMovementUnits"`
	DependsOn                       []ActivityDependencyResponse             `pulumi:"dependsOn"`
	Description                     *string                                  `pulumi:"description"`
	EnableSkipIncompatibleRow       interface{}                              `pulumi:"enableSkipIncompatibleRow"`
	EnableStaging                   interface{}                              `pulumi:"enableStaging"`
	Inputs                          []DatasetReferenceResponse               `pulumi:"inputs"`
	LinkedServiceName               *LinkedServiceReferenceResponse          `pulumi:"linkedServiceName"`
	Name                            string                                   `pulumi:"name"`
	Outputs                         []DatasetReferenceResponse               `pulumi:"outputs"`
	ParallelCopies                  interface{}                              `pulumi:"parallelCopies"`
	Policy                          *ActivityPolicyResponse                  `pulumi:"policy"`
	RedirectIncompatibleRowSettings *RedirectIncompatibleRowSettingsResponse `pulumi:"redirectIncompatibleRowSettings"`
	Sink                            interface{}                              `pulumi:"sink"`
	Source                          interface{}                              `pulumi:"source"`
	StagingSettings                 *StagingSettingsResponse                 `pulumi:"stagingSettings"`
	Translator                      interface{}                              `pulumi:"translator"`
	Type                            string                                   `pulumi:"type"`
}

type CosmosDbLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Type                string                            `pulumi:"type"`
}

type CosmosDbLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
}

type CouchbaseLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Type                string                            `pulumi:"type"`
}

type CouchbaseLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
}

type CouchbaseSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type CouchbaseSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type CouchbaseTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type CouchbaseTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type CustomActivity struct {
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
	Type                  string                         `pulumi:"type"`
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
	Type                  string                                 `pulumi:"type"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type CustomDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
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
}

type DatabricksNotebookActivity struct {
	BaseParameters    map[string]interface{}  `pulumi:"baseParameters"`
	DependsOn         []ActivityDependency    `pulumi:"dependsOn"`
	Description       *string                 `pulumi:"description"`
	LinkedServiceName *LinkedServiceReference `pulumi:"linkedServiceName"`
	Name              string                  `pulumi:"name"`
	NotebookPath      interface{}             `pulumi:"notebookPath"`
	Policy            *ActivityPolicy         `pulumi:"policy"`
	Type              string                  `pulumi:"type"`
}

type DatabricksNotebookActivityResponse struct {
	BaseParameters    map[string]interface{}          `pulumi:"baseParameters"`
	DependsOn         []ActivityDependencyResponse    `pulumi:"dependsOn"`
	Description       *string                         `pulumi:"description"`
	LinkedServiceName *LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	Name              string                          `pulumi:"name"`
	NotebookPath      interface{}                     `pulumi:"notebookPath"`
	Policy            *ActivityPolicyResponse         `pulumi:"policy"`
	Type              string                          `pulumi:"type"`
}

type DatasetBZip2Compression struct {
	Type string `pulumi:"type"`
}

type DatasetBZip2CompressionResponse struct {
	Type string `pulumi:"type"`
}

type DatasetDeflateCompression struct {
	Level *string `pulumi:"level"`
	Type  string  `pulumi:"type"`
}

type DatasetDeflateCompressionResponse struct {
	Level *string `pulumi:"level"`
	Type  string  `pulumi:"type"`
}

type DatasetGZipCompression struct {
	Level *string `pulumi:"level"`
	Type  string  `pulumi:"type"`
}

type DatasetGZipCompressionResponse struct {
	Level *string `pulumi:"level"`
	Type  string  `pulumi:"type"`
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

type DatasetZipDeflateCompression struct {
	Level *string `pulumi:"level"`
	Type  string  `pulumi:"type"`
}

type DatasetZipDeflateCompressionResponse struct {
	Level *string `pulumi:"level"`
	Type  string  `pulumi:"type"`
}

type Db2LinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	AuthenticationType  *string                           `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Database            interface{}                       `pulumi:"database"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Server              interface{}                       `pulumi:"server"`
	Type                string                            `pulumi:"type"`
	Username            interface{}                       `pulumi:"username"`
}

type Db2LinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	AuthenticationType  *string                                   `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Database            interface{}                               `pulumi:"database"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Server              interface{}                               `pulumi:"server"`
	Type                string                                    `pulumi:"type"`
	Username            interface{}                               `pulumi:"username"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type DocumentDbCollectionDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	CollectionName    interface{}                               `pulumi:"collectionName"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type DocumentDbCollectionSink struct {
	NestingSeparator  interface{} `pulumi:"nestingSeparator"`
	SinkRetryCount    interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait     interface{} `pulumi:"sinkRetryWait"`
	Type              string      `pulumi:"type"`
	WriteBatchSize    interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout interface{} `pulumi:"writeBatchTimeout"`
}

type DocumentDbCollectionSinkResponse struct {
	NestingSeparator  interface{} `pulumi:"nestingSeparator"`
	SinkRetryCount    interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait     interface{} `pulumi:"sinkRetryWait"`
	Type              string      `pulumi:"type"`
	WriteBatchSize    interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout interface{} `pulumi:"writeBatchTimeout"`
}

type DocumentDbCollectionSource struct {
	NestingSeparator interface{} `pulumi:"nestingSeparator"`
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type DocumentDbCollectionSourceResponse struct {
	NestingSeparator interface{} `pulumi:"nestingSeparator"`
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type DrillLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Type                string                            `pulumi:"type"`
}

type DrillLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
}

type DrillSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type DrillSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type DrillTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type DrillTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type DynamicsEntityDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	EntityName        interface{}                       `pulumi:"entityName"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type DynamicsEntityDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	EntityName        interface{}                               `pulumi:"entityName"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type DynamicsLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	AuthenticationType  string                            `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	DeploymentType      string                            `pulumi:"deploymentType"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	HostName            interface{}                       `pulumi:"hostName"`
	OrganizationName    interface{}                       `pulumi:"organizationName"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Port                interface{}                       `pulumi:"port"`
	ServiceUri          interface{}                       `pulumi:"serviceUri"`
	Type                string                            `pulumi:"type"`
	Username            interface{}                       `pulumi:"username"`
}

type DynamicsLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	AuthenticationType  string                                    `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	DeploymentType      string                                    `pulumi:"deploymentType"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	HostName            interface{}                               `pulumi:"hostName"`
	OrganizationName    interface{}                               `pulumi:"organizationName"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Port                interface{}                               `pulumi:"port"`
	ServiceUri          interface{}                               `pulumi:"serviceUri"`
	Type                string                                    `pulumi:"type"`
	Username            interface{}                               `pulumi:"username"`
}

type DynamicsSink struct {
	IgnoreNullValues  interface{} `pulumi:"ignoreNullValues"`
	SinkRetryCount    interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait     interface{} `pulumi:"sinkRetryWait"`
	Type              string      `pulumi:"type"`
	WriteBatchSize    interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior     string      `pulumi:"writeBehavior"`
}

type DynamicsSinkResponse struct {
	IgnoreNullValues  interface{} `pulumi:"ignoreNullValues"`
	SinkRetryCount    interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait     interface{} `pulumi:"sinkRetryWait"`
	Type              string      `pulumi:"type"`
	WriteBatchSize    interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior     string      `pulumi:"writeBehavior"`
}

type DynamicsSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type DynamicsSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type EloquaObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type EloquaSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type EloquaSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type EntityReference struct {
	ReferenceName *string `pulumi:"referenceName"`
	Type          *string `pulumi:"type"`
}

type EntityReferenceResponse struct {
	ReferenceName *string `pulumi:"referenceName"`
	Type          *string `pulumi:"type"`
}

type ExecutePipelineActivity struct {
	DependsOn        []ActivityDependency   `pulumi:"dependsOn"`
	Description      *string                `pulumi:"description"`
	Name             string                 `pulumi:"name"`
	Parameters       map[string]interface{} `pulumi:"parameters"`
	Pipeline         PipelineReference      `pulumi:"pipeline"`
	Type             string                 `pulumi:"type"`
	WaitOnCompletion *bool                  `pulumi:"waitOnCompletion"`
}

type ExecutePipelineActivityResponse struct {
	DependsOn        []ActivityDependencyResponse `pulumi:"dependsOn"`
	Description      *string                      `pulumi:"description"`
	Name             string                       `pulumi:"name"`
	Parameters       map[string]interface{}       `pulumi:"parameters"`
	Pipeline         PipelineReferenceResponse    `pulumi:"pipeline"`
	Type             string                       `pulumi:"type"`
	WaitOnCompletion *bool                        `pulumi:"waitOnCompletion"`
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
}

type ExecutionActivity struct {
	DependsOn         []ActivityDependency    `pulumi:"dependsOn"`
	Description       *string                 `pulumi:"description"`
	LinkedServiceName *LinkedServiceReference `pulumi:"linkedServiceName"`
	Name              string                  `pulumi:"name"`
	Policy            *ActivityPolicy         `pulumi:"policy"`
	Type              string                  `pulumi:"type"`
}

type ExecutionActivityResponse struct {
	DependsOn         []ActivityDependencyResponse    `pulumi:"dependsOn"`
	Description       *string                         `pulumi:"description"`
	LinkedServiceName *LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	Name              string                          `pulumi:"name"`
	Policy            *ActivityPolicyResponse         `pulumi:"policy"`
	Type              string                          `pulumi:"type"`
}

type Expression struct {
	Type  string `pulumi:"type"`
	Value string `pulumi:"value"`
}

type ExpressionResponse struct {
	Type  string `pulumi:"type"`
	Value string `pulumi:"value"`
}

type FactoryIdentity struct {
	Type FactoryIdentityType `pulumi:"type"`
}





type FactoryIdentityInput interface {
	pulumi.Input

	ToFactoryIdentityOutput() FactoryIdentityOutput
	ToFactoryIdentityOutputWithContext(context.Context) FactoryIdentityOutput
}

type FactoryIdentityArgs struct {
	Type FactoryIdentityTypeInput `pulumi:"type"`
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

func (o FactoryIdentityOutput) Type() FactoryIdentityTypeOutput {
	return o.ApplyT(func(v FactoryIdentity) FactoryIdentityType { return v.Type }).(FactoryIdentityTypeOutput)
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

func (o FactoryIdentityPtrOutput) Type() FactoryIdentityTypePtrOutput {
	return o.ApplyT(func(v *FactoryIdentity) *FactoryIdentityType {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(FactoryIdentityTypePtrOutput)
}

type FactoryIdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
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

type FactoryVSTSConfiguration struct {
	AccountName         *string `pulumi:"accountName"`
	CollaborationBranch *string `pulumi:"collaborationBranch"`
	LastCommitId        *string `pulumi:"lastCommitId"`
	ProjectName         *string `pulumi:"projectName"`
	RepositoryName      *string `pulumi:"repositoryName"`
	RootFolder          *string `pulumi:"rootFolder"`
	TenantId            *string `pulumi:"tenantId"`
}





type FactoryVSTSConfigurationInput interface {
	pulumi.Input

	ToFactoryVSTSConfigurationOutput() FactoryVSTSConfigurationOutput
	ToFactoryVSTSConfigurationOutputWithContext(context.Context) FactoryVSTSConfigurationOutput
}

type FactoryVSTSConfigurationArgs struct {
	AccountName         pulumi.StringPtrInput `pulumi:"accountName"`
	CollaborationBranch pulumi.StringPtrInput `pulumi:"collaborationBranch"`
	LastCommitId        pulumi.StringPtrInput `pulumi:"lastCommitId"`
	ProjectName         pulumi.StringPtrInput `pulumi:"projectName"`
	RepositoryName      pulumi.StringPtrInput `pulumi:"repositoryName"`
	RootFolder          pulumi.StringPtrInput `pulumi:"rootFolder"`
	TenantId            pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (FactoryVSTSConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FactoryVSTSConfiguration)(nil)).Elem()
}

func (i FactoryVSTSConfigurationArgs) ToFactoryVSTSConfigurationOutput() FactoryVSTSConfigurationOutput {
	return i.ToFactoryVSTSConfigurationOutputWithContext(context.Background())
}

func (i FactoryVSTSConfigurationArgs) ToFactoryVSTSConfigurationOutputWithContext(ctx context.Context) FactoryVSTSConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FactoryVSTSConfigurationOutput)
}

func (i FactoryVSTSConfigurationArgs) ToFactoryVSTSConfigurationPtrOutput() FactoryVSTSConfigurationPtrOutput {
	return i.ToFactoryVSTSConfigurationPtrOutputWithContext(context.Background())
}

func (i FactoryVSTSConfigurationArgs) ToFactoryVSTSConfigurationPtrOutputWithContext(ctx context.Context) FactoryVSTSConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FactoryVSTSConfigurationOutput).ToFactoryVSTSConfigurationPtrOutputWithContext(ctx)
}









type FactoryVSTSConfigurationPtrInput interface {
	pulumi.Input

	ToFactoryVSTSConfigurationPtrOutput() FactoryVSTSConfigurationPtrOutput
	ToFactoryVSTSConfigurationPtrOutputWithContext(context.Context) FactoryVSTSConfigurationPtrOutput
}

type factoryVSTSConfigurationPtrType FactoryVSTSConfigurationArgs

func FactoryVSTSConfigurationPtr(v *FactoryVSTSConfigurationArgs) FactoryVSTSConfigurationPtrInput {
	return (*factoryVSTSConfigurationPtrType)(v)
}

func (*factoryVSTSConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FactoryVSTSConfiguration)(nil)).Elem()
}

func (i *factoryVSTSConfigurationPtrType) ToFactoryVSTSConfigurationPtrOutput() FactoryVSTSConfigurationPtrOutput {
	return i.ToFactoryVSTSConfigurationPtrOutputWithContext(context.Background())
}

func (i *factoryVSTSConfigurationPtrType) ToFactoryVSTSConfigurationPtrOutputWithContext(ctx context.Context) FactoryVSTSConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FactoryVSTSConfigurationPtrOutput)
}

type FactoryVSTSConfigurationOutput struct{ *pulumi.OutputState }

func (FactoryVSTSConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FactoryVSTSConfiguration)(nil)).Elem()
}

func (o FactoryVSTSConfigurationOutput) ToFactoryVSTSConfigurationOutput() FactoryVSTSConfigurationOutput {
	return o
}

func (o FactoryVSTSConfigurationOutput) ToFactoryVSTSConfigurationOutputWithContext(ctx context.Context) FactoryVSTSConfigurationOutput {
	return o
}

func (o FactoryVSTSConfigurationOutput) ToFactoryVSTSConfigurationPtrOutput() FactoryVSTSConfigurationPtrOutput {
	return o.ToFactoryVSTSConfigurationPtrOutputWithContext(context.Background())
}

func (o FactoryVSTSConfigurationOutput) ToFactoryVSTSConfigurationPtrOutputWithContext(ctx context.Context) FactoryVSTSConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FactoryVSTSConfiguration) *FactoryVSTSConfiguration {
		return &v
	}).(FactoryVSTSConfigurationPtrOutput)
}

func (o FactoryVSTSConfigurationOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FactoryVSTSConfiguration) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationOutput) CollaborationBranch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FactoryVSTSConfiguration) *string { return v.CollaborationBranch }).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationOutput) LastCommitId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FactoryVSTSConfiguration) *string { return v.LastCommitId }).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FactoryVSTSConfiguration) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FactoryVSTSConfiguration) *string { return v.RepositoryName }).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationOutput) RootFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FactoryVSTSConfiguration) *string { return v.RootFolder }).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FactoryVSTSConfiguration) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type FactoryVSTSConfigurationPtrOutput struct{ *pulumi.OutputState }

func (FactoryVSTSConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FactoryVSTSConfiguration)(nil)).Elem()
}

func (o FactoryVSTSConfigurationPtrOutput) ToFactoryVSTSConfigurationPtrOutput() FactoryVSTSConfigurationPtrOutput {
	return o
}

func (o FactoryVSTSConfigurationPtrOutput) ToFactoryVSTSConfigurationPtrOutputWithContext(ctx context.Context) FactoryVSTSConfigurationPtrOutput {
	return o
}

func (o FactoryVSTSConfigurationPtrOutput) Elem() FactoryVSTSConfigurationOutput {
	return o.ApplyT(func(v *FactoryVSTSConfiguration) FactoryVSTSConfiguration {
		if v != nil {
			return *v
		}
		var ret FactoryVSTSConfiguration
		return ret
	}).(FactoryVSTSConfigurationOutput)
}

func (o FactoryVSTSConfigurationPtrOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FactoryVSTSConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AccountName
	}).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationPtrOutput) CollaborationBranch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FactoryVSTSConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.CollaborationBranch
	}).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationPtrOutput) LastCommitId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FactoryVSTSConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.LastCommitId
	}).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationPtrOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FactoryVSTSConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.ProjectName
	}).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationPtrOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FactoryVSTSConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RepositoryName
	}).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationPtrOutput) RootFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FactoryVSTSConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RootFolder
	}).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FactoryVSTSConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type FactoryVSTSConfigurationResponse struct {
	AccountName         *string `pulumi:"accountName"`
	CollaborationBranch *string `pulumi:"collaborationBranch"`
	LastCommitId        *string `pulumi:"lastCommitId"`
	ProjectName         *string `pulumi:"projectName"`
	RepositoryName      *string `pulumi:"repositoryName"`
	RootFolder          *string `pulumi:"rootFolder"`
	TenantId            *string `pulumi:"tenantId"`
}

type FactoryVSTSConfigurationResponseOutput struct{ *pulumi.OutputState }

func (FactoryVSTSConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FactoryVSTSConfigurationResponse)(nil)).Elem()
}

func (o FactoryVSTSConfigurationResponseOutput) ToFactoryVSTSConfigurationResponseOutput() FactoryVSTSConfigurationResponseOutput {
	return o
}

func (o FactoryVSTSConfigurationResponseOutput) ToFactoryVSTSConfigurationResponseOutputWithContext(ctx context.Context) FactoryVSTSConfigurationResponseOutput {
	return o
}

func (o FactoryVSTSConfigurationResponseOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FactoryVSTSConfigurationResponse) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationResponseOutput) CollaborationBranch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FactoryVSTSConfigurationResponse) *string { return v.CollaborationBranch }).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationResponseOutput) LastCommitId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FactoryVSTSConfigurationResponse) *string { return v.LastCommitId }).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationResponseOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FactoryVSTSConfigurationResponse) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationResponseOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FactoryVSTSConfigurationResponse) *string { return v.RepositoryName }).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationResponseOutput) RootFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FactoryVSTSConfigurationResponse) *string { return v.RootFolder }).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FactoryVSTSConfigurationResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type FactoryVSTSConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (FactoryVSTSConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FactoryVSTSConfigurationResponse)(nil)).Elem()
}

func (o FactoryVSTSConfigurationResponsePtrOutput) ToFactoryVSTSConfigurationResponsePtrOutput() FactoryVSTSConfigurationResponsePtrOutput {
	return o
}

func (o FactoryVSTSConfigurationResponsePtrOutput) ToFactoryVSTSConfigurationResponsePtrOutputWithContext(ctx context.Context) FactoryVSTSConfigurationResponsePtrOutput {
	return o
}

func (o FactoryVSTSConfigurationResponsePtrOutput) Elem() FactoryVSTSConfigurationResponseOutput {
	return o.ApplyT(func(v *FactoryVSTSConfigurationResponse) FactoryVSTSConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret FactoryVSTSConfigurationResponse
		return ret
	}).(FactoryVSTSConfigurationResponseOutput)
}

func (o FactoryVSTSConfigurationResponsePtrOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FactoryVSTSConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccountName
	}).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationResponsePtrOutput) CollaborationBranch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FactoryVSTSConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.CollaborationBranch
	}).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationResponsePtrOutput) LastCommitId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FactoryVSTSConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastCommitId
	}).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationResponsePtrOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FactoryVSTSConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProjectName
	}).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationResponsePtrOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FactoryVSTSConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RepositoryName
	}).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationResponsePtrOutput) RootFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FactoryVSTSConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RootFolder
	}).(pulumi.StringPtrOutput)
}

func (o FactoryVSTSConfigurationResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FactoryVSTSConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
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

type FileShareDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Compression       interface{}                       `pulumi:"compression"`
	Description       *string                           `pulumi:"description"`
	FileFilter        interface{}                       `pulumi:"fileFilter"`
	FileName          interface{}                       `pulumi:"fileName"`
	FolderPath        interface{}                       `pulumi:"folderPath"`
	Format            interface{}                       `pulumi:"format"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type FileShareDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Compression       interface{}                               `pulumi:"compression"`
	Description       *string                                   `pulumi:"description"`
	FileFilter        interface{}                               `pulumi:"fileFilter"`
	FileName          interface{}                               `pulumi:"fileName"`
	FolderPath        interface{}                               `pulumi:"folderPath"`
	Format            interface{}                               `pulumi:"format"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type FileSystemSink struct {
	CopyBehavior      interface{} `pulumi:"copyBehavior"`
	SinkRetryCount    interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait     interface{} `pulumi:"sinkRetryWait"`
	Type              string      `pulumi:"type"`
	WriteBatchSize    interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout interface{} `pulumi:"writeBatchTimeout"`
}

type FileSystemSinkResponse struct {
	CopyBehavior      interface{} `pulumi:"copyBehavior"`
	SinkRetryCount    interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait     interface{} `pulumi:"sinkRetryWait"`
	Type              string      `pulumi:"type"`
	WriteBatchSize    interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout interface{} `pulumi:"writeBatchTimeout"`
}

type FileSystemSource struct {
	Recursive        interface{} `pulumi:"recursive"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type FileSystemSourceResponse struct {
	Recursive        interface{} `pulumi:"recursive"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type FilterActivity struct {
	Condition   Expression           `pulumi:"condition"`
	DependsOn   []ActivityDependency `pulumi:"dependsOn"`
	Description *string              `pulumi:"description"`
	Items       Expression           `pulumi:"items"`
	Name        string               `pulumi:"name"`
	Type        string               `pulumi:"type"`
}

type FilterActivityResponse struct {
	Condition   ExpressionResponse           `pulumi:"condition"`
	DependsOn   []ActivityDependencyResponse `pulumi:"dependsOn"`
	Description *string                      `pulumi:"description"`
	Items       ExpressionResponse           `pulumi:"items"`
	Name        string                       `pulumi:"name"`
	Type        string                       `pulumi:"type"`
}

type ForEachActivity struct {
	Activities   []interface{}        `pulumi:"activities"`
	BatchCount   *int                 `pulumi:"batchCount"`
	DependsOn    []ActivityDependency `pulumi:"dependsOn"`
	Description  *string              `pulumi:"description"`
	IsSequential *bool                `pulumi:"isSequential"`
	Items        Expression           `pulumi:"items"`
	Name         string               `pulumi:"name"`
	Type         string               `pulumi:"type"`
}

type ForEachActivityResponse struct {
	Activities   []interface{}                `pulumi:"activities"`
	BatchCount   *int                         `pulumi:"batchCount"`
	DependsOn    []ActivityDependencyResponse `pulumi:"dependsOn"`
	Description  *string                      `pulumi:"description"`
	IsSequential *bool                        `pulumi:"isSequential"`
	Items        ExpressionResponse           `pulumi:"items"`
	Name         string                       `pulumi:"name"`
	Type         string                       `pulumi:"type"`
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

type GetMetadataActivity struct {
	Dataset           DatasetReference        `pulumi:"dataset"`
	DependsOn         []ActivityDependency    `pulumi:"dependsOn"`
	Description       *string                 `pulumi:"description"`
	FieldList         []interface{}           `pulumi:"fieldList"`
	LinkedServiceName *LinkedServiceReference `pulumi:"linkedServiceName"`
	Name              string                  `pulumi:"name"`
	Policy            *ActivityPolicy         `pulumi:"policy"`
	Type              string                  `pulumi:"type"`
}

type GetMetadataActivityResponse struct {
	Dataset           DatasetReferenceResponse        `pulumi:"dataset"`
	DependsOn         []ActivityDependencyResponse    `pulumi:"dependsOn"`
	Description       *string                         `pulumi:"description"`
	FieldList         []interface{}                   `pulumi:"fieldList"`
	LinkedServiceName *LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	Name              string                          `pulumi:"name"`
	Policy            *ActivityPolicyResponse         `pulumi:"policy"`
	Type              string                          `pulumi:"type"`
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
	Description       *string                           `pulumi:"description"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type GoogleBigQueryObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type GoogleBigQuerySource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type GoogleBigQuerySourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type GreenplumLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Type                string                            `pulumi:"type"`
}

type GreenplumLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
}

type GreenplumSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type GreenplumSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type GreenplumTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type GreenplumTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type HBaseObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type HBaseSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type HBaseSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
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
	ScriptLinkedService   *LinkedServiceReference  `pulumi:"scriptLinkedService"`
	ScriptPath            interface{}              `pulumi:"scriptPath"`
	StorageLinkedServices []LinkedServiceReference `pulumi:"storageLinkedServices"`
	Type                  string                   `pulumi:"type"`
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
	ScriptLinkedService   *LinkedServiceReferenceResponse  `pulumi:"scriptLinkedService"`
	ScriptPath            interface{}                      `pulumi:"scriptPath"`
	StorageLinkedServices []LinkedServiceReferenceResponse `pulumi:"storageLinkedServices"`
	Type                  string                           `pulumi:"type"`
}

type HDInsightLinkedService struct {
	Annotations               []interface{}                     `pulumi:"annotations"`
	ClusterUri                interface{}                       `pulumi:"clusterUri"`
	ConnectVia                *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description               *string                           `pulumi:"description"`
	EncryptedCredential       interface{}                       `pulumi:"encryptedCredential"`
	HcatalogLinkedServiceName *LinkedServiceReference           `pulumi:"hcatalogLinkedServiceName"`
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
	HcatalogLinkedServiceName *LinkedServiceReferenceResponse           `pulumi:"hcatalogLinkedServiceName"`
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
	ServicePrincipalId           interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey          interface{}                       `pulumi:"servicePrincipalKey"`
	SparkVersion                 interface{}                       `pulumi:"sparkVersion"`
	StormConfiguration           interface{}                       `pulumi:"stormConfiguration"`
	Tenant                       interface{}                       `pulumi:"tenant"`
	TimeToLive                   interface{}                       `pulumi:"timeToLive"`
	Type                         string                            `pulumi:"type"`
	Version                      interface{}                       `pulumi:"version"`
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
	ServicePrincipalId           interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey          interface{}                               `pulumi:"servicePrincipalKey"`
	SparkVersion                 interface{}                               `pulumi:"sparkVersion"`
	StormConfiguration           interface{}                               `pulumi:"stormConfiguration"`
	Tenant                       interface{}                               `pulumi:"tenant"`
	TimeToLive                   interface{}                               `pulumi:"timeToLive"`
	Type                         string                                    `pulumi:"type"`
	Version                      interface{}                               `pulumi:"version"`
	YarnConfiguration            interface{}                               `pulumi:"yarnConfiguration"`
	ZookeeperNodeSize            interface{}                               `pulumi:"zookeeperNodeSize"`
}

type HDInsightPigActivity struct {
	Arguments             []interface{}            `pulumi:"arguments"`
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
}

type HDInsightPigActivityResponse struct {
	Arguments             []interface{}                    `pulumi:"arguments"`
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

type HdfsSource struct {
	DistcpSettings   *DistcpSettings `pulumi:"distcpSettings"`
	Recursive        interface{}     `pulumi:"recursive"`
	SourceRetryCount interface{}     `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{}     `pulumi:"sourceRetryWait"`
	Type             string          `pulumi:"type"`
}

type HdfsSourceResponse struct {
	DistcpSettings   *DistcpSettingsResponse `pulumi:"distcpSettings"`
	Recursive        interface{}             `pulumi:"recursive"`
	SourceRetryCount interface{}             `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{}             `pulumi:"sourceRetryWait"`
	Type             string                  `pulumi:"type"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type HiveObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type HiveSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type HiveSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type HttpDataset struct {
	AdditionalHeaders interface{}                       `pulumi:"additionalHeaders"`
	Annotations       []interface{}                     `pulumi:"annotations"`
	Compression       interface{}                       `pulumi:"compression"`
	Description       *string                           `pulumi:"description"`
	Format            interface{}                       `pulumi:"format"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	RelativeUrl       interface{}                       `pulumi:"relativeUrl"`
	RequestBody       interface{}                       `pulumi:"requestBody"`
	RequestMethod     interface{}                       `pulumi:"requestMethod"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type HttpDatasetResponse struct {
	AdditionalHeaders interface{}                               `pulumi:"additionalHeaders"`
	Annotations       []interface{}                             `pulumi:"annotations"`
	Compression       interface{}                               `pulumi:"compression"`
	Description       *string                                   `pulumi:"description"`
	Format            interface{}                               `pulumi:"format"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	RelativeUrl       interface{}                               `pulumi:"relativeUrl"`
	RequestBody       interface{}                               `pulumi:"requestBody"`
	RequestMethod     interface{}                               `pulumi:"requestMethod"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type HttpLinkedService struct {
	Annotations                       []interface{}                     `pulumi:"annotations"`
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

type HttpSource struct {
	HttpRequestTimeout interface{} `pulumi:"httpRequestTimeout"`
	SourceRetryCount   interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait    interface{} `pulumi:"sourceRetryWait"`
	Type               string      `pulumi:"type"`
}

type HttpSourceResponse struct {
	HttpRequestTimeout interface{} `pulumi:"httpRequestTimeout"`
	SourceRetryCount   interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait    interface{} `pulumi:"sourceRetryWait"`
	Type               string      `pulumi:"type"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type HubspotObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type HubspotSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type HubspotSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type IfConditionActivity struct {
	DependsOn         []ActivityDependency `pulumi:"dependsOn"`
	Description       *string              `pulumi:"description"`
	Expression        Expression           `pulumi:"expression"`
	IfFalseActivities []interface{}        `pulumi:"ifFalseActivities"`
	IfTrueActivities  []interface{}        `pulumi:"ifTrueActivities"`
	Name              string               `pulumi:"name"`
	Type              string               `pulumi:"type"`
}

type IfConditionActivityResponse struct {
	DependsOn         []ActivityDependencyResponse `pulumi:"dependsOn"`
	Description       *string                      `pulumi:"description"`
	Expression        ExpressionResponse           `pulumi:"expression"`
	IfFalseActivities []interface{}                `pulumi:"ifFalseActivities"`
	IfTrueActivities  []interface{}                `pulumi:"ifTrueActivities"`
	Name              string                       `pulumi:"name"`
	Type              string                       `pulumi:"type"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type ImpalaObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type ImpalaSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type ImpalaSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type IntegrationRuntimeComputeProperties struct {
	Location                     *string                           `pulumi:"location"`
	MaxParallelExecutionsPerNode *int                              `pulumi:"maxParallelExecutionsPerNode"`
	NodeSize                     *string                           `pulumi:"nodeSize"`
	NumberOfNodes                *int                              `pulumi:"numberOfNodes"`
	VNetProperties               *IntegrationRuntimeVNetProperties `pulumi:"vNetProperties"`
}

type IntegrationRuntimeComputePropertiesResponse struct {
	Location                     *string                                   `pulumi:"location"`
	MaxParallelExecutionsPerNode *int                                      `pulumi:"maxParallelExecutionsPerNode"`
	NodeSize                     *string                                   `pulumi:"nodeSize"`
	NumberOfNodes                *int                                      `pulumi:"numberOfNodes"`
	VNetProperties               *IntegrationRuntimeVNetPropertiesResponse `pulumi:"vNetProperties"`
}

type IntegrationRuntimeCustomSetupScriptProperties struct {
	BlobContainerUri *string       `pulumi:"blobContainerUri"`
	SasToken         *SecureString `pulumi:"sasToken"`
}

type IntegrationRuntimeCustomSetupScriptPropertiesResponse struct {
	BlobContainerUri *string               `pulumi:"blobContainerUri"`
	SasToken         *SecureStringResponse `pulumi:"sasToken"`
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
}

type IntegrationRuntimeSsisCatalogInfoResponse struct {
	CatalogAdminPassword  *SecureStringResponse `pulumi:"catalogAdminPassword"`
	CatalogAdminUserName  *string               `pulumi:"catalogAdminUserName"`
	CatalogPricingTier    *string               `pulumi:"catalogPricingTier"`
	CatalogServerEndpoint *string               `pulumi:"catalogServerEndpoint"`
}

type IntegrationRuntimeSsisProperties struct {
	CatalogInfo                 *IntegrationRuntimeSsisCatalogInfo             `pulumi:"catalogInfo"`
	CustomSetupScriptProperties *IntegrationRuntimeCustomSetupScriptProperties `pulumi:"customSetupScriptProperties"`
	DataProxyProperties         *IntegrationRuntimeDataProxyProperties         `pulumi:"dataProxyProperties"`
	Edition                     *string                                        `pulumi:"edition"`
	LicenseType                 *string                                        `pulumi:"licenseType"`
}

type IntegrationRuntimeSsisPropertiesResponse struct {
	CatalogInfo                 *IntegrationRuntimeSsisCatalogInfoResponse             `pulumi:"catalogInfo"`
	CustomSetupScriptProperties *IntegrationRuntimeCustomSetupScriptPropertiesResponse `pulumi:"customSetupScriptProperties"`
	DataProxyProperties         *IntegrationRuntimeDataProxyPropertiesResponse         `pulumi:"dataProxyProperties"`
	Edition                     *string                                                `pulumi:"edition"`
	LicenseType                 *string                                                `pulumi:"licenseType"`
}

type IntegrationRuntimeVNetProperties struct {
	Subnet *string `pulumi:"subnet"`
	VNetId *string `pulumi:"vNetId"`
}

type IntegrationRuntimeVNetPropertiesResponse struct {
	Subnet *string `pulumi:"subnet"`
	VNetId *string `pulumi:"vNetId"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type JiraObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type JiraSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type JiraSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type JsonFormat struct {
	Deserializer       interface{} `pulumi:"deserializer"`
	EncodingName       interface{} `pulumi:"encodingName"`
	FilePattern        *string     `pulumi:"filePattern"`
	JsonNodeReference  interface{} `pulumi:"jsonNodeReference"`
	JsonPathDefinition interface{} `pulumi:"jsonPathDefinition"`
	NestingSeparator   interface{} `pulumi:"nestingSeparator"`
	Serializer         interface{} `pulumi:"serializer"`
	Type               string      `pulumi:"type"`
}

type JsonFormatResponse struct {
	Deserializer       interface{} `pulumi:"deserializer"`
	EncodingName       interface{} `pulumi:"encodingName"`
	FilePattern        *string     `pulumi:"filePattern"`
	JsonNodeReference  interface{} `pulumi:"jsonNodeReference"`
	JsonPathDefinition interface{} `pulumi:"jsonPathDefinition"`
	NestingSeparator   interface{} `pulumi:"nestingSeparator"`
	Serializer         interface{} `pulumi:"serializer"`
	Type               string      `pulumi:"type"`
}

type LinkedIntegrationRuntimeKey struct {
	AuthorizationType string       `pulumi:"authorizationType"`
	Key               SecureString `pulumi:"key"`
}

type LinkedIntegrationRuntimeKeyResponse struct {
	AuthorizationType string               `pulumi:"authorizationType"`
	Key               SecureStringResponse `pulumi:"key"`
}

type LinkedIntegrationRuntimeRbac struct {
	AuthorizationType string `pulumi:"authorizationType"`
	ResourceId        string `pulumi:"resourceId"`
}

type LinkedIntegrationRuntimeRbacResponse struct {
	AuthorizationType string `pulumi:"authorizationType"`
	ResourceId        string `pulumi:"resourceId"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type MagentoObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type MagentoSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type MagentoSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type ManagedIntegrationRuntime struct {
	ComputeProperties *IntegrationRuntimeComputeProperties `pulumi:"computeProperties"`
	Description       *string                              `pulumi:"description"`
	SsisProperties    *IntegrationRuntimeSsisProperties    `pulumi:"ssisProperties"`
	Type              string                               `pulumi:"type"`
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
	ComputeProperties *IntegrationRuntimeComputePropertiesResponse `pulumi:"computeProperties"`
	Description       *string                                      `pulumi:"description"`
	SsisProperties    *IntegrationRuntimeSsisPropertiesResponse    `pulumi:"ssisProperties"`
	State             string                                       `pulumi:"state"`
	Type              string                                       `pulumi:"type"`
}

type ManagedIntegrationRuntimeStatusResponse struct {
	CreateTime      string                                           `pulumi:"createTime"`
	DataFactoryName string                                           `pulumi:"dataFactoryName"`
	LastOperation   ManagedIntegrationRuntimeOperationResultResponse `pulumi:"lastOperation"`
	Nodes           []ManagedIntegrationRuntimeNodeResponse          `pulumi:"nodes"`
	OtherErrors     []ManagedIntegrationRuntimeErrorResponse         `pulumi:"otherErrors"`
	State           string                                           `pulumi:"state"`
	Type            *string                                          `pulumi:"type"`
}

type MariaDBLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Type                string                            `pulumi:"type"`
}

type MariaDBLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
}

type MariaDBSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type MariaDBSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type MariaDBTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type MariaDBTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type MarketoObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type MarketoSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type MarketoSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type MongoDbCollectionDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	CollectionName    interface{}                       `pulumi:"collectionName"`
	Description       *string                           `pulumi:"description"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type MongoDbCollectionDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	CollectionName    interface{}                               `pulumi:"collectionName"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
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
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type MongoDbSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type MultiplePipelineTrigger struct {
	Description *string                    `pulumi:"description"`
	Pipelines   []TriggerPipelineReference `pulumi:"pipelines"`
	Type        string                     `pulumi:"type"`
}

type MultiplePipelineTriggerResponse struct {
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
	Type                string                            `pulumi:"type"`
}

type MySqlLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
}

type NetezzaLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Type                string                            `pulumi:"type"`
}

type NetezzaLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
}

type NetezzaSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type NetezzaSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type NetezzaTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type NetezzaTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type ODataLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	AuthenticationType  *string                           `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Type                string                            `pulumi:"type"`
	Url                 interface{}                       `pulumi:"url"`
	UserName            interface{}                       `pulumi:"userName"`
}

type ODataLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	AuthenticationType  *string                                   `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
	Url                 interface{}                               `pulumi:"url"`
	UserName            interface{}                               `pulumi:"userName"`
}

type ODataResourceDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Path              interface{}                       `pulumi:"path"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type ODataResourceDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Path              interface{}                               `pulumi:"path"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
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
	PreCopyScript     interface{} `pulumi:"preCopyScript"`
	SinkRetryCount    interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait     interface{} `pulumi:"sinkRetryWait"`
	Type              string      `pulumi:"type"`
	WriteBatchSize    interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout interface{} `pulumi:"writeBatchTimeout"`
}

type OdbcSinkResponse struct {
	PreCopyScript     interface{} `pulumi:"preCopyScript"`
	SinkRetryCount    interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait     interface{} `pulumi:"sinkRetryWait"`
	Type              string      `pulumi:"type"`
	WriteBatchSize    interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout interface{} `pulumi:"writeBatchTimeout"`
}

type OracleLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Type                string                            `pulumi:"type"`
}

type OracleLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
}

type OracleSink struct {
	PreCopyScript     interface{} `pulumi:"preCopyScript"`
	SinkRetryCount    interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait     interface{} `pulumi:"sinkRetryWait"`
	Type              string      `pulumi:"type"`
	WriteBatchSize    interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout interface{} `pulumi:"writeBatchTimeout"`
}

type OracleSinkResponse struct {
	PreCopyScript     interface{} `pulumi:"preCopyScript"`
	SinkRetryCount    interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait     interface{} `pulumi:"sinkRetryWait"`
	Type              string      `pulumi:"type"`
	WriteBatchSize    interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout interface{} `pulumi:"writeBatchTimeout"`
}

type OracleSource struct {
	OracleReaderQuery interface{} `pulumi:"oracleReaderQuery"`
	QueryTimeout      interface{} `pulumi:"queryTimeout"`
	SourceRetryCount  interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait   interface{} `pulumi:"sourceRetryWait"`
	Type              string      `pulumi:"type"`
}

type OracleSourceResponse struct {
	OracleReaderQuery interface{} `pulumi:"oracleReaderQuery"`
	QueryTimeout      interface{} `pulumi:"queryTimeout"`
	SourceRetryCount  interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait   interface{} `pulumi:"sourceRetryWait"`
	Type              string      `pulumi:"type"`
}

type OracleTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type OracleTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type PaypalObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type PaypalSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type PaypalSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type PhoenixObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type PhoenixSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type PhoenixSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
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
	Type                string                            `pulumi:"type"`
}

type PostgreSqlLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type PrestoObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type PrestoSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type PrestoSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type QuickBooksLinkedService struct {
	AccessToken           interface{}                       `pulumi:"accessToken"`
	AccessTokenSecret     interface{}                       `pulumi:"accessTokenSecret"`
	Annotations           []interface{}                     `pulumi:"annotations"`
	CompanyId             interface{}                       `pulumi:"companyId"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type QuickBooksObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type QuickBooksSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type QuickBooksSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
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
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type RelationalSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type RelationalTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type RelationalTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type ResponsysObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type ResponsysSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type ResponsysSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
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
	AccessCredential  *SSISAccessCredential `pulumi:"accessCredential"`
	ConfigurationPath interface{}           `pulumi:"configurationPath"`
	PackagePassword   interface{}           `pulumi:"packagePassword"`
	PackagePath       interface{}           `pulumi:"packagePath"`
	Type              *string               `pulumi:"type"`
}

type SSISPackageLocationResponse struct {
	AccessCredential  *SSISAccessCredentialResponse `pulumi:"accessCredential"`
	ConfigurationPath interface{}                   `pulumi:"configurationPath"`
	PackagePassword   interface{}                   `pulumi:"packagePassword"`
	PackagePath       interface{}                   `pulumi:"packagePath"`
	Type              *string                       `pulumi:"type"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type SalesforceMarketingCloudObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type SalesforceMarketingCloudSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type SalesforceMarketingCloudSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type SalesforceObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	ObjectApiName     interface{}                       `pulumi:"objectApiName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type SalesforceObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	ObjectApiName     interface{}                               `pulumi:"objectApiName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type SalesforceSink struct {
	ExternalIdFieldName interface{} `pulumi:"externalIdFieldName"`
	IgnoreNullValues    interface{} `pulumi:"ignoreNullValues"`
	SinkRetryCount      interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait       interface{} `pulumi:"sinkRetryWait"`
	Type                string      `pulumi:"type"`
	WriteBatchSize      interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout   interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior       *string     `pulumi:"writeBehavior"`
}

type SalesforceSinkResponse struct {
	ExternalIdFieldName interface{} `pulumi:"externalIdFieldName"`
	IgnoreNullValues    interface{} `pulumi:"ignoreNullValues"`
	SinkRetryCount      interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait       interface{} `pulumi:"sinkRetryWait"`
	Type                string      `pulumi:"type"`
	WriteBatchSize      interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout   interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior       *string     `pulumi:"writeBehavior"`
}

type SalesforceSource struct {
	Query            interface{} `pulumi:"query"`
	ReadBehavior     *string     `pulumi:"readBehavior"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type SalesforceSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	ReadBehavior     *string     `pulumi:"readBehavior"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Path              interface{}                       `pulumi:"path"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type SapCloudForCustomerResourceDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Path              interface{}                               `pulumi:"path"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type SapCloudForCustomerSink struct {
	SinkRetryCount    interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait     interface{} `pulumi:"sinkRetryWait"`
	Type              string      `pulumi:"type"`
	WriteBatchSize    interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior     *string     `pulumi:"writeBehavior"`
}

type SapCloudForCustomerSinkResponse struct {
	SinkRetryCount    interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait     interface{} `pulumi:"sinkRetryWait"`
	Type              string      `pulumi:"type"`
	WriteBatchSize    interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior     *string     `pulumi:"writeBehavior"`
}

type SapCloudForCustomerSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type SapCloudForCustomerSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Path              interface{}                       `pulumi:"path"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type SapEccResourceDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Path              interface{}                               `pulumi:"path"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type SapEccSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type SapEccSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type SapHanaLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	AuthenticationType  *string                           `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
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
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Server              interface{}                               `pulumi:"server"`
	Type                string                                    `pulumi:"type"`
	UserName            interface{}                               `pulumi:"userName"`
}

type ScheduleTrigger struct {
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
	Description  *string                            `pulumi:"description"`
	Pipelines    []TriggerPipelineReferenceResponse `pulumi:"pipelines"`
	Recurrence   ScheduleTriggerRecurrenceResponse  `pulumi:"recurrence"`
	RuntimeState string                             `pulumi:"runtimeState"`
	Type         string                             `pulumi:"type"`
}

type SecureString struct {
	Type  string `pulumi:"type"`
	Value string `pulumi:"value"`
}

type SecureStringResponse struct {
	Type  string `pulumi:"type"`
	Value string `pulumi:"value"`
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
	Capabilities              map[string]string                          `pulumi:"capabilities"`
	CreateTime                string                                     `pulumi:"createTime"`
	DataFactoryName           string                                     `pulumi:"dataFactoryName"`
	InternalChannelEncryption string                                     `pulumi:"internalChannelEncryption"`
	Links                     []LinkedIntegrationRuntimeResponse         `pulumi:"links"`
	LocalTimeZoneOffset       string                                     `pulumi:"localTimeZoneOffset"`
	Nodes                     []SelfHostedIntegrationRuntimeNodeResponse `pulumi:"nodes"`
	ScheduledUpdateDate       string                                     `pulumi:"scheduledUpdateDate"`
	ServiceUrls               []string                                   `pulumi:"serviceUrls"`
	State                     string                                     `pulumi:"state"`
	TaskQueueId               string                                     `pulumi:"taskQueueId"`
	Type                      *string                                    `pulumi:"type"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type ServiceNowObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type ServiceNowSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type ServiceNowSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type ShopifyObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type ShopifySource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type ShopifySourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type SparkObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type SparkSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type SparkSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type SqlDWSink struct {
	AllowPolyBase     interface{}       `pulumi:"allowPolyBase"`
	PolyBaseSettings  *PolybaseSettings `pulumi:"polyBaseSettings"`
	PreCopyScript     interface{}       `pulumi:"preCopyScript"`
	SinkRetryCount    interface{}       `pulumi:"sinkRetryCount"`
	SinkRetryWait     interface{}       `pulumi:"sinkRetryWait"`
	Type              string            `pulumi:"type"`
	WriteBatchSize    interface{}       `pulumi:"writeBatchSize"`
	WriteBatchTimeout interface{}       `pulumi:"writeBatchTimeout"`
}

type SqlDWSinkResponse struct {
	AllowPolyBase     interface{}               `pulumi:"allowPolyBase"`
	PolyBaseSettings  *PolybaseSettingsResponse `pulumi:"polyBaseSettings"`
	PreCopyScript     interface{}               `pulumi:"preCopyScript"`
	SinkRetryCount    interface{}               `pulumi:"sinkRetryCount"`
	SinkRetryWait     interface{}               `pulumi:"sinkRetryWait"`
	Type              string                    `pulumi:"type"`
	WriteBatchSize    interface{}               `pulumi:"writeBatchSize"`
	WriteBatchTimeout interface{}               `pulumi:"writeBatchTimeout"`
}

type SqlDWSource struct {
	SourceRetryCount             interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait              interface{} `pulumi:"sourceRetryWait"`
	SqlReaderQuery               interface{} `pulumi:"sqlReaderQuery"`
	SqlReaderStoredProcedureName interface{} `pulumi:"sqlReaderStoredProcedureName"`
	StoredProcedureParameters    interface{} `pulumi:"storedProcedureParameters"`
	Type                         string      `pulumi:"type"`
}

type SqlDWSourceResponse struct {
	SourceRetryCount             interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait              interface{} `pulumi:"sourceRetryWait"`
	SqlReaderQuery               interface{} `pulumi:"sqlReaderQuery"`
	SqlReaderStoredProcedureName interface{} `pulumi:"sqlReaderStoredProcedureName"`
	StoredProcedureParameters    interface{} `pulumi:"storedProcedureParameters"`
	Type                         string      `pulumi:"type"`
}

type SqlServerLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Type                string                            `pulumi:"type"`
	UserName            interface{}                       `pulumi:"userName"`
}

type SqlServerLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
	UserName            interface{}                               `pulumi:"userName"`
}

type SqlServerStoredProcedureActivity struct {
	DependsOn                 []ActivityDependency                `pulumi:"dependsOn"`
	Description               *string                             `pulumi:"description"`
	LinkedServiceName         LinkedServiceReference              `pulumi:"linkedServiceName"`
	Name                      string                              `pulumi:"name"`
	Policy                    *ActivityPolicy                     `pulumi:"policy"`
	StoredProcedureName       interface{}                         `pulumi:"storedProcedureName"`
	StoredProcedureParameters map[string]StoredProcedureParameter `pulumi:"storedProcedureParameters"`
	Type                      string                              `pulumi:"type"`
}

type SqlServerStoredProcedureActivityResponse struct {
	DependsOn                 []ActivityDependencyResponse                `pulumi:"dependsOn"`
	Description               *string                                     `pulumi:"description"`
	LinkedServiceName         LinkedServiceReferenceResponse              `pulumi:"linkedServiceName"`
	Name                      string                                      `pulumi:"name"`
	Policy                    *ActivityPolicyResponse                     `pulumi:"policy"`
	StoredProcedureName       interface{}                                 `pulumi:"storedProcedureName"`
	StoredProcedureParameters map[string]StoredProcedureParameterResponse `pulumi:"storedProcedureParameters"`
	Type                      string                                      `pulumi:"type"`
}

type SqlServerTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type SqlServerTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type SqlSink struct {
	PreCopyScript                interface{}                         `pulumi:"preCopyScript"`
	SinkRetryCount               interface{}                         `pulumi:"sinkRetryCount"`
	SinkRetryWait                interface{}                         `pulumi:"sinkRetryWait"`
	SqlWriterStoredProcedureName interface{}                         `pulumi:"sqlWriterStoredProcedureName"`
	SqlWriterTableType           interface{}                         `pulumi:"sqlWriterTableType"`
	StoredProcedureParameters    map[string]StoredProcedureParameter `pulumi:"storedProcedureParameters"`
	Type                         string                              `pulumi:"type"`
	WriteBatchSize               interface{}                         `pulumi:"writeBatchSize"`
	WriteBatchTimeout            interface{}                         `pulumi:"writeBatchTimeout"`
}

type SqlSinkResponse struct {
	PreCopyScript                interface{}                                 `pulumi:"preCopyScript"`
	SinkRetryCount               interface{}                                 `pulumi:"sinkRetryCount"`
	SinkRetryWait                interface{}                                 `pulumi:"sinkRetryWait"`
	SqlWriterStoredProcedureName interface{}                                 `pulumi:"sqlWriterStoredProcedureName"`
	SqlWriterTableType           interface{}                                 `pulumi:"sqlWriterTableType"`
	StoredProcedureParameters    map[string]StoredProcedureParameterResponse `pulumi:"storedProcedureParameters"`
	Type                         string                                      `pulumi:"type"`
	WriteBatchSize               interface{}                                 `pulumi:"writeBatchSize"`
	WriteBatchTimeout            interface{}                                 `pulumi:"writeBatchTimeout"`
}

type SqlSource struct {
	SourceRetryCount             interface{}                         `pulumi:"sourceRetryCount"`
	SourceRetryWait              interface{}                         `pulumi:"sourceRetryWait"`
	SqlReaderQuery               interface{}                         `pulumi:"sqlReaderQuery"`
	SqlReaderStoredProcedureName interface{}                         `pulumi:"sqlReaderStoredProcedureName"`
	StoredProcedureParameters    map[string]StoredProcedureParameter `pulumi:"storedProcedureParameters"`
	Type                         string                              `pulumi:"type"`
}

type SqlSourceResponse struct {
	SourceRetryCount             interface{}                                 `pulumi:"sourceRetryCount"`
	SourceRetryWait              interface{}                                 `pulumi:"sourceRetryWait"`
	SqlReaderQuery               interface{}                                 `pulumi:"sqlReaderQuery"`
	SqlReaderStoredProcedureName interface{}                                 `pulumi:"sqlReaderStoredProcedureName"`
	StoredProcedureParameters    map[string]StoredProcedureParameterResponse `pulumi:"storedProcedureParameters"`
	Type                         string                                      `pulumi:"type"`
}

type SquareLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	ClientId              interface{}                       `pulumi:"clientId"`
	ClientSecret          interface{}                       `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type SquareObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type SquareSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type SquareSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
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

type TeradataLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	AuthenticationType  *string                           `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
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
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Server              interface{}                               `pulumi:"server"`
	Type                string                                    `pulumi:"type"`
	Username            interface{}                               `pulumi:"username"`
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

type TriggerPipelineReference struct {
	Parameters        map[string]interface{} `pulumi:"parameters"`
	PipelineReference *PipelineReference     `pulumi:"pipelineReference"`
}

type TriggerPipelineReferenceResponse struct {
	Parameters        map[string]interface{}     `pulumi:"parameters"`
	PipelineReference *PipelineReferenceResponse `pulumi:"pipelineReference"`
}

type TumblingWindowTrigger struct {
	Delay          interface{}              `pulumi:"delay"`
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

type TumblingWindowTriggerResponse struct {
	Delay          interface{}                      `pulumi:"delay"`
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

type UntilActivity struct {
	Activities  []interface{}        `pulumi:"activities"`
	DependsOn   []ActivityDependency `pulumi:"dependsOn"`
	Description *string              `pulumi:"description"`
	Expression  Expression           `pulumi:"expression"`
	Name        string               `pulumi:"name"`
	Timeout     interface{}          `pulumi:"timeout"`
	Type        string               `pulumi:"type"`
}

type UntilActivityResponse struct {
	Activities  []interface{}                `pulumi:"activities"`
	DependsOn   []ActivityDependencyResponse `pulumi:"dependsOn"`
	Description *string                      `pulumi:"description"`
	Expression  ExpressionResponse           `pulumi:"expression"`
	Name        string                       `pulumi:"name"`
	Timeout     interface{}                  `pulumi:"timeout"`
	Type        string                       `pulumi:"type"`
}

type VerticaLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Type                string                            `pulumi:"type"`
}

type VerticaLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
}

type VerticaSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type VerticaSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type VerticaTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type VerticaTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type WaitActivity struct {
	DependsOn         []ActivityDependency `pulumi:"dependsOn"`
	Description       *string              `pulumi:"description"`
	Name              string               `pulumi:"name"`
	Type              string               `pulumi:"type"`
	WaitTimeInSeconds int                  `pulumi:"waitTimeInSeconds"`
}

type WaitActivityResponse struct {
	DependsOn         []ActivityDependencyResponse `pulumi:"dependsOn"`
	Description       *string                      `pulumi:"description"`
	Name              string                       `pulumi:"name"`
	Type              string                       `pulumi:"type"`
	WaitTimeInSeconds int                          `pulumi:"waitTimeInSeconds"`
}

type WebActivity struct {
	Authentication    *WebActivityAuthentication `pulumi:"authentication"`
	Body              interface{}                `pulumi:"body"`
	Datasets          []DatasetReference         `pulumi:"datasets"`
	DependsOn         []ActivityDependency       `pulumi:"dependsOn"`
	Description       *string                    `pulumi:"description"`
	Headers           interface{}                `pulumi:"headers"`
	LinkedServiceName *LinkedServiceReference    `pulumi:"linkedServiceName"`
	LinkedServices    []LinkedServiceReference   `pulumi:"linkedServices"`
	Method            string                     `pulumi:"method"`
	Name              string                     `pulumi:"name"`
	Policy            *ActivityPolicy            `pulumi:"policy"`
	Type              string                     `pulumi:"type"`
	Url               interface{}                `pulumi:"url"`
}

type WebActivityAuthentication struct {
	Password *SecureString `pulumi:"password"`
	Pfx      *SecureString `pulumi:"pfx"`
	Resource *string       `pulumi:"resource"`
	Type     string        `pulumi:"type"`
	Username *string       `pulumi:"username"`
}

type WebActivityAuthenticationResponse struct {
	Password *SecureStringResponse `pulumi:"password"`
	Pfx      *SecureStringResponse `pulumi:"pfx"`
	Resource *string               `pulumi:"resource"`
	Type     string                `pulumi:"type"`
	Username *string               `pulumi:"username"`
}

type WebActivityResponse struct {
	Authentication    *WebActivityAuthenticationResponse `pulumi:"authentication"`
	Body              interface{}                        `pulumi:"body"`
	Datasets          []DatasetReferenceResponse         `pulumi:"datasets"`
	DependsOn         []ActivityDependencyResponse       `pulumi:"dependsOn"`
	Description       *string                            `pulumi:"description"`
	Headers           interface{}                        `pulumi:"headers"`
	LinkedServiceName *LinkedServiceReferenceResponse    `pulumi:"linkedServiceName"`
	LinkedServices    []LinkedServiceReferenceResponse   `pulumi:"linkedServices"`
	Method            string                             `pulumi:"method"`
	Name              string                             `pulumi:"name"`
	Policy            *ActivityPolicyResponse            `pulumi:"policy"`
	Type              string                             `pulumi:"type"`
	Url               interface{}                        `pulumi:"url"`
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
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type WebSourceResponse struct {
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type WebTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Index             interface{}                       `pulumi:"index"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Path              interface{}                       `pulumi:"path"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type WebTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Index             interface{}                               `pulumi:"index"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Path              interface{}                               `pulumi:"path"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type XeroLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
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
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type XeroObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type XeroSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type XeroSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type ZohoLinkedService struct {
	AccessToken           interface{}                       `pulumi:"accessToken"`
	Annotations           []interface{}                     `pulumi:"annotations"`
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

type ZohoLinkedServiceResponse struct {
	AccessToken           interface{}                               `pulumi:"accessToken"`
	Annotations           []interface{}                             `pulumi:"annotations"`
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

type ZohoObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type ZohoObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type ZohoSource struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

type ZohoSourceResponse struct {
	Query            interface{} `pulumi:"query"`
	SourceRetryCount interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait  interface{} `pulumi:"sourceRetryWait"`
	Type             string      `pulumi:"type"`
}

func init() {
	pulumi.RegisterOutputType(FactoryIdentityOutput{})
	pulumi.RegisterOutputType(FactoryIdentityPtrOutput{})
	pulumi.RegisterOutputType(FactoryIdentityResponseOutput{})
	pulumi.RegisterOutputType(FactoryIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(FactoryVSTSConfigurationOutput{})
	pulumi.RegisterOutputType(FactoryVSTSConfigurationPtrOutput{})
	pulumi.RegisterOutputType(FactoryVSTSConfigurationResponseOutput{})
	pulumi.RegisterOutputType(FactoryVSTSConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ParameterSpecificationOutput{})
	pulumi.RegisterOutputType(ParameterSpecificationMapOutput{})
	pulumi.RegisterOutputType(ParameterSpecificationResponseOutput{})
	pulumi.RegisterOutputType(ParameterSpecificationResponseMapOutput{})
}
