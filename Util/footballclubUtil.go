package Util

//all utility functions are related to footballclub

/*Need --> Generate Random club_name*/
func RandomfootballclubName() string {
	return RandomStringGenerator(10)
}

/*Need --> Generate Random country_fc*/
func Randomcountryfc() string {
	return RandomStringGenerator(4)
}

/* Need --> Generate Random balance */
func Randombalance() int64 {
	return RandomIntGenerator(10000000, 100000000000)
}
