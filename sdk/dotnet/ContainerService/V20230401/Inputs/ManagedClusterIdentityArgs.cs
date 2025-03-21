// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20230401.Inputs
{

    /// <summary>
    /// Identity for the managed cluster.
    /// </summary>
    public sealed class ManagedClusterIdentityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// For more information see [use managed identities in AKS](https://docs.microsoft.com/azure/aks/use-managed-identity).
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AzureNative.ContainerService.V20230401.ResourceIdentityType>? Type { get; set; }

        [Input("userAssignedIdentities")]
        private InputList<string>? _userAssignedIdentities;

        /// <summary>
        /// The keys must be ARM resource IDs in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
        /// </summary>
        public InputList<string> UserAssignedIdentities
        {
            get => _userAssignedIdentities ?? (_userAssignedIdentities = new InputList<string>());
            set => _userAssignedIdentities = value;
        }

        public ManagedClusterIdentityArgs()
        {
        }
        public static new ManagedClusterIdentityArgs Empty => new ManagedClusterIdentityArgs();
    }
}
