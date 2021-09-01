// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an X.509 certificate using the specified certificate signing request.
// Note: The CSR must include a public key that is either an RSA key with a length
// of at least 2048 bits or an ECC key from NIST P-256 or NIST P-384 curves. Note:
// Reusing the same certificate signing request (CSR) results in a distinct
// certificate. Requires permission to access the CreateCertificateFromCsr
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action. You can create multiple certificates in a batch by creating a directory,
// copying multiple .csr files into that directory, and then specifying that
// directory on the command line. The following commands show how to create a batch
// of certificates given a batch of CSRs. Assuming a set of CSRs are located inside
// of the directory my-csr-directory: On Linux and OS X, the command is: $ ls
// my-csr-directory/ | xargs -I {} aws iot create-certificate-from-csr
// --certificate-signing-request file://my-csr-directory/{} This command lists all
// of the CSRs in my-csr-directory and pipes each CSR file name to the aws iot
// create-certificate-from-csr Amazon Web Services CLI command to create a
// certificate for the corresponding CSR. The aws iot create-certificate-from-csr
// part of the command can also be run in parallel to speed up the certificate
// creation process: $ ls my-csr-directory/ | xargs -P 10 -I {} aws iot
// create-certificate-from-csr --certificate-signing-request
// file://my-csr-directory/{} On Windows PowerShell, the command to create
// certificates for all CSRs in my-csr-directory is: > ls -Name my-csr-directory |
// %{aws iot create-certificate-from-csr --certificate-signing-request
// file://my-csr-directory/$_} On a Windows command prompt, the command to create
// certificates for all CSRs in my-csr-directory is: > forfiles /p my-csr-directory
// /c "cmd /c aws iot create-certificate-from-csr --certificate-signing-request
// file://@path"
func (c *Client) CreateCertificateFromCsr(ctx context.Context, params *CreateCertificateFromCsrInput, optFns ...func(*Options)) (*CreateCertificateFromCsrOutput, error) {
	if params == nil {
		params = &CreateCertificateFromCsrInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCertificateFromCsr", params, optFns, c.addOperationCreateCertificateFromCsrMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCertificateFromCsrOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the CreateCertificateFromCsr operation.
type CreateCertificateFromCsrInput struct {

	// The certificate signing request (CSR).
	//
	// This member is required.
	CertificateSigningRequest *string

	// Specifies whether the certificate is active.
	SetAsActive bool

	noSmithyDocumentSerde
}

// The output from the CreateCertificateFromCsr operation.
type CreateCertificateFromCsrOutput struct {

	// The Amazon Resource Name (ARN) of the certificate. You can use the ARN as a
	// principal for policy operations.
	CertificateArn *string

	// The ID of the certificate. Certificate management operations only take a
	// certificateId.
	CertificateId *string

	// The certificate data, in PEM format.
	CertificatePem *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCertificateFromCsrMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCertificateFromCsr{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCertificateFromCsr{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateCertificateFromCsrValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCertificateFromCsr(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateCertificateFromCsr(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateCertificateFromCsr",
	}
}
