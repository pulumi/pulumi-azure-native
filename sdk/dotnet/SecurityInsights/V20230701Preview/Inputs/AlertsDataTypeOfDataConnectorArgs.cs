// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.V20230701Preview.Inputs
{

    /// <summary>
    /// Alerts data type for data connectors.
    /// </summary>
    public sealed class AlertsDataTypeOfDataConnectorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alerts data type connection.
        /// </summary>
        [Input("alerts", required: true)]
        public Input<Inputs.DataConnectorDataTypeCommonArgs> Alerts { get; set; } = null!;

        public AlertsDataTypeOfDataConnectorArgs()
        {
        }
        public static new AlertsDataTypeOfDataConnectorArgs Empty => new AlertsDataTypeOfDataConnectorArgs();
    }
}
