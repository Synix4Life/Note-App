# Note App

![Version](https://img.shields.io/badge/Version-3.0.1-%2300ADD8.svg?style=for-the-badge&)
![License](https://img.shields.io/badge/LICENSE-MIT-%2300ADD8.svg?style=for-the-badge&)

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

![CSS](https://img.shields.io/badge/CSS-563d7c?style=for-the-badge&logo=CSS&logoColor=white)
![HTML5](https://img.shields.io/badge/html-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white)
![JavaScript](https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E)
![SQLite](https://img.shields.io/badge/sqlite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white)

---

This is a small note-taking application written in Go.

It doesn't feature security mechanisms or similar, it's for personal use only.

---

### Requirements

This program uses the go module `modernc.org/sqlite`

Therefore, you need to install it first:

```shell
go get modernc.org/sqlite
go mod tidy
```

---

### Changelog

- $\textsf{\color{aqua}Version 1.0.0}$
  - Initial commit
- $\textsf{\color{aqua}Version 2.0.0}$
  - Merge: Web- GUI added
    - $\textsf{\color{aqua}Version 2.0.1}$
      - Small fixes
      - Improved outsourcing
- $\textsf{\color{aqua}Version 3.0.0}$
  - Switch to SQLite3 Database
  - Improved Frontend → Buttons reconfig + Help Button
    - $\textsf{\color{aqua}Version 3.0.1}$
      - .gitignore updated
