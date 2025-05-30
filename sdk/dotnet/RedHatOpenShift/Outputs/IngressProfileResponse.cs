// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RedHatOpenShift.Outputs
{

    /// <summary>
    /// IngressProfile represents an ingress profile.
    /// </summary>
    [OutputType]
    public sealed class IngressProfileResponse
    {
        /// <summary>
        /// The IP of the ingress.
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// The ingress profile name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Ingress visibility.
        /// </summary>
        public readonly string? Visibility;

        [OutputConstructor]
        private IngressProfileResponse(
            string ip,

            string? name,

            string? visibility)
        {
            Ip = ip;
            Name = name;
            Visibility = visibility;
        }
    }
}
