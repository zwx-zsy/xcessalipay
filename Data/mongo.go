package Data

import (
	"XcessAlipay/Config"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)


var DB *mgo.Database

type Middleware struct {
	session *mgo.Session
}

func NewMiddleware(session *mgo.Session) *Middleware {
	return &Middleware{
		session: session,
	}
}

func (m *Middleware) Connect() *mgo.Session {
	s := m.session.Clone()
	DB = s.DB(Config.ServerConf.DBConf.DatabaseName)
	//defer s.Close()
	//c.Next()
	return s
}

//func Dial(router *gin.Engine) {
//	Session, err := mgo.Dial(Config.MongoDB{}.String())
//	if err != nil {
//		log.Fatal(err)
//	}
//	Session.SetMode(mgo.Eventual, true)
//	if err != nil {
//		log.Fatal(err)
//	}
//	// middleware
//	middleware := NewMiddleware(Session)
//	router.Use(middleware.Connect)
//
//}



func init() {

	yamlFile, err := ioutil.ReadFile(Config.CONFPATH)
	if err != nil {
		log.Fatal(err)
	} else {
		err = yaml.Unmarshal(yamlFile, &Config.ServerConf)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func init() {
	Session, err := mgo.Dial(Config.ServerConf.DBConf.String())
	fmt.Println(Session)
	if err != nil {
		log.Fatal(err)
	}
	Session.SetMode(mgo.Eventual, true)
	if err != nil {
		log.Fatal(err)
	}

	middleware := NewMiddleware(Session)
	middleware.Connect()
}
