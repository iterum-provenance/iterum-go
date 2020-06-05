package minio

import (
	"fmt"
	"io"

	desc "github.com/iterum-provenance/iterum-go/descriptors"
	"github.com/iterum-provenance/iterum-go/env"
	"github.com/iterum-provenance/iterum-go/util"
)

const (
	dataPrefix   string = "iterum_data"
	configPrefix string = "iterum_config"
)

var configBucket string = env.PipelineHash + "-ITERUM-CONFIG"

func (config Config) _putFile(localFile desc.LocalFileDesc, putMechanism func() (err error), prefix string) (remoteFile desc.RemoteFileDesc, err error) {
	defer util.ReturnErrOnPanic(&err)()

	if !config.IsConnected() {
		return remoteFile, fmt.Errorf("Minio client not initialized, cannot put file")
	}

	err = config.EnsureBucket(config.TargetBucket, 10)
	util.PanicIfErr(err, "")

	err = putMechanism()
	util.PanicIfErr(err, "")

	remoteFile = desc.RemoteFileDesc{
		Name:       localFile.Name,
		RemotePath: localFile.ToRemotePath(prefix),
		Bucket:     config.TargetBucket,
	}

	return
}

// PutFileFromReader send the data associated with a fileHandler into the minioStorage
// It assumes that the target bucket exists and access is granted to it ia the config
// remotePath is the target remote path to store to. fileName is used in the RemoteFileDesc
func (config Config) PutFileFromReader(fileHandle io.ReadCloser, contentSize int64, localFile desc.LocalFileDesc) (remoteFile desc.RemoteFileDesc, err error) {
	putMechanism := func() (err error) {
		defer fileHandle.Close()
		_, err = config.Client.PutObject(config.TargetBucket, localFile.ToRemotePath(dataPrefix), fileHandle, contentSize, config.PutOptions)
		return
	}
	return config._putFile(localFile, putMechanism, dataPrefix)
}

// PutFile send the file associated with localPath into the minioStorage
// It ensures that the target bucket exists and otherwise creates it
// filePath is the target remote path
func (config Config) PutFile(localFile desc.LocalFileDesc) (remoteFile desc.RemoteFileDesc, err error) {
	putMechanism := func() (err error) {
		_, err = config.Client.FPutObject(config.TargetBucket, localFile.ToRemotePath(dataPrefix), localFile.LocalPath, config.PutOptions)
		return
	}
	return config._putFile(localFile, putMechanism, dataPrefix)
}

// PutConfigFile sends the file associated with localPath into the minioStorage
// in the config bucket of this pipeline run
func (config Config) PutConfigFile(localFile desc.LocalFileDesc) (remoteFile desc.RemoteFileDesc, err error) {
	config.TargetBucket = configBucket
	putMechanism := func() (err error) {
		_, err = config.Client.FPutObject(config.TargetBucket, localFile.ToRemotePath(configPrefix), localFile.LocalPath, config.PutOptions)
		return
	}
	return config._putFile(localFile, putMechanism, configPrefix)
}
