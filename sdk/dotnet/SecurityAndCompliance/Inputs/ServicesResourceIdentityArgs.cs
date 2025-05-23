// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityAndCompliance.Inputs
{

    /// <summary>
    /// Setting indicating whether the service has a managed identity associated with it.
    /// </summary>
    public sealed class ServicesResourceIdentityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of identity being specified, currently SystemAssigned and None are allowed.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNative.SecurityAndCompliance.ManagedServiceIdentityType>? Type { get; set; }

        public ServicesResourceIdentityArgs()
        {
        }
        public static new ServicesResourceIdentityArgs Empty => new ServicesResourceIdentityArgs();
    }
}
