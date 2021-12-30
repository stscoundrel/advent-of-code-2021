package transmissions

const LITERAL_PACKAGE = 4

type Packet struct {
	version    int
	typeId     int
	value      int
	subPackets []Packet
}

func (packet *Packet) sumPacketVersions() int {
	sum := packet.version

	for _, subpacket := range packet.subPackets {
		sum += subpacket.sumPacketVersions()
	}

	return sum
}
