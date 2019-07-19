package sunspec

import "fmt"

func (p Eui48) String() string {
	buf := []byte{}
	for x, b := range p[2:] {
		if x != 0 {
			buf = append(buf, ':')
		}
		buf = append(buf, fmt.Sprintf("%02x", b)...)
	}
	return string(buf)
}

func (p Ipaddr) String() string {
	buf := []byte{}
	for x, b := range p {
		if x != 0 {
			buf = append(buf, '.')
		}
		buf = append(buf, fmt.Sprintf("%d", b)...)
	}
	return string(buf)
}

func (p Ipv6addr) String() string {
	buf := []byte{}
	for x := range p {
		if x%2 == 1 {
			continue
		}
		if x != 0 {
			buf = append(buf, ':')
		}
		buf = append(buf, fmt.Sprintf("%04x", uint16(p[x])<<8|uint16(p[x+1]))...)
	}
	return string(buf)
}

func (p Bitfield16) String() string {
	return fmt.Sprintf("0x%04x", uint16(p))
}

func (p Bitfield32) String() string {
	return fmt.Sprintf("0x%08x", uint32(p))
}

func (p Pad) String() string {
	return fmt.Sprintf("0x%04x", uint16(p))
}
