###################################################################################################
# Golang
###################################################################################################
# Golang tools (compiler).
export PATH=$PATH:/usr/local/go/bin

# Golang path.
export GOPATH=$HOME/code/golang
export PATH=$PATH:$GOPATH/bin

###################################################################################################
# Webserver
###################################################################################################
export WEBSERVER_SRC=$HOME/code/webserver
export WEBSERVER_DATA=~/.local/share/webserver

export WEBSERVER_HOST_DEV=http://localhost:8080
export WEBSERVER_USER_DEV=webserver
export WEBSERVER_PASS_DEV=site1q2w3e4R
export WEBSERVER_HOST_PROD=https://www.webserver.com.br
export WEBSERVER_USER_PROD=webserver
export WEBSERVER_PASS_PROD=siteq1w2e3R4

export WEBSERVER_DB=postgres://app:zunKa4a@localhost:27017/zunka?authSource=admin 
export WEBSERVER_DB_NAME=webserver
export WEBSERVER_DB_USER=webserver
export WEBSERVER_DB_NAME=webserver
