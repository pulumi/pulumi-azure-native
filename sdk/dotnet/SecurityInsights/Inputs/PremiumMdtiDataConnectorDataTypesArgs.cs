// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.Inputs
{

    /// <summary>
    /// The available data types for Premium Microsoft Defender for Threat Intelligence data connector.
    /// </summary>
    public sealed class PremiumMdtiDataConnectorDataTypesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Data type for Premium Microsoft Defender for Threat Intelligence data connector.
        /// </summary>
        [Input("connector", required: true)]
        public Input<Inputs.PremiumMdtiDataConnectorDataTypesConnectorArgs> Connector { get; set; } = null!;

        public PremiumMdtiDataConnectorDataTypesArgs()
        {
        }
        public static new PremiumMdtiDataConnectorDataTypesArgs Empty => new PremiumMdtiDataConnectorDataTypesArgs();
    }
}
