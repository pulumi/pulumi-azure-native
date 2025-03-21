// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.V20240401Preview.Inputs
{

    /// <summary>
    /// Logs data type.
    /// </summary>
    public sealed class AwsCloudTrailDataConnectorDataTypesLogsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describe whether this data type connection is enabled or not.
        /// </summary>
        [Input("state", required: true)]
        public InputUnion<string, Pulumi.AzureNative.SecurityInsights.V20240401Preview.DataTypeState> State { get; set; } = null!;

        public AwsCloudTrailDataConnectorDataTypesLogsArgs()
        {
        }
        public static new AwsCloudTrailDataConnectorDataTypesLogsArgs Empty => new AwsCloudTrailDataConnectorDataTypesLogsArgs();
    }
}
