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

__all__ = ['GalleryInVMAccessControlProfileVersionArgs', 'GalleryInVMAccessControlProfileVersion']

@pulumi.input_type
class GalleryInVMAccessControlProfileVersionArgs:
    def __init__(__self__, *,
                 default_access: pulumi.Input[Union[builtins.str, 'EndpointAccess']],
                 gallery_name: pulumi.Input[builtins.str],
                 in_vm_access_control_profile_name: pulumi.Input[builtins.str],
                 mode: pulumi.Input[Union[builtins.str, 'AccessControlRulesMode']],
                 resource_group_name: pulumi.Input[builtins.str],
                 exclude_from_latest: Optional[pulumi.Input[builtins.bool]] = None,
                 in_vm_access_control_profile_version_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 rules: Optional[pulumi.Input['AccessControlRulesArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 target_locations: Optional[pulumi.Input[Sequence[pulumi.Input['TargetRegionArgs']]]] = None):
        """
        The set of arguments for constructing a GalleryInVMAccessControlProfileVersion resource.
        :param pulumi.Input[Union[builtins.str, 'EndpointAccess']] default_access: This property allows you to specify if the requests will be allowed to access the host endpoints. Possible values are: 'Allow', 'Deny'.
        :param pulumi.Input[builtins.str] gallery_name: The name of the Shared Image Gallery.
        :param pulumi.Input[builtins.str] in_vm_access_control_profile_name: The name of the gallery inVMAccessControlProfile to be retrieved.
        :param pulumi.Input[Union[builtins.str, 'AccessControlRulesMode']] mode: This property allows you to specify whether the access control rules are in Audit mode, in Enforce mode or Disabled. Possible values are: 'Audit', 'Enforce' or 'Disabled'.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.bool] exclude_from_latest: If set to true, Virtual Machines deployed from the latest version of the Resource Profile won't use this Profile version.
        :param pulumi.Input[builtins.str] in_vm_access_control_profile_version_name: The name of the gallery inVMAccessControlProfile version to be retrieved.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input['AccessControlRulesArgs'] rules: This is the Access Control Rules specification for an inVMAccessControlProfile version.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[Sequence[pulumi.Input['TargetRegionArgs']]] target_locations: The target regions where the Resource Profile version is going to be replicated to. This property is updatable.
        """
        pulumi.set(__self__, "default_access", default_access)
        pulumi.set(__self__, "gallery_name", gallery_name)
        pulumi.set(__self__, "in_vm_access_control_profile_name", in_vm_access_control_profile_name)
        pulumi.set(__self__, "mode", mode)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if exclude_from_latest is not None:
            pulumi.set(__self__, "exclude_from_latest", exclude_from_latest)
        if in_vm_access_control_profile_version_name is not None:
            pulumi.set(__self__, "in_vm_access_control_profile_version_name", in_vm_access_control_profile_version_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if target_locations is not None:
            pulumi.set(__self__, "target_locations", target_locations)

    @property
    @pulumi.getter(name="defaultAccess")
    def default_access(self) -> pulumi.Input[Union[builtins.str, 'EndpointAccess']]:
        """
        This property allows you to specify if the requests will be allowed to access the host endpoints. Possible values are: 'Allow', 'Deny'.
        """
        return pulumi.get(self, "default_access")

    @default_access.setter
    def default_access(self, value: pulumi.Input[Union[builtins.str, 'EndpointAccess']]):
        pulumi.set(self, "default_access", value)

    @property
    @pulumi.getter(name="galleryName")
    def gallery_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the Shared Image Gallery.
        """
        return pulumi.get(self, "gallery_name")

    @gallery_name.setter
    def gallery_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "gallery_name", value)

    @property
    @pulumi.getter(name="inVMAccessControlProfileName")
    def in_vm_access_control_profile_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the gallery inVMAccessControlProfile to be retrieved.
        """
        return pulumi.get(self, "in_vm_access_control_profile_name")

    @in_vm_access_control_profile_name.setter
    def in_vm_access_control_profile_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "in_vm_access_control_profile_name", value)

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Input[Union[builtins.str, 'AccessControlRulesMode']]:
        """
        This property allows you to specify whether the access control rules are in Audit mode, in Enforce mode or Disabled. Possible values are: 'Audit', 'Enforce' or 'Disabled'.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: pulumi.Input[Union[builtins.str, 'AccessControlRulesMode']]):
        pulumi.set(self, "mode", value)

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
    @pulumi.getter(name="excludeFromLatest")
    def exclude_from_latest(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If set to true, Virtual Machines deployed from the latest version of the Resource Profile won't use this Profile version.
        """
        return pulumi.get(self, "exclude_from_latest")

    @exclude_from_latest.setter
    def exclude_from_latest(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "exclude_from_latest", value)

    @property
    @pulumi.getter(name="inVMAccessControlProfileVersionName")
    def in_vm_access_control_profile_version_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the gallery inVMAccessControlProfile version to be retrieved.
        """
        return pulumi.get(self, "in_vm_access_control_profile_version_name")

    @in_vm_access_control_profile_version_name.setter
    def in_vm_access_control_profile_version_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "in_vm_access_control_profile_version_name", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input['AccessControlRulesArgs']]:
        """
        This is the Access Control Rules specification for an inVMAccessControlProfile version.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input['AccessControlRulesArgs']]):
        pulumi.set(self, "rules", value)

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
    @pulumi.getter(name="targetLocations")
    def target_locations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TargetRegionArgs']]]]:
        """
        The target regions where the Resource Profile version is going to be replicated to. This property is updatable.
        """
        return pulumi.get(self, "target_locations")

    @target_locations.setter
    def target_locations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TargetRegionArgs']]]]):
        pulumi.set(self, "target_locations", value)


@pulumi.type_token("azure-native:compute:GalleryInVMAccessControlProfileVersion")
class GalleryInVMAccessControlProfileVersion(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_access: Optional[pulumi.Input[Union[builtins.str, 'EndpointAccess']]] = None,
                 exclude_from_latest: Optional[pulumi.Input[builtins.bool]] = None,
                 gallery_name: Optional[pulumi.Input[builtins.str]] = None,
                 in_vm_access_control_profile_name: Optional[pulumi.Input[builtins.str]] = None,
                 in_vm_access_control_profile_version_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 mode: Optional[pulumi.Input[Union[builtins.str, 'AccessControlRulesMode']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 rules: Optional[pulumi.Input[Union['AccessControlRulesArgs', 'AccessControlRulesArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 target_locations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TargetRegionArgs', 'TargetRegionArgsDict']]]]] = None,
                 __props__=None):
        """
        Specifies information about the gallery inVMAccessControlProfile version that you want to create or update.

        Uses Azure REST API version 2024-03-03. In version 2.x of the Azure Native provider, it used API version 2024-03-03.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union[builtins.str, 'EndpointAccess']] default_access: This property allows you to specify if the requests will be allowed to access the host endpoints. Possible values are: 'Allow', 'Deny'.
        :param pulumi.Input[builtins.bool] exclude_from_latest: If set to true, Virtual Machines deployed from the latest version of the Resource Profile won't use this Profile version.
        :param pulumi.Input[builtins.str] gallery_name: The name of the Shared Image Gallery.
        :param pulumi.Input[builtins.str] in_vm_access_control_profile_name: The name of the gallery inVMAccessControlProfile to be retrieved.
        :param pulumi.Input[builtins.str] in_vm_access_control_profile_version_name: The name of the gallery inVMAccessControlProfile version to be retrieved.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[Union[builtins.str, 'AccessControlRulesMode']] mode: This property allows you to specify whether the access control rules are in Audit mode, in Enforce mode or Disabled. Possible values are: 'Audit', 'Enforce' or 'Disabled'.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Union['AccessControlRulesArgs', 'AccessControlRulesArgsDict']] rules: This is the Access Control Rules specification for an inVMAccessControlProfile version.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[Sequence[pulumi.Input[Union['TargetRegionArgs', 'TargetRegionArgsDict']]]] target_locations: The target regions where the Resource Profile version is going to be replicated to. This property is updatable.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GalleryInVMAccessControlProfileVersionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Specifies information about the gallery inVMAccessControlProfile version that you want to create or update.

        Uses Azure REST API version 2024-03-03. In version 2.x of the Azure Native provider, it used API version 2024-03-03.

        :param str resource_name: The name of the resource.
        :param GalleryInVMAccessControlProfileVersionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GalleryInVMAccessControlProfileVersionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_access: Optional[pulumi.Input[Union[builtins.str, 'EndpointAccess']]] = None,
                 exclude_from_latest: Optional[pulumi.Input[builtins.bool]] = None,
                 gallery_name: Optional[pulumi.Input[builtins.str]] = None,
                 in_vm_access_control_profile_name: Optional[pulumi.Input[builtins.str]] = None,
                 in_vm_access_control_profile_version_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 mode: Optional[pulumi.Input[Union[builtins.str, 'AccessControlRulesMode']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 rules: Optional[pulumi.Input[Union['AccessControlRulesArgs', 'AccessControlRulesArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 target_locations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TargetRegionArgs', 'TargetRegionArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GalleryInVMAccessControlProfileVersionArgs.__new__(GalleryInVMAccessControlProfileVersionArgs)

            if default_access is None and not opts.urn:
                raise TypeError("Missing required property 'default_access'")
            __props__.__dict__["default_access"] = default_access
            __props__.__dict__["exclude_from_latest"] = exclude_from_latest
            if gallery_name is None and not opts.urn:
                raise TypeError("Missing required property 'gallery_name'")
            __props__.__dict__["gallery_name"] = gallery_name
            if in_vm_access_control_profile_name is None and not opts.urn:
                raise TypeError("Missing required property 'in_vm_access_control_profile_name'")
            __props__.__dict__["in_vm_access_control_profile_name"] = in_vm_access_control_profile_name
            __props__.__dict__["in_vm_access_control_profile_version_name"] = in_vm_access_control_profile_version_name
            __props__.__dict__["location"] = location
            if mode is None and not opts.urn:
                raise TypeError("Missing required property 'mode'")
            __props__.__dict__["mode"] = mode
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["rules"] = rules
            __props__.__dict__["tags"] = tags
            __props__.__dict__["target_locations"] = target_locations
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["published_date"] = None
            __props__.__dict__["replication_status"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:compute/v20240303:GalleryInVMAccessControlProfileVersion")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(GalleryInVMAccessControlProfileVersion, __self__).__init__(
            'azure-native:compute:GalleryInVMAccessControlProfileVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'GalleryInVMAccessControlProfileVersion':
        """
        Get an existing GalleryInVMAccessControlProfileVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = GalleryInVMAccessControlProfileVersionArgs.__new__(GalleryInVMAccessControlProfileVersionArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["default_access"] = None
        __props__.__dict__["exclude_from_latest"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["mode"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["published_date"] = None
        __props__.__dict__["replication_status"] = None
        __props__.__dict__["rules"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["target_locations"] = None
        __props__.__dict__["type"] = None
        return GalleryInVMAccessControlProfileVersion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="defaultAccess")
    def default_access(self) -> pulumi.Output[builtins.str]:
        """
        This property allows you to specify if the requests will be allowed to access the host endpoints. Possible values are: 'Allow', 'Deny'.
        """
        return pulumi.get(self, "default_access")

    @property
    @pulumi.getter(name="excludeFromLatest")
    def exclude_from_latest(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        If set to true, Virtual Machines deployed from the latest version of the Resource Profile won't use this Profile version.
        """
        return pulumi.get(self, "exclude_from_latest")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[builtins.str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[builtins.str]:
        """
        This property allows you to specify whether the access control rules are in Audit mode, in Enforce mode or Disabled. Possible values are: 'Audit', 'Enforce' or 'Disabled'.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        The provisioning state, which only appears in the response.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publishedDate")
    def published_date(self) -> pulumi.Output[builtins.str]:
        """
        The timestamp for when the Resource Profile Version is published.
        """
        return pulumi.get(self, "published_date")

    @property
    @pulumi.getter(name="replicationStatus")
    def replication_status(self) -> pulumi.Output['outputs.ReplicationStatusResponse']:
        """
        This is the replication status of the gallery image version.
        """
        return pulumi.get(self, "replication_status")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Optional['outputs.AccessControlRulesResponse']]:
        """
        This is the Access Control Rules specification for an inVMAccessControlProfile version.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
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
    @pulumi.getter(name="targetLocations")
    def target_locations(self) -> pulumi.Output[Optional[Sequence['outputs.TargetRegionResponse']]]:
        """
        The target regions where the Resource Profile version is going to be replicated to. This property is updatable.
        """
        return pulumi.get(self, "target_locations")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

