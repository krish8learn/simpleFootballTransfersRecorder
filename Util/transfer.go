package Util

//this will generate season transfer
func RandomSeason() int64 {
	return RandomIntGenerator(2000, 2022)
}

//this will generate transfer amount
func RandomAmount() int64 {
	return RandomIntGenerator(10000, 10000000)
}
