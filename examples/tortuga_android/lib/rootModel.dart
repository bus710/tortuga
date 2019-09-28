import 'dart:async';
import 'dart:convert'; // for JSON/struct conversion

import 'package:flutter/material.dart';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:web_socket_channel/io.dart';
import 'package:web_socket_channel/status.dart' as status;
import 'package:http/http.dart' as http;

enum Status {
  init,
  connecting,
  connected,
  disconnected,
}

enum Request {
  dial,
  connect,
  send,
  disconnect,
}

class RootModel with ChangeNotifier {
  Timer _timer;
  Status _state;
  IOWebSocketChannel _ws;
  String _host;
  String _buttonName;

  RootModel() {
    _timer = Timer.periodic(Duration(milliseconds: 1000), timerHandler);
    _state = Status.init;
    _buttonName = "none/none";
  }

  getStatus() => _state;

  void timerHandler(Timer timer) async {
    send();
  }

  void pressHandler(Request request, String param) async {
    switch (_state) {
      case Status.init:
        if (request == Request.dial) {
          // TODO: check the format of IP address
          _host = param;
          _state = Status.connecting;
          notifyListeners();
          await Future.delayed(Duration(seconds: 2));
          socketInit();
        }
        break;
      case Status.connecting:
        break;
      case Status.connected:
        if (request == Request.send) {
          // TODO: check the format of button name
          _buttonName = param;
          send();
        } else if (request == Request.disconnect) {
          if (_ws != null) {
            _ws.sink.close(status.goingAway);
          }
        }
        break;
      case Status.disconnected:
        _buttonName = "none/none";
        _state = Status.init;
        break;
      default:
        break;
    }
    notifyListeners();
  }

  void socketInit() {
    // https://gist.github.com/pyzenberg/4037e11627a8cac1c442183cc7cf172a

    // _ws = IOWebSocketChannel.connect('ws://echo.websocket.org',
    _ws = IOWebSocketChannel.connect(
      'ws://' + _host + ':8080/message',
      pingInterval: Duration(seconds: 2),
    );

    _state = Status.connected;
    _ws.stream.listen(this.onData, onError: onError, onDone: onDone);

    notifyListeners();
  }

  void onData(event) {
    debugPrint("received: " + event);
  }

  void onError(err) {
    debugPrint(err.runtimeType.toString());
    WebSocketChannelException ex = err;
    debugPrint(ex.message);
  }

  void onDone() {
    _state = Status.disconnected;
    _ws = null;
  }

  void send() {
    if (_state == Status.connected && _ws != null) {
      _ws.sink.add(
        json.encode({
          "ButtonName": _buttonName,
        }),
      );
    }
  }
}
