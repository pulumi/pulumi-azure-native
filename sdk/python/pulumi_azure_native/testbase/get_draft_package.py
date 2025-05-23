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
    'GetDraftPackageResult',
    'AwaitableGetDraftPackageResult',
    'get_draft_package',
    'get_draft_package_output',
]

@pulumi.output_type
class GetDraftPackageResult:
    """
    The Test Base Draft Package resource.
    """
    def __init__(__self__, app_file_name=None, application_name=None, azure_api_version=None, comments=None, draft_package_path=None, edit_package=None, executable_launch_command=None, first_party_apps=None, flighting_ring=None, gallery_apps=None, highlighted_files=None, id=None, inplace_upgrade_os_pair=None, intune_enrollment_metadata=None, intune_metadata=None, last_modified_time=None, name=None, package_id=None, package_tags=None, process_name=None, provisioning_state=None, source_type=None, system_data=None, tab_state=None, target_os_list=None, test_types=None, tests=None, type=None, use_autofill=None, use_sample=None, version=None, working_path=None):
        if app_file_name and not isinstance(app_file_name, str):
            raise TypeError("Expected argument 'app_file_name' to be a str")
        pulumi.set(__self__, "app_file_name", app_file_name)
        if application_name and not isinstance(application_name, str):
            raise TypeError("Expected argument 'application_name' to be a str")
        pulumi.set(__self__, "application_name", application_name)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if comments and not isinstance(comments, str):
            raise TypeError("Expected argument 'comments' to be a str")
        pulumi.set(__self__, "comments", comments)
        if draft_package_path and not isinstance(draft_package_path, str):
            raise TypeError("Expected argument 'draft_package_path' to be a str")
        pulumi.set(__self__, "draft_package_path", draft_package_path)
        if edit_package and not isinstance(edit_package, bool):
            raise TypeError("Expected argument 'edit_package' to be a bool")
        pulumi.set(__self__, "edit_package", edit_package)
        if executable_launch_command and not isinstance(executable_launch_command, str):
            raise TypeError("Expected argument 'executable_launch_command' to be a str")
        pulumi.set(__self__, "executable_launch_command", executable_launch_command)
        if first_party_apps and not isinstance(first_party_apps, list):
            raise TypeError("Expected argument 'first_party_apps' to be a list")
        pulumi.set(__self__, "first_party_apps", first_party_apps)
        if flighting_ring and not isinstance(flighting_ring, str):
            raise TypeError("Expected argument 'flighting_ring' to be a str")
        pulumi.set(__self__, "flighting_ring", flighting_ring)
        if gallery_apps and not isinstance(gallery_apps, list):
            raise TypeError("Expected argument 'gallery_apps' to be a list")
        pulumi.set(__self__, "gallery_apps", gallery_apps)
        if highlighted_files and not isinstance(highlighted_files, list):
            raise TypeError("Expected argument 'highlighted_files' to be a list")
        pulumi.set(__self__, "highlighted_files", highlighted_files)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if inplace_upgrade_os_pair and not isinstance(inplace_upgrade_os_pair, dict):
            raise TypeError("Expected argument 'inplace_upgrade_os_pair' to be a dict")
        pulumi.set(__self__, "inplace_upgrade_os_pair", inplace_upgrade_os_pair)
        if intune_enrollment_metadata and not isinstance(intune_enrollment_metadata, dict):
            raise TypeError("Expected argument 'intune_enrollment_metadata' to be a dict")
        pulumi.set(__self__, "intune_enrollment_metadata", intune_enrollment_metadata)
        if intune_metadata and not isinstance(intune_metadata, dict):
            raise TypeError("Expected argument 'intune_metadata' to be a dict")
        pulumi.set(__self__, "intune_metadata", intune_metadata)
        if last_modified_time and not isinstance(last_modified_time, str):
            raise TypeError("Expected argument 'last_modified_time' to be a str")
        pulumi.set(__self__, "last_modified_time", last_modified_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if package_id and not isinstance(package_id, str):
            raise TypeError("Expected argument 'package_id' to be a str")
        pulumi.set(__self__, "package_id", package_id)
        if package_tags and not isinstance(package_tags, dict):
            raise TypeError("Expected argument 'package_tags' to be a dict")
        pulumi.set(__self__, "package_tags", package_tags)
        if process_name and not isinstance(process_name, str):
            raise TypeError("Expected argument 'process_name' to be a str")
        pulumi.set(__self__, "process_name", process_name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if source_type and not isinstance(source_type, str):
            raise TypeError("Expected argument 'source_type' to be a str")
        pulumi.set(__self__, "source_type", source_type)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tab_state and not isinstance(tab_state, dict):
            raise TypeError("Expected argument 'tab_state' to be a dict")
        pulumi.set(__self__, "tab_state", tab_state)
        if target_os_list and not isinstance(target_os_list, list):
            raise TypeError("Expected argument 'target_os_list' to be a list")
        pulumi.set(__self__, "target_os_list", target_os_list)
        if test_types and not isinstance(test_types, list):
            raise TypeError("Expected argument 'test_types' to be a list")
        pulumi.set(__self__, "test_types", test_types)
        if tests and not isinstance(tests, list):
            raise TypeError("Expected argument 'tests' to be a list")
        pulumi.set(__self__, "tests", tests)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if use_autofill and not isinstance(use_autofill, bool):
            raise TypeError("Expected argument 'use_autofill' to be a bool")
        pulumi.set(__self__, "use_autofill", use_autofill)
        if use_sample and not isinstance(use_sample, bool):
            raise TypeError("Expected argument 'use_sample' to be a bool")
        pulumi.set(__self__, "use_sample", use_sample)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)
        if working_path and not isinstance(working_path, str):
            raise TypeError("Expected argument 'working_path' to be a str")
        pulumi.set(__self__, "working_path", working_path)

    @property
    @pulumi.getter(name="appFileName")
    def app_file_name(self) -> Optional[builtins.str]:
        """
        The name of the app file.
        """
        return pulumi.get(self, "app_file_name")

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> Optional[builtins.str]:
        """
        Application name
        """
        return pulumi.get(self, "application_name")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def comments(self) -> Optional[builtins.str]:
        """
        Comments added by user.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter(name="draftPackagePath")
    def draft_package_path(self) -> builtins.str:
        """
        The relative path of the folder hosting package files.
        """
        return pulumi.get(self, "draft_package_path")

    @property
    @pulumi.getter(name="editPackage")
    def edit_package(self) -> Optional[builtins.bool]:
        """
        Specifies whether this draft package is used to edit a package.
        """
        return pulumi.get(self, "edit_package")

    @property
    @pulumi.getter(name="executableLaunchCommand")
    def executable_launch_command(self) -> Optional[builtins.str]:
        """
        The executable launch command for script auto-fill. Will be used to run the application.
        """
        return pulumi.get(self, "executable_launch_command")

    @property
    @pulumi.getter(name="firstPartyApps")
    def first_party_apps(self) -> Optional[Sequence['outputs.FirstPartyAppDefinitionResponse']]:
        """
        The list of first party applications to test along with user application.
        """
        return pulumi.get(self, "first_party_apps")

    @property
    @pulumi.getter(name="flightingRing")
    def flighting_ring(self) -> Optional[builtins.str]:
        """
        The flighting ring for feature update.
        """
        return pulumi.get(self, "flighting_ring")

    @property
    @pulumi.getter(name="galleryApps")
    def gallery_apps(self) -> Optional[Sequence['outputs.GalleryAppDefinitionResponse']]:
        """
        The list of gallery apps to test along with user application.
        """
        return pulumi.get(self, "gallery_apps")

    @property
    @pulumi.getter(name="highlightedFiles")
    def highlighted_files(self) -> Optional[Sequence['outputs.HighlightedFileResponse']]:
        """
        The highlight files in the package.
        """
        return pulumi.get(self, "highlighted_files")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="inplaceUpgradeOSPair")
    def inplace_upgrade_os_pair(self) -> Optional['outputs.InplaceUpgradeOSInfoResponse']:
        """
        Specifies the baseline os and target os for inplace upgrade.
        """
        return pulumi.get(self, "inplace_upgrade_os_pair")

    @property
    @pulumi.getter(name="intuneEnrollmentMetadata")
    def intune_enrollment_metadata(self) -> Optional['outputs.IntuneEnrollmentMetadataResponse']:
        """
        The metadata of Intune enrollment.
        """
        return pulumi.get(self, "intune_enrollment_metadata")

    @property
    @pulumi.getter(name="intuneMetadata")
    def intune_metadata(self) -> Optional['outputs.DraftPackageIntuneAppMetadataResponse']:
        """
        Metadata used to generate draft package folder and scripts.
        """
        return pulumi.get(self, "intune_metadata")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> builtins.str:
        """
        The UTC timestamp when the package was last modified.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="packageId")
    def package_id(self) -> Optional[builtins.str]:
        """
        Specifies the package id from which the draft package copied.
        """
        return pulumi.get(self, "package_id")

    @property
    @pulumi.getter(name="packageTags")
    def package_tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Tags of the package to be created.
        """
        return pulumi.get(self, "package_tags")

    @property
    @pulumi.getter(name="processName")
    def process_name(self) -> Optional[builtins.str]:
        """
        The process name for script auto-fill. Will be used to identify the application process.
        """
        return pulumi.get(self, "process_name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        The provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> Optional[builtins.str]:
        """
        The source type.
        """
        return pulumi.get(self, "source_type")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter(name="tabState")
    def tab_state(self) -> Optional['outputs.TabStateResponse']:
        """
        Tab state.
        """
        return pulumi.get(self, "tab_state")

    @property
    @pulumi.getter(name="targetOSList")
    def target_os_list(self) -> Optional[Sequence['outputs.TargetOSInfoResponse']]:
        """
        Specifies the target OSs of specific OS Update types.
        """
        return pulumi.get(self, "target_os_list")

    @property
    @pulumi.getter(name="testTypes")
    def test_types(self) -> Optional[Sequence[builtins.str]]:
        """
        OOB, functional or flow driven. Mapped to the data in 'tests' property.
        """
        return pulumi.get(self, "test_types")

    @property
    @pulumi.getter
    def tests(self) -> Optional[Sequence['outputs.TestResponse']]:
        """
        The detailed test information.
        """
        return pulumi.get(self, "tests")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="useAutofill")
    def use_autofill(self) -> Optional[builtins.bool]:
        """
        Indicates whether user choose to enable script auto-fill.
        """
        return pulumi.get(self, "use_autofill")

    @property
    @pulumi.getter(name="useSample")
    def use_sample(self) -> Optional[builtins.bool]:
        """
        Specifies whether a sample package should be used instead of the one uploaded by the user.
        """
        return pulumi.get(self, "use_sample")

    @property
    @pulumi.getter
    def version(self) -> Optional[builtins.str]:
        """
        Application version
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="workingPath")
    def working_path(self) -> builtins.str:
        """
        The relative path for a temporarily folder for package creation work.
        """
        return pulumi.get(self, "working_path")


class AwaitableGetDraftPackageResult(GetDraftPackageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDraftPackageResult(
            app_file_name=self.app_file_name,
            application_name=self.application_name,
            azure_api_version=self.azure_api_version,
            comments=self.comments,
            draft_package_path=self.draft_package_path,
            edit_package=self.edit_package,
            executable_launch_command=self.executable_launch_command,
            first_party_apps=self.first_party_apps,
            flighting_ring=self.flighting_ring,
            gallery_apps=self.gallery_apps,
            highlighted_files=self.highlighted_files,
            id=self.id,
            inplace_upgrade_os_pair=self.inplace_upgrade_os_pair,
            intune_enrollment_metadata=self.intune_enrollment_metadata,
            intune_metadata=self.intune_metadata,
            last_modified_time=self.last_modified_time,
            name=self.name,
            package_id=self.package_id,
            package_tags=self.package_tags,
            process_name=self.process_name,
            provisioning_state=self.provisioning_state,
            source_type=self.source_type,
            system_data=self.system_data,
            tab_state=self.tab_state,
            target_os_list=self.target_os_list,
            test_types=self.test_types,
            tests=self.tests,
            type=self.type,
            use_autofill=self.use_autofill,
            use_sample=self.use_sample,
            version=self.version,
            working_path=self.working_path)


def get_draft_package(draft_package_name: Optional[builtins.str] = None,
                      resource_group_name: Optional[builtins.str] = None,
                      test_base_account_name: Optional[builtins.str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDraftPackageResult:
    """
    Gets a Test Base Draft Package.

    Uses Azure REST API version 2023-11-01-preview.


    :param builtins.str draft_package_name: The resource name of the Test Base Draft Package.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str test_base_account_name: The resource name of the Test Base Account.
    """
    __args__ = dict()
    __args__['draftPackageName'] = draft_package_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['testBaseAccountName'] = test_base_account_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:testbase:getDraftPackage', __args__, opts=opts, typ=GetDraftPackageResult).value

    return AwaitableGetDraftPackageResult(
        app_file_name=pulumi.get(__ret__, 'app_file_name'),
        application_name=pulumi.get(__ret__, 'application_name'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        comments=pulumi.get(__ret__, 'comments'),
        draft_package_path=pulumi.get(__ret__, 'draft_package_path'),
        edit_package=pulumi.get(__ret__, 'edit_package'),
        executable_launch_command=pulumi.get(__ret__, 'executable_launch_command'),
        first_party_apps=pulumi.get(__ret__, 'first_party_apps'),
        flighting_ring=pulumi.get(__ret__, 'flighting_ring'),
        gallery_apps=pulumi.get(__ret__, 'gallery_apps'),
        highlighted_files=pulumi.get(__ret__, 'highlighted_files'),
        id=pulumi.get(__ret__, 'id'),
        inplace_upgrade_os_pair=pulumi.get(__ret__, 'inplace_upgrade_os_pair'),
        intune_enrollment_metadata=pulumi.get(__ret__, 'intune_enrollment_metadata'),
        intune_metadata=pulumi.get(__ret__, 'intune_metadata'),
        last_modified_time=pulumi.get(__ret__, 'last_modified_time'),
        name=pulumi.get(__ret__, 'name'),
        package_id=pulumi.get(__ret__, 'package_id'),
        package_tags=pulumi.get(__ret__, 'package_tags'),
        process_name=pulumi.get(__ret__, 'process_name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        source_type=pulumi.get(__ret__, 'source_type'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tab_state=pulumi.get(__ret__, 'tab_state'),
        target_os_list=pulumi.get(__ret__, 'target_os_list'),
        test_types=pulumi.get(__ret__, 'test_types'),
        tests=pulumi.get(__ret__, 'tests'),
        type=pulumi.get(__ret__, 'type'),
        use_autofill=pulumi.get(__ret__, 'use_autofill'),
        use_sample=pulumi.get(__ret__, 'use_sample'),
        version=pulumi.get(__ret__, 'version'),
        working_path=pulumi.get(__ret__, 'working_path'))
def get_draft_package_output(draft_package_name: Optional[pulumi.Input[builtins.str]] = None,
                             resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                             test_base_account_name: Optional[pulumi.Input[builtins.str]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDraftPackageResult]:
    """
    Gets a Test Base Draft Package.

    Uses Azure REST API version 2023-11-01-preview.


    :param builtins.str draft_package_name: The resource name of the Test Base Draft Package.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str test_base_account_name: The resource name of the Test Base Account.
    """
    __args__ = dict()
    __args__['draftPackageName'] = draft_package_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['testBaseAccountName'] = test_base_account_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:testbase:getDraftPackage', __args__, opts=opts, typ=GetDraftPackageResult)
    return __ret__.apply(lambda __response__: GetDraftPackageResult(
        app_file_name=pulumi.get(__response__, 'app_file_name'),
        application_name=pulumi.get(__response__, 'application_name'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        comments=pulumi.get(__response__, 'comments'),
        draft_package_path=pulumi.get(__response__, 'draft_package_path'),
        edit_package=pulumi.get(__response__, 'edit_package'),
        executable_launch_command=pulumi.get(__response__, 'executable_launch_command'),
        first_party_apps=pulumi.get(__response__, 'first_party_apps'),
        flighting_ring=pulumi.get(__response__, 'flighting_ring'),
        gallery_apps=pulumi.get(__response__, 'gallery_apps'),
        highlighted_files=pulumi.get(__response__, 'highlighted_files'),
        id=pulumi.get(__response__, 'id'),
        inplace_upgrade_os_pair=pulumi.get(__response__, 'inplace_upgrade_os_pair'),
        intune_enrollment_metadata=pulumi.get(__response__, 'intune_enrollment_metadata'),
        intune_metadata=pulumi.get(__response__, 'intune_metadata'),
        last_modified_time=pulumi.get(__response__, 'last_modified_time'),
        name=pulumi.get(__response__, 'name'),
        package_id=pulumi.get(__response__, 'package_id'),
        package_tags=pulumi.get(__response__, 'package_tags'),
        process_name=pulumi.get(__response__, 'process_name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        source_type=pulumi.get(__response__, 'source_type'),
        system_data=pulumi.get(__response__, 'system_data'),
        tab_state=pulumi.get(__response__, 'tab_state'),
        target_os_list=pulumi.get(__response__, 'target_os_list'),
        test_types=pulumi.get(__response__, 'test_types'),
        tests=pulumi.get(__response__, 'tests'),
        type=pulumi.get(__response__, 'type'),
        use_autofill=pulumi.get(__response__, 'use_autofill'),
        use_sample=pulumi.get(__response__, 'use_sample'),
        version=pulumi.get(__response__, 'version'),
        working_path=pulumi.get(__response__, 'working_path')))
