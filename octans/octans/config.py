#!/usr/bin/env python
#
#    Copyright (C) 2016 Weibo Inc.
#
#    This file is part of Opendcp.
#
#    Opendcp is free software: you can redistribute it and/or modify
#    it under the terms of the GNU General Public License as published by
#    the Free Software Foundation; version 2 of the License.
#
#    Opendcp is distributed in the hope that it will be useful,
#    but WITHOUT ANY WARRANTY; without even the implied warranty of
#    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#    GNU General Public License for more details.
#
#    You should have received a copy of the GNU General Public License
#    along with Opendcp.  if not, write to the Free Software
#    Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301  USA
#
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
