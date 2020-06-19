package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/sampado/bookstore_utils-go/logger"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(c *elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {

	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),

		elastic.SetSniff(false),
		elastic.SetErrorLog(logger.GetLogger()),
		//log.New(os.Stderr, "ELASTIC - ERROR: ", log.LstdFlags)
		elastic.SetInfoLog(logger.GetLogger()),
	)

	if err != nil {
		panic(err)
	}

	Client.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.
		Index().
		Index(index).
		BodyJson(doc).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when tying to index document in index %s", index), err)
		return nil, err
	}

	return result, nil
}
