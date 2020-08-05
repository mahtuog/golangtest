package iteration

const repeatCount = 5

func Repeat(character string, repeatTimes int) (repeated string){

	if repeatTimes < 0{
		repeatTimes = repeatCount
	}

	for i := 0; i< repeatTimes; i++{
		repeated += character
	}
	return
}
