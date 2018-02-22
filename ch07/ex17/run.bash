#/bin/bash
go build ../fetch/fetch.go
go build 
./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | ./ex17 div p class="toc" 
#./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | ./ex17 div div h2
