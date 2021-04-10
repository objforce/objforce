// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/data/data.proto

package data

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/emptypb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *SObject) Validate() error {
	return nil
}
func (this *CreateSObjectRequest) Validate() error {
	if this.Object != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Object); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Object", err)
		}
	}
	return nil
}
func (this *UpdateSObjectRequest) Validate() error {
	if this.Object != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Object); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Object", err)
		}
	}
	return nil
}
func (this *UpsertSObjectRequest) Validate() error {
	if this.Object != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Object); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Object", err)
		}
	}
	return nil
}
func (this *UpsertSObjectResponse) Validate() error {
	for _, item := range this.Errors {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Errors", err)
			}
		}
	}
	return nil
}
func (this *GetSObjectRequest) Validate() error {
	return nil
}
func (this *DeleteSObjectRequest) Validate() error {
	return nil
}
func (this *Error) Validate() error {
	return nil
}