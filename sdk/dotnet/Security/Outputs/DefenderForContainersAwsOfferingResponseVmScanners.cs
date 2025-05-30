// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Outputs
{

    /// <summary>
    /// The Microsoft Defender for Container K8s VM host scanning configuration
    /// </summary>
    [OutputType]
    public sealed class DefenderForContainersAwsOfferingResponseVmScanners
    {
        /// <summary>
        /// The cloud role ARN in AWS for this feature
        /// </summary>
        public readonly string? CloudRoleArn;
        /// <summary>
        /// Configuration for VM scanning
        /// </summary>
        public readonly Outputs.VmScannersBaseResponseConfiguration? Configuration;
        /// <summary>
        /// Is VM scanning enabled
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private DefenderForContainersAwsOfferingResponseVmScanners(
            string? cloudRoleArn,

            Outputs.VmScannersBaseResponseConfiguration? configuration,

            bool? enabled)
        {
            CloudRoleArn = cloudRoleArn;
            Configuration = configuration;
            Enabled = enabled;
        }
    }
}
