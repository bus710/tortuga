import 'package:flutter_web/gestures.dart';
import 'package:flutter_web/material.dart';
import 'package:tortuga_frontend/AppBLoC.dart';
import 'package:tortuga_frontend/AppEvent.dart';
import 'package:tortuga_frontend/main.dart';

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

  // _bloc.backwardSink.add(GestureEvent(0, 0, 0, 0)),
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
                _getFeedback(),
                _getController(),
              ]));
    } else {
      return Text("The space available is too small");
    }
  }

  Widget _getFeedback() {
    return Container(
      width: maxW - 10,
      height: maxH - controllerW - 30,
      margin: EdgeInsets.only(left: 10, top: 20, right: 10, bottom: 10),
      decoration: BoxDecoration(
        color: Color(PRIMARY_COLOR),
        // border: Border.all(color: Colors.black, width: 10),
        // borderRadius: BorderRadius.all(Radius.circular(20)),
      ),
      child: Center(
          child: Text(
        "TORTUGA",
        style: TextStyle(
          fontSize: 48,
          fontWeight: FontWeight.w100,
          color: Colors.white,
        ),
      )),
    );
  }

  Widget _getController() {
    return Container(
      width: controllerW - 30,
      height: controllerW - 30,
      margin: EdgeInsets.all(5),
      decoration: BoxDecoration(
        color: Color(PRIMARY_COLOR),
        border: Border.all(color: Colors.grey[50], width: 3),
        borderRadius: BorderRadius.all(Radius.circular(5)),
      ),
      child: Stack(children: <Widget>[]),
    );
  }
}
