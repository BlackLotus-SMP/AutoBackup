# AutoBackup

[![Build status](https://github.com/BlackLotus-SMP/AutoBackup/actions/workflows/build.yml/badge.svg)](https://github.com/BlackLotus-SMP/AutoBackup/actions/workflows/build.yml)

An API Based auto backup system using `rsync` and `sshpass`

**rsync and sshpass needs to be installed**

## Start

```bash
./backup -p 8555 -c config.json
```
`-p is optional`, **default port is 8462**

`-c is optional`, **default config file is config/config.json**

## Endpoints

You probably want to use cron to schedule backups!

### Create
- **method**: `GET`
- **endpoint**: `/backup/create/{name}`
- **command**: `curl 127.0.0.1:{port}/backup/create/{name}` (Read Config to understand what {name} is)
- **run**: Will create a backup using rsync and zip with tar

### List
- **method**: `GET`
- **endpoint**: `/backup/list`
- **command**: `curl 127.0.0.1:{port}/backup/list`
- **run**: Will list configs
- **response**:
```json
        {
            "code": 200,
            "data": [
                {
                    "name": "test",
                    "n_backups": 7
                }
            ]
        }
```

### Delete
- **method**: `GET`
- **endpoint**: `/backup/delete/{name}`
- **command**: `curl 127.0.0.1:{port}/backup/delete/{name}` (Read Config to understand what {name} is)
- **run**: Will delete the json object from the config file

### Reload
- **method**: `GET`
- **endpoint**: `/reload`
- **command**: `curl 127.0.0.1:{port}/reload`
- **run**: Will reload the config file so you don't need to restart the process on config modify

### HealthCheck
- **method**: `GET`
- **endpoint**: `/healthcheck`
- **command**: `curl 127.0.0.1:{port}/healthcheck`
- **run**: Just returns 200, this is for docker/kubernetes integration

## Config
The config file needs to be in the config directory, you can find an example on [config/config.json](https://github.com/BlackLotus-SMP/AutoBackup/blob/master/config/config.json)

You can create as many json objects as you want

```json
[
  {
    "name": "test",
    "ssh_remote_path": "1.2.3.4:/home/test/bck/",
    "ssh_user": "user",
    "ssh_pass": "pass",
    "local_path": "/home/bck/",
    "n_backups": 5
  },
  {
    "name": "server",
    "ssh_remote_path": "1.2.3.4:/home/test/smp/",
    "ssh_user": "user",
    "ssh_pass": "pass",
    "local_path": "/home/smp-server/",
    "n_backups": 2
  }
]
```

- **name**: name of the endpoint on /backup/{create/delete}/{name} **must be unique**
- **ssh_remote_path**: path of the dir/file we want to make a backup of
- **ssh_user**: ssh user
- **ssh_pass**: ssh password
- **local_path**: path we want to copy the data to
- **n_backups**: history of backups we want to keep

**Whenever a backup is made, the script will zip it with tar and set a name based on the name and date**

**If the n_backups history is reached, the oldest tar will be deleted, it will only keep the last n backups**

---

# TODO

- [ ] check if sshpass and rsync are installed
- [ ] login with file (sshpass only now)
- [x] set config file as arg on run (./build -c config.json)
- [x] fix build.sh
- [x] probably workflow to upload artifacts as release
- [ ] endpoint to register/delete a config
- [x] unique config name validation
- [ ] tests
- [ ] capture error msg from rsync if happens
- [x] refactor config logic
- [ ] destination of the zip file (defaults to the folder where the command has been run)