package test

import (
	"bytes"
	"cointiger.com/verification/app/verify/api/internal/types"
	"cointiger.com/verification/common/middleware"
	"cointiger.com/verification/common/sign"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/threading"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

//var host = "http://127.0.0.1:8888"

//var host = "http://172.16.10.70:30021" //nodePort
//var host = "http://172.16.10.12:94" //Nginx
var host = "http://172.16.10.12:90" //Nginx

//done1
func TestMakeDetail(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	for i := 0; i < 1; i++ {
		threading.GoSafe(func() {
			req := &types.MakeDetaiRequest{
				Symbol:  "ethusdt",
				OrderId: "73",
			}
			jsonReq, _ := json.Marshal(req)
			//m := make(map[string]interface{})
			//json.Unmarshal(jsonReq, &m)
			//str := middleware.SortMap(m)
			str := string(jsonReq)
			apikey := "005f712e-5e79-4251-8bb2-213dfc36e4ad"
			apiSercet := "YTY5ZGM4NWYxODE4ZjQ2NzA0ZTRlZjJmY2M0MWYxODBhMDZhMDUwMzgzM2Y2YmQ3YzAxOTZmMmJmNjA4OWM5OQ=="
			str += apiSercet
			signStr := sign.Sign(apikey, str)
			fmt.Println("test:apikey:", apikey)
			fmt.Println("test:waitSign:", str)
			fmt.Println("test:signstr:", signStr)
			path := fmt.Sprintf("%s/api/v3/order/makedetail", host)
			fmt.Println(string(jsonReq))
			httpReq, _ := http.NewRequest("POST", path, bytes.NewReader(jsonReq))
			httpReq.Header.Add("Content-Type", "application/json")
			httpReq.Header.Add("api-key", apikey)
			httpReq.Header.Add("sign", signStr)
			cli := http.Client{}
			resp, err := cli.Do(httpReq)
			if err != nil {
				t.Error("err:", err)
			}
			defer resp.Body.Close()
			vals, err := io.ReadAll(resp.Body)
			t.Log("body:", string(vals), "err:", err)
			wg.Done()
		})
	}
	wg.Wait()
}

//done1
func TestMatchResult(t *testing.T) {
	//{"symbol":"ethusdt","starDate":"2022-09-03 00:00:00","endDate":"2022-09-06 23:59:59","from":"8054320","direct":"prev","size":"100"}
	//ISODate("2022-03-17T22:12:23.000+0000")
	req := &types.MatchResultRequest{
		StartDate: "2022-09-03 00:00:00",
		EndDate:   "2022-09-06 23:59:59",
		From:      "8054320",
		Direct:    "prev",
		Size:      "100",
		Symbol:    "ethusdt",
	}
	jsonReq, _ := json.Marshal(req)
	//m := make(map[string]interface{})
	//json.Unmarshal(jsonReq, &m)
	//str := middleware.SortMap(m)
	str := string(jsonReq)
	apikey := "56ac3e1a-108f-4f4f-8571-224731aaf996"
	apiSercet := "MjVmZWFkZjU1MGIwZTVmYTI0OGUzNzY3YmEyYzEzN2NhOTRmODRjNjE5ZDg2N2Y3ZWM5MjY4NTg1Y2IwOTlhNA=="
	str += apiSercet
	signStr := sign.Sign(apikey, str)
	fmt.Println("test:apikey:", apikey)
	fmt.Println("test:waitSign:", str)
	fmt.Println("test:signstr:", signStr)
	path := fmt.Sprintf("%s/api/v3/order/matchresults", host)
	fmt.Println(string(jsonReq))
	httpReq, _ := http.NewRequest("POST", path, bytes.NewReader(jsonReq))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("api-key", apikey)
	httpReq.Header.Add("sign", signStr)
	cli := http.Client{}
	resp, err := cli.Do(httpReq)
	if err != nil {
		t.Error("err:", err)
	}
	defer resp.Body.Close()
	vals, err := io.ReadAll(resp.Body)
	t.Log("body:", string(vals), "err:", err)
}

//done1
func TestOrderHistory(t *testing.T) {
	req := &types.OrderHistoryRequest{
		Symbol:    "ethusdt",
		Types:     "buy-limit",           //查询的订单类型组合，使用','分割	全部	buy-market：市价买, sell-market：市价卖, buy-limit：限价买, sell-limit：限价卖
		StartDate: "2022-09-04 00:00:00", //查询开始日期, 日期格式yyyy-mm-dd	当前日期减2天	起始日期和结束日志时间最大间隔2天
		EndDate:   "2022-09-06 00:00:00", //查询结束日期, 日期格式yyyy-mm-dd	当前日期	起始日期和结束日志时间最大间隔2天
		States:    "filled",              //查询的订单状态组合，使用','分割		new ：新订单, part_filled ：部分成交
		From:      "8054320",             //查询起始 ID （订单ID）		配合direct 使用
		Direct:    "prev",                //查询方向	next	prev 向前，next 向后
		Size:      "100",                 //查询记录大小100	最多一次查询100条数据
	}
	jsonReq, _ := json.Marshal(req)
	//m := make(map[string]interface{})
	//json.Unmarshal(jsonReq, &m)
	//str := middleware.SortMap(m)
	str := string(jsonReq)
	apikey := "56ac3e1a-108f-4f4f-8571-224731aaf996"
	apiSercet := "MjVmZWFkZjU1MGIwZTVmYTI0OGUzNzY3YmEyYzEzN2NhOTRmODRjNjE5ZDg2N2Y3ZWM5MjY4NTg1Y2IwOTlhNA=="
	str += apiSercet
	signStr := sign.Sign(apikey, str)
	fmt.Println("test:apikey:", apikey)
	fmt.Println("test:waitSign:", str)
	fmt.Println("test:signstr:", signStr)
	path := fmt.Sprintf("%s/api/v3/order/history", host)
	fmt.Println(string(jsonReq))
	httpReq, _ := http.NewRequest("POST", path, bytes.NewReader(jsonReq))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("api-key", apikey)
	httpReq.Header.Add("sign", signStr)
	cli := http.Client{}
	resp, err := cli.Do(httpReq)
	if err != nil {
		t.Error("err:", err)
	}
	defer resp.Body.Close()
	vals, err := io.ReadAll(resp.Body)
	t.Log("body:", string(vals), "err:", err)
}

//done1
func TestOrderCurrent(t *testing.T) {
	//{"symbol":"ethusdt","types":"buy-limit","states":"part_filled","from":"8054137","direct":"prev","size":"50"}
	//{"symbol":"ethusdt","types":"sell-limit","states":"part_filled","from":"8054349","direct":"next","size":"50"}

	var wg sync.WaitGroup
	count := 1
	wg.Add(count)
	for i := 0; i < count; i++ {
		threading.GoSafe(func() {
			req := &types.OrderCurrentReq{
				Symbol: "ethusdt",
				Types:  "sell-limit",  //查询的订单类型组合，使用','分割	全部	buy-market：市价买, sell-market：市价卖, buy-limit：限价买, sell-limit：限价卖
				States: "part_filled", //查询的订单状态组合，使用','分割		new ：新订单, part_filled ：部分成交
				From:   "8054349",     //查询起始 ID （订单ID）		配合direct 使用
				Direct: "next",        //查询方向	next	prev 向前，next 向后
				Size:   "50",          //查询记录大小	50	最多一次查询50条数据
			}
			jsonReq, _ := json.Marshal(req)
			//m := make(map[string]interface{})
			//json.Unmarshal(jsonReq, &m)
			//str := middleware.SortMap(m)
			str := string(jsonReq)
			apikey := "56ac3e1a-108f-4f4f-8571-224731aaf996"
			apiSercet := "MjVmZWFkZjU1MGIwZTVmYTI0OGUzNzY3YmEyYzEzN2NhOTRmODRjNjE5ZDg2N2Y3ZWM5MjY4NTg1Y2IwOTlhNA=="
			str += apiSercet
			signStr := sign.Sign(apikey, str)
			fmt.Println("test:apikey:", apikey)
			fmt.Println("test:waitSign:", str)
			fmt.Println("test:signstr:", signStr)
			path := fmt.Sprintf("%s/api/v3/order/current", host)
			fmt.Println(string(jsonReq))
			httpReq, _ := http.NewRequest("POST", path, bytes.NewReader(jsonReq))
			httpReq.Header.Add("Content-Type", "application/json")
			httpReq.Header.Add("api-key", apikey)
			httpReq.Header.Add("sign", signStr)
			cli := http.Client{}
			resp, err := cli.Do(httpReq)
			if err != nil {
				t.Error("err:", err)
			}
			defer resp.Body.Close()
			vals, err := io.ReadAll(resp.Body)
			t.Log("body:", string(vals), "err:", err)
			wg.Done()
		})
	}
	wg.Wait()
}

//done1
func TestOrderDetail(t *testing.T) {
	req := &types.MakeDetaiRequest{
		Symbol:  "ethusdt",
		OrderId: "8054146",
	}
	jsonReq, _ := json.Marshal(req)
	//m := make(map[string]interface{})
	//json.Unmarshal(jsonReq, &m)
	//str := middleware.SortMap(m)
	str := string(jsonReq)
	apikey := "56ac3e1a-108f-4f4f-8571-224731aaf996"
	apiSercet := "MjVmZWFkZjU1MGIwZTVmYTI0OGUzNzY3YmEyYzEzN2NhOTRmODRjNjE5ZDg2N2Y3ZWM5MjY4NTg1Y2IwOTlhNA=="
	str += apiSercet
	signStr := sign.Sign(apikey, str)
	fmt.Println("test:apikey:", apikey)
	fmt.Println("test:waitSign:", str)
	fmt.Println("test:signstr:", signStr)
	path := fmt.Sprintf("%s/api/v3/order/detail", host)
	fmt.Println(string(jsonReq))
	httpReq, _ := http.NewRequest("POST", path, bytes.NewReader(jsonReq))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("api-key", apikey)
	httpReq.Header.Add("sign", signStr)
	cli := http.Client{}
	resp, err := cli.Do(httpReq)
	if err != nil {
		t.Error("err:", err)
	}
	defer resp.Body.Close()
	vals, err := io.ReadAll(resp.Body)
	t.Log("body:", string(vals), "err:", err)
}

//done1
func TestCreateOrder(t *testing.T) {
	//{"symbol":"ethusdt","price":"2588","volumn":"3.1","side":"BUY","type":"1","time":"1662431449"}
	req := &types.OrderRequest{
		Symbol: "ethusdt",
		Price:  "2588",
		Volume: "0.123",
		Side:   "BUY", //买卖方向 买BUY 卖SELL
		Type:   "1",   //1 ：限价交易，2：市价交易
		Time:   "1662431449",
	}
	//{"symbol":"ethusdt","price":"2388","volumn":"1","side":"buy","type":"1","time":"1662431449"}
	jsonReq, _ := json.Marshal(req)
	//m := make(map[string]interface{})
	//json.Unmarshal(jsonReq, &m)
	//str := middleware.SortMap(m)
	str := string(jsonReq)
	apikey := "56ac3e1a-108f-4f4f-8571-224731aaf996"
	apiSercet := "MjVmZWFkZjU1MGIwZTVmYTI0OGUzNzY3YmEyYzEzN2NhOTRmODRjNjE5ZDg2N2Y3ZWM5MjY4NTg1Y2IwOTlhNA=="
	str += apiSercet
	signStr := sign.Sign(apikey, str)
	fmt.Println("test:apikey:", apikey)
	fmt.Println("test:waitSign:", str)
	fmt.Println("test:signstr:", signStr)
	path := fmt.Sprintf("%s/api/v3/order", host)
	fmt.Println(string(jsonReq))
	httpReq, _ := http.NewRequest("POST", path, bytes.NewReader(jsonReq))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("api-key", apikey)
	httpReq.Header.Add("sign", signStr)
	cli := http.Client{}
	resp, err := cli.Do(httpReq)
	if err != nil {
		t.Error("err:", err)
	}
	defer resp.Body.Close()
	vals, err := io.ReadAll(resp.Body)
	t.Log("body:", string(vals), "err:", err)
}

func TestBatchCreateOrder(t *testing.T) {
	orders := types.BatchOrderRequest{
		Orders: []*types.BatchOrder{
			{
				COrderId: "1",
				Tag:      "1",
				Symbol:   "ethusdt",
				Price:    "2",
				Volume:   "3",
				Side:     "BUY", //买卖方向 买BUY 卖SELL
				Type:     "1",   //1 ：限价交易，2：市价交易
			},
			{
				COrderId: "2",
				Tag:      "2",
				Symbol:   "ethusdt",
				Price:    "2",
				Volume:   "3",
				Side:     "BUY", //买卖方向 买BUY 卖SELL
				Type:     "1",   //1 ：限价交易，2：市价交易
			},
			{
				COrderId: "3",
				Tag:      "3",
				Symbol:   "ethusdt",
				Price:    "2",
				Volume:   "3",
				Side:     "BUY", //买卖方向 买BUY 卖SELL
				Type:     "1",   //1 ：限价交易，2：市价交易
			},
		},
		Time: strconv.FormatInt(time.Now().Unix(), 10),
	}
	jsonReq, _ := json.Marshal(orders)
	//m := make(map[string]interface{})
	//json.Unmarshal(jsonReq, &m)
	//str := middleware.SortMap(m)
	str := string(jsonReq)
	apikey := "005f712e-5e79-4251-8bb2-213dfc36e4ad"
	apiSercet := "YTY5ZGM4NWYxODE4ZjQ2NzA0ZTRlZjJmY2M0MWYxODBhMDZhMDUwMzgzM2Y2YmQ3YzAxOTZmMmJmNjA4OWM5OQ=="
	str += apiSercet
	signStr := sign.Sign(apikey, str)
	fmt.Println("test:apikey:", apikey)
	fmt.Println("test:waitSign:", str)
	fmt.Println("test:signstr:", signStr)
	path := fmt.Sprintf("%s/api/v3/order/batch", host)
	fmt.Println(string(jsonReq))
	fmt.Println(path)
	httpReq, _ := http.NewRequest("POST", path, bytes.NewReader(jsonReq))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("api-key", apikey)
	httpReq.Header.Add("sign", signStr)
	cli := http.Client{}
	resp, err := cli.Do(httpReq)
	if err != nil {
		t.Error("err:", err)
	}
	defer resp.Body.Close()
	vals, err := io.ReadAll(resp.Body)
	t.Log("body:", string(vals), "err:", err)
}

//done1
func TestBatchCancelOrder(t *testing.T) {
	req := &types.OrderCancelRequest{
		OrderIdList: []*types.OrderCancelReqEntity{
			{Symbol: "ethusdt",
				OrderId: "73",
			},
		},
	}
	jsonReq, _ := json.Marshal(req)
	m := make(map[string]interface{})
	json.Unmarshal(jsonReq, &m)
	str := middleware.SortMap(m)
	apikey := "005f712e-5e79-4251-8bb2-213dfc36e4ad"
	apiSercet := "YTY5ZGM4NWYxODE4ZjQ2NzA0ZTRlZjJmY2M0MWYxODBhMDZhMDUwMzgzM2Y2YmQ3YzAxOTZmMmJmNjA4OWM5OQ=="
	str += apiSercet
	signStr := sign.Sign(apikey, str)
	fmt.Println("test:apikey:", apikey)
	fmt.Println("test:waitSign:", str)
	fmt.Println("test:signstr:", signStr)
	path := fmt.Sprintf("%s/api/v3/order/batchcancel", host)
	fmt.Println(string(jsonReq))
	httpReq, _ := http.NewRequest("POST", path, bytes.NewReader(jsonReq))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("api-key", apikey)
	httpReq.Header.Add("sign", signStr)
	cli := http.Client{}
	resp, err := cli.Do(httpReq)
	if err != nil {
		t.Error("err:", err)
	}
	defer resp.Body.Close()
	vals, err := io.ReadAll(resp.Body)
	t.Log("body:", string(vals), "err:", err)
}

func TestMarkdown(t *testing.T) {
	m := make(map[string]interface{})
	bs, _ := ioutil.ReadFile("/Users/chengjun/Desktop/solution/cointiger.com/cointiger-apis-services/doc/maker.json")
	json.Unmarshal(bs, &m)
	m1 := make(map[string]string, 0)
	for k, v := range m["definitions"].(map[string]interface{})["MakeDetailResponse"].(map[string]interface{})["properties"].(map[string]interface{}) {
		m1[k] = v.(map[string]interface{})["description"].(string)
	}
	v := &types.MakeDetailResponse{}
	val := reflect.ValueOf(v).Elem()

	var b strings.Builder
	b.WriteString("|  字段   | 类型  | 是否必须  | 描述 | \n")
	b.WriteString("|  :----|  :----:  |  :----:  |  :----: | \n")
	for i := 0; i < val.NumField(); i++ {
		name := val.Type().Field(i).Tag.Get("json")
		t := val.Type().Field(i).Type
		b.WriteString(fmt.Sprintf("|  %s   | %s  |  true |  %s |\n", name, t, m1[name]))
	}
	s := b.String()
	fmt.Println(s)
}

func TestNewBatchCancelOrder(t *testing.T) {
	req := &types.OrderCancelRequest{
		OrderIdList: []*types.OrderCancelReqEntity{
			{Symbol: "ethusdt",
				OrderId: "8054359",
			},
			{Symbol: "ethusdt",
				OrderId: "8054358",
			},
		},
	}
	jsonReq, _ := json.Marshal(req)
	//m := make(map[string]interface{})
	//json.Unmarshal(jsonReq, &m)
	//str := middleware.SortMap(m)
	str := string(jsonReq)
	apikey := "56ac3e1a-108f-4f4f-8571-224731aaf996"
	apiSercet := "MjVmZWFkZjU1MGIwZTVmYTI0OGUzNzY3YmEyYzEzN2NhOTRmODRjNjE5ZDg2N2Y3ZWM5MjY4NTg1Y2IwOTlhNA=="
	str += apiSercet
	signStr := sign.Sign(apikey, str)
	fmt.Println("test:apikey:", apikey)
	fmt.Println("test:waitSign:", str)
	fmt.Println("test:signstr:", signStr)
	path := fmt.Sprintf("%s/api/v3/order/newbatchcancel", host)
	fmt.Println(string(jsonReq))
	httpReq, _ := http.NewRequest("POST", path, bytes.NewReader(jsonReq))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("api-key", apikey)
	httpReq.Header.Add("sign", signStr)
	cli := http.Client{}
	resp, err := cli.Do(httpReq)
	if err != nil {
		t.Error("err:", err)
	}
	defer resp.Body.Close()
	vals, err := io.ReadAll(resp.Body)
	t.Log("body:", string(vals), "err:", err)
}
