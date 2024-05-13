package readwrite

import (
	"Job_Uz/methods"
	"encoding/json"
	"os"
)

func ReadUsers(nameOfFile string) methods.Users {
	f, err := os.OpenFile(nameOfFile, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	length, err := f.Stat()
	if err != nil {
		panic(err)
	}

	data := make([]byte, length.Size())
	_, err = f.Read(data)
	if err != nil {
		panic(err)
	}

	if !json.Valid(data) {
		panic("Invalid Json")
	}
	users := methods.Users{}
	err = json.Unmarshal(data, &users)
	if err != nil {
		panic(err)
	}
	f.Close()

	return users
}

func WriteUsers(nameOfFile string, Users methods.Users) {
	f, err := os.OpenFile(nameOfFile, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	dataInJson, err := json.MarshalIndent(Users, "", "\t")
	if err != nil {
		panic(err)
	}
	// fmt.Println("Json = ", string(dataInJson))
	f.Truncate(0)
	f.Seek(0, 0)

	_, err = f.Write(dataInJson)
	if err != nil {
		panic(err)
	}

	f.Close()
}
