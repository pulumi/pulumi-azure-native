// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManagedServices.Inputs
{

    /// <summary>
    /// The Azure Active Directory principal identifier, Azure built-in role, and just-in-time access policy that describes the just-in-time access the principal will receive on the delegated resource in the managed tenant.
    /// </summary>
    public sealed class EligibleAuthorizationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The just-in-time access policy setting.
        /// </summary>
        [Input("justInTimeAccessPolicy")]
        public Input<Inputs.JustInTimeAccessPolicyArgs>? JustInTimeAccessPolicy { get; set; }

        /// <summary>
        /// The identifier of the Azure Active Directory principal.
        /// </summary>
        [Input("principalId", required: true)]
        public Input<string> PrincipalId { get; set; } = null!;

        /// <summary>
        /// The display name of the Azure Active Directory principal.
        /// </summary>
        [Input("principalIdDisplayName")]
        public Input<string>? PrincipalIdDisplayName { get; set; }

        /// <summary>
        /// The identifier of the Azure built-in role that defines the permissions that the Azure Active Directory principal will have on the projected scope.
        /// </summary>
        [Input("roleDefinitionId", required: true)]
        public Input<string> RoleDefinitionId { get; set; } = null!;

        public EligibleAuthorizationArgs()
        {
        }
        public static new EligibleAuthorizationArgs Empty => new EligibleAuthorizationArgs();
    }
}
