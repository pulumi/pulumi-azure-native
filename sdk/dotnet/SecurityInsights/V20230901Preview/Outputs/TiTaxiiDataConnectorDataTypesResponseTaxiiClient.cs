// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.V20230901Preview.Outputs
{

    /// <summary>
    /// Data type for TAXII connector.
    /// </summary>
    [OutputType]
    public sealed class TiTaxiiDataConnectorDataTypesResponseTaxiiClient
    {
        /// <summary>
        /// Describe whether this data type connection is enabled or not.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private TiTaxiiDataConnectorDataTypesResponseTaxiiClient(string state)
        {
            State = state;
        }
    }
}
