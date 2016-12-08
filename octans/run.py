# !/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: WhiteBlue
# Time  : 2016/07/20

from octans import App, start_app

if __name__ == '__main__':
    start_app()

    #App.run("0.0.0.0", 8000, debug=False)
    App.run("0.0.0.0", 8000, debug=True)
