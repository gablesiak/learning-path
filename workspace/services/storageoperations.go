package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/Azure/azure-pipeline-go/pipeline"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/gablesiak/datatypes"
	"github.com/google/uuid"
)

type StorageAccess struct {
	StorageAccount string
	AccessKey      string
}

func SetStorageAccess() StorageAccess {
	azureAccess := StorageAccess{}

	azureAccess.StorageAccount = os.Getenv("AZURE_STORAGE_ACCOUNT")
	azureAccess.AccessKey = os.Getenv("AZURE_STORAGE_ACCESS_KEY")

	return azureAccess
}

func createPipeline(azureAccess StorageAccess) pipeline.Pipeline {

	credential, err := azblob.NewSharedKeyCredential(azureAccess.StorageAccount, azureAccess.AccessKey)
	if err != nil {
		fmt.Print(err)
	}
	blobPipeline := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	return blobPipeline
}

func createStorageURL(azureAccess StorageAccess, containerName string) *url.URL {
	URL, err := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net/%s", azureAccess.StorageAccount, containerName))
	if err != nil {
		fmt.Print(err)
	}

	return URL
}

func getContainerURL(azureAccess StorageAccess, containerName string) azblob.ContainerURL {
	blobPipeline := createPipeline(azureAccess)
	URL := createStorageURL(azureAccess, containerName)

	containerURL := azblob.NewContainerURL(*URL, blobPipeline)

	return containerURL
}


func UploadFile(azureAccess StorageAccess, inputData datatypes.InputUserData) {
	uuidString := uuid.NewString()
	containerName := "rel-project"
	containerURL := getContainerURL(azureAccess, containerName)
	outputData := GenerateOutputStruct(inputData)

	multipleOutputData, err := json.MarshalIndent(outputData, "", " ")
	if err != nil {
		fmt.Print(err)
	}

	ctx := context.Background()
	_, err = containerURL.Create(ctx, azblob.Metadata{}, azblob.PublicAccessNone)
	if err != nil {
		fmt.Print(err)
	}

	blobURL := containerURL.NewBlockBlobURL(uuidString + ".json")

	_, err = azblob.UploadBufferToBlockBlob(ctx, multipleOutputData, blobURL, azblob.UploadToBlockBlobOptions{})
	if err != nil {
		fmt.Print(err)
	}
}
