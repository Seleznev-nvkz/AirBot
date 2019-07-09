### Configuration
Main part of configuration contains in yaml file. 
### Environment
| name  | default  | description  |
| :------------: | :------------: | :------------: |
| TOKEN  |  required | token of telegram bot  |
| CONFIG_PATH  | config.yaml  | path to config  |
| URL  | 192.168.88.192  | url to check sensors (esp8266-mh-z19)  |
| GROUP_ID  | none  | will be added to subscribers and saved in config  |


### Install
## With Docker
* copy provided docker-compose.yml and customize for your needs
* compile from the sources and run - docker-compose build && docker-compose up -d
* or just run latest version from dockerhub - docker-compose pull && docker-compose up -d