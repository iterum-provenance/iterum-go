// Package transmit contains the transmission protocol between Iterum sidecars and user-defined containers.
// It encodes and decodes any Go struct that implements the Serializable interface.
//
// The protocol is fairly simple. It first Serializes the structure into a byte slice. This slice is
// prepended with 4 byte unsigned int specifying the total length of the message. If the message is larger
// than 2^16, then it is send in chunks of 2^16.
package transmit
