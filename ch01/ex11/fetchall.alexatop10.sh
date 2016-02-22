#!/bin/bash

set -eux

go run main.go \
   http://www.alexa.com/siteinfo/google.com \
   http://www.alexa.com/siteinfo/facebook.com \
   http://www.alexa.com/siteinfo/youtube.com \
   http://www.alexa.com/siteinfo/baidu.com \
   http://www.alexa.com/siteinfo/yahoo.com \
   http://www.alexa.com/siteinfo/amazon.com \
   http://www.alexa.com/siteinfo/wikipedia.org \
   http://www.alexa.com/siteinfo/qq.com \
   http://www.alexa.com/siteinfo/google.co.in \
   http://www.alexa.com/siteinfo/twitter.com
