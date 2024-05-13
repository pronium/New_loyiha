package readwrite

import (
	"Job_Uz/methods"
	"encoding/json"
	"os"
)

func ReadCompanies(nameOfFile string) methods.Companies {
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
	companies := methods.Companies{}
	err = json.Unmarshal(data, &companies)
	if err != nil {
		panic(err)
	}
	f.Close()

	return companies
}

func WriteCompanies(nameOfFile string, Companies methods.Companies) {
	f, err := os.OpenFile(nameOfFile, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	dataInJson, err := json.MarshalIndent(Companies, "", "\t")
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
