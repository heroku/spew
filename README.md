Used in [log-shuttle](http://github.com/heroku/log-shuttle) integration testing

## Usage

```bash
heroku create -b https://github.com/kr/heroku-buildpack-go.git <name>
heroku config:set DURATION=100ms # default
git push heroku master
```
