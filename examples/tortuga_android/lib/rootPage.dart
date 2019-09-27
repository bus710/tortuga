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
        return InitPage(data: data);
        break;
      case Status.dial:
        return DialPage(data: data);
        break;
      case Status.init:
      default:
        return InitPage(data: data);
    }
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
          child: Container(
            decoration: BoxDecoration(
              color: Theme.of(context).primaryColor,
            ),
            child: SpinKitFoldingCube(
              color: Colors.orange[600],
              size: 80.0,
            ),
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
                    rootModel.pressHandler(Status.dial, address);
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
