package mypic

import "golang.org/x/tour/pic"

// Pic returns a slice of length dy, each element of which is a slice of dx
// 8-bit unsigned integers.
// https://tour.golang.org/moretypes/18
func Pic(dx, dy int) [][]uint8 {
	// Generate a slice of length dy
	// Each element of this slice is a slice of 8-bit unsigned integer
	sliceMain := make([][]uint8, dy)

	for i := range sliceMain {
		// Generate a slice of 8-bit unsigned integer of length dx
		sliceSub := make([]uint8, dx)

		for j := range sliceSub {
			// Generate an 8-bit unsigned integer using a function
			// Good examples are: i ^ j, i * j, (i + j) / 2
			sliceSub[j] = uint8(i ^ j)
		}

		sliceMain[i] = sliceSub
	}

	return sliceMain
}

// ShowPic returns a Base-64 encoded string that can be used to display an
// image.
// To display on the web, use the below format, with ${encodedString} being the
// string returned by this function:
// <img src={`data:image/png;base64,${encodedString}`} />
func ShowPic() {
	pic.Show(Pic)
}
