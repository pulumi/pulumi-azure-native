// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Elastic.Outputs
{

    /// <summary>
    /// Details of the user's elastic deployment associated with the monitor resource.
    /// </summary>
    [OutputType]
    public sealed class ElasticCloudDeploymentResponse
    {
        /// <summary>
        /// Associated Azure subscription Id for the elastic deployment.
        /// </summary>
        public readonly string AzureSubscriptionId;
        /// <summary>
        /// Elastic deployment Id
        /// </summary>
        public readonly string DeploymentId;
        /// <summary>
        /// Region where Deployment at Elastic side took place.
        /// </summary>
        public readonly string ElasticsearchRegion;
        /// <summary>
        /// Elasticsearch ingestion endpoint of the Elastic deployment.
        /// </summary>
        public readonly string ElasticsearchServiceUrl;
        /// <summary>
        /// Kibana endpoint of the Elastic deployment.
        /// </summary>
        public readonly string KibanaServiceUrl;
        /// <summary>
        /// Kibana dashboard sso URL of the Elastic deployment.
        /// </summary>
        public readonly string KibanaSsoUrl;
        /// <summary>
        /// Elastic deployment name
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private ElasticCloudDeploymentResponse(
            string azureSubscriptionId,

            string deploymentId,

            string elasticsearchRegion,

            string elasticsearchServiceUrl,

            string kibanaServiceUrl,

            string kibanaSsoUrl,

            string name)
        {
            AzureSubscriptionId = azureSubscriptionId;
            DeploymentId = deploymentId;
            ElasticsearchRegion = elasticsearchRegion;
            ElasticsearchServiceUrl = elasticsearchServiceUrl;
            KibanaServiceUrl = kibanaServiceUrl;
            KibanaSsoUrl = kibanaSsoUrl;
            Name = name;
        }
    }
}
