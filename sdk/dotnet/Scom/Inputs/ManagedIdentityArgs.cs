// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Scom.Inputs
{

    /// <summary>
    /// Azure Active Directory identity configuration for a resource.
    /// </summary>
    public sealed class ManagedIdentityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identity type
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNative.Scom.ManagedIdentityType>? Type { get; set; }

        [Input("userAssignedIdentities")]
        private InputList<string>? _userAssignedIdentities;

        /// <summary>
        /// The resource ids of the user assigned identities to use
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
