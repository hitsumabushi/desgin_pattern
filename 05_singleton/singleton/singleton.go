package singleton

// instance を保持する変数
// singleton の中や、 GetInstance内に書いてしまうと、一意性を保証することが難しいので、まずここで宣言する
// nilで初期化されるので、nilかどうかですでに生成済みか判別できる
// また、ここで初期値を入れてしまうことで、thread safeにできる
// 詳細: <http://marcio.io/2015/07/singleton-pattern-in-go/>
var instance *singleton = newInstance()

func newInstance() *singleton {
	return &singleton{}
}

// Singleton は公開しない
type singleton struct{}

// Singletonインスタンスを得るためのGetInstance は公開する
func GetInstance() *singleton {
	return instance
}
