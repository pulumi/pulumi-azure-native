// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Storage.Outputs
{

    /// <summary>
    /// Target helps provide filter parameters for the objects in the storage account and forms the execution context for the storage task
    /// </summary>
    [OutputType]
    public sealed class ExecutionTargetResponse
    {
        /// <summary>
        /// List of object prefixes to be excluded from task execution. If there is a conflict between include and exclude prefixes, the exclude prefix will be the determining factor
        /// </summary>
        public readonly ImmutableArray<string> ExcludePrefix;
        /// <summary>
        /// Required list of object prefixes to be included for task execution
        /// </summary>
        public readonly ImmutableArray<string> Prefix;

        [OutputConstructor]
        private ExecutionTargetResponse(
            ImmutableArray<string> excludePrefix,

            ImmutableArray<string> prefix)
        {
            ExcludePrefix = excludePrefix;
            Prefix = prefix;
        }
    }
}
