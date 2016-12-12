#!/usr/bin/env python
#
#  Copyright 2009-2016 Weibo, Inc.
#
#    Licensed under the Apache License, Version 2.0 (the "License");
#    you may not use this file except in compliance with the License.
#    You may obtain a copy of the License at
#
#        http://www.apache.org/licenses/LICENSE-2.0
#
#    Unless required by applicable law or agreed to in writing, software
#    distributed under the License is distributed on an "AS IS" BASIS,
#    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#    See the License for the specific language governing permissions and
#    limitations under the License.
#
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
