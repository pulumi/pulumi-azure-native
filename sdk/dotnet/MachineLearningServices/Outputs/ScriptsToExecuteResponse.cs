// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Outputs
{

    /// <summary>
    /// Customized setup scripts
    /// </summary>
    [OutputType]
    public sealed class ScriptsToExecuteResponse
    {
        /// <summary>
        /// Script that's run only once during provision of the compute.
        /// </summary>
        public readonly Outputs.ScriptReferenceResponse? CreationScript;
        /// <summary>
        /// Script that's run every time the machine starts.
        /// </summary>
        public readonly Outputs.ScriptReferenceResponse? StartupScript;

        [OutputConstructor]
        private ScriptsToExecuteResponse(
            Outputs.ScriptReferenceResponse? creationScript,

            Outputs.ScriptReferenceResponse? startupScript)
        {
            CreationScript = creationScript;
            StartupScript = startupScript;
        }
    }
}
