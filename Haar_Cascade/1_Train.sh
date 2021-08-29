#!/bin/bash

rm -rf output_vec/*
mkdir output_vec/positive/

positive_num="ls positive | wc -l"
positive_file=haar_positive.txt
negative_num="ls negative | wc -l"
negative_file=haar_negative.tx

# make vec file
for file in positive/*.jpg
do
    if [ -f $file ]; then
        echo "filename = ${file}"
        opencv_createsamples -img ${file} -vec output_vec/${file}.vec -bg ${negative_file} -num 100 -maxxangle 0.1 -maxyangle 0.1 -maxzangle 0.1 -maxidev 40 -bgcolor 0 -bgthresh 0 -w 24 -h 24
    fi
done


python mergevec.py -v output_vec/positive/ -o output_vec/phase1.vec



