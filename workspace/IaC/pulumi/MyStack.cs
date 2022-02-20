using Pulumi;
using Pulumi.AzureNative.Resources;
using Pulumi.AzureNative.Storage;
using Pulumi.AzureNative.Storage.Inputs;
using Snowflake = Pulumi.Snowflake;

class MyStack : Stack
{
    public MyStack()
    {
        var resourceGroup = new ResourceGroup("glresourceGroup", new ResourceGroupArgs
        {
            ResourceGroupName = "glresourceGroup",
            Location = "eastus",
        });

        var storageAccount = new StorageAccount("gllearningpulumi", new StorageAccountArgs
        {
            AccountName = "gllearningpulumi",
            Kind = "BlockBlobStorage",
            Location = "eastus",
            ResourceGroupName = resourceGroup.Name,
            Sku = new SkuArgs
            {
                Name = "Standard_LRS",
            },
        });

        var blobContainer = new BlobContainer("snowdatablob", new BlobContainerArgs
        {
            AccountName = storageAccount.Name,
            ContainerName = "snowdatablob",
            ResourceGroupName = resourceGroup.Name,
        });
   
        var rawSourceTable = new Snowflake.Table("RAW_SOURCE", new Snowflake.TableArgs
        {
            Name = "RAW_SOURCE",
            Database = "BLOB_DB",
            Schema = "PUBLIC",

            Columns =
            {
                new Snowflake.Inputs.TableColumnArgs
                {
                    Name = "SRC",
                    Type = "VARIANT",
                    Nullable = false,
                },
                new Snowflake.Inputs.TableColumnArgs
                {
                    Name     = "SHA",
                    Type     = "VARCHAR(100)",
                    Nullable = false,
                },
            },
        });


        var usersTransformedTable = new Snowflake.Table("USERS_TRANSFORMED", new Snowflake.TableArgs
        {
            Name = "USERS_TRANSFORMED",
            Database = "BLOB_DB",
            Schema = "PUBLIC",

            Columns =
            {
                new Snowflake.Inputs.TableColumnArgs
                {
                    Name     = "SHA",
                    Type     = "VARCHAR(100)",
                    Nullable = false,
                },
                new Snowflake.Inputs.TableColumnArgs
                {
                    Name     = "FullName",
                    Type     = "VARCHAR(46)",
                    Nullable = false,
                },
                new Snowflake.Inputs.TableColumnArgs
                {
                    Name     = "Age",
                    Type     = "NUMBER(38,0)",
                    Nullable = false,
                },
                new Snowflake.Inputs.TableColumnArgs
                {
                    Name     = "City",
                    Type     = "VARCHAR(30)",
                    Nullable = false,
                },
                new Snowflake.Inputs.TableColumnArgs
                {
                    Name     = "Organization",
                    Type     = "VARCHAR(30)",
                    Nullable = false,
                },
                new Snowflake.Inputs.TableColumnArgs
                {
                    Name     = "Department",
                    Type     = "VARCHAR(30)",
                    Nullable = false,
                },
                new Snowflake.Inputs.TableColumnArgs
                {
                    Name     = "Subdepartment",
                    Type     = "VARCHAR(30)",
                    Nullable = false,
                },
                new Snowflake.Inputs.TableColumnArgs
                {
                    Name     = "Team",
                    Type     = "VARCHAR(30)",
                    Nullable = false,
                },
            },
        });

    var azureStageSas = new Snowflake.Stage("azure_stage_sas", new Snowflake.StageArgs{
            Name = "azure_stage_sas",
            Url  = $"az_container_url",
            Database = "BLOB_DB",
            Schema   = "PUBLIC",
            Credentials="AZURE_SAS_TOKEN='$az_container_url'",
        });

    var copyFromAzureTask = new Snowflake.Task("copy_data_from_azure", new Snowflake.TaskArgs
        {
            Name ="copy_data_from_azure",
            Database = "BLOB_DB",
            Schema = "PUBLIC",
            Warehouse = "TEST_WH",
            Schedule = "10 MINUTE",
            SqlStatement = "copy into RAW_SOURCE(SRC, SHA) from (select $1, SHA2($1) from @azure_stage_sas)",
        });   
    

    var insertIntoUsersTransformedTask = new Snowflake.Task("insert_data_into_users_transformed", new Snowflake.TaskArgs
        {
            Name = "insert_data_into_users_transformed",
            Database = "BLOB_DB",
            Schema = "PUBLIC",
            Warehouse = "TEST_WH",
            Schedule = "10 MINUTE",
            SqlStatement = "insert into USERS_TRANSFORMED select SHA, src:FULLNAME, src:AGE, src:CITY, src:ORGANIZATION, src:DEPARTMENT, src:SUBDEPARTMENT, src:TEAM from RAW_SOURCE where not exists (select 'x' from USERS_TRANSFORMED where USERS_TRANSFORMED.SHA = RAW_SOURCE.SHA);",
        });   
    }

}
