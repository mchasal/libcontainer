// +build ppc64 ppc64le
// +build !go1.5

package netlink

func ifrDataByte(b byte) int8 {
    return int8(b)
}
