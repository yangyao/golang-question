package config

import "golang-question/errorx"

type Manager[T any] interface {
	Get() T
	Update(T) errorx.Error
	OnChange(func(T)) (cancel func())
	Watch() Manager[T]     //執行Watch後，會開始監聽配置的變化，並在變化時自動更新 否則每次Get都會從數據源取得最新資料
	InitData(T) Manager[T] //如果數據源沒有資料，則使用InitData put資料
}

func Local[T any]() Manager[T] {
	//TODO: implement
	return nil
}

func Etcd[T any]() Manager[T] {
	//TODO: implement
	return nil
}
