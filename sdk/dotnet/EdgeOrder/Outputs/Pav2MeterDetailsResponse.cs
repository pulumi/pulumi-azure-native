// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EdgeOrder.Outputs
{

    /// <summary>
    /// Billing type PAV2 meter details.
    /// </summary>
    [OutputType]
    public sealed class Pav2MeterDetailsResponse
    {
        /// <summary>
        /// Represents billing type.
        /// Expected value is 'Pav2'.
        /// </summary>
        public readonly string BillingType;
        /// <summary>
        /// Charging type.
        /// </summary>
        public readonly string ChargingType;
        /// <summary>
        /// Validation status of requested data center and transport.
        /// </summary>
        public readonly string MeterGuid;
        /// <summary>
        /// Billing unit applicable for Pav2 billing.
        /// </summary>
        public readonly double Multiplier;

        [OutputConstructor]
        private Pav2MeterDetailsResponse(
            string billingType,

            string chargingType,

            string meterGuid,

            double multiplier)
        {
            BillingType = billingType;
            ChargingType = chargingType;
            MeterGuid = meterGuid;
            Multiplier = multiplier;
        }
    }
}
