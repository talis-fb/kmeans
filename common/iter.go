package common;


func Map[V any, T any](array []V, fn func(V) T) []T {
  output := make([]T, len(array))
  for i := range output {
    output[i] = fn(array[i])
  }
  return output;
}


func Filter[V any](array []V, fn func(V) bool) []V {
  output := make([]V, 0)

  for _, v := range array{
    if fn(v) {
      output = append(output, v)
    }
  }
  return output;
}

func Any[V any](array []V, fn func(V) bool) bool {
  for _, v := range array{
    if fn(v) {
      return true;
    }
  }
  return false;
}

func All[V any](array []V, fn func(V) bool) bool {
  // return !Any(array, fn);
  for _, v := range array{
    if !fn(v) {
      return false;
    }
  }
  return true;
}

func Sum(array []float64) float64 {
  var sum float64 = 0
  for _, v := range array{
    sum += v
  }
  return sum;
}


func IndexOfMin(array []float64) int {
  if array == nil {
    return -1
  }

	lowerValue := array[0]
  targetIndex := 0

	for i, v := range array[1:] {
		if v < lowerValue {
			lowerValue = v
      targetIndex = i
		}
	}

	return targetIndex
}

func Min(array []float64) float64 {
  ind := IndexOfMin(array)
	return array[ind]
}
