// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStack
{
    public static class GetRegistrationActivationKey
    {
        /// <summary>
        /// Returns Azure Stack Activation Key.
        /// 
        /// Uses Azure REST API version 2022-06-01.
        /// 
        /// Other available API versions: 2020-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurestack [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetRegistrationActivationKeyResult> InvokeAsync(GetRegistrationActivationKeyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRegistrationActivationKeyResult>("azure-native:azurestack:getRegistrationActivationKey", args ?? new GetRegistrationActivationKeyArgs(), options.WithDefaults());

        /// <summary>
        /// Returns Azure Stack Activation Key.
        /// 
        /// Uses Azure REST API version 2022-06-01.
        /// 
        /// Other available API versions: 2020-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurestack [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetRegistrationActivationKeyResult> Invoke(GetRegistrationActivationKeyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRegistrationActivationKeyResult>("azure-native:azurestack:getRegistrationActivationKey", args ?? new GetRegistrationActivationKeyInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Returns Azure Stack Activation Key.
        /// 
        /// Uses Azure REST API version 2022-06-01.
        /// 
        /// Other available API versions: 2020-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurestack [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetRegistrationActivationKeyResult> Invoke(GetRegistrationActivationKeyInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRegistrationActivationKeyResult>("azure-native:azurestack:getRegistrationActivationKey", args ?? new GetRegistrationActivationKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegistrationActivationKeyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Azure Stack registration.
        /// </summary>
        [Input("registrationName", required: true)]
        public string RegistrationName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group.
        /// </summary>
        [Input("resourceGroup", required: true)]
        public string ResourceGroup { get; set; } = null!;

        public GetRegistrationActivationKeyArgs()
        {
        }
        public static new GetRegistrationActivationKeyArgs Empty => new GetRegistrationActivationKeyArgs();
    }

    public sealed class GetRegistrationActivationKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Azure Stack registration.
        /// </summary>
        [Input("registrationName", required: true)]
        public Input<string> RegistrationName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group.
        /// </summary>
        [Input("resourceGroup", required: true)]
        public Input<string> ResourceGroup { get; set; } = null!;

        public GetRegistrationActivationKeyInvokeArgs()
        {
        }
        public static new GetRegistrationActivationKeyInvokeArgs Empty => new GetRegistrationActivationKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetRegistrationActivationKeyResult
    {
        /// <summary>
        /// Azure Stack activation key.
        /// </summary>
        public readonly string? ActivationKey;

        [OutputConstructor]
        private GetRegistrationActivationKeyResult(string? activationKey)
        {
            ActivationKey = activationKey;
        }
    }
}
