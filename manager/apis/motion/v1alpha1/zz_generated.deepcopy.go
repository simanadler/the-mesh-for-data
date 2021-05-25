// +build !ignore_autogenerated

// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	appv1alpha1 "github.com/mesh-for-data/mesh-for-data/manager/apis/app/v1alpha1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchTransfer) DeepCopyInto(out *BatchTransfer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchTransfer.
func (in *BatchTransfer) DeepCopy() *BatchTransfer {
	if in == nil {
		return nil
	}
	out := new(BatchTransfer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BatchTransfer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchTransferList) DeepCopyInto(out *BatchTransferList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BatchTransfer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchTransferList.
func (in *BatchTransferList) DeepCopy() *BatchTransferList {
	if in == nil {
		return nil
	}
	out := new(BatchTransferList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BatchTransferList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchTransferSpec) DeepCopyInto(out *BatchTransferSpec) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	in.Destination.DeepCopyInto(&out.Destination)
	if in.Transformation != nil {
		in, out := &in.Transformation, &out.Transformation
		*out = make([]Transformation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Spark != nil {
		in, out := &in.Spark, &out.Spark
		*out = new(Spark)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchTransferSpec.
func (in *BatchTransferSpec) DeepCopy() *BatchTransferSpec {
	if in == nil {
		return nil
	}
	out := new(BatchTransferSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchTransferStatus) DeepCopyInto(out *BatchTransferStatus) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.LastCompleted != nil {
		in, out := &in.LastCompleted, &out.LastCompleted
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.LastFailed != nil {
		in, out := &in.LastFailed, &out.LastFailed
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.LastSuccessTime != nil {
		in, out := &in.LastSuccessTime, &out.LastSuccessTime
		*out = (*in).DeepCopy()
	}
	if in.LastRecordTime != nil {
		in, out := &in.LastRecordTime, &out.LastRecordTime
		*out = (*in).DeepCopy()
	}
	if in.LastScheduleTime != nil {
		in, out := &in.LastScheduleTime, &out.LastScheduleTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchTransferStatus.
func (in *BatchTransferStatus) DeepCopy() *BatchTransferStatus {
	if in == nil {
		return nil
	}
	out := new(BatchTransferStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cloudant) DeepCopyInto(out *Cloudant) {
	*out = *in
	if in.SecretImport != nil {
		in, out := &in.SecretImport, &out.SecretImport
		*out = new(string)
		**out = **in
	}
	if in.Vault != nil {
		in, out := &in.Vault, &out.Vault
		*out = new(appv1alpha1.Vault)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cloudant.
func (in *Cloudant) DeepCopy() *Cloudant {
	if in == nil {
		return nil
	}
	out := new(Cloudant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataStore) DeepCopyInto(out *DataStore) {
	*out = *in
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(Database)
		(*in).DeepCopyInto(*out)
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(S3)
		(*in).DeepCopyInto(*out)
	}
	if in.Kafka != nil {
		in, out := &in.Kafka, &out.Kafka
		*out = new(Kafka)
		(*in).DeepCopyInto(*out)
	}
	if in.Cloudant != nil {
		in, out := &in.Cloudant, &out.Cloudant
		*out = new(Cloudant)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataStore.
func (in *DataStore) DeepCopy() *DataStore {
	if in == nil {
		return nil
	}
	out := new(DataStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Database) DeepCopyInto(out *Database) {
	*out = *in
	if in.SecretImport != nil {
		in, out := &in.SecretImport, &out.SecretImport
		*out = new(string)
		**out = **in
	}
	if in.Vault != nil {
		in, out := &in.Vault, &out.Vault
		*out = new(appv1alpha1.Vault)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Database.
func (in *Database) DeepCopy() *Database {
	if in == nil {
		return nil
	}
	out := new(Database)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kafka) DeepCopyInto(out *Kafka) {
	*out = *in
	if in.SecretImport != nil {
		in, out := &in.SecretImport, &out.SecretImport
		*out = new(string)
		**out = **in
	}
	if in.Vault != nil {
		in, out := &in.Vault, &out.Vault
		*out = new(appv1alpha1.Vault)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kafka.
func (in *Kafka) DeepCopy() *Kafka {
	if in == nil {
		return nil
	}
	out := new(Kafka)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3) DeepCopyInto(out *S3) {
	*out = *in
	if in.PartitionBy != nil {
		in, out := &in.PartitionBy, &out.PartitionBy
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.SecretImport != nil {
		in, out := &in.SecretImport, &out.SecretImport
		*out = new(string)
		**out = **in
	}
	if in.Vault != nil {
		in, out := &in.Vault, &out.Vault
		*out = new(appv1alpha1.Vault)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3.
func (in *S3) DeepCopy() *S3 {
	if in == nil {
		return nil
	}
	out := new(S3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Spark) DeepCopyInto(out *Spark) {
	*out = *in
	if in.AdditionalOptions != nil {
		in, out := &in.AdditionalOptions, &out.AdditionalOptions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Spark.
func (in *Spark) DeepCopy() *Spark {
	if in == nil {
		return nil
	}
	out := new(Spark)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamTransfer) DeepCopyInto(out *StreamTransfer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamTransfer.
func (in *StreamTransfer) DeepCopy() *StreamTransfer {
	if in == nil {
		return nil
	}
	out := new(StreamTransfer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StreamTransfer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamTransferList) DeepCopyInto(out *StreamTransferList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StreamTransfer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamTransferList.
func (in *StreamTransferList) DeepCopy() *StreamTransferList {
	if in == nil {
		return nil
	}
	out := new(StreamTransferList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StreamTransferList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamTransferSpec) DeepCopyInto(out *StreamTransferSpec) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	in.Destination.DeepCopyInto(&out.Destination)
	if in.Transformation != nil {
		in, out := &in.Transformation, &out.Transformation
		*out = make([]Transformation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamTransferSpec.
func (in *StreamTransferSpec) DeepCopy() *StreamTransferSpec {
	if in == nil {
		return nil
	}
	out := new(StreamTransferSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamTransferStatus) DeepCopyInto(out *StreamTransferStatus) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamTransferStatus.
func (in *StreamTransferStatus) DeepCopy() *StreamTransferStatus {
	if in == nil {
		return nil
	}
	out := new(StreamTransferStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Transformation) DeepCopyInto(out *Transformation) {
	*out = *in
	if in.Columns != nil {
		in, out := &in.Columns, &out.Columns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Transformation.
func (in *Transformation) DeepCopy() *Transformation {
	if in == nil {
		return nil
	}
	out := new(Transformation)
	in.DeepCopyInto(out)
	return out
}
