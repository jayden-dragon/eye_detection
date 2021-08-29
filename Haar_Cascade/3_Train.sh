#!/bin/bash


# make multi sample
for file in output_samples/*.dat
do
    positive_num="ls positive | wc -l"

    if [ -f $file ]; then
        opencv_createsamples -info ${file} -vec ${file}.vec -w 24 -h 24 -num ${positive_num}
    fi
done

python mergevec.py -v output_samples/ -o output_samples/phase2.vec

