package scripts

func OpenURL(url string) string {
	return `
	var error = ''
	try {
		location.href = '` + url + `'
	} catch(err) {
		error + err
	}
	`
}

func GetBool(jsBool string) string {
	return `
		try {
			` + jsBool + `
		} catch(err) {}
	`
}

func GetStringsSlice(jsString string) string {
	return `
		var result = [];
		try {
			` + jsString + `
		} catch(err) {
  			result.push(err)
			result
		}
	`
}
func GetString(jsString string) string {
	return `
		var error = ''
		try {
			` + jsString + `
		} catch(err) {
			error + err
		}
	`
}
