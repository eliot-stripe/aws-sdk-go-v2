// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a parallel data resource in Amazon Translate by importing an input file
// from Amazon S3. Parallel data files contain examples that show how you want
// segments of text to be translated. By adding parallel data, you can influence
// the style, tone, and word choice in your translation output.
func (c *Client) CreateParallelData(ctx context.Context, params *CreateParallelDataInput, optFns ...func(*Options)) (*CreateParallelDataOutput, error) {
	if params == nil {
		params = &CreateParallelDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateParallelData", params, optFns, c.addOperationCreateParallelDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateParallelDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateParallelDataInput struct {

	// A unique identifier for the request. This token is automatically generated when
	// you use Amazon Translate through an AWS SDK.
	//
	// This member is required.
	ClientToken *string

	// A custom name for the parallel data resource in Amazon Translate. You must
	// assign a name that is unique in the account and region.
	//
	// This member is required.
	Name *string

	// Specifies the format and S3 location of the parallel data input file.
	//
	// This member is required.
	ParallelDataConfig *types.ParallelDataConfig

	// A custom description for the parallel data resource in Amazon Translate.
	Description *string

	// The encryption key used to encrypt this object.
	EncryptionKey *types.EncryptionKey

	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateParallelDataOutput struct {

	// The custom name that you assigned to the parallel data resource.
	Name *string

	// The status of the parallel data resource. When the resource is ready for you to
	// use, the status is ACTIVE.
	Status types.ParallelDataStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateParallelDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateParallelData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateParallelData{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateParallelDataMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateParallelDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateParallelData(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateParallelData struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateParallelData) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateParallelData) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateParallelDataInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateParallelDataInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateParallelDataMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateParallelData{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateParallelData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "translate",
		OperationName: "CreateParallelData",
	}
}
