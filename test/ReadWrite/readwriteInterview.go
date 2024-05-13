package readwrite

import (
	"Job_Uz/methods"
	"encoding/json"
	"os"
)

func ReadInterviews(nameOfFile string) methods.Interviews {
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
	interviews := methods.Interviews{}
	err = json.Unmarshal(data, &interviews)
	if err != nil {
		panic(err)
	}
	f.Close()

	return interviews
}

func WriteInterviews(nameOfFile string, interviews methods.Interviews) {
	f, err := os.OpenFile(nameOfFile, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	dataInJson, err := json.MarshalIndent(interviews, "", "\t")
	if err != nil {
		panic(err)
	}

	f.Truncate(0)
	f.Seek(0, 0)

	_, err = f.Write(dataInJson)
	if err != nil {
		panic(err)
	}

	f.Close()
}
