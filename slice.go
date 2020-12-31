package gotool

func intSliceDelete(slice []int, position int) (output []int) {
	length := len(slice)
	output = append(slice[0:position-1], slice[position:length]...)
	return output
}

func stringSliceDelete(slice []string, position int) (output []string) {
	length := len(slice)
	output = append(slice[0:position-1], slice[position:length]...)
	return output
}
