//go:generate protoc -I=../../schemas --go_out=.. ../../schemas/message.proto

package protobuf

import (
	"log"

	"github.com/kcchu/buffer-benchmarks/constants"
	"github.com/kcchu/buffer-benchmarks/puffer/generated"
)

func Encode() []byte {
	md := &generated.MessageData{
		Body: &generated.CastAddBody{
			Parent: &generated.CastId{
				Fid:     constants.SampleFid,
				Ts_hash: constants.SampleTsHash,
			},
			Text: constants.SampleText,
		},
		Type:      1,
		Timestamp: constants.SampleTimestamp,
		Fid:       constants.SampleFid,
		Network:   3,
	}
	mdBytes := md.Pack()

	m := &generated.Message{
		Data:             mdBytes,
		Hash:             constants.SampleHash,
		Hash_scheme:      1,
		Signature:        constants.SampleSignature,
		Signature_scheme: 1,
		Signer:           constants.SampleSigner,
	}
	mBytes := m.Pack()
	return mBytes
}

func Decode(buf []byte) {
	m := new(generated.Message)
	if err := m.Unpack(buf); err != nil {
		log.Fatalln("Failed to decode Message", err)
	}
	md := new(generated.MessageData)
	if err := md.Unpack(m.Data); err != nil {
		log.Fatalln("Failed to decode MessageData", err)
	}
	castAddBody := md.Body
	if castAddBody.Text != constants.SampleText {
		log.Fatalln("Unexpected decoded text")
	}
}
