[![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/EricNeid/go-weatherstation) 

# go-WeatherStation -- go away bad weather

Tool to display your pictures, as well as the current weather information. Intended to be run on your local raspberry pi.

## Dependencies

Application uses the very nice ui framework fyne.
Check out their git repository: <https://github.com/fyne-io/fyne>

Requires:

* go 1.12 or later
* c compiler
* For Debian/Ubuntu: libegl1-mesa-dev and xorg-dev

## Weather data

Weather data is retrieved from OpenWeatherMap: <https://openweathermap.org/>
To use the application, grep yourself a free key and put it in a file named
**api.key**.

## Getting started

```bash
cd cmd/weatherstation

weatherstation.exe
weatherstation.exe -h
```
