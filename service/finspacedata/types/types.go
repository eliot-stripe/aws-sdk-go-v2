// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A changeset is unit of data in a dataset.
type ChangesetInfo struct {

	// Change type indicates how a changeset is applied to a dataset.
	//
	// * REPLACE -
	// Changeset is considered as a replacement to all prior loaded changesets.
	//
	// *
	// APPEND - Changeset is considered as an addition to the end of all prior loaded
	// changesets.
	//
	// * MODIFY - Changeset is considered as a replacement to a specific
	// prior ingested changeset.
	ChangeType ChangeType

	// The ARN identifier of the changeset.
	ChangesetArn *string

	// Tags associated with the changeset.
	ChangesetLabels map[string]string

	// The timestamp at which the changeset was created in FinSpace.
	CreateTimestamp *time.Time

	// The unique identifier for the FinSpace dataset in which the changeset is
	// created.
	DatasetId *string

	// The structure with error messages.
	ErrorInfo *ErrorInfo

	// Structure of the source file(s).
	FormatParams map[string]string

	// Format type of the input files loaded into the changeset.
	FormatType FormatType

	// Unique identifier for a changeset.
	Id *string

	// Source path from which the files to create the changeset are sourced.
	SourceParams map[string]string

	// Type of the data source from which the files to create the changeset are
	// sourced.
	//
	// * S3 - Amazon S3.
	SourceType SourceType

	// The status of changeset creation operation.
	Status ChangesetStatus

	// Unique identifier of the changeset that is updated a changeset.
	UpdatedByChangesetId *string

	// Unique identifier of the changeset that is updated.
	UpdatesChangesetId *string

	noSmithyDocumentSerde
}

// Set short term API credentials.
type Credentials struct {

	// The access key identifier.
	AccessKeyId *string

	// The access key.
	SecretAccessKey *string

	// The session token.
	SessionToken *string

	noSmithyDocumentSerde
}

// Error message.
type ErrorInfo struct {

	// The category of the error.
	//
	// * VALIDATION -The inputs to this request are
	// invalid.
	//
	// * SERVICE_QUOTA_EXCEEDED - Service quotas have been exceeded. Please
	// contact AWS support to increase quotas.
	//
	// * ACCESS_DENIED - Missing required
	// permission to perform this request.
	//
	// * RESOURCE_NOT_FOUND - One or more inputs
	// to this request were not found.
	//
	// * THROTTLING - The system temporarily lacks
	// sufficient resources to process the request.
	//
	// * INTERNAL_SERVICE_EXCEPTION - An
	// internal service error has occurred.
	//
	// * CANCELLED - A user recoverable error has
	// occurred.
	ErrorCategory ErrorCategory

	// The text of the error message.
	ErrorMessage *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
