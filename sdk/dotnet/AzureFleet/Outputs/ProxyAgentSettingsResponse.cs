// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureFleet.Outputs
{

    /// <summary>
    /// Specifies ProxyAgent settings while creating the virtual machine. Minimum
    /// api-version: 2023-09-01.
    /// </summary>
    [OutputType]
    public sealed class ProxyAgentSettingsResponse
    {
        /// <summary>
        /// Specifies whether ProxyAgent feature should be enabled on the virtual machine
        /// or virtual machine scale set.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Increase the value of this property allows user to reset the key used for
        /// securing communication channel between guest and host.
        /// </summary>
        public readonly int? KeyIncarnationId;
        /// <summary>
        /// Specifies the mode that ProxyAgent will execute on if the feature is enabled.
        /// ProxyAgent will start to audit or monitor but not enforce access control over
        /// requests to host endpoints in Audit mode, while in Enforce mode it will enforce
        /// access control. The default value is Enforce mode.
        /// </summary>
        public readonly string? Mode;

        [OutputConstructor]
        private ProxyAgentSettingsResponse(
            bool? enabled,

            int? keyIncarnationId,

            string? mode)
        {
            Enabled = enabled;
            KeyIncarnationId = keyIncarnationId;
            Mode = mode;
        }
    }
}
