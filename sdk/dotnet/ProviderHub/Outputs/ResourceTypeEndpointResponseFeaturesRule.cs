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
    /// The features rule.
    /// </summary>
    [OutputType]
    public sealed class ResourceTypeEndpointResponseFeaturesRule
    {
        /// <summary>
        /// The required feature policy.
        /// </summary>
        public readonly string RequiredFeaturesPolicy;

        [OutputConstructor]
        private ResourceTypeEndpointResponseFeaturesRule(string requiredFeaturesPolicy)
        {
            RequiredFeaturesPolicy = requiredFeaturesPolicy;
        }
    }
}
