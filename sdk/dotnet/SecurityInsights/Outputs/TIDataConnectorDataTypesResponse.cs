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
    /// The available data types for TI (Threat Intelligence) data connector.
    /// </summary>
    [OutputType]
    public sealed class TIDataConnectorDataTypesResponse
    {
        /// <summary>
        /// Data type for indicators connection.
        /// </summary>
        public readonly Outputs.TIDataConnectorDataTypesResponseIndicators? Indicators;

        [OutputConstructor]
        private TIDataConnectorDataTypesResponse(Outputs.TIDataConnectorDataTypesResponseIndicators? indicators)
        {
            Indicators = indicators;
        }
    }
}
