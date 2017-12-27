# regexRange
This package will generate regular expressions for [a...b] ranges. Please see the test-cases for a howto.

	regex := regexRange.GetRegex(196608, 197631)
	positiveMatch, err := regexp.Compile(regex)
	if err != nil {
		t.Error("Cant compile regex")
	}

	if !positiveMatch.MatchString("_196922") {
		t.Error("Regex did not match _196922!")
	} 
