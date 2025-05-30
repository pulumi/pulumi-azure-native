// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.KeyVault.Outputs
{

    [OutputType]
    public sealed class TriggerResponse
    {
        /// <summary>
        /// The time duration after key creation to rotate the key. It only applies to rotate. It will be in ISO 8601 duration format. Eg: 'P90D', 'P1Y'.
        /// </summary>
        public readonly string? TimeAfterCreate;
        /// <summary>
        /// The time duration before key expiring to rotate or notify. It will be in ISO 8601 duration format. Eg: 'P90D', 'P1Y'.
        /// </summary>
        public readonly string? TimeBeforeExpiry;

        [OutputConstructor]
        private TriggerResponse(
            string? timeAfterCreate,

            string? timeBeforeExpiry)
        {
            TimeAfterCreate = timeAfterCreate;
            TimeBeforeExpiry = timeBeforeExpiry;
        }
    }
}
