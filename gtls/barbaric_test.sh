#!/bin/bash


# the -w "\n%{http_code}\n" simple appends the resultant http status code to the output

echo ""
echo "should give us a 400, bad request"
curl -w "\n%{http_code}\n" "http://127.0.0.1:8080"

echo ""
echo ""
echo "should give us hello world"
curl -w "\n%{http_code}\n" "http://127.0.0.1:8080/admin/"

echo ""
echo ""
echo "should give us hello world"
curl -w "\n%{http_code}\n" "http://127.0.0.1:8080/admin/hello"

