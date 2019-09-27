import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'package:provider/provider.dart';
import 'package:flutter_spinkit/flutter_spinkit.dart';

import 'main.dart';
import 'rootModel.dart';

class RootPage extends StatelessWidget {
  final AppData data;

  RootPage({Key key, this.data}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    var rootModel = Provider.of<RootModel>(context);

    switch (rootModel.getStatus()) {
      case Status.disconnected:
        return InitPage(data: data);
        break;
      case Status.connected:
        return ConnectedPage(data: data);
        break;
      case Status.connecting:
        return DialPage(data: data);
        break;
      case Status.init:
      default:
        return InitPage(data: data);
    }
  }
}

class ConnectedPage extends StatefulWidget {
  final AppData data;

  ConnectedPage({Key key, this.data}) : super(key: key);

  @override
  State<StatefulWidget> createState() {
    return ConnectedPageState();
  }
}

class BoxSize {
  double height;
  double width;

  BoxSize({this.height, this.width});
}

class ConnectedPageState extends State<ConnectedPage> {
  var rootModel;
  Timer _timer;

  BoxSize _mxSize; // the size of the entire box
  BoxSize _fbSize; // the size of the feedback box
  BoxSize _ctSize; // the size of the controller box

  Color _blinkColor;
  // Map<String, bool> _buttonState;
  List<ButtonData> _buttonDataList;

  @override
  void initState() {
    _timer = Timer.periodic(Duration(milliseconds: 500), timerCallback);

    _mxSize = BoxSize(height: 0, width: 0);
    _fbSize = BoxSize(height: 0, width: 0);
    _ctSize = BoxSize(height: 0, width: 0);
    _blinkColor = Colors.black;
    // _buttonState = Map<String, bool>();
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
        .add(ButtonData("none/none", 110, 110, false, this.buttonCallback));
    _buttonDataList
        .add(ButtonData("none/right", 200, 110, false, this.buttonCallback));

    _buttonDataList
        .add(ButtonData("backward.left", 20, 200, false, this.buttonCallback));
    _buttonDataList
        .add(ButtonData("backward/none", 110, 200, false, this.buttonCallback));
    _buttonDataList.add(
        ButtonData("backward/right", 200, 200, false, this.buttonCallback));

    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    rootModel = Provider.of<RootModel>(context);

    return Center(
      child: LayoutBuilder(builder: (context, constraint) {
        _mxSize.height = constraint.maxWidth;
        _mxSize.width = constraint.maxHeight;
        // print(maxW.toString() + " / " + maxH.toString());
        return getInterface();
      }),
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

    rootModel.pressHandler.send(Request.send, pressedButtonName);

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
              mainAxisAlignment: MainAxisAlignment.start,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: <Widget>[
                _getTitleInterface(),
                _getControllerInterface(),
              ]));
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
          color: Colors.grey[200],
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
    _blinkColor = Colors.black;
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

class DialPage extends StatefulWidget {
  final AppData data;

  DialPage({Key key, this.data}) : super(key: key);

  @override
  State<StatefulWidget> createState() {
    return DialPageState();
  }
}

class DialPageState extends State<DialPage> {
  Timer _timer;
  bool _show;

  @override
  void initState() {
    _timer = Timer.periodic(Duration(milliseconds: 100), timerHandler);
    _show = false;
    super.initState();
  }

  @override
  void dispose() {
    _timer.cancel();
    super.dispose();
  }

  void timerHandler(Timer timer) async {
    _show = true;
    setState(() {});
  }

  @override
  Widget build(BuildContext context) {
    var rootModel = Provider.of<RootModel>(context);

    if (_show) {
      return Scaffold(
        appBar: PreferredSize(
          preferredSize: Size.fromHeight(0),
          child: AppBar(),
        ),
        body: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              Container(
                height: 60,
              ),
              Container(
                decoration: BoxDecoration(
                  color: Theme.of(context).primaryColor,
                ),
                child: SpinKitFoldingCube(
                  color: Colors.grey[300],
                  size: 80.0,
                ),
              ),
              Container(
                height: 100,
              ),
            ],
          ),
        ),
      );
    } else {
      return Scaffold(
        appBar: PreferredSize(
          preferredSize: Size.fromHeight(0),
          child: AppBar(),
        ),
        body: Center(
          child: Container(
            decoration: BoxDecoration(
              color: Theme.of(context).primaryColor,
            ),
            child: Center(),
          ),
        ),
      );
    }
  }
}

class InitPage extends StatefulWidget {
  final AppData data;

  InitPage({Key key, this.data}) : super(key: key);

  @override
  State<StatefulWidget> createState() {
    return InitPageState();
  }
}

class InitPageState extends State<InitPage> {
  String mainText;

  @override
  void initState() {
    mainText = "Host IP";
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    var rootModel = Provider.of<RootModel>(context);

    return Scaffold(
      // Give 0 to remove appBar from the actual screen
      appBar: PreferredSize(
        preferredSize: Size.fromHeight(0),
        child: AppBar(),
      ),
      body: Center(
        child: Container(
          decoration: BoxDecoration(
            color: Theme.of(context).primaryColor,
          ),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: <Widget>[
              Container(
                margin: EdgeInsets.all(40),
                child: SizedBox(
                  width: 150,
                  height: 60,
                  child: Text(
                    mainText,
                    style: TextStyle(fontSize: 44, fontWeight: FontWeight.bold),
                    textAlign: TextAlign.center,
                  ),
                ),
              ),
              Container(
                width: 300,
                height: 64,
                margin: EdgeInsets.all(10),
                child: TextFormField(
                  textAlign: TextAlign.center,
                  keyboardType: TextInputType.number,
                  initialValue: "",
                  style: TextStyle(
                      fontSize: 24,
                      color: Colors.black,
                      fontFamily: widget.data.fontFamily),
                  strutStyle: StrutStyle.fromTextStyle(TextStyle(height: 0.1)),
                  decoration: InputDecoration(
                    filled: true,
                    fillColor: Colors.grey[100],
                    border: OutlineInputBorder(
                        borderSide: BorderSide(color: Colors.white),
                        borderRadius: BorderRadius.circular(5)),
                    focusedBorder: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(5)),
                    enabledBorder: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(5)),
                  ),
                  onFieldSubmitted: (address) {
                    rootModel.pressHandler(Request.dial, address);
                  },
                ),
              ),
              Container(
                height: 100,
              )
            ],
          ),
        ),
      ),
    );
  }
}
