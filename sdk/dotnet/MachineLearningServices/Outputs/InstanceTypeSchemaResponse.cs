// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Outputs
{

    /// <summary>
    /// Instance type schema.
    /// </summary>
    [OutputType]
    public sealed class InstanceTypeSchemaResponse
    {
        /// <summary>
        /// Node Selector
        /// </summary>
        public readonly ImmutableDictionary<string, string>? NodeSelector;
        /// <summary>
        /// Resource requests/limits for this instance type
        /// </summary>
        public readonly Outputs.InstanceTypeSchemaResponseResources? Resources;

        [OutputConstructor]
        private InstanceTypeSchemaResponse(
            ImmutableDictionary<string, string>? nodeSelector,

            Outputs.InstanceTypeSchemaResponseResources? resources)
        {
            NodeSelector = nodeSelector;
            Resources = resources;
        }
    }
}
