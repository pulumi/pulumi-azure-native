// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Synapse.V20210601Preview.Inputs
{

    /// <summary>
    /// The workspace managed identity
    /// </summary>
    public sealed class ManagedIdentityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of managed identity for the workspace
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AzureNative.Synapse.V20210601Preview.ResourceIdentityType>? Type { get; set; }

        [Input("userAssignedIdentities")]
        private InputList<string>? _userAssignedIdentities;

        /// <summary>
        /// The user assigned managed identities.
        /// </summary>
        public InputList<string> UserAssignedIdentities
        {
            get => _userAssignedIdentities ?? (_userAssignedIdentities = new InputList<string>());
            set => _userAssignedIdentities = value;
        }

        public ManagedIdentityArgs()
        {
        }
        public static new ManagedIdentityArgs Empty => new ManagedIdentityArgs();
    }
}
