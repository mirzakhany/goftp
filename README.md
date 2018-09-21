# goFtp

A Very Simple file sharing tools with golang, based on memory of ```python -m SimpleHTTPServer``` but
very faster and multi thread with simple auth support

## Simple install with curl on linux

    curl -L https://github.com/mirzakhany/goftp/releases/download/v0.2/goftp_linux -o goftp
    chmod +x goftp
    sudo cp goftp /usr/local/bin


## Build from source

    git clone githup.com/mirzakhany/goftp
    cd goftp
    make all
    sudo cp bin/goftp /usr/local/bin    
    

## Share a folder

in any folder you want to share run flowing command
    
    goftp serv [-d . -u username -p password -P port -i ip-address]

and open the ip address in browser (default is 0.0.0.0:9090)

for more information run ```goftp -help```
