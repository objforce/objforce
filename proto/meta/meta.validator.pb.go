// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/meta/meta.proto

package meta

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

func (this *CreateCustomObjectRequest) Validate() error {
	for _, item := range this.Fields {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Fields", err)
			}
		}
	}
	for _, item := range this.Indexes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Indexes", err)
			}
		}
	}
	return nil
}
func (this *FindByObjNameRequest) Validate() error {
	return nil
}
func (this *GetCustomObjectRequest) Validate() error {
	return nil
}
func (this *DeleteCustomObjectRequest) Validate() error {
	return nil
}
func (this *GetCustomFieldRequest) Validate() error {
	return nil
}
func (this *DeleteCustomFieldRequest) Validate() error {
	return nil
}
