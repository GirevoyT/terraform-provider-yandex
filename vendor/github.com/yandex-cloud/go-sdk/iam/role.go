// Code generated by sdkgen. DO NOT EDIT.

// nolint
package iam

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
)

// RoleServiceClient is a iam.RoleServiceClient with
// lazy GRPC connection initialization.
type RoleServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

var _ iam.RoleServiceClient = &RoleServiceClient{}

// Get implements iam.RoleServiceClient
func (c *RoleServiceClient) Get(ctx context.Context, in *iam.GetRoleRequest, opts ...grpc.CallOption) (*iam.Role, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewRoleServiceClient(conn).Get(ctx, in, opts...)
}

// List implements iam.RoleServiceClient
func (c *RoleServiceClient) List(ctx context.Context, in *iam.ListRolesRequest, opts ...grpc.CallOption) (*iam.ListRolesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewRoleServiceClient(conn).List(ctx, in, opts...)
}
