// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Outputs
{

    /// <summary>
    /// Azure ML WebService Input/Output file
    /// </summary>
    [OutputType]
    public sealed class AzureMLWebServiceFileResponse
    {
        /// <summary>
        /// The relative file path, including container name, in the Azure Blob Storage specified by the LinkedService. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object FilePath;
        /// <summary>
        /// Reference to an Azure Storage LinkedService, where Azure ML WebService Input/Output file located.
        /// </summary>
        public readonly Outputs.LinkedServiceReferenceResponse LinkedServiceName;

        [OutputConstructor]
        private AzureMLWebServiceFileResponse(
            object filePath,

            Outputs.LinkedServiceReferenceResponse linkedServiceName)
        {
            FilePath = filePath;
            LinkedServiceName = linkedServiceName;
        }
    }
}
