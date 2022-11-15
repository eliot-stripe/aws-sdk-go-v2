// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

type Application struct {

	//
	AppRegistryArn *string

	//
	Arn *string

	//
	Components []string

	//
	Id *string

	//
	LastUpdated *time.Time

	//
	Status ApplicationStatus

	//
	StatusMessage *string

	//
	Type ApplicationType

	noSmithyDocumentSerde
}

type ApplicationCredential struct {

	//
	//
	// This member is required.
	CredentialType CredentialType

	//
	//
	// This member is required.
	DatabaseName *string

	//
	//
	// This member is required.
	SecretId *string

	noSmithyDocumentSerde
}

type ApplicationSummary struct {

	//
	Arn *string

	//
	Id *string

	//
	Tags map[string]string

	//
	Type ApplicationType

	noSmithyDocumentSerde
}

type Component struct {

	//
	ApplicationId *string

	//
	ComponentId *string

	//
	ComponentType ComponentType

	//
	Databases []string

	//
	Hosts []Host

	//
	LastUpdated *time.Time

	//
	PrimaryHost *string

	//
	Status ComponentStatus

	noSmithyDocumentSerde
}

type ComponentSummary struct {

	//
	ApplicationId *string

	//
	ComponentId *string

	//
	ComponentType ComponentType

	//
	Tags map[string]string

	noSmithyDocumentSerde
}

type Database struct {

	//
	ApplicationId *string

	//
	Arn *string

	//
	ComponentId *string

	//
	Credentials []ApplicationCredential

	//
	DatabaseId *string

	//
	DatabaseName *string

	//
	DatabaseType DatabaseType

	//
	LastUpdated *time.Time

	//
	PrimaryHost *string

	//
	SQLPort *int32

	//
	Status DatabaseStatus

	noSmithyDocumentSerde
}

type DatabaseSummary struct {

	//
	ApplicationId *string

	//
	Arn *string

	//
	ComponentId *string

	//
	DatabaseId *string

	//
	DatabaseType DatabaseType

	//
	Tags map[string]string

	noSmithyDocumentSerde
}

type Host struct {

	//
	HostIp *string

	//
	HostName *string

	//
	HostRole HostRole

	//
	InstanceId *string

	noSmithyDocumentSerde
}

type Operation struct {

	//
	EndTime *time.Time

	//
	Id *string

	//
	LastUpdatedTime *time.Time

	//
	Properties map[string]*string

	//
	ResourceArn *string

	//
	ResourceId *string

	//
	ResourceType *string

	//
	StartTime *time.Time

	//
	Status OperationStatus

	//
	StatusMessage *string

	//
	Type *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
