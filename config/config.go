package config

import (
	"golang-question/errorx"
	"reflect"
)

type Manager[T any] interface {
	Get() T
	Update(T) errorx.Error
	OnChange(func(T)) (cancel func())
	Watch() Manager[T]     //執行Watch後，會開始監聽配置的變化，並在變化時自動更新 否則每次Get都會從數據源取得最新資料
	InitData(T) Manager[T] //如果數據源沒有資料，則使用InitData put資料
}

type localConfig[T any] struct {
	data     T
	watchers []func(T)
	isWatch  bool
}

func (l *localConfig[T]) Get() T {
	return l.data
}

// Update 更新配置，依次调用全部的OnChange的回調
func (l *localConfig[T]) Update(newData T) errorx.Error {
	l.data = newData
	if l.isWatch {
		for _, watcher := range l.watchers {
			// run watcher
			watcher(newData)
		}
	}
	return nil
}

// OnChange 註冊變更回調, 並返回解除註冊的函數
func (l *localConfig[T]) OnChange(callback func(T)) func() {
	l.watchers = append(l.watchers, callback)
	return func() {
		// Remove the callback from watchers
		for i, w := range l.watchers {
			if &w == &callback {
				l.watchers = append(l.watchers[:i], l.watchers[i+1:]...)
				break
			}
		}
	}
}

// Watch 開始監聽配置的變化
func (l *localConfig[T]) Watch() Manager[T] {
	l.isWatch = true
	return l
}

// InitData 如果數據源沒有資料，則使用InitData put資料
func (l *localConfig[T]) InitData(initialData T) Manager[T] {
	if reflect.ValueOf(l.data).IsZero() {
		l.Update(initialData)
	}
	return l
}

func Local[T any]() Manager[T] {
	return &localConfig[T]{}
}
