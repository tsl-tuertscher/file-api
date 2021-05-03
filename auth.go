package main

func CheckKey(keys []string, key string) bool {
	for i := 0; i < len(keys); i++ {
		if key == keys[i] {
			return true
		}
	}
	return false
}
