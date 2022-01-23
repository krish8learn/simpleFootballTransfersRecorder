package Util

//this will generate player name
func RandomPlayername() string {
	return RandomStringGenerator(8) + " " + RandomStringGenerator(8)
}

//this will generate player position
func RandomPosition() string {
	positions := []string{"GK", "DEF", "MID", "FW", "WING", "DM", "SB", "CB", "AM", "CM"}
	max := len(positions)
	return positions[RandomIntGenerator(0, int64(max)-1)]
}

//this will generate player country
func RandomCountryPl() string {
	return RandomStringGenerator(4)
}

//this will generate player value
func RandomPlayerValue() int64 {
	return RandomIntGenerator(10000, 10000000)
}
