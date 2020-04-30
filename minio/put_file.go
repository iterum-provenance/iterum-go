package minio

import (
	"fmt"
	"io"

	desc "github.com/iterum-provenance/iterum-go/descriptors"
	"github.com/iterum-provenance/iterum-go/util"
)

func (config Config) _putFile(localFile desc.LocalFileDesc, putMechanism func() (err error)) (remoteFile desc.RemoteFileDesc, err error) {
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
		RemotePath: localFile.ToRemotePath("iterum"),
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
		_, err = config.Client.PutObject(config.TargetBucket, localFile.ToRemotePath("iterum"), fileHandle, contentSize, config.PutOptions)
		return
	}
	return config._putFile(localFile, putMechanism)
}

// PutFile send the file associated with localPath into the minioStorage
// It ensures that the target bucket exists and otherwise creates it
// filePath is the target remote path
func (config Config) PutFile(localFile desc.LocalFileDesc) (remoteFile desc.RemoteFileDesc, err error) {
	putMechanism := func() (err error) {
		_, err = config.Client.FPutObject(config.TargetBucket, localFile.ToRemotePath("iterum"), localFile.LocalPath, config.PutOptions)
		return
	}
	return config._putFile(localFile, putMechanism)
}
