#!/bin/bash

PREDATTR=Survived
EVAL=rf-gini

mkdir evals

mkdir evals/$EVAL

SIZE=20
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test1.csv -train datasets/$SIZE/train1.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null
echo  >> evals/$EVAL/$SIZE.txt
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test2.csv -train datasets/$SIZE/train2.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null
echo  >> evals/$EVAL/$SIZE.txt
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test3.csv -train datasets/$SIZE/train3.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null
echo  >> evals/$EVAL/$SIZE.txt
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test4.csv -train datasets/$SIZE/train4.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null
echo  >> evals/$EVAL/$SIZE.txt
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test5.csv -train datasets/$SIZE/train5.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null

SIZE=25
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test1.csv -train datasets/$SIZE/train1.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null
echo  >> evals/$EVAL/$SIZE.txt
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test2.csv -train datasets/$SIZE/train2.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null
echo  >> evals/$EVAL/$SIZE.txt
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test3.csv -train datasets/$SIZE/train3.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null
echo  >> evals/$EVAL/$SIZE.txt
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test4.csv -train datasets/$SIZE/train4.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null
echo  >> evals/$EVAL/$SIZE.txt
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test5.csv -train datasets/$SIZE/train5.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null

SIZE=30
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test1.csv -train datasets/$SIZE/train1.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null
echo  >> evals/$EVAL/$SIZE.txt
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test2.csv -train datasets/$SIZE/train2.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null
echo  >> evals/$EVAL/$SIZE.txt
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test3.csv -train datasets/$SIZE/train3.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null
echo  >> evals/$EVAL/$SIZE.txt
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test4.csv -train datasets/$SIZE/train4.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null
echo  >> evals/$EVAL/$SIZE.txt
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test5.csv -train datasets/$SIZE/train5.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null

SIZE=35
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test1.csv -train datasets/$SIZE/train1.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null
echo  >> evals/$EVAL/$SIZE.txt
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test2.csv -train datasets/$SIZE/train2.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null
echo  >> evals/$EVAL/$SIZE.txt
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test3.csv -train datasets/$SIZE/train3.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null
echo  >> evals/$EVAL/$SIZE.txt
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test4.csv -train datasets/$SIZE/train4.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null
echo  >> evals/$EVAL/$SIZE.txt
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test5.csv -train datasets/$SIZE/train5.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null

SIZE=40
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test1.csv -train datasets/$SIZE/train1.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null
echo  >> evals/$EVAL/$SIZE.txt
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test2.csv -train datasets/$SIZE/train2.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null
echo  >> evals/$EVAL/$SIZE.txt
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test3.csv -train datasets/$SIZE/train3.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null
echo  >> evals/$EVAL/$SIZE.txt
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test4.csv -train datasets/$SIZE/train4.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null
echo  >> evals/$EVAL/$SIZE.txt
go run cmd/confusionmatrix/confusion_matrix.go -pred-attr Survived -test datasets/$SIZE/test5.csv -train datasets/$SIZE/train5.csv >> evals/$EVAL/$SIZE.txt 2> /dev/null