package elasticsearch

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/olivere/elastic/v7"
	"sync"
)

var (
	once     sync.Once
	conf     Elasticsearch
	esClient *elastic.Client
	err      error
)

// es配置信息
type Elasticsearch struct {
	Hosts     []string  //host地址
	Sniff     bool      //是否使用监听机制，新加入节点或者有节点死掉
	BasicAuth basicAuth //如果设置了用户名密码认证
}

type basicAuth struct {
	UserName string //用户名
	PassWord string //密码
}

func Init() (*elastic.Client, error) {
	once.Do(func() {
		conf = Elasticsearch{}
		err = config.Get("elasticsearch").Scan(&conf)
		if err != nil {
			log.Fatal(err)
		}
		var options []elastic.ClientOptionFunc
		// 配置host
		if conf.Hosts != nil {
			options = append(options, elastic.SetURL(conf.Hosts...))
		}
		// 如果配置了不检测地址，在调试的时候可用
		if conf.Sniff == false {
			options = append(options, elastic.SetSniff(conf.Sniff))
		}
		// 如果有用户名和密码
		if conf.BasicAuth.UserName != "" {
			elastic.SetBasicAuth(conf.BasicAuth.UserName, conf.BasicAuth.PassWord)
		}
		// 创建一个es客户端
		esClient, err = elastic.NewClient(options...)
		if err != nil {
			log.Fatal(err)
		}
	})
	return esClient, err
}
