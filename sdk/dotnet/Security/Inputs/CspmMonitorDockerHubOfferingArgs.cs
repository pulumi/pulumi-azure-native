// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Inputs
{

    /// <summary>
    /// The CSPM (Cloud security posture management) monitoring for Docker Hub offering
    /// </summary>
    public sealed class CspmMonitorDockerHubOfferingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of the security offering.
        /// Expected value is 'CspmMonitorDockerHub'.
        /// </summary>
        [Input("offeringType", required: true)]
        public Input<string> OfferingType { get; set; } = null!;

        public CspmMonitorDockerHubOfferingArgs()
        {
        }
        public static new CspmMonitorDockerHubOfferingArgs Empty => new CspmMonitorDockerHubOfferingArgs();
    }
}
