package core

func router(message string, addr string) ([]byte, error) {

	operationCode := message[:4]
	if operationCode == "0000" {
		return RegisterComponent(message[4:], addr)
	}

	token := "test123"
	if !checkToken(addr, token) {
		return nil, nil
	}

	return nil, nil

}

func checkToken(addr string, token string) bool {
	return true
}

func RegisterComponent(name string, addr string) ([]byte, error) {
	registery := GetRegistery()

	token, err := registery.AddComponent(name, addr)

	return []byte(token), err
}
