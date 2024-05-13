package validators

const (
	IntegerRegex = `^[\d]*$`
	StringRegex  = `^[\s\S]*$`
	EmailRegex   = `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`
	PhoneRegex   = `^[\d]{10,15}$`
	FloatRegex   = `^[-]?[\d]*\.[\d]*$`
)

var tagRegexMap = map[string]string{
	IntegerTag: IntegerRegex,
	StringTag:  StringRegex,
	EmailTag:   EmailRegex,
	PhoneTag:   PhoneRegex,
}

func RegexForTagGet(tag string) string {
	regex, exists := tagRegexMap[tag]
	if !exists {
		panic("regex not found for tag " + tag)
	}
	return regex
}
