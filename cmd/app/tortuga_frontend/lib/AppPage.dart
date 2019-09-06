import 'package:flutter_web/material.dart';
import 'package:flutter_web/gestures.dart';

import 'dart:async'; // for streams and timer

import 'package:tortuga_frontend/main.dart';
import 'package:tortuga_frontend/AppBLoC.dart';
import 'package:tortuga_frontend/AppEvent.dart';

class AppPage extends StatefulWidget {
  AppPage({Key key, this.title}) : super(key: key);

  final String title;

  @override
  State<StatefulWidget> createState() {
    return _AppState();
  }
}

class _AppState extends State<AppPage> {
  final _bloc = AppBLoC();

  Timer timer;

  int primaryColor;

  // max screen size to calcurate the gesture area
  double maxW;
  double maxH;

  // actual width of the feedback area
  double feedbackH = 0;
  double feedbackW = 0;

  // actual width of the controller area as square
  double controllerH = 0;
  double controllerW = 0;

  Map<String, bool> buttonState = Map<String, bool>();
  List<ButtonData> buttonDataList = List<ButtonData>();

  @override
  void initState() {
    // Init the button data list
    // First row
    buttonDataList.add(ButtonData("0/0", 20, 20, false, this.callback));
    buttonDataList.add(ButtonData("1/0", 70, 20, false, this.callback));
    buttonDataList.add(ButtonData("2/0", 120, 20, false, this.callback));
    buttonDataList.add(ButtonData("3/0", 170, 20, false, this.callback));
    buttonDataList.add(ButtonData("4/0", 220, 20, false, this.callback));
    // Second row
    buttonDataList.add(ButtonData("0/1", 20, 70, false, this.callback));
    buttonDataList.add(ButtonData("1/1", 70, 70, false, this.callback));
    buttonDataList.add(ButtonData("2/1", 120, 70, false, this.callback));
    buttonDataList.add(ButtonData("3/1", 170, 70, false, this.callback));
    buttonDataList.add(ButtonData("4/1", 220, 70, false, this.callback));
    // Third row
    buttonDataList.add(ButtonData("0/2", 20, 120, false, this.callback));
    buttonDataList.add(ButtonData("1/2", 70, 120, false, this.callback));
    buttonDataList.add(ButtonData("2/2", 120, 120, true, this.callback));
    buttonDataList.add(ButtonData("3/2", 170, 120, false, this.callback));
    buttonDataList.add(ButtonData("4/2", 220, 120, false, this.callback));
    // Fourth row
    buttonDataList.add(ButtonData("0/3", 20, 170, false, this.callback));
    buttonDataList.add(ButtonData("1/3", 70, 170, false, this.callback));
    buttonDataList.add(ButtonData("2/3", 120, 170, false, this.callback));
    buttonDataList.add(ButtonData("3/3", 170, 170, false, this.callback));
    buttonDataList.add(ButtonData("4/3", 220, 170, false, this.callback));
    // Fifth row
    buttonDataList.add(ButtonData("0/4", 20, 220, false, this.callback));
    buttonDataList.add(ButtonData("1/4", 70, 220, false, this.callback));
    buttonDataList.add(ButtonData("2/4", 120, 220, false, this.callback));
    buttonDataList.add(ButtonData("3/4", 170, 220, false, this.callback));
    buttonDataList.add(ButtonData("4/4", 220, 220, false, this.callback));

    timer = Timer.periodic(Duration(milliseconds: 500), timerCallback);
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Center(
      child: LayoutBuilder(builder: (context, constraint) {
        maxW = constraint.maxWidth;
        maxH = constraint.maxHeight;
        // print(maxW.toString() + " / " + maxH.toString());
        return _getInterface();
      }),
    );
  }

  Widget _getInterface() {
    if (maxW > 319) {
      feedbackW = controllerH = controllerW = 319;
      return Container(
          decoration: BoxDecoration(
            color: Color(PRIMARY_COLOR),
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

    if (maxH < 700) {
      targetH = maxH - controllerW - 30;
    } else {
      targetH = 340;
    }

    return Container(
      width: maxW - 10,
      height: targetH,
      margin: EdgeInsets.only(left: 10, top: 20, right: 10, bottom: 10),
      decoration: BoxDecoration(
        color: Color(PRIMARY_COLOR),
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
      width: controllerW - 30,
      height: controllerW - 30,
      margin: EdgeInsets.all(5),
      decoration: BoxDecoration(
        color: Color(PRIMARY_COLOR),
        /* boundary for debugging purpose */
        // border: Border.all(color: Colors.grey[900], width: 3),
        // borderRadius: BorderRadius.all(Radius.circular(5)),
      ),
      child: Stack(
        children: List.generate(buttonDataList.length, (index) {
          return _getButton(index);
        }),
      ),
    );
  }

  // Widget _getButton(String name, double x, double y, bool state) {
  Widget _getButton(int index) {
    return buttonDataList[index].Get();
  }

  void callback(String pressedButtonName) {
    // Only the button pressed last time should be set
    // so that this handler clears everyone and sets the last one.
    buttonDataList.forEach((b) => {
          b.ClearState(),
          if (b.name == pressedButtonName) {b.SetState()}
        });

            _bloc.backwardSink.add(ButtonEvent(pressedButtonName));
    setState(() {});
  }

  void timerCallback(Timer timer) async {
    buttonDataList.forEach((b) => {
          b.FlipBlinkState(),
        });
    setState(() {});
  }
}

class ButtonData {
  String name;
  double x, y;
  bool state;
  bool blinkState;
  Function callback;

  ButtonData(String name, double x, double y, bool state, Function callback) {
    this.name = name;
    this.x = x;
    this.y = y;
    this.state = state;
    this.blinkState = false;
    this.callback = callback;
  }

  Widget Get() {
    Color stateColor = Colors.white;
    if (this.state && this.blinkState) {
      stateColor = Colors.orange;
    } else {
      stateColor = Colors.grey[200];
    }

    return Positioned(
      left: x,
      top: y,
      child: GestureDetector(
          child: Container(
            width: 40,
            height: 40,
            decoration: BoxDecoration(
              color: stateColor,
              border: Border.all(color: stateColor, width: 3),
              borderRadius: BorderRadius.all(Radius.circular(5)),
            ),
          ),
          onTap: () {
            if (this.state) {
              this.state = false;
            } else {
              this.state = true;
            }
            // print("${name}");
            this.callback(this.name);
          }),
    );
  }

  void ClearState() {
    this.state = false;
  }

  void SetState() {
    this.state = true;
  }

  void FlipBlinkState() {
    this.blinkState = !this.blinkState;
  }
}
