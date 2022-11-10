


package v20170901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
	Authentication        *WebActivityAuthentication `pulumi:"authentication"`
	Body                  interface{}                `pulumi:"body"`
	Datasets              []DatasetReference         `pulumi:"datasets"`
	DependsOn             []ActivityDependency       `pulumi:"dependsOn"`
	Description           *string                    `pulumi:"description"`
	DisableCertValidation *bool                      `pulumi:"disableCertValidation"`
	Headers               interface{}                `pulumi:"headers"`
	LinkedServiceName     *LinkedServiceReference    `pulumi:"linkedServiceName"`
	LinkedServices        []LinkedServiceReference   `pulumi:"linkedServices"`
	Method                string                     `pulumi:"method"`
	Name                  string                     `pulumi:"name"`
	Policy                *ActivityPolicy            `pulumi:"policy"`
	Type                  string                     `pulumi:"type"`
	Url                   interface{}                `pulumi:"url"`
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
	Authentication        *WebActivityAuthenticationResponse `pulumi:"authentication"`
	Body                  interface{}                        `pulumi:"body"`
	Datasets              []DatasetReferenceResponse         `pulumi:"datasets"`
	DependsOn             []ActivityDependencyResponse       `pulumi:"dependsOn"`
	Description           *string                            `pulumi:"description"`
	DisableCertValidation *bool                              `pulumi:"disableCertValidation"`
	Headers               interface{}                        `pulumi:"headers"`
	LinkedServiceName     *LinkedServiceReferenceResponse    `pulumi:"linkedServiceName"`
	LinkedServices        []LinkedServiceReferenceResponse   `pulumi:"linkedServices"`
	Method                string                             `pulumi:"method"`
	Name                  string                             `pulumi:"name"`
	Policy                *ActivityPolicyResponse            `pulumi:"policy"`
	Type                  string                             `pulumi:"type"`
	Url                   interface{}                        `pulumi:"url"`
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
}
