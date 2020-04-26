package char

const NumlowerToUpper = uint8('a' - 'A')

func IsUpper(char byte) bool {
	return 'A' <= char && char <= 'Z'
}

func ToLower(char byte) byte {
	if IsUpper(char) {
		return char + NumlowerToUpper
	}

	return char
}
