# kafka to redisearch

Given a stream of CMS content from kafka, tidy it up and index it into redisearch.

This project is packaged as a helm chart, at: https://github.com/jspc/helm-charts

## schema

Given the body:

```json
{
    "operation": "CREATE",
    "message": {
        "slug": "hello-world!",
        "title": "Hello, world!",
        "author": "jspc",
        "body": "<h1>Hello World</h1>"
    }
}
```

Generate the redis hash:

```
 HGETALL hello-world!
1) "body"
2) "<h1>Hello World</h1>"
3) "author"
4) "jspc"
5) "title"
6) "Hello, world!"
7) "date"
8) "-62135596800"
```

Which can be searched as per:

```
127.0.0.1:6379> FT.SEARCH content "Hello, World!" LIMIT 0 1
1) (integer) 2
2) "hello-world!"
3) 1) "body"
   2) "<h1>Hello World</h1>"
   3) "author"
   4) "jspc"
   5) "title"
   6) "Hello, world!"
   7) "date"
   8) "-62135596800"
```
