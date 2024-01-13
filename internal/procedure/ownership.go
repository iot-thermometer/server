package procedure

import (
	"context"
	"fmt"
	"github.com/iot-thermometer/server/internal/model"

	"github.com/iot-thermometer/server/gen"
	"github.com/iot-thermometer/server/internal/service"
	"google.golang.org/grpc/metadata"
)

type Ownership interface {
	AddMember(ctx context.Context, request *gen.AddMemberRequest) (*gen.AddMemberResponse, error)
	ListMembers(ctx context.Context, request *gen.ListMembersRequest) (*gen.ListMembersResponse, error)
	RemoveMember(ctx context.Context, request *gen.RemoveMemberRequest) (*gen.RemoveMemberResponse, error)
}

type ownership struct {
	ownershipService service.Ownership
	userService      service.User
}

func newOwnershipProcedure(ownershipService service.Ownership, userService service.User) Ownership {
	return ownership{
		ownershipService: ownershipService,
		userService:      userService,
	}
}

func (o ownership) AddMember(ctx context.Context, request *gen.AddMemberRequest) (*gen.AddMemberResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing context metadata")
	}
	token := md.Get("Authorization")[0]
	userID, err := o.userService.DecodeToken(token)
	if err != nil {
		return nil, err
	}
	role := request.Role.Enum()
	err = o.ownershipService.AddMember(userID, uint(request.DeviceID), request.Email, model.OwnershipRole(*role))
	if err != nil {
		return nil, err
	}
	return &gen.AddMemberResponse{}, nil
}

func (o ownership) ListMembers(ctx context.Context, request *gen.ListMembersRequest) (*gen.ListMembersResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing context metadata")
	}
	token := md.Get("Authorization")[0]
	userID, err := o.userService.DecodeToken(token)
	if err != nil {
		return nil, err
	}
	ownerships, err := o.ownershipService.ListMembers(userID, uint(request.DeviceID))
	if err != nil {
		return nil, err
	}

	var pbOwnerships []*gen.Ownership
	for _, ownership := range ownerships {
		pbOwnership := &gen.Ownership{
			DeviceID: int64(ownership.DeviceID),
			UserID:   int64(ownership.UserID),
			Email:    ownership.User.Email,
			Role:     gen.OwnershipRole(ownership.Role),
			IsMe:     ownership.User.ID == userID,
		}
		pbOwnerships = append(pbOwnerships, pbOwnership)
	}

	return &gen.ListMembersResponse{
		Ownerships: pbOwnerships,
	}, nil
}

func (o ownership) RemoveMember(ctx context.Context, request *gen.RemoveMemberRequest) (*gen.RemoveMemberResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing context metadata")
	}
	token := md.Get("Authorization")[0]
	userID, err := o.userService.DecodeToken(token)
	if err != nil {
		return nil, err
	}
	err = o.ownershipService.RemoveMember(userID, uint(request.DeviceID), uint(request.UserID))
	if err != nil {
		return nil, err
	}
	return &gen.RemoveMemberResponse{}, nil
}
