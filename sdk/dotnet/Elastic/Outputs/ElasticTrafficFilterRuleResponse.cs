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
    /// Elastic traffic filter rule object
    /// </summary>
    [OutputType]
    public sealed class ElasticTrafficFilterRuleResponse
    {
        /// <summary>
        /// Guid of Private Endpoint in the elastic filter rule
        /// </summary>
        public readonly string? AzureEndpointGuid;
        /// <summary>
        /// Name of the Private Endpoint in the elastic filter rule
        /// </summary>
        public readonly string? AzureEndpointName;
        /// <summary>
        /// Description of the elastic filter rule
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Id of the elastic filter rule
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// IP of the elastic filter rule
        /// </summary>
        public readonly string? Source;

        [OutputConstructor]
        private ElasticTrafficFilterRuleResponse(
            string? azureEndpointGuid,

            string? azureEndpointName,

            string? description,

            string? id,

            string? source)
        {
            AzureEndpointGuid = azureEndpointGuid;
            AzureEndpointName = azureEndpointName;
            Description = description;
            Id = id;
            Source = source;
        }
    }
}
