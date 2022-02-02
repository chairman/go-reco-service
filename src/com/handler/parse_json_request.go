package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Rule struct {
	ruleId      int
	description string
}

type AutotaskRequest struct {
	ruleId      int    `json:"ruleId"`
	description string `json:"description"`
}

func Test(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	defer fmt.Fprintf(w, "ok\n")

	fmt.Println("method:", r.Method)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	println("json:", string(body))

	var a AutotaskRequest
	if err = json.Unmarshal(body, &a); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return
	}
	fmt.Printf("%+v", a)

}

func ParseJsonRequestHandler(w http.ResponseWriter, r *http.Request) {
	var p Rule
	//
	//curl -H "Content-Type:application/json" -XPOST http://127.0.0.1:8080/index/parse_json_request -d '{"ruleId":1,"order":1,"description":"好友圈抽样","op":"zhencw","selector":"{\"$eq":["\humanSysBizScene\",\"feed\"]}"}'
	//curl -H "Content-Type:application/json" -XPOST http://127.0.0.1:8080/index/parse_json_request -d '{"ruleId"=1}'
	// 将请求体中的 JSON 数据解析到结构体中
	// 发生错误，返回400 错误码
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Person: %+v", p)
}
