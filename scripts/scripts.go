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

func GetValue(js string) string {
	return `
		try {
			` + js + `
		} catch(err) {}
	`
}
