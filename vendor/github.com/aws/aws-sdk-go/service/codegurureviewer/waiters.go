// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codegurureviewer

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// WaitUntilCodeReviewCompleted uses the CodeGuruReviewer API operation
// DescribeCodeReview to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *CodeGuruReviewer) WaitUntilCodeReviewCompleted(input *DescribeCodeReviewInput) error {
	return c.WaitUntilCodeReviewCompletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilCodeReviewCompletedWithContext is an extended version of WaitUntilCodeReviewCompleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CodeGuruReviewer) WaitUntilCodeReviewCompletedWithContext(ctx aws.Context, input *DescribeCodeReviewInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilCodeReviewCompleted",
		MaxAttempts: 180,
		Delay:       request.ConstantWaiterDelay(10 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "CodeReview.State",
				Expected: "Completed",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "CodeReview.State",
				Expected: "Failed",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "CodeReview.State",
				Expected: "Pending",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeCodeReviewInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeCodeReviewRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilRepositoryAssociationSucceeded uses the CodeGuruReviewer API operation
// DescribeRepositoryAssociation to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *CodeGuruReviewer) WaitUntilRepositoryAssociationSucceeded(input *DescribeRepositoryAssociationInput) error {
	return c.WaitUntilRepositoryAssociationSucceededWithContext(aws.BackgroundContext(), input)
}

// WaitUntilRepositoryAssociationSucceededWithContext is an extended version of WaitUntilRepositoryAssociationSucceeded.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CodeGuruReviewer) WaitUntilRepositoryAssociationSucceededWithContext(ctx aws.Context, input *DescribeRepositoryAssociationInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilRepositoryAssociationSucceeded",
		MaxAttempts: 30,
		Delay:       request.ConstantWaiterDelay(10 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "RepositoryAssociation.State",
				Expected: "Associated",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "RepositoryAssociation.State",
				Expected: "Failed",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "RepositoryAssociation.State",
				Expected: "Associating",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeRepositoryAssociationInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeRepositoryAssociationRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
