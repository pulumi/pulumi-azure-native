// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Inputs
{

    /// <summary>
    /// Managed service identity.
    /// </summary>
    public sealed class ManagedServiceIdentityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of managed service identity.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AzureNative.Web.ManagedServiceIdentityType>? Type { get; set; }

        [Input("userAssignedIdentities")]
        private InputList<string>? _userAssignedIdentities;

        /// <summary>
        /// The list of user assigned identities associated with the resource. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}
        /// </summary>
        public InputList<string> UserAssignedIdentities
        {
            get => _userAssignedIdentities ?? (_userAssignedIdentities = new InputList<string>());
            set => _userAssignedIdentities = value;
        }

        public ManagedServiceIdentityArgs()
        {
        }
        public static new ManagedServiceIdentityArgs Empty => new ManagedServiceIdentityArgs();
    }
}
