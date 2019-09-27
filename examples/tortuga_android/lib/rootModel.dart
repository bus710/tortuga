import 'package:flutter/material.dart';
import 'dart:async';

enum Status {
  init,
  dial,
  connected,
  disconnected,
}

class RootModel with ChangeNotifier {
  Timer _timer;
  Status _state;

  RootModel() {
    _timer = Timer.periodic(Duration(milliseconds: 300), timerHandler);
    _state = Status.init;
  }

  getStatus() => _state;

  void timerHandler(Timer timer) async {
    // notifyListeners();
  }

  void pressHandler(Status request, String param) {
    switch (_state) {
      case Status.init:
        if (request == Status.dial) {
          // TODO: check the format of IP address
          _state = Status.dial;
        }
        break;
      case Status.dial:
        break;
      case Status.connected:
        break;
      case Status.disconnected:
        break;
      default:
        break;
    }
    notifyListeners();
  }
}
