# GoFiber Mongo Boilerplate

<p align="center">
  <a href="https://golang.org/doc/go1.16">
    <img src="https://img.shields.io/badge/Go-1.16+-00ADD8?style=flat&logo=go">
  </a>
  <a href="https://github.com/gofiber/fiber/releases">
    <img src="https://img.shields.io/github/v/release/gofiber/fiber?color=00ADD8&label=%F0%9F%9A%80%20">
  </a>
  <a href="https://jwt.io/">
    <img src="https://img.shields.io/badge/JWT-black?style=flat&logo=JSON%20web%20tokens">
  </a>
  <a href="https://www.mongodb.com">
    <img src="https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=flat&logo=mongodb&logoColor=white">
  </a>
  <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/License-MIT-green.svg?style=flat">
  </a>
</p>

Is a golang based boilerplate application with Fiber Go web framework.
For any fiber go application, just clone the repo & rename the application name.

[Fiber](https://gofiber.io/) is a Go web framework built on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for **fast** development with **zero memory allocation** and **performance** in mind. 

This boilerplate uses [MongoDB](https://www.mongodb.com/) as database service. And for authentication this boilerplate uses [JWT](https://jwt.io/) where it has Access Token and Refresh Token schema.

For better coding experience, this boilerplate using [Air](https://github.com/cosmtrek/air) to handle hot reloading. If you already installed [Air](https://github.com/cosmtrek/air), you can simply run it by command,
```bash
  air
```

If not you can run, ```go run main.go```, but i don't recomend it because it sucks to re-run everytime you just need to print something 😜.

# For production use
It is recommended to use docker compose. Just simply run:
```bash
  docker compose up
```

## 🚧 WORK IN PROGRESS

- I will update(try) this regularly to add functionality and new features.

**Used libraries:**

- [jwt-go](https://github.com/dgrijalva/jwt-go) v3.2.0+incompatible
- [godotenv](https://github.com/joho/godotenv) v1.5.1
- [uuid](https://github.com/google/uuid) v1.6.0
- [Go Mongo Driver](https://go.mongodb.org/mongo-driver) v1.14.0
- [Go Playground Validator](https://github.com/go-playground/validator) v9.31.0+incompatible
- [goccy/go-json](https://github.com/goccy/go-json) 
