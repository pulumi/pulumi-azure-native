// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid.V20241215Preview.Inputs
{

    /// <summary>
    /// The identity information for retrieving the certificate for custom JWT authentication.
    /// </summary>
    public sealed class CustomJwtAuthenticationManagedIdentityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of managed identity used. Can be either 'SystemAssigned' or 'UserAssigned'.
        /// </summary>
        [Input("type", required: true)]
        public InputUnion<string, Pulumi.AzureNative.EventGrid.V20241215Preview.CustomJwtAuthenticationManagedIdentityType> Type { get; set; } = null!;

        /// <summary>
        /// The user identity associated with the resource.
        /// </summary>
        [Input("userAssignedIdentity")]
        public Input<string>? UserAssignedIdentity { get; set; }

        public CustomJwtAuthenticationManagedIdentityArgs()
        {
        }
        public static new CustomJwtAuthenticationManagedIdentityArgs Empty => new CustomJwtAuthenticationManagedIdentityArgs();
    }
}
