package buttgenerator

const svg_js = `
function load() {
	var svg = document.getElementsByTagName("svg")[0]
	svg.onclick = function() {
		console.log("Click!")
	}
}
window.onload = load
`
