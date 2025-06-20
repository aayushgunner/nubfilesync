# 🗂️ FileShare - File Synchronization Tool in Go

FileShare is a lightweight command-line tool written in Go that allows you to **synchronize a text file (`new.txt`) between machines** using a simple HTTP server. It supports basic push and pull operations over IPv6 and can be easily extended.

---

## 🔧 Features

- 🖥️ Run as a simple HTTP server to serve or receive files
- ⬇️ Pull a file from a remote server using its IPv6 address
- ⬆️ Push a file to a remote server over HTTP
- 📄 View file metadata using `/status` endpoint
- 🌐 IPv6 compatible

---

## 📁 File in Use

- The tool operates on a single file named `new.txt`.
- If the file doesn't exist, it will be created automatically when the server starts.

---

## 🚀 Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.18+ installed

### Run the tool

```bash
go run fileshare.go <command> [<IPv6_address>]
```

---

## 📦 Commands

### 🔹 Run as Server

```bash
go run fileshare.go server
```

- Starts a server on port `8000`
- Serves `new.txt` at the `/file` endpoint
- Shows file status at `/status`

---

### 🔹 Pull a File from Server

```bash
go run fileshare.go pull <IPv6_address>
```

- Downloads `new.txt` from the server and saves it locally

---

### 🔹 Push a File to Server

```bash
go run fileshare.go push <IPv6_address>
```

- Uploads the local `new.txt` to the server

---

## 🌐 Example (with IPv6)

Make sure the IPv6 address is enclosed in square brackets:

```bash
go run fileshare.go pull [fe80::1%eth0]
```

> ⚠️ Replace `fe80::1%eth0` with your actual server IPv6 address (with interface scope if link-local).

---
