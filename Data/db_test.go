package Data

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func GetUser(i int) {
	fmt.Println(i)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.wsgjx.top/user/content/phonelist?num=1000",nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Set("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySW5mbyI6eyJpZCI6IjVjYWRhZDI4YjEyNDBjMzFkMzU0NTM3MyIsIm1vYmlsZSI6IjE4NjU1MTIwMjQwIiwiY291bnRyeV9jb2RlIjoiODYiLCJjcmVhdGVfdGltZSI6MTU1NDg4NTkyOH0sImlhdCI6MTU1NjI2Njc0MywiZXhwIjoxNTU4ODU4NzQzfQ.ZKFyvsVKBkSn7xEo4JCUV898Y0Ax_VJK93e40JJEXYk")
	resp, err := client.Do(req)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
	}
	result := map[string]interface{}{}
	json.Unmarshal([]byte(string(body)),&result)
	//fmt.Println(result["data"].(map[string]interface {})["phonelist"])
	//fmt.Println(reflect.TypeOf(result["data"].(map[string]interface {})["phonelist"]))
	userone := &User{}
	userone2 := &User{}
	for _, v := range result["data"].(map[string]interface {})["phonelist"].([]interface {}) {
		Err := Users().Find(bson.M{"UserCode": v.(map[string]interface{})["phone"].(string)}).One(userone2)
		//fmt.Println(userone2)
		if Err != nil{
			fmt.Println(Err.Error())
			userone.WxOpenId = v.(map[string]interface{})["nickname"].(string)
			userone.UserCode = v.(map[string]interface{})["phone"].(string)
			userone.PersonId = v.(map[string]interface{})["phone"].(string)
			userone.Id = bson.NewObjectId()
			userone.CreateDateTime = time.Now()
			fmt.Println(userone)
			e := Users().Insert(userone)
			if e != nil {
				fmt.Println(e)
			}
		}else{
			//fmt.Println(Err)
			//fmt.Println(userone2)
			continue

		}


	}
}

func TestDb(t *testing.T)  {
	//slices := make([]int,10,15)
	i:=1
	for i<10{
		GetUser(i)
		i+=1
	}
}
