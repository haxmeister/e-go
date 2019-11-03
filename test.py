#!/usr/bin/env python3
import urllib.request, json

with urllib.request.urlopen("https://gitlab.com/api/v4/projects/liguros%2Fkit-fixups/releases") as url:
    data = json.loads(url.read().decode("utf-8"))
    #print(json.dumps(data, indent=4))
    #print(type(data))

for item in data:
    print("Release name:", item['name'])
    print("Tag:", item['tag_name'])
    print("Released:", item['released_at'])
    print("Release notes:\n", item['description'])
    print()
