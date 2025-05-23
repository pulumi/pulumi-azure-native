// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Storage.Inputs
{

    /// <summary>
    /// Resource Access Rule.
    /// </summary>
    public sealed class ResourceAccessRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource Id
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// Tenant Id
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public ResourceAccessRuleArgs()
        {
        }
        public static new ResourceAccessRuleArgs Empty => new ResourceAccessRuleArgs();
    }
}
