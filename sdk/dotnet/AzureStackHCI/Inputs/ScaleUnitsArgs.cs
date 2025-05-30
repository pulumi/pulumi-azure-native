// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.Inputs
{

    /// <summary>
    /// Scale units will contains list of deployment data
    /// </summary>
    public sealed class ScaleUnitsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Deployment Data to deploy AzureStackHCI Cluster.
        /// </summary>
        [Input("deploymentData", required: true)]
        public Input<Inputs.DeploymentDataArgs> DeploymentData { get; set; } = null!;

        /// <summary>
        /// Solution builder extension (SBE) partner properties
        /// </summary>
        [Input("sbePartnerInfo")]
        public Input<Inputs.SbePartnerInfoArgs>? SbePartnerInfo { get; set; }

        public ScaleUnitsArgs()
        {
        }
        public static new ScaleUnitsArgs Empty => new ScaleUnitsArgs();
    }
}
