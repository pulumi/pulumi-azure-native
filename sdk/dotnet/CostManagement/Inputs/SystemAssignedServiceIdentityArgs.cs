// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CostManagement.Inputs
{

    /// <summary>
    /// Managed service identity (either system assigned, or none)
    /// </summary>
    public sealed class SystemAssignedServiceIdentityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of managed service identity (either system assigned, or none).
        /// </summary>
        [Input("type", required: true)]
        public InputUnion<string, Pulumi.AzureNative.CostManagement.SystemAssignedServiceIdentityType> Type { get; set; } = null!;

        public SystemAssignedServiceIdentityArgs()
        {
        }
        public static new SystemAssignedServiceIdentityArgs Empty => new SystemAssignedServiceIdentityArgs();
    }
}
