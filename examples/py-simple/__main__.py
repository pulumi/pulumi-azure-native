import pulumi
import pulumi_random
from pulumi_azure_nextgen.storage import latest as storage
from pulumi_azure_nextgen.resources import latest as resources

random_string = pulumi_random.RandomString('random',
    length=12,
    special=False,
    upper=False)

# Create an Azure Resource Group
resource_group = resources.ResourceGroup('resource_group',
    resource_group_name=random_string.result,
    location='westus')

# Create an Azure resource (Storage Account)
account = storage.StorageAccount('sa',
    account_name=random_string.result,
    resource_group_name=resource_group.name,
    location=resource_group.location,
    sku=storage.SkuArgs(
        name='Standard_LRS',
    ),
    kind='StorageV2')

primary_key = pulumi.Output.all(resource_group.name, account.name) \
    .apply(lambda args: storage.list_storage_account_keys(
        resource_group_name=args[0],
        account_name=args[1]
    )).apply(lambda accountKeys: accountKeys.keys[0].value)

pulumi.export("primary_storage_key", primary_key)
