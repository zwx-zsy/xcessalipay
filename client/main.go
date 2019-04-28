package main

import (
	pb "XcessAlipay/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

// 此处应与服务器端对应
const address  = "127.0.0.1:8080"

/**
    1. 创建groc连接器
    2. 创建grpc客户端,并将连接器赋值给客户端
    3. 向grpc服务端发起请求
    4. 获取grpc服务端返回的结果

 */
type SomeOne struct {
	Subject string `json:"subject"`
	Out_trade_no string `json:"out_trade_no"`
	Total_amount string `json:"total_amount"`
	Product_code string `json:"product_code"`
	Body string `json:"body"`

}

type result struct {
	AppId string `json:"app_id"`
}

func main()  {

	// 创建一个grpc连接器
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil{
		fmt.Println("asdsa")
		fmt.Println(err)
	}
	// 当请求完毕后记得关闭连接,否则大量连接会占用资源

	defer conn.Close()
	// 创建grpc客户端
	c := pb.NewAliPayServiceClient(conn)

	//name := &[]pb.AliPayData{}

	// 客户端向grpc服务端发起请求
	result, err := c.TradePagePay(context.Background(), &pb.AliPayRequest{Mid:"hdkjhsjkf",Data: map[string]string{
		"NotifyURL":"http://203.86.24.181:3000/alipay",
		"ReturnURL":"http://liveadmin.cczhou.cn:32080",
		"Body":"skdyhfiuwehif",
		"Subject":"测试支付系统-支付宝",
		"OutTradeNo":"ueoiruwo7981789231",
		"TotalAmount":"1000.00",
		"ProductCode":"FAST_INSTANT_TRADE_PAY",
	}})
	response, err := c.TradeRefund(context.Background(), &pb.AliRefundRequest{OutTradeNo: "ueoiruwo7981789231", RefundAmount: "1000.00"})
	if err !=nil{
		fmt.Println("shdjkfsd",response)
	}
	fmt.Println(response)
	fmt.Println(result)
	//result, err := c.TradeApp(context.Background(), &pb.AliRequest{Name: "shdjkfsd"})
	if err != nil{
		fmt.Println(err.Error())
		fmt.Println("请求失败!!!")
		return
	}

	refoundResponse, err := c.FundTransToAccountTransfer(context.Background(), &pb.AliPayRequest{Data: map[string]string{"OutBizNo": "shgfhdksfasd", "PayeeAccount": "lsvhnj7045@sandbox.com", "Amount": "1.00"}})
	if err !=nil{
		fmt.Println(err.Error())
	}else{
		fmt.Println(refoundResponse)
	}
	// 获取服务端返回的结果
	//fmt.Println(result.Data["sign"])
	//fmt.Println(result.Data["method"])
	//fmt.Println(result.Data["biz_content"])
	//fmt.Println(result.Data["app_id"])
	//var someOne SomeOne
	//if err := json.Unmarshal([]byte(result.Data["biz_content"]), &someOne); err == nil {
	//	fmt.Println(someOne.Out_trade_no)
	//} else {
	//	fmt.Println(err)
	//}
}