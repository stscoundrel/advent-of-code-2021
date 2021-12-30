package transmissions

import (
	"strconv"
	"strings"
)

func joinBits(bitRange []string) string {
	return strings.Join(bitRange, "")
}

func binaryToDecimal(binary string) int {
	number, _ := strconv.ParseInt(binary, 2, 64)
	return int(number)
}

func getVersion(transmission []string, bitsRead *int) int {
	version := joinBits(transmission[*bitsRead : *bitsRead+3])
	*bitsRead += 3
	return binaryToDecimal(version)
}

func getTypeId(transmission []string, bitsRead *int) int {
	typeId := joinBits(transmission[*bitsRead : *bitsRead+3])
	*bitsRead += 3
	return binaryToDecimal(typeId)
}

func parseSubPacketLength(transmission []string, bitsRead *int) (int, int) {
	subpacket := ""
	lengthType, _ := strconv.Atoi(transmission[*bitsRead])
	*bitsRead += 1

	if lengthType == 0 {
		subpacket = joinBits(transmission[*bitsRead : *bitsRead+15])
		*bitsRead += 15
	} else {
		subpacket = joinBits(transmission[*bitsRead : *bitsRead+11])
		*bitsRead += 11
	}

	return lengthType, binaryToDecimal(subpacket)
}

func parseLiteralValuePacket(packet Packet, transmission []string, bitsRead *int) Packet {
	isProcessing := true
	hexNumber := ""

	for isProcessing {
		fiveBits := transmission[*bitsRead : *bitsRead+5]
		hexNumber += joinBits(fiveBits[1:5])
		if fiveBits[0] == "0" {
			isProcessing = false
		}

		*bitsRead += 5
	}

	packet.value = binaryToDecimal(hexNumber)

	return packet
}

func parseOperatorPacket(packet Packet, transmission []string, bitsRead *int) Packet {
	packets := []Packet{}
	lengthTypeId, lengthValue := parseSubPacketLength(transmission, bitsRead)

	if lengthTypeId == 0 {
		packetEndBit := *bitsRead + lengthValue

		for *bitsRead < packetEndBit {
			packets = append(packets, parsePacket(transmission, bitsRead))
		}
	} else {
		for i := 0; i < lengthValue; i += 1 {
			packets = append(packets, parsePacket(transmission, bitsRead))
		}
	}

	packet.subPackets = packets

	return packet
}

func parsePacket(transmission []string, bitsRead *int) Packet {
	packet := Packet{
		version: getVersion(transmission, bitsRead),
		typeId:  getTypeId(transmission, bitsRead),
	}

	if packet.isLiteralPackage() {
		packet = parseLiteralValuePacket(packet, transmission, bitsRead)
	} else {
		packet = parseOperatorPacket(packet, transmission, bitsRead)
	}

	return packet
}
