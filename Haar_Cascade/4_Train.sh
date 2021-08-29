#!/bin/bash#!/bin/bash

# make output_xml
rm -rf output_xml/*

cp output_vec/phase1.vec output_xml/
cp output_samples/phase2.vec output_xml/
python mergevec.py -v output_xml/ -o output_xml/phase3.vec

opencv_traincascade -data output_xml -vec output_xml/phase3.vec -bg haar_negative.txt -numPos 1010 -numNeg 10000 -numStages 20 -featureType LBP -mode ALL -w 24 -h 24

cp output_xml/cascade.xml .
