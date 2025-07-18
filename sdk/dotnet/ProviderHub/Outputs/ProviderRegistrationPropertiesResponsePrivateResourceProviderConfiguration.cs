// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Outputs
{

    /// <summary>
    /// The private resource provider configuration.
    /// </summary>
    [OutputType]
    public sealed class ProviderRegistrationPropertiesResponsePrivateResourceProviderConfiguration
    {
        /// <summary>
        /// The allowed subscriptions.
        /// </summary>
        public readonly ImmutableArray<string> AllowedSubscriptions;

        [OutputConstructor]
        private ProviderRegistrationPropertiesResponsePrivateResourceProviderConfiguration(ImmutableArray<string> allowedSubscriptions)
        {
            AllowedSubscriptions = allowedSubscriptions;
        }
    }
}
