// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Outputs
{

    [OutputType]
    public sealed class EndpointModelSkuCapacityPropertiesResponse
    {
        /// <summary>
        /// The default capacity.
        /// </summary>
        public readonly int? Default;
        /// <summary>
        /// The maximum capacity.
        /// </summary>
        public readonly int? Maximum;

        [OutputConstructor]
        private EndpointModelSkuCapacityPropertiesResponse(
            int? @default,

            int? maximum)
        {
            Default = @default;
            Maximum = maximum;
        }
    }
}
