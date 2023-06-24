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

func ScrollDown() {
	return `
	// Get the current scroll position
	var currentScrollPos = window.pageYOffset;
	// Scroll down by the height of the viewport
	window.scrollTo(0, currentScrollPos + window.innerHeight);
	`
}
