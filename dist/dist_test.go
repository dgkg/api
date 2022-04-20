package dist

import (
	"testing"
)

// func main() {
// 	fmt.Printf("%f Miles\n", )
// 	fmt.Printf("%f Kilometers\n", distance(32.9697, -96.80322, 29.46786, -98.53506, "K"))
// 	fmt.Printf("%f Nautical Miles\n", distance(32.9697, -96.80322, 29.46786, -98.53506, "N"))
// }

func TestDistance(t *testing.T) {
	distance(32.9697, -96.80322, 29.46786, -98.53506, "M")
}
func BenchmarkDistance(b *testing.B) {
	for n := 0; n < b.N; n++ {
		distance(32.9697, -96.80322, 29.46786, -98.53506, "M")
	}
}

func BenchmarkDistance1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		distance(32.9697, -96.80322, 29.46786, -98.53506, "M")
	}
}

func BenchmarkDistance2(b *testing.B) {
	c := Coordinates{32.9697, -96.80322}
	d := Coordinates{32.9697, -96.80322}
	for n := 0; n < b.N; n++ {
		c.distance2(d)
	}
}

func BenchmarkDistance3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		distance3()
	}
}
