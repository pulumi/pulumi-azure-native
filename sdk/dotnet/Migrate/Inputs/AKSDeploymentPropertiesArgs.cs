// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.Inputs
{

    /// <summary>
    /// Class for AKSDeployment Properties.
    /// </summary>
    public sealed class AKSDeploymentPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the AKS cluster name.
        /// </summary>
        [Input("aksClusterName")]
        public Input<string>? AksClusterName { get; set; }

        /// <summary>
        /// Gets or sets the resource group of the resource.
        /// </summary>
        [Input("resourceGroup")]
        public Input<string>? ResourceGroup { get; set; }

        /// <summary>
        /// Gets or sets the subscription id of the resource.
        /// </summary>
        [Input("subscriptionId")]
        public Input<string>? SubscriptionId { get; set; }

        /// <summary>
        /// Gets or sets the tenant id.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public AKSDeploymentPropertiesArgs()
        {
        }
        public static new AKSDeploymentPropertiesArgs Empty => new AKSDeploymentPropertiesArgs();
    }
}
