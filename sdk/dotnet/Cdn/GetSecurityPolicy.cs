// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cdn
{
    public static class GetSecurityPolicy
    {
        /// <summary>
        /// Gets an existing security policy within a profile.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// 
        /// Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetSecurityPolicyResult> InvokeAsync(GetSecurityPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecurityPolicyResult>("azure-native:cdn:getSecurityPolicy", args ?? new GetSecurityPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an existing security policy within a profile.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// 
        /// Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSecurityPolicyResult> Invoke(GetSecurityPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityPolicyResult>("azure-native:cdn:getSecurityPolicy", args ?? new GetSecurityPolicyInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an existing security policy within a profile.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// 
        /// Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSecurityPolicyResult> Invoke(GetSecurityPolicyInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityPolicyResult>("azure-native:cdn:getSecurityPolicy", args ?? new GetSecurityPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecurityPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
        /// </summary>
        [Input("profileName", required: true)]
        public string ProfileName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the security policy under the profile.
        /// </summary>
        [Input("securityPolicyName", required: true)]
        public string SecurityPolicyName { get; set; } = null!;

        public GetSecurityPolicyArgs()
        {
        }
        public static new GetSecurityPolicyArgs Empty => new GetSecurityPolicyArgs();
    }

    public sealed class GetSecurityPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
        /// </summary>
        [Input("profileName", required: true)]
        public Input<string> ProfileName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the security policy under the profile.
        /// </summary>
        [Input("securityPolicyName", required: true)]
        public Input<string> SecurityPolicyName { get; set; } = null!;

        public GetSecurityPolicyInvokeArgs()
        {
        }
        public static new GetSecurityPolicyInvokeArgs Empty => new GetSecurityPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecurityPolicyResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        public readonly string DeploymentStatus;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// object which contains security policy parameters
        /// </summary>
        public readonly Outputs.SecurityPolicyWebApplicationFirewallParametersResponse? Parameters;
        /// <summary>
        /// The name of the profile which holds the security policy.
        /// </summary>
        public readonly string ProfileName;
        /// <summary>
        /// Provisioning status
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Read only system data
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSecurityPolicyResult(
            string azureApiVersion,

            string deploymentStatus,

            string id,

            string name,

            Outputs.SecurityPolicyWebApplicationFirewallParametersResponse? parameters,

            string profileName,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            DeploymentStatus = deploymentStatus;
            Id = id;
            Name = name;
            Parameters = parameters;
            ProfileName = profileName;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Type = type;
        }
    }
}
