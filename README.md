# IoT server

### HTTP REST API Specification
```
POST /api/auth/login
{
    "email": "",
    "password": ""
}
```
```
POST /api/auth/register
{
    "email": "",
    "password": ""
}
```
```
GET /api/devices
```
```
POST /api/devices
{
    "name": ""
}
```
```
PUT /api/devices/:id
{
    "name": ""
}
```
```
DELETE /api/devices/:id
```
```
GET /api/devices/:id/readings
```

MQTT message:
```json
{
    "deviceId": 1,
    "type": "TEMPERATURE",
    "value": 0,
    "time": 1696699342
}
```