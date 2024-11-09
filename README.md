# Khatru Pyramid

A relay based on [Khatru](https://github.com/fiatjaf/khatru) with an invite hierarchy feature.

### Deploy with docker

```sh
docker run \
    -p 3334:3334 \
    -v ./users.json:/app/users.json \
    -v ./db:/app/db \
    -e DOMAIN="yourdomain.example.com" \
    -e RELAY_NAME="your relay name" \
    -e RELAY_PUBKEY="your nostr hex pubkey" \
    tijlxyz/khatru-pyramid:latest
```

### Using Docker compose

```shell
docker compose build
docker compose up 
```

You can now access the frontend of the relay at http://localhost:<EXPOSED_DOCKER_PORT>.

### Nginx reverse proxy config

@TODO

### Deploy with

 - [YunoHost](https://github.com/YunoHost-Apps/khatru-pyramid_ynh) ([app catalog](https://apps.yunohost.org/catalog) [pending](https://github.com/YunoHost/apps/pull/2077))
 - [Cloudron](https://github.com/github-tijlxyz/khatru-pyramid_cloudron) ([app catalog](https://www.cloudron.io/store/index.html) [pending](https://forum.cloudron.io/topic/11146/khatru-pyramid-a-nostr-relay))

### Manually build

```sh
git clone https://github.com/github-tijlxyz/khatru-pyramid && cd khatru-pyramid
just build
DOMAIN="example.com" RELAY_NAME="my relay" RELAY_PUBKEY=yourpubkey ./khatru-pyramid
```

### Configuration

Look at [example.env](./example.env) for all configuration options.

You can also manually edit the `users.json` file. Do this only when the server is down.
`users.json` is formatted as follows:
```json
{ 
  "[user_pubkey_hex]": "[invited_by_pubkey_hex]"
}
```

