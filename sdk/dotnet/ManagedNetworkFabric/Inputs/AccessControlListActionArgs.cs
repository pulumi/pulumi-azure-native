// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManagedNetworkFabric.Inputs
{

    /// <summary>
    /// Action that need to performed.
    /// </summary>
    public sealed class AccessControlListActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the counter block to get match count information.
        /// </summary>
        [Input("counterName")]
        public Input<string>? CounterName { get; set; }

        /// <summary>
        /// Type of actions that can be performed.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNative.ManagedNetworkFabric.AclActionType>? Type { get; set; }

        public AccessControlListActionArgs()
        {
        }
        public static new AccessControlListActionArgs Empty => new AccessControlListActionArgs();
    }
}
