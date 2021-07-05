package device

func GetData(b []byte) (count int, err error) {
	str := `{ "msg": "Don't run with scissors." }`
	copy(b, []byte(str))

	return len(b), nil
}
