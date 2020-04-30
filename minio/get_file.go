package minio

import (
	"fmt"

	desc "github.com/iterum-provenance/iterum-go/descriptors"
	"github.com/iterum-provenance/iterum-go/util"
)

// GetFile retrieves the file associated with the RemoteFileDesc onto local disk
// It does not ensure any existing connection neither the bucket. This is the responsibility of the user
// targetFolder is the folder in which to store the data
func (config Config) GetFile(descriptor desc.RemoteFileDesc, targetFolder string) (localFile desc.LocalFileDesc, err error) {
	defer util.ReturnErrOnPanic(&err)()

	if !config.IsConnected() {
		return localFile, fmt.Errorf("Minio client not initialized, cannot get file")
	}
	err = config.EnsureBucket(descriptor.Bucket, 10)
	util.PanicIfErr(err, "")

	localFilePath := descriptor.ToLocalPath(targetFolder)
	err = config.Client.FGetObject(descriptor.Bucket, descriptor.RemotePath, localFilePath, config.GetOptions)
	util.PanicIfErr(err, fmt.Sprintf("Download failed due to '%v'\n", err))

	localFile = desc.LocalFileDesc{
		Name:      descriptor.Name,
		LocalPath: localFilePath,
	}

	return
}
