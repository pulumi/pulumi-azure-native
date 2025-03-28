# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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

__all__ = ['PackageArgs', 'Package']

@pulumi.input_type
class PackageArgs:
    def __init__(__self__, *,
                 application_name: pulumi.Input[str],
                 blob_path: pulumi.Input[str],
                 flighting_ring: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 target_os_list: pulumi.Input[Sequence[pulumi.Input['TargetOSInfoArgs']]],
                 test_base_account_name: pulumi.Input[str],
                 tests: pulumi.Input[Sequence[pulumi.Input['TestArgs']]],
                 version: pulumi.Input[str],
                 location: Optional[pulumi.Input[str]] = None,
                 package_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Package resource.
        :param pulumi.Input[str] application_name: Application name
        :param pulumi.Input[str] blob_path: The file path of the package.
        :param pulumi.Input[str] flighting_ring: The flighting ring for feature update.
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the resource.
        :param pulumi.Input[Sequence[pulumi.Input['TargetOSInfoArgs']]] target_os_list: Specifies the target OSs of specific OS Update types.
        :param pulumi.Input[str] test_base_account_name: The resource name of the Test Base Account.
        :param pulumi.Input[Sequence[pulumi.Input['TestArgs']]] tests: The detailed test information.
        :param pulumi.Input[str] version: Application version
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input[str] package_name: The resource name of the Test Base Package.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags of the resource.
        """
        pulumi.set(__self__, "application_name", application_name)
        pulumi.set(__self__, "blob_path", blob_path)
        pulumi.set(__self__, "flighting_ring", flighting_ring)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "target_os_list", target_os_list)
        pulumi.set(__self__, "test_base_account_name", test_base_account_name)
        pulumi.set(__self__, "tests", tests)
        pulumi.set(__self__, "version", version)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if package_name is not None:
            pulumi.set(__self__, "package_name", package_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> pulumi.Input[str]:
        """
        Application name
        """
        return pulumi.get(self, "application_name")

    @application_name.setter
    def application_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "application_name", value)

    @property
    @pulumi.getter(name="blobPath")
    def blob_path(self) -> pulumi.Input[str]:
        """
        The file path of the package.
        """
        return pulumi.get(self, "blob_path")

    @blob_path.setter
    def blob_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "blob_path", value)

    @property
    @pulumi.getter(name="flightingRing")
    def flighting_ring(self) -> pulumi.Input[str]:
        """
        The flighting ring for feature update.
        """
        return pulumi.get(self, "flighting_ring")

    @flighting_ring.setter
    def flighting_ring(self, value: pulumi.Input[str]):
        pulumi.set(self, "flighting_ring", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group that contains the resource.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="targetOSList")
    def target_os_list(self) -> pulumi.Input[Sequence[pulumi.Input['TargetOSInfoArgs']]]:
        """
        Specifies the target OSs of specific OS Update types.
        """
        return pulumi.get(self, "target_os_list")

    @target_os_list.setter
    def target_os_list(self, value: pulumi.Input[Sequence[pulumi.Input['TargetOSInfoArgs']]]):
        pulumi.set(self, "target_os_list", value)

    @property
    @pulumi.getter(name="testBaseAccountName")
    def test_base_account_name(self) -> pulumi.Input[str]:
        """
        The resource name of the Test Base Account.
        """
        return pulumi.get(self, "test_base_account_name")

    @test_base_account_name.setter
    def test_base_account_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "test_base_account_name", value)

    @property
    @pulumi.getter
    def tests(self) -> pulumi.Input[Sequence[pulumi.Input['TestArgs']]]:
        """
        The detailed test information.
        """
        return pulumi.get(self, "tests")

    @tests.setter
    def tests(self, value: pulumi.Input[Sequence[pulumi.Input['TestArgs']]]):
        pulumi.set(self, "tests", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[str]:
        """
        Application version
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[str]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="packageName")
    def package_name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the Test Base Package.
        """
        return pulumi.get(self, "package_name")

    @package_name.setter
    def package_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "package_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class Package(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_name: Optional[pulumi.Input[str]] = None,
                 blob_path: Optional[pulumi.Input[str]] = None,
                 flighting_ring: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 package_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_os_list: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TargetOSInfoArgs', 'TargetOSInfoArgsDict']]]]] = None,
                 test_base_account_name: Optional[pulumi.Input[str]] = None,
                 tests: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TestArgs', 'TestArgsDict']]]]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The Test Base Package resource.

        Uses Azure REST API version 2022-04-01-preview. In version 1.x of the Azure Native provider, it used API version 2022-04-01-preview.

        Other available API versions: 2023-11-01-preview.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_name: Application name
        :param pulumi.Input[str] blob_path: The file path of the package.
        :param pulumi.Input[str] flighting_ring: The flighting ring for feature update.
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input[str] package_name: The resource name of the Test Base Package.
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags of the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['TargetOSInfoArgs', 'TargetOSInfoArgsDict']]]] target_os_list: Specifies the target OSs of specific OS Update types.
        :param pulumi.Input[str] test_base_account_name: The resource name of the Test Base Account.
        :param pulumi.Input[Sequence[pulumi.Input[Union['TestArgs', 'TestArgsDict']]]] tests: The detailed test information.
        :param pulumi.Input[str] version: Application version
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PackageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The Test Base Package resource.

        Uses Azure REST API version 2022-04-01-preview. In version 1.x of the Azure Native provider, it used API version 2022-04-01-preview.

        Other available API versions: 2023-11-01-preview.

        :param str resource_name: The name of the resource.
        :param PackageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PackageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_name: Optional[pulumi.Input[str]] = None,
                 blob_path: Optional[pulumi.Input[str]] = None,
                 flighting_ring: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 package_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_os_list: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TargetOSInfoArgs', 'TargetOSInfoArgsDict']]]]] = None,
                 test_base_account_name: Optional[pulumi.Input[str]] = None,
                 tests: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TestArgs', 'TestArgsDict']]]]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PackageArgs.__new__(PackageArgs)

            if application_name is None and not opts.urn:
                raise TypeError("Missing required property 'application_name'")
            __props__.__dict__["application_name"] = application_name
            if blob_path is None and not opts.urn:
                raise TypeError("Missing required property 'blob_path'")
            __props__.__dict__["blob_path"] = blob_path
            if flighting_ring is None and not opts.urn:
                raise TypeError("Missing required property 'flighting_ring'")
            __props__.__dict__["flighting_ring"] = flighting_ring
            __props__.__dict__["location"] = location
            __props__.__dict__["package_name"] = package_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            if target_os_list is None and not opts.urn:
                raise TypeError("Missing required property 'target_os_list'")
            __props__.__dict__["target_os_list"] = target_os_list
            if test_base_account_name is None and not opts.urn:
                raise TypeError("Missing required property 'test_base_account_name'")
            __props__.__dict__["test_base_account_name"] = test_base_account_name
            if tests is None and not opts.urn:
                raise TypeError("Missing required property 'tests'")
            __props__.__dict__["tests"] = tests
            if version is None and not opts.urn:
                raise TypeError("Missing required property 'version'")
            __props__.__dict__["version"] = version
            __props__.__dict__["etag"] = None
            __props__.__dict__["is_enabled"] = None
            __props__.__dict__["last_modified_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["package_status"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["test_types"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["validation_results"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:testbase/v20201216preview:Package"), pulumi.Alias(type_="azure-native:testbase/v20220401preview:Package"), pulumi.Alias(type_="azure-native:testbase/v20231101preview:Package")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Package, __self__).__init__(
            'azure-native:testbase:Package',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Package':
        """
        Get an existing Package resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PackageArgs.__new__(PackageArgs)

        __props__.__dict__["application_name"] = None
        __props__.__dict__["blob_path"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["flighting_ring"] = None
        __props__.__dict__["is_enabled"] = None
        __props__.__dict__["last_modified_time"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["package_status"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["target_os_list"] = None
        __props__.__dict__["test_types"] = None
        __props__.__dict__["tests"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["validation_results"] = None
        __props__.__dict__["version"] = None
        return Package(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> pulumi.Output[str]:
        """
        Application name
        """
        return pulumi.get(self, "application_name")

    @property
    @pulumi.getter(name="blobPath")
    def blob_path(self) -> pulumi.Output[str]:
        """
        The file path of the package.
        """
        return pulumi.get(self, "blob_path")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        Resource Etag.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="flightingRing")
    def flighting_ring(self) -> pulumi.Output[str]:
        """
        The flighting ring for feature update.
        """
        return pulumi.get(self, "flighting_ring")

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> pulumi.Output[bool]:
        """
        Flag showing that whether the package is enabled. It doesn't schedule test for package which is not enabled.
        """
        return pulumi.get(self, "is_enabled")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> pulumi.Output[str]:
        """
        The UTC timestamp when the package was last modified.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="packageStatus")
    def package_status(self) -> pulumi.Output[str]:
        """
        The status of the package.
        """
        return pulumi.get(self, "package_status")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        The system metadata relating to this resource
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetOSList")
    def target_os_list(self) -> pulumi.Output[Sequence['outputs.TargetOSInfoResponse']]:
        """
        Specifies the target OSs of specific OS Update types.
        """
        return pulumi.get(self, "target_os_list")

    @property
    @pulumi.getter(name="testTypes")
    def test_types(self) -> pulumi.Output[Sequence[str]]:
        """
        OOB, functional or both. Mapped to the data in 'tests' property.
        """
        return pulumi.get(self, "test_types")

    @property
    @pulumi.getter
    def tests(self) -> pulumi.Output[Sequence['outputs.TestResponse']]:
        """
        The detailed test information.
        """
        return pulumi.get(self, "tests")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="validationResults")
    def validation_results(self) -> pulumi.Output[Sequence['outputs.PackageValidationResultResponse']]:
        """
        The validation results. There's validation on package when it's created or updated.
        """
        return pulumi.get(self, "validation_results")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        Application version
        """
        return pulumi.get(self, "version")

