#!/bin/bash

cd webclient

echo "> Compiling SCSS files"

sass static/sass/aigera.scss static/aigera.css

echo "> Building Web Client"

yarn run build
