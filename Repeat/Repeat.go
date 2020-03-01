package Repeat

//Repeat charactor for {repeatTimes} times
func Repeat(charactor string, repeatTimes int) string {
	var repeated string
	for i := 0; i < repeatTimes; i++ {
		repeated += charactor
	}
	return repeated
}
