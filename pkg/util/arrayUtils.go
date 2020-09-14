package util

import "strings"

func RemoveRepeatedStringElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func ContainsStringElement(arr []string, ele string) bool {
	for _, v := range arr {
		if v == ele {
			return true
		}
	}
	return false
}

func ContainsIntElement(arr []int, ele int) bool {
	for _, v := range arr {
		if v == ele {
			return true
		}
	}
	return false
}

func GetValueFromExtraArgs(args []string, key string) string {
	for _, arg := range args {
		split := strings.Split(arg, "=")
		if split[0] == key {
			return split[1]
		}
	}
	return ""
}
