/*
MySQL实例cpu使用率(占操作系统总数 )，
MySQL实例内存使用率(占操作系统总数)。
*/

package info

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func GetMemCpuInfo(addr string) {
	var s Signature
	log.Debug(addr)
	rc := ReadConf()
	s.AccessKeyId = rc.Ak
	s.Action = rc.Action
	s.DBInstanceId = rc.DBInstanceId
	s.Format = "JSON"
	s.Key = rc.Key
	s.SignatureMethod = "HMAC-SHA1"
	s.SignatureVersion = "1.0"
	s.Version = "2014-08-15"
	now := time.Now()
	year, mon, day := now.UTC().Date()
	hour, min, sec := now.UTC().Clock()
	s.Timestamp = fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02dZ", year, mon, day, hour, min, sec)

	t := time.Now()
	d, _ := time.ParseDuration("-482m")
	tm := t.Add(d)
	s.StartTime = tm.Format("2006-01-02T15:04Z")
	k, _ := time.ParseDuration("-480m")
	tk := t.Add(k)
	s.EndTime = tk.Format("2006-01-02T15:04Z")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s.SignatureNonce = strconv.Itoa(r.Intn(10000000000000))
	value := reflect.ValueOf(&s).Elem()
	typeOfT := value.Type()
	var str string
	for i := 0; i < value.NumField(); i++ {
		f := value.Field(i)
		str += "&" + percentEncode(typeOfT.Field(i).Name) + "=" + percentEncode(fmt.Sprintf("%v", f.Interface()))

	}
	sk := string(str[1:])
	stringToSign := "GET&%2F&" + percentEncode(sk)
	Sk := rc.Sk + "&"
	sh1 := sha1.New
	hash := hmac.New(sh1, []byte(Sk))
	hash.Write([]byte(stringToSign))
	sha1 := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	sha1 = url.QueryEscape(sha1)
	url := "https://rds.aliyuncs.com" + "/?" + string(str[1:]) + "&Signature=" + sha1
	log.Debug(url)
	response, err := http.Get(url)
	var mc ALiYunReturnMC
	if err != nil {
		log.Error(err)

	} else {
		content, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(content, &mc)
		log.Debug(mc)
		response.Body.Close()
	}
	var mem SendALiYunINFO
	var cpu SendALiYunINFO
	var memcpu []SendALiYunINFO
	key := mc.PerformanceKeys.PerformanceKey
	for _, k := range key {
		for _, v := range k.Values.PerformanceValue {
			log.Infoln(v.Date)
			str := strings.Split(v.Value, "&")
			mem.Value, _ = strconv.ParseFloat(str[0], 64)
			mem.Counter = fmt.Sprintf("aliyun.rds.mem.memuse.percent/instanceId=%s", mc.DBInstanceID)
			memcpu = append(memcpu, mem)
			cpu.Value, _ = strconv.ParseFloat(str[1], 64)
			cpu.Counter = fmt.Sprintf("aliyun.rds.cpu.busy/instanceId=%s", mc.DBInstanceID)
			memcpu = append(memcpu, cpu)

		}
	}

	b, _ := json.Marshal(memcpu)
	log.Infof(string(b))
	res, _ := http.NewRequest("POST", fmt.Sprintf("http://%s/push", addr), strings.NewReader(string(b)))
	res.Header.Set("Content-Type", "application/json")
	res.Header.Set("Connection", "close")
	res.Header.Set("Token", rc.Token)
	client := &http.Client{}

	resp, _ := client.Do(res)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	log.Infof(string(data))

}

func percentEncode(r string) string {
	u := url.QueryEscape(r)
	u = strings.Replace(u, "+", "%20", -1)
	u = strings.Replace(u, "*", "%2A", -1)
	u = strings.Replace(u, "%7E", "~", -1)
	return u
}
