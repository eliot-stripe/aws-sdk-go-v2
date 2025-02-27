// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the invitations that you have received for resource shares.
func (c *Client) GetResourceShareInvitations(ctx context.Context, params *GetResourceShareInvitationsInput, optFns ...func(*Options)) (*GetResourceShareInvitationsOutput, error) {
	if params == nil {
		params = &GetResourceShareInvitationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResourceShareInvitations", params, optFns, c.addOperationGetResourceShareInvitationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourceShareInvitationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourceShareInvitationsInput struct {

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The Amazon Resource Names (ARN) of the resource shares.
	ResourceShareArns []string

	// The Amazon Resource Names (ARN) of the invitations.
	ResourceShareInvitationArns []string

	noSmithyDocumentSerde
}

type GetResourceShareInvitationsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the invitations.
	ResourceShareInvitations []types.ResourceShareInvitation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetResourceShareInvitationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetResourceShareInvitations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetResourceShareInvitations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResourceShareInvitations(options.Region), middleware.Before); err != nil {
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

// GetResourceShareInvitationsAPIClient is a client that implements the
// GetResourceShareInvitations operation.
type GetResourceShareInvitationsAPIClient interface {
	GetResourceShareInvitations(context.Context, *GetResourceShareInvitationsInput, ...func(*Options)) (*GetResourceShareInvitationsOutput, error)
}

var _ GetResourceShareInvitationsAPIClient = (*Client)(nil)

// GetResourceShareInvitationsPaginatorOptions is the paginator options for
// GetResourceShareInvitations
type GetResourceShareInvitationsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetResourceShareInvitationsPaginator is a paginator for
// GetResourceShareInvitations
type GetResourceShareInvitationsPaginator struct {
	options   GetResourceShareInvitationsPaginatorOptions
	client    GetResourceShareInvitationsAPIClient
	params    *GetResourceShareInvitationsInput
	nextToken *string
	firstPage bool
}

// NewGetResourceShareInvitationsPaginator returns a new
// GetResourceShareInvitationsPaginator
func NewGetResourceShareInvitationsPaginator(client GetResourceShareInvitationsAPIClient, params *GetResourceShareInvitationsInput, optFns ...func(*GetResourceShareInvitationsPaginatorOptions)) *GetResourceShareInvitationsPaginator {
	if params == nil {
		params = &GetResourceShareInvitationsInput{}
	}

	options := GetResourceShareInvitationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetResourceShareInvitationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetResourceShareInvitationsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetResourceShareInvitations page.
func (p *GetResourceShareInvitationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetResourceShareInvitationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.GetResourceShareInvitations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetResourceShareInvitations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "GetResourceShareInvitations",
	}
}
