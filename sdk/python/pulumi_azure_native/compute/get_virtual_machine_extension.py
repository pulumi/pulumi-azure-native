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

__all__ = [
    'GetVirtualMachineExtensionResult',
    'AwaitableGetVirtualMachineExtensionResult',
    'get_virtual_machine_extension',
    'get_virtual_machine_extension_output',
]

@pulumi.output_type
class GetVirtualMachineExtensionResult:
    """
    Describes a Virtual Machine Extension.
    """
    def __init__(__self__, auto_upgrade_minor_version=None, azure_api_version=None, enable_automatic_upgrade=None, force_update_tag=None, id=None, instance_view=None, location=None, name=None, protected_settings=None, protected_settings_from_key_vault=None, provision_after_extensions=None, provisioning_state=None, publisher=None, settings=None, suppress_failures=None, system_data=None, tags=None, type=None, type_handler_version=None):
        if auto_upgrade_minor_version and not isinstance(auto_upgrade_minor_version, bool):
            raise TypeError("Expected argument 'auto_upgrade_minor_version' to be a bool")
        pulumi.set(__self__, "auto_upgrade_minor_version", auto_upgrade_minor_version)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if enable_automatic_upgrade and not isinstance(enable_automatic_upgrade, bool):
            raise TypeError("Expected argument 'enable_automatic_upgrade' to be a bool")
        pulumi.set(__self__, "enable_automatic_upgrade", enable_automatic_upgrade)
        if force_update_tag and not isinstance(force_update_tag, str):
            raise TypeError("Expected argument 'force_update_tag' to be a str")
        pulumi.set(__self__, "force_update_tag", force_update_tag)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_view and not isinstance(instance_view, dict):
            raise TypeError("Expected argument 'instance_view' to be a dict")
        pulumi.set(__self__, "instance_view", instance_view)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if protected_settings and not isinstance(protected_settings, dict):
            raise TypeError("Expected argument 'protected_settings' to be a dict")
        pulumi.set(__self__, "protected_settings", protected_settings)
        if protected_settings_from_key_vault and not isinstance(protected_settings_from_key_vault, dict):
            raise TypeError("Expected argument 'protected_settings_from_key_vault' to be a dict")
        pulumi.set(__self__, "protected_settings_from_key_vault", protected_settings_from_key_vault)
        if provision_after_extensions and not isinstance(provision_after_extensions, list):
            raise TypeError("Expected argument 'provision_after_extensions' to be a list")
        pulumi.set(__self__, "provision_after_extensions", provision_after_extensions)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if publisher and not isinstance(publisher, str):
            raise TypeError("Expected argument 'publisher' to be a str")
        pulumi.set(__self__, "publisher", publisher)
        if settings and not isinstance(settings, dict):
            raise TypeError("Expected argument 'settings' to be a dict")
        pulumi.set(__self__, "settings", settings)
        if suppress_failures and not isinstance(suppress_failures, bool):
            raise TypeError("Expected argument 'suppress_failures' to be a bool")
        pulumi.set(__self__, "suppress_failures", suppress_failures)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if type_handler_version and not isinstance(type_handler_version, str):
            raise TypeError("Expected argument 'type_handler_version' to be a str")
        pulumi.set(__self__, "type_handler_version", type_handler_version)

    @property
    @pulumi.getter(name="autoUpgradeMinorVersion")
    def auto_upgrade_minor_version(self) -> Optional[builtins.bool]:
        """
        Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
        """
        return pulumi.get(self, "auto_upgrade_minor_version")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="enableAutomaticUpgrade")
    def enable_automatic_upgrade(self) -> Optional[builtins.bool]:
        """
        Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
        """
        return pulumi.get(self, "enable_automatic_upgrade")

    @property
    @pulumi.getter(name="forceUpdateTag")
    def force_update_tag(self) -> Optional[builtins.str]:
        """
        How the extension handler should be forced to update even if the extension configuration has not changed.
        """
        return pulumi.get(self, "force_update_tag")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceView")
    def instance_view(self) -> Optional['outputs.VirtualMachineExtensionInstanceViewResponse']:
        """
        The virtual machine extension instance view.
        """
        return pulumi.get(self, "instance_view")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="protectedSettings")
    def protected_settings(self) -> Optional[Any]:
        """
        The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
        """
        return pulumi.get(self, "protected_settings")

    @property
    @pulumi.getter(name="protectedSettingsFromKeyVault")
    def protected_settings_from_key_vault(self) -> Optional['outputs.KeyVaultSecretReferenceResponse']:
        """
        The extensions protected settings that are passed by reference, and consumed from key vault
        """
        return pulumi.get(self, "protected_settings_from_key_vault")

    @property
    @pulumi.getter(name="provisionAfterExtensions")
    def provision_after_extensions(self) -> Optional[Sequence[builtins.str]]:
        """
        Collection of extension names after which this extension needs to be provisioned.
        """
        return pulumi.get(self, "provision_after_extensions")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        The provisioning state, which only appears in the response.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def publisher(self) -> Optional[builtins.str]:
        """
        The name of the extension handler publisher.
        """
        return pulumi.get(self, "publisher")

    @property
    @pulumi.getter
    def settings(self) -> Optional[Any]:
        """
        Json formatted public settings for the extension.
        """
        return pulumi.get(self, "settings")

    @property
    @pulumi.getter(name="suppressFailures")
    def suppress_failures(self) -> Optional[builtins.bool]:
        """
        Indicates whether failures stemming from the extension will be suppressed (Operational failures such as not connecting to the VM will not be suppressed regardless of this value). The default is false.
        """
        return pulumi.get(self, "suppress_failures")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="typeHandlerVersion")
    def type_handler_version(self) -> Optional[builtins.str]:
        """
        Specifies the version of the script handler.
        """
        return pulumi.get(self, "type_handler_version")


class AwaitableGetVirtualMachineExtensionResult(GetVirtualMachineExtensionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVirtualMachineExtensionResult(
            auto_upgrade_minor_version=self.auto_upgrade_minor_version,
            azure_api_version=self.azure_api_version,
            enable_automatic_upgrade=self.enable_automatic_upgrade,
            force_update_tag=self.force_update_tag,
            id=self.id,
            instance_view=self.instance_view,
            location=self.location,
            name=self.name,
            protected_settings=self.protected_settings,
            protected_settings_from_key_vault=self.protected_settings_from_key_vault,
            provision_after_extensions=self.provision_after_extensions,
            provisioning_state=self.provisioning_state,
            publisher=self.publisher,
            settings=self.settings,
            suppress_failures=self.suppress_failures,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type,
            type_handler_version=self.type_handler_version)


def get_virtual_machine_extension(expand: Optional[builtins.str] = None,
                                  resource_group_name: Optional[builtins.str] = None,
                                  vm_extension_name: Optional[builtins.str] = None,
                                  vm_name: Optional[builtins.str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVirtualMachineExtensionResult:
    """
    The operation to get the extension.

    Uses Azure REST API version 2024-11-01.

    Other available API versions: 2022-08-01, 2022-11-01, 2023-03-01, 2023-07-01, 2023-09-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str expand: The expand expression to apply on the operation.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str vm_extension_name: The name of the virtual machine extension.
    :param builtins.str vm_name: The name of the virtual machine.
    """
    __args__ = dict()
    __args__['expand'] = expand
    __args__['resourceGroupName'] = resource_group_name
    __args__['vmExtensionName'] = vm_extension_name
    __args__['vmName'] = vm_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:compute:getVirtualMachineExtension', __args__, opts=opts, typ=GetVirtualMachineExtensionResult).value

    return AwaitableGetVirtualMachineExtensionResult(
        auto_upgrade_minor_version=pulumi.get(__ret__, 'auto_upgrade_minor_version'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        enable_automatic_upgrade=pulumi.get(__ret__, 'enable_automatic_upgrade'),
        force_update_tag=pulumi.get(__ret__, 'force_update_tag'),
        id=pulumi.get(__ret__, 'id'),
        instance_view=pulumi.get(__ret__, 'instance_view'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        protected_settings=pulumi.get(__ret__, 'protected_settings'),
        protected_settings_from_key_vault=pulumi.get(__ret__, 'protected_settings_from_key_vault'),
        provision_after_extensions=pulumi.get(__ret__, 'provision_after_extensions'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        publisher=pulumi.get(__ret__, 'publisher'),
        settings=pulumi.get(__ret__, 'settings'),
        suppress_failures=pulumi.get(__ret__, 'suppress_failures'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        type_handler_version=pulumi.get(__ret__, 'type_handler_version'))
def get_virtual_machine_extension_output(expand: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                         resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                         vm_extension_name: Optional[pulumi.Input[builtins.str]] = None,
                                         vm_name: Optional[pulumi.Input[builtins.str]] = None,
                                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetVirtualMachineExtensionResult]:
    """
    The operation to get the extension.

    Uses Azure REST API version 2024-11-01.

    Other available API versions: 2022-08-01, 2022-11-01, 2023-03-01, 2023-07-01, 2023-09-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str expand: The expand expression to apply on the operation.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str vm_extension_name: The name of the virtual machine extension.
    :param builtins.str vm_name: The name of the virtual machine.
    """
    __args__ = dict()
    __args__['expand'] = expand
    __args__['resourceGroupName'] = resource_group_name
    __args__['vmExtensionName'] = vm_extension_name
    __args__['vmName'] = vm_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:compute:getVirtualMachineExtension', __args__, opts=opts, typ=GetVirtualMachineExtensionResult)
    return __ret__.apply(lambda __response__: GetVirtualMachineExtensionResult(
        auto_upgrade_minor_version=pulumi.get(__response__, 'auto_upgrade_minor_version'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        enable_automatic_upgrade=pulumi.get(__response__, 'enable_automatic_upgrade'),
        force_update_tag=pulumi.get(__response__, 'force_update_tag'),
        id=pulumi.get(__response__, 'id'),
        instance_view=pulumi.get(__response__, 'instance_view'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        protected_settings=pulumi.get(__response__, 'protected_settings'),
        protected_settings_from_key_vault=pulumi.get(__response__, 'protected_settings_from_key_vault'),
        provision_after_extensions=pulumi.get(__response__, 'provision_after_extensions'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        publisher=pulumi.get(__response__, 'publisher'),
        settings=pulumi.get(__response__, 'settings'),
        suppress_failures=pulumi.get(__response__, 'suppress_failures'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type'),
        type_handler_version=pulumi.get(__response__, 'type_handler_version')))
