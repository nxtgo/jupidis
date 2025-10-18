# jupidis

simple redis server implementation written in go.

# commands

> active checkbox means the command is implemented into jupidis.

## general

* [x] command
* [x] ping
* [x] flush
* [x] del
* [x] exists
* [x] keys
* [x] type
* [ ] expire
* [ ] pexpire
* [ ] ttl
* [ ] pttl
* [ ] persist
* [ ] rename
* [ ] renamenx
* [ ] save
* [ ] bgsave
* [ ] lastsave
* [ ] info
* [ ] config get
* [ ] config set
* [ ] dbsize
* [ ] flushdb
* [ ] flushall
* [ ] randomkey
* [ ] dump
* [ ] restore
* [ ] memory usage
* [ ] scan
* [ ] sscan
* [ ] hscan
* [ ] zscan

## strings

* [x] append
* [x] get
* [x] set
* [x] incr
* [x] incrby
* [x] decr
* [x] decrby
* [ ] mget
* [ ] mset
* [ ] strlen
* [ ] getset
* [ ] getrange
* [ ] setrange

## hashes

* [x] hget
* [x] hgetall
* [x] hset
* [ ] hdel
* [ ] hexists
* [ ] hkeys
* [ ] hvals
* [ ] hlen
* [ ] hmget
* [ ] hmset

## lists

* [x] lpush
* [ ] rpush
* [ ] lpop
* [ ] rpop
* [ ] llen
* [ ] lrange
* [ ] lindex
* [ ] lset
* [ ] lrem

## sets

* [x] sadd
* [x] srem
* [x] smove
* [x] scard
* [x] sismember
* [x] smembers
* [x] smismember
* [x] sunion
* [x] sunionstore
* [x] sinter
* [x] sinterstore
* [x] sdiff
* [x] sdiffstore
* [ ] spop
* [ ] srandmember

## sorted sets

* [ ] zadd
* [ ] zcard
* [ ] zrange
* [ ] zrevrange
* [ ] zrem
* [ ] zscore
* [ ] zincrby
* [ ] zrangebyscore
* [ ] zrevrangebyscore
* [ ] zunionstore
* [ ] zinterstore

## pub/sub

* [ ] publish
* [ ] subscribe
* [ ] unsubscribe
* [ ] psubscribe
* [ ] punsubscribe

## JSON

* [ ] json.del
* [ ] json.forget
* [ ] json.get
* [ ] json.mget
* [ ] json.set
* [ ] json.merge
* [ ] json.type
* [ ] json.numincrby
* [ ] json.nummultby
* [ ] json.strappend
* [ ] json.strlen
* [ ] json.arrappend
* [ ] json.arrindex
* [ ] json.arrinsert
* [ ] json.arrlen
* [ ] json.arrpop
* [ ] json.arrtrim
* [ ] json.objkeys
* [ ] json.objlen
* [ ] json.toggle
* [ ] json.clear
* [ ] json.debug
* [ ] json.resp

# license

CC0 1.0
