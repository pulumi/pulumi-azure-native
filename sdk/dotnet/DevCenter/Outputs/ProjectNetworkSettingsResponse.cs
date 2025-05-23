// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevCenter.Outputs
{

    /// <summary>
    /// Network settings for the project.
    /// </summary>
    [OutputType]
    public sealed class ProjectNetworkSettingsResponse
    {
        /// <summary>
        /// Indicates whether pools in this Dev Center can use Microsoft Hosted Networks. Defaults to Enabled if not set.
        /// </summary>
        public readonly string MicrosoftHostedNetworkEnableStatus;

        [OutputConstructor]
        private ProjectNetworkSettingsResponse(string microsoftHostedNetworkEnableStatus)
        {
            MicrosoftHostedNetworkEnableStatus = microsoftHostedNetworkEnableStatus;
        }
    }
}
