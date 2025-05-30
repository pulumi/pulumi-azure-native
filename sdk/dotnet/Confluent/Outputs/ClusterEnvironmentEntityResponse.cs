// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Confluent.Outputs
{

    /// <summary>
    /// The environment to which cluster belongs
    /// </summary>
    [OutputType]
    public sealed class ClusterEnvironmentEntityResponse
    {
        /// <summary>
        /// Environment of the referred resource
        /// </summary>
        public readonly string? Environment;
        /// <summary>
        /// ID of the referred resource
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// API URL for accessing or modifying the referred object
        /// </summary>
        public readonly string? Related;
        /// <summary>
        /// CRN reference to the referred resource
        /// </summary>
        public readonly string? ResourceName;

        [OutputConstructor]
        private ClusterEnvironmentEntityResponse(
            string? environment,

            string? id,

            string? related,

            string? resourceName)
        {
            Environment = environment;
            Id = id;
            Related = related;
            ResourceName = resourceName;
        }
    }
}
