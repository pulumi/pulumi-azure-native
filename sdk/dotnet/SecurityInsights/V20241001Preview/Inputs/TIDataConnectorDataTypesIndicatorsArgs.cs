// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.V20241001Preview.Inputs
{

    /// <summary>
    /// Data type for indicators connection.
    /// </summary>
    public sealed class TIDataConnectorDataTypesIndicatorsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describe whether this data type connection is enabled or not.
        /// </summary>
        [Input("state", required: true)]
        public InputUnion<string, Pulumi.AzureNative.SecurityInsights.V20241001Preview.DataTypeState> State { get; set; } = null!;

        public TIDataConnectorDataTypesIndicatorsArgs()
        {
        }
        public static new TIDataConnectorDataTypesIndicatorsArgs Empty => new TIDataConnectorDataTypesIndicatorsArgs();
    }
}
