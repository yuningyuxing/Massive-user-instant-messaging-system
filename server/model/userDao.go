package model

//本文件用于编写对User对象操作的各种方法 主要来说就是增删改查
import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 声明一个全局变量
var (
	MyUserDao *UserDao
)

// 定义一个UserDao结构体 用于对User结构体的各种操作
type UserDao struct {
	//使用连接池进行链接管理
	pool *redis.Pool
}

// 使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// 根据用户id返回一个用户实例
// 我们传入一个链接通道，因为一般是有链接通道函数的函数调用这个函数
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	//从数据库中获取信息
	res, err := redis.String(conn.Do("HGet", "users", id)) //从redis数据库中获取信息
	if err != nil {
		//此时在users哈希中，没有找到对应id
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	//创建一个User实例
	user = &User{}
	//将获取的信息反序列化为User类型的结构体
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

// 完成登陆的校验Login 并返回对应登陆信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := this.pool.Get() //我们要在数据库中进行登陆校验 所以我们需要先从链接池中获取一个链接
	defer conn.Close()      //记得关闭链接
	//我们先查询这个用户是否存在 这里就可以通过链接和函数 获取这个用户的实例
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	//看看密码对不对
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
