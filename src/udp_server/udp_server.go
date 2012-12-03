package udp_server

import (
  "net"
  )

// address is a host:port string, response is a chan string
func CreateServer(address string, response chan string) {
  laddr, err := net.ResolveUDPAddr("udp", address)
  if err != nil {
    panic(err)
  }

  conn, err := net.ListenUDP("udp", laddr)
  if err != nil {
    panic(err)
  }

  go func(conn *net.UDPConn, c chan string) {
    for {
      buf := make([]byte, 100)

      _, err := conn.Read(buf)
      if err != nil {
        panic(err)
      }

      c <- string(buf)
    }
  }(conn, response)
}
