// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DocumentDB.V20231115.Outputs
{

    /// <summary>
    /// Parameters to indicate the information about the restore.
    /// </summary>
    [OutputType]
    public sealed class ResourceRestoreParametersResponse
    {
        /// <summary>
        /// The id of the restorable database account from which the restore has to be initiated. For example: /subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{restorableDatabaseAccountName}
        /// </summary>
        public readonly string? RestoreSource;
        /// <summary>
        /// Time to which the account has to be restored (ISO-8601 format).
        /// </summary>
        public readonly string? RestoreTimestampInUtc;

        [OutputConstructor]
        private ResourceRestoreParametersResponse(
            string? restoreSource,

            string? restoreTimestampInUtc)
        {
            RestoreSource = restoreSource;
            RestoreTimestampInUtc = restoreTimestampInUtc;
        }
    }
}
