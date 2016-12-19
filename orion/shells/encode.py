#!python

import json
import sys

filename = sys.argv[1]

lines = list(open(filename))
text = ''.join(lines)
encoded = json.dumps(text)
print encoded[1:len(encoded)-1]
