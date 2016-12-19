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

import logging
from logging.handlers import RotatingFileHandler




_formatter = logging.Formatter('[%(name)s] %(asctime)s %(filename)s[line:%(lineno)d] %(levelname)s %(message)s',
                               datefmt='%a, %d %b %Y %H:%M:%S', )

_stream_handler = logging.StreamHandler()
_file_handler = RotatingFileHandler(filename="log/channel.log", maxBytes=1024 * 1024, backupCount=5)

_stream_handler.setFormatter(_formatter)
_file_handler.setFormatter(_formatter)


class LogManager:
    Level = logging.DEBUG

    def __init__(self):
        pass

    @staticmethod
    def get_logger(name):
        logger = logging.getLogger(name)

        logger.setLevel(LogManager.Level)

        logger.addHandler(_stream_handler)
        logger.addHandler(_file_handler)

        return logger
