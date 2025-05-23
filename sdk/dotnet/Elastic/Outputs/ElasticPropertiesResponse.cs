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
    /// Elastic Resource Properties.
    /// </summary>
    [OutputType]
    public sealed class ElasticPropertiesResponse
    {
        /// <summary>
        /// Details of the elastic cloud deployment.
        /// </summary>
        public readonly Outputs.ElasticCloudDeploymentResponse? ElasticCloudDeployment;
        /// <summary>
        /// Details of the user's elastic account.
        /// </summary>
        public readonly Outputs.ElasticCloudUserResponse? ElasticCloudUser;

        [OutputConstructor]
        private ElasticPropertiesResponse(
            Outputs.ElasticCloudDeploymentResponse? elasticCloudDeployment,

            Outputs.ElasticCloudUserResponse? elasticCloudUser)
        {
            ElasticCloudDeployment = elasticCloudDeployment;
            ElasticCloudUser = elasticCloudUser;
        }
    }
}
