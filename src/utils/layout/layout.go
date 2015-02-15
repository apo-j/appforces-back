package layout

const (
	Top int = 1
	Left int = 2
	Center int = 4
	Right int = 8
	Bottom int = 16
)

func HasTop(layout int)  bool{
	return layout & Top > 0
}

func HasLeft(layout int)  bool{
	return layout & Left > 0
}

func HasRight(layout int)  bool{
	return layout & Right > 0
}

func HasBottom(layout int)  bool{
	return layout & Bottom > 0
}

func HasCenter(layout int)  bool{
	return layout & Center > 0
}
