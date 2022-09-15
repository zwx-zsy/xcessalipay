package Data

import (
	"github.com/yancyzhou/xcessalipay/Config"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

var (
	session *mgo.Session
)

func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:          Config.ServerConf.DBConf.Address,
		Direct:         false,
		Timeout:        time.Second * 60,
		FailFast:       false,
		Database:       Config.ServerConf.DBConf.AuthDBName,
		ReplicaSetName: Config.ServerConf.DBConf.ReplicaSetName,
		Source:         "admin",
		Service:        "",
		ServiceHost:    "",
		Mechanism:      "",
		Username:       Config.ServerConf.DBConf.User,
		Password:       Config.ServerConf.DBConf.Password,
		PoolLimit:      4096,
		DialServer:     nil,
		Dial:           nil,
	}

	sess, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal(err.Error())
	}
	session = sess
	session.SetMode(mgo.Monotonic, true)
}

type SessionStore struct {
	session *mgo.Session
}

func NewSessionStore() *SessionStore {
	return &SessionStore{
		session: session.Copy(),
	}
}

func (s *SessionStore) C(name string) *mgo.Collection {
	return s.session.DB(Config.ServerConf.DBConf.DatabaseName).C(name)
}

func (s *SessionStore) Close() {
	s.session.Close()
}

func GetErrNotFound() error {
	return mgo.ErrNotFound
}
