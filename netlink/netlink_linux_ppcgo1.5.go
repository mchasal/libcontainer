// +build ppc64 ppc64le
// +build go1.5

package netlink

func ifrDataByte(b byte) uint8 {
	return uint8(b)
}
