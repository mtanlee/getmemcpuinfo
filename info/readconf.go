package info

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadConf() GetInfoConf {
	var rconf GetInfoConf
	filepath := "/etc/getinfoconf/"
	name := "getinfo.conf"
	filename := filepath + name
	f, err := os.OpenFile(filename, os.O_RDONLY|os.O_SYNC, os.ModeType)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	r, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(r, &rconf)
	return rconf
}

func CreateConf() {
	var (
		f        *os.File
		filepath = "/etc/getinfoconf/"
		fname    = "getinfo.conf"
		wconf    []GetInfoConf
	)

	err := os.MkdirAll(filepath, 0777)
	if err != nil {
		return
	}
	filename := filepath + fname
	if !checkFile(filename) {
		f, err = os.Create(filename)
		if err != nil {
			panic(err)
		}
		wc, _ := json.Marshal(wconf)
		f.Write([]byte(wc))
	}
	defer f.Close()
}
func checkFile(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist

}
