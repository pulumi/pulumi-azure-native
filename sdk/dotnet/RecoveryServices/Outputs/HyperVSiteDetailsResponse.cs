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
    /// HyperVSite fabric specific details.
    /// </summary>
    [OutputType]
    public sealed class HyperVSiteDetailsResponse
    {
        /// <summary>
        /// The list of Hyper-V hosts associated with the fabric.
        /// </summary>
        public readonly ImmutableArray<Outputs.HyperVHostDetailsResponse> HyperVHosts;
        /// <summary>
        /// Gets the class type. Overridden in derived classes.
        /// Expected value is 'HyperVSite'.
        /// </summary>
        public readonly string InstanceType;

        [OutputConstructor]
        private HyperVSiteDetailsResponse(
            ImmutableArray<Outputs.HyperVHostDetailsResponse> hyperVHosts,

            string instanceType)
        {
            HyperVHosts = hyperVHosts;
            InstanceType = instanceType;
        }
    }
}
