package minio

import (
	"fmt"
	"path"

	desc "github.com/iterum-provenance/iterum-go/descriptors"
	"github.com/iterum-provenance/iterum-go/env"
	"github.com/iterum-provenance/iterum-go/util"
	"github.com/prometheus/common/log"
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

// GetConfigFile gets the file associated with filename from the minioStorage
func (config Config) GetConfigFile(filename string) (localFile desc.LocalFileDesc, err error) {
	descriptor := desc.RemoteFileDesc{
		Bucket:     configBucket,
		Name:       filename,
		RemotePath: path.Join(configPrefix, filename),
	}
	if env.ProcessConfigPath == env.DataVolumePath {
		log.Fatalf("EnvironmentError: '%v' is not a valid value for ITERUM_CONFIG_PATH", env.ProcessConfigPath)
	}
	return config.GetFile(descriptor, env.ProcessConfigPath)
}
