// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.V20231201Preview.Inputs
{

    /// <summary>
    /// Logs data type.
    /// </summary>
    public sealed class MicrosoftPurviewInformationProtectionConnectorDataTypesLogsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describe whether this data type connection is enabled or not.
        /// </summary>
        [Input("state", required: true)]
        public InputUnion<string, Pulumi.AzureNative.SecurityInsights.V20231201Preview.DataTypeState> State { get; set; } = null!;

        public MicrosoftPurviewInformationProtectionConnectorDataTypesLogsArgs()
        {
        }
        public static new MicrosoftPurviewInformationProtectionConnectorDataTypesLogsArgs Empty => new MicrosoftPurviewInformationProtectionConnectorDataTypesLogsArgs();
    }
}
