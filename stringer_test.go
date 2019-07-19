package sunspec

import (
	"testing"
)

func TestEui48(t *testing.T) {
	p := Eui48{1, 2, 3, 4, 5, 6, 7, 8}

	if v := p.String(); v != "03:04:05:06:07:08" {
		t.Errorf("unexpected result, got %v", v)
	}
}

func TestIpaddr(t *testing.T) {
	p := Ipaddr{1, 2, 3, 4}

	if v := p.String(); v != "1.2.3.4" {
		t.Errorf("unexpected result, got %v", v)
	}
}

func TestIpv6addr(t *testing.T) {
	p := Ipv6addr{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8}

	if v := p.String(); v != "0102:0304:0506:0708:0102:0304:0506:0708" {
		t.Errorf("unexpected result, got %v", v)
	}
}

func TestBitfield16(t *testing.T) {
	p := Bitfield16(0x1234)

	if v := p.String(); v != "0x1234" {
		t.Errorf("unexpected result, got %v", v)
	}
}

func TestBitfield32(t *testing.T) {
	p := Bitfield32(0x12345678)

	if v := p.String(); v != "0x12345678" {
		t.Errorf("unexpected result, got %v", v)
	}
}

func TestPad(t *testing.T) {
	p := Pad(0x1234)

	if v := p.String(); v != "0x1234" {
		t.Errorf("unexpected result, got %v", v)
	}}
