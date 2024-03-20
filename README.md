# GoFiber Mongo Boilerplate

<p align="center">
  <a href="https://golang.org/doc/go1.16">
    <img src="https://img.shields.io/badge/Go-1.16+-00ADD8?style=flat&logo=go">
  </a>
  <a href="https://github.com/gofiber/fiber/releases">
    <img src="https://img.shields.io/github/v/release/gofiber/fiber?color=00ADD8&label=%F0%9F%9A%80%20">
  </a>
  <a href="https://jwt.io/">
    <img src="https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens">
  </a>
  <a href="https://www.mongodb.com">
    <img src="https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white">
  </a>
  <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/License-MIT-green.svg">
  </a>
</p>

Is a golang based boilerplate application with Fiber Go web framework.
For any fiber go application, just clone the repo & rename the application name.

[Fiber](https://gofiber.io/) is an Express.js inspired web framework build on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for **fast** development with **zero memory allocation** and **performance** in mind. 

This boilerplate application uses [MongoDB](https://www.mongodb.com/) as database service. And for authentication this boilerplate uses [JWT](https://jwt.io/) where it has Access Token and Refresh Token schema.

For better coding experience, this boilerplate using [Air](https://github.com/cosmtrek/air) to handle hot reloading. If you already installed [Air](https://github.com/cosmtrek/air), you can simply run it by command
```bash
  air
```