#!/usr/bin/python
from __future__ import print_function

import json
import os
import subprocess

REDASH_API_KEY = os.environ.get('REDASH_API_KEY')
REDASH_URL = os.environ.get('REDASH_URL')
BACKUP_DIR = os.path.join(os.path.dirname(os.path.abspath(__file__)), 'backup')
PAGE_SIZE = 50

page = 1

if not os.path.exists(BACKUP_DIR):
    os.mkdir(BACKUP_DIR)

while True:
    cmd = 'redashman query list %d %d --api-key=%s --url=%s --json' % (PAGE_SIZE, page, REDASH_API_KEY, REDASH_URL)
    stdout, stderr = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, shell=True).communicate()
    if stderr:
        break

    queries = json.loads(stdout)
    for query in queries['results']:
        backup_file_path = '%s/%04d.sql' % (BACKUP_DIR, query['id'])
        with open(backup_file_path, 'w') as f:
            f.write(query['query'])

    page += 1
