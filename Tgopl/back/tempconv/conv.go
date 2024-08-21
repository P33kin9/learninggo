package tempconv

// CToF converts a Celsius temprature to Fathrentheit.

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// Add Kelvin function
func CToK(c Celsius) Kelvin { return Kelvin(c + AbsoluteZeroC) }
func KToC(k Kelvin) Celsius { return Celsius(k) - AbsoluteZeroC }
