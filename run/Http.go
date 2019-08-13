package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	hessian "github.com/apache/dubbo-go-hessian2"
	"github.com/apache/dubbo-go-hessian2/java_exception"
)

func main() {
	fmt.Println("server start")
	http.HandleFunc("/", test)
	_ = http.ListenAndServe(":30122", nil)
}

type TriggerParam struct {
	JobId                 int
	ExecutorHandler       string
	ExecutorParams        string
	ExecutorBlockStrategy string
	ExecutorTimeout       int
	LogId                 int
	LogDateTim            int
	GlueType              string
	GlueSource            string
	GlueUpdatetime        int
	BroadcastIndex        int
	BroadcastTotal        int
}

func (TriggerParam) JavaClassName() string {
	return "com.xxl.job.core.biz.model.TriggerParam"
}

type XxlRpcRequest struct {
	RequestId        string
	CreateMillisTime int
	AccessToken      string
	ClassName        string
	MethodName       string
	ParameterTypes   []java_exception.Class
	Parameters       []interface{}
	Version          string
}

func (XxlRpcRequest) JavaClassName() string {
	return "com.xxl.rpc.remoting.net.params.XxlRpcRequest"
}

func test(w http.ResponseWriter, r *http.Request) {

	fmt.Println("aasdf")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	fmt.Println(body)
	fmt.Println(string(body))
	f, _ := os.Create("test.dat")
	_, _ = f.Write(body)
	_ = f.Close()
	SerializerHttp(body)
}

func SerializerHttp(bt []byte) {

	hessian.RegisterPOJO(XxlRpcRequest{})
	hessian.RegisterPOJO(TriggerParam{})
	d := hessian.NewDecoder(bt)
	res, err := d.Decode()
	resJson, _ := json.Marshal(res)
	if err != nil {
		fmt.Printf("Decode() = %+v", err)
	}
	fmt.Printf("decode = %v, %v\n", string(resJson), err)
}
