# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['AzurePowerShellScriptArgs', 'AzurePowerShellScript']

@pulumi.input_type
class AzurePowerShellScriptArgs:
    def __init__(__self__, *,
                 az_power_shell_version: pulumi.Input[builtins.str],
                 kind: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 retention_interval: pulumi.Input[builtins.str],
                 arguments: Optional[pulumi.Input[builtins.str]] = None,
                 cleanup_preference: Optional[pulumi.Input[Union[builtins.str, 'CleanupOptions']]] = None,
                 container_settings: Optional[pulumi.Input['ContainerConfigurationArgs']] = None,
                 environment_variables: Optional[pulumi.Input[Sequence[pulumi.Input['EnvironmentVariableArgs']]]] = None,
                 force_update_tag: Optional[pulumi.Input[builtins.str]] = None,
                 identity: Optional[pulumi.Input['ManagedServiceIdentityArgs']] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 primary_script_uri: Optional[pulumi.Input[builtins.str]] = None,
                 script_content: Optional[pulumi.Input[builtins.str]] = None,
                 script_name: Optional[pulumi.Input[builtins.str]] = None,
                 storage_account_settings: Optional[pulumi.Input['StorageAccountConfigurationArgs']] = None,
                 supporting_script_uris: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 timeout: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a AzurePowerShellScript resource.
        :param pulumi.Input[builtins.str] az_power_shell_version: Azure PowerShell module version to be used.
        :param pulumi.Input[builtins.str] kind: Type of the script.
               Expected value is 'AzurePowerShell'.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] retention_interval: Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. Duration is based on ISO 8601 pattern (for example P1D means one day).
        :param pulumi.Input[builtins.str] arguments: Command line arguments to pass to the script. Arguments are separated by spaces. ex: -Name blue* -Location 'West US 2' 
        :param pulumi.Input[Union[builtins.str, 'CleanupOptions']] cleanup_preference: The clean up preference when the script execution gets in a terminal state. Default setting is 'Always'.
        :param pulumi.Input['ContainerConfigurationArgs'] container_settings: Container settings.
        :param pulumi.Input[Sequence[pulumi.Input['EnvironmentVariableArgs']]] environment_variables: The environment variables to pass over to the script.
        :param pulumi.Input[builtins.str] force_update_tag: Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID.
        :param pulumi.Input['ManagedServiceIdentityArgs'] identity: Optional property. Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
        :param pulumi.Input[builtins.str] location: The location of the ACI and the storage account for the deployment script.
        :param pulumi.Input[builtins.str] primary_script_uri: Uri for the script. This is the entry point for the external script.
        :param pulumi.Input[builtins.str] script_content: Script body.
        :param pulumi.Input[builtins.str] script_name: Name of the deployment script.
        :param pulumi.Input['StorageAccountConfigurationArgs'] storage_account_settings: Storage Account settings.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] supporting_script_uris: Supporting files for the external script.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[builtins.str] timeout: Maximum allowed script execution time specified in ISO 8601 format. Default value is P1D
        """
        pulumi.set(__self__, "az_power_shell_version", az_power_shell_version)
        pulumi.set(__self__, "kind", 'AzurePowerShell')
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "retention_interval", retention_interval)
        if arguments is not None:
            pulumi.set(__self__, "arguments", arguments)
        if cleanup_preference is None:
            cleanup_preference = 'Always'
        if cleanup_preference is not None:
            pulumi.set(__self__, "cleanup_preference", cleanup_preference)
        if container_settings is not None:
            pulumi.set(__self__, "container_settings", container_settings)
        if environment_variables is not None:
            pulumi.set(__self__, "environment_variables", environment_variables)
        if force_update_tag is not None:
            pulumi.set(__self__, "force_update_tag", force_update_tag)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if primary_script_uri is not None:
            pulumi.set(__self__, "primary_script_uri", primary_script_uri)
        if script_content is not None:
            pulumi.set(__self__, "script_content", script_content)
        if script_name is not None:
            pulumi.set(__self__, "script_name", script_name)
        if storage_account_settings is not None:
            pulumi.set(__self__, "storage_account_settings", storage_account_settings)
        if supporting_script_uris is not None:
            pulumi.set(__self__, "supporting_script_uris", supporting_script_uris)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if timeout is None:
            timeout = 'P1D'
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter(name="azPowerShellVersion")
    def az_power_shell_version(self) -> pulumi.Input[builtins.str]:
        """
        Azure PowerShell module version to be used.
        """
        return pulumi.get(self, "az_power_shell_version")

    @az_power_shell_version.setter
    def az_power_shell_version(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "az_power_shell_version", value)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Input[builtins.str]:
        """
        Type of the script.
        Expected value is 'AzurePowerShell'.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group. The name is case insensitive.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="retentionInterval")
    def retention_interval(self) -> pulumi.Input[builtins.str]:
        """
        Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. Duration is based on ISO 8601 pattern (for example P1D means one day).
        """
        return pulumi.get(self, "retention_interval")

    @retention_interval.setter
    def retention_interval(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "retention_interval", value)

    @property
    @pulumi.getter
    def arguments(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Command line arguments to pass to the script. Arguments are separated by spaces. ex: -Name blue* -Location 'West US 2' 
        """
        return pulumi.get(self, "arguments")

    @arguments.setter
    def arguments(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "arguments", value)

    @property
    @pulumi.getter(name="cleanupPreference")
    def cleanup_preference(self) -> Optional[pulumi.Input[Union[builtins.str, 'CleanupOptions']]]:
        """
        The clean up preference when the script execution gets in a terminal state. Default setting is 'Always'.
        """
        return pulumi.get(self, "cleanup_preference")

    @cleanup_preference.setter
    def cleanup_preference(self, value: Optional[pulumi.Input[Union[builtins.str, 'CleanupOptions']]]):
        pulumi.set(self, "cleanup_preference", value)

    @property
    @pulumi.getter(name="containerSettings")
    def container_settings(self) -> Optional[pulumi.Input['ContainerConfigurationArgs']]:
        """
        Container settings.
        """
        return pulumi.get(self, "container_settings")

    @container_settings.setter
    def container_settings(self, value: Optional[pulumi.Input['ContainerConfigurationArgs']]):
        pulumi.set(self, "container_settings", value)

    @property
    @pulumi.getter(name="environmentVariables")
    def environment_variables(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EnvironmentVariableArgs']]]]:
        """
        The environment variables to pass over to the script.
        """
        return pulumi.get(self, "environment_variables")

    @environment_variables.setter
    def environment_variables(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EnvironmentVariableArgs']]]]):
        pulumi.set(self, "environment_variables", value)

    @property
    @pulumi.getter(name="forceUpdateTag")
    def force_update_tag(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID.
        """
        return pulumi.get(self, "force_update_tag")

    @force_update_tag.setter
    def force_update_tag(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "force_update_tag", value)

    @property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input['ManagedServiceIdentityArgs']]:
        """
        Optional property. Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input['ManagedServiceIdentityArgs']]):
        pulumi.set(self, "identity", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The location of the ACI and the storage account for the deployment script.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="primaryScriptUri")
    def primary_script_uri(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Uri for the script. This is the entry point for the external script.
        """
        return pulumi.get(self, "primary_script_uri")

    @primary_script_uri.setter
    def primary_script_uri(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "primary_script_uri", value)

    @property
    @pulumi.getter(name="scriptContent")
    def script_content(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Script body.
        """
        return pulumi.get(self, "script_content")

    @script_content.setter
    def script_content(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "script_content", value)

    @property
    @pulumi.getter(name="scriptName")
    def script_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the deployment script.
        """
        return pulumi.get(self, "script_name")

    @script_name.setter
    def script_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "script_name", value)

    @property
    @pulumi.getter(name="storageAccountSettings")
    def storage_account_settings(self) -> Optional[pulumi.Input['StorageAccountConfigurationArgs']]:
        """
        Storage Account settings.
        """
        return pulumi.get(self, "storage_account_settings")

    @storage_account_settings.setter
    def storage_account_settings(self, value: Optional[pulumi.Input['StorageAccountConfigurationArgs']]):
        pulumi.set(self, "storage_account_settings", value)

    @property
    @pulumi.getter(name="supportingScriptUris")
    def supporting_script_uris(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Supporting files for the external script.
        """
        return pulumi.get(self, "supporting_script_uris")

    @supporting_script_uris.setter
    def supporting_script_uris(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "supporting_script_uris", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Maximum allowed script execution time specified in ISO 8601 format. Default value is P1D
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "timeout", value)


@pulumi.type_token("azure-native:resources:AzurePowerShellScript")
class AzurePowerShellScript(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arguments: Optional[pulumi.Input[builtins.str]] = None,
                 az_power_shell_version: Optional[pulumi.Input[builtins.str]] = None,
                 cleanup_preference: Optional[pulumi.Input[Union[builtins.str, 'CleanupOptions']]] = None,
                 container_settings: Optional[pulumi.Input[Union['ContainerConfigurationArgs', 'ContainerConfigurationArgsDict']]] = None,
                 environment_variables: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EnvironmentVariableArgs', 'EnvironmentVariableArgsDict']]]]] = None,
                 force_update_tag: Optional[pulumi.Input[builtins.str]] = None,
                 identity: Optional[pulumi.Input[Union['ManagedServiceIdentityArgs', 'ManagedServiceIdentityArgsDict']]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 primary_script_uri: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 retention_interval: Optional[pulumi.Input[builtins.str]] = None,
                 script_content: Optional[pulumi.Input[builtins.str]] = None,
                 script_name: Optional[pulumi.Input[builtins.str]] = None,
                 storage_account_settings: Optional[pulumi.Input[Union['StorageAccountConfigurationArgs', 'StorageAccountConfigurationArgsDict']]] = None,
                 supporting_script_uris: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 timeout: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Object model for the Azure PowerShell script.

        Uses Azure REST API version 2023-08-01. In version 2.x of the Azure Native provider, it used API version 2020-10-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] arguments: Command line arguments to pass to the script. Arguments are separated by spaces. ex: -Name blue* -Location 'West US 2' 
        :param pulumi.Input[builtins.str] az_power_shell_version: Azure PowerShell module version to be used.
        :param pulumi.Input[Union[builtins.str, 'CleanupOptions']] cleanup_preference: The clean up preference when the script execution gets in a terminal state. Default setting is 'Always'.
        :param pulumi.Input[Union['ContainerConfigurationArgs', 'ContainerConfigurationArgsDict']] container_settings: Container settings.
        :param pulumi.Input[Sequence[pulumi.Input[Union['EnvironmentVariableArgs', 'EnvironmentVariableArgsDict']]]] environment_variables: The environment variables to pass over to the script.
        :param pulumi.Input[builtins.str] force_update_tag: Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID.
        :param pulumi.Input[Union['ManagedServiceIdentityArgs', 'ManagedServiceIdentityArgsDict']] identity: Optional property. Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
        :param pulumi.Input[builtins.str] kind: Type of the script.
               Expected value is 'AzurePowerShell'.
        :param pulumi.Input[builtins.str] location: The location of the ACI and the storage account for the deployment script.
        :param pulumi.Input[builtins.str] primary_script_uri: Uri for the script. This is the entry point for the external script.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] retention_interval: Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. Duration is based on ISO 8601 pattern (for example P1D means one day).
        :param pulumi.Input[builtins.str] script_content: Script body.
        :param pulumi.Input[builtins.str] script_name: Name of the deployment script.
        :param pulumi.Input[Union['StorageAccountConfigurationArgs', 'StorageAccountConfigurationArgsDict']] storage_account_settings: Storage Account settings.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] supporting_script_uris: Supporting files for the external script.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[builtins.str] timeout: Maximum allowed script execution time specified in ISO 8601 format. Default value is P1D
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AzurePowerShellScriptArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Object model for the Azure PowerShell script.

        Uses Azure REST API version 2023-08-01. In version 2.x of the Azure Native provider, it used API version 2020-10-01.

        :param str resource_name: The name of the resource.
        :param AzurePowerShellScriptArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AzurePowerShellScriptArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arguments: Optional[pulumi.Input[builtins.str]] = None,
                 az_power_shell_version: Optional[pulumi.Input[builtins.str]] = None,
                 cleanup_preference: Optional[pulumi.Input[Union[builtins.str, 'CleanupOptions']]] = None,
                 container_settings: Optional[pulumi.Input[Union['ContainerConfigurationArgs', 'ContainerConfigurationArgsDict']]] = None,
                 environment_variables: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EnvironmentVariableArgs', 'EnvironmentVariableArgsDict']]]]] = None,
                 force_update_tag: Optional[pulumi.Input[builtins.str]] = None,
                 identity: Optional[pulumi.Input[Union['ManagedServiceIdentityArgs', 'ManagedServiceIdentityArgsDict']]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 primary_script_uri: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 retention_interval: Optional[pulumi.Input[builtins.str]] = None,
                 script_content: Optional[pulumi.Input[builtins.str]] = None,
                 script_name: Optional[pulumi.Input[builtins.str]] = None,
                 storage_account_settings: Optional[pulumi.Input[Union['StorageAccountConfigurationArgs', 'StorageAccountConfigurationArgsDict']]] = None,
                 supporting_script_uris: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 timeout: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AzurePowerShellScriptArgs.__new__(AzurePowerShellScriptArgs)

            __props__.__dict__["arguments"] = arguments
            if az_power_shell_version is None and not opts.urn:
                raise TypeError("Missing required property 'az_power_shell_version'")
            __props__.__dict__["az_power_shell_version"] = az_power_shell_version
            if cleanup_preference is None:
                cleanup_preference = 'Always'
            __props__.__dict__["cleanup_preference"] = cleanup_preference
            __props__.__dict__["container_settings"] = container_settings
            __props__.__dict__["environment_variables"] = environment_variables
            __props__.__dict__["force_update_tag"] = force_update_tag
            __props__.__dict__["identity"] = identity
            if kind is None and not opts.urn:
                raise TypeError("Missing required property 'kind'")
            __props__.__dict__["kind"] = 'AzurePowerShell'
            __props__.__dict__["location"] = location
            __props__.__dict__["primary_script_uri"] = primary_script_uri
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if retention_interval is None and not opts.urn:
                raise TypeError("Missing required property 'retention_interval'")
            __props__.__dict__["retention_interval"] = retention_interval
            __props__.__dict__["script_content"] = script_content
            __props__.__dict__["script_name"] = script_name
            __props__.__dict__["storage_account_settings"] = storage_account_settings
            __props__.__dict__["supporting_script_uris"] = supporting_script_uris
            __props__.__dict__["tags"] = tags
            if timeout is None:
                timeout = 'P1D'
            __props__.__dict__["timeout"] = timeout
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["outputs"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:resources/v20191001preview:AzureCliScript"), pulumi.Alias(type_="azure-native:resources/v20191001preview:AzurePowerShellScript"), pulumi.Alias(type_="azure-native:resources/v20201001:AzureCliScript"), pulumi.Alias(type_="azure-native:resources/v20201001:AzurePowerShellScript"), pulumi.Alias(type_="azure-native:resources/v20230801:AzureCliScript"), pulumi.Alias(type_="azure-native:resources/v20230801:AzurePowerShellScript"), pulumi.Alias(type_="azure-native:resources:AzureCliScript")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AzurePowerShellScript, __self__).__init__(
            'azure-native:resources:AzurePowerShellScript',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AzurePowerShellScript':
        """
        Get an existing AzurePowerShellScript resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AzurePowerShellScriptArgs.__new__(AzurePowerShellScriptArgs)

        __props__.__dict__["arguments"] = None
        __props__.__dict__["az_power_shell_version"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["cleanup_preference"] = None
        __props__.__dict__["container_settings"] = None
        __props__.__dict__["environment_variables"] = None
        __props__.__dict__["force_update_tag"] = None
        __props__.__dict__["identity"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["outputs"] = None
        __props__.__dict__["primary_script_uri"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["retention_interval"] = None
        __props__.__dict__["script_content"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["storage_account_settings"] = None
        __props__.__dict__["supporting_script_uris"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["timeout"] = None
        __props__.__dict__["type"] = None
        return AzurePowerShellScript(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arguments(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Command line arguments to pass to the script. Arguments are separated by spaces. ex: -Name blue* -Location 'West US 2' 
        """
        return pulumi.get(self, "arguments")

    @property
    @pulumi.getter(name="azPowerShellVersion")
    def az_power_shell_version(self) -> pulumi.Output[builtins.str]:
        """
        Azure PowerShell module version to be used.
        """
        return pulumi.get(self, "az_power_shell_version")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="cleanupPreference")
    def cleanup_preference(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The clean up preference when the script execution gets in a terminal state. Default setting is 'Always'.
        """
        return pulumi.get(self, "cleanup_preference")

    @property
    @pulumi.getter(name="containerSettings")
    def container_settings(self) -> pulumi.Output[Optional['outputs.ContainerConfigurationResponse']]:
        """
        Container settings.
        """
        return pulumi.get(self, "container_settings")

    @property
    @pulumi.getter(name="environmentVariables")
    def environment_variables(self) -> pulumi.Output[Optional[Sequence['outputs.EnvironmentVariableResponse']]]:
        """
        The environment variables to pass over to the script.
        """
        return pulumi.get(self, "environment_variables")

    @property
    @pulumi.getter(name="forceUpdateTag")
    def force_update_tag(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID.
        """
        return pulumi.get(self, "force_update_tag")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.ManagedServiceIdentityResponse']]:
        """
        Optional property. Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[builtins.str]:
        """
        Type of the script.
        Expected value is 'AzurePowerShell'.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[builtins.str]:
        """
        The location of the ACI and the storage account for the deployment script.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Name of this resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def outputs(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        List of script outputs.
        """
        return pulumi.get(self, "outputs")

    @property
    @pulumi.getter(name="primaryScriptUri")
    def primary_script_uri(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Uri for the script. This is the entry point for the external script.
        """
        return pulumi.get(self, "primary_script_uri")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        State of the script execution. This only appears in the response.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="retentionInterval")
    def retention_interval(self) -> pulumi.Output[builtins.str]:
        """
        Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. Duration is based on ISO 8601 pattern (for example P1D means one day).
        """
        return pulumi.get(self, "retention_interval")

    @property
    @pulumi.getter(name="scriptContent")
    def script_content(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Script body.
        """
        return pulumi.get(self, "script_content")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['outputs.ScriptStatusResponse']:
        """
        Contains the results of script execution.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="storageAccountSettings")
    def storage_account_settings(self) -> pulumi.Output[Optional['outputs.StorageAccountConfigurationResponse']]:
        """
        Storage Account settings.
        """
        return pulumi.get(self, "storage_account_settings")

    @property
    @pulumi.getter(name="supportingScriptUris")
    def supporting_script_uris(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        Supporting files for the external script.
        """
        return pulumi.get(self, "supporting_script_uris")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        The system metadata related to this resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Maximum allowed script execution time specified in ISO 8601 format. Default value is P1D
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Type of this resource.
        """
        return pulumi.get(self, "type")

