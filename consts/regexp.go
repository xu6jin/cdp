package consts

const (
	RegexpPatternIPv4 = `\b(?:(?:2(?:[0-4][0-9]|5[0-5])|[0-1]?[0-9]?[0-9])\.){3}(?:(?:2([0-4][0-9]|5[0-5])|[0-1]?[0-9]?[0-9]))\b`

	RegexpPatternIPv6 = `(([a-fA-F0-9]{1,4}|):){1,7}([a-fA-F0-9]{1,4}|:)`

	RegexpPatternIP = `(` + RegexpPatternIPv4 + `)|(` + RegexpPatternIPv6 + `)`

	RegexpPatternDomain = `([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}`

	RegexpPatternIPorDomain = `(` + RegexpPatternIP + `)|(` + RegexpPatternDomain + `)`

	RegexpPatternURL = `([--:\w?@%&+~#=]*\.[a-z]{2,4}\/{0,2})((?:[?&](?:\w+)=(?:\w+))+|[--:\w?@%&+~#=]+)?`
)
