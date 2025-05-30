// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridCompute
{
    public static class GetMachineRunCommand
    {
        /// <summary>
        /// The operation to get a run command.
        /// 
        /// Uses Azure REST API version 2024-07-31-preview.
        /// 
        /// Other available API versions: 2023-10-03-preview, 2024-03-31-preview, 2024-05-20-preview, 2024-09-10-preview, 2024-11-10-preview, 2025-01-13, 2025-02-19-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hybridcompute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetMachineRunCommandResult> InvokeAsync(GetMachineRunCommandArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMachineRunCommandResult>("azure-native:hybridcompute:getMachineRunCommand", args ?? new GetMachineRunCommandArgs(), options.WithDefaults());

        /// <summary>
        /// The operation to get a run command.
        /// 
        /// Uses Azure REST API version 2024-07-31-preview.
        /// 
        /// Other available API versions: 2023-10-03-preview, 2024-03-31-preview, 2024-05-20-preview, 2024-09-10-preview, 2024-11-10-preview, 2025-01-13, 2025-02-19-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hybridcompute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetMachineRunCommandResult> Invoke(GetMachineRunCommandInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMachineRunCommandResult>("azure-native:hybridcompute:getMachineRunCommand", args ?? new GetMachineRunCommandInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The operation to get a run command.
        /// 
        /// Uses Azure REST API version 2024-07-31-preview.
        /// 
        /// Other available API versions: 2023-10-03-preview, 2024-03-31-preview, 2024-05-20-preview, 2024-09-10-preview, 2024-11-10-preview, 2025-01-13, 2025-02-19-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hybridcompute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetMachineRunCommandResult> Invoke(GetMachineRunCommandInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMachineRunCommandResult>("azure-native:hybridcompute:getMachineRunCommand", args ?? new GetMachineRunCommandInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMachineRunCommandArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the hybrid machine.
        /// </summary>
        [Input("machineName", required: true)]
        public string MachineName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the run command.
        /// </summary>
        [Input("runCommandName", required: true)]
        public string RunCommandName { get; set; } = null!;

        public GetMachineRunCommandArgs()
        {
        }
        public static new GetMachineRunCommandArgs Empty => new GetMachineRunCommandArgs();
    }

    public sealed class GetMachineRunCommandInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the hybrid machine.
        /// </summary>
        [Input("machineName", required: true)]
        public Input<string> MachineName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the run command.
        /// </summary>
        [Input("runCommandName", required: true)]
        public Input<string> RunCommandName { get; set; } = null!;

        public GetMachineRunCommandInvokeArgs()
        {
        }
        public static new GetMachineRunCommandInvokeArgs Empty => new GetMachineRunCommandInvokeArgs();
    }


    [OutputType]
    public sealed class GetMachineRunCommandResult
    {
        /// <summary>
        /// Optional. If set to true, provisioning will complete as soon as script starts and will not wait for script to complete.
        /// </summary>
        public readonly bool? AsyncExecution;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// User-assigned managed identity that has access to errorBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged 
        /// </summary>
        public readonly Outputs.RunCommandManagedIdentityResponse? ErrorBlobManagedIdentity;
        /// <summary>
        /// Specifies the Azure storage blob where script error stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer errorBlobManagedIdentity parameter.
        /// </summary>
        public readonly string? ErrorBlobUri;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The machine run command instance view.
        /// </summary>
        public readonly Outputs.MachineRunCommandInstanceViewResponse InstanceView;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// User-assigned managed identity that has access to outputBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged 
        /// </summary>
        public readonly Outputs.RunCommandManagedIdentityResponse? OutputBlobManagedIdentity;
        /// <summary>
        /// Specifies the Azure storage blob where script output stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer outputBlobManagedIdentity parameter. 
        /// </summary>
        public readonly string? OutputBlobUri;
        /// <summary>
        /// The parameters used by the script.
        /// </summary>
        public readonly ImmutableArray<Outputs.RunCommandInputParameterResponse> Parameters;
        /// <summary>
        /// The parameters used by the script.
        /// </summary>
        public readonly ImmutableArray<Outputs.RunCommandInputParameterResponse> ProtectedParameters;
        /// <summary>
        /// The provisioning state, which only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Specifies the user account password on the machine when executing the run command.
        /// </summary>
        public readonly string? RunAsPassword;
        /// <summary>
        /// Specifies the user account on the machine when executing the run command.
        /// </summary>
        public readonly string? RunAsUser;
        /// <summary>
        /// The source of the run command script.
        /// </summary>
        public readonly Outputs.MachineRunCommandScriptSourceResponse? Source;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The timeout in seconds to execute the run command.
        /// </summary>
        public readonly int? TimeoutInSeconds;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetMachineRunCommandResult(
            bool? asyncExecution,

            string azureApiVersion,

            Outputs.RunCommandManagedIdentityResponse? errorBlobManagedIdentity,

            string? errorBlobUri,

            string id,

            Outputs.MachineRunCommandInstanceViewResponse instanceView,

            string location,

            string name,

            Outputs.RunCommandManagedIdentityResponse? outputBlobManagedIdentity,

            string? outputBlobUri,

            ImmutableArray<Outputs.RunCommandInputParameterResponse> parameters,

            ImmutableArray<Outputs.RunCommandInputParameterResponse> protectedParameters,

            string provisioningState,

            string? runAsPassword,

            string? runAsUser,

            Outputs.MachineRunCommandScriptSourceResponse? source,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            int? timeoutInSeconds,

            string type)
        {
            AsyncExecution = asyncExecution;
            AzureApiVersion = azureApiVersion;
            ErrorBlobManagedIdentity = errorBlobManagedIdentity;
            ErrorBlobUri = errorBlobUri;
            Id = id;
            InstanceView = instanceView;
            Location = location;
            Name = name;
            OutputBlobManagedIdentity = outputBlobManagedIdentity;
            OutputBlobUri = outputBlobUri;
            Parameters = parameters;
            ProtectedParameters = protectedParameters;
            ProvisioningState = provisioningState;
            RunAsPassword = runAsPassword;
            RunAsUser = runAsUser;
            Source = source;
            SystemData = systemData;
            Tags = tags;
            TimeoutInSeconds = timeoutInSeconds;
            Type = type;
        }
    }
}
