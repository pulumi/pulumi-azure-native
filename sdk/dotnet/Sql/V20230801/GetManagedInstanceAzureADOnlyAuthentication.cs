// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql.V20230801
{
    public static class GetManagedInstanceAzureADOnlyAuthentication
    {
        /// <summary>
        /// Gets a specific Azure Active Directory only authentication property.
        /// </summary>
        public static Task<GetManagedInstanceAzureADOnlyAuthenticationResult> InvokeAsync(GetManagedInstanceAzureADOnlyAuthenticationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetManagedInstanceAzureADOnlyAuthenticationResult>("azure-native:sql/v20230801:getManagedInstanceAzureADOnlyAuthentication", args ?? new GetManagedInstanceAzureADOnlyAuthenticationArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a specific Azure Active Directory only authentication property.
        /// </summary>
        public static Output<GetManagedInstanceAzureADOnlyAuthenticationResult> Invoke(GetManagedInstanceAzureADOnlyAuthenticationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetManagedInstanceAzureADOnlyAuthenticationResult>("azure-native:sql/v20230801:getManagedInstanceAzureADOnlyAuthentication", args ?? new GetManagedInstanceAzureADOnlyAuthenticationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a specific Azure Active Directory only authentication property.
        /// </summary>
        public static Output<GetManagedInstanceAzureADOnlyAuthenticationResult> Invoke(GetManagedInstanceAzureADOnlyAuthenticationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetManagedInstanceAzureADOnlyAuthenticationResult>("azure-native:sql/v20230801:getManagedInstanceAzureADOnlyAuthentication", args ?? new GetManagedInstanceAzureADOnlyAuthenticationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetManagedInstanceAzureADOnlyAuthenticationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of server azure active directory only authentication.
        /// </summary>
        [Input("authenticationName", required: true)]
        public string AuthenticationName { get; set; } = null!;

        /// <summary>
        /// The name of the managed instance.
        /// </summary>
        [Input("managedInstanceName", required: true)]
        public string ManagedInstanceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetManagedInstanceAzureADOnlyAuthenticationArgs()
        {
        }
        public static new GetManagedInstanceAzureADOnlyAuthenticationArgs Empty => new GetManagedInstanceAzureADOnlyAuthenticationArgs();
    }

    public sealed class GetManagedInstanceAzureADOnlyAuthenticationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of server azure active directory only authentication.
        /// </summary>
        [Input("authenticationName", required: true)]
        public Input<string> AuthenticationName { get; set; } = null!;

        /// <summary>
        /// The name of the managed instance.
        /// </summary>
        [Input("managedInstanceName", required: true)]
        public Input<string> ManagedInstanceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetManagedInstanceAzureADOnlyAuthenticationInvokeArgs()
        {
        }
        public static new GetManagedInstanceAzureADOnlyAuthenticationInvokeArgs Empty => new GetManagedInstanceAzureADOnlyAuthenticationInvokeArgs();
    }


    [OutputType]
    public sealed class GetManagedInstanceAzureADOnlyAuthenticationResult
    {
        /// <summary>
        /// Azure Active Directory only Authentication enabled.
        /// </summary>
        public readonly bool AzureADOnlyAuthentication;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetManagedInstanceAzureADOnlyAuthenticationResult(
            bool azureADOnlyAuthentication,

            string id,

            string name,

            string type)
        {
            AzureADOnlyAuthentication = azureADOnlyAuthentication;
            Id = id;
            Name = name;
            Type = type;
        }
    }
}
