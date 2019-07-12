### Configuration

`.yaml` file contains the main part of the configuration. 

### Environment

| name  | default  | description  |
| :------------: | :------------: | :------------: |
| `TOKEN`  |  required | telegram bot token  |
| `CONFIG_PATH`  | `config.yaml`  | path to config  |
| `URL`  | `192.168.88.192`  | sensor url (esp8266-mh-z19)  |
| `GROUP_ID`  | none  | will be added to `config.Subscribers`  |

### Install With Docker
* copy provided `docker-compose.yml` and customize for your needs
* compile the sources and run - `docker-compose build && docker-compose up -d`
* or just run the latest version from dockerhub - `docker-compose pull && docker-compose up -d`

### Usage

`buildGraph(mode string)` 

| telegram command | mode       |
| :-------------   | :--------- |
|  `/graph`        | any string |
|  `/graph_temp`   | `"temp"`   | 
|  `/graph_co2`    | `"co2"`    |
|  `/graph_hum`    | `"hum"`    |