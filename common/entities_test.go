package common

import "testing"

func TestPoint(t *testing.T) {
  p1 := Point{1, 2}
  p2 := Point{3, 4}

  t.Log(p1.EuclideanDistance(p2))
}
