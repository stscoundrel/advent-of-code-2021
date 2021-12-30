package transmissions

func GetValue(hexTransmissions string) int {
	binaryTransmissions := toBinaryTransmission(hexTransmissions)
	bitsRead := 0
	packet := parsePacket(binaryTransmissions, &bitsRead)

	return packet.getValue()
}
