package buttgenerator

import (
	"strings"
)

func Butt(spot bool) string {
	butt := strings.Replace(`
	<svg class="butt" version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="640px" height="480px" viewBox="0 0 640 480" enable-background="new 0 0 640 480" xml:space="preserve">
	<path id="butt-body" fill="#DEC276" d="M171,91l291-16c0,0,34,98,28,120s-16,72-54,80s-72.29,6-105.145-36c0,0-14.855,74-64.855,64S156,277,171,91z "/>
	<polyline id="right-arm" fill="none" stroke="#00093F" stroke-width="4" stroke-miterlimit="10" points="514.667,202.333 502.667,197 544,135 486.311,141.616 "/>
	<polyline id="left-arm" fill="none" stroke="#00093F" stroke-width="4" stroke-miterlimit="10" points="171,294.333 181.333,273 96.666,269 168.657,199.667 "/>
	<path id="left-leg" fill="none" stroke="#00093F" stroke-width="4" stroke-miterlimit="10" d="M268,405l40-18c0,0-24-55.333-20-79.333"/>
	<path id="right-leg" fill="none" stroke="#00093F" stroke-width="4" stroke-miterlimit="10" d="M408.666,386.667c0,0,26-12.666,28-12 c0,0-15.999-91.334-3.333-91.334"/>
	<path id="mouth" fill="none" stroke="#001F0E" stroke-width="2" stroke-miterlimit="10" d="M300,189.465c0,0,4,12.201,24.666,10.202 c20.667-2,20.001-10.202,20.001-10.202"/>
	<path id="right-eye" fill="#1F003F" stroke="#001F0E" stroke-miterlimit="10" d="M356.667,127c0,0,5.333,2,3.333,8S347.334,130.333,356.667,127z"/>
	<path id="left-eye" fill="#13003F" stroke="#001F0E" stroke-miterlimit="10" d="M252.667,147c0,0,10.667,2.667,2.667,9.333 S244,148.333,252.667,147z"/>`,
	"\n", "", -1)

	if spot {
		butt += `<circle id="spot" color=#1F003F" stroke-miterlimit="10" cx=360 cy=150 r=2>`
	}

	butt += `</svg>`

	return butt
}

