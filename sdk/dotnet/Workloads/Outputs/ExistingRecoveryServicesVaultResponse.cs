// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Workloads.Outputs
{

    /// <summary>
    /// Existing recovery services vault.
    /// </summary>
    [OutputType]
    public sealed class ExistingRecoveryServicesVaultResponse
    {
        /// <summary>
        /// The resource ID of the recovery services vault that has been created.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The vault type, whether it is existing or has to be created.
        /// Expected value is 'Existing'.
        /// </summary>
        public readonly string VaultType;

        [OutputConstructor]
        private ExistingRecoveryServicesVaultResponse(
            string id,

            string vaultType)
        {
            Id = id;
            VaultType = vaultType;
        }
    }
}
