// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Outputs
{

    /// <summary>
    /// Configuration properties Keda component
    /// </summary>
    [OutputType]
    public sealed class KedaConfigurationResponse
    {
        /// <summary>
        /// The version of Keda
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private KedaConfigurationResponse(string version)
        {
            Version = version;
        }
    }
}
