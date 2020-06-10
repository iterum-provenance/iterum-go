package transmit

import (
	"encoding/binary"
	"net"
)

// EncodeSend encodes a serializable object via the Iterum defaults:
//     unsigned 32bit int msg length, followed by the encoded object
// If the message size exceeds 2^16 it is send in chunks of 2^16 bytes
// Then it sends it on the target connection
func EncodeSend(conn net.Conn, obj Serializable) (err error) {
	// Encoding
	data, err := obj.Serialize()
	if err != nil {
		return
	}

	size := make([]byte, FragmentSizeLength)
	binary.LittleEndian.PutUint32(size, uint32(len(data)))
	data = append(size, data...)
	dataLen := len(data)

	// Setup sending in chunks
	bytesWrittenTotal := 0
	lwb := 0
	maxChunkSize := 65536 // 2^16
	upb := dataLen
	if dataLen > maxChunkSize {
		upb = maxChunkSize
	}

	// Send all chunks
	for bytesWrittenTotal != dataLen {
		bytesWritten, err := conn.Write(data[lwb:upb])
		if err != nil {
			return ErrConnection(err)
		}
		lwb = upb
		if maxChunkSize+upb > dataLen {
			upb = dataLen
		} else {
			upb += maxChunkSize
		}
		bytesWrittenTotal += bytesWritten
	}

	return
}

// ReadMessage reads one message worth of data from the passed connection.
// it leaves the data serialized, but it is guaranteed to be of the right length
func ReadMessage(conn net.Conn) (encMsg []byte, err error) {
	encMsgSize := make([]byte, FragmentSizeLength)
	_, err = conn.Read(encMsgSize)
	if err != nil {
		err = ErrConnection(err)
		return
	}
	msgSize := int(binary.LittleEndian.Uint32(encMsgSize))

	encMsg = make([]byte, msgSize)
	_, err = conn.Read(encMsg)

	if err != nil {
		err = ErrConnection(err)
		return
	}
	return encMsg, nil
}

// DecodeRead tries to decode a serialized object that was encoded
// via the Iterum defaults as described in `transmit.Encode`
// and Read from the passed connection
func DecodeRead(conn net.Conn, obj Serializable) (err error) {
	// Reading
	encMsg, err := ReadMessage(conn)
	// Decoding
	err = obj.Deserialize(encMsg)
	return
}
