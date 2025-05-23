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
    public sealed class ResourceTypeRegistrationPropertiesResponseResourceGraphConfiguration
    {
        public readonly string? ApiVersion;
        public readonly bool? Enabled;

        [OutputConstructor]
        private ResourceTypeRegistrationPropertiesResponseResourceGraphConfiguration(
            string? apiVersion,

            bool? enabled)
        {
            ApiVersion = apiVersion;
            Enabled = enabled;
        }
    }
}
