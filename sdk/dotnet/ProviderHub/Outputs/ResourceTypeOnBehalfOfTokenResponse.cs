// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Outputs
{

    [OutputType]
    public sealed class ResourceTypeOnBehalfOfTokenResponse
    {
        /// <summary>
        /// The action name.
        /// </summary>
        public readonly string? ActionName;
        /// <summary>
        /// This is a TimeSpan property.
        /// </summary>
        public readonly string? LifeTime;

        [OutputConstructor]
        private ResourceTypeOnBehalfOfTokenResponse(
            string? actionName,

            string? lifeTime)
        {
            ActionName = actionName;
            LifeTime = lifeTime;
        }
    }
}
