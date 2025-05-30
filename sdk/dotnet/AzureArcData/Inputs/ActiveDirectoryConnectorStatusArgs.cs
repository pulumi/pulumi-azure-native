// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureArcData.Inputs
{

    /// <summary>
    /// The status of the Kubernetes custom resource.
    /// </summary>
    public sealed class ActiveDirectoryConnectorStatusArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time that the custom resource was last updated.
        /// </summary>
        [Input("lastUpdateTime")]
        public Input<string>? LastUpdateTime { get; set; }

        /// <summary>
        /// The version of the replicaSet associated with the AD connector custom resource.
        /// </summary>
        [Input("observedGeneration")]
        public Input<double>? ObservedGeneration { get; set; }

        /// <summary>
        /// The state of the AD connector custom resource.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public ActiveDirectoryConnectorStatusArgs()
        {
        }
        public static new ActiveDirectoryConnectorStatusArgs Empty => new ActiveDirectoryConnectorStatusArgs();
    }
}
