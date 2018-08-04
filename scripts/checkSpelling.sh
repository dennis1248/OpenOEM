#!/bin/bash

# Use this script to automaticly check spelling from readmes

# WHY? 
# Because my (mjarkk) spelling is realy bad and this is a fast wahy to check my spelling

if ! [ -x "$(command -v npm)" ]; then
  echo 'Error: nodejs install: https://nodejs.org/en/download/package-manager/#debian-and-ubuntu-based-linux-distributions' >&2
  exit 1
fi

if ! [ -x "$(command -v mdspell)" ]; then
  echo 'Error: markdown spellcheck not found trying to install' >&2
  npm i markdown-spellcheck -g
fi

if ! [ -x "$(command -v mdspell)" ]; then
  echo 'can not install markdown spellcheck try to install it yourself: `npm i -g markdown-spellcheck`' >&2
  exit 1
fi

mdspell '../**/*.md' '!../**/node_modules/**/*.md'
