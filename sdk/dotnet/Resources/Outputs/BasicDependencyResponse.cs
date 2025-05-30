// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Resources.Outputs
{

    /// <summary>
    /// Deployment dependency information.
    /// </summary>
    [OutputType]
    public sealed class BasicDependencyResponse
    {
        /// <summary>
        /// The ID of the dependency.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The dependency resource name.
        /// </summary>
        public readonly string? ResourceName;
        /// <summary>
        /// The dependency resource type.
        /// </summary>
        public readonly string? ResourceType;

        [OutputConstructor]
        private BasicDependencyResponse(
            string? id,

            string? resourceName,

            string? resourceType)
        {
            Id = id;
            ResourceName = resourceName;
            ResourceType = resourceType;
        }
    }
}
