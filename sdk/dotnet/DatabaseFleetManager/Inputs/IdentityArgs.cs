// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DatabaseFleetManager.Inputs
{

    /// <summary>
    /// Database Identity.
    /// </summary>
    public sealed class IdentityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The federated client id for the SQL Database. It is used for cross tenant CMK scenario.
        /// </summary>
        [Input("federatedClientId")]
        public Input<string>? FederatedClientId { get; set; }

        /// <summary>
        /// Identity type of the main principal.
        /// </summary>
        [Input("identityType")]
        public InputUnion<string, Pulumi.AzureNative.DatabaseFleetManager.IdentityType>? IdentityType { get; set; }

        [Input("userAssignedIdentities")]
        private InputList<Inputs.DatabaseIdentityArgs>? _userAssignedIdentities;

        /// <summary>
        /// User identity ids
        /// </summary>
        public InputList<Inputs.DatabaseIdentityArgs> UserAssignedIdentities
        {
            get => _userAssignedIdentities ?? (_userAssignedIdentities = new InputList<Inputs.DatabaseIdentityArgs>());
            set => _userAssignedIdentities = value;
        }

        public IdentityArgs()
        {
        }
        public static new IdentityArgs Empty => new IdentityArgs();
    }
}
