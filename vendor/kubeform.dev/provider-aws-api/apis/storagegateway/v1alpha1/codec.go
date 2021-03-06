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
		jsoniter.MustGetKind(reflect2.TypeOf(FileSystemAssociationSpecCacheAttributes{}).Type1()): FileSystemAssociationSpecCacheAttributesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecSmbActiveDirectorySettings{}).Type1()):    GatewaySpecSmbActiveDirectorySettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(NfsFileShareSpecCacheAttributes{}).Type1()):          NfsFileShareSpecCacheAttributesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(NfsFileShareSpecNfsFileShareDefaults{}).Type1()):     NfsFileShareSpecNfsFileShareDefaultsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SmbFileShareSpecCacheAttributes{}).Type1()):          SmbFileShareSpecCacheAttributesCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(FileSystemAssociationSpecCacheAttributes{}).Type1()): FileSystemAssociationSpecCacheAttributesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecSmbActiveDirectorySettings{}).Type1()):    GatewaySpecSmbActiveDirectorySettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(NfsFileShareSpecCacheAttributes{}).Type1()):          NfsFileShareSpecCacheAttributesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(NfsFileShareSpecNfsFileShareDefaults{}).Type1()):     NfsFileShareSpecNfsFileShareDefaultsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SmbFileShareSpecCacheAttributes{}).Type1()):          SmbFileShareSpecCacheAttributesCodec{},
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
type FileSystemAssociationSpecCacheAttributesCodec struct {
}

func (FileSystemAssociationSpecCacheAttributesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*FileSystemAssociationSpecCacheAttributes)(ptr) == nil
}

func (FileSystemAssociationSpecCacheAttributesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*FileSystemAssociationSpecCacheAttributes)(ptr)
	var objs []FileSystemAssociationSpecCacheAttributes
	if obj != nil {
		objs = []FileSystemAssociationSpecCacheAttributes{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FileSystemAssociationSpecCacheAttributes{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (FileSystemAssociationSpecCacheAttributesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*FileSystemAssociationSpecCacheAttributes)(ptr) = FileSystemAssociationSpecCacheAttributes{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []FileSystemAssociationSpecCacheAttributes

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FileSystemAssociationSpecCacheAttributes{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*FileSystemAssociationSpecCacheAttributes)(ptr) = objs[0]
			} else {
				*(*FileSystemAssociationSpecCacheAttributes)(ptr) = FileSystemAssociationSpecCacheAttributes{}
			}
		} else {
			*(*FileSystemAssociationSpecCacheAttributes)(ptr) = FileSystemAssociationSpecCacheAttributes{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj FileSystemAssociationSpecCacheAttributes

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FileSystemAssociationSpecCacheAttributes{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*FileSystemAssociationSpecCacheAttributes)(ptr) = obj
		} else {
			*(*FileSystemAssociationSpecCacheAttributes)(ptr) = FileSystemAssociationSpecCacheAttributes{}
		}
	default:
		iter.ReportError("decode FileSystemAssociationSpecCacheAttributes", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GatewaySpecSmbActiveDirectorySettingsCodec struct {
}

func (GatewaySpecSmbActiveDirectorySettingsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewaySpecSmbActiveDirectorySettings)(ptr) == nil
}

func (GatewaySpecSmbActiveDirectorySettingsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewaySpecSmbActiveDirectorySettings)(ptr)
	var objs []GatewaySpecSmbActiveDirectorySettings
	if obj != nil {
		objs = []GatewaySpecSmbActiveDirectorySettings{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecSmbActiveDirectorySettings{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewaySpecSmbActiveDirectorySettingsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewaySpecSmbActiveDirectorySettings)(ptr) = GatewaySpecSmbActiveDirectorySettings{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewaySpecSmbActiveDirectorySettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecSmbActiveDirectorySettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewaySpecSmbActiveDirectorySettings)(ptr) = objs[0]
			} else {
				*(*GatewaySpecSmbActiveDirectorySettings)(ptr) = GatewaySpecSmbActiveDirectorySettings{}
			}
		} else {
			*(*GatewaySpecSmbActiveDirectorySettings)(ptr) = GatewaySpecSmbActiveDirectorySettings{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewaySpecSmbActiveDirectorySettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecSmbActiveDirectorySettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewaySpecSmbActiveDirectorySettings)(ptr) = obj
		} else {
			*(*GatewaySpecSmbActiveDirectorySettings)(ptr) = GatewaySpecSmbActiveDirectorySettings{}
		}
	default:
		iter.ReportError("decode GatewaySpecSmbActiveDirectorySettings", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type NfsFileShareSpecCacheAttributesCodec struct {
}

func (NfsFileShareSpecCacheAttributesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*NfsFileShareSpecCacheAttributes)(ptr) == nil
}

func (NfsFileShareSpecCacheAttributesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*NfsFileShareSpecCacheAttributes)(ptr)
	var objs []NfsFileShareSpecCacheAttributes
	if obj != nil {
		objs = []NfsFileShareSpecCacheAttributes{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NfsFileShareSpecCacheAttributes{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (NfsFileShareSpecCacheAttributesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*NfsFileShareSpecCacheAttributes)(ptr) = NfsFileShareSpecCacheAttributes{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []NfsFileShareSpecCacheAttributes

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NfsFileShareSpecCacheAttributes{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*NfsFileShareSpecCacheAttributes)(ptr) = objs[0]
			} else {
				*(*NfsFileShareSpecCacheAttributes)(ptr) = NfsFileShareSpecCacheAttributes{}
			}
		} else {
			*(*NfsFileShareSpecCacheAttributes)(ptr) = NfsFileShareSpecCacheAttributes{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj NfsFileShareSpecCacheAttributes

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NfsFileShareSpecCacheAttributes{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*NfsFileShareSpecCacheAttributes)(ptr) = obj
		} else {
			*(*NfsFileShareSpecCacheAttributes)(ptr) = NfsFileShareSpecCacheAttributes{}
		}
	default:
		iter.ReportError("decode NfsFileShareSpecCacheAttributes", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type NfsFileShareSpecNfsFileShareDefaultsCodec struct {
}

func (NfsFileShareSpecNfsFileShareDefaultsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*NfsFileShareSpecNfsFileShareDefaults)(ptr) == nil
}

func (NfsFileShareSpecNfsFileShareDefaultsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*NfsFileShareSpecNfsFileShareDefaults)(ptr)
	var objs []NfsFileShareSpecNfsFileShareDefaults
	if obj != nil {
		objs = []NfsFileShareSpecNfsFileShareDefaults{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NfsFileShareSpecNfsFileShareDefaults{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (NfsFileShareSpecNfsFileShareDefaultsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*NfsFileShareSpecNfsFileShareDefaults)(ptr) = NfsFileShareSpecNfsFileShareDefaults{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []NfsFileShareSpecNfsFileShareDefaults

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NfsFileShareSpecNfsFileShareDefaults{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*NfsFileShareSpecNfsFileShareDefaults)(ptr) = objs[0]
			} else {
				*(*NfsFileShareSpecNfsFileShareDefaults)(ptr) = NfsFileShareSpecNfsFileShareDefaults{}
			}
		} else {
			*(*NfsFileShareSpecNfsFileShareDefaults)(ptr) = NfsFileShareSpecNfsFileShareDefaults{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj NfsFileShareSpecNfsFileShareDefaults

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NfsFileShareSpecNfsFileShareDefaults{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*NfsFileShareSpecNfsFileShareDefaults)(ptr) = obj
		} else {
			*(*NfsFileShareSpecNfsFileShareDefaults)(ptr) = NfsFileShareSpecNfsFileShareDefaults{}
		}
	default:
		iter.ReportError("decode NfsFileShareSpecNfsFileShareDefaults", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type SmbFileShareSpecCacheAttributesCodec struct {
}

func (SmbFileShareSpecCacheAttributesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*SmbFileShareSpecCacheAttributes)(ptr) == nil
}

func (SmbFileShareSpecCacheAttributesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*SmbFileShareSpecCacheAttributes)(ptr)
	var objs []SmbFileShareSpecCacheAttributes
	if obj != nil {
		objs = []SmbFileShareSpecCacheAttributes{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SmbFileShareSpecCacheAttributes{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (SmbFileShareSpecCacheAttributesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*SmbFileShareSpecCacheAttributes)(ptr) = SmbFileShareSpecCacheAttributes{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []SmbFileShareSpecCacheAttributes

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SmbFileShareSpecCacheAttributes{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*SmbFileShareSpecCacheAttributes)(ptr) = objs[0]
			} else {
				*(*SmbFileShareSpecCacheAttributes)(ptr) = SmbFileShareSpecCacheAttributes{}
			}
		} else {
			*(*SmbFileShareSpecCacheAttributes)(ptr) = SmbFileShareSpecCacheAttributes{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj SmbFileShareSpecCacheAttributes

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SmbFileShareSpecCacheAttributes{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*SmbFileShareSpecCacheAttributes)(ptr) = obj
		} else {
			*(*SmbFileShareSpecCacheAttributes)(ptr) = SmbFileShareSpecCacheAttributes{}
		}
	default:
		iter.ReportError("decode SmbFileShareSpecCacheAttributes", "unexpected JSON type")
	}
}
