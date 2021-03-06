import 'dart:async'; // for streams and timer
import 'package:flutter/material.dart';
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
  StreamSubscription<String> subscription;

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
  Color blinkColor = Colors.black;

  @override
  void initState() {
    // Init the button data list
    // First row
    buttonDataList
        .add(ButtonData("forward/left", 20, 20, false, this.buttonCallback));
    buttonDataList
        .add(ButtonData("forward/none", 110, 20, false, this.buttonCallback));
    buttonDataList
        .add(ButtonData("forward/right", 200, 20, false, this.buttonCallback));
    // Second row
    buttonDataList
        .add(ButtonData("none/left", 20, 110, false, this.buttonCallback));
    buttonDataList
        .add(ButtonData("none/none", 110, 110, true, this.buttonCallback));
    buttonDataList
        .add(ButtonData("none/right", 200, 110, false, this.buttonCallback));
    // Third row
    buttonDataList
        .add(ButtonData("backward/left", 20, 200, false, this.buttonCallback));
    buttonDataList
        .add(ButtonData("backward/none", 110, 200, false, this.buttonCallback));
    buttonDataList.add(
        ButtonData("backward/right", 200, 200, false, this.buttonCallback));

    timer = Timer.periodic(Duration(milliseconds: 500), timerCallback);

    subscription = _bloc.forwardStream.listen((data) {
      socketCallback(data);
    }, onDone: () {
      print("");
    }, onError: (error) {
      print("");
    });
    super.initState();
  }

  @override
  void dispose() {
    subscription.cancel();
    super.dispose();
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
    return buttonDataList[index].getButton();
  }

  void buttonCallback(String pressedButtonName) {
    // To the button indicate the color immediately
    timer.cancel();
    timer = Timer.periodic(Duration(milliseconds: 500), timerCallback);

    // Only the button pressed last time should be set
    // so that this handler clears everyone and sets the last one.
    buttonDataList.forEach((b) => {
          b.clearButtonState(),
          b.setBlinkState(),
          if (b.name == pressedButtonName) {b.setButtonState()}
        });
    _bloc.backwardSink.add(ButtonEvent(pressedButtonName));

    setState(() {});
  }

  void timerCallback(Timer timer) async {
    // To blink the button pressed
    buttonDataList.forEach((b) => {
          b.flipBlinkState(),
        });
    setState(() {});
  }

  void socketCallback(String data) {
    if (data == "active") {
      blinkColor = Colors.orange;
    } else if (data == "inactive") {
      blinkColor = Colors.grey[500];
    }
    buttonDataList.forEach((b) => {
          b.setBlinkColor(blinkColor),
        });
  }
}

class ButtonData {
  String name;
  double x, y;
  bool state;
  bool blinkState;
  Color blinkColor;
  Function callback;

  ButtonData(String name, double x, double y, bool state, Function callback) {
    this.name = name;
    this.x = x;
    this.y = y;
    this.state = state;
    this.blinkState = false;
    this.blinkColor = Colors.black;
    this.callback = callback;
  }

  Widget getButton() {
    Color stateColor = Colors.white;
    if (this.state && this.blinkState) {
      stateColor = blinkColor;
    } else {
      stateColor = Colors.grey[200];
    }

    return Positioned(
      left: x,
      top: y,
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
            debugPrint("In Page: " + this.name);
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

  void clearButtonState() {
    this.state = false;
  }

  void setButtonState() {
    this.state = true;
  }

  void flipBlinkState() {
    this.blinkState = !this.blinkState;
  }

  void setBlinkState() {
    this.blinkState = true;
  }

  void setBlinkColor(Color color) {
    this.blinkColor = color;
  }
}
