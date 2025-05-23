// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Outputs
{

    /// <summary>
    /// Azure API management (APIM) configuration linked to the app.
    /// </summary>
    [OutputType]
    public sealed class ApiManagementConfigResponse
    {
        /// <summary>
        /// APIM-Api Identifier.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private ApiManagementConfigResponse(string? id)
        {
            Id = id;
        }
    }
}
