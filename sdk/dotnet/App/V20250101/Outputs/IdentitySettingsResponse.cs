// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.V20250101.Outputs
{

    /// <summary>
    /// Optional settings for a Managed Identity that is assigned to the Container App.
    /// </summary>
    [OutputType]
    public sealed class IdentitySettingsResponse
    {
        /// <summary>
        /// The resource ID of a user-assigned managed identity that is assigned to the Container App, or 'system' for system-assigned identity.
        /// </summary>
        public readonly string Identity;
        /// <summary>
        /// Use to select the lifecycle stages of a Container App during which the Managed Identity should be available.
        /// </summary>
        public readonly string? Lifecycle;

        [OutputConstructor]
        private IdentitySettingsResponse(
            string identity,

            string? lifecycle)
        {
            Identity = identity;
            Lifecycle = lifecycle;
        }
    }
}
