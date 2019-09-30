/* This page shows the buttons with some dynamics */

import 'dart:async';

import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../utils/appData.dart';
import '../utils/boxSize.dart';
import '../utils/enums.dart';
import '../rootModel.dart';

class ConnectedPage extends StatefulWidget {
  final AppData data;

  ConnectedPage({Key key, this.data}) : super(key: key);

  @override
  State<StatefulWidget> createState() {
    return ConnectedPageState();
  }
}

class ConnectedPageState extends State<ConnectedPage> {
  var rootModel;
  Timer _timer;

  BoxSize _mxSize; // the size of the entire box
  BoxSize _fbSize; // the size of the feedback box
  BoxSize _ctSize; // the size of the controller box

  Color _blinkColor;
  List<ButtonData> _buttonDataList;

  @override
  void initState() {
    _timer = Timer.periodic(Duration(milliseconds: 500), timerCallback);

    _mxSize = BoxSize(height: 0, width: 0);
    _fbSize = BoxSize(height: 0, width: 0);
    _ctSize = BoxSize(height: 0, width: 0);
    _buttonDataList = List<ButtonData>();

    _buttonDataList
        .add(ButtonData("forward/left", 20, 20, false, this.buttonCallback));
    _buttonDataList
        .add(ButtonData("forward/none", 110, 20, false, this.buttonCallback));
    _buttonDataList
        .add(ButtonData("forward/right", 200, 20, false, this.buttonCallback));

    _buttonDataList
        .add(ButtonData("none/left", 20, 110, false, this.buttonCallback));
    _buttonDataList
        .add(ButtonData("none/none", 110, 110, true, this.buttonCallback));
    _buttonDataList
        .add(ButtonData("none/right", 200, 110, false, this.buttonCallback));

    _buttonDataList
        .add(ButtonData("backward/left", 20, 200, false, this.buttonCallback));
    _buttonDataList
        .add(ButtonData("backward/none", 110, 200, false, this.buttonCallback));
    _buttonDataList.add(
        ButtonData("backward/right", 200, 200, false, this.buttonCallback));

    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    rootModel = Provider.of<RootModel>(context);
    return Scaffold(
      appBar: PreferredSize(
        preferredSize: Size.fromHeight(0),
        child: AppBar(),
      ),
      body: Center(
        child: LayoutBuilder(
          builder: (context, constraint) {
            _mxSize.height = constraint.maxWidth;
            _mxSize.width = constraint.maxHeight;
            return getInterface();
          },
        ),
      ),
    );
  }

  void buttonCallback(String pressedButtonName) {
    // To the button indicate the color immediately
    _timer.cancel();
    _timer = Timer.periodic(Duration(milliseconds: 500), timerCallback);

    // Only the button pressed last time should be set
    // so that this handler clears everyone and sets the last one.
    _buttonDataList.forEach((b) => {
          b.clearButtonState(),
          b.setBlinkState(),
          if (b._name == pressedButtonName) {b.setButtonState()}
        });

    rootModel.pressHandler(Request.send, pressedButtonName);

    setState(() {});
  }

  void timerCallback(Timer timer) {
    // To blink the button pressed
    _buttonDataList.forEach((b) => {
          b.flipBlinkState(),
        });
    setState(() {});
  }

  Widget getInterface() {
    if (_mxSize.width > 319) {
      _fbSize.width = _ctSize.height = _ctSize.width = 319;

      return Container(
        decoration: BoxDecoration(
          color: Color(widget.data.primaryColor),
        ),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: <Widget>[
            _getTitleInterface(),
            Container(
              height: 64,
            ),
            _getControllerInterface(),
          ],
        ),
      );
    } else {
      return Text("The space available is too small");
    }
  }

  Widget _getTitleInterface() {
    double targetH;

    if (_mxSize.height < 700) {
      targetH = _mxSize.height - _ctSize.height - 30;
    } else {
      targetH = 340;
    }

    return Container(
      width: _mxSize.width - 10,
      height: targetH,
      margin: EdgeInsets.only(left: 10, top: 20, right: 10, bottom: 10),
      decoration: BoxDecoration(
        color: Color(widget.data.primaryColor),
      ),
      child: Center(
          child: Text(
        "TORTUGA",
        style: TextStyle(
          fontSize: 48,
          fontWeight: FontWeight.w100,
          color: Colors.grey[100],
        ),
      )),
    );
  }

  Widget _getControllerInterface() {
    return Container(
      width: _ctSize.width - 30,
      height: _ctSize.height - 30,
      margin: EdgeInsets.all(5),
      decoration: BoxDecoration(
        color: Color(widget.data.primaryColor),
        /* boundary for debugging purpose */
        // border: Border.all(color: Colors.grey[900], width: 3),
        // borderRadius: BorderRadius.all(Radius.circular(5)),
      ),
      child: Stack(
        children: List.generate(_buttonDataList.length, (index) {
          return getButton(index);
        }),
      ),
    );
  }

  // Widget _getButton(String name, double x, double y, bool state) {
  Widget getButton(int index) {
    return _buttonDataList[index].getButton();
  }
}

class ButtonData {
  String _name;
  double _x, _y;
  bool _state;
  bool _blinkState;
  Color _blinkColor;
  Function _callback;

  ButtonData(String name, double x, double y, bool state, Function callback) {
    _name = name;
    _x = x;
    _y = y;
    _state = state;
    _blinkState = false;
    _blinkColor = Colors.green;
    _callback = callback;
  }

  Widget getButton() {
    Color stateColor = Colors.white;
    if (_state && _blinkState) {
      stateColor = _blinkColor;
    } else {
      stateColor = Colors.grey[200];
    }

    return Positioned(
      left: _x,
      top: _y,
      child: GestureDetector(
          child: Container(
            width: 70,
            height: 70,
            decoration: BoxDecoration(
              color: stateColor,
              border: Border.all(color: stateColor, width: 3),
              borderRadius: BorderRadius.all(Radius.circular(5)),
            ),
          ),
          onTap: () {
            if (_state) {
              _state = false;
            } else {
              _state = true;
            }
            // print("${name}");
            _callback(_name);
          }),
    );
  }

  void clearButtonState() {
    _state = false;
  }

  void setButtonState() {
    _state = true;
  }

  void flipBlinkState() {
    _blinkState = !_blinkState;
  }

  void setBlinkState() {
    _blinkState = true;
  }

  void setBlinkColor(Color color) {
    _blinkColor = color;
  }
}
