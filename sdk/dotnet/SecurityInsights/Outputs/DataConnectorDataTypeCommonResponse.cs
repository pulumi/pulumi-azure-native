// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.Outputs
{

    /// <summary>
    /// Common field for data type in data connectors.
    /// </summary>
    [OutputType]
    public sealed class DataConnectorDataTypeCommonResponse
    {
        /// <summary>
        /// Describe whether this data type connection is enabled or not.
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private DataConnectorDataTypeCommonResponse(string? state)
        {
            State = state;
        }
    }
}
