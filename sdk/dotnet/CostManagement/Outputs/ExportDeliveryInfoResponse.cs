// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CostManagement.Outputs
{

    /// <summary>
    /// The delivery information associated with a export.
    /// </summary>
    [OutputType]
    public sealed class ExportDeliveryInfoResponse
    {
        /// <summary>
        /// Has destination for the export being delivered.
        /// </summary>
        public readonly Outputs.ExportDeliveryDestinationResponse Destination;

        [OutputConstructor]
        private ExportDeliveryInfoResponse(Outputs.ExportDeliveryDestinationResponse destination)
        {
            Destination = destination;
        }
    }
}
