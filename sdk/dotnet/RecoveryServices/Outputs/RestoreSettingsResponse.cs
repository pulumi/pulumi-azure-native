// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.Outputs
{

    /// <summary>
    /// Restore Settings  of the vault
    /// </summary>
    [OutputType]
    public sealed class RestoreSettingsResponse
    {
        /// <summary>
        /// Settings for CrossSubscriptionRestore
        /// </summary>
        public readonly Outputs.CrossSubscriptionRestoreSettingsResponse? CrossSubscriptionRestoreSettings;

        [OutputConstructor]
        private RestoreSettingsResponse(Outputs.CrossSubscriptionRestoreSettingsResponse? crossSubscriptionRestoreSettings)
        {
            CrossSubscriptionRestoreSettings = crossSubscriptionRestoreSettings;
        }
    }
}
