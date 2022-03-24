# Youtube-dl-Server

---

## Get Started


## Config

```yaml
firebase:
  token_path: "./firebase_token.json"
ngrok:
  port: "8444"
  token: "Your Ngrok Token"
youtube_dl:
  audio_format: "mp3"
  audio_quality: 5

  # Api Url created like : http://localhost:{ngrok.port}/{api.version}/{api_name}
api:
  version: "v1"
  config_api: "/config"
  audio_api: "/audio/{videoID}"
  log_api: "/logger/{page}"
  melon_api: "/melon"
view:
  path: "./view/web/"
```