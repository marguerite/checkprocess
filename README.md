systemd user service to check if a specific program is ready by checking a process

## Usage

    go build checkprocess.go
    sudo cp -r checkprocess /usr/bin
    cp -r checkprocess@.service ~/.config/systemd/user/
    systemctl enable checkprocess@kdeinit5.service
    systemctl start checkprocess@kdeinit5.service

