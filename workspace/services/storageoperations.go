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

func CreatePipeline(azureAccess StorageAccess) pipeline.Pipeline {

	credential, err := azblob.NewSharedKeyCredential(azureAccess.StorageAccount, azureAccess.AccessKey)
	if err != nil {
		fmt.Print(err)
	}
	blobPipeline := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	return blobPipeline
}

func CreateStorageURL(azureAccess StorageAccess, containerName string) *url.URL {
	URL, err := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net/%s", azureAccess.StorageAccount, containerName))
	if err != nil {
		fmt.Print(err)
	}

	return URL
}

func GetContainerURL(azureAccess StorageAccess, containerName string) azblob.ContainerURL {
	blobPipeline := CreatePipeline(azureAccess)
	URL := CreateStorageURL(azureAccess, containerName)

	containerURL := azblob.NewContainerURL(*URL, blobPipeline)

	return containerURL
}

func CreateContainerName() string {
	uuidString := uuid.NewString()

	containerName := fmt.Sprintf("c-%s", uuidString)

	return containerName
}

func UploadFile(stcr StorageAccess, inputData datatypes.InputUserData) {
	uuidString := uuid.NewString()
	containerName := CreateContainerName()
	containerURL := GetContainerURL(stcr, containerName)
	outputData := GenerateOutputStruct(inputData)

	multipleOutputData, err := json.MarshalIndent(outputData, "", " ")

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

func GenerateOutputStruct(inputData datatypes.InputUserData) {
	panic("unimplemented")
}
