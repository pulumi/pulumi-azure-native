// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorageSync.Outputs
{

    /// <summary>
    /// Server endpoint provisioning step status object.
    /// </summary>
    [OutputType]
    public sealed class ServerEndpointProvisioningStepStatusResponse
    {
        /// <summary>
        /// Additional information for the provisioning step
        /// </summary>
        public readonly ImmutableDictionary<string, string> AdditionalInformation;
        /// <summary>
        /// End time of the provisioning step
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// Error code (HResult) for the provisioning step
        /// </summary>
        public readonly int ErrorCode;
        /// <summary>
        /// Estimated completion time of the provisioning step in minutes
        /// </summary>
        public readonly int MinutesLeft;
        /// <summary>
        /// Name of the provisioning step
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Estimated progress percentage
        /// </summary>
        public readonly int ProgressPercentage;
        /// <summary>
        /// Start time of the provisioning step
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// Status of the provisioning step
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private ServerEndpointProvisioningStepStatusResponse(
            ImmutableDictionary<string, string> additionalInformation,

            string endTime,

            int errorCode,

            int minutesLeft,

            string name,

            int progressPercentage,

            string startTime,

            string status)
        {
            AdditionalInformation = additionalInformation;
            EndTime = endTime;
            ErrorCode = errorCode;
            MinutesLeft = minutesLeft;
            Name = name;
            ProgressPercentage = progressPercentage;
            StartTime = startTime;
            Status = status;
        }
    }
}
