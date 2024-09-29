package main

import (
  "bufio"
  "crypto/sha1"
  "encoding/hex"
  "log"
  "net"
  "os"
  "strings"
)

var sofname = "pwnd"
var version = "0.0.0"
var serverhost = "0.0.0.0"
var serverport = "9951"
var pwnroot = "/mnt/pwned/hashes/"

func checkPwnedHash(hash string) string {
  prefix := strings.ToUpper(hash[:4])
  filePath := pwnroot + prefix + ".txt"

  file, err := os.Open(filePath)
  if err != nil {
    log.Printf("ファイル「%s」を開けられません: %v\n", filePath, err)
    return "-1"
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    parts := strings.Split(line, ":")
    if len(parts) != 2 {
      continue
    }

    if parts[0] == strings.ToUpper(hash) {
      return parts[1]
    }
  }

  if err := scanner.Err(); err != nil {
    log.Printf("ファイル「%s]を読み込まれません: %v\n", filePath, err)
  }

  return "0"
}

func handleConnection(conn net.Conn) {
  defer conn.Close()

  buf := make([]byte, 256)
  n, err := conn.Read(buf)
  if err != nil {
    log.Println("クライアントからのエラー:", err)
    return
  }
  password := strings.TrimSpace(string(buf[:n]))

  sha1Hash := sha1.New()
  sha1Hash.Write([]byte(password))
  hash := hex.EncodeToString(sha1Hash.Sum(nil))

  res := checkPwnedHash(hash)
  conn.Write([]byte(res))
}

func main() {
  listener, err := net.Listen("tcp", serverhost + ":" + serverport)
  if err != nil {
    log.Fatal(err)
  }
  defer listener.Close()

  log.Println("サーバーは " + serverhost + ":" + serverport + " で実行中・・・")

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Println("接続エラー:", err)
      continue
    }
    go handleConnection(conn)
  }
}
