# Hermetix

Hermetix is a Docker container that provides simple REST API for sending iMessages. The messages will be sent from the current user on the Docker host machine.

This is quite useful for development environments where you need a quick way to send messages via iMessage, as an alternative to using SMS.

## Limitations

- Works only on OS X.
- Messages will only be sent to recipients if there is *already an existing dialogue* with the user on the host. Otherwise, the request will be ignored.

## Setup

Make sure the host is running the SSH daemon

    sudo systemsetup -setremotelogin on

There must be a valid RSA keypair in the current user's ~/.ssh/ directory. It's assumed that the private key is named id_rsa: 

    ssh-keygen -t rsa

The current user's SSH public key must be added to ~/.ssh/authorized_keys.

    cat ~/.ssh/id_rsa.pub >> ~/.ssh/authorized_keys

## Run

To start the container:

    docker run -p 8009:8009 -e USER=`whoami` -e HOST=`ipconfig getifaddr en0` -v ~/.ssh/:/app/ssh  augustzf/hermetix

## Usage

To send a message from the host:

    curl "http://localhost:8009?msg=<message>&rec=<email or mobile number>"

Where `msg` is the message to be sent and `rec` is the recipient. `rec` can be a mobile number or email address.

To send messages from other orchestrated containers, make sure you use the right hostname.

