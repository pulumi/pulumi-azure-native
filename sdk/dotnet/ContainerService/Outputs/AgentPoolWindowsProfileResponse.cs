// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Outputs
{

    /// <summary>
    /// The Windows agent pool's specific profile.
    /// </summary>
    [OutputType]
    public sealed class AgentPoolWindowsProfileResponse
    {
        /// <summary>
        /// The default value is false. Outbound NAT can only be disabled if the cluster outboundType is NAT Gateway and the Windows agent pool does not have node public IP enabled.
        /// </summary>
        public readonly bool? DisableOutboundNat;

        [OutputConstructor]
        private AgentPoolWindowsProfileResponse(bool? disableOutboundNat)
        {
            DisableOutboundNat = disableOutboundNat;
        }
    }
}
