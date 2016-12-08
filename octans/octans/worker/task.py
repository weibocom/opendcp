# !/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: WhiteBlue
# Time  : 2016/07/20


class Task:
    """
    Base "task" Class
    """

    def __init__(self, id):
        self.id = id

    def get_id(self):
        return self.id

    def before(self):
        pass

    def run(self):
        pass

    def success(self, result):
        pass

    def failed(self, error):
        pass

    def final(self):
        pass
