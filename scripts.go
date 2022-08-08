package chromedp_wrapper

func scriptOpenURL(url string) string {
	return `
	var error = ''
	try {
		location.href = '` + url + `'
	} catch(err) {
		error + err
	}
	`
}

func scriptGetBool(jsBool string) string {
	return `
		try {
			` + jsBool + `
		} catch(err) {}
	`
}

func scriptGetStringsSlice(jsString string) string {
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
func scriptGetString(jsString string) string {
	return `
		var error = ''
		try {
			` + jsString + `
		} catch(err) {
			error + err
		}
	`
}
