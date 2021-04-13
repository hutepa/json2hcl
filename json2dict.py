#!/usr/bin/env python

from argparse import ArgumentParser
import json

parser = ArgumentParser()
parser.add_argument('filename', type=str, metavar="FILE")

args = parser.parse_args()

filename = args.filename

with open(filename, 'r') as fname:
    data = json.load(fname)        

print(data)
