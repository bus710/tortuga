import 'dart:async';
import 'dart:convert'; // for JSON/struct conversion

import 'package:flutter/material.dart';
import 'package:universal_html/html.dart' as html;

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
  html.WebSocket _socket;
  String _host;
  String _buttonName;

  RootModel() {
    _timer = Timer.periodic(Duration(milliseconds: 1000), timerHandler);
    _state = Status.init;
    _buttonName = "none/none";
  }

  getStatus() => _state;
  getSocketStatus() => _socket.readyState;

  void timerHandler(Timer timer) async {
    send();
  }

  void pressHandler(Request request, String param) {
    debugPrint(_state.toString() + " / " + request.toString() + " / " + param);
    debugPrint(">>> " + _socket.readyState.toString());
    switch (_state) {
      case Status.init:
        if (request == Request.dial) {
          // TODO: check the format of IP address
          _host = param;
          socketInit();
          _state = Status.connecting;
        }
        break;
      case Status.connecting:
        break;
      case Status.connected:
        if (request == Request.send) {
          // TODO: check the format of button name
          _buttonName = param;
          send();
        }
        break;
      case Status.disconnected:
        _buttonName = "none/none";
        send();
        _socket.close();
        _state = Status.init;
        break;
      default:
        break;
    }
    notifyListeners();
  }

  void socketInit() {

    if (_socket.readyState == html.WebSocket.OPEN) {
      _socket.close();
    }

    debugPrint(">>> 1");

    _socket = html.WebSocket('ws://' + _host + ':3000/message');

    debugPrint(">>> 2");

    _socket.onOpen.listen((e) async {
      debugPrint("websocket: opened");
      await Future.delayed(Duration(seconds: 2));
      _state = Status.connected;
      notifyListeners();
    });

    _socket.onClose.listen((e) {
      debugPrint("websocket: closed");
      _state = Status.disconnected;
      notifyListeners();
    });

    _socket.onMessage.listen((e) {
      // Do nothing
    });
  }

  void send() {
    if (_state == Status.connected &&
        _socket != null &&
        _socket.readyState == html.WebSocket.OPEN) {
      _socket.send(json.encode({
        "ButtonName": _buttonName,
      }));
    }
  }
}
