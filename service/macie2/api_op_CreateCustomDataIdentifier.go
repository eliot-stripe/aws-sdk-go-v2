// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates and defines the criteria and other settings for a custom data
// identifier.
func (c *Client) CreateCustomDataIdentifier(ctx context.Context, params *CreateCustomDataIdentifierInput, optFns ...func(*Options)) (*CreateCustomDataIdentifierOutput, error) {
	if params == nil {
		params = &CreateCustomDataIdentifierInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCustomDataIdentifier", params, optFns, c.addOperationCreateCustomDataIdentifierMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCustomDataIdentifierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCustomDataIdentifierInput struct {

	// A unique, case-sensitive token that you provide to ensure the idempotency of the
	// request.
	ClientToken *string

	// A custom description of the custom data identifier. The description can contain
	// as many as 512 characters. We strongly recommend that you avoid including any
	// sensitive data in the description of a custom data identifier. Other users of
	// your account might be able to see the identifier's description, depending on the
	// actions that they're allowed to perform in Amazon Macie.
	Description *string

	// An array that lists specific character sequences (ignore words) to exclude from
	// the results. If the text matched by the regular expression is the same as any
	// string in this array, Amazon Macie ignores it. The array can contain as many as
	// 10 ignore words. Each ignore word can contain 4-90 UTF-8 characters. Ignore
	// words are case sensitive.
	IgnoreWords []string

	// An array that lists specific character sequences (keywords), one of which must
	// be within proximity (maximumMatchDistance) of the regular expression to match.
	// The array can contain as many as 50 keywords. Each keyword can contain 3-90
	// UTF-8 characters. Keywords aren't case sensitive.
	Keywords []string

	// The maximum number of characters that can exist between text that matches the
	// regex pattern and the character sequences specified by the keywords array.
	// Amazon Macie includes or excludes a result based on the proximity of a keyword
	// to text that matches the regex pattern. The distance can be 1-300 characters.
	// The default value is 50.
	MaximumMatchDistance int32

	// A custom name for the custom data identifier. The name can contain as many as
	// 128 characters. We strongly recommend that you avoid including any sensitive
	// data in the name of a custom data identifier. Other users of your account might
	// be able to see the identifier's name, depending on the actions that they're
	// allowed to perform in Amazon Macie.
	Name *string

	// The regular expression (regex) that defines the pattern to match. The expression
	// can contain as many as 512 characters.
	Regex *string

	// A map of key-value pairs that specifies the tags to associate with the custom
	// data identifier. A custom data identifier can have a maximum of 50 tags. Each
	// tag consists of a tag key and an associated tag value. The maximum length of a
	// tag key is 128 characters. The maximum length of a tag value is 256 characters.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateCustomDataIdentifierOutput struct {

	// The unique identifier for the custom data identifier that was created.
	CustomDataIdentifierId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCustomDataIdentifierMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCustomDataIdentifier{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCustomDataIdentifier{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateCustomDataIdentifierMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCustomDataIdentifier(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateCustomDataIdentifier struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateCustomDataIdentifier) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateCustomDataIdentifier) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateCustomDataIdentifierInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateCustomDataIdentifierInput ")
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
func addIdempotencyToken_opCreateCustomDataIdentifierMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateCustomDataIdentifier{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateCustomDataIdentifier(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "CreateCustomDataIdentifier",
	}
}
