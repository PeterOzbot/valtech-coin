package transactions

// private/public key used for testing
var testPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`
var testPublicKey = `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAhAeIc5VZRD0Ef0u9KtyL
JtMujPGL4v8kyDlZdQ4tnhQLaU7yaQREQvckzF2bOtj0pY7oWYQopnW+jpSGsif6
ckNk6A6Lcw+wUY2JfuFsOnb3NAm7S48ZIg0owryZdk1NvV2LiApXNALRd5EZwhp4
aVrFRFVv2g+hgmHaJBuzujAyc7QHnS7/WUAKSGfQ3ItcF1fNMEBa2007kUIx8ka5
PmRofiIdb5Ld89LpaXbGDUseW+oCZwf3wbThOBBub02EDlLIeo6Gi0UCAXjGtvT1
nyMXlvUUmojs5em4dzH+DMptG0GzSxvP5RJUtKzdnm/8l6sBANNjzmWFRDp+uCZv
TJEaWmBXX8hl9QS7sp9UF9N9pwVLdihYw7FMp0q1HnvvQ3L7Ed4fcRWQLvPTnsFg
qgCTXgt0ZYJGk4tC3fmqmfjCxvvPL+uNEpiZKJFjyLAMwQDNLvKQpmcgHDgIw/qY
qinisLSUgPQgqY9jx9NM52L6spLffVKsb0Khg0LTIiWTso7pAXKBIFt1RmdPwTTH
RGmG0lcehULnBU3gxWSEDnBEkhqjVnq2T0vGOfpc4JOWhKDNa7O5tgUJxHnaXZBY
dHnupKnClnc+GTZdfPh9dpRQN995ayhGDbbQlOM2ct7eS1JFou8nZgCDAp1FVULc
GPbm8Zuw6sDa2QTOPFFYdNECAwEAAQ==
-----END PUBLIC KEY-----`
