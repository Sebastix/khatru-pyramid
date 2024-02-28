# Khatru Invite

A relay based on [Khatru](https://github.com/fiatjaf/khatru) with a invite hierarchy feature.

### Deploy with docker

```
docker run \
    -p 3334:3334 \
    -v ./users.json:/app/users.json \
    -v ./db:/app/db \
    -e DOMAIN=yourdomain.example.com \
    -e RELAY_NAME="your relay name" \
    -e RELAY_PUBKEY="your nostr hex pubkey" \
    tijlxyz/khatru-pyramid:latest
```

