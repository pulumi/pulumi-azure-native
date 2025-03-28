// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevCenter.V20250201.Outputs
{

    /// <summary>
    /// Stop on no connect configuration settings for Dev Boxes created in this pool.
    /// </summary>
    [OutputType]
    public sealed class StopOnNoConnectConfigurationResponse
    {
        /// <summary>
        /// The specified time in minutes to wait before stopping a Dev Box if no connection is made.
        /// </summary>
        public readonly int? GracePeriodMinutes;
        /// <summary>
        /// Enables the feature to stop a started Dev Box when it has not been connected to, once the grace period has lapsed.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private StopOnNoConnectConfigurationResponse(
            int? gracePeriodMinutes,

            string? status)
        {
            GracePeriodMinutes = gracePeriodMinutes;
            Status = status;
        }
    }
}
