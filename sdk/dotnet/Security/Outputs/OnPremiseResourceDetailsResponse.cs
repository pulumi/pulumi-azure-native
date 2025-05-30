// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Outputs
{

    /// <summary>
    /// Details of the On Premise resource that was assessed
    /// </summary>
    [OutputType]
    public sealed class OnPremiseResourceDetailsResponse
    {
        /// <summary>
        /// The name of the machine
        /// </summary>
        public readonly string MachineName;
        /// <summary>
        /// The platform where the assessed resource resides
        /// Expected value is 'OnPremise'.
        /// </summary>
        public readonly string Source;
        /// <summary>
        /// The oms agent Id installed on the machine
        /// </summary>
        public readonly string SourceComputerId;
        /// <summary>
        /// The unique Id of the machine
        /// </summary>
        public readonly string Vmuuid;
        /// <summary>
        /// Azure resource Id of the workspace the machine is attached to
        /// </summary>
        public readonly string WorkspaceId;

        [OutputConstructor]
        private OnPremiseResourceDetailsResponse(
            string machineName,

            string source,

            string sourceComputerId,

            string vmuuid,

            string workspaceId)
        {
            MachineName = machineName;
            Source = source;
            SourceComputerId = sourceComputerId;
            Vmuuid = vmuuid;
            WorkspaceId = workspaceId;
        }
    }
}
