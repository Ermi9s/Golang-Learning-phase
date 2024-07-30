package util

func Course_Score(score float32) bool{
	return score > -1 && score < 101;
}

func Numbers(num int) bool {
	return num >= 0;
}

func Names(name string) bool {
	for i := range len(name) {
		if int(name[i]) - int('a') > 26 || int(name[i]) - int('a') < 0 {
			return false;
		} 
	}
	return true;
}