# 🎸 Groupie Tracker

A simple web project that consumes the **Groupie Trackers API** to display information about various bands, their locations, and concert dates. The project features:

* A backend written in **Go**
* A responsive frontend with **plain HTML and CSS**
* Visual styling inspired by the anime **Bocchi the Rock!**
* Error handling for **404**, **500**, and **405** HTTP status codes

---

## 📦 Features

* View artist and band information fetched from the [Groupie Trackers API](https://groupietrackers.herokuapp.com/api).
* Mobile-friendly and responsive design.
* Anime-themed visual style based on *Bocchi the Rock!*.
* Server-side error handling:

  * **404** – Page not found
  * **500** – Internal server error
  * **405** – Method not allowed

---

## 🚀 Getting Started

### Prerequisites

Make sure you have:

* [Go](https://golang.org/dl/) installed (version 1.18 or later recommended)

### Running the Project

1. Open your terminal.
2. Navigate to the project root directory.
3. Run:

```bash
go run ./cmd
```

4. The server should start immediately. By default, it listens on `localhost` and the port `:8080`.

---

## ⚠️ Troubleshooting

* If any errors occur during startup, they will be printed directly to the terminal.
* Make sure all required files are present under `./frontend`.
* Ensure you have a stable internet connection to access the Groupie Trackers API.
