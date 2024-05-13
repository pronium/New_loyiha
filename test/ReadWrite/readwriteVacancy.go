package readwrite

import (
	"Job_Uz/methods"
	"encoding/json"
	"os"
)

func ReadVacancies(nameOfFile string) methods.Vacancies {
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
	vacancies := methods.Vacancies{}
	err = json.Unmarshal(data, &vacancies)
	if err != nil {
		panic(err)
	}
	f.Close()

	return vacancies
}

func WriteVacancies(nameOfFile string, vacancies methods.Vacancies) {
	f, err := os.OpenFile(nameOfFile, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	dataInJson, err := json.MarshalIndent(vacancies, "", "\t")
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
