package transmissions

func SumVersionNumbers(hexTransmissions string) int {
	binaryTransmissions := toBinaryTransmission(hexTransmissions)
	bitsRead := 0
	packet := parsePacket(binaryTransmissions, &bitsRead)

	return packet.sumPacketVersions()
}
