import 'package:flutter/material.dart';
import 'dart:async';

enum Status{
  init,
  dial,
  connected,
  disconnected,
}


class RootModel with ChangeNotifier {
  Timer _timer;
  int _counter;
  Status _state;

  RootModel() {
    _timer = Timer.periodic(Duration(milliseconds: 300), timerHandler);
    _counter = 0;
    _state = Status.init; 
  }

  getCounter() => _counter;
  getStatus() => _state;

  void timerHandler(Timer timer) async {
    if (_counter > 0) {
      debugPrint(">>> Timer log from the init model");
      _counter--;
      notifyListeners();
    }
  }

  void pressHandler() {
    _counter++;
    debugPrint(">>> Press log from the init model");
    notifyListeners();
  }
}
