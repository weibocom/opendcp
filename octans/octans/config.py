# !/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: WhiteBlue
# Time  : 2016/07/20

import yaml

'''
Load config items in config.yml
'''
class Config:
    def __init__(self):
        self._params = {
            "mysql": "",
            "get_key_url": "",
            "pool_size": "",
            "pool_recycle": ""
        }

# read config
    def get(self, key):
        return self._params[key]

# Load config file   
# param    config_file: config file with path
    def load(self, config_file):
        with open(config_file) as f:
            params = yaml.safe_load(f)

        for key in self._params.keys():
            if key not in params:
                raise Exception("param '{}' not found in config file")
            self._params[key] = params[key]
