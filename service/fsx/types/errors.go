// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// An Active Directory error.
type ActiveDirectoryError struct {
	Message *string

	ActiveDirectoryId *string
	Type              ActiveDirectoryErrorType

	noSmithyDocumentSerde
}

func (e *ActiveDirectoryError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ActiveDirectoryError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ActiveDirectoryError) ErrorCode() string             { return "ActiveDirectoryError" }
func (e *ActiveDirectoryError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You can't delete a backup while it's being copied.
type BackupBeingCopied struct {
	Message *string

	BackupId *string

	noSmithyDocumentSerde
}

func (e *BackupBeingCopied) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BackupBeingCopied) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BackupBeingCopied) ErrorCode() string             { return "BackupBeingCopied" }
func (e *BackupBeingCopied) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Another backup is already under way. Wait for completion before initiating
// additional backups of this file system.
type BackupInProgress struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *BackupInProgress) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BackupInProgress) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BackupInProgress) ErrorCode() string             { return "BackupInProgress" }
func (e *BackupInProgress) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// No Amazon FSx backups were found based upon the supplied parameters.
type BackupNotFound struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *BackupNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BackupNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BackupNotFound) ErrorCode() string             { return "BackupNotFound" }
func (e *BackupNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You can't delete a backup while it's being used to restore a file system.
type BackupRestoring struct {
	Message *string

	FileSystemId *string

	noSmithyDocumentSerde
}

func (e *BackupRestoring) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BackupRestoring) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BackupRestoring) ErrorCode() string             { return "BackupRestoring" }
func (e *BackupRestoring) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A generic error indicating a failure with a client request.
type BadRequest struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *BadRequest) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadRequest) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadRequest) ErrorCode() string             { return "BadRequest" }
func (e *BadRequest) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// No data repository associations were found based upon the supplied parameters.
type DataRepositoryAssociationNotFound struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DataRepositoryAssociationNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DataRepositoryAssociationNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DataRepositoryAssociationNotFound) ErrorCode() string {
	return "DataRepositoryAssociationNotFound"
}
func (e *DataRepositoryAssociationNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The data repository task could not be canceled because the task has already
// ended.
type DataRepositoryTaskEnded struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DataRepositoryTaskEnded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DataRepositoryTaskEnded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DataRepositoryTaskEnded) ErrorCode() string             { return "DataRepositoryTaskEnded" }
func (e *DataRepositoryTaskEnded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An existing data repository task is currently executing on the file system. Wait
// until the existing task has completed, then create the new task.
type DataRepositoryTaskExecuting struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DataRepositoryTaskExecuting) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DataRepositoryTaskExecuting) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DataRepositoryTaskExecuting) ErrorCode() string             { return "DataRepositoryTaskExecuting" }
func (e *DataRepositoryTaskExecuting) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The data repository task or tasks you specified could not be found.
type DataRepositoryTaskNotFound struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DataRepositoryTaskNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DataRepositoryTaskNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DataRepositoryTaskNotFound) ErrorCode() string             { return "DataRepositoryTaskNotFound" }
func (e *DataRepositoryTaskNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// No caches were found based upon supplied parameters.
type FileCacheNotFound struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *FileCacheNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FileCacheNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FileCacheNotFound) ErrorCode() string             { return "FileCacheNotFound" }
func (e *FileCacheNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// No Amazon FSx file systems were found based upon supplied parameters.
type FileSystemNotFound struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *FileSystemNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FileSystemNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FileSystemNotFound) ErrorCode() string             { return "FileSystemNotFound" }
func (e *FileSystemNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The error returned when a second request is received with the same client
// request token but different parameters settings. A client request token should
// always uniquely identify a single request.
type IncompatibleParameterError struct {
	Message *string

	Parameter *string

	noSmithyDocumentSerde
}

func (e *IncompatibleParameterError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IncompatibleParameterError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IncompatibleParameterError) ErrorCode() string             { return "IncompatibleParameterError" }
func (e *IncompatibleParameterError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Amazon FSx doesn't support Multi-AZ Windows File Server copy backup in the
// destination Region, so the copied backup can't be restored.
type IncompatibleRegionForMultiAZ struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *IncompatibleRegionForMultiAZ) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IncompatibleRegionForMultiAZ) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IncompatibleRegionForMultiAZ) ErrorCode() string             { return "IncompatibleRegionForMultiAZ" }
func (e *IncompatibleRegionForMultiAZ) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A generic error indicating a server-side failure.
type InternalServerError struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerError) ErrorCode() string             { return "InternalServerError" }
func (e *InternalServerError) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// You have filtered the response to a data repository type that is not supported.
type InvalidDataRepositoryType struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidDataRepositoryType) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDataRepositoryType) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDataRepositoryType) ErrorCode() string             { return "InvalidDataRepositoryType" }
func (e *InvalidDataRepositoryType) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The Key Management Service (KMS) key of the destination backup is not valid.
type InvalidDestinationKmsKey struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidDestinationKmsKey) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDestinationKmsKey) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDestinationKmsKey) ErrorCode() string             { return "InvalidDestinationKmsKey" }
func (e *InvalidDestinationKmsKey) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The path provided for data repository export isn't valid.
type InvalidExportPath struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidExportPath) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidExportPath) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidExportPath) ErrorCode() string             { return "InvalidExportPath" }
func (e *InvalidExportPath) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The path provided for data repository import isn't valid.
type InvalidImportPath struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidImportPath) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidImportPath) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidImportPath) ErrorCode() string             { return "InvalidImportPath" }
func (e *InvalidImportPath) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more network settings specified in the request are invalid.
type InvalidNetworkSettings struct {
	Message *string

	InvalidSubnetId        *string
	InvalidSecurityGroupId *string
	InvalidRouteTableId    *string

	noSmithyDocumentSerde
}

func (e *InvalidNetworkSettings) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNetworkSettings) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNetworkSettings) ErrorCode() string             { return "InvalidNetworkSettings" }
func (e *InvalidNetworkSettings) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An invalid value for PerUnitStorageThroughput was provided. Please create your
// file system again, using a valid value.
type InvalidPerUnitStorageThroughput struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidPerUnitStorageThroughput) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidPerUnitStorageThroughput) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidPerUnitStorageThroughput) ErrorCode() string {
	return "InvalidPerUnitStorageThroughput"
}
func (e *InvalidPerUnitStorageThroughput) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The Region provided for SourceRegion is not valid or is in a different Amazon
// Web Services partition.
type InvalidRegion struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidRegion) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRegion) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRegion) ErrorCode() string             { return "InvalidRegion" }
func (e *InvalidRegion) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The Key Management Service (KMS) key of the source backup is not valid.
type InvalidSourceKmsKey struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidSourceKmsKey) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSourceKmsKey) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSourceKmsKey) ErrorCode() string             { return "InvalidSourceKmsKey" }
func (e *InvalidSourceKmsKey) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A cache configuration is required for this operation.
type MissingFileCacheConfiguration struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *MissingFileCacheConfiguration) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MissingFileCacheConfiguration) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MissingFileCacheConfiguration) ErrorCode() string             { return "MissingFileCacheConfiguration" }
func (e *MissingFileCacheConfiguration) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A file system configuration is required for this operation.
type MissingFileSystemConfiguration struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *MissingFileSystemConfiguration) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MissingFileSystemConfiguration) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MissingFileSystemConfiguration) ErrorCode() string             { return "MissingFileSystemConfiguration" }
func (e *MissingFileSystemConfiguration) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A volume configuration is required for this operation.
type MissingVolumeConfiguration struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *MissingVolumeConfiguration) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MissingVolumeConfiguration) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MissingVolumeConfiguration) ErrorCode() string             { return "MissingVolumeConfiguration" }
func (e *MissingVolumeConfiguration) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource specified for the tagging operation is not a resource type owned by
// Amazon FSx. Use the API of the relevant service to perform the operation.
type NotServiceResourceError struct {
	Message *string

	ResourceARN *string

	noSmithyDocumentSerde
}

func (e *NotServiceResourceError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotServiceResourceError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotServiceResourceError) ErrorCode() string             { return "NotServiceResourceError" }
func (e *NotServiceResourceError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource specified does not support tagging.
type ResourceDoesNotSupportTagging struct {
	Message *string

	ResourceARN *string

	noSmithyDocumentSerde
}

func (e *ResourceDoesNotSupportTagging) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceDoesNotSupportTagging) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceDoesNotSupportTagging) ErrorCode() string             { return "ResourceDoesNotSupportTagging" }
func (e *ResourceDoesNotSupportTagging) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource specified by the Amazon Resource Name (ARN) can't be found.
type ResourceNotFound struct {
	Message *string

	ResourceARN *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFound) ErrorCode() string             { return "ResourceNotFound" }
func (e *ResourceNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An error indicating that a particular service limit was exceeded. You can
// increase some service limits by contacting Amazon Web Services Support.
type ServiceLimitExceeded struct {
	Message *string

	Limit ServiceLimit

	noSmithyDocumentSerde
}

func (e *ServiceLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceLimitExceeded) ErrorCode() string             { return "ServiceLimitExceeded" }
func (e *ServiceLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// No Amazon FSx snapshots were found based on the supplied parameters.
type SnapshotNotFound struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *SnapshotNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SnapshotNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SnapshotNotFound) ErrorCode() string             { return "SnapshotNotFound" }
func (e *SnapshotNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because the lifecycle status of the source backup isn't
// AVAILABLE.
type SourceBackupUnavailable struct {
	Message *string

	BackupId *string

	noSmithyDocumentSerde
}

func (e *SourceBackupUnavailable) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SourceBackupUnavailable) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SourceBackupUnavailable) ErrorCode() string             { return "SourceBackupUnavailable" }
func (e *SourceBackupUnavailable) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// No FSx for ONTAP SVMs were found based upon the supplied parameters.
type StorageVirtualMachineNotFound struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *StorageVirtualMachineNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StorageVirtualMachineNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StorageVirtualMachineNotFound) ErrorCode() string             { return "StorageVirtualMachineNotFound" }
func (e *StorageVirtualMachineNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested operation is not supported for this resource or API.
type UnsupportedOperation struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UnsupportedOperation) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedOperation) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedOperation) ErrorCode() string             { return "UnsupportedOperation" }
func (e *UnsupportedOperation) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// No Amazon FSx volumes were found based upon the supplied parameters.
type VolumeNotFound struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *VolumeNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *VolumeNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *VolumeNotFound) ErrorCode() string             { return "VolumeNotFound" }
func (e *VolumeNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
