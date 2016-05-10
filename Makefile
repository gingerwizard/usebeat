BEATNAME=usebeat
BEAT_DIR=github.com/gingerwizard
SYSTEM_TESTS=false
TEST_ENVIRONMENT=false
ES_BEATS=./vendor/github.com/elastic/beats
GOPACKAGES=$(shell glide novendor)
PREFIX?=.

# Path to the libbeat Makefile
include $(ES_BEATS)/libbeat/scripts/Makefile