package main

const svg_js = `
function random_color() {
	return '#'+(0x1000000+(Math.random())*0xffffff).toString(16).substr(1,6)
}

function load() {
	// When you click on these, change the stroke or fill to a random color!
	var onclicks = {
		"stroke": ["left-arm", "right-arm", "left-leg", "right-leg", "mouth"],
		"fill":   ["butt-body", "left-eye", "right-eye"],
	}

	for (var attr in onclicks) {
		for (var i = 0; i < onclicks[attr].length; i++) {
			var onclick = function(attr) {
				return function() { this.setAttribute(attr, random_color()) }
			}
			document.getElementById(onclicks[attr][i]).onclick = onclick(attr)
		}
	}
}
window.onload = load
`
