// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cdn.Outputs
{

    /// <summary>
    /// Defines the parameters for the Url Signing action.
    /// </summary>
    [OutputType]
    public sealed class UrlSigningActionParametersResponse
    {
        /// <summary>
        /// Algorithm to use for URL signing
        /// </summary>
        public readonly string? Algorithm;
        /// <summary>
        /// Defines which query string parameters in the url to be considered for expires, key id etc. 
        /// </summary>
        public readonly ImmutableArray<Outputs.UrlSigningParamIdentifierResponse> ParameterNameOverride;
        /// <summary>
        /// 
        /// Expected value is 'DeliveryRuleUrlSigningActionParameters'.
        /// </summary>
        public readonly string TypeName;

        [OutputConstructor]
        private UrlSigningActionParametersResponse(
            string? algorithm,

            ImmutableArray<Outputs.UrlSigningParamIdentifierResponse> parameterNameOverride,

            string typeName)
        {
            Algorithm = algorithm;
            ParameterNameOverride = parameterNameOverride;
            TypeName = typeName;
        }
    }
}
