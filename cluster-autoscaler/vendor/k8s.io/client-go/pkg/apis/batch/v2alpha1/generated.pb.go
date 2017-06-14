/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/pkg/apis/batch/v2alpha1/generated.proto
// DO NOT EDIT!

/*
	Package v2alpha1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/pkg/apis/batch/v2alpha1/generated.proto

	It has these top-level messages:
		CronJob
		CronJobList
		CronJobSpec
		CronJobStatus
		JobTemplate
		JobTemplateSpec
*/
package v2alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import k8s_io_apimachinery_pkg_apis_meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

import k8s_io_kubernetes_pkg_api_v1 "k8s.io/client-go/pkg/api/v1"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *CronJob) Reset()                    { *m = CronJob{} }
func (*CronJob) ProtoMessage()               {}
func (*CronJob) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *CronJobList) Reset()                    { *m = CronJobList{} }
func (*CronJobList) ProtoMessage()               {}
func (*CronJobList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *CronJobSpec) Reset()                    { *m = CronJobSpec{} }
func (*CronJobSpec) ProtoMessage()               {}
func (*CronJobSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *CronJobStatus) Reset()                    { *m = CronJobStatus{} }
func (*CronJobStatus) ProtoMessage()               {}
func (*CronJobStatus) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func (m *JobTemplate) Reset()                    { *m = JobTemplate{} }
func (*JobTemplate) ProtoMessage()               {}
func (*JobTemplate) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{4} }

func (m *JobTemplateSpec) Reset()                    { *m = JobTemplateSpec{} }
func (*JobTemplateSpec) ProtoMessage()               {}
func (*JobTemplateSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{5} }

func init() {
	proto.RegisterType((*CronJob)(nil), "k8s.io.client-go.pkg.apis.batch.v2alpha1.CronJob")
	proto.RegisterType((*CronJobList)(nil), "k8s.io.client-go.pkg.apis.batch.v2alpha1.CronJobList")
	proto.RegisterType((*CronJobSpec)(nil), "k8s.io.client-go.pkg.apis.batch.v2alpha1.CronJobSpec")
	proto.RegisterType((*CronJobStatus)(nil), "k8s.io.client-go.pkg.apis.batch.v2alpha1.CronJobStatus")
	proto.RegisterType((*JobTemplate)(nil), "k8s.io.client-go.pkg.apis.batch.v2alpha1.JobTemplate")
	proto.RegisterType((*JobTemplateSpec)(nil), "k8s.io.client-go.pkg.apis.batch.v2alpha1.JobTemplateSpec")
}
func (m *CronJob) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CronJob) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
	n2, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Status.Size()))
	n3, err := m.Status.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *CronJobList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CronJobList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ListMeta.Size()))
	n4, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *CronJobSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CronJobSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Schedule)))
	i += copy(dAtA[i:], m.Schedule)
	if m.StartingDeadlineSeconds != nil {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(*m.StartingDeadlineSeconds))
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ConcurrencyPolicy)))
	i += copy(dAtA[i:], m.ConcurrencyPolicy)
	if m.Suspend != nil {
		dAtA[i] = 0x20
		i++
		if *m.Suspend {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	dAtA[i] = 0x2a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.JobTemplate.Size()))
	n5, err := m.JobTemplate.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	if m.SuccessfulJobsHistoryLimit != nil {
		dAtA[i] = 0x30
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(*m.SuccessfulJobsHistoryLimit))
	}
	if m.FailedJobsHistoryLimit != nil {
		dAtA[i] = 0x38
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(*m.FailedJobsHistoryLimit))
	}
	return i, nil
}

func (m *CronJobStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CronJobStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Active) > 0 {
		for _, msg := range m.Active {
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.LastScheduleTime != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.LastScheduleTime.Size()))
		n6, err := m.LastScheduleTime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}

func (m *JobTemplate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JobTemplate) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n7, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Template.Size()))
	n8, err := m.Template.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n8
	return i, nil
}

func (m *JobTemplateSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JobTemplateSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n9, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n9
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
	n10, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n10
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CronJob) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *CronJobList) Size() (n int) {
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *CronJobSpec) Size() (n int) {
	var l int
	_ = l
	l = len(m.Schedule)
	n += 1 + l + sovGenerated(uint64(l))
	if m.StartingDeadlineSeconds != nil {
		n += 1 + sovGenerated(uint64(*m.StartingDeadlineSeconds))
	}
	l = len(m.ConcurrencyPolicy)
	n += 1 + l + sovGenerated(uint64(l))
	if m.Suspend != nil {
		n += 2
	}
	l = m.JobTemplate.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.SuccessfulJobsHistoryLimit != nil {
		n += 1 + sovGenerated(uint64(*m.SuccessfulJobsHistoryLimit))
	}
	if m.FailedJobsHistoryLimit != nil {
		n += 1 + sovGenerated(uint64(*m.FailedJobsHistoryLimit))
	}
	return n
}

func (m *CronJobStatus) Size() (n int) {
	var l int
	_ = l
	if len(m.Active) > 0 {
		for _, e := range m.Active {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.LastScheduleTime != nil {
		l = m.LastScheduleTime.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *JobTemplate) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Template.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *JobTemplateSpec) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CronJob) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CronJob{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "CronJobSpec", "CronJobSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "CronJobStatus", "CronJobStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CronJobList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CronJobList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "CronJob", "CronJob", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CronJobSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CronJobSpec{`,
		`Schedule:` + fmt.Sprintf("%v", this.Schedule) + `,`,
		`StartingDeadlineSeconds:` + valueToStringGenerated(this.StartingDeadlineSeconds) + `,`,
		`ConcurrencyPolicy:` + fmt.Sprintf("%v", this.ConcurrencyPolicy) + `,`,
		`Suspend:` + valueToStringGenerated(this.Suspend) + `,`,
		`JobTemplate:` + strings.Replace(strings.Replace(this.JobTemplate.String(), "JobTemplateSpec", "JobTemplateSpec", 1), `&`, ``, 1) + `,`,
		`SuccessfulJobsHistoryLimit:` + valueToStringGenerated(this.SuccessfulJobsHistoryLimit) + `,`,
		`FailedJobsHistoryLimit:` + valueToStringGenerated(this.FailedJobsHistoryLimit) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CronJobStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CronJobStatus{`,
		`Active:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Active), "ObjectReference", "k8s_io_kubernetes_pkg_api_v1.ObjectReference", 1), `&`, ``, 1) + `,`,
		`LastScheduleTime:` + strings.Replace(fmt.Sprintf("%v", this.LastScheduleTime), "Time", "k8s_io_apimachinery_pkg_apis_meta_v1.Time", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *JobTemplate) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&JobTemplate{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Template:` + strings.Replace(strings.Replace(this.Template.String(), "JobTemplateSpec", "JobTemplateSpec", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *JobTemplateSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&JobTemplateSpec{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "JobSpec", "k8s_io_kubernetes_pkg_apis_batch_v1.JobSpec", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CronJob) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CronJob: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CronJob: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CronJobList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CronJobList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CronJobList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, CronJob{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CronJobSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CronJobSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CronJobSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schedule", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Schedule = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartingDeadlineSeconds", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.StartingDeadlineSeconds = &v
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConcurrencyPolicy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConcurrencyPolicy = ConcurrencyPolicy(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Suspend", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Suspend = &b
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobTemplate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.JobTemplate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuccessfulJobsHistoryLimit", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SuccessfulJobsHistoryLimit = &v
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailedJobsHistoryLimit", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FailedJobsHistoryLimit = &v
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CronJobStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CronJobStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CronJobStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Active = append(m.Active, k8s_io_kubernetes_pkg_api_v1.ObjectReference{})
			if err := m.Active[len(m.Active)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastScheduleTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastScheduleTime == nil {
				m.LastScheduleTime = &k8s_io_apimachinery_pkg_apis_meta_v1.Time{}
			}
			if err := m.LastScheduleTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *JobTemplate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: JobTemplate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JobTemplate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Template", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Template.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *JobTemplateSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: JobTemplateSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JobTemplateSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/client-go/pkg/apis/batch/v2alpha1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 783 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x4d, 0x4f, 0xdb, 0x48,
	0x18, 0xc7, 0xe3, 0x90, 0x37, 0x26, 0xcb, 0x2e, 0x78, 0x57, 0x10, 0x65, 0x25, 0x27, 0x8a, 0xb4,
	0x52, 0x16, 0x81, 0xbd, 0x84, 0x15, 0x62, 0xf7, 0x56, 0x53, 0x55, 0x2d, 0xa2, 0x2f, 0x72, 0x40,
	0xad, 0x2a, 0x54, 0x31, 0x9e, 0x4c, 0x92, 0x21, 0x7e, 0xab, 0x67, 0x1c, 0x29, 0xb7, 0x7e, 0x84,
	0x7e, 0x8b, 0x7e, 0x8b, 0x5e, 0xda, 0x03, 0x47, 0x0e, 0x3d, 0xb4, 0x97, 0xa8, 0xb8, 0xdf, 0xa2,
	0xa7, 0xca, 0x13, 0x27, 0x0e, 0x38, 0x2e, 0xa1, 0x95, 0xb8, 0x79, 0xc6, 0xcf, 0xff, 0x37, 0xcf,
	0xf3, 0x7f, 0x9e, 0x19, 0xf0, 0x5f, 0x6f, 0x97, 0xca, 0xc4, 0x56, 0x7a, 0x9e, 0x8e, 0x5d, 0x0b,
	0x33, 0x4c, 0x15, 0xa7, 0xd7, 0x51, 0xa0, 0x43, 0xa8, 0xa2, 0x43, 0x86, 0xba, 0x4a, 0xbf, 0x01,
	0x0d, 0xa7, 0x0b, 0xb7, 0x94, 0x0e, 0xb6, 0xb0, 0x0b, 0x19, 0x6e, 0xc9, 0x8e, 0x6b, 0x33, 0x5b,
	0xfc, 0x7b, 0x24, 0x95, 0x23, 0xa9, 0xec, 0xf4, 0x3a, 0x72, 0x20, 0x95, 0xb9, 0x54, 0x1e, 0x4b,
	0xcb, 0x9b, 0x1d, 0xc2, 0xba, 0x9e, 0x2e, 0x23, 0xdb, 0x54, 0x3a, 0x76, 0xc7, 0x56, 0x38, 0x41,
	0xf7, 0xda, 0x7c, 0xc5, 0x17, 0xfc, 0x6b, 0x44, 0x2e, 0xff, 0x1b, 0x26, 0x05, 0x1d, 0x62, 0x42,
	0xd4, 0x25, 0x16, 0x76, 0x07, 0x51, 0x5a, 0x26, 0x66, 0x50, 0xe9, 0xc7, 0xf2, 0x29, 0x2b, 0x49,
	0x2a, 0xd7, 0xb3, 0x18, 0x31, 0x71, 0x4c, 0xb0, 0x73, 0x9d, 0x80, 0xa2, 0x2e, 0x36, 0x61, 0x4c,
	0xb7, 0x9d, 0xa4, 0xf3, 0x18, 0x31, 0x14, 0x62, 0x31, 0xca, 0xdc, 0x98, 0x68, 0x23, 0xd1, 0xe8,
	0x59, 0xb5, 0x6c, 0x5f, 0xdf, 0x96, 0x98, 0xa8, 0xf6, 0x26, 0x0d, 0xf2, 0x7b, 0xae, 0x6d, 0xed,
	0xdb, 0xba, 0x78, 0x02, 0x0a, 0x81, 0x4f, 0x2d, 0xc8, 0x60, 0x49, 0xa8, 0x0a, 0xf5, 0x62, 0xe3,
	0x1f, 0x39, 0xec, 0xd7, 0x74, 0xda, 0x51, 0xc7, 0x82, 0x68, 0xb9, 0xbf, 0x25, 0x3f, 0xd6, 0x4f,
	0x31, 0x62, 0x0f, 0x31, 0x83, 0xaa, 0x78, 0x36, 0xac, 0xa4, 0xfc, 0x61, 0x05, 0x44, 0x7b, 0xda,
	0x84, 0x2a, 0x3e, 0x03, 0x19, 0xea, 0x60, 0x54, 0x4a, 0x73, 0xfa, 0x8e, 0x3c, 0xf7, 0x34, 0xc8,
	0x61, 0x8e, 0x4d, 0x07, 0x23, 0xf5, 0x97, 0xf0, 0x8c, 0x4c, 0xb0, 0xd2, 0x38, 0x51, 0x3c, 0x01,
	0x39, 0xca, 0x20, 0xf3, 0x68, 0x69, 0x81, 0xb3, 0x77, 0x7f, 0x80, 0xcd, 0xf5, 0xea, 0xaf, 0x21,
	0x3d, 0x37, 0x5a, 0x6b, 0x21, 0xb7, 0xf6, 0x5e, 0x00, 0xc5, 0x30, 0xf2, 0x80, 0x50, 0x26, 0x1e,
	0xc7, 0xdc, 0x92, 0xe7, 0x73, 0x2b, 0x50, 0x73, 0xaf, 0x96, 0xc3, 0x93, 0x0a, 0xe3, 0x9d, 0x29,
	0xa7, 0x9e, 0x82, 0x2c, 0x61, 0xd8, 0xa4, 0xa5, 0x74, 0x75, 0xa1, 0x5e, 0x6c, 0x34, 0x6e, 0x5e,
	0x8e, 0xba, 0x14, 0xe2, 0xb3, 0x0f, 0x02, 0x90, 0x36, 0xe2, 0xd5, 0xde, 0x66, 0x26, 0x65, 0x04,
	0xf6, 0x89, 0x1b, 0xa0, 0x10, 0x8c, 0x6c, 0xcb, 0x33, 0x30, 0x2f, 0x63, 0x31, 0x4a, 0xab, 0x19,
	0xee, 0x6b, 0x93, 0x08, 0xf1, 0x08, 0xac, 0x51, 0x06, 0x5d, 0x46, 0xac, 0xce, 0x5d, 0x0c, 0x5b,
	0x06, 0xb1, 0x70, 0x13, 0x23, 0xdb, 0x6a, 0x51, 0xde, 0xd3, 0x05, 0xf5, 0x4f, 0x7f, 0x58, 0x59,
	0x6b, 0xce, 0x0e, 0xd1, 0x92, 0xb4, 0xe2, 0x31, 0x58, 0x41, 0xb6, 0x85, 0x3c, 0xd7, 0xc5, 0x16,
	0x1a, 0x3c, 0xb1, 0x0d, 0x82, 0x06, 0xbc, 0x91, 0x8b, 0xaa, 0x1c, 0x66, 0xb3, 0xb2, 0x77, 0x35,
	0xe0, 0xeb, 0xac, 0x4d, 0x2d, 0x0e, 0x12, 0xff, 0x02, 0x79, 0xea, 0x51, 0x07, 0x5b, 0xad, 0x52,
	0xa6, 0x2a, 0xd4, 0x0b, 0x6a, 0xd1, 0x1f, 0x56, 0xf2, 0xcd, 0xd1, 0x96, 0x36, 0xfe, 0x27, 0xbe,
	0x04, 0xc5, 0x53, 0x5b, 0x3f, 0xc4, 0xa6, 0x63, 0x40, 0x86, 0x4b, 0x59, 0xde, 0xd3, 0xff, 0x6f,
	0x60, 0xfc, 0x7e, 0xa4, 0xe6, 0x73, 0xfa, 0x7b, 0x98, 0x7a, 0x71, 0xea, 0x87, 0x36, 0x7d, 0x86,
	0xf8, 0x02, 0x94, 0xa9, 0x87, 0x10, 0xa6, 0xb4, 0xed, 0x19, 0xfb, 0xb6, 0x4e, 0xef, 0x13, 0xca,
	0x6c, 0x77, 0x70, 0x40, 0x4c, 0xc2, 0x4a, 0xb9, 0xaa, 0x50, 0xcf, 0xaa, 0x92, 0x3f, 0xac, 0x94,
	0x9b, 0x89, 0x51, 0xda, 0x77, 0x08, 0xa2, 0x06, 0x56, 0xdb, 0x90, 0x18, 0xb8, 0x15, 0x63, 0xe7,
	0x39, 0xbb, 0xec, 0x0f, 0x2b, 0xab, 0xf7, 0x66, 0x46, 0x68, 0x09, 0xca, 0xda, 0x07, 0x01, 0x2c,
	0x5d, 0xba, 0x31, 0xe2, 0x11, 0xc8, 0x41, 0xc4, 0x48, 0x3f, 0x18, 0xa0, 0x60, 0x58, 0x37, 0x93,
	0x3d, 0x8b, 0x5e, 0x0b, 0x0d, 0xb7, 0x71, 0xd0, 0x24, 0x1c, 0x5d, 0xb8, 0x3b, 0x1c, 0xa2, 0x85,
	0x30, 0xd1, 0x00, 0xcb, 0x06, 0xa4, 0x6c, 0x3c, 0x85, 0x87, 0xc4, 0xc4, 0xbc, 0x7f, 0xc5, 0xc6,
	0xfa, 0x7c, 0x17, 0x2d, 0x50, 0xa8, 0x7f, 0xf8, 0xc3, 0xca, 0xf2, 0xc1, 0x15, 0x8e, 0x16, 0x23,
	0xd7, 0x3e, 0x09, 0x60, 0xba, 0x4f, 0xb7, 0xf0, 0x18, 0x76, 0x41, 0x81, 0x8d, 0x87, 0x2d, 0xfd,
	0xd3, 0xc3, 0x36, 0xb9, 0xb5, 0x93, 0x49, 0x9b, 0xd0, 0x6b, 0xef, 0x04, 0xf0, 0xdb, 0x95, 0xf8,
	0x5b, 0xa8, 0xef, 0xd1, 0xa5, 0xc7, 0x7e, 0x63, 0x8e, 0xda, 0x78, 0x55, 0x49, 0x4f, 0xbc, 0xba,
	0x7e, 0x76, 0x21, 0xa5, 0xce, 0x2f, 0xa4, 0xd4, 0xc7, 0x0b, 0x29, 0xf5, 0xca, 0x97, 0x84, 0x33,
	0x5f, 0x12, 0xce, 0x7d, 0x49, 0xf8, 0xec, 0x4b, 0xc2, 0xeb, 0x2f, 0x52, 0xea, 0x79, 0x61, 0xec,
	0xce, 0xb7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0x9b, 0x4c, 0x2b, 0xa3, 0x08, 0x00, 0x00,
}