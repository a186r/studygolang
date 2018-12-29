package tempconv

// 包级别的常量名都是以大写字母开头，
type Celsius float64
type Fahrenheit float64

const{
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
}

func (c Celsius) String() string {
	return fmt.Sprintf("%gC",c)
}

func (f Fahrenheit) String() string{
	return fmt.Sprintf("%gF",f)
}