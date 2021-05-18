#!/bin/bash

# generate model with cache from ddl
fromDDLWithCache:
	taoctl template clean
	taoctl model mysql ddl -src="./sql/*.sql" -dir="./sql/model/cache" -cache

fromDDLWithoutCache:
	taoctl template clean;
	taoctl model mysql ddl -src="./sql/*.sql" -dir="./sql/model/nocache"


# generate model with cache from data source
user=root
password=password
datasource=127.0.0.1:3306
database=gozero

fromDataSource:
	taoctl template clean
	taoctl model mysql datasource -url="$(user):$(password)@tcp($(datasource))/$(database)" -table="*" -dir ./model/cache -c -style gozero
