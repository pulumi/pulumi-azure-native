// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DBforPostgreSQL.V20230301Preview.Outputs
{

    /// <summary>
    /// Migration sub state details.
    /// </summary>
    [OutputType]
    public sealed class MigrationSubStateDetailsResponse
    {
        /// <summary>
        /// Migration sub state.
        /// </summary>
        public readonly string CurrentSubState;

        [OutputConstructor]
        private MigrationSubStateDetailsResponse(string currentSubState)
        {
            CurrentSubState = currentSubState;
        }
    }
}
