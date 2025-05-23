// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.Inputs
{

    /// <summary>
    /// Enable migration input properties.
    /// </summary>
    public sealed class EnableMigrationInputPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy Id.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        /// <summary>
        /// The provider specific details.
        /// </summary>
        [Input("providerSpecificDetails", required: true)]
        public Input<Inputs.VMwareCbtEnableMigrationInputArgs> ProviderSpecificDetails { get; set; } = null!;

        public EnableMigrationInputPropertiesArgs()
        {
        }
        public static new EnableMigrationInputPropertiesArgs Empty => new EnableMigrationInputPropertiesArgs();
    }
}
