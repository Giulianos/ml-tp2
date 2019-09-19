#!/bin/bash

DATASET=datasets/titanic-discretized.csv

SIZE=20
mkdir datasets/$SIZE
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test1.csv -train datasets/$SIZE/train1.csv -test-size 0.1 -seed 4795
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test2.csv -train datasets/$SIZE/train2.csv -test-size 0.1 -seed 7642
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test3.csv -train datasets/$SIZE/train3.csv -test-size 0.1 -seed 9898
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test4.csv -train datasets/$SIZE/train4.csv -test-size 0.1 -seed 3954
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test5.csv -train datasets/$SIZE/train5.csv -test-size 0.1 -seed 9724

SIZE=25
mkdir datasets/$SIZE
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test1.csv -train datasets/$SIZE/train1.csv -test-size 0.2 -seed 6390
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test2.csv -train datasets/$SIZE/train2.csv -test-size 0.2 -seed 5383
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test3.csv -train datasets/$SIZE/train3.csv -test-size 0.2 -seed 458
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test4.csv -train datasets/$SIZE/train4.csv -test-size 0.2 -seed 9669
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test5.csv -train datasets/$SIZE/train5.csv -test-size 0.2 -seed 6002

SIZE=30
mkdir datasets/$SIZE
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test1.csv -train datasets/$SIZE/train1.csv -test-size 0.3 -seed 3879
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test2.csv -train datasets/$SIZE/train2.csv -test-size 0.3 -seed 3779
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test3.csv -train datasets/$SIZE/train3.csv -test-size 0.3 -seed 6685
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test4.csv -train datasets/$SIZE/train4.csv -test-size 0.3 -seed 2109
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test5.csv -train datasets/$SIZE/train5.csv -test-size 0.3 -seed 2150

SIZE=35
mkdir datasets/$SIZE
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test1.csv -train datasets/$SIZE/train1.csv -test-size 0.4 -seed 5651
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test2.csv -train datasets/$SIZE/train2.csv -test-size 0.4 -seed 4554
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test3.csv -train datasets/$SIZE/train3.csv -test-size 0.4 -seed 3195
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test4.csv -train datasets/$SIZE/train4.csv -test-size 0.4 -seed 6974
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test5.csv -train datasets/$SIZE/train5.csv -test-size 0.4 -seed 3227

SIZE=40
mkdir datasets/$SIZE
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test1.csv -train datasets/$SIZE/train1.csv -test-size 0.4 -seed 444
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test2.csv -train datasets/$SIZE/train2.csv -test-size 0.4 -seed 7074
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test3.csv -train datasets/$SIZE/train3.csv -test-size 0.4 -seed 5655
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test4.csv -train datasets/$SIZE/train4.csv -test-size 0.4 -seed 5591
go run cmd/splitset/split_set.go -ds $DATASET -test datasets/$SIZE/test5.csv -train datasets/$SIZE/train5.csv -test-size 0.4 -seed 2604