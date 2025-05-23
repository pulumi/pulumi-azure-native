// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperations.Outputs
{

    /// <summary>
    /// State Store Resource Rule properties.
    /// </summary>
    [OutputType]
    public sealed class StateStoreResourceRuleResponse
    {
        /// <summary>
        /// Allowed keyTypes pattern, string, binary. The key type used for matching, for example pattern tries to match the key to a glob-style pattern and string checks key is equal to value provided in keys.
        /// </summary>
        public readonly string KeyType;
        /// <summary>
        /// Give access to state store keys for the corresponding principals defined. When key type is pattern set glob-style pattern (e.g., '*', 'clients/*').
        /// </summary>
        public readonly ImmutableArray<string> Keys;
        /// <summary>
        /// Give access for `Read`, `Write` and `ReadWrite` access level.
        /// </summary>
        public readonly string Method;

        [OutputConstructor]
        private StateStoreResourceRuleResponse(
            string keyType,

            ImmutableArray<string> keys,

            string method)
        {
            KeyType = keyType;
            Keys = keys;
            Method = method;
        }
    }
}
