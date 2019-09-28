import 'dart:async';
import 'dart:convert'; // for JSON/struct conversion

import 'package:flutter/material.dart';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:web_socket_channel/io.dart';
import 'package:web_socket_channel/status.dart' as status;

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
          // Checking for the input string's format.
          if (param.length == 0) {
            // empty or not
            return;
          } else if (param.split('.').length != 4) {
            // normal ipv4 string?
            return;
          } else {
            // each cell has numerics only?
            bool anyError = false;
            param.split('.').forEach(
                  (p) => {
                    if (double.tryParse(p) == null)
                      {
                        anyError = true,
                      }
                  },
                );
            if (anyError) {
              return;
            }
          }

          _host = param;
          _state = Status.connecting;
          notifyListeners();
          await Future.delayed(Duration(seconds: 3));
          socketInit();
        }
        break;
      case Status.connecting:
        break;
      case Status.connected:
        if (request == Request.send) {
          // Checking the format of button name
          if (param.split('/').length != 2) {
            return;
          }
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
        notifyListeners();
        break;
      default:
        break;
    }
  }

  void socketInit() {
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
    notifyListeners();
  }

  void send() {
    if (_state == Status.connected && _ws != null) {
      var tmp = json.encode({"ButtonName": _buttonName});
      _ws.sink.add(tmp);
    }
  }
}
