# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing
# Export this package's modules as members:
from ._enums import *
from .get_billing_info import *
from .get_monitor import *
from .get_monitor_default_key import *
from .get_monitored_subscription import *
from .list_monitor_api_keys import *
from .list_monitor_hosts import *
from .list_monitor_linked_resources import *
from .list_monitor_monitored_resources import *
from .monitor import *
from .monitored_subscription import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_azure_native.datadog.v20220601 as __v20220601
    v20220601 = __v20220601
    import pulumi_azure_native.datadog.v20220801 as __v20220801
    v20220801 = __v20220801
    import pulumi_azure_native.datadog.v20230101 as __v20230101
    v20230101 = __v20230101
    import pulumi_azure_native.datadog.v20230707 as __v20230707
    v20230707 = __v20230707
    import pulumi_azure_native.datadog.v20231020 as __v20231020
    v20231020 = __v20231020
else:
    v20220601 = _utilities.lazy_import('pulumi_azure_native.datadog.v20220601')
    v20220801 = _utilities.lazy_import('pulumi_azure_native.datadog.v20220801')
    v20230101 = _utilities.lazy_import('pulumi_azure_native.datadog.v20230101')
    v20230707 = _utilities.lazy_import('pulumi_azure_native.datadog.v20230707')
    v20231020 = _utilities.lazy_import('pulumi_azure_native.datadog.v20231020')

