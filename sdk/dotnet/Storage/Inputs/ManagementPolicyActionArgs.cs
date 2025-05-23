// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Storage.Inputs
{

    /// <summary>
    /// Actions are applied to the filtered blobs when the execution condition is met.
    /// </summary>
    public sealed class ManagementPolicyActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The management policy action for base blob
        /// </summary>
        [Input("baseBlob")]
        public Input<Inputs.ManagementPolicyBaseBlobArgs>? BaseBlob { get; set; }

        /// <summary>
        /// The management policy action for snapshot
        /// </summary>
        [Input("snapshot")]
        public Input<Inputs.ManagementPolicySnapShotArgs>? Snapshot { get; set; }

        /// <summary>
        /// The management policy action for version
        /// </summary>
        [Input("version")]
        public Input<Inputs.ManagementPolicyVersionArgs>? Version { get; set; }

        public ManagementPolicyActionArgs()
        {
        }
        public static new ManagementPolicyActionArgs Empty => new ManagementPolicyActionArgs();
    }
}
