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
		jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinition{}).Type1()):                                  ClassificationJobSpecS3JobDefinitionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScoping{}).Type1()):                           ClassificationJobSpecS3JobDefinitionScopingCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingExcludes{}).Type1()):                   ClassificationJobSpecS3JobDefinitionScopingExcludesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm{}).Type1()): ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTermCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm{}).Type1()):    ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTermCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingIncludes{}).Type1()):                   ClassificationJobSpecS3JobDefinitionScopingIncludesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm{}).Type1()): ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTermCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm{}).Type1()):    ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTermCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecScheduleFrequency{}).Type1()):                                ClassificationJobSpecScheduleFrequencyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FindingsFilterSpecFindingCriteria{}).Type1()):                                     FindingsFilterSpecFindingCriteriaCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinition{}).Type1()):                                  ClassificationJobSpecS3JobDefinitionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScoping{}).Type1()):                           ClassificationJobSpecS3JobDefinitionScopingCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingExcludes{}).Type1()):                   ClassificationJobSpecS3JobDefinitionScopingExcludesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm{}).Type1()): ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTermCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm{}).Type1()):    ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTermCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingIncludes{}).Type1()):                   ClassificationJobSpecS3JobDefinitionScopingIncludesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm{}).Type1()): ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTermCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm{}).Type1()):    ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTermCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecScheduleFrequency{}).Type1()):                                ClassificationJobSpecScheduleFrequencyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FindingsFilterSpecFindingCriteria{}).Type1()):                                     FindingsFilterSpecFindingCriteriaCodec{},
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
type ClassificationJobSpecS3JobDefinitionCodec struct {
}

func (ClassificationJobSpecS3JobDefinitionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClassificationJobSpecS3JobDefinition)(ptr) == nil
}

func (ClassificationJobSpecS3JobDefinitionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClassificationJobSpecS3JobDefinition)(ptr)
	var objs []ClassificationJobSpecS3JobDefinition
	if obj != nil {
		objs = []ClassificationJobSpecS3JobDefinition{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinition{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClassificationJobSpecS3JobDefinitionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClassificationJobSpecS3JobDefinition)(ptr) = ClassificationJobSpecS3JobDefinition{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClassificationJobSpecS3JobDefinition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClassificationJobSpecS3JobDefinition)(ptr) = objs[0]
			} else {
				*(*ClassificationJobSpecS3JobDefinition)(ptr) = ClassificationJobSpecS3JobDefinition{}
			}
		} else {
			*(*ClassificationJobSpecS3JobDefinition)(ptr) = ClassificationJobSpecS3JobDefinition{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClassificationJobSpecS3JobDefinition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClassificationJobSpecS3JobDefinition)(ptr) = obj
		} else {
			*(*ClassificationJobSpecS3JobDefinition)(ptr) = ClassificationJobSpecS3JobDefinition{}
		}
	default:
		iter.ReportError("decode ClassificationJobSpecS3JobDefinition", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClassificationJobSpecS3JobDefinitionScopingCodec struct {
}

func (ClassificationJobSpecS3JobDefinitionScopingCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClassificationJobSpecS3JobDefinitionScoping)(ptr) == nil
}

func (ClassificationJobSpecS3JobDefinitionScopingCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClassificationJobSpecS3JobDefinitionScoping)(ptr)
	var objs []ClassificationJobSpecS3JobDefinitionScoping
	if obj != nil {
		objs = []ClassificationJobSpecS3JobDefinitionScoping{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScoping{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClassificationJobSpecS3JobDefinitionScopingCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClassificationJobSpecS3JobDefinitionScoping)(ptr) = ClassificationJobSpecS3JobDefinitionScoping{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClassificationJobSpecS3JobDefinitionScoping

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScoping{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClassificationJobSpecS3JobDefinitionScoping)(ptr) = objs[0]
			} else {
				*(*ClassificationJobSpecS3JobDefinitionScoping)(ptr) = ClassificationJobSpecS3JobDefinitionScoping{}
			}
		} else {
			*(*ClassificationJobSpecS3JobDefinitionScoping)(ptr) = ClassificationJobSpecS3JobDefinitionScoping{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClassificationJobSpecS3JobDefinitionScoping

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScoping{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClassificationJobSpecS3JobDefinitionScoping)(ptr) = obj
		} else {
			*(*ClassificationJobSpecS3JobDefinitionScoping)(ptr) = ClassificationJobSpecS3JobDefinitionScoping{}
		}
	default:
		iter.ReportError("decode ClassificationJobSpecS3JobDefinitionScoping", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClassificationJobSpecS3JobDefinitionScopingExcludesCodec struct {
}

func (ClassificationJobSpecS3JobDefinitionScopingExcludesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClassificationJobSpecS3JobDefinitionScopingExcludes)(ptr) == nil
}

func (ClassificationJobSpecS3JobDefinitionScopingExcludesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClassificationJobSpecS3JobDefinitionScopingExcludes)(ptr)
	var objs []ClassificationJobSpecS3JobDefinitionScopingExcludes
	if obj != nil {
		objs = []ClassificationJobSpecS3JobDefinitionScopingExcludes{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingExcludes{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClassificationJobSpecS3JobDefinitionScopingExcludesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClassificationJobSpecS3JobDefinitionScopingExcludes)(ptr) = ClassificationJobSpecS3JobDefinitionScopingExcludes{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClassificationJobSpecS3JobDefinitionScopingExcludes

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingExcludes{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClassificationJobSpecS3JobDefinitionScopingExcludes)(ptr) = objs[0]
			} else {
				*(*ClassificationJobSpecS3JobDefinitionScopingExcludes)(ptr) = ClassificationJobSpecS3JobDefinitionScopingExcludes{}
			}
		} else {
			*(*ClassificationJobSpecS3JobDefinitionScopingExcludes)(ptr) = ClassificationJobSpecS3JobDefinitionScopingExcludes{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClassificationJobSpecS3JobDefinitionScopingExcludes

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingExcludes{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClassificationJobSpecS3JobDefinitionScopingExcludes)(ptr) = obj
		} else {
			*(*ClassificationJobSpecS3JobDefinitionScopingExcludes)(ptr) = ClassificationJobSpecS3JobDefinitionScopingExcludes{}
		}
	default:
		iter.ReportError("decode ClassificationJobSpecS3JobDefinitionScopingExcludes", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTermCodec struct {
}

func (ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTermCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm)(ptr) == nil
}

func (ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTermCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm)(ptr)
	var objs []ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm
	if obj != nil {
		objs = []ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTermCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm)(ptr) = ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm)(ptr) = objs[0]
			} else {
				*(*ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm)(ptr) = ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm{}
			}
		} else {
			*(*ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm)(ptr) = ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm)(ptr) = obj
		} else {
			*(*ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm)(ptr) = ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm{}
		}
	default:
		iter.ReportError("decode ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTermCodec struct {
}

func (ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTermCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm)(ptr) == nil
}

func (ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTermCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm)(ptr)
	var objs []ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm
	if obj != nil {
		objs = []ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTermCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm)(ptr) = ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm)(ptr) = objs[0]
			} else {
				*(*ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm)(ptr) = ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm{}
			}
		} else {
			*(*ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm)(ptr) = ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm)(ptr) = obj
		} else {
			*(*ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm)(ptr) = ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm{}
		}
	default:
		iter.ReportError("decode ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClassificationJobSpecS3JobDefinitionScopingIncludesCodec struct {
}

func (ClassificationJobSpecS3JobDefinitionScopingIncludesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClassificationJobSpecS3JobDefinitionScopingIncludes)(ptr) == nil
}

func (ClassificationJobSpecS3JobDefinitionScopingIncludesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClassificationJobSpecS3JobDefinitionScopingIncludes)(ptr)
	var objs []ClassificationJobSpecS3JobDefinitionScopingIncludes
	if obj != nil {
		objs = []ClassificationJobSpecS3JobDefinitionScopingIncludes{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingIncludes{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClassificationJobSpecS3JobDefinitionScopingIncludesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClassificationJobSpecS3JobDefinitionScopingIncludes)(ptr) = ClassificationJobSpecS3JobDefinitionScopingIncludes{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClassificationJobSpecS3JobDefinitionScopingIncludes

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingIncludes{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClassificationJobSpecS3JobDefinitionScopingIncludes)(ptr) = objs[0]
			} else {
				*(*ClassificationJobSpecS3JobDefinitionScopingIncludes)(ptr) = ClassificationJobSpecS3JobDefinitionScopingIncludes{}
			}
		} else {
			*(*ClassificationJobSpecS3JobDefinitionScopingIncludes)(ptr) = ClassificationJobSpecS3JobDefinitionScopingIncludes{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClassificationJobSpecS3JobDefinitionScopingIncludes

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingIncludes{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClassificationJobSpecS3JobDefinitionScopingIncludes)(ptr) = obj
		} else {
			*(*ClassificationJobSpecS3JobDefinitionScopingIncludes)(ptr) = ClassificationJobSpecS3JobDefinitionScopingIncludes{}
		}
	default:
		iter.ReportError("decode ClassificationJobSpecS3JobDefinitionScopingIncludes", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTermCodec struct {
}

func (ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTermCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm)(ptr) == nil
}

func (ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTermCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm)(ptr)
	var objs []ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm
	if obj != nil {
		objs = []ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTermCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm)(ptr) = ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm)(ptr) = objs[0]
			} else {
				*(*ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm)(ptr) = ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm{}
			}
		} else {
			*(*ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm)(ptr) = ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm)(ptr) = obj
		} else {
			*(*ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm)(ptr) = ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm{}
		}
	default:
		iter.ReportError("decode ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTermCodec struct {
}

func (ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTermCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm)(ptr) == nil
}

func (ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTermCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm)(ptr)
	var objs []ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm
	if obj != nil {
		objs = []ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTermCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm)(ptr) = ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm)(ptr) = objs[0]
			} else {
				*(*ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm)(ptr) = ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm{}
			}
		} else {
			*(*ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm)(ptr) = ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm)(ptr) = obj
		} else {
			*(*ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm)(ptr) = ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm{}
		}
	default:
		iter.ReportError("decode ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClassificationJobSpecScheduleFrequencyCodec struct {
}

func (ClassificationJobSpecScheduleFrequencyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClassificationJobSpecScheduleFrequency)(ptr) == nil
}

func (ClassificationJobSpecScheduleFrequencyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClassificationJobSpecScheduleFrequency)(ptr)
	var objs []ClassificationJobSpecScheduleFrequency
	if obj != nil {
		objs = []ClassificationJobSpecScheduleFrequency{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecScheduleFrequency{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClassificationJobSpecScheduleFrequencyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClassificationJobSpecScheduleFrequency)(ptr) = ClassificationJobSpecScheduleFrequency{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClassificationJobSpecScheduleFrequency

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecScheduleFrequency{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClassificationJobSpecScheduleFrequency)(ptr) = objs[0]
			} else {
				*(*ClassificationJobSpecScheduleFrequency)(ptr) = ClassificationJobSpecScheduleFrequency{}
			}
		} else {
			*(*ClassificationJobSpecScheduleFrequency)(ptr) = ClassificationJobSpecScheduleFrequency{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClassificationJobSpecScheduleFrequency

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClassificationJobSpecScheduleFrequency{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClassificationJobSpecScheduleFrequency)(ptr) = obj
		} else {
			*(*ClassificationJobSpecScheduleFrequency)(ptr) = ClassificationJobSpecScheduleFrequency{}
		}
	default:
		iter.ReportError("decode ClassificationJobSpecScheduleFrequency", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type FindingsFilterSpecFindingCriteriaCodec struct {
}

func (FindingsFilterSpecFindingCriteriaCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*FindingsFilterSpecFindingCriteria)(ptr) == nil
}

func (FindingsFilterSpecFindingCriteriaCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*FindingsFilterSpecFindingCriteria)(ptr)
	var objs []FindingsFilterSpecFindingCriteria
	if obj != nil {
		objs = []FindingsFilterSpecFindingCriteria{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FindingsFilterSpecFindingCriteria{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (FindingsFilterSpecFindingCriteriaCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*FindingsFilterSpecFindingCriteria)(ptr) = FindingsFilterSpecFindingCriteria{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []FindingsFilterSpecFindingCriteria

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FindingsFilterSpecFindingCriteria{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*FindingsFilterSpecFindingCriteria)(ptr) = objs[0]
			} else {
				*(*FindingsFilterSpecFindingCriteria)(ptr) = FindingsFilterSpecFindingCriteria{}
			}
		} else {
			*(*FindingsFilterSpecFindingCriteria)(ptr) = FindingsFilterSpecFindingCriteria{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj FindingsFilterSpecFindingCriteria

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FindingsFilterSpecFindingCriteria{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*FindingsFilterSpecFindingCriteria)(ptr) = obj
		} else {
			*(*FindingsFilterSpecFindingCriteria)(ptr) = FindingsFilterSpecFindingCriteria{}
		}
	default:
		iter.ReportError("decode FindingsFilterSpecFindingCriteria", "unexpected JSON type")
	}
}
