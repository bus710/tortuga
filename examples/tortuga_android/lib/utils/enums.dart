// Status can be used in the rootModel
enum Status {
  init,
  connecting,
  connected,
  disconnected,
}

// Request can be used in the pages as the param to rootModel
enum Request {
  dial,
  connect,
  send,
  disconnect,
}
