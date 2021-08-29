#!/bin/bash

# make sample file
rm -rf output_samples/*
mkdir output_samples/positive/

positive_num="ls positive | wc -l"
positive_file=haar_positive.txt
negative_num="ls negative | wc -l"
negative_file=haar_negative.txt

for file in positive/*.jpg
do
    if [ -f $file ]; then
        opencv_createsamples -info ${file}.dat -img ${file} -bg ${negative_file} -num ${negative_num} -maxxangle 0.1 -maxyangle 0.1 -maxzangle 0.1 -maxidev 40 -bgcolor 0 -bgthresh 0
    
        if [ $? -eq 0 ]; then
            echo "OK............"
            mv positive/00* output_samples/
            mv positive/*dat output_samples/
        else
            echo "FAIL.........."
            rm -f positive/00*
            rm -f positive/*dat
        fi
    fi
done


