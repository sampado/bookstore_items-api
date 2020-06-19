# bookstore_items-api


## Environment setup

### Elastic Search
* Download image
https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html

```
docker pull docker.elastic.co/elasticsearch/elasticsearch:7.7.1
```
* Run a local instance
```
docker run --name my-elastic -p 9200:9200 -p 9300:9300 --network dev-network -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.7.1
```

* Create Index
```
PUT http://localhost:9200/items
```
Body
```
{
    "settings" : {
        "index" : {
            "number_of_shards" : 4, 
            "number_of_replicas" : 2
        }
    }
}
```
