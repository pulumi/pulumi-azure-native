// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevTestLab.Inputs
{

    /// <summary>
    /// Properties of a managed identity
    /// </summary>
    public sealed class IdentityPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The client secret URL of the identity.
        /// </summary>
        [Input("clientSecretUrl")]
        public Input<string>? ClientSecretUrl { get; set; }

        /// <summary>
        /// The principal id of resource identity.
        /// </summary>
        [Input("principalId")]
        public Input<string>? PrincipalId { get; set; }

        /// <summary>
        /// The tenant identifier of resource.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// Managed identity.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNative.DevTestLab.ManagedIdentityType>? Type { get; set; }

        public IdentityPropertiesArgs()
        {
        }
        public static new IdentityPropertiesArgs Empty => new IdentityPropertiesArgs();
    }
}
