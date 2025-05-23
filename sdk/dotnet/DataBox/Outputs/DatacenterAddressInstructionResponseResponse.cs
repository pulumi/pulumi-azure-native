// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataBox.Outputs
{

    /// <summary>
    /// Datacenter instruction for given storage location.
    /// </summary>
    [OutputType]
    public sealed class DatacenterAddressInstructionResponseResponse
    {
        /// <summary>
        /// Data center communication instruction
        /// </summary>
        public readonly string CommunicationInstruction;
        /// <summary>
        /// Azure Location where the Data Center serves primarily.
        /// </summary>
        public readonly string DataCenterAzureLocation;
        /// <summary>
        /// Data center address type
        /// Expected value is 'DatacenterAddressInstruction'.
        /// </summary>
        public readonly string DatacenterAddressType;
        /// <summary>
        /// List of supported carriers for return shipment.
        /// </summary>
        public readonly ImmutableArray<string> SupportedCarriersForReturnShipment;

        [OutputConstructor]
        private DatacenterAddressInstructionResponseResponse(
            string communicationInstruction,

            string dataCenterAzureLocation,

            string datacenterAddressType,

            ImmutableArray<string> supportedCarriersForReturnShipment)
        {
            CommunicationInstruction = communicationInstruction;
            DataCenterAzureLocation = dataCenterAzureLocation;
            DatacenterAddressType = datacenterAddressType;
            SupportedCarriersForReturnShipment = supportedCarriersForReturnShipment;
        }
    }
}
