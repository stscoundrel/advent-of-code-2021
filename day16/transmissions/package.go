package transmissions

const LITERAL_PACKAGE = 4

type Packet struct {
	version    int
	typeId     int
	value      int
	subPackets []Packet
}

func (packet *Packet) isLiteralPackage() bool {
	return packet.typeId == LITERAL_PACKAGE
}

func (packet *Packet) sumPacketVersions() int {
	sum := packet.version

	for _, subpacket := range packet.subPackets {
		sum += subpacket.sumPacketVersions()
	}

	return sum
}

func (packet *Packet) getValue() int {
	switch packet.typeId {
	case 0:
		value := 0
		for _, subPacket := range packet.subPackets {
			value += subPacket.getValue()
		}
		return value
	case 1:
		value := 1
		for _, subPacket := range packet.subPackets {
			value *= subPacket.getValue()
		}
		return value
	case 2:
		value := -1
		for _, subPacket := range packet.subPackets {
			subValue := subPacket.getValue()
			if value == -1 || subValue < value {
				value = subValue
			}
		}
		return value
	case 3:
		value := -1
		for _, subPacket := range packet.subPackets {
			subValue := subPacket.getValue()
			if value == -1 || subValue > value {
				value = subValue
			}
		}
		return value
	case 5:
		if packet.subPackets[0].getValue() > packet.subPackets[1].getValue() {
			return 1
		}
		return 0
	case 6:
		if packet.subPackets[0].getValue() < packet.subPackets[1].getValue() {
			return 1
		}
		return 0
	case 7:
		if packet.subPackets[0].getValue() == packet.subPackets[1].getValue() {
			return 1
		}
		return 0
	default:
		return packet.value
	}
}
