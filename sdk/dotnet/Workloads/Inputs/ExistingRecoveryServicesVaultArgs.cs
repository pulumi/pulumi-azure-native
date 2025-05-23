// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Workloads.Inputs
{

    /// <summary>
    /// Existing recovery services vault.
    /// </summary>
    public sealed class ExistingRecoveryServicesVaultArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID of the recovery services vault that has been created.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The vault type, whether it is existing or has to be created.
        /// Expected value is 'Existing'.
        /// </summary>
        [Input("vaultType", required: true)]
        public Input<string> VaultType { get; set; } = null!;

        public ExistingRecoveryServicesVaultArgs()
        {
        }
        public static new ExistingRecoveryServicesVaultArgs Empty => new ExistingRecoveryServicesVaultArgs();
    }
}
