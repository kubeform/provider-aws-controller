/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func GetEncoder() map[string]jsoniter.ValEncoder {
	return map[string]jsoniter.ValEncoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecAutoTerminationPolicy{}).Type1()):                   ClusterSpecAutoTerminationPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecCoreInstanceFleet{}).Type1()):                       ClusterSpecCoreInstanceFleetCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecCoreInstanceFleetLaunchSpecifications{}).Type1()):   ClusterSpecCoreInstanceFleetLaunchSpecificationsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecCoreInstanceGroup{}).Type1()):                       ClusterSpecCoreInstanceGroupCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecEc2Attributes{}).Type1()):                           ClusterSpecEc2AttributesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecKerberosAttributes{}).Type1()):                      ClusterSpecKerberosAttributesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecMasterInstanceFleet{}).Type1()):                     ClusterSpecMasterInstanceFleetCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecMasterInstanceFleetLaunchSpecifications{}).Type1()): ClusterSpecMasterInstanceFleetLaunchSpecificationsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecMasterInstanceGroup{}).Type1()):                     ClusterSpecMasterInstanceGroupCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecStepHadoopJarStep{}).Type1()):                       ClusterSpecStepHadoopJarStepCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceFleetSpecLaunchSpecifications{}).Type1()):              InstanceFleetSpecLaunchSpecificationsCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecAutoTerminationPolicy{}).Type1()):                   ClusterSpecAutoTerminationPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecCoreInstanceFleet{}).Type1()):                       ClusterSpecCoreInstanceFleetCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecCoreInstanceFleetLaunchSpecifications{}).Type1()):   ClusterSpecCoreInstanceFleetLaunchSpecificationsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecCoreInstanceGroup{}).Type1()):                       ClusterSpecCoreInstanceGroupCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecEc2Attributes{}).Type1()):                           ClusterSpecEc2AttributesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecKerberosAttributes{}).Type1()):                      ClusterSpecKerberosAttributesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecMasterInstanceFleet{}).Type1()):                     ClusterSpecMasterInstanceFleetCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecMasterInstanceFleetLaunchSpecifications{}).Type1()): ClusterSpecMasterInstanceFleetLaunchSpecificationsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecMasterInstanceGroup{}).Type1()):                     ClusterSpecMasterInstanceGroupCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecStepHadoopJarStep{}).Type1()):                       ClusterSpecStepHadoopJarStepCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceFleetSpecLaunchSpecifications{}).Type1()):              InstanceFleetSpecLaunchSpecificationsCodec{},
	}
}

func getEncodersWithout(typ string) map[string]jsoniter.ValEncoder {
	origMap := GetEncoder()
	delete(origMap, typ)
	return origMap
}

func getDecodersWithout(typ string) map[string]jsoniter.ValDecoder {
	origMap := GetDecoder()
	delete(origMap, typ)
	return origMap
}

// +k8s:deepcopy-gen=false
type ClusterSpecAutoTerminationPolicyCodec struct {
}

func (ClusterSpecAutoTerminationPolicyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecAutoTerminationPolicy)(ptr) == nil
}

func (ClusterSpecAutoTerminationPolicyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecAutoTerminationPolicy)(ptr)
	var objs []ClusterSpecAutoTerminationPolicy
	if obj != nil {
		objs = []ClusterSpecAutoTerminationPolicy{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecAutoTerminationPolicy{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecAutoTerminationPolicyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecAutoTerminationPolicy)(ptr) = ClusterSpecAutoTerminationPolicy{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecAutoTerminationPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecAutoTerminationPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecAutoTerminationPolicy)(ptr) = objs[0]
			} else {
				*(*ClusterSpecAutoTerminationPolicy)(ptr) = ClusterSpecAutoTerminationPolicy{}
			}
		} else {
			*(*ClusterSpecAutoTerminationPolicy)(ptr) = ClusterSpecAutoTerminationPolicy{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecAutoTerminationPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecAutoTerminationPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecAutoTerminationPolicy)(ptr) = obj
		} else {
			*(*ClusterSpecAutoTerminationPolicy)(ptr) = ClusterSpecAutoTerminationPolicy{}
		}
	default:
		iter.ReportError("decode ClusterSpecAutoTerminationPolicy", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClusterSpecCoreInstanceFleetCodec struct {
}

func (ClusterSpecCoreInstanceFleetCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecCoreInstanceFleet)(ptr) == nil
}

func (ClusterSpecCoreInstanceFleetCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecCoreInstanceFleet)(ptr)
	var objs []ClusterSpecCoreInstanceFleet
	if obj != nil {
		objs = []ClusterSpecCoreInstanceFleet{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecCoreInstanceFleet{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecCoreInstanceFleetCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecCoreInstanceFleet)(ptr) = ClusterSpecCoreInstanceFleet{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecCoreInstanceFleet

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecCoreInstanceFleet{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecCoreInstanceFleet)(ptr) = objs[0]
			} else {
				*(*ClusterSpecCoreInstanceFleet)(ptr) = ClusterSpecCoreInstanceFleet{}
			}
		} else {
			*(*ClusterSpecCoreInstanceFleet)(ptr) = ClusterSpecCoreInstanceFleet{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecCoreInstanceFleet

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecCoreInstanceFleet{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecCoreInstanceFleet)(ptr) = obj
		} else {
			*(*ClusterSpecCoreInstanceFleet)(ptr) = ClusterSpecCoreInstanceFleet{}
		}
	default:
		iter.ReportError("decode ClusterSpecCoreInstanceFleet", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClusterSpecCoreInstanceFleetLaunchSpecificationsCodec struct {
}

func (ClusterSpecCoreInstanceFleetLaunchSpecificationsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecCoreInstanceFleetLaunchSpecifications)(ptr) == nil
}

func (ClusterSpecCoreInstanceFleetLaunchSpecificationsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecCoreInstanceFleetLaunchSpecifications)(ptr)
	var objs []ClusterSpecCoreInstanceFleetLaunchSpecifications
	if obj != nil {
		objs = []ClusterSpecCoreInstanceFleetLaunchSpecifications{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecCoreInstanceFleetLaunchSpecifications{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecCoreInstanceFleetLaunchSpecificationsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecCoreInstanceFleetLaunchSpecifications)(ptr) = ClusterSpecCoreInstanceFleetLaunchSpecifications{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecCoreInstanceFleetLaunchSpecifications

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecCoreInstanceFleetLaunchSpecifications{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecCoreInstanceFleetLaunchSpecifications)(ptr) = objs[0]
			} else {
				*(*ClusterSpecCoreInstanceFleetLaunchSpecifications)(ptr) = ClusterSpecCoreInstanceFleetLaunchSpecifications{}
			}
		} else {
			*(*ClusterSpecCoreInstanceFleetLaunchSpecifications)(ptr) = ClusterSpecCoreInstanceFleetLaunchSpecifications{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecCoreInstanceFleetLaunchSpecifications

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecCoreInstanceFleetLaunchSpecifications{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecCoreInstanceFleetLaunchSpecifications)(ptr) = obj
		} else {
			*(*ClusterSpecCoreInstanceFleetLaunchSpecifications)(ptr) = ClusterSpecCoreInstanceFleetLaunchSpecifications{}
		}
	default:
		iter.ReportError("decode ClusterSpecCoreInstanceFleetLaunchSpecifications", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClusterSpecCoreInstanceGroupCodec struct {
}

func (ClusterSpecCoreInstanceGroupCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecCoreInstanceGroup)(ptr) == nil
}

func (ClusterSpecCoreInstanceGroupCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecCoreInstanceGroup)(ptr)
	var objs []ClusterSpecCoreInstanceGroup
	if obj != nil {
		objs = []ClusterSpecCoreInstanceGroup{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecCoreInstanceGroup{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecCoreInstanceGroupCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecCoreInstanceGroup)(ptr) = ClusterSpecCoreInstanceGroup{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecCoreInstanceGroup

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecCoreInstanceGroup{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecCoreInstanceGroup)(ptr) = objs[0]
			} else {
				*(*ClusterSpecCoreInstanceGroup)(ptr) = ClusterSpecCoreInstanceGroup{}
			}
		} else {
			*(*ClusterSpecCoreInstanceGroup)(ptr) = ClusterSpecCoreInstanceGroup{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecCoreInstanceGroup

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecCoreInstanceGroup{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecCoreInstanceGroup)(ptr) = obj
		} else {
			*(*ClusterSpecCoreInstanceGroup)(ptr) = ClusterSpecCoreInstanceGroup{}
		}
	default:
		iter.ReportError("decode ClusterSpecCoreInstanceGroup", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClusterSpecEc2AttributesCodec struct {
}

func (ClusterSpecEc2AttributesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecEc2Attributes)(ptr) == nil
}

func (ClusterSpecEc2AttributesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecEc2Attributes)(ptr)
	var objs []ClusterSpecEc2Attributes
	if obj != nil {
		objs = []ClusterSpecEc2Attributes{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecEc2Attributes{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecEc2AttributesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecEc2Attributes)(ptr) = ClusterSpecEc2Attributes{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecEc2Attributes

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecEc2Attributes{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecEc2Attributes)(ptr) = objs[0]
			} else {
				*(*ClusterSpecEc2Attributes)(ptr) = ClusterSpecEc2Attributes{}
			}
		} else {
			*(*ClusterSpecEc2Attributes)(ptr) = ClusterSpecEc2Attributes{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecEc2Attributes

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecEc2Attributes{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecEc2Attributes)(ptr) = obj
		} else {
			*(*ClusterSpecEc2Attributes)(ptr) = ClusterSpecEc2Attributes{}
		}
	default:
		iter.ReportError("decode ClusterSpecEc2Attributes", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClusterSpecKerberosAttributesCodec struct {
}

func (ClusterSpecKerberosAttributesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecKerberosAttributes)(ptr) == nil
}

func (ClusterSpecKerberosAttributesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecKerberosAttributes)(ptr)
	var objs []ClusterSpecKerberosAttributes
	if obj != nil {
		objs = []ClusterSpecKerberosAttributes{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecKerberosAttributes{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecKerberosAttributesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecKerberosAttributes)(ptr) = ClusterSpecKerberosAttributes{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecKerberosAttributes

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecKerberosAttributes{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecKerberosAttributes)(ptr) = objs[0]
			} else {
				*(*ClusterSpecKerberosAttributes)(ptr) = ClusterSpecKerberosAttributes{}
			}
		} else {
			*(*ClusterSpecKerberosAttributes)(ptr) = ClusterSpecKerberosAttributes{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecKerberosAttributes

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecKerberosAttributes{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecKerberosAttributes)(ptr) = obj
		} else {
			*(*ClusterSpecKerberosAttributes)(ptr) = ClusterSpecKerberosAttributes{}
		}
	default:
		iter.ReportError("decode ClusterSpecKerberosAttributes", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClusterSpecMasterInstanceFleetCodec struct {
}

func (ClusterSpecMasterInstanceFleetCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecMasterInstanceFleet)(ptr) == nil
}

func (ClusterSpecMasterInstanceFleetCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecMasterInstanceFleet)(ptr)
	var objs []ClusterSpecMasterInstanceFleet
	if obj != nil {
		objs = []ClusterSpecMasterInstanceFleet{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecMasterInstanceFleet{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecMasterInstanceFleetCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecMasterInstanceFleet)(ptr) = ClusterSpecMasterInstanceFleet{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecMasterInstanceFleet

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecMasterInstanceFleet{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecMasterInstanceFleet)(ptr) = objs[0]
			} else {
				*(*ClusterSpecMasterInstanceFleet)(ptr) = ClusterSpecMasterInstanceFleet{}
			}
		} else {
			*(*ClusterSpecMasterInstanceFleet)(ptr) = ClusterSpecMasterInstanceFleet{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecMasterInstanceFleet

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecMasterInstanceFleet{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecMasterInstanceFleet)(ptr) = obj
		} else {
			*(*ClusterSpecMasterInstanceFleet)(ptr) = ClusterSpecMasterInstanceFleet{}
		}
	default:
		iter.ReportError("decode ClusterSpecMasterInstanceFleet", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClusterSpecMasterInstanceFleetLaunchSpecificationsCodec struct {
}

func (ClusterSpecMasterInstanceFleetLaunchSpecificationsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecMasterInstanceFleetLaunchSpecifications)(ptr) == nil
}

func (ClusterSpecMasterInstanceFleetLaunchSpecificationsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecMasterInstanceFleetLaunchSpecifications)(ptr)
	var objs []ClusterSpecMasterInstanceFleetLaunchSpecifications
	if obj != nil {
		objs = []ClusterSpecMasterInstanceFleetLaunchSpecifications{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecMasterInstanceFleetLaunchSpecifications{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecMasterInstanceFleetLaunchSpecificationsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecMasterInstanceFleetLaunchSpecifications)(ptr) = ClusterSpecMasterInstanceFleetLaunchSpecifications{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecMasterInstanceFleetLaunchSpecifications

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecMasterInstanceFleetLaunchSpecifications{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecMasterInstanceFleetLaunchSpecifications)(ptr) = objs[0]
			} else {
				*(*ClusterSpecMasterInstanceFleetLaunchSpecifications)(ptr) = ClusterSpecMasterInstanceFleetLaunchSpecifications{}
			}
		} else {
			*(*ClusterSpecMasterInstanceFleetLaunchSpecifications)(ptr) = ClusterSpecMasterInstanceFleetLaunchSpecifications{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecMasterInstanceFleetLaunchSpecifications

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecMasterInstanceFleetLaunchSpecifications{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecMasterInstanceFleetLaunchSpecifications)(ptr) = obj
		} else {
			*(*ClusterSpecMasterInstanceFleetLaunchSpecifications)(ptr) = ClusterSpecMasterInstanceFleetLaunchSpecifications{}
		}
	default:
		iter.ReportError("decode ClusterSpecMasterInstanceFleetLaunchSpecifications", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClusterSpecMasterInstanceGroupCodec struct {
}

func (ClusterSpecMasterInstanceGroupCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecMasterInstanceGroup)(ptr) == nil
}

func (ClusterSpecMasterInstanceGroupCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecMasterInstanceGroup)(ptr)
	var objs []ClusterSpecMasterInstanceGroup
	if obj != nil {
		objs = []ClusterSpecMasterInstanceGroup{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecMasterInstanceGroup{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecMasterInstanceGroupCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecMasterInstanceGroup)(ptr) = ClusterSpecMasterInstanceGroup{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecMasterInstanceGroup

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecMasterInstanceGroup{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecMasterInstanceGroup)(ptr) = objs[0]
			} else {
				*(*ClusterSpecMasterInstanceGroup)(ptr) = ClusterSpecMasterInstanceGroup{}
			}
		} else {
			*(*ClusterSpecMasterInstanceGroup)(ptr) = ClusterSpecMasterInstanceGroup{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecMasterInstanceGroup

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecMasterInstanceGroup{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecMasterInstanceGroup)(ptr) = obj
		} else {
			*(*ClusterSpecMasterInstanceGroup)(ptr) = ClusterSpecMasterInstanceGroup{}
		}
	default:
		iter.ReportError("decode ClusterSpecMasterInstanceGroup", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClusterSpecStepHadoopJarStepCodec struct {
}

func (ClusterSpecStepHadoopJarStepCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecStepHadoopJarStep)(ptr) == nil
}

func (ClusterSpecStepHadoopJarStepCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecStepHadoopJarStep)(ptr)
	var objs []ClusterSpecStepHadoopJarStep
	if obj != nil {
		objs = []ClusterSpecStepHadoopJarStep{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecStepHadoopJarStep{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecStepHadoopJarStepCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecStepHadoopJarStep)(ptr) = ClusterSpecStepHadoopJarStep{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecStepHadoopJarStep

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecStepHadoopJarStep{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecStepHadoopJarStep)(ptr) = objs[0]
			} else {
				*(*ClusterSpecStepHadoopJarStep)(ptr) = ClusterSpecStepHadoopJarStep{}
			}
		} else {
			*(*ClusterSpecStepHadoopJarStep)(ptr) = ClusterSpecStepHadoopJarStep{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecStepHadoopJarStep

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecStepHadoopJarStep{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecStepHadoopJarStep)(ptr) = obj
		} else {
			*(*ClusterSpecStepHadoopJarStep)(ptr) = ClusterSpecStepHadoopJarStep{}
		}
	default:
		iter.ReportError("decode ClusterSpecStepHadoopJarStep", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type InstanceFleetSpecLaunchSpecificationsCodec struct {
}

func (InstanceFleetSpecLaunchSpecificationsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InstanceFleetSpecLaunchSpecifications)(ptr) == nil
}

func (InstanceFleetSpecLaunchSpecificationsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InstanceFleetSpecLaunchSpecifications)(ptr)
	var objs []InstanceFleetSpecLaunchSpecifications
	if obj != nil {
		objs = []InstanceFleetSpecLaunchSpecifications{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceFleetSpecLaunchSpecifications{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InstanceFleetSpecLaunchSpecificationsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InstanceFleetSpecLaunchSpecifications)(ptr) = InstanceFleetSpecLaunchSpecifications{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InstanceFleetSpecLaunchSpecifications

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceFleetSpecLaunchSpecifications{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InstanceFleetSpecLaunchSpecifications)(ptr) = objs[0]
			} else {
				*(*InstanceFleetSpecLaunchSpecifications)(ptr) = InstanceFleetSpecLaunchSpecifications{}
			}
		} else {
			*(*InstanceFleetSpecLaunchSpecifications)(ptr) = InstanceFleetSpecLaunchSpecifications{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj InstanceFleetSpecLaunchSpecifications

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceFleetSpecLaunchSpecifications{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*InstanceFleetSpecLaunchSpecifications)(ptr) = obj
		} else {
			*(*InstanceFleetSpecLaunchSpecifications)(ptr) = InstanceFleetSpecLaunchSpecifications{}
		}
	default:
		iter.ReportError("decode InstanceFleetSpecLaunchSpecifications", "unexpected JSON type")
	}
}
