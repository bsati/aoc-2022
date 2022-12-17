package solutions

type Problem interface {
	Parse(input string)
	Part1()
	Part2()
}

func GetSolver(day int) Problem {
	switch day {
	case 1:
		return newDay1()
	case 2:
		return newDay2()
	case 9:
		return newDay9()
	case 14:
		return newDay14()
	}
	return nil
}
