// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of awsOpenSearchDomainStatus
    /// </summary>
    [OutputType]
    public sealed class AwsOpenSearchDomainStatusPropertiesResponse
    {
        /// <summary>
        /// &lt;p&gt;Identity and Access Management (IAM) policy document specifying the access policies for the domain.&lt;/p&gt;
        /// </summary>
        public readonly string? AccessPolicies;
        /// <summary>
        /// &lt;p&gt;Key-value pairs that specify advanced configuration options.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableDictionary<string, string>? AdvancedOptions;
        /// <summary>
        /// &lt;p&gt;Settings for fine-grained access control.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.AdvancedSecurityOptionsResponse? AdvancedSecurityOptions;
        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name (ARN) of the domain. For more information, see &lt;a href='https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html'&gt;IAM identifiers &lt;/a&gt; in the &lt;i&gt;AWS Identity and Access Management User Guide&lt;/i&gt;.&lt;/p&gt;
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// &lt;p&gt;Auto-Tune settings for the domain.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.AutoTuneOptionsOutputResponse? AutoTuneOptions;
        /// <summary>
        /// &lt;p&gt;Information about a configuration change happening on the domain.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.ChangeProgressDetailsResponse? ChangeProgressDetails;
        /// <summary>
        /// &lt;p&gt;Container for the cluster configuration of the domain.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.ClusterConfigResponse? ClusterConfig;
        /// <summary>
        /// &lt;p&gt;Key-value pairs to configure Amazon Cognito authentication for OpenSearch Dashboards.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.CognitoOptionsResponse? CognitoOptions;
        /// <summary>
        /// &lt;p&gt;Creation status of an OpenSearch Service domain. True if domain creation is complete. False if domain creation is still in progress.&lt;/p&gt;
        /// </summary>
        public readonly bool? Created;
        /// <summary>
        /// &lt;p&gt;Deletion status of an OpenSearch Service domain. True if domain deletion is complete. False if domain deletion is still in progress. Once deletion is complete, the status of the domain is no longer returned.&lt;/p&gt;
        /// </summary>
        public readonly bool? Deleted;
        /// <summary>
        /// &lt;p&gt;Additional options for the domain endpoint, such as whether to require HTTPS for all traffic.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.DomainEndpointOptionsResponse? DomainEndpointOptions;
        /// <summary>
        /// &lt;p&gt;Unique identifier for the domain.&lt;/p&gt;
        /// </summary>
        public readonly string? DomainId;
        /// <summary>
        /// &lt;p&gt;Name of the domain. Domain names are unique across all domains owned by the same account within an Amazon Web Services Region.&lt;/p&gt;
        /// </summary>
        public readonly string? DomainName;
        /// <summary>
        /// &lt;p&gt;The status of any changes that are currently in progress for the domain.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.DomainProcessingStatusTypeEnumValueResponse? DomainProcessingStatus;
        /// <summary>
        /// &lt;p&gt;Container for EBS-based storage settings for the domain.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.EBSOptionsResponse? EbsOptions;
        /// <summary>
        /// &lt;p&gt;Encryption at rest settings for the domain.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.EncryptionAtRestOptionsResponse? EncryptionAtRestOptions;
        /// <summary>
        /// &lt;p&gt;Domain-specific endpoint used to submit index, search, and data upload requests to the domain.&lt;/p&gt;
        /// </summary>
        public readonly string? Endpoint;
        /// <summary>
        /// &lt;p&gt;If &lt;code&gt;IPAddressType&lt;/code&gt; to set to &lt;code&gt;dualstack&lt;/code&gt;, a version 2 domain endpoint is provisioned. This endpoint functions like a normal endpoint, except that it works with both IPv4 and IPv6 IP addresses. Normal endpoints work only with IPv4 IP addresses. &lt;/p&gt;
        /// </summary>
        public readonly string? EndpointV2;
        /// <summary>
        /// &lt;p&gt;The key-value pair that exists if the OpenSearch Service domain uses VPC endpoints. Example &lt;code&gt;key, value&lt;/code&gt;: &lt;code&gt;'vpc','vpc-endpoint-h2dsd34efgyghrtguk5gt6j2foh4.us-east-1.es.amazonaws.com'&lt;/code&gt;.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Endpoints;
        /// <summary>
        /// &lt;p&gt;Version of OpenSearch or Elasticsearch that the domain is running, in the format &lt;code&gt;Elasticsearch_X.Y&lt;/code&gt; or &lt;code&gt;OpenSearch_X.Y&lt;/code&gt;.&lt;/p&gt;
        /// </summary>
        public readonly string? EngineVersion;
        /// <summary>
        /// &lt;p&gt;The type of IP addresses supported by the endpoint for the domain.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.IPAddressTypeEnumValueResponse? IpAddressType;
        /// <summary>
        /// &lt;p&gt;Log publishing options for the domain.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.LogPublishingOptionResponse>? LogPublishingOptions;
        /// <summary>
        /// &lt;p&gt;Information about the domain properties that are currently being modified.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.ModifyingPropertiesResponse> ModifyingProperties;
        /// <summary>
        /// &lt;p&gt;Whether node-to-node encryption is enabled or disabled.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.NodeToNodeEncryptionOptionsResponse? NodeToNodeEncryptionOptions;
        /// <summary>
        /// &lt;p&gt;Options that specify a custom 10-hour window during which OpenSearch Service can perform configuration changes on the domain.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.OffPeakWindowOptionsResponse? OffPeakWindowOptions;
        /// <summary>
        /// &lt;p&gt;The status of the domain configuration. True if OpenSearch Service is processing configuration changes. False if the configuration is active.&lt;/p&gt;
        /// </summary>
        public readonly bool? Processing;
        /// <summary>
        /// &lt;p&gt;The current status of the domain's service software.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.ServiceSoftwareOptionsResponse? ServiceSoftwareOptions;
        /// <summary>
        /// &lt;p&gt;DEPRECATED. Container for parameters required to configure automated snapshots of domain indexes.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.SnapshotOptionsResponse? SnapshotOptions;
        /// <summary>
        /// &lt;p&gt;Service software update options for the domain.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.SoftwareUpdateOptionsResponse? SoftwareUpdateOptions;
        /// <summary>
        /// &lt;p&gt;The status of a domain version upgrade to a new version of OpenSearch or Elasticsearch. True if OpenSearch Service is in the process of a version upgrade. False if the configuration is active.&lt;/p&gt;
        /// </summary>
        public readonly bool? UpgradeProcessing;
        /// <summary>
        /// &lt;p&gt;The VPC configuration for the domain.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.VPCDerivedInfoResponse? VpcOptions;

        [OutputConstructor]
        private AwsOpenSearchDomainStatusPropertiesResponse(
            string? accessPolicies,

            ImmutableDictionary<string, string>? advancedOptions,

            Outputs.AdvancedSecurityOptionsResponse? advancedSecurityOptions,

            string? arn,

            Outputs.AutoTuneOptionsOutputResponse? autoTuneOptions,

            Outputs.ChangeProgressDetailsResponse? changeProgressDetails,

            Outputs.ClusterConfigResponse? clusterConfig,

            Outputs.CognitoOptionsResponse? cognitoOptions,

            bool? created,

            bool? deleted,

            Outputs.DomainEndpointOptionsResponse? domainEndpointOptions,

            string? domainId,

            string? domainName,

            Outputs.DomainProcessingStatusTypeEnumValueResponse? domainProcessingStatus,

            Outputs.EBSOptionsResponse? ebsOptions,

            Outputs.EncryptionAtRestOptionsResponse? encryptionAtRestOptions,

            string? endpoint,

            string? endpointV2,

            ImmutableDictionary<string, string>? endpoints,

            string? engineVersion,

            Outputs.IPAddressTypeEnumValueResponse? ipAddressType,

            ImmutableDictionary<string, Outputs.LogPublishingOptionResponse>? logPublishingOptions,

            ImmutableArray<Outputs.ModifyingPropertiesResponse> modifyingProperties,

            Outputs.NodeToNodeEncryptionOptionsResponse? nodeToNodeEncryptionOptions,

            Outputs.OffPeakWindowOptionsResponse? offPeakWindowOptions,

            bool? processing,

            Outputs.ServiceSoftwareOptionsResponse? serviceSoftwareOptions,

            Outputs.SnapshotOptionsResponse? snapshotOptions,

            Outputs.SoftwareUpdateOptionsResponse? softwareUpdateOptions,

            bool? upgradeProcessing,

            Outputs.VPCDerivedInfoResponse? vpcOptions)
        {
            AccessPolicies = accessPolicies;
            AdvancedOptions = advancedOptions;
            AdvancedSecurityOptions = advancedSecurityOptions;
            Arn = arn;
            AutoTuneOptions = autoTuneOptions;
            ChangeProgressDetails = changeProgressDetails;
            ClusterConfig = clusterConfig;
            CognitoOptions = cognitoOptions;
            Created = created;
            Deleted = deleted;
            DomainEndpointOptions = domainEndpointOptions;
            DomainId = domainId;
            DomainName = domainName;
            DomainProcessingStatus = domainProcessingStatus;
            EbsOptions = ebsOptions;
            EncryptionAtRestOptions = encryptionAtRestOptions;
            Endpoint = endpoint;
            EndpointV2 = endpointV2;
            Endpoints = endpoints;
            EngineVersion = engineVersion;
            IpAddressType = ipAddressType;
            LogPublishingOptions = logPublishingOptions;
            ModifyingProperties = modifyingProperties;
            NodeToNodeEncryptionOptions = nodeToNodeEncryptionOptions;
            OffPeakWindowOptions = offPeakWindowOptions;
            Processing = processing;
            ServiceSoftwareOptions = serviceSoftwareOptions;
            SnapshotOptions = snapshotOptions;
            SoftwareUpdateOptions = softwareUpdateOptions;
            UpgradeProcessing = upgradeProcessing;
            VpcOptions = vpcOptions;
        }
    }
}
