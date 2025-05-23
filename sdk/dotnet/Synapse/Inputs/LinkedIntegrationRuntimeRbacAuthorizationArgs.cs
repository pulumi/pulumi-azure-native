// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Synapse.Inputs
{

    /// <summary>
    /// The role based access control (RBAC) authorization type integration runtime.
    /// </summary>
    public sealed class LinkedIntegrationRuntimeRbacAuthorizationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authorization type for integration runtime sharing.
        /// Expected value is 'RBAC'.
        /// </summary>
        [Input("authorizationType", required: true)]
        public Input<string> AuthorizationType { get; set; } = null!;

        /// <summary>
        /// The resource identifier of the integration runtime to be shared.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        public LinkedIntegrationRuntimeRbacAuthorizationArgs()
        {
        }
        public static new LinkedIntegrationRuntimeRbacAuthorizationArgs Empty => new LinkedIntegrationRuntimeRbacAuthorizationArgs();
    }
}
