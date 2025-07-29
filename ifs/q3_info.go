package ifs

//import (
//	"fmt"
//	"strconv"
//	"tty28/base"
//)
//
//type Info struct {
//	TotalBet string
//	Values   map[int]float64
//}
//
//func getInfo(id int) (*Info, error) {
//	var resp struct {
//		Code int    `json:"code"`
//		Msg  string `json:"msg"`
//		Data struct {
//			Id       int    `json:"id"`
//			TotalBet string `json:"total_bet"`
//			List     []struct {
//				No  int         `json:"no"`
//				Odd interface{} `json:"odd"`
//			} `json:"list"`
//		} `json:"data"`
//	}
//
//	if err := Exec("templates/info.tpl", struct {
//		Token string
//		Id    int
//	}{Token: base.Config.Token, Id: id}, &resp); err != nil {
//		return nil, err
//	}
//
//	if resp.Code != 0 {
//		return nil, fmt.Errorf("错误代码 [%d] ，错误信息[%s]", resp.Code, resp.Msg)
//	}
//
//	values := make(map[int]float64)
//	for _, value := range resp.Data.List {
//		odd, err := strconv.ParseFloat(fmt.Sprintf("%v", value.Odd), 64)
//		if err != nil {
//			return nil, err
//		}
//
//		values[value.No] = odd
//	}
//
//	return &Info{TotalBet: resp.Data.TotalBet, Values: values}, nil
//}
