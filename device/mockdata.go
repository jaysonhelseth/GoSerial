package device

func GetData() (bytes []byte, err error) {
	str := `{ "msg": "Don't run with scissors." }`
	return []byte(str), nil
}
