// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/assistant/embedded/v1alpha1/embedded_assistant.proto

/*
Package embedded is a generated protocol buffer package.

It is generated from these files:
	google/assistant/embedded/v1alpha1/embedded_assistant.proto

It has these top-level messages:
	ConverseConfig
	AudioInConfig
	AudioOutConfig
	ConverseState
	AudioOut
	ConverseResult
	ConverseRequest
	ConverseResponse
*/
package embedded

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_rpc "google.golang.org/genproto/googleapis/rpc/status"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Audio encoding of the data sent in the audio message.
// Audio must be one-channel (mono). The only language supported is "en-US".
type AudioInConfig_Encoding int32

const (
	// Not specified. Will return result [google.rpc.Code.INVALID_ARGUMENT][].
	AudioInConfig_ENCODING_UNSPECIFIED AudioInConfig_Encoding = 0
	// Uncompressed 16-bit signed little-endian samples (Linear PCM).
	// This encoding includes no header, only the raw audio bytes.
	AudioInConfig_LINEAR16 AudioInConfig_Encoding = 1
	// [`FLAC`](https://xiph.org/flac/documentation.html) (Free Lossless Audio
	// Codec) is the recommended encoding because it is
	// lossless--therefore recognition is not compromised--and
	// requires only about half the bandwidth of `LINEAR16`. This encoding
	// includes the `FLAC` stream header followed by audio data. It supports
	// 16-bit and 24-bit samples, however, not all fields in `STREAMINFO` are
	// supported.
	AudioInConfig_FLAC AudioInConfig_Encoding = 2
)

var AudioInConfig_Encoding_name = map[int32]string{
	0: "ENCODING_UNSPECIFIED",
	1: "LINEAR16",
	2: "FLAC",
}
var AudioInConfig_Encoding_value = map[string]int32{
	"ENCODING_UNSPECIFIED": 0,
	"LINEAR16":             1,
	"FLAC":                 2,
}

func (x AudioInConfig_Encoding) String() string {
	return proto.EnumName(AudioInConfig_Encoding_name, int32(x))
}
func (AudioInConfig_Encoding) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

// Audio encoding of the data returned in the audio message. All encodings are
// raw audio bytes with no header, except as indicated below.
type AudioOutConfig_Encoding int32

const (
	// Not specified. Will return result [google.rpc.Code.INVALID_ARGUMENT][].
	AudioOutConfig_ENCODING_UNSPECIFIED AudioOutConfig_Encoding = 0
	// Uncompressed 16-bit signed little-endian samples (Linear PCM).
	AudioOutConfig_LINEAR16 AudioOutConfig_Encoding = 1
	// MP3 audio encoding. The sample rate is encoded in the payload.
	AudioOutConfig_MP3 AudioOutConfig_Encoding = 2
	// Opus-encoded audio wrapped in an ogg container. The result will be a
	// file which can be played natively on Android and in some browsers (such
	// as Chrome). The quality of the encoding is considerably higher than MP3
	// while using the same bitrate. The sample rate is encoded in the payload.
	AudioOutConfig_OPUS_IN_OGG AudioOutConfig_Encoding = 3
)

var AudioOutConfig_Encoding_name = map[int32]string{
	0: "ENCODING_UNSPECIFIED",
	1: "LINEAR16",
	2: "MP3",
	3: "OPUS_IN_OGG",
}
var AudioOutConfig_Encoding_value = map[string]int32{
	"ENCODING_UNSPECIFIED": 0,
	"LINEAR16":             1,
	"MP3":                  2,
	"OPUS_IN_OGG":          3,
}

func (x AudioOutConfig_Encoding) String() string {
	return proto.EnumName(AudioOutConfig_Encoding_name, int32(x))
}
func (AudioOutConfig_Encoding) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

// Possible states of the microphone after a `Converse` RPC completes.
type ConverseResult_MicrophoneMode int32

const (
	// No mode specified.
	ConverseResult_MICROPHONE_MODE_UNSPECIFIED ConverseResult_MicrophoneMode = 0
	// The service is not expecting a follow-on question from the user.
	// The microphone should remain off until the user re-activates it.
	ConverseResult_CLOSE_MICROPHONE ConverseResult_MicrophoneMode = 1
	// The service is expecting a follow-on question from the user. The
	// microphone should be re-opened when the `AudioOut` playback completes
	// (by starting a new `Converse` RPC call to send the new audio).
	ConverseResult_DIALOG_FOLLOW_ON ConverseResult_MicrophoneMode = 2
)

var ConverseResult_MicrophoneMode_name = map[int32]string{
	0: "MICROPHONE_MODE_UNSPECIFIED",
	1: "CLOSE_MICROPHONE",
	2: "DIALOG_FOLLOW_ON",
}
var ConverseResult_MicrophoneMode_value = map[string]int32{
	"MICROPHONE_MODE_UNSPECIFIED": 0,
	"CLOSE_MICROPHONE":            1,
	"DIALOG_FOLLOW_ON":            2,
}

func (x ConverseResult_MicrophoneMode) String() string {
	return proto.EnumName(ConverseResult_MicrophoneMode_name, int32(x))
}
func (ConverseResult_MicrophoneMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0}
}

// Indicates the type of event.
type ConverseResponse_EventType int32

const (
	// No event specified.
	ConverseResponse_EVENT_TYPE_UNSPECIFIED ConverseResponse_EventType = 0
	// This event indicates that the server has detected the end of the user's
	// speech utterance and expects no additional speech. Therefore, the server
	// will not process additional audio (although it may subsequently return
	// additional results). The client should stop sending additional audio
	// data, half-close the gRPC connection, and wait for any additional results
	// until the server closes the gRPC connection.
	ConverseResponse_END_OF_UTTERANCE ConverseResponse_EventType = 1
)

var ConverseResponse_EventType_name = map[int32]string{
	0: "EVENT_TYPE_UNSPECIFIED",
	1: "END_OF_UTTERANCE",
}
var ConverseResponse_EventType_value = map[string]int32{
	"EVENT_TYPE_UNSPECIFIED": 0,
	"END_OF_UTTERANCE":       1,
}

func (x ConverseResponse_EventType) String() string {
	return proto.EnumName(ConverseResponse_EventType_name, int32(x))
}
func (ConverseResponse_EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{7, 0}
}

// Specifies how to process the `ConverseRequest` messages.
type ConverseConfig struct {
	// *Required* Specifies how to process the subsequent incoming audio.
	AudioInConfig *AudioInConfig `protobuf:"bytes,1,opt,name=audio_in_config,json=audioInConfig" json:"audio_in_config,omitempty"`
	// *Required* Specifies how to format the audio that will be returned.
	AudioOutConfig *AudioOutConfig `protobuf:"bytes,2,opt,name=audio_out_config,json=audioOutConfig" json:"audio_out_config,omitempty"`
	// *Required* Represents the current dialog state.
	ConverseState *ConverseState `protobuf:"bytes,3,opt,name=converse_state,json=converseState" json:"converse_state,omitempty"`
}

func (m *ConverseConfig) Reset()                    { *m = ConverseConfig{} }
func (m *ConverseConfig) String() string            { return proto.CompactTextString(m) }
func (*ConverseConfig) ProtoMessage()               {}
func (*ConverseConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ConverseConfig) GetAudioInConfig() *AudioInConfig {
	if m != nil {
		return m.AudioInConfig
	}
	return nil
}

func (m *ConverseConfig) GetAudioOutConfig() *AudioOutConfig {
	if m != nil {
		return m.AudioOutConfig
	}
	return nil
}

func (m *ConverseConfig) GetConverseState() *ConverseState {
	if m != nil {
		return m.ConverseState
	}
	return nil
}

// Specifies how to process the `audio_in` data that will be provided in
// subsequent requests. For recommended settings, see the Google Assistant SDK
// [best practices](https://developers.google.com/assistant/sdk/develop/grpc/best-practices/audio).
type AudioInConfig struct {
	// *Required* Encoding of audio data sent in all `audio_in` messages.
	Encoding AudioInConfig_Encoding `protobuf:"varint,1,opt,name=encoding,enum=google.assistant.embedded.v1alpha1.AudioInConfig_Encoding" json:"encoding,omitempty"`
	// *Required* Sample rate (in Hertz) of the audio data sent in all `audio_in`
	// messages. Valid values are from 16000-24000, but 16000 is optimal.
	// For best results, set the sampling rate of the audio source to 16000 Hz.
	// If that's not possible, use the native sample rate of the audio source
	// (instead of re-sampling).
	SampleRateHertz int32 `protobuf:"varint,2,opt,name=sample_rate_hertz,json=sampleRateHertz" json:"sample_rate_hertz,omitempty"`
}

func (m *AudioInConfig) Reset()                    { *m = AudioInConfig{} }
func (m *AudioInConfig) String() string            { return proto.CompactTextString(m) }
func (*AudioInConfig) ProtoMessage()               {}
func (*AudioInConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AudioInConfig) GetEncoding() AudioInConfig_Encoding {
	if m != nil {
		return m.Encoding
	}
	return AudioInConfig_ENCODING_UNSPECIFIED
}

func (m *AudioInConfig) GetSampleRateHertz() int32 {
	if m != nil {
		return m.SampleRateHertz
	}
	return 0
}

// Specifies the desired format for the server to use when it returns
// `audio_out` messages.
type AudioOutConfig struct {
	// *Required* The encoding of audio data to be returned in all `audio_out`
	// messages.
	Encoding AudioOutConfig_Encoding `protobuf:"varint,1,opt,name=encoding,enum=google.assistant.embedded.v1alpha1.AudioOutConfig_Encoding" json:"encoding,omitempty"`
	// *Required* The sample rate in Hertz of the audio data returned in
	// `audio_out` messages. Valid values are: 16000-24000.
	SampleRateHertz int32 `protobuf:"varint,2,opt,name=sample_rate_hertz,json=sampleRateHertz" json:"sample_rate_hertz,omitempty"`
	// *Required* Current volume setting of the device's audio output.
	// Valid values are 1 to 100 (corresponding to 1% to 100%).
	VolumePercentage int32 `protobuf:"varint,3,opt,name=volume_percentage,json=volumePercentage" json:"volume_percentage,omitempty"`
}

func (m *AudioOutConfig) Reset()                    { *m = AudioOutConfig{} }
func (m *AudioOutConfig) String() string            { return proto.CompactTextString(m) }
func (*AudioOutConfig) ProtoMessage()               {}
func (*AudioOutConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AudioOutConfig) GetEncoding() AudioOutConfig_Encoding {
	if m != nil {
		return m.Encoding
	}
	return AudioOutConfig_ENCODING_UNSPECIFIED
}

func (m *AudioOutConfig) GetSampleRateHertz() int32 {
	if m != nil {
		return m.SampleRateHertz
	}
	return 0
}

func (m *AudioOutConfig) GetVolumePercentage() int32 {
	if m != nil {
		return m.VolumePercentage
	}
	return 0
}

// Provides information about the current dialog state.
type ConverseState struct {
	// *Required* The `conversation_state` value returned in the prior
	// `ConverseResponse`. Omit (do not set the field) if there was no prior
	// `ConverseResponse`. If there was a prior `ConverseResponse`, do not omit
	// this field; doing so will end that conversation (and this new request will
	// start a new conversation).
	ConversationState []byte `protobuf:"bytes,1,opt,name=conversation_state,json=conversationState,proto3" json:"conversation_state,omitempty"`
}

func (m *ConverseState) Reset()                    { *m = ConverseState{} }
func (m *ConverseState) String() string            { return proto.CompactTextString(m) }
func (*ConverseState) ProtoMessage()               {}
func (*ConverseState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ConverseState) GetConversationState() []byte {
	if m != nil {
		return m.ConversationState
	}
	return nil
}

// The audio containing the assistant's response to the query. Sequential chunks
// of audio data are received in sequential `ConverseResponse` messages.
type AudioOut struct {
	// *Output-only* The audio data containing the assistant's response to the
	// query. Sequential chunks of audio data are received in sequential
	// `ConverseResponse` messages.
	AudioData []byte `protobuf:"bytes,1,opt,name=audio_data,json=audioData,proto3" json:"audio_data,omitempty"`
}

func (m *AudioOut) Reset()                    { *m = AudioOut{} }
func (m *AudioOut) String() string            { return proto.CompactTextString(m) }
func (*AudioOut) ProtoMessage()               {}
func (*AudioOut) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AudioOut) GetAudioData() []byte {
	if m != nil {
		return m.AudioData
	}
	return nil
}

// The semantic result for the user's spoken query.
type ConverseResult struct {
	// *Output-only* The recognized transcript of what the user said.
	SpokenRequestText string `protobuf:"bytes,1,opt,name=spoken_request_text,json=spokenRequestText" json:"spoken_request_text,omitempty"`
	// *Output-only* The text of the assistant's spoken response. This is only
	// returned for an IFTTT action.
	SpokenResponseText string `protobuf:"bytes,2,opt,name=spoken_response_text,json=spokenResponseText" json:"spoken_response_text,omitempty"`
	// *Output-only* State information for subsequent `ConverseRequest`. This
	// value should be saved in the client and returned in the
	// `conversation_state` with the next `ConverseRequest`. (The client does not
	// need to interpret or otherwise use this value.) There is no need to save
	// this information across device restarts.
	ConversationState []byte `protobuf:"bytes,3,opt,name=conversation_state,json=conversationState,proto3" json:"conversation_state,omitempty"`
	// *Output-only* Specifies the mode of the microphone after this `Converse`
	// RPC is processed.
	MicrophoneMode ConverseResult_MicrophoneMode `protobuf:"varint,4,opt,name=microphone_mode,json=microphoneMode,enum=google.assistant.embedded.v1alpha1.ConverseResult_MicrophoneMode" json:"microphone_mode,omitempty"`
	// *Output-only* Updated volume level. The value will be 0 or omitted
	// (indicating no change) unless a voice command such as "Increase the volume"
	// or "Set volume level 4" was recognized, in which case the value will be
	// between 1 and 100 (corresponding to the new volume level of 1% to 100%).
	// Typically, a client should use this volume level when playing the
	// `audio_out` data, and retain this value as the current volume level and
	// supply it in the `AudioOutConfig` of the next `ConverseRequest`. (Some
	// clients may also implement other ways to allow the current volume level to
	// be changed, for example, by providing a knob that the user can turn.)
	VolumePercentage int32 `protobuf:"varint,5,opt,name=volume_percentage,json=volumePercentage" json:"volume_percentage,omitempty"`
}

func (m *ConverseResult) Reset()                    { *m = ConverseResult{} }
func (m *ConverseResult) String() string            { return proto.CompactTextString(m) }
func (*ConverseResult) ProtoMessage()               {}
func (*ConverseResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ConverseResult) GetSpokenRequestText() string {
	if m != nil {
		return m.SpokenRequestText
	}
	return ""
}

func (m *ConverseResult) GetSpokenResponseText() string {
	if m != nil {
		return m.SpokenResponseText
	}
	return ""
}

func (m *ConverseResult) GetConversationState() []byte {
	if m != nil {
		return m.ConversationState
	}
	return nil
}

func (m *ConverseResult) GetMicrophoneMode() ConverseResult_MicrophoneMode {
	if m != nil {
		return m.MicrophoneMode
	}
	return ConverseResult_MICROPHONE_MODE_UNSPECIFIED
}

func (m *ConverseResult) GetVolumePercentage() int32 {
	if m != nil {
		return m.VolumePercentage
	}
	return 0
}

// The top-level message sent by the client. Clients must send at least two, and
// typically numerous `ConverseRequest` messages. The first message must
// contain a `config` message and must not contain `audio_in` data. All
// subsequent messages must contain `audio_in` data and must not contain a
// `config` message.
type ConverseRequest struct {
	// Exactly one of these fields must be specified in each `ConverseRequest`.
	//
	// Types that are valid to be assigned to ConverseRequest:
	//	*ConverseRequest_Config
	//	*ConverseRequest_AudioIn
	ConverseRequest isConverseRequest_ConverseRequest `protobuf_oneof:"converse_request"`
}

func (m *ConverseRequest) Reset()                    { *m = ConverseRequest{} }
func (m *ConverseRequest) String() string            { return proto.CompactTextString(m) }
func (*ConverseRequest) ProtoMessage()               {}
func (*ConverseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isConverseRequest_ConverseRequest interface {
	isConverseRequest_ConverseRequest()
}

type ConverseRequest_Config struct {
	Config *ConverseConfig `protobuf:"bytes,1,opt,name=config,oneof"`
}
type ConverseRequest_AudioIn struct {
	AudioIn []byte `protobuf:"bytes,2,opt,name=audio_in,json=audioIn,proto3,oneof"`
}

func (*ConverseRequest_Config) isConverseRequest_ConverseRequest()  {}
func (*ConverseRequest_AudioIn) isConverseRequest_ConverseRequest() {}

func (m *ConverseRequest) GetConverseRequest() isConverseRequest_ConverseRequest {
	if m != nil {
		return m.ConverseRequest
	}
	return nil
}

func (m *ConverseRequest) GetConfig() *ConverseConfig {
	if x, ok := m.GetConverseRequest().(*ConverseRequest_Config); ok {
		return x.Config
	}
	return nil
}

func (m *ConverseRequest) GetAudioIn() []byte {
	if x, ok := m.GetConverseRequest().(*ConverseRequest_AudioIn); ok {
		return x.AudioIn
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ConverseRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ConverseRequest_OneofMarshaler, _ConverseRequest_OneofUnmarshaler, _ConverseRequest_OneofSizer, []interface{}{
		(*ConverseRequest_Config)(nil),
		(*ConverseRequest_AudioIn)(nil),
	}
}

func _ConverseRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ConverseRequest)
	// converse_request
	switch x := m.ConverseRequest.(type) {
	case *ConverseRequest_Config:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Config); err != nil {
			return err
		}
	case *ConverseRequest_AudioIn:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.AudioIn)
	case nil:
	default:
		return fmt.Errorf("ConverseRequest.ConverseRequest has unexpected type %T", x)
	}
	return nil
}

func _ConverseRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ConverseRequest)
	switch tag {
	case 1: // converse_request.config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ConverseConfig)
		err := b.DecodeMessage(msg)
		m.ConverseRequest = &ConverseRequest_Config{msg}
		return true, err
	case 2: // converse_request.audio_in
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.ConverseRequest = &ConverseRequest_AudioIn{x}
		return true, err
	default:
		return false, nil
	}
}

func _ConverseRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ConverseRequest)
	// converse_request
	switch x := m.ConverseRequest.(type) {
	case *ConverseRequest_Config:
		s := proto.Size(x.Config)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ConverseRequest_AudioIn:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.AudioIn)))
		n += len(x.AudioIn)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The top-level message received by the client. A series of one or more
// `ConverseResponse` messages are streamed back to the client.
type ConverseResponse struct {
	// Exactly one of these fields will be populated in each `ConverseResponse`.
	//
	// Types that are valid to be assigned to ConverseResponse:
	//	*ConverseResponse_Error
	//	*ConverseResponse_EventType_
	//	*ConverseResponse_AudioOut
	//	*ConverseResponse_Result
	ConverseResponse isConverseResponse_ConverseResponse `protobuf_oneof:"converse_response"`
}

func (m *ConverseResponse) Reset()                    { *m = ConverseResponse{} }
func (m *ConverseResponse) String() string            { return proto.CompactTextString(m) }
func (*ConverseResponse) ProtoMessage()               {}
func (*ConverseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isConverseResponse_ConverseResponse interface {
	isConverseResponse_ConverseResponse()
}

type ConverseResponse_Error struct {
	Error *google_rpc.Status `protobuf:"bytes,1,opt,name=error,oneof"`
}
type ConverseResponse_EventType_ struct {
	EventType ConverseResponse_EventType `protobuf:"varint,2,opt,name=event_type,json=eventType,enum=google.assistant.embedded.v1alpha1.ConverseResponse_EventType,oneof"`
}
type ConverseResponse_AudioOut struct {
	AudioOut *AudioOut `protobuf:"bytes,3,opt,name=audio_out,json=audioOut,oneof"`
}
type ConverseResponse_Result struct {
	Result *ConverseResult `protobuf:"bytes,5,opt,name=result,oneof"`
}

func (*ConverseResponse_Error) isConverseResponse_ConverseResponse()      {}
func (*ConverseResponse_EventType_) isConverseResponse_ConverseResponse() {}
func (*ConverseResponse_AudioOut) isConverseResponse_ConverseResponse()   {}
func (*ConverseResponse_Result) isConverseResponse_ConverseResponse()     {}

func (m *ConverseResponse) GetConverseResponse() isConverseResponse_ConverseResponse {
	if m != nil {
		return m.ConverseResponse
	}
	return nil
}

func (m *ConverseResponse) GetError() *google_rpc.Status {
	if x, ok := m.GetConverseResponse().(*ConverseResponse_Error); ok {
		return x.Error
	}
	return nil
}

func (m *ConverseResponse) GetEventType() ConverseResponse_EventType {
	if x, ok := m.GetConverseResponse().(*ConverseResponse_EventType_); ok {
		return x.EventType
	}
	return ConverseResponse_EVENT_TYPE_UNSPECIFIED
}

func (m *ConverseResponse) GetAudioOut() *AudioOut {
	if x, ok := m.GetConverseResponse().(*ConverseResponse_AudioOut); ok {
		return x.AudioOut
	}
	return nil
}

func (m *ConverseResponse) GetResult() *ConverseResult {
	if x, ok := m.GetConverseResponse().(*ConverseResponse_Result); ok {
		return x.Result
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ConverseResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ConverseResponse_OneofMarshaler, _ConverseResponse_OneofUnmarshaler, _ConverseResponse_OneofSizer, []interface{}{
		(*ConverseResponse_Error)(nil),
		(*ConverseResponse_EventType_)(nil),
		(*ConverseResponse_AudioOut)(nil),
		(*ConverseResponse_Result)(nil),
	}
}

func _ConverseResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ConverseResponse)
	// converse_response
	switch x := m.ConverseResponse.(type) {
	case *ConverseResponse_Error:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case *ConverseResponse_EventType_:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.EventType))
	case *ConverseResponse_AudioOut:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AudioOut); err != nil {
			return err
		}
	case *ConverseResponse_Result:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Result); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ConverseResponse.ConverseResponse has unexpected type %T", x)
	}
	return nil
}

func _ConverseResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ConverseResponse)
	switch tag {
	case 1: // converse_response.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_rpc.Status)
		err := b.DecodeMessage(msg)
		m.ConverseResponse = &ConverseResponse_Error{msg}
		return true, err
	case 2: // converse_response.event_type
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ConverseResponse = &ConverseResponse_EventType_{ConverseResponse_EventType(x)}
		return true, err
	case 3: // converse_response.audio_out
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AudioOut)
		err := b.DecodeMessage(msg)
		m.ConverseResponse = &ConverseResponse_AudioOut{msg}
		return true, err
	case 5: // converse_response.result
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ConverseResult)
		err := b.DecodeMessage(msg)
		m.ConverseResponse = &ConverseResponse_Result{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ConverseResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ConverseResponse)
	// converse_response
	switch x := m.ConverseResponse.(type) {
	case *ConverseResponse_Error:
		s := proto.Size(x.Error)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ConverseResponse_EventType_:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.EventType))
	case *ConverseResponse_AudioOut:
		s := proto.Size(x.AudioOut)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ConverseResponse_Result:
		s := proto.Size(x.Result)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ConverseConfig)(nil), "google.assistant.embedded.v1alpha1.ConverseConfig")
	proto.RegisterType((*AudioInConfig)(nil), "google.assistant.embedded.v1alpha1.AudioInConfig")
	proto.RegisterType((*AudioOutConfig)(nil), "google.assistant.embedded.v1alpha1.AudioOutConfig")
	proto.RegisterType((*ConverseState)(nil), "google.assistant.embedded.v1alpha1.ConverseState")
	proto.RegisterType((*AudioOut)(nil), "google.assistant.embedded.v1alpha1.AudioOut")
	proto.RegisterType((*ConverseResult)(nil), "google.assistant.embedded.v1alpha1.ConverseResult")
	proto.RegisterType((*ConverseRequest)(nil), "google.assistant.embedded.v1alpha1.ConverseRequest")
	proto.RegisterType((*ConverseResponse)(nil), "google.assistant.embedded.v1alpha1.ConverseResponse")
	proto.RegisterEnum("google.assistant.embedded.v1alpha1.AudioInConfig_Encoding", AudioInConfig_Encoding_name, AudioInConfig_Encoding_value)
	proto.RegisterEnum("google.assistant.embedded.v1alpha1.AudioOutConfig_Encoding", AudioOutConfig_Encoding_name, AudioOutConfig_Encoding_value)
	proto.RegisterEnum("google.assistant.embedded.v1alpha1.ConverseResult_MicrophoneMode", ConverseResult_MicrophoneMode_name, ConverseResult_MicrophoneMode_value)
	proto.RegisterEnum("google.assistant.embedded.v1alpha1.ConverseResponse_EventType", ConverseResponse_EventType_name, ConverseResponse_EventType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EmbeddedAssistant service

type EmbeddedAssistantClient interface {
	// Initiates or continues a conversation with the embedded assistant service.
	// Each call performs one round-trip, sending an audio request to the service
	// and receiving the audio response. Uses bidirectional streaming to receive
	// results, such as the `END_OF_UTTERANCE` event, while sending audio.
	//
	// A conversation is one or more gRPC connections, each consisting of several
	// streamed requests and responses.
	// For example, the user says *Add to my shopping list* and the assistant
	// responds *What do you want to add?*. The sequence of streamed requests and
	// responses in the first gRPC message could be:
	//
	// *   ConverseRequest.config
	// *   ConverseRequest.audio_in
	// *   ConverseRequest.audio_in
	// *   ConverseRequest.audio_in
	// *   ConverseRequest.audio_in
	// *   ConverseResponse.event_type.END_OF_UTTERANCE
	// *   ConverseResponse.result.microphone_mode.DIALOG_FOLLOW_ON
	// *   ConverseResponse.audio_out
	// *   ConverseResponse.audio_out
	// *   ConverseResponse.audio_out
	//
	// The user then says *bagels* and the assistant responds
	// *OK, I've added bagels to your shopping list*. This is sent as another gRPC
	// connection call to the `Converse` method, again with streamed requests and
	// responses, such as:
	//
	// *   ConverseRequest.config
	// *   ConverseRequest.audio_in
	// *   ConverseRequest.audio_in
	// *   ConverseRequest.audio_in
	// *   ConverseResponse.event_type.END_OF_UTTERANCE
	// *   ConverseResponse.result.microphone_mode.CLOSE_MICROPHONE
	// *   ConverseResponse.audio_out
	// *   ConverseResponse.audio_out
	// *   ConverseResponse.audio_out
	// *   ConverseResponse.audio_out
	//
	// Although the precise order of responses is not guaranteed, sequential
	// ConverseResponse.audio_out messages will always contain sequential portions
	// of audio.
	Converse(ctx context.Context, opts ...grpc.CallOption) (EmbeddedAssistant_ConverseClient, error)
}

type embeddedAssistantClient struct {
	cc *grpc.ClientConn
}

func NewEmbeddedAssistantClient(cc *grpc.ClientConn) EmbeddedAssistantClient {
	return &embeddedAssistantClient{cc}
}

func (c *embeddedAssistantClient) Converse(ctx context.Context, opts ...grpc.CallOption) (EmbeddedAssistant_ConverseClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EmbeddedAssistant_serviceDesc.Streams[0], c.cc, "/google.assistant.embedded.v1alpha1.EmbeddedAssistant/Converse", opts...)
	if err != nil {
		return nil, err
	}
	x := &embeddedAssistantConverseClient{stream}
	return x, nil
}

type EmbeddedAssistant_ConverseClient interface {
	Send(*ConverseRequest) error
	Recv() (*ConverseResponse, error)
	grpc.ClientStream
}

type embeddedAssistantConverseClient struct {
	grpc.ClientStream
}

func (x *embeddedAssistantConverseClient) Send(m *ConverseRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *embeddedAssistantConverseClient) Recv() (*ConverseResponse, error) {
	m := new(ConverseResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for EmbeddedAssistant service

type EmbeddedAssistantServer interface {
	// Initiates or continues a conversation with the embedded assistant service.
	// Each call performs one round-trip, sending an audio request to the service
	// and receiving the audio response. Uses bidirectional streaming to receive
	// results, such as the `END_OF_UTTERANCE` event, while sending audio.
	//
	// A conversation is one or more gRPC connections, each consisting of several
	// streamed requests and responses.
	// For example, the user says *Add to my shopping list* and the assistant
	// responds *What do you want to add?*. The sequence of streamed requests and
	// responses in the first gRPC message could be:
	//
	// *   ConverseRequest.config
	// *   ConverseRequest.audio_in
	// *   ConverseRequest.audio_in
	// *   ConverseRequest.audio_in
	// *   ConverseRequest.audio_in
	// *   ConverseResponse.event_type.END_OF_UTTERANCE
	// *   ConverseResponse.result.microphone_mode.DIALOG_FOLLOW_ON
	// *   ConverseResponse.audio_out
	// *   ConverseResponse.audio_out
	// *   ConverseResponse.audio_out
	//
	// The user then says *bagels* and the assistant responds
	// *OK, I've added bagels to your shopping list*. This is sent as another gRPC
	// connection call to the `Converse` method, again with streamed requests and
	// responses, such as:
	//
	// *   ConverseRequest.config
	// *   ConverseRequest.audio_in
	// *   ConverseRequest.audio_in
	// *   ConverseRequest.audio_in
	// *   ConverseResponse.event_type.END_OF_UTTERANCE
	// *   ConverseResponse.result.microphone_mode.CLOSE_MICROPHONE
	// *   ConverseResponse.audio_out
	// *   ConverseResponse.audio_out
	// *   ConverseResponse.audio_out
	// *   ConverseResponse.audio_out
	//
	// Although the precise order of responses is not guaranteed, sequential
	// ConverseResponse.audio_out messages will always contain sequential portions
	// of audio.
	Converse(EmbeddedAssistant_ConverseServer) error
}

func RegisterEmbeddedAssistantServer(s *grpc.Server, srv EmbeddedAssistantServer) {
	s.RegisterService(&_EmbeddedAssistant_serviceDesc, srv)
}

func _EmbeddedAssistant_Converse_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EmbeddedAssistantServer).Converse(&embeddedAssistantConverseServer{stream})
}

type EmbeddedAssistant_ConverseServer interface {
	Send(*ConverseResponse) error
	Recv() (*ConverseRequest, error)
	grpc.ServerStream
}

type embeddedAssistantConverseServer struct {
	grpc.ServerStream
}

func (x *embeddedAssistantConverseServer) Send(m *ConverseResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *embeddedAssistantConverseServer) Recv() (*ConverseRequest, error) {
	m := new(ConverseRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EmbeddedAssistant_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.assistant.embedded.v1alpha1.EmbeddedAssistant",
	HandlerType: (*EmbeddedAssistantServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Converse",
			Handler:       _EmbeddedAssistant_Converse_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "google/assistant/embedded/v1alpha1/embedded_assistant.proto",
}

func init() {
	proto.RegisterFile("google/assistant/embedded/v1alpha1/embedded_assistant.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 892 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x51, 0x73, 0xdb, 0x44,
	0x10, 0xb6, 0xec, 0xa6, 0xb5, 0xb7, 0x89, 0x2c, 0x5f, 0x33, 0x90, 0x49, 0x61, 0x60, 0xf4, 0xc0,
	0x94, 0x02, 0x72, 0xe3, 0x30, 0x3c, 0x10, 0xe8, 0x8c, 0x63, 0x2b, 0xb1, 0xc1, 0x96, 0x3c, 0x67,
	0xa7, 0xa5, 0x0c, 0xcc, 0xcd, 0x55, 0x3e, 0x1c, 0x81, 0x7d, 0x27, 0xa4, 0x73, 0xa6, 0xe1, 0x07,
	0xf0, 0xd8, 0xe1, 0x95, 0x67, 0x7e, 0x11, 0xff, 0x88, 0xd1, 0x9d, 0xa4, 0xd8, 0x90, 0x42, 0x1c,
	0x1e, 0x6f, 0xf7, 0xbe, 0x4f, 0xbb, 0xdf, 0x7e, 0xb7, 0x23, 0x38, 0x9a, 0x09, 0x31, 0x9b, 0xb3,
	0x26, 0x4d, 0x92, 0x30, 0x91, 0x94, 0xcb, 0x26, 0x5b, 0xbc, 0x64, 0xd3, 0x29, 0x9b, 0x36, 0x2f,
	0x0e, 0xe8, 0x3c, 0x3a, 0xa7, 0x07, 0x45, 0x84, 0x14, 0x97, 0x9c, 0x28, 0x16, 0x52, 0x20, 0x5b,
	0x83, 0x9d, 0xab, 0x78, 0x7e, 0xd5, 0xc9, 0xc1, 0xfb, 0xef, 0xe4, 0x1f, 0x88, 0xc2, 0x26, 0xe5,
	0x5c, 0x48, 0x2a, 0x43, 0xc1, 0x13, 0xcd, 0xb0, 0xff, 0x76, 0x96, 0x8d, 0xa3, 0xa0, 0x99, 0x48,
	0x2a, 0x97, 0x59, 0xc2, 0xfe, 0xa3, 0x0c, 0x66, 0x47, 0xf0, 0x0b, 0x16, 0x27, 0xac, 0x23, 0xf8,
	0x0f, 0xe1, 0x0c, 0xbd, 0x80, 0x3a, 0x5d, 0x4e, 0x43, 0x41, 0x42, 0x4e, 0x02, 0x15, 0xda, 0x33,
	0xde, 0x37, 0x1e, 0xdd, 0x6f, 0x1d, 0x38, 0xff, 0x5d, 0x87, 0xd3, 0x4e, 0xa1, 0x7d, 0xae, 0xb9,
	0xf0, 0x0e, 0x5d, 0x3d, 0xa2, 0xef, 0xc0, 0xd2, 0xd4, 0x62, 0x29, 0x73, 0xee, 0xb2, 0xe2, 0x6e,
	0xdd, 0x98, 0xdb, 0x5f, 0xca, 0x8c, 0xdc, 0xa4, 0x6b, 0x67, 0xf4, 0x0d, 0x98, 0x41, 0xd6, 0x0a,
	0x49, 0x9b, 0x64, 0x7b, 0x95, 0x9b, 0xd7, 0x9d, 0x8b, 0x30, 0x4e, 0x81, 0x78, 0x27, 0x58, 0x3d,
	0xda, 0x7f, 0x1a, 0xb0, 0xb3, 0xd6, 0x18, 0x7a, 0x06, 0x55, 0xc6, 0x03, 0x31, 0x0d, 0xb9, 0x56,
	0xc7, 0x6c, 0x7d, 0xbe, 0xb1, 0x3a, 0x8e, 0x9b, 0x31, 0xe0, 0x82, 0x0b, 0x3d, 0x86, 0x46, 0x42,
	0x17, 0xd1, 0x9c, 0x91, 0x98, 0x4a, 0x46, 0xce, 0x59, 0x2c, 0x7f, 0x51, 0x12, 0x6d, 0xe1, 0xba,
	0x4e, 0x60, 0x2a, 0x59, 0x2f, 0x0d, 0xdb, 0x5f, 0x40, 0x35, 0x67, 0x40, 0x7b, 0xb0, 0xeb, 0x7a,
	0x1d, 0xbf, 0xdb, 0xf7, 0x4e, 0xc9, 0x99, 0x37, 0x1e, 0xb9, 0x9d, 0xfe, 0x49, 0xdf, 0xed, 0x5a,
	0x25, 0xb4, 0x0d, 0xd5, 0x41, 0xdf, 0x73, 0xdb, 0xf8, 0xe0, 0x33, 0xcb, 0x40, 0x55, 0xb8, 0x73,
	0x32, 0x68, 0x77, 0xac, 0xb2, 0xfd, 0x5b, 0x19, 0xcc, 0x75, 0x41, 0xd1, 0xf3, 0x7f, 0x34, 0x75,
	0xb4, 0xf9, 0x58, 0xfe, 0x67, 0x57, 0xe8, 0x23, 0x68, 0x5c, 0x88, 0xf9, 0x72, 0xc1, 0x48, 0xc4,
	0xe2, 0x80, 0x71, 0x49, 0x67, 0x7a, 0x90, 0x5b, 0xd8, 0xd2, 0x89, 0x51, 0x11, 0xb7, 0x07, 0xb7,
	0x90, 0xe0, 0x1e, 0x54, 0x86, 0xa3, 0x43, 0xab, 0x8c, 0xea, 0x70, 0xdf, 0x1f, 0x9d, 0x8d, 0x49,
	0xdf, 0x23, 0xfe, 0xe9, 0xa9, 0x55, 0xb1, 0x9f, 0xc2, 0xce, 0x9a, 0x0d, 0xd0, 0x27, 0x80, 0x32,
	0x23, 0xa8, 0xd7, 0x94, 0xb9, 0x2a, 0x95, 0x66, 0x1b, 0x37, 0x56, 0x33, 0xda, 0x26, 0x1f, 0x42,
	0x35, 0xd7, 0x02, 0xbd, 0x0b, 0xa0, 0xad, 0x3e, 0xa5, 0x92, 0x66, 0x90, 0x9a, 0x8a, 0x74, 0xa9,
	0xa4, 0xf6, 0xef, 0x95, 0xab, 0x77, 0x87, 0x59, 0xb2, 0x9c, 0x4b, 0xe4, 0xc0, 0x83, 0x24, 0x12,
	0x3f, 0x31, 0x4e, 0x62, 0xf6, 0xf3, 0x92, 0x25, 0x92, 0x48, 0xf6, 0x4a, 0x2a, 0x68, 0x0d, 0x37,
	0x74, 0x0a, 0xeb, 0xcc, 0x84, 0xbd, 0x92, 0xe8, 0x09, 0xec, 0x16, 0xf7, 0x93, 0x48, 0xf0, 0x84,
	0x69, 0x40, 0x59, 0x01, 0x50, 0x0e, 0xd0, 0x29, 0x85, 0xb8, 0xbe, 0x9d, 0xca, 0x1b, 0xda, 0x41,
	0x3f, 0x42, 0x7d, 0x11, 0x06, 0xb1, 0x88, 0xce, 0x05, 0x67, 0x64, 0x21, 0xa6, 0x6c, 0xef, 0x8e,
	0x72, 0x45, 0x7b, 0x93, 0x07, 0xa5, 0xbb, 0x73, 0x86, 0x05, 0xd3, 0x50, 0x4c, 0x19, 0x36, 0x17,
	0x6b, 0xe7, 0xeb, 0xa7, 0xbe, 0xf5, 0x86, 0xa9, 0x7f, 0x0f, 0xe6, 0x3a, 0x1d, 0x7a, 0x0f, 0x1e,
	0x0e, 0xfb, 0x1d, 0xec, 0x8f, 0x7a, 0xbe, 0xe7, 0x92, 0xa1, 0xdf, 0x75, 0xff, 0x66, 0x81, 0x5d,
	0xb0, 0x3a, 0x03, 0x7f, 0xec, 0x92, 0xab, 0x6b, 0x96, 0x91, 0x46, 0xbb, 0xfd, 0xf6, 0xc0, 0x3f,
	0x25, 0x27, 0xfe, 0x60, 0xe0, 0x3f, 0x27, 0xbe, 0x97, 0xbe, 0x0c, 0x03, 0xea, 0x57, 0xd5, 0x2b,
	0xc1, 0xd1, 0x00, 0xee, 0xae, 0xed, 0xc2, 0xd6, 0x26, 0x12, 0xe8, 0x87, 0xd1, 0x2b, 0xe1, 0x8c,
	0x03, 0x3d, 0x84, 0x6a, 0xbe, 0x62, 0xd5, 0xb8, 0xb6, 0x7b, 0x25, 0x7c, 0x2f, 0x5b, 0x95, 0xc7,
	0x08, 0xac, 0x62, 0x8d, 0x65, 0x4e, 0xb0, 0x5f, 0x57, 0xc0, 0x5a, 0x11, 0x54, 0x8d, 0x14, 0x3d,
	0x86, 0x2d, 0x16, 0xc7, 0x22, 0xce, 0x4a, 0x42, 0x79, 0x49, 0x71, 0x14, 0x38, 0x63, 0xb5, 0xe4,
	0x7b, 0x25, 0xac, 0xaf, 0x20, 0x02, 0xc0, 0x2e, 0x18, 0x97, 0x44, 0x5e, 0x46, 0x4c, 0x7d, 0xd3,
	0x6c, 0x3d, 0xdd, 0x70, 0x8c, 0xea, 0xab, 0x8e, 0x9b, 0xd2, 0x4c, 0x2e, 0x23, 0xd6, 0x2b, 0xe1,
	0x1a, 0xcb, 0x0f, 0xe8, 0x6b, 0xa8, 0x15, 0xab, 0x3d, 0xdb, 0xbb, 0x1f, 0x6f, 0xb2, 0x3c, 0x7a,
	0x25, 0x5c, 0xcd, 0xf7, 0x79, 0xaa, 0x76, 0xac, 0x6c, 0xa3, 0x2c, 0xb0, 0xa1, 0xda, 0xda, 0x70,
	0xa9, 0xda, 0x9a, 0xc3, 0xfe, 0x12, 0x6a, 0x45, 0xd1, 0x68, 0x1f, 0xde, 0x72, 0x9f, 0xb9, 0xde,
	0x84, 0x4c, 0x5e, 0x8c, 0xae, 0x31, 0x89, 0xeb, 0x75, 0x89, 0x7f, 0x42, 0xce, 0x26, 0x13, 0x17,
	0xb7, 0xbd, 0x8e, 0x6b, 0x19, 0xc7, 0x0f, 0xa0, 0xb1, 0x32, 0x0f, 0xad, 0x42, 0xeb, 0xb5, 0x01,
	0x0d, 0x37, 0x2b, 0xa1, 0x9d, 0x17, 0x85, 0x2e, 0xa1, 0x9a, 0x57, 0x81, 0x0e, 0x37, 0xab, 0x59,
	0xcd, 0x79, 0xff, 0xd3, 0xdb, 0x8c, 0xe4, 0x91, 0xf1, 0xc4, 0x38, 0xfe, 0xd5, 0x80, 0x0f, 0x02,
	0xb1, 0xb8, 0x01, 0xfe, 0xd8, 0x2c, 0x0a, 0x1e, 0xa5, 0xff, 0x00, 0x23, 0xe3, 0xdb, 0xaf, 0x32,
	0xd4, 0x4c, 0xcc, 0x29, 0x9f, 0x39, 0x22, 0x9e, 0x35, 0x67, 0x8c, 0xab, 0x3f, 0x84, 0xa6, 0x4e,
	0xd1, 0x28, 0x4c, 0xfe, 0xed, 0xe7, 0xe5, 0x28, 0x8f, 0xbc, 0xbc, 0xab, 0x60, 0x87, 0x7f, 0x05,
	0x00, 0x00, 0xff, 0xff, 0xec, 0x7a, 0x68, 0xfa, 0xf2, 0x08, 0x00, 0x00,
}
