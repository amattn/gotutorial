#!/bin/bash


# the -w "\n%{http_code}\n" simple appends the resultant http status code to the output

echo ""
echo "should give us a 200, welcome message"
curl -w "\n%{http_code}\n" "http://127.0.0.1:8080"

echo ""
echo ""
echo "should give us hello world"
curl -w "\n%{http_code}\n" "http://127.0.0.1:8080/admin/"

echo ""
echo ""
echo "should give us hello world"
curl -w "\n%{http_code}\n" "http://127.0.0.1:8080/admin/hello"

echo ""
echo "should give us a 301, golang.org"
curl -w "\n%{http_code}\n" "http://127.0.0.1:8080/a"
curl --head --silent "http://127.0.0.1:8080/a" | grep Location
echo ""
echo "should give us a 301, tour.golang.org"
curl -w "\n%{http_code}\n" "http://127.0.0.1:8080/b"
curl --head --silent "http://127.0.0.1:8080/a" | grep Location
echo ""
echo "should give us a 301, gotutorial.net"
curl -w "\n%{http_code}\n" "http://127.0.0.1:8080/c"
curl --head --silent "http://127.0.0.1:8080/a" | grep Location

echo ""
echo "should give us a 404, Not Found"
curl -w "\n%{http_code}\n" "http://127.0.0.1:8080/bogus/bogus"
curl --head --silent "http://127.0.0.1:8080/bogus/bogus" | grep Location


echo ""
